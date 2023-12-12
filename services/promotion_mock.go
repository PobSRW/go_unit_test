package services

import "github.com/stretchr/testify/mock"

type promotionServiceMock struct {
	mock.Mock
}

func NewPromotionServiceMock() *promotionServiceMock {
	return &promotionServiceMock{}
}

func (s *promotionServiceMock) CalculateDiscount(amount int) (int, error) {
	args := s.Called(amount)
	return args.Int(0), args.Error(1)
}
