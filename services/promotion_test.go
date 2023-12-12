package services_test

import (
	"errors"
	"obp_unit_test/repositories"
	"obp_unit_test/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		Name            string
		PurchaseMin     int
		DiscountPercent int
		Amount          int
		Expect          int
	}

	cases := []testCase{
		{Name: "applied 100", PurchaseMin: 100, DiscountPercent: 20, Amount: 100, Expect: 80},
		{Name: "applied 200", PurchaseMin: 100, DiscountPercent: 20, Amount: 200, Expect: 160},
		{Name: "applied 500", PurchaseMin: 100, DiscountPercent: 20, Amount: 500, Expect: 400},
		{Name: "not applied", PurchaseMin: 100, DiscountPercent: 20, Amount: 80, Expect: 80},
	}

	promoRepo := repositories.NewPromotionRepositoryMock()

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			// Arrage
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.PurchaseMin,
				DiscountPercent: c.DiscountPercent,
			}, nil)

			promoService := services.NewPromotionService(promoRepo)

			// Action
			discount, _ := promoService.CalculateDiscount(c.Amount)
			expected := c.Expect

			// Assert
			assert.Equal(t, expected, discount)
		})
	}

	t.Run("purchase amount zero", func(t *testing.T) {
		// Arrage
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)

		// Action
		_, err := promoService.CalculateDiscount(0)

		// Assert
		assert.ErrorIs(t, err, services.ErrZeroAmount)

		// ทดสอบว่า call func GetPromotion รึเปล่า กันเขียนโค้ดผิด
		// ไปเรียก func GetPromotion ก่อน check error
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repository error", func(t *testing.T) {
		// Arrage
		// ไม่ได้สนใจว่า error อะไร อาจจะ network ขาด, น้ำไม่ไหล,ไฟดับ ก็นับว่าเป็น error
		// จึง config แค่ให้มันเกิด error เฉยๆ
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New(""))

		promoService := services.NewPromotionService(promoRepo)

		// Action
		_, err := promoService.CalculateDiscount(10)

		// Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
