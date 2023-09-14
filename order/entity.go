package order

import "time"

type Base struct {
	OrderId  string    `json:"orderId" gorm:"column:order_id;primary_key"` // 订单id
	Price    float64   `yaml:"price" gorm:"column:price"`                  // 订单价格
	CreateAt time.Time `json:"createAt"  gorm:"column:create_at"`          // 创建时间
	UpdateAt time.Time `json:"updateAt"  gorm:"column:update_at"`          // 更新时间
}
