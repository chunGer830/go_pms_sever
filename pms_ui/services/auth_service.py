from __future__ import annotations

import json
import os
from urllib import error, request


class AuthService:
    def __init__(self, base_url: str | None = None) -> None:
        self.base_url = (base_url or os.getenv("PMS_API_BASE_URL") or "http://127.0.0.1").rstrip("/")
        self.access_token = ""
        self.token_type = "Bearer"

    def login(self, payload: dict[str, str]) -> dict[str, object]:
        return self._post_json("/project/login", payload, fallback_message="登录请求失败")

    def change_password(self, payload: dict[str, str]) -> dict[str, object]:
        return self._post_json("/project/login/changePassword", payload, fallback_message="修改密码请求失败")

    def get_room_types(self) -> dict[str, object]:
        return self._request_json("GET", "/project/room/roomType", fallback_message="获取房型数据失败")

    def get_hotel_rooms(self) -> dict[str, object]:
        return self._request_json("GET", "/project/room/hotelRoom", fallback_message="获取房间数据失败")

    def get_room_guest_stays(self) -> dict[str, object]:
        return self._request_json("GET", "/project/room/roomGuestStay", fallback_message="获取房态数据失败")

    def update_room_guest_stay(self, payload: dict[str, object]) -> dict[str, object]:
        request_payload = {
            "guest_room_no": str(payload.get("guest_room_no", "")).strip(),
            "guest_name": str(payload.get("guest_name", "")).strip(),
            "guest_id_no": str(payload.get("guest_id_no", "")).strip(),
            "real_price": int(round(self._to_float(payload.get("real_price")) * 100)),
            "mobile": str(payload.get("mobile", "")).strip(),
            "check_in_time": str(payload.get("check_in_time", "")).strip(),
            "check_out_time": str(payload.get("check_out_time", "")).strip(),
            "stay_status": self._to_int(payload.get("stay_status")),
            "description": str(payload.get("description", "")).strip(),
        }
        return self._request_json(
            "POST",
            "/project/room/updateRoomGuestStay",
            payload=request_payload,
            fallback_message="更新房态失败",
        )

    def checkout_room_guest_stay(self, payload: dict[str, object]) -> dict[str, object]:
        request_payload = {
            "guest_room_no": str(payload.get("guest_room_no", "")).strip(),
            "guest_name": str(payload.get("guest_name", "")).strip(),
            "guest_id_no": str(payload.get("guest_id_no", "")).strip(),
            "real_price": int(round(self._to_float(payload.get("real_price")) * 100)),
            "mobile": str(payload.get("mobile", "")).strip(),
            "check_in_time": str(payload.get("check_in_time", "")).strip(),
            "check_out_time": str(payload.get("check_out_time", "")).strip(),
            "stay_status": self._to_int(payload.get("stay_status")),
            "description": str(payload.get("description", "")).strip(),
        }
        return self._request_json(
            "POST",
            "/project/room/checkoutRoomGuestStay",
            payload=request_payload,
            fallback_message="退房失败",
        )

    def clean_room_guest_stay(self, payload: dict[str, object]) -> dict[str, object]:
        request_payload = {
            "guest_room_no": str(payload.get("guest_room_no", "")).strip(),
            "guest_name": str(payload.get("guest_name", "")).strip(),
            "guest_id_no": str(payload.get("guest_id_no", "")).strip(),
            "real_price": int(round(self._to_float(payload.get("real_price")) * 100)),
            "mobile": str(payload.get("mobile", "")).strip(),
            "check_in_time": str(payload.get("check_in_time", "")).strip(),
            "check_out_time": str(payload.get("check_out_time", "")).strip(),
            "stay_status": self._to_int(payload.get("stay_status")),
            "description": str(payload.get("description", "")).strip(),
        }
        return self._request_json(
            "POST",
            "/project/room/cleanRoomGuestStay",
            payload=request_payload,
            fallback_message="清理完成失败",
        )

    def disable_room_guest_stay(self, payload: dict[str, object]) -> dict[str, object]:
        request_payload = {
            "guest_room_no": str(payload.get("guest_room_no", "")).strip(),
            "guest_name": str(payload.get("guest_name", "")).strip(),
            "guest_id_no": str(payload.get("guest_id_no", "")).strip(),
            "real_price": int(round(self._to_float(payload.get("real_price")) * 100)),
            "mobile": str(payload.get("mobile", "")).strip(),
            "check_in_time": str(payload.get("check_in_time", "")).strip(),
            "check_out_time": str(payload.get("check_out_time", "")).strip(),
            "stay_status": self._to_int(payload.get("stay_status")),
            "description": str(payload.get("description", "")).strip(),
        }
        return self._request_json(
            "POST",
            "/project/room/disableRoomGuestStay",
            payload=request_payload,
            fallback_message="禁用失败",
        )

    def save_hotel_room(self, payload: dict[str, object]) -> dict[str, object]:
        request_payload = {
            "room_no": str(payload.get("room_no", "")).strip(),
            "room_type_name": str(payload.get("room_type_name", "")).strip(),
            "room_type_code": str(payload.get("room_type_code", "")).strip(),
            "floor_no": str(payload.get("floor_no", "")).strip(),
            "phone_ext": str(payload.get("phone_ext", "")).strip(),
            "description": str(payload.get("description", "")).strip(),
        }
        return self._request_json(
            "POST",
            "/project/room/saveHotelRoom",
            payload=request_payload,
            fallback_message="新增房间失败",
        )

    def update_hotel_room(self, payload: dict[str, object]) -> dict[str, object]:
        raw_id = str(payload.get("id", "")).strip()
        try:
            room_id = int(raw_id)
        except (TypeError, ValueError):
            room_id = 0

        request_payload = {
            "id": room_id,
            "room_no": str(payload.get("room_no", "")).strip(),
            "room_type_name": str(payload.get("room_type_name", "")).strip(),
            "room_type_code": str(payload.get("room_type_code", "")).strip(),
            "floor_no": str(payload.get("floor_no", "")).strip(),
            "phone_ext": str(payload.get("phone_ext", "")).strip(),
            "description": str(payload.get("description", "")).strip(),
        }
        return self._request_json(
            "POST",
            "/project/room/updateHotelRoom",
            payload=request_payload,
            fallback_message="修改房间失败",
        )

    def delete_hotel_room(self, payload: dict[str, object]) -> dict[str, object]:
        request_payload = {
            "room_no": str(payload.get("room_no", "")).strip(),
        }
        return self._request_json(
            "POST",
            "/project/room/deleteHotelRoom",
            payload=request_payload,
            fallback_message="删除房间失败",
        )

    def save_room_type(self, payload: dict[str, object]) -> dict[str, object]:
        request_payload = {
            "type_name": str(payload.get("name", "")).strip(),
            "type_code": str(payload.get("code", "")).strip(),
            "max_occupancy": int(payload.get("occupancy", 0) or 0),
            "base_price": int(payload.get("base_price", 0) or 0) * 100,
            "quantity": int(payload.get("total_rooms", 0) or 0),
            "status": 1 if str(payload.get("status", "")).strip() == "启用" else 0,
        }
        return self._request_json(
            "POST",
            "/project/room/saveRoomType",
            payload=request_payload,
            fallback_message="新增房型失败",
        )

    def update_room_type(self, payload: dict[str, object]) -> dict[str, object]:
        raw_id = str(payload.get("id", "")).strip()
        try:
            room_type_id = int(raw_id)
        except (TypeError, ValueError):
            room_type_id = 0

        request_payload = {
            "id": room_type_id,
            "type_name": str(payload.get("name", "")).strip(),
            "type_code": str(payload.get("code", "")).strip(),
            "max_occupancy": int(payload.get("occupancy", 0) or 0),
            "base_price": int(payload.get("base_price", 0) or 0) * 100,
            "quantity": int(payload.get("total_rooms", 0) or 0),
            "status": 1 if str(payload.get("status", "")).strip() == "启用" else 0,
        }
        return self._request_json(
            "POST",
            "/project/room/updateRoomType",
            payload=request_payload,
            fallback_message="修改房型失败",
        )

    def delete_room_type(self, payload: dict[str, object]) -> dict[str, object]:
        request_payload = {
            "type_code": str(payload.get("type_code", "")).strip(),
        }
        return self._request_json(
            "POST",
            "/project/room/deleteRoomType",
            payload=request_payload,
            fallback_message="删除房型失败",
        )

    def set_auth_token(self, access_token: str, token_type: str = "Bearer") -> None:
        self.access_token = access_token
        self.token_type = token_type or "Bearer"

    def clear_auth_token(self) -> None:
        self.access_token = ""
        self.token_type = "Bearer"

    def _post_json(self, path: str, payload: dict[str, str], *, fallback_message: str) -> dict[str, object]:
        return self._request_json("POST", path, payload=payload, fallback_message=fallback_message)

    @staticmethod
    def _to_int(value: object) -> int:
        text = str(value or "").strip()
        if not text:
            return 0
        try:
            return int(float(text))
        except (TypeError, ValueError):
            return 0

    @staticmethod
    def _to_float(value: object) -> float:
        text = str(value or "").strip()
        if not text:
            return 0.0
        try:
            return float(text)
        except (TypeError, ValueError):
            return 0.0

    def _request_json(
        self,
        method: str,
        path: str,
        payload: dict[str, object] | None = None,
        *,
        fallback_message: str,
    ) -> dict[str, object]:
        headers = {"Content-Type": "application/json"}
        if self.access_token:
            headers["Authorization"] = self.access_token

        req = request.Request(
            url=f"{self.base_url}{path}",
            data=json.dumps(payload).encode("utf-8") if payload is not None else None,
            headers=headers,
            method=method,
        )
        try:
            with request.urlopen(req, timeout=10) as response:
                body = response.read().decode("utf-8")
        except error.HTTPError as exc:
            body = exc.read().decode("utf-8", errors="ignore")
            result = self._parse_response(body, fallback_message=f"请求失败: HTTP {exc.code}")
            if exc.code == 401:
                result["unauthorized"] = True
                result["message"] = str(result.get("message") or "登录已失效，请重新登录。")
            return result
        except error.URLError as exc:
            return {"success": False, "message": f"无法连接后端: {exc.reason}"}

        return self._parse_response(body, fallback_message=fallback_message)

    @staticmethod
    def _parse_response(body: str, fallback_message: str) -> dict[str, object]:
        if not body:
            return {"success": False, "message": fallback_message}
        try:
            data = json.loads(body)
        except json.JSONDecodeError:
            message = body.strip()
            return {"success": message in {"修改成功", "添加成功", "删除成功"}, "message": message or fallback_message}

        if isinstance(data, dict):
            message = str(data.get("message") or data.get("msg") or fallback_message)
            success = data.get("code") == 0 or message in {"修改成功", "添加成功", "删除成功"}
            unauthorized = data.get("code") in {401, 1001, 1002} or "未登录" in message or "token" in message.lower()
            return {
                "success": success,
                "message": message,
                "data": data.get("data"),
                "unauthorized": unauthorized,
            }

        return {"success": False, "message": fallback_message}
