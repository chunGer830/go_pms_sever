package order_inf_data

import "time"

type OrderInfo struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                                                        // 主键ID
	HotelId      int64     `gorm:"column:hotel_id;not null" json:"hotel_id"`                                                            // 入住酒店id
	GuestRoomNo  string    `gorm:"column:guest_room_no;type:varchar(20);not null" json:"guest_room_no"`                                 // 入住房间号
	GuestName    string    `gorm:"column:guest_name;type:varchar(30);not null" json:"guest_name"`                                       // 住客姓名
	GuestIdNo    string    `gorm:"column:guest_id_no;type:varchar(20);not null;default:''" json:"guest_id_no"`                          // 证件号
	RealPrice    uint      `gorm:"column:real_price;type:int unsigned;not null;default:0" json:"real_price"`                            // 实际价格 (int unsigned → uint)
	Mobile       string    `gorm:"column:mobile;type:varchar(20);default:''" json:"mobile"`                                             // 手机号
	CheckInTime  time.Time `gorm:"column:check_in_time;type:datetime;not null" json:"check_in_time"`                                    // 入住时间
	CheckOutTime time.Time `gorm:"column:check_out_time;type:datetime;default:not null" json:"check_out_time,omitempty"`                // 退房时间
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"` // 更新时间
}

func (OrderInfo) TableName() string {
	return "order_info"
}
