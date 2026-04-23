package room

type HotelRoomReq struct {
	Id           int64  `json:"id" form:"id"`
	HotelID      int64  `json:"hotel_id" form:"hotel_id"`
	RoomNo       string `json:"room_no" form:"room_no"`
	RoomTypeName string `json:"room_type_name" form:"room_type_name"`
	RoomTypeCode string `json:"room_type_code" form:"room_type_code"`
	FloorNo      string `json:"floor_no" form:"floor_no"`
	Building     string `json:"building" form:"building"`
	PhoneExt     string `json:"phone_ext" form:"phone_ext"`
	Description  string `json:"description" form:"description"`
}

type SaveRoomTypeReq struct {
	Id           int64  `json:"id" form:"id"`
	HotelID      int64  `json:"hotel_id" form:"hotel_id"`
	TypeName     string `json:"type_name" form:"type_name"`
	TypeCode     string `json:"type_code" form:"type_code"`
	MaxOccupancy int    `json:"max_occupancy" form:"max_occupancy"`
	BasePrice    int64  `json:"base_price" form:"base_price"`
	Quantity     int    `json:"quantity" form:"quantity"`
	Status       int    `json:"status" form:"status"`
}

type RoomGuestStayReq struct {
	Id           int64  `json:"id" form:"id"`
	HotelId      int64  `json:"hotel_id" form:"hotel_id"`
	GuestRoomNo  string `json:"guest_room_no" form:"guest_room_no"`
	GuestName    string `json:"guest_name" form:"guest_name"`
	GuestIdNo    string `json:"guest_id_no" form:"guest_id_no"`
	RealPrice    int64  `json:"real_price" form:"real_price"`
	Mobile       string `json:"mobile" form:"mobile"`
	CheckInTime  string `json:"check_in_time" form:"check_in_time"`
	CheckOutTime string `json:"check_out_time" form:"check_out_time"`
	StayStatus   int8   `json:"stay_status" form:"stay_status"`
	Description  string `json:"description" form:"description"`
}
