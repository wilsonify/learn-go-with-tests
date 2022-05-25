package handlers

import (
	"S04-go-echo-server/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Strength - summary: signal strength
func (c *Container) Strength(ctx echo.Context) error {
	var s models.StrengthInput
	ctx.Bind(&inputStrengthInput)

	strength_float := float32(s.Actual) / float32(s.Expected)
	outputStrengthOutput := models.StrengthOutput{
		s.Expected,
		s.Actual,
		strength_float,
	}

	fmt.Println("Expected=", s.Expected)
	fmt.Println("Actual=", s.Actual)
	fmt.Println("strength_float=", strength_float)

	return ctx.JSON(http.StatusOK, outputStrengthOutput)
}
