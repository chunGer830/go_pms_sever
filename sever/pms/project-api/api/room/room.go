package room

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	roomModel "pms.com/project-api/pkg/model/room"
	common "pms.com/project-common"
	"pms.com/project-common/errs"
	"pms.com/project-grpc/room/room_type"
	"time"
)

type HandlerRoom struct {
}

func New() *HandlerRoom {
	return &HandlerRoom{}
}

func (r *HandlerRoom) roomType(c *gin.Context) {
	result := &common.Result{}

	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}

	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.RoomTypeMessage{
		HotelId: id,
	}

	roomRsp, err := RoomServiceClient.RoomType(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess(roomRsp))
}

func (r *HandlerRoom) saveRoomType(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.SaveRoomTypeReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.SaveRoomTypeMessage{
		HotelId:      id,
		TypeName:     req.TypeName,
		TypeCode:     req.TypeCode,
		MaxOccupancy: int32(req.MaxOccupancy),
		BasePrice:    req.BasePrice,
		Quantity:     int32(req.Quantity),
		Status:       int32(req.Status),
	}

	_, err = RoomServiceClient.SaveRoomType(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("添加成功"))
}

func (r *HandlerRoom) updateRoomType(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.SaveRoomTypeReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println(req)
	msg := &room_type.UpdateRoomTypeMessage{
		Id:           req.Id,
		HotelId:      id,
		TypeName:     req.TypeName,
		TypeCode:     req.TypeCode,
		MaxOccupancy: int32(req.MaxOccupancy),
		BasePrice:    req.BasePrice,
		Quantity:     int32(req.Quantity),
		Status:       int32(req.Status),
	}
	fmt.Println(msg)
	_, err = RoomServiceClient.UpdateRoomType(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("修改成功"))
}

func (r *HandlerRoom) deleteRoomType(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.SaveRoomTypeReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.DeleteRoomTypeMessage{
		HotelId:  id,
		TypeCode: req.TypeCode,
	}

	_, err = RoomServiceClient.DeleteRoomType(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("删除成功"))
}

func (r *HandlerRoom) hotelRoom(c *gin.Context) {
	result := &common.Result{}

	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.HotelRoomMessage{
		HotelId: id,
	}

	roomRsp, err := RoomServiceClient.HotelRoom(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess(roomRsp))
}

func (r *HandlerRoom) saveHotelRoom(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.HotelRoomReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.SaveHotelRoomMessage{
		HotelId:      id,
		RoomNo:       req.RoomNo,
		RoomTypeName: req.RoomTypeName,
		RoomTypeCode: req.RoomTypeCode,
		FloorNo:      req.FloorNo,
		PhoneExt:     req.PhoneExt,
		Description:  req.Description,
	}

	_, err = RoomServiceClient.SaveHotelRoom(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("添加成功"))
}

func (r *HandlerRoom) updateHotelRoom(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.HotelRoomReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.UpdateHotelRoomMessage{
		Id:           req.Id,
		HotelId:      id,
		RoomNo:       req.RoomNo,
		RoomTypeName: req.RoomTypeName,
		RoomTypeCode: req.RoomTypeCode,
		FloorNo:      req.FloorNo,
		PhoneExt:     req.PhoneExt,
		Description:  req.Description,
	}

	_, err = RoomServiceClient.UpdateHotelRoom(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("修改成功"))
}

func (r *HandlerRoom) deleteHotelRoom(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.HotelRoomReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.DeleteHotelRoomMessage{
		HotelId: id,
		RoomNo:  req.RoomNo,
	}

	_, err = RoomServiceClient.DeleteHotelRoom(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("删除成功"))
}

func (r *HandlerRoom) roomGuestStay(c *gin.Context) {
	result := &common.Result{}

	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.RoomGuestStayMessage{
		HotelId: id,
	}

	roomRsp, err := RoomServiceClient.RoomGuestStay(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess(roomRsp))
}

func (r *HandlerRoom) updateRoomGuestStay(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.RoomGuestStayReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.UpdateRoomGuestStayMessage{
		Id:           req.Id,
		HotelId:      id,
		GuestRoomNo:  req.GuestRoomNo,
		GuestName:    req.GuestName,
		GuestIdNo:    req.GuestIdNo,
		RealPrice:    req.RealPrice,
		Mobile:       req.Mobile,
		CheckInTime:  req.CheckInTime,
		CheckOutTime: req.CheckOutTime,
		StayStatus:   2, //1空置 2在住 3待清理 4禁用
		Description:  req.Description,
	}

	_, err = RoomServiceClient.UpdateRoomGuestStay(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("入住成功"))
}

func (r *HandlerRoom) checkoutRoomGuestStay(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.RoomGuestStayReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.UpdateRoomGuestStayMessage{
		Id:           req.Id,
		HotelId:      id,
		GuestRoomNo:  req.GuestRoomNo,
		GuestName:    req.GuestName,
		GuestIdNo:    req.GuestIdNo,
		RealPrice:    req.RealPrice,
		Mobile:       req.Mobile,
		CheckInTime:  req.CheckInTime,
		CheckOutTime: req.CheckOutTime,
		StayStatus:   3, //1空置 2在住 3待清理 4禁用
		Description:  req.Description,
	}

	_, err = RoomServiceClient.UpdateRoomGuestStay(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("退房成功"))
}

func (r *HandlerRoom) cleanRoomGuestStay(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.RoomGuestStayReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.UpdateRoomGuestStayMessage{
		Id:           req.Id,
		HotelId:      id,
		GuestRoomNo:  req.GuestRoomNo,
		GuestName:    "",
		GuestIdNo:    "",
		RealPrice:    0,
		Mobile:       "",
		CheckInTime:  "",
		CheckOutTime: "",
		StayStatus:   1,
		Description:  "",
	}

	_, err = RoomServiceClient.UpdateRoomGuestStay(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("清理成功"))
}

func (r *HandlerRoom) disableRoomGuestStay(c *gin.Context) {
	result := &common.Result{}
	//接收参数
	var req roomModel.RoomGuestStayReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	//获取id
	idValue, exists := c.Get("hotel_id")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "未登录"))
		return
	}
	id, ok := idValue.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, result.Fail(http.StatusBadRequest, "用户ID类型错误"))
		return
	}

	//
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	msg := &room_type.UpdateRoomGuestStayMessage{
		Id:           req.Id,
		HotelId:      id,
		GuestRoomNo:  req.GuestRoomNo,
		GuestName:    "",
		GuestIdNo:    "",
		RealPrice:    0,
		Mobile:       "",
		CheckInTime:  "",
		CheckOutTime: "",
		StayStatus:   4,
		Description:  req.Description,
	}

	_, err = RoomServiceClient.UpdateRoomGuestStay(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	c.JSON(http.StatusOK, result.Sucess("禁用成功"))
}
