package handlers

import (
	"github.com/kelbwah/GoChatter/ui/pages/home"
	"github.com/kelbwah/GoChatter/utils"
	"github.com/labstack/echo/v4"
)

func HandleShowHome(c echo.Context) error {
	return utils.Render(c, home.Show())
}
