package hoteluser

import (
	"time"
)

type HotelUser struct {
	ID           int64      `gorm:"primaryKey;autoIncrement" json:"id"`                                    // 主键ID
	HotelName    string     `gorm:"column:hotel_name;type:varchar(100);not null" json:"hotel_name"`        // 酒店名称
	Username     string     `gorm:"column:username;type:varchar(50);not null;uniqueIndex" json:"username"` // 登录账号
	PasswordHash string     `gorm:"column:password_hash;type:varchar(255);not null" json:"password_hash"`  // 密码哈希
	ContactName  *string    `gorm:"column:contact_name;type:varchar(50)" json:"contact_name,omitempty"`    // 联系人（可空）
	Mobile       *string    `gorm:"column:mobile;type:varchar(20)" json:"mobile,omitempty"`                // 手机号（可空）
	Email        *string    `gorm:"column:email;type:varchar(100)" json:"email,omitempty"`                 // 邮箱（可空）
	Address      *string    `gorm:"column:address;type:varchar(255)" json:"address,omitempty"`             // 酒店地址（可空）
	Status       int8       `gorm:"column:status;type:tinyint;not null;default:1" json:"status"`           // 状态 1启用 0停用
	LastLoginAt  *time.Time `gorm:"column:last_login_at" json:"last_login_at,omitempty"`                   // 最后登录时间（可空）
	LastLoginIp  *string    `gorm:"column:last_login_ip;type:varchar(45)" json:"last_login_ip,omitempty"`  // 最后登录IP（可空）
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`                    // 创建时间
	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`                    // 更新时间
}

func (*HotelUser) TableName() string {
	return "hotel_user"
}
