package handlers
import (
    "go-echo-server/models"
    "github.com/labstack/echo/v4"
    "net/http"
)

// Strength - summary: signal strength
func (c *Container) Strength(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.HelloWorld {
        Message: "Hello World",
    })
}
