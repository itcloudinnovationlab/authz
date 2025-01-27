package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/eko/authz/backend/internal/entity/manager"
	"github.com/eko/authz/backend/internal/entity/repository"
	"github.com/eko/authz/backend/internal/http/handler/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Lists actions.
//
//	@security	Authentication
//	@Summary	Lists actions
//	@Tags		Action
//	@Produce	json
//	@Param		page	query		int		false	"page number"			example(1)
//	@Param		size	query		int		false	"page size"				minimum(1)	maximum(1000)	default(100)
//	@Param		filter	query		string	false	"filter on a field"		example(name:contains:something)
//	@Param		sort	query		string	false	"sort field and order"	example(name:desc)
//	@Success	200		{object}	[]model.Action
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/actions [Get]
func ActionList(
	actionManager manager.Action,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page, size, err := paginate(c)
		if err != nil {
			return returnError(c, http.StatusInternalServerError, err)
		}

		// List actions
		action, total, err := actionManager.GetRepository().Find(
			repository.WithPage(page),
			repository.WithSize(size),
			repository.WithFilter(httpFilterToORM(c)),
			repository.WithSort(httpSortToORM(c)),
		)
		if err != nil {
			return returnError(c, http.StatusInternalServerError, err)
		}

		return c.JSON(model.NewPaginated(action, total, page, size))
	}
}

// Retrieve an action.
//
//	@security	Authentication
//	@Summary	Retrieve an action
//	@Tags		Action
//	@Produce	json
//	@Success	200	{object}	model.Action
//	@Failure	404	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/actions/{identifier} [Get]
func ActionGet(
	actionManager manager.Action,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		identifier := c.Params("identifier")

		// Retrieve action
		action, err := actionManager.GetRepository().Get(identifier)
		if err != nil {
			statusCode := http.StatusInternalServerError

			if errors.Is(err, gorm.ErrRecordNotFound) {
				statusCode = http.StatusNotFound
			}

			return returnError(c, statusCode,
				fmt.Errorf("cannot retrieve action: %v", err),
			)
		}

		return c.JSON(action)
	}
}
