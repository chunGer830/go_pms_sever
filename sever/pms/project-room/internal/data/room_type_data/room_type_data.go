package room_type_data

import "time"

type RoomType struct {
	ID             int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                     // 主键ID
	HotelID        int64     `gorm:"column:hotel_id;not null" json:"hotel_id"`                         // 酒店ID
	TypeName       string    `gorm:"column:type_name;type:varchar(100);not null" json:"type_name"`     // 房型名称
	TypeCode       string    `gorm:"column:type_code;type:varchar(50);not null" json:"type_code"`      // 房型编码
	MaxOccupancy   int       `gorm:"column:max_occupancy;not null;default:2" json:"max_occupancy"`     // 最大入住人数
	BedType        *string   `gorm:"column:bed_type;type:varchar(50)" json:"bed_type"`                 // 床型
	BasePrice      int64     `gorm:"column:base_price;not null;default:0" json:"base_price"`           // 基础价
	BreakfastCount int       `gorm:"column:breakfast_count;not null;default:0" json:"breakfast_count"` // 早餐份数
	HasWindow      int       `gorm:"column:has_window;not null;default:1" json:"has_window"`           // 是否有窗 1是0否
	Quantity       int       `gorm:"column:quantity;default:1" json:"quantity"`                        // 数量
	Description    *string   `gorm:"column:description;type:varchar(500)" json:"description"`          // 描述
	Status         int       `gorm:"column:status;not null;default:1" json:"status"`                   // 1启用 0停用 2维护
	IsDeleted      int       `gorm:"column:is_deleted;not null;default:0" json:"is_deleted"`           // 逻辑删除
	CreatedAt      time.Time `gorm:"column:created_at;not null;autoCreateTime" json:"created_at"`      // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;autoUpdateTime" json:"updated_at"`      // 更新时间
}

func (RoomType) TableName() string {
	return "room_type"
}

type HotelRoom struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	HotelID      int64     `gorm:"not null;column:hotel_id;comment:酒店ID" json:"hotel_id"`
	RoomNo       string    `gorm:"not null;column:room_no;comment:房号" json:"room_no"`
	RoomTypeCode string    `gorm:"not null;column:room_type_code;comment:房型编号" json:"room_type_code"`
	RoomTypeName string    `gorm:"column:room_type_name;type:varchar(100);not null" json:"room_type_name"` // 房型名称
	FloorNo      *string   `gorm:"not null;column:floor_no;comment:楼层，如18F" json:"floor_no"`
	Building     *string   `gorm:"column:building;comment:楼栋/区域" json:"building"`    // 可为空
	PhoneExt     *string   `gorm:"column:phone_ext;comment:分机号" json:"phone_ext"`    // 可为空
	Description  *string   `gorm:"column:description;comment:备注" json:"description"` // 可为空
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP;not null;column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP;not null;column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (HotelRoom) TableName() string {
	return "hotel_room"
}

// RoomGuestStay 房间入住表
type RoomGuestStay struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	HotelID      int64     `gorm:"column:hotel_id;not null" json:"hotel_id"`
	GuestRoomNo  string    `gorm:"column:guest_room_no;not null" json:"guest_room_no"`
	GuestName    string    `gorm:"column:guest_name;not null;comment:住客姓名" json:"guest_name"`
	GuestIDNo    string    `gorm:"column:guest_id_no" json:"guest_id_no,omitempty"` // 可空
	RealPrice    int64     `gorm:"column:real_price;not null;default:0;comment:实际价格" json:"real_price"`
	Mobile       string    `gorm:"column:mobile" json:"mobile,omitempty"` // 可空
	CheckInTime  string    `gorm:"column:check_in_time;not null;comment:入住时间" json:"check_in_time"`
	CheckOutTime string    `gorm:"column:check_out_time" json:"check_out_time,omitempty"` // 可空
	StayStatus   int8      `gorm:"column:stay_status;not null;default:1;comment:1空置 2在住 3待清理" json:"stay_status"`
	Description  string    `gorm:"column:description" json:"description,omitempty"` // 可空
	CreatedAt    time.Time `gorm:"column:created_at;not null;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;autoUpdateTime" json:"updated_at"`
}

func (RoomGuestStay) TableName() string {
	return "room_guest_stay"
}
