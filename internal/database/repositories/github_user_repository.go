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

package repositories

import (
	"github.com/google/uuid"
	"github.com/l3montree-dev/devguard/internal/core"
	"github.com/l3montree-dev/devguard/internal/database/models"
)

type githubUserRepository struct {
	db core.DB
	Repository[int, models.GithubUser, core.DB]
}

func NewGithubUserRepository(db core.DB) *githubUserRepository {
	if err := db.AutoMigrate(&models.GithubUser{}); err != nil {
		panic(err)
	}
	return &githubUserRepository{
		db:         db,
		Repository: newGormRepository[int, models.GithubUser](db),
	}
}

func (r *githubUserRepository) FindByOrgID(tx core.DB, orgID uuid.UUID) ([]models.GithubUser, error) {
	var users []models.GithubUser
	if err := r.GetDB(tx).Raw("SELECT gh.* FROM github_users gh WHERE EXISTS(SELECT 1 from github_user_orgs where github_user_id = gh.id AND org_id = ?)", orgID).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
