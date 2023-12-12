package handlers_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"obp_unit_test/handlers"
	"obp_unit_test/services"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrage
		amount := 100
		expect := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expect, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		//ทำ http test
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%[1]d", amount), nil)

		// Action
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			// statusOk ค่อยมาอ่านค่าใน body
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expect), string(body))
		}
	})

	t.Run("error status bad request", func(t *testing.T) {

		// Arrage
		amount := "test"
		expect := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expect, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		//ทำ http test
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Action
		res, err := app.Test(req)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})
}
