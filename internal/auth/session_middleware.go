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

package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ory/client-go"
)

type tokenRepository interface {
	GetUserIDByToken(tokenStr string) (string, error)
}

func getCookie(name string, cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

func cookieAuth(ctx context.Context, oryApiClient *client.APIClient, oryKratosSessionCookie string) (string, error) {
	// check if we have a session
	session, _, err := oryApiClient.FrontendApi.ToSession(ctx).Cookie(oryKratosSessionCookie).Execute()
	if err != nil {
		return "", err
	}
	return session.Identity.Id, nil
}

func tokenAuth(tokenRepository tokenRepository, header string) (string, error) {
	// get the user id from the database.
	// check if we need to strip a bearer prefix
	if len(header) > 7 && header[:7] == "Bearer " {
		header = header[7:]
	}
	return tokenRepository.GetUserIDByToken(header)
}

func SessionMiddleware(oryApiClient *client.APIClient, tokenRepository tokenRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			oryKratosSessionCookie := getCookie("ory_kratos_session", c.Cookies())

			var userID string

			if oryKratosSessionCookie == nil {
				// check for authorization header
				authorizationHeader := c.Request().Header.Get("Authorization")
				if authorizationHeader == "" {
					return c.JSON(401, map[string]string{"error": "no session, missing authorization header"})
				}
				userID, err = tokenAuth(tokenRepository, authorizationHeader)
			} else {
				userID, err = cookieAuth(c.Request().Context(), oryApiClient, oryKratosSessionCookie.String())
			}

			if err != nil {
				return c.JSON(401, map[string]string{"error": "no session, could not authenticate"})
			}

			c.Set("session", NewSession(userID))
			c.Set("sessionCookie", oryKratosSessionCookie)

			return next(c)
		}
	}
}
