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

package main

import (
	"github.com/l3montree-dev/flawfix/cmd/flawfix/api"
	"github.com/l3montree-dev/flawfix/internal/core"
	"github.com/l3montree-dev/flawfix/internal/core/config"
	"github.com/l3montree-dev/flawfix/internal/core/leaderelection"
	"github.com/l3montree-dev/flawfix/internal/core/vulndb"

	_ "github.com/lib/pq"
)

//	@title			FlawFix API
//	@version		v1
//	@description	FlawFix API

//	@contact.name	Support
//	@contact.url	https://github.com/l3montree-dev/flawfix/issues

//	@license.name	AGPL-3
//	@license.url	https://github.com/l3montree-dev/flawfix/blob/main/LICENSE.txt

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	core.LoadConfig() // nolint: errcheck
	core.InitLogger()

	db, err := core.DatabaseFactory()

	if err != nil {
		panic(err)
	}

	configService := config.NewService(db)
	leaderElector := leaderelection.NewDatabaseLeaderElector(configService)
	vulndb.StartMirror(db, leaderElector, configService)
	api.Start(db)
}
