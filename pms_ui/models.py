from __future__ import annotations

from dataclasses import dataclass, field


@dataclass(slots=True)
class RoomType:
    name: str
    code: str
    total_rooms: int
    base_price: int
    occupancy: int
    breakfast: bool
    window: bool
    status: str


@dataclass(slots=True)
class Reservation:
    guest_name: str
    room_type: str
    check_in: str
    check_out: str
    channel: str
    amount: int
    status: str


@dataclass(slots=True)
class RoomStatus:
    room_no: str
    room_type: str
    floor: str
    state: str
    guest: str
    guest_id_no: str = ""
    stay_label: str = ""
    last_action: str = ""
    logs: list[str] = field(default_factory=list)


@dataclass(slots=True)
class Member:
    member_no: str
    name: str
    level: str
    phone: str
    balance: int
    points: int
    status: str


@dataclass(slots=True)
class RechargeRule:
    name: str
    amount: int
    bonus: int
    channel: str
    level_limit: str
    status: str
