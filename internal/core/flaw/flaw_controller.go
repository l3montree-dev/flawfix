package flaw

import (
	"github.com/l3montree-dev/flawfix/internal/core"
	"github.com/labstack/echo/v4"
)

type FlawHttpController struct {
	flawRepository Repository
}

func NewHttpController(flawRepository Repository) *FlawHttpController {
	return &FlawHttpController{
		flawRepository: flawRepository,
	}
}

func (c FlawHttpController) ListPaged(ctx core.Context) error {
	// get the env
	env := core.GetEnv(ctx)

	pagedResp, err := c.flawRepository.GetByEnvIdPaged(
		nil,
		core.GetPageInfo(ctx),
		core.GetFilterQuery(ctx),
		core.GetSortQuery(ctx),
		env.GetID(),
	)

	if err != nil {
		return echo.NewHTTPError(500, "could not get flaws").WithInternal(err)
	}

	return ctx.JSON(200, pagedResp)
}

func (c FlawHttpController) Read(ctx core.Context) error {
	flawId, err := core.GetFlawID(ctx)
	if err != nil {
		return echo.NewHTTPError(400, "invalid flaw id")
	}

	flaw, err := c.flawRepository.Read(flawId)
	if err != nil {
		return echo.NewHTTPError(404, "could not find flaw")
	}

	// get all the associated cwes

	return ctx.JSON(200, flaw)
}
