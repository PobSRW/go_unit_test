package repositories

type PromotionRepository interface {
	GetPromotion() (Promotion, error)
}

// entity เหมือนเป็น table ใน database
type Promotion struct {
	ID              int
	PurchaseMin     int
	DiscountPercent int
}
