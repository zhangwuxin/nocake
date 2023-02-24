package app

type Order struct {
	Id            int     `gorm:"id"`
	OpenId        int     `gorm:"open_id"`        // 用户表的用户ID
	OrderSn       string  `gorm:"order_sn"`       // 订单编号
	OrderStatus   int     `gorm:"order_status"`   // 订单状态
	Consignee     string  `gorm:"consignee"`      // 收货人名称
	Mobile        string  `gorm:"mobile"`         // 收货人手机号
	Address       string  `gorm:"address"`        // 收货具体地址
	Message       string  `gorm:"message"`        // 用户订单留言
	GoodsPrice    float64 `gorm:"goods_price"`    // 商品总费用
	Freighprice   float64 `gorm:"freight_price"`  // 配送费用
	CouponPrice   float64 `gorm:"coupon_price"`   // 优惠券减免
	IntegralPrice float64 `gorm:"integral_price"` // 用户积分减免
	GrouponPrice  float64 `gorm:"groupon_price"`  // 团购优惠价减免
	OrderPrice    float64 `gorm:"order_price"`    // 订单费用， = goods_price + freight_price - coupon_price
	ActualPrice   float64 `gorm:"actual_price"`   // 实付费用， = order_price - integral_price
	PayId         string  `gorm:"pay_id"`         // 微信付款编号
	PayTime       string  `gorm:"pay_time"`       // 微信付款时间
	ShipSn        string  `gorm:"ship_sn"`        // 发货编号
	ShipChannel   string  `gorm:"ship_channel"`   // 发货快递公司
	ShipTime      string  `gorm:"ship_time"`      // 发货开始时间
	ConfirmTime   string  `gorm:"confirm_time"`   // 用户确认收货时间
	Comments      int     `gorm:"comments"`       // 待评价订单商品数量
	EndTime       string  `gorm:"end_time"`       // 订单关闭时间
	Deleted       int     `gorm:"deleted"`        // 逻辑删除
	CreatedTime   string  `gorm:"created_time"`   // 创建时间
	UpdatedTime   string  `gorm:"updated_time"`   // 更新时间
	DeletedTime   string  `gorm:"deleted_time"`   // 删除时间
}