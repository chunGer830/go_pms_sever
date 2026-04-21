from __future__ import annotations

from pms_ui.models import Member, RechargeRule, Reservation, RoomRecord, RoomStatus, RoomType


def get_room_types() -> list[RoomType]:
    return [
        RoomType("行政大床房", "DLX-K", 18, 1280, 2, True, True, "开放预订"),
        RoomType("城景双床房", "CTY-T", 24, 980, 2, True, True, "开放预订"),
        RoomType("套房", "STE-S", 8, 2380, 3, True, True, "限量预订"),
        RoomType("轻奢单人房", "SGL-L", 16, 760, 1, False, False, "维护中"),
    ]


def get_reservations() -> list[Reservation]:
    return [
        Reservation("周晓晨", "行政大床房", "03-23", "03-25", "官网直订", 2560, "已确认"),
        Reservation("李晗", "套房", "03-23", "03-24", "企业协议", 2380, "待入住"),
        Reservation("陈若然", "城景双床房", "03-24", "03-27", "OTA", 2940, "担保中"),
        Reservation("王可欣", "轻奢单人房", "03-25", "03-26", "前台散客", 760, "待付款"),
    ]


def get_room_statuses() -> list[RoomStatus]:
    return [
        RoomStatus(
            "1801",
            "套房",
            "18F",
            "在住",
            "周晓晨",
            "320***********2211",
            "03-23 入住 / 03-25 离店",
            "已完成入住登记",
            ["03-23 14:05 前台办理入住", "03-23 14:12 已录入证件信息", "03-24 09:10 补录早餐权益"],
        ),
        RoomStatus(
            "1802",
            "套房",
            "18F",
            "待清理",
            "",
            "",
            "上一位客人 11:20 退房",
            "待客房清扫",
            ["03-24 11:20 客人退房", "03-24 11:21 状态切换为待清理"],
        ),
        RoomStatus(
            "1506",
            "行政大床房",
            "15F",
            "待检",
            "",
            "",
            "客房已清扫，待主管查房",
            "待主管放房",
            ["03-24 10:15 客房完成清扫", "03-24 10:18 状态切换为待检"],
        ),
        RoomStatus(
            "1208",
            "城景双床房",
            "12F",
            "预抵",
            "李晗",
            "510***********8703",
            "今日 18:00 预计到店",
            "已保留房间",
            ["03-24 09:00 创建预抵房", "03-24 09:05 锁定房号 1208"],
        ),
        RoomStatus(
            "907",
            "轻奢单人房",
            "9F",
            "维修",
            "",
            "",
            "空调检修中",
            "暂停销售",
            ["03-23 20:30 报修空调", "03-24 08:10 工程部处理中"],
        ),
        RoomStatus(
            "1012",
            "城景双床房",
            "10F",
            "空净",
            "",
            "",
            "可直接出售",
            "已放房",
            ["03-24 08:40 完成查房", "03-24 08:45 状态切换为空净"],
        ),
    ]


def get_room_records() -> list[RoomRecord]:
    return [
        RoomRecord(
            room_no=room.room_no,
            room_type=room.room_type,
            floor=room.floor,
            id="",
            phone_ext=f"6{room.room_no[-3:]}",
            remark=room.last_action,
        )
        for room in get_room_statuses()
    ]


def get_dashboard_metrics() -> list[tuple[str, str, str]]:
    return [
        ("今日入住", "46", "+12%"),
        ("今日离店", "31", "+4%"),
        ("出租率", "82%", "+6%"),
        ("RevPAR", "¥684", "+9%"),
    ]


def get_alerts() -> list[str]:
    return [
        "03 间房待夜审复核",
        "2 个企业协议价将于 48 小时内过期",
        "套房库存低于预警阈值",
    ]


def get_members() -> list[Member]:
    return [
        Member("M202603001", "林若琪", "铂金会员", "13800138001", 5680, 12800, "正常"),
        Member("M202603002", "陈知行", "金卡会员", "13800138002", 3200, 8600, "正常"),
        Member("M202603003", "苏念", "银卡会员", "13800138003", 980, 2600, "休眠"),
        Member("M202603004", "顾远", "黑金会员", "13800138004", 12880, 26800, "正常"),
    ]


def get_recharge_rules() -> list[RechargeRule]:
    return [
        RechargeRule("新客首充", 1000, 80, "前台", "全部会员", "启用"),
        RechargeRule("金卡充值礼", 2000, 220, "小程序", "金卡及以上", "启用"),
        RechargeRule("企业会员储值", 5000, 800, "协议客户", "企业会员", "启用"),
        RechargeRule("节庆活动", 3000, 388, "全渠道", "全部会员", "停用"),
    ]
