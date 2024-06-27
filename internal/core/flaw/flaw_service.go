// Copyright (C) 2024 Tim Bastin, l3montree UG (haftungsbeschränkt)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package flaw

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/l3montree-dev/flawfix/internal/core"
	"github.com/l3montree-dev/flawfix/internal/core/risk"
	"github.com/l3montree-dev/flawfix/internal/database"
	"github.com/l3montree-dev/flawfix/internal/database/models"
)

type assetRepository interface {
	Update(tx core.DB, asset *models.Asset) error
}

type flawRepository interface {
	SaveBatch(db core.DB, flaws []models.Flaw) error
	Save(db core.DB, flaws *models.Flaw) error
	Transaction(txFunc func(core.DB) error) error
	GetAllFlawsByAssetID(tx core.DB, assetID uuid.UUID) ([]models.Flaw, error)
}

type flawEventRepository interface {
	SaveBatch(db core.DB, events []models.FlawEvent) error
	Save(db core.DB, event *models.FlawEvent) error
}
type cveRepository interface {
	FindCVE(tx database.DB, cveId string) (any, error)
}
type service struct {
	flawRepository      flawRepository
	flawEventRepository flawEventRepository
	assetRepository     assetRepository
	cveRepository       cveRepository
}

func NewService(flawRepository flawRepository, flawEventRepository flawEventRepository, assetRepository assetRepository, cveRepository cveRepository) *service {
	return &service{
		flawRepository:      flawRepository,
		flawEventRepository: flawEventRepository,
		assetRepository:     assetRepository,
		cveRepository:       cveRepository,
	}
}

// expect a transaction to be passed
func (s *service) UserFixedFlaws(tx core.DB, userID string, flaws []models.Flaw) error {
	if len(flaws) == 0 {
		return nil
	}
	// create a new flawevent for each fixed flaw
	events := make([]models.FlawEvent, len(flaws))
	for i, flaw := range flaws {
		ev := models.NewFixedEvent(flaw.CalculateHash(), userID)
		// apply the event on the flaw
		ev.Apply(&flaws[i])
		events[i] = ev
	}

	err := s.flawRepository.SaveBatch(tx, flaws)
	if err != nil {
		return err
	}
	return s.flawEventRepository.SaveBatch(tx, events)
}

// expect a transaction to be passed
func (s *service) UserDetectedFlaws(tx core.DB, userID string, flaws []models.Flaw, asset models.Asset) error {
	if len(flaws) == 0 {
		return nil
	}
	// create a new flawevent for each detected flaw
	events := make([]models.FlawEvent, len(flaws))
	for i, flaw := range flaws {
		ev := models.NewDetectedEvent(flaw.CalculateHash(), userID)
		// apply the event on the flaw
		ev.Apply(&flaws[i])
		events[i] = ev

		e := core.Environmental{
			ConfidentialityRequirements: string(asset.ConfidentialityRequirement),
			IntegrityRequirements:       string(asset.IntegrityRequirement),
			AvailabilityRequirements:    string(asset.AvailabilityRequirement),
		}

		flaws[i].RawRiskAssessment = risk.RawRisk(*flaw.CVE, e)
	}

	// run the updates in the transaction to keep a valid state
	err := s.flawRepository.SaveBatch(tx, flaws)
	if err != nil {
		return err
	}
	return s.flawEventRepository.SaveBatch(tx, events)
}

func (s *service) RecalculateRawRiskAssessment(tx core.DB, userID string, flaws []models.Flaw, justification string, asset models.Asset) error {

	/*
		if len(flaws) == 0 {
			return fmt.Errorf("no flaws to update")
		}
	*/
	env := core.Environmental{
		ConfidentialityRequirements: string(asset.ConfidentialityRequirement),
		IntegrityRequirements:       string(asset.IntegrityRequirement),
		AvailabilityRequirements:    string(asset.AvailabilityRequirement),
	}

	// create a new flawevent for each updated flaw
	events := make([]models.FlawEvent, len(flaws))
	for i, flaw := range flaws {
		cviID := flaw.CVEID
		cve, err := s.cveRepository.FindCVE(nil, cviID)
		if err != nil {
			slog.Info("Error getting CVE: %v", err)
			continue
		}

		cve2 := cve.(models.CVE)
		oldRiskAssessment := flaw.RawRiskAssessment
		newRiskAssessment := risk.RawRisk(cve2, env)

		ev := models.NewRawRiskAssessmentUpdatedEvent(flaw.CalculateHash(), userID, justification, *oldRiskAssessment, *newRiskAssessment)
		// apply the event on the flaw
		ev.Apply(&flaws[i])
		events[i] = ev
	}

	err := s.flawRepository.SaveBatch(tx, flaws)
	if err != nil {
		return fmt.Errorf("could not save flaws: %v", err)
	}

	err = s.flawEventRepository.SaveBatch(tx, events)
	if err != nil {
		return fmt.Errorf("could not save events: %v", err)

	}
	return nil
}

func (s *service) UpdateFlawState(tx core.DB, userID string, flaw *models.Flaw, statusType string, justification *string) error {
	if tx == nil {
		// we are not part of a parent transaction - create a new one
		return s.flawRepository.Transaction(func(d core.DB) error {
			return s.updateFlawState(d, userID, flaw, statusType, justification)
		})
	}
	return s.updateFlawState(tx, userID, flaw, statusType, justification)
}

func (s *service) updateFlawState(tx core.DB, userID string, flaw *models.Flaw, statusType string, justification *string) error {
	ev := models.FlawEvent{
		Type:          models.FlawEventType(statusType),
		FlawID:        flaw.CalculateHash(),
		UserID:        userID,
		Justification: justification,
	}
	// apply the event on the flaw
	ev.Apply(flaw)

	// run the updates in the transaction to keep a valid state
	err := s.flawRepository.Save(tx, flaw)
	if err != nil {
		return err
	}
	if err := s.flawEventRepository.Save(tx, &ev); err != nil {
		return err
	}
	flaw.Events = append(flaw.Events, ev)
	return nil
}
