// Copyright (C) 2023 Tim Bastin, l3montree UG (haftungsbeschränkt)
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
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package repositories

import (
	"github.com/google/uuid"
	"github.com/l3montree-dev/flawfix/internal/models"
	"gorm.io/gorm"
)

type GormOrganizationRepository struct {
	db *gorm.DB
	GormRepository[uuid.UUID, models.Organization]
}

func NewGormOrganizationRepository(db *gorm.DB) *GormOrganizationRepository {
	return &GormOrganizationRepository{
		db: db,
	}
}

func (o *GormOrganizationRepository) Create(org *models.Organization) error {
	return o.db.Create(org).Error
}

func (o *GormOrganizationRepository) Delete(id uuid.UUID) error {
	return o.db.Delete(&models.Organization{}, id).Error
}

func (o *GormOrganizationRepository) Read(id uuid.UUID) (models.Organization, error) {
	var org models.Organization
	err := o.db.First(&org, id).Error
	if err != nil {
		return org, err
	}
	return org, nil
}

func (o *GormOrganizationRepository) Update(org *models.Organization) error {
	return o.db.Save(org).Error
}
