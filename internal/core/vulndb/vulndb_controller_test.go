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

package vulndb

import (
	"math"
	"testing"

	"github.com/l3montree-dev/flawfix/internal/core"
	"github.com/l3montree-dev/flawfix/internal/database/models"
	"github.com/l3montree-dev/flawfix/internal/obj"
)

type tableTest struct {
	vector             string
	metrics            obj.RiskMetrics
	env                core.Environmental
	exploits           []*models.Exploit
	expectedVector     string
	cvss               float32
	affectedComponents []models.AffectedComponent
}

func ptr[T any](s T) *T {
	return &s
}

func TestCalculateRisk(t *testing.T) {
	t.Run("should not panic if no vector is defined", func(t *testing.T) {
		sut := models.CVE{
			CVSS:   5,
			Vector: "",
		}
		env := core.Environmental{}
		riskMetrics, vector := RiskCalculation(sut, env)

		if riskMetrics.BaseScore != 0 {
			t.Errorf("Expected base score to be 5, got %f", riskMetrics.BaseScore)
		}

		if riskMetrics.WithEnvironment != obj.CannotCalculateRisk {
			t.Errorf("Expected with environment score to be %f, got %f", obj.CannotCalculateRisk, riskMetrics.WithEnvironment)
		}

		if riskMetrics.WithThreatIntelligence != obj.CannotCalculateRisk {
			t.Errorf("Expected with threat intelligence score to be %f, got %f", obj.CannotCalculateRisk, riskMetrics.WithThreatIntelligence)
		}

		if riskMetrics.WithEnvironmentAndThreatIntelligence != obj.CannotCalculateRisk {
			t.Errorf("Expected with environment and threat intelligence score to be %f, got %f", obj.CannotCalculateRisk, riskMetrics.WithEnvironmentAndThreatIntelligence)
		}

		if vector != "" {
			t.Errorf("Expected vector to be empty, got %s", vector)
		}
	})

	table := []tableTest{
		{
			vector: "AV:L/AC:H/Au:M/C:C/I:C/A:C",
			metrics: obj.RiskMetrics{
				BaseScore:                            5.9,
				WithEnvironment:                      5.9,
				WithThreatIntelligence:               5.0,
				WithEnvironmentAndThreatIntelligence: 5.0,
			},
			env:            core.Environmental{},
			expectedVector: "AV:L/AC:H/Au:M/C:C/I:C/A:C/E:U/RL:ND/RC:C",
			cvss:           5.9,
		},
		{
			vector: "AV:L/AC:H/Au:M/C:C/I:C/A:C",
			metrics: obj.RiskMetrics{
				BaseScore:                            5.9,
				WithEnvironment:                      4.0,
				WithThreatIntelligence:               5.0,
				WithEnvironmentAndThreatIntelligence: 3.4,
			},
			env: core.Environmental{
				ConfidentialityRequirements: "L",
				IntegrityRequirements:       "L",
				AvailabilityRequirements:    "L",
			},
			expectedVector: "AV:L/AC:H/Au:M/C:C/I:C/A:C/E:U/RL:ND/RC:C/CDP:ND/TD:ND/CR:L/IR:L/AR:L",
			cvss:           5.9,
		},
		{
			vector: "AV:L/AC:H/Au:M/C:C/I:C/A:C",
			metrics: obj.RiskMetrics{
				BaseScore:                            5.9,
				WithEnvironment:                      4.0,
				WithThreatIntelligence:               5.6,
				WithEnvironmentAndThreatIntelligence: 3.8,
			},
			env: core.Environmental{
				ConfidentialityRequirements: "L",
				IntegrityRequirements:       "L",
				AvailabilityRequirements:    "L",
			},
			exploits: []*models.Exploit{
				{
					Verified: true,
				},
			},
			expectedVector: "AV:L/AC:H/Au:M/C:C/I:C/A:C/E:F/RL:ND/RC:C/CDP:ND/TD:ND/CR:L/IR:L/AR:L",
			cvss:           5.9,
		},
		{
			vector: "AV:L/AC:H/Au:M/C:C/I:C/A:C",
			metrics: obj.RiskMetrics{
				BaseScore:                            5.9,
				WithEnvironment:                      4.0,
				WithThreatIntelligence:               4.9,
				WithEnvironmentAndThreatIntelligence: 3.3,
			},
			env: core.Environmental{
				ConfidentialityRequirements: "L",
				IntegrityRequirements:       "L",
				AvailabilityRequirements:    "L",
			},
			exploits: []*models.Exploit{
				{
					Verified: true,
				},
			},
			expectedVector: "AV:L/AC:H/Au:M/C:C/I:C/A:C/E:F/RL:OF/RC:C/CDP:ND/TD:ND/CR:L/IR:L/AR:L",
			cvss:           5.9,
			affectedComponents: []models.AffectedComponent{{
				SemverFixed: ptr("v1.0.0"),
			}},
		},
		{
			vector: "AV:L/AC:H/Au:M/C:C/I:C/A:C",
			metrics: obj.RiskMetrics{
				BaseScore:                            5.9,
				WithEnvironment:                      4.0,
				WithThreatIntelligence:               5.6,
				WithEnvironmentAndThreatIntelligence: 3.8,
			},
			env: core.Environmental{
				ConfidentialityRequirements: "L",
				IntegrityRequirements:       "L",
				AvailabilityRequirements:    "L",
			},
			exploits: []*models.Exploit{
				{
					Verified: true,
				},
			},
			expectedVector: "AV:L/AC:H/Au:M/C:C/I:C/A:C/E:F/RL:U/RC:C/CDP:ND/TD:ND/CR:L/IR:L/AR:L",
			cvss:           5.9,
			affectedComponents: []models.AffectedComponent{{
				SemverFixed: nil,
			}},
		},
		{
			vector: "CVSS:3.0/AV:N/AC:H/PR:L/UI:R/S:U/C:N/I:N/A:L",
			metrics: obj.RiskMetrics{
				BaseScore:                            2.6,
				WithEnvironment:                      2.6,
				WithThreatIntelligence:               2.4,
				WithEnvironmentAndThreatIntelligence: 2.4,
			},
			env:            core.Environmental{},
			expectedVector: "CVSS:3.0/AV:N/AC:H/PR:L/UI:R/S:U/C:N/I:N/A:L/E:U/RC:C",
			cvss:           2.6,
		},
		{
			vector: "CVSS:3.1/AV:N/AC:H/PR:L/UI:R/S:U/C:N/I:N/A:L",
			metrics: obj.RiskMetrics{
				BaseScore:                            2.6,
				WithEnvironment:                      1.9,
				WithThreatIntelligence:               2.4,
				WithEnvironmentAndThreatIntelligence: 1.8,
			},
			env: core.Environmental{
				IntegrityRequirements:       "L",
				ConfidentialityRequirements: "L",
				AvailabilityRequirements:    "L",
			},
			expectedVector: "CVSS:3.1/AV:N/AC:H/PR:L/UI:R/S:U/C:N/I:N/A:L/E:U/RC:C/CR:L/IR:L/AR:L",
			cvss:           2.6,
		},
		{
			vector: "CVSS:3.1/AV:N/AC:H/PR:L/UI:R/S:U/C:N/I:N/A:L",
			metrics: obj.RiskMetrics{
				BaseScore:                            2.6,
				WithEnvironment:                      3.4,
				WithThreatIntelligence:               2.4,
				WithEnvironmentAndThreatIntelligence: 3.2,
			},
			env: core.Environmental{
				IntegrityRequirements:       "H",
				ConfidentialityRequirements: "H",
				AvailabilityRequirements:    "H",
			},
			exploits: []*models.Exploit{
				{
					Verified: true,
				},
			},
			affectedComponents: []models.AffectedComponent{{
				SemverFixed: ptr("v1.0.0"),
			}},
			expectedVector: "CVSS:3.1/AV:N/AC:H/PR:L/UI:R/S:U/C:N/I:N/A:L/E:F/RL:O/RC:C/CR:H/IR:H/AR:H",
			cvss:           2.6,
		},
		{
			vector: "CVSS:3.1/AV:N/AC:H/PR:L/UI:R/S:U/C:N/I:N/A:L",
			metrics: obj.RiskMetrics{
				BaseScore:                            2.6,
				WithEnvironment:                      3.4,
				WithThreatIntelligence:               2.6,
				WithEnvironmentAndThreatIntelligence: 3.3,
			},
			env: core.Environmental{
				IntegrityRequirements:       "H",
				ConfidentialityRequirements: "H",
				AvailabilityRequirements:    "H",
			},
			exploits: []*models.Exploit{
				{
					Verified: true,
				},
			},
			affectedComponents: []models.AffectedComponent{{
				SemverFixed: nil,
			}},
			expectedVector: "CVSS:3.1/AV:N/AC:H/PR:L/UI:R/S:U/C:N/I:N/A:L/E:F/RL:U/RC:C/CR:H/IR:H/AR:H",
			cvss:           2.6,
		},
		{
			vector: "CVSS:4.0/AV:A/AC:H/AT:P/PR:L/UI:N/VC:H/VI:H/VA:H/SC:N/SI:N/SA:N",
			metrics: obj.RiskMetrics{
				BaseScore:                            7.5,
				WithEnvironment:                      6.2,
				WithThreatIntelligence:               4.8,
				WithEnvironmentAndThreatIntelligence: 2.4,
			},
			env: core.Environmental{
				IntegrityRequirements:       "L",
				ConfidentialityRequirements: "L",
				AvailabilityRequirements:    "L",
			},
			cvss:           7.5,
			expectedVector: "CVSS:4.0/AV:A/AC:H/AT:P/PR:L/UI:N/VC:H/VI:H/VA:H/SC:N/SI:N/SA:N/E:U/CR:L/IR:L/AR:L",
		},
		{
			vector: "CVSS:4.0/AV:A/AC:H/AT:P/PR:L/UI:N/VC:H/VI:H/VA:H/SC:N/SI:N/SA:N",
			metrics: obj.RiskMetrics{
				BaseScore:                            7.5,
				WithEnvironment:                      6.2,
				WithThreatIntelligence:               6.6,
				WithEnvironmentAndThreatIntelligence: 5.1,
			},
			env: core.Environmental{
				IntegrityRequirements:       "L",
				ConfidentialityRequirements: "L",
				AvailabilityRequirements:    "L",
			},
			exploits:       []*models.Exploit{{Verified: true}},
			cvss:           7.5,
			expectedVector: "CVSS:4.0/AV:A/AC:H/AT:P/PR:L/UI:N/VC:H/VI:H/VA:H/SC:N/SI:N/SA:N/E:P/CR:L/IR:L/AR:L",
		},
		{
			vector: "CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N",
			metrics: obj.RiskMetrics{
				BaseScore:                            6.1,
				WithEnvironment:                      6.8,
				WithThreatIntelligence:               6.0,
				WithEnvironmentAndThreatIntelligence: 6.6,
			},
			env: core.Environmental{
				IntegrityRequirements:       "M",
				ConfidentialityRequirements: "H",
				AvailabilityRequirements:    "M",
			},
			expectedVector: "CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N/E:F/RC:C/CR:H/IR:M/AR:M",
			cvss:           6.1,
			exploits:       []*models.Exploit{{Verified: true}},
		},
	}

	for _, tableTest := range table {
		vector := tableTest.vector
		t.Run("should return same values, if no env metrics and threat metrics are defined. Vector: "+vector, func(t *testing.T) {
			sut := models.CVE{
				CVSS:               tableTest.cvss,
				Vector:             vector,
				Exploits:           tableTest.exploits,
				AffectedComponents: tableTest.affectedComponents,
			}
			env := tableTest.env
			expectedRiskMetrics := tableTest.metrics
			riskMetrics, vector := RiskCalculation(sut, env)

			if !floatsEqual(riskMetrics.BaseScore, expectedRiskMetrics.BaseScore) {
				t.Errorf("Expected base score to be %f, got %f", expectedRiskMetrics.BaseScore, riskMetrics.BaseScore)
			}

			if !floatsEqual(riskMetrics.WithEnvironment, expectedRiskMetrics.WithEnvironment) {
				t.Errorf("Expected with environment score to be %f, got %f", expectedRiskMetrics.WithEnvironment, riskMetrics.WithEnvironment)
			}

			if !floatsEqual(riskMetrics.WithThreatIntelligence, expectedRiskMetrics.WithThreatIntelligence) {
				t.Errorf("Expected with threat intelligence score to be %f, got %f", expectedRiskMetrics.WithThreatIntelligence, riskMetrics.WithThreatIntelligence)
			}

			if !floatsEqual(riskMetrics.WithEnvironmentAndThreatIntelligence, expectedRiskMetrics.WithEnvironmentAndThreatIntelligence) {
				t.Errorf("Expected with environment and threat intelligence score to be %f, got %f", expectedRiskMetrics.WithEnvironmentAndThreatIntelligence, riskMetrics.WithEnvironmentAndThreatIntelligence)
			}

			if vector != tableTest.expectedVector {
				t.Errorf("Expected vector to be %s, got %s", tableTest.expectedVector, vector)
			}
		})
	}
}

func floatsEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.01
}
