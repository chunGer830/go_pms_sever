from __future__ import annotations

from PyQt6.QtCore import Qt
from PyQt6.QtWidgets import (
    QPushButton,
    QFrame,
    QGraphicsDropShadowEffect,
    QHBoxLayout,
    QLabel,
    QVBoxLayout,
    QWidget,
)


class GlassCard(QFrame):
    def __init__(self, *, panel: bool = False, parent: QWidget | None = None) -> None:
        super().__init__(parent)
        self.setProperty("card", not panel)
        self.setProperty("panel", panel)
        shadow = QGraphicsDropShadowEffect(self)
        shadow.setBlurRadius(36)
        shadow.setOffset(0, 12)
        shadow.setColor(Qt.GlobalColor.gray)
        self.setGraphicsEffect(shadow)


class MetricCard(GlassCard):
    def __init__(self, title: str, value: str, delta: str, parent: QWidget | None = None) -> None:
        super().__init__(parent=parent)
        layout = QVBoxLayout(self)
        layout.setContentsMargins(18, 18, 18, 18)
        layout.setSpacing(10)

        title_label = QLabel(title)
        title_label.setProperty("role", "subtle")
        value_label = QLabel(value)
        value_label.setProperty("role", "headline")
        value_label.setStyleSheet("font-size: 28px;")
        delta_label = QLabel(delta)
        delta_label.setStyleSheet(
            "color: #198754; background: rgba(25, 135, 84, 0.12);"
            "padding: 4px 8px; border-radius: 8px; font-weight: 600;"
        )

        layout.addWidget(title_label)
        layout.addWidget(value_label)
        layout.addStretch()
        layout.addWidget(delta_label, alignment=Qt.AlignmentFlag.AlignLeft)


class SectionTitle(QWidget):
    def __init__(self, title: str, description: str, parent: QWidget | None = None) -> None:
        super().__init__(parent)
        layout = QVBoxLayout(self)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(4)
        heading = QLabel(title)
        heading.setProperty("role", "headline")
        heading.setStyleSheet("font-size: 22px;")
        desc = QLabel(description)
        desc.setProperty("role", "subtle")
        layout.addWidget(heading)
        layout.addWidget(desc)


class BulletRow(QWidget):
    def __init__(self, text: str, parent: QWidget | None = None) -> None:
        super().__init__(parent)
        layout = QHBoxLayout(self)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(10)
        dot = QLabel("●")
        dot.setStyleSheet("color: #4E86F7;")
        label = QLabel(text)
        label.setWordWrap(True)
        label.setProperty("role", "subtle")
        layout.addWidget(dot, alignment=Qt.AlignmentFlag.AlignTop)
        layout.addWidget(label)


class SidebarSection(QFrame):
    def __init__(self, title: str, items: list[tuple[str, callable]], expanded: bool = True, parent: QWidget | None = None) -> None:
        super().__init__(parent)
        self.setStyleSheet("background: transparent;")
        self._expanded = expanded

        layout = QVBoxLayout(self)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(6)

        self.header_button = QPushButton(f"{title}  {'▾' if expanded else '▸'}")
        self.header_button.setProperty("variant", "ghost")
        self.header_button.setStyleSheet("font-size: 15px; font-weight: 700;")
        self.header_button.clicked.connect(self._toggle)
        layout.addWidget(self.header_button)

        self.body = QWidget()
        body_layout = QVBoxLayout(self.body)
        body_layout.setContentsMargins(14, 0, 0, 0)
        body_layout.setSpacing(6)
        for text, callback in items:
            button = QPushButton(text)
            button.setProperty("variant", "ghost")
            button.setStyleSheet(
                "font-size: 14px; font-weight: 500; color: #D9E6F8; text-align: left; padding-left: 18px;"
            )
            button.clicked.connect(callback)
            body_layout.addWidget(button)
        layout.addWidget(self.body)
        self.body.setVisible(expanded)

    def _toggle(self) -> None:
        self._expanded = not self._expanded
        self.header_button.setText(self.header_button.text()[:-1] + ("▾" if self._expanded else "▸"))
        self.body.setVisible(self._expanded)


class RoomStatusCard(QPushButton):
    STATUS_COLORS = {
        "空净": ("#EAF8EE", "#2E8B57"),
        "在住": ("#E9F1FF", "#2F6BDE"),
        "预抵": ("#FFF6E8", "#D68A28"),
        "待清理": ("#FFF1F0", "#D35B57"),
        "待检": ("#F3F0FF", "#7A5AF8"),
        "维修": ("#F0F3F6", "#6B7280"),
        "禁用": ("#F0F3F6", "#6B7280"),
    }

    def __init__(self, room, parent: QWidget | None = None) -> None:
        super().__init__(parent)
        self.room = room
        self.setCursor(Qt.CursorShape.PointingHandCursor)
        self.setMinimumHeight(148)
        self._build_content()
        self.setStyleSheet(self._style_for_room())
        self._update_content()

    def _style_for_room(self) -> str:
        bg, accent = self.STATUS_COLORS.get(self.room.state, ("#FFFFFF", "#2F7AF8"))
        return f"""
        QPushButton {{
            background: {bg};
            border: 1px solid #DDE5EF;
            border-left: 5px solid {accent};
            border-radius: 16px;
            padding: 0px;
            text-align: left;
        }}
        QPushButton:hover {{
            border: 1px solid {accent};
            border-left: 5px solid {accent};
        }}
        QLabel {{
            background: transparent;
            color: #142033;
        }}
        QLabel#roomTitle {{
            font-size: 15px;
            font-weight: 700;
        }}
        QLabel#roomType {{
            font-size: 13px;
            font-weight: 600;
            color: #2F425A;
        }}
        QLabel#roomGuest {{
            font-size: 13px;
            font-weight: 500;
            color: #223248;
        }}
        QLabel#roomMeta {{
            font-size: 12px;
            font-weight: 500;
            color: #51657E;
        }}
        """

    def _build_content(self) -> None:
        layout = QVBoxLayout(self)
        layout.setContentsMargins(14, 12, 12, 12)
        layout.setSpacing(6)

        self.title_label = QLabel()
        self.title_label.setObjectName("roomTitle")
        self.type_label = QLabel()
        self.type_label.setObjectName("roomType")
        self.guest_label = QLabel()
        self.guest_label.setObjectName("roomGuest")
        self.meta_label = QLabel()
        self.meta_label.setObjectName("roomMeta")
        self.meta_label.setWordWrap(True)

        for label in (self.title_label, self.type_label, self.guest_label, self.meta_label):
            label.setAttribute(Qt.WidgetAttribute.WA_TransparentForMouseEvents, True)
            layout.addWidget(label)
        layout.addStretch()

    def _update_content(self) -> None:
        guest = self.room.guest if self.room.guest else "当前无住客"
        self.title_label.setText(f"{self.room.room_no}  {self.room.state}")
        self.type_label.setText(self.room.room_type)
        self.guest_label.setText(guest)
        self.meta_label.setText(self.room.stay_label or self.room.last_action or "暂无动态")

    def refresh(self) -> None:
        self.setStyleSheet(self._style_for_room())
        self._update_content()
