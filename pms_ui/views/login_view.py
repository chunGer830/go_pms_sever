from __future__ import annotations

from PyQt6.QtCore import Qt, pyqtSignal
from PyQt6.QtGui import QFont
from PyQt6.QtWidgets import (
    QDialog,
    QDialogButtonBox,
    QFrame,
    QHBoxLayout,
    QLabel,
    QLineEdit,
    QPushButton,
    QVBoxLayout,
    QWidget,
)

from pms_ui.widgets import BulletRow, GlassCard


class ChangePasswordDialog(QDialog):
    def __init__(self, username: str = "", parent: QWidget | None = None) -> None:
        super().__init__(parent)
        self.setWindowTitle("修改密码")
        self.setModal(True)
        self.resize(420, 320)
        self.setStyleSheet(
            """
            QDialog {
                background-color: #FFFFFF;
            }
            QLabel {
                color: #1E2A3B;
            }
            QLineEdit {
                background-color: #FFFFFF;
                color: #1F2E42;
                border: 1px solid #D8E1EC;
                border-radius: 12px;
                padding: 10px 12px;
            }
            QPushButton {
                min-width: 88px;
            }
            """
        )
        self._build_ui(username)

    def _build_ui(self, username: str) -> None:
        layout = QVBoxLayout(self)
        layout.setContentsMargins(20, 20, 20, 20)
        layout.setSpacing(14)

        title = QLabel("修改密码")
        title.setStyleSheet("font-size: 20px; font-weight: 700;")
        self.username_edit = QLineEdit()
        self.username_edit.setPlaceholderText("用户名")
        self.username_edit.setText(username)
        self.old_password_edit = QLineEdit()
        self.old_password_edit.setPlaceholderText("原密码")
        self.old_password_edit.setEchoMode(QLineEdit.EchoMode.Password)
        self.new_password_edit = QLineEdit()
        self.new_password_edit.setPlaceholderText("新密码")
        self.new_password_edit.setEchoMode(QLineEdit.EchoMode.Password)
        self.confirm_password_edit = QLineEdit()
        self.confirm_password_edit.setPlaceholderText("确认新密码")
        self.confirm_password_edit.setEchoMode(QLineEdit.EchoMode.Password)

        self.error_label = QLabel("")
        self.error_label.setStyleSheet("color: #D35B57; font-size: 12px;")
        self.error_label.setWordWrap(True)

        buttons = QDialogButtonBox(QDialogButtonBox.StandardButton.Save | QDialogButtonBox.StandardButton.Cancel)
        save_button = buttons.button(QDialogButtonBox.StandardButton.Save)
        cancel_button = buttons.button(QDialogButtonBox.StandardButton.Cancel)
        if save_button is not None:
            save_button.setText("确认修改")
        if cancel_button is not None:
            cancel_button.setText("取消")
        buttons.accepted.connect(self._validate_and_accept)
        buttons.rejected.connect(self.reject)

        layout.addWidget(title)
        layout.addWidget(self.username_edit)
        layout.addWidget(self.old_password_edit)
        layout.addWidget(self.new_password_edit)
        layout.addWidget(self.confirm_password_edit)
        layout.addWidget(self.error_label)
        layout.addWidget(buttons)

    def _validate_and_accept(self) -> None:
        username = self.username_edit.text().strip()
        old_password = self.old_password_edit.text()
        new_password = self.new_password_edit.text()
        confirm_password = self.confirm_password_edit.text()

        if not username or not old_password or not new_password or not confirm_password:
            self.error_label.setText("请完整填写用户名、原密码和新密码。")
            return
        if new_password != confirm_password:
            self.error_label.setText("两次输入的新密码不一致。")
            return
        if len(new_password) < 6:
            self.error_label.setText("新密码长度至少 6 位。")
            return
        self.accept()

    def get_payload(self) -> dict[str, str]:
        return {
            "username": self.username_edit.text().strip(),
            "old_password": self.old_password_edit.text(),
            "password": self.new_password_edit.text(),
            "password2": self.new_password_edit.text(),
        }


class LoginView(QWidget):
    login_requested = pyqtSignal(str, str)
    password_change_requested = pyqtSignal(dict)

    def __init__(self, parent: QWidget | None = None) -> None:
        super().__init__(parent)
        self._build_ui()

    def _build_ui(self) -> None:
        root = QHBoxLayout(self)
        root.setContentsMargins(48, 40, 48, 40)
        root.setSpacing(24)

        hero = QFrame()
        hero_layout = QVBoxLayout(hero)
        hero_layout.setContentsMargins(20, 16, 20, 16)
        hero_layout.setSpacing(18)

        badge = QLabel("LUXURY PMS SUITE")
        badge.setProperty("role", "eyebrow")
        title = QLabel("为高端酒店打造的\n下一代前台工作台")
        title_font = QFont()
        title_font.setPointSize(28)
        title_font.setBold(True)
        title.setFont(title_font)
        title.setStyleSheet("color: #122033; line-height: 1.2;")
        desc = QLabel("统一承接预订、房态、房型定价与住客运营。前端已预留接口位，便于你后续接入自定义后端。")
        desc.setWordWrap(True)
        desc.setProperty("role", "subtle")
        desc.setStyleSheet("font-size: 15px;")

        feature_card = GlassCard(panel=True)
        feature_layout = QVBoxLayout(feature_card)
        feature_layout.setContentsMargins(20, 20, 20, 20)
        feature_layout.setSpacing(12)
        feature_title = QLabel("核心亮点")
        feature_title.setStyleSheet("font-size: 18px; font-weight: 700; color: #F3F7FF;")
        feature_layout.addWidget(feature_title)
        feature_layout.addWidget(BulletRow("接近 JetBrains 官网的深色科技感布局与分层卡片。"))
        feature_layout.addWidget(BulletRow("房型控制页已具备检索、表格浏览、策略录入与通知区。"))
        feature_layout.addWidget(BulletRow("所有数据来自前端 mock，可平滑替换为真实 API。"))

        hero_layout.addStretch()
        hero_layout.addWidget(badge)
        hero_layout.addWidget(title)
        hero_layout.addWidget(desc)
        hero_layout.addWidget(feature_card)
        hero_layout.addStretch()

        login_card = GlassCard()
        login_card.setMaximumWidth(420)
        login_layout = QVBoxLayout(login_card)
        login_layout.setContentsMargins(28, 28, 28, 28)
        login_layout.setSpacing(16)

        welcome = QLabel("登录 PMS")
        welcome.setProperty("role", "headline")
        welcome.setStyleSheet("font-size: 26px;")
        subtitle = QLabel("使用你的酒店账号进入系统。当前登录逻辑为前端占位，可替换为真实认证。")
        subtitle.setProperty("role", "subtle")
        subtitle.setWordWrap(True)

        self.hotel_edit = QLineEdit()
        self.hotel_edit.setPlaceholderText("酒店代码 / Hotel Code")
        self.username_edit = QLineEdit()
        self.username_edit.setPlaceholderText("用户名")
        self.password_edit = QLineEdit()
        self.password_edit.setPlaceholderText("密码")
        self.password_edit.setEchoMode(QLineEdit.EchoMode.Password)

        sign_in = QPushButton("进入系统")
        sign_in.setProperty("variant", "primary")
        sign_in.clicked.connect(self._emit_login)
        change_password_button = QPushButton("修改密码")
        change_password_button.clicked.connect(self._open_change_password_dialog)

        hint = QLabel("演示账号可直接输入任意内容登录。")
        hint.setProperty("role", "subtle")

        login_layout.addWidget(welcome)
        login_layout.addWidget(subtitle)
        login_layout.addSpacing(8)
        login_layout.addWidget(self.hotel_edit)
        login_layout.addWidget(self.username_edit)
        login_layout.addWidget(self.password_edit)
        login_layout.addSpacing(8)
        login_layout.addWidget(sign_in)
        login_layout.addWidget(change_password_button)
        login_layout.addWidget(hint)
        login_layout.addStretch()

        root.addWidget(hero, stretch=7)
        root.addWidget(login_card, stretch=4, alignment=Qt.AlignmentFlag.AlignCenter)

    def _emit_login(self) -> None:
        self.login_requested.emit(self.username_edit.text().strip(), self.password_edit.text())

    def _open_change_password_dialog(self) -> None:
        dialog = ChangePasswordDialog(self.username_edit.text().strip(), self)
        if dialog.exec() == QDialog.DialogCode.Accepted:
            self.password_change_requested.emit(dialog.get_payload())
