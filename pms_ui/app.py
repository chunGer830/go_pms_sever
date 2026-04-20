from __future__ import annotations

from PyQt6.QtCore import Qt
from PyQt6.QtGui import QColor, QLinearGradient, QPainter
from PyQt6.QtWidgets import QApplication, QLabel, QMainWindow, QMessageBox, QStackedWidget, QWidget

from pms_ui.services.auth_service import AuthService
from pms_ui.theme import APP_STYLE
from pms_ui.views.login_view import LoginView
from pms_ui.views.main_window import PMSMainWindow


class GradientHost(QWidget):
    def paintEvent(self, event) -> None:  # noqa: N802
        painter = QPainter(self)
        gradient = QLinearGradient(0, 0, self.width(), self.height())
        gradient.setColorAt(0.0, QColor("#F8FBFF"))
        gradient.setColorAt(0.55, QColor("#F4F7FB"))
        gradient.setColorAt(1.0, QColor("#EEF3F9"))
        painter.fillRect(self.rect(), gradient)
        painter.setPen(QColor(48, 92, 160, 20))
        painter.drawEllipse(-140, -40, 420, 420)
        painter.drawEllipse(self.width() - 360, self.height() - 260, 420, 420)
        super().paintEvent(event)


class AppShell(QMainWindow):
    def __init__(self) -> None:
        super().__init__()
        self.setWindowTitle("Hotel PMS")
        self.resize(1520, 920)

        self.auth_service = AuthService()
        self.access_token = ""
        self.refresh_token = ""
        self.token_type = ""
        self.access_token_exp = 0
        self.current_user: dict[str, object] | None = None

        self.stack = QStackedWidget()
        host = GradientHost()
        self.setCentralWidget(host)
        self.stack.setParent(host)
        self.stack.setGeometry(host.rect())

        self.login_view = LoginView()
        self.login_view.login_requested.connect(self._handle_login)
        self.login_view.password_change_requested.connect(self._handle_password_change)
        self.main_window = PMSMainWindow()
        self.main_window.logout_requested.connect(self._handle_logout)
        self.main_window.room_type_refresh_requested.connect(self._handle_room_type_refresh)

        self.stack.addWidget(self.login_view)
        self.stack.addWidget(self.main_window)
        self.stack.setCurrentIndex(0)

    def resizeEvent(self, event) -> None:  # noqa: N802
        super().resizeEvent(event)
        self.stack.setGeometry(self.centralWidget().rect())

    def _handle_login(self, username: str, password: str) -> None:
        if not username or not password:
            self._show_message_dialog("登录失败", "请输入用户名和密码。", success=False)
            return

        result = self.auth_service.login({"username": username, "password": password})
        if not result.get("success"):
            self.auth_service.clear_auth_token()
            self._show_message_dialog("登录失败", str(result.get("message", "登录失败")), success=False)
            return

        data = result.get("data")
        if not isinstance(data, dict):
            self.auth_service.clear_auth_token()
            self._show_message_dialog("登录失败", "登录返回数据格式不正确。", success=False)
            return

        member = data.get("member")
        token_list = data.get("tokenList")
        if not isinstance(member, dict) or not isinstance(token_list, dict):
            self.auth_service.clear_auth_token()
            self._show_message_dialog("登录失败", "登录返回缺少用户或令牌信息。", success=False)
            return

        self.current_user = member
        self.access_token = str(token_list.get("accessToken", ""))
        self.refresh_token = str(token_list.get("refreshToken", ""))
        self.token_type = str(token_list.get("tokenType", ""))
        self.access_token_exp = int(token_list.get("accessTokenExp", 0) or 0)
        self.auth_service.set_auth_token(self.access_token, self.token_type)
        self.stack.setCurrentIndex(1)

    def _handle_unauthorized(self, result: dict[str, object]) -> bool:
        if not result.get("unauthorized"):
            return False
        self._handle_logout(show_message=False)
        self._show_message_dialog("登录已失效", str(result.get("message", "登录已失效，请重新登录。")), success=False)
        return True

    def _handle_room_type_refresh(self) -> None:
        result = self.auth_service.get_room_types()
        if self._handle_unauthorized(result):
            return
        if not result.get("success"):
            self._show_message_dialog("房型控制", str(result.get("message", "获取房型数据失败")), success=False)
            return

        data = result.get("data")
        if isinstance(data, dict) and isinstance(data.get("list"), list):
            items = [item for item in data["list"] if isinstance(item, dict)]
        elif isinstance(data, dict):
            items = [data]
        elif isinstance(data, list):
            items = [item for item in data if isinstance(item, dict)]
        else:
            self._show_message_dialog("房型控制", "房型数据格式不正确。", success=False)
            return

        self.main_window.set_room_type_items(items)

    def _show_message_dialog(self, title: str, message: str, *, success: bool) -> None:
        dialog = QMessageBox(self)
        dialog.setWindowTitle(title)
        dialog.setText(message)
        dialog.setIcon(QMessageBox.Icon.NoIcon)
        dialog.setMinimumSize(460, 220)
        dialog.setStyleSheet(
            """
            QMessageBox {
                background-color: #FFFFFF;
            }
            QLabel {
                color: #1E2A3B;
                min-width: 260px;
                min-height: 72px;
                font-size: 14px;
            }
            QPushButton {
                background-color: #2F7AF8;
                color: white;
                border: none;
                border-radius: 8px;
                padding: 8px 18px;
                min-width: 88px;
            }
            QPushButton:hover {
                background-color: #1F68E5;
            }
            """
        )
        for label in dialog.findChildren(QLabel):
            if label.text() == message:
                label.setAlignment(Qt.AlignmentFlag.AlignCenter)
                label.setStyleSheet(
                    "color: #223248; font-size: 18px; font-weight: 500; "
                    'font-family: "Microsoft YaHei UI", "Segoe UI";'
                )
        dialog.exec()

    def _handle_password_change(self, payload: dict[str, str]) -> None:
        result = self.auth_service.change_password(payload)
        if self._handle_unauthorized(result):
            return
        if result.get("success"):
            self._show_message_dialog("修改密码", "修改成功", success=True)
            return
        self._show_message_dialog("修改密码失败", str(result.get("message", "修改失败")), success=False)

    def _handle_logout(self, show_message: bool = False) -> None:
        self.access_token = ""
        self.refresh_token = ""
        self.token_type = ""
        self.access_token_exp = 0
        self.current_user = None
        self.auth_service.clear_auth_token()
        self.login_view.password_edit.clear()
        self.login_view.username_edit.clear()
        self.stack.setCurrentIndex(0)
        if show_message:
            self._show_message_dialog("退出登录", "已退出登录。", success=True)


def run() -> int:
    app = QApplication([])
    app.setStyleSheet(APP_STYLE)
    window = AppShell()
    window.show()
    return app.exec()
