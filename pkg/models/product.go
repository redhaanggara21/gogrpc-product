package models

type Product struct {
	Id                int64            `json:"id" gorm:"primaryKey"`
	Name              string           `json:"name"`
	SKU               string           `json:"sku"`
	Stock             int64            `json:"stock"`
	Price             int64            `json:"price"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}
