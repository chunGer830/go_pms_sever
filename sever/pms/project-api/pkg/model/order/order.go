package order

import "time"

type OrdInf struct {
	Id           int64     `json:"id" form:"id"`
	HotelId      int64     `json:"hotel_id" form:"hotel_id"`
	GuestRoomNo  string    `json:"guest_room_no" form:"guest_room_no"`
	GuestName    string    `json:"guest_name" form:"guest_name"`
	GuestIdNo    string    `json:"guest_id_no" form:"guest_id_no"`
	RealPrice    uint64    `json:"real_price" form:"real_price"` // int unsigned → uint64
	Mobile       string    `json:"mobile" form:"mobile"`
	CheckInTime  time.Time `json:"check_in_time" form:"check_in_time"`
	CheckOutTime time.Time `json:"check_out_time,omitempty" form:"check_out_time"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
}
