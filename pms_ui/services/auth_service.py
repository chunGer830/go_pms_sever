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

    def set_auth_token(self, access_token: str, token_type: str = "Bearer") -> None:
        self.access_token = access_token
        self.token_type = token_type or "Bearer"

    def clear_auth_token(self) -> None:
        self.access_token = ""
        self.token_type = "Bearer"

    def _post_json(self, path: str, payload: dict[str, str], *, fallback_message: str) -> dict[str, object]:
        return self._request_json("POST", path, payload=payload, fallback_message=fallback_message)

    def _request_json(
        self,
        method: str,
        path: str,
        payload: dict[str, str] | None = None,
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
            return {"success": message == "修改成功", "message": message or fallback_message}

        if isinstance(data, dict):
            message = str(data.get("message") or data.get("msg") or fallback_message)
            success = data.get("code") == 0 or message == "修改成功"
            unauthorized = data.get("code") in {401, 1001, 1002} or "未登录" in message or "token" in message.lower()
            return {
                "success": success,
                "message": message,
                "data": data.get("data"),
                "unauthorized": unauthorized,
            }

        return {"success": False, "message": fallback_message}
