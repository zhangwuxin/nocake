package app

type Address struct {
	Id            int    `gorm:"id"`
	AddressName   string `gorm:"address_name"`   // 收货人名称
	OpenId        int    `gorm:"open_id"`        // 用户表的用户ID
	Province      string `gorm:"province"`       // 行政区域表的省ID
	City          string `gorm:"city"`           // 行政区域表的市ID
	County        string `gorm:"county"`         // 行政区域表的区县ID
	AddressDetail string `gorm:"address_detail"` // 详细收货地址
	AreaCode      string `gorm:"area_code"`      // 地区编码
	PostalCode    string `gorm:"postal_code"`    // 邮政编码
	Tel           string `gorm:"tel"`            // 手机号码
	IsDefault     int    `gorm:"is_default"`     // 是否默认地址
	Deleted       int    `gorm:"deleted"`        // 逻辑删除
	CreatedTime   string `gorm:"created_time"`   // 创建时间
	UpdatedTime   string `gorm:"updated_time"`   // 更新时间
	DeletedTime   string `gorm:"deleted_time"`   // 删除时间
}