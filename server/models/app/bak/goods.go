package app

import "time"

type Goods struct {
	Id          int       `gorm:"id"`
	Name        string    `gorm:"name"`        // 商品名称
	Brief       string    `gorm:"brief"`       // 商品简介
	CategoryId  int       `gorm:"category_id"` // 商品所属类目ID
	Keywords    string    `gorm:"keywords"`    // 商品关键字，采用逗号间隔
	Status      int       `gorm:"status"`      // 状态 1未上架 2已上架
	IsHot       int       `gorm:"is_hot"`      // 状态 1不推荐 2推荐
	Weight      int       `gorm:"weight"`
	PicUrl      string    `gorm:"pic_url"`      // 商品主图
	PicUrls     string    `gorm:"pic_urls"`     // 商品附图JSON数组格式
	Unit        string    `gorm:"unit"`         // 商品单位，例如件、盒
	Quantity    int       `gorm:"quantity"`     // 商品库存
	Price       float64   `gorm:"price"`        // 价格
	Deleted     int       `gorm:"deleted"`      // 逻辑删除
	CreatedTime time.Time `gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time"` // 更新时间
	DeletedTime time.Time `gorm:"deleted_time"` // 删除时间
}
