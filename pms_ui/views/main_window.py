from __future__ import annotations

from PyQt6.QtCore import Qt, pyqtSignal
from PyQt6.QtWidgets import (
    QAbstractItemView,
    QComboBox,
    QDialog,
    QDialogButtonBox,
    QFrame,
    QGridLayout,
    QHBoxLayout,
    QLabel,
    QLineEdit,
    QListWidget,
    QPushButton,
    QScrollArea,
    QSpinBox,
    QStackedWidget,
    QTableWidget,
    QTableWidgetItem,
    QVBoxLayout,
    QWidget,
)

from pms_ui import data
from pms_ui.widgets import GlassCard, MetricCard, RoomStatusCard, SectionTitle, SidebarSection


class AddRoomTypeDialog(QDialog):
    def __init__(self, parent: QWidget | None = None) -> None:
        super().__init__(parent)
        self.setWindowTitle("新增房型")
        self.setModal(True)
        self.resize(460, 520)
        self.setStyleSheet(
            """
            QDialog {
                background-color: #FFFFFF;
            }
            QLabel {
                color: #1E2A3B;
            }
            QLineEdit, QComboBox, QSpinBox {
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
        self._build_ui()

    def _build_ui(self) -> None:
        layout = QVBoxLayout(self)
        layout.setContentsMargins(20, 20, 20, 20)
        layout.setSpacing(14)

        title = QLabel("新增房型")
        title.setStyleSheet("font-size: 20px; font-weight: 700; color: #1E2A3B;")
        hint = QLabel("录入房型基础信息。当前保存为前端临时数据。")
        hint.setProperty("role", "subtle")
        hint.setWordWrap(True)

        self.name_edit = QLineEdit()
        self.name_edit.setPlaceholderText("房型名称")
        self.code_edit = QLineEdit()
        self.code_edit.setPlaceholderText("房型编码")

        self.status_box = QComboBox()
        self.status_box.addItems(["开放预订", "限量预订", "维护中"])

        self.room_count_spin = QSpinBox()
        self.room_count_spin.setRange(1, 999)
        self.room_count_spin.setPrefix("房量 ")
        self.room_count_spin.setValue(10)

        self.price_spin = QSpinBox()
        self.price_spin.setRange(0, 99999)
        self.price_spin.setPrefix("¥")
        self.price_spin.setValue(888)

        self.occupancy_box = QComboBox()
        self.occupancy_box.addItems(["1 人", "2 人", "3 人", "4 人"])

        fields = [
            self.name_edit,
            self.code_edit,
            self.room_count_spin,
            self.price_spin,
            self.occupancy_box,
            self.status_box,
        ]

        buttons = QDialogButtonBox(QDialogButtonBox.StandardButton.Save | QDialogButtonBox.StandardButton.Cancel)
        save_button = buttons.button(QDialogButtonBox.StandardButton.Save)
        cancel_button = buttons.button(QDialogButtonBox.StandardButton.Cancel)
        if save_button is not None:
            save_button.setText("保存")
        if cancel_button is not None:
            cancel_button.setText("取消")
        buttons.accepted.connect(self.accept)
        buttons.rejected.connect(self.reject)

        layout.addWidget(title)
        layout.addWidget(hint)
        for widget in fields:
            layout.addWidget(widget)
        layout.addWidget(buttons)

    def get_payload(self) -> dict[str, object]:
        return {
            "name": self.name_edit.text().strip(),
            "code": self.code_edit.text().strip(),
            "total_rooms": self.room_count_spin.value(),
            "base_price": self.price_spin.value(),
            "occupancy": int(self.occupancy_box.currentText().split()[0]),
            "status": self.status_box.currentText(),
        }


class PMSMainWindow(QWidget):
    logout_requested = pyqtSignal()
    room_type_refresh_requested = pyqtSignal()

    def __init__(self) -> None:
        super().__init__()
        self.room_statuses = data.get_room_statuses()
        self.selected_room = None
        self.room_status_cards: list[RoomStatusCard] = []
        self._build_ui()

    def _build_ui(self) -> None:
        root = QFrame()
        root.setObjectName("rootFrame")
        outer = QVBoxLayout(self)
        outer.setContentsMargins(0, 0, 0, 0)
        outer.addWidget(root)

        layout = QHBoxLayout(root)
        layout.setContentsMargins(24, 20, 24, 20)
        layout.setSpacing(20)

        sidebar = self._build_sidebar()
        self.stack = QStackedWidget()
        self.stack.addWidget(self._build_dashboard_page())
        self.stack.addWidget(self._build_room_type_page())
        self.stack.addWidget(self._build_room_status_page())
        self.stack.addWidget(self._build_reservation_page())
        self.stack.addWidget(self._build_member_list_page())
        self.stack.addWidget(self._build_member_recharge_page())
        self.stack.currentChanged.connect(self._handle_stack_changed)

        layout.addWidget(sidebar, stretch=2)
        layout.addWidget(self.stack, stretch=9)

    def _open_add_room_type_dialog(self) -> None:
        dialog = AddRoomTypeDialog(self)
        if dialog.exec() != QDialog.DialogCode.Accepted:
            return

        payload = dialog.get_payload()
        if not payload["name"] or not payload["code"]:
            return

        self.room_type_table.insertRow(self.room_type_table.rowCount())
        values = [
            str(payload["name"]),
            str(payload["code"]),
            str(payload["total_rooms"]),
            f"¥{payload['base_price']}",
            str(payload["occupancy"]),
            str(payload["status"]),
        ]
        row_index = self.room_type_table.rowCount() - 1
        for col_index, value in enumerate(values):
            self.room_type_table.setItem(row_index, col_index, QTableWidgetItem(value))

    def _reload_room_type_table(self) -> None:
        self.room_type_refresh_requested.emit()

    def _delete_selected_room_type(self) -> None:
        selected_rows = sorted({index.row() for index in self.room_type_table.selectedIndexes()}, reverse=True)
        for row_index in selected_rows:
            self.room_type_table.removeRow(row_index)

    def set_room_type_items(self, items: list[dict[str, object]]) -> None:
        self.room_type_table.setColumnCount(6)
        self.room_type_table.setHorizontalHeaderLabels(["房型", "编码", "房量", "门市价", "入住人数", "状态"])
        self.room_type_table.setColumnHidden(5, True)
        self.room_type_table.setRowCount(0)
        for row_index, item in enumerate(items):
            self.room_type_table.insertRow(row_index)
            values = [
                str(item.get("type_name", "")),
                str(item.get("type_code", "")),
                str(item.get("quantity", "")),
                f"¥{self._format_room_type_price(item.get('base_price'))}",
                str(item.get("max_occupancy", "")),
                self._map_room_type_status(item.get("status")),
            ]
            for col_index, value in enumerate(values):
                self.room_type_table.setItem(row_index, col_index, QTableWidgetItem(value))

    def _handle_stack_changed(self, index: int) -> None:
        if index == 1:
            self.room_type_refresh_requested.emit()

    @staticmethod
    def _map_room_type_status(status: object) -> str:
        if status == 1:
            return "启用"
        if status == 0:
            return "停用"
        return str(status)

    @staticmethod
    def _format_room_type_price(price: object) -> str:
        try:
            amount = float(price) / 100
        except (TypeError, ValueError):
            return ""
        return f"{amount:.2f}"

    def _filter_room_status_cards(self) -> None:
        floor = self.room_status_floor_box.currentText()
        state = self.room_status_state_box.currentText()
        keyword = self.room_status_keyword_edit.text().strip()

        for card in self.room_status_cards:
            room = card.room
            visible = True
            if floor != "全部楼层" and room.floor != floor:
                visible = False
            if state != "全部状态" and room.state != state:
                visible = False
            if keyword:
                haystack = f"{room.room_no} {room.guest} {room.guest_id_no}"
                if keyword not in haystack:
                    visible = False
            card.setVisible(visible)

    def _select_room_status(self, room) -> None:
        self.selected_room = room
        self.room_status_title.setText(f"房间 {room.room_no}")
        self.room_status_desc.setText(f"{room.room_type} | {room.floor} | 当前状态：{room.state}")
        self.room_status_guest_edit.setText(room.guest)
        self.room_status_id_edit.setText(room.guest_id_no)
        self.room_status_note_label.setText(room.stay_label or room.last_action or "暂无动态")
        self.room_status_logs.clear()
        self.room_status_logs.addItems(room.logs or ["暂无房间日志"])

    def _checkin_selected_room(self) -> None:
        if self.selected_room is None:
            return
        guest = self.room_status_guest_edit.text().strip()
        guest_id = self.room_status_id_edit.text().strip()
        if not guest:
            return
        self.selected_room.guest = guest
        self.selected_room.guest_id_no = guest_id
        self.selected_room.state = "在住"
        self.selected_room.stay_label = "刚完成入住登记"
        self.selected_room.last_action = "已录入住客信息"
        self.selected_room.logs.insert(0, f"前台入住登记: {guest}")
        self._refresh_room_status_cards()
        self._select_room_status(self.selected_room)

    def _checkout_selected_room(self) -> None:
        if self.selected_room is None:
            return
        guest_name = self.selected_room.guest or "当前住客"
        self.selected_room.guest = ""
        self.selected_room.guest_id_no = ""
        self.selected_room.state = "待清理"
        self.selected_room.stay_label = f"{guest_name} 已退房，待客房清扫"
        self.selected_room.last_action = "状态切换为待清理"
        self.selected_room.logs.insert(0, f"前台办理退房: {guest_name}")
        self._refresh_room_status_cards()
        self._select_room_status(self.selected_room)

    def _refresh_room_status_cards(self) -> None:
        for card in self.room_status_cards:
            card.refresh()
        self._filter_room_status_cards()

    def _build_sidebar(self) -> QWidget:
        card = GlassCard(panel=True)
        card.setMaximumWidth(260)
        layout = QVBoxLayout(card)
        layout.setContentsMargins(18, 18, 18, 18)
        layout.setSpacing(12)

        brand = QLabel("Asteria PMS")
        brand.setStyleSheet("font-size: 24px; font-weight: 800; color: #F8FAFF;")
        desc = QLabel("Boutique hospitality cockpit")
        desc.setProperty("role", "subtle")
        desc.setStyleSheet("color: #AFC0D8;")

        layout.addWidget(brand)
        layout.addWidget(desc)
        layout.addSpacing(8)

        hotel_section = SidebarSection(
            "酒店管理",
            [
                ("运营总览", lambda: self.stack.setCurrentIndex(0)),
                ("房型控制", lambda: self.stack.setCurrentIndex(1)),
                ("房态中心", lambda: self.stack.setCurrentIndex(2)),
                ("预订中心", lambda: self.stack.setCurrentIndex(3)),
            ],
        )
        member_section = SidebarSection(
            "会员管理",
            [
                ("会员列表", lambda: self.stack.setCurrentIndex(4)),
                ("会员充值设置", lambda: self.stack.setCurrentIndex(5)),
                ("会员等级设置", lambda: self.stack.setCurrentIndex(4)),
                ("会员积分设置", lambda: self.stack.setCurrentIndex(5)),
                ("会员余额变动记录", lambda: self.stack.setCurrentIndex(4)),
            ],
        )
        layout.addWidget(hotel_section)
        layout.addWidget(member_section)

        logout_button = QPushButton("退出登录")
        logout_button.clicked.connect(self.logout_requested.emit)
        layout.addWidget(logout_button)

        notice = GlassCard()
        notice_layout = QVBoxLayout(notice)
        notice_layout.setContentsMargins(16, 16, 16, 16)
        notice_layout.setSpacing(10)
        notice_title = QLabel("值班提醒")
        notice_title.setStyleSheet("font-weight: 700; font-size: 16px; color: #F3F7FF;")
        notice_layout.addWidget(notice_title)
        for alert in data.get_alerts():
            row = QLabel("• " + alert)
            row.setWordWrap(True)
            row.setProperty("role", "subtle")
            row.setStyleSheet("color: #D6E2F3;")
            notice_layout.addWidget(row)

        layout.addStretch()
        layout.addWidget(notice)
        return card

    def _build_dashboard_page(self) -> QWidget:
        page = QWidget()
        layout = QVBoxLayout(page)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(18)

        header = SectionTitle("运营总览", "集中查看经营指标、重点客人和当班提醒。")
        metrics_grid = QGridLayout()
        metrics_grid.setSpacing(16)
        for idx, (title, value, delta) in enumerate(data.get_dashboard_metrics()):
            metrics_grid.addWidget(MetricCard(title, value, delta), idx // 2, idx % 2)

        lower = QHBoxLayout()
        lower.setSpacing(16)

        guests_card = GlassCard()
        guests_layout = QVBoxLayout(guests_card)
        guests_layout.setContentsMargins(18, 18, 18, 18)
        guests_layout.setSpacing(12)
        guests_layout.addWidget(QLabel("重点住客"))
        guest_list = QListWidget()
        guest_list.addItems(
            [
                "周晓晨 | 套房 | VIP 回头客",
                "李晗 | 企业协议 | 今日预抵",
                "王可欣 | 轻奢单人房 | 待支付押金",
            ]
        )
        guests_layout.addWidget(guest_list)

        todo_card = GlassCard()
        todo_layout = QVBoxLayout(todo_card)
        todo_layout.setContentsMargins(18, 18, 18, 18)
        todo_layout.setSpacing(12)
        todo_layout.addWidget(QLabel("班次任务"))
        todo_list = QListWidget()
        todo_list.addItems(
            [
                "核对夜审差异报表",
                "处理 1802 脏房催扫",
                "确认套房超售保护策略",
                "检查 OTA 同步价格",
            ]
        )
        todo_layout.addWidget(todo_list)

        lower.addWidget(guests_card)
        lower.addWidget(todo_card)

        layout.addWidget(header)
        layout.addLayout(metrics_grid)
        layout.addLayout(lower)
        return page

    def _build_room_type_page(self) -> QWidget:
        page = QWidget()
        layout = QVBoxLayout(page)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(18)

        layout.addWidget(SectionTitle("房型控制", "管理房型档案并查看后端同步数据。"))

        top = QHBoxLayout()
        top.setSpacing(16)

        search = QLineEdit()
        search.setPlaceholderText("搜索房型、编码或标签")
        room_area_box = QComboBox()
        room_area_box.addItems(["全部楼区", "行政楼层", "景观楼层", "标准楼层"])
        filter_box = QComboBox()
        filter_box.addItems(["全部状态", "开放预订", "限量预订", "维护中"])
        channel_box = QComboBox()
        channel_box.addItems(["全部渠道", "官网直销", "OTA", "企业协议", "旅行社"])
        top.addWidget(search, stretch=3)
        top.addWidget(filter_box, stretch=1)
        top.addWidget(room_area_box, stretch=1)
        top.addWidget(channel_box, stretch=1)

        table_card = GlassCard()
        table_layout = QVBoxLayout(table_card)
        table_layout.setContentsMargins(18, 18, 18, 18)
        table_layout.setSpacing(12)

        toolbar = QHBoxLayout()
        add_button = QPushButton("新增房型")
        add_button.clicked.connect(self._open_add_room_type_dialog)
        toolbar.addWidget(add_button)
        refresh_button = QPushButton("刷新")
        refresh_button.clicked.connect(self._reload_room_type_table)
        toolbar.addWidget(refresh_button)
        delete_button = QPushButton("删除")
        delete_button.clicked.connect(self._delete_selected_room_type)
        toolbar.addWidget(delete_button)
        toolbar.addStretch()

        self.room_type_table = QTableWidget(0, 6)
        self.room_type_table.setHorizontalHeaderLabels(["房型", "编码", "房量", "门市价", "入住人数", "状态"])
        self.room_type_table.horizontalHeader().setStretchLastSection(True)
        self.room_type_table.verticalHeader().setVisible(False)
        self.room_type_table.setSelectionBehavior(QAbstractItemView.SelectionBehavior.SelectRows)
        self.room_type_table.setEditTriggers(QAbstractItemView.EditTrigger.NoEditTriggers)
        self.room_type_table.setAlternatingRowColors(False)
        self.room_type_table.setColumnHidden(5, True)

        table_layout.addLayout(toolbar)
        table_layout.addWidget(self.room_type_table)

        layout.addLayout(top)
        layout.addWidget(table_card)
        return page

    def _build_room_status_page(self) -> QWidget:
        page = QWidget()
        layout = QVBoxLayout(page)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(18)

        layout.addWidget(SectionTitle("房态中心", "快速浏览楼层状态，便于前台、客房与夜审协作。"))

        filters = QHBoxLayout()
        filters.setSpacing(16)
        self.room_status_floor_box = QComboBox()
        self.room_status_floor_box.addItems(["全部楼层", "18F", "15F", "12F", "10F", "9F"])
        self.room_status_state_box = QComboBox()
        self.room_status_state_box.addItems(["全部状态", "空净", "在住", "预抵", "待清理", "待检", "维修"])
        self.room_status_keyword_edit = QLineEdit()
        self.room_status_keyword_edit.setPlaceholderText("输入楼层、房号、入住人姓名或证件号")
        filters.addWidget(self.room_status_floor_box)
        filters.addWidget(self.room_status_state_box)
        filters.addWidget(self.room_status_keyword_edit, stretch=2)
        filters.addStretch()

        content = QHBoxLayout()
        content.setSpacing(16)

        left_card = GlassCard()
        left_layout = QVBoxLayout(left_card)
        left_layout.setContentsMargins(18, 18, 18, 18)
        left_layout.setSpacing(14)

        scroll = QScrollArea()
        scroll.setWidgetResizable(True)
        scroll.setFrameShape(QFrame.Shape.NoFrame)
        container = QWidget()
        self.room_status_grid = QGridLayout(container)
        self.room_status_grid.setContentsMargins(0, 0, 0, 0)
        self.room_status_grid.setSpacing(14)
        scroll.setWidget(container)

        self.room_status_cards = []
        for index, room in enumerate(self.room_statuses):
            card = RoomStatusCard(room)
            card.clicked.connect(lambda _checked=False, r=room: self._select_room_status(r))
            self.room_status_cards.append(card)
            self.room_status_grid.addWidget(card, index // 3, index % 3)

        left_layout.addWidget(scroll)

        detail_card = GlassCard(panel=True)
        detail_layout = QVBoxLayout(detail_card)
        detail_layout.setContentsMargins(18, 18, 18, 18)
        detail_layout.setSpacing(12)

        self.room_status_title = QLabel("房间详情")
        self.room_status_title.setStyleSheet("font-size: 20px; font-weight: 700; color: #F3F7FF;")
        self.room_status_desc = QLabel("点击左侧卡片查看房间状态与操作入口。")
        self.room_status_desc.setWordWrap(True)
        self.room_status_desc.setStyleSheet("color: #D6E2F3;")
        self.room_status_note_label = QLabel("暂无动态")
        self.room_status_note_label.setWordWrap(True)
        self.room_status_note_label.setStyleSheet("color: #D6E2F3;")
        self.room_status_guest_edit = QLineEdit()
        self.room_status_guest_edit.setPlaceholderText("入住人姓名")
        self.room_status_id_edit = QLineEdit()
        self.room_status_id_edit.setPlaceholderText("入住人证件号")

        action_row = QHBoxLayout()
        checkin_btn = QPushButton("入住")
        checkin_btn.setProperty("variant", "primary")
        checkin_btn.clicked.connect(self._checkin_selected_room)
        checkout_btn = QPushButton("退房")
        checkout_btn.clicked.connect(self._checkout_selected_room)
        action_row.addWidget(checkin_btn)
        action_row.addWidget(checkout_btn)

        log_title = QLabel("房间日志")
        log_title.setStyleSheet("font-size: 16px; font-weight: 700; color: #F3F7FF;")
        self.room_status_logs = QListWidget()

        detail_layout.addWidget(self.room_status_title)
        detail_layout.addWidget(self.room_status_desc)
        detail_layout.addWidget(self.room_status_note_label)
        detail_layout.addWidget(self.room_status_guest_edit)
        detail_layout.addWidget(self.room_status_id_edit)
        detail_layout.addLayout(action_row)
        detail_layout.addWidget(log_title)
        detail_layout.addWidget(self.room_status_logs)
        detail_layout.addStretch()

        content.addWidget(left_card, stretch=7)
        content.addWidget(detail_card, stretch=3)

        self.room_status_floor_box.currentTextChanged.connect(self._filter_room_status_cards)
        self.room_status_state_box.currentTextChanged.connect(self._filter_room_status_cards)
        self.room_status_keyword_edit.textChanged.connect(self._filter_room_status_cards)

        layout.addLayout(filters)
        layout.addLayout(content)
        if self.room_statuses:
            self._select_room_status(self.room_statuses[0])
        return page

    def _build_reservation_page(self) -> QWidget:
        page = QWidget()
        layout = QVBoxLayout(page)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(18)

        layout.addWidget(SectionTitle("预订中心", "整合散客、会员和协议客户的预订信息。"))

        filters = QHBoxLayout()
        filters.setSpacing(16)
        channel_box = QComboBox()
        channel_box.addItems(["全部渠道", "官网直订", "OTA", "企业协议", "前台散客"])
        arrival_box = QComboBox()
        arrival_box.addItems(["抵店状态", "今日预抵", "未来预抵", "已入住", "已取消"])
        payment_box = QComboBox()
        payment_box.addItems(["支付状态", "已付款", "待付款", "担保中", "挂账"])
        filters.addWidget(channel_box)
        filters.addWidget(arrival_box)
        filters.addWidget(payment_box)
        filters.addStretch()

        table = QTableWidget(0, 7)
        table.setHorizontalHeaderLabels(["住客", "房型", "入住", "离店", "渠道", "金额", "状态"])
        table.horizontalHeader().setStretchLastSection(True)
        table.verticalHeader().setVisible(False)
        table.setSelectionBehavior(QAbstractItemView.SelectionBehavior.SelectRows)
        table.setEditTriggers(QAbstractItemView.EditTrigger.NoEditTriggers)
        for row_index, reservation in enumerate(data.get_reservations()):
            table.insertRow(row_index)
            values = [
                reservation.guest_name,
                reservation.room_type,
                reservation.check_in,
                reservation.check_out,
                reservation.channel,
                f"¥{reservation.amount}",
                reservation.status,
            ]
            for col_index, value in enumerate(values):
                table.setItem(row_index, col_index, QTableWidgetItem(value))
        layout.addLayout(filters)
        layout.addWidget(table)
        return page

    def _build_member_list_page(self) -> QWidget:
        page = QWidget()
        layout = QVBoxLayout(page)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(18)

        layout.addWidget(SectionTitle("会员列表", "查看会员档案、等级、余额与积分状态。"))

        filters = QHBoxLayout()
        filters.setSpacing(16)
        keyword_edit = QLineEdit()
        keyword_edit.setPlaceholderText("搜索会员姓名、手机号或会员编号")
        level_box = QComboBox()
        level_box.addItems(["全部等级", "黑金会员", "铂金会员", "金卡会员", "银卡会员"])
        status_box = QComboBox()
        status_box.addItems(["全部状态", "正常", "休眠", "冻结"])
        source_box = QComboBox()
        source_box.addItems(["全部来源", "前台注册", "小程序", "OTA 转化", "协议客户"])
        filters.addWidget(keyword_edit, stretch=3)
        filters.addWidget(level_box)
        filters.addWidget(status_box)
        filters.addWidget(source_box)

        content = QHBoxLayout()
        content.setSpacing(16)

        table_card = GlassCard()
        table_layout = QVBoxLayout(table_card)
        table_layout.setContentsMargins(18, 18, 18, 18)
        table_layout.setSpacing(12)

        member_table = QTableWidget(0, 7)
        member_table.setHorizontalHeaderLabels(["会员号", "姓名", "等级", "手机号", "余额", "积分", "状态"])
        member_table.horizontalHeader().setStretchLastSection(True)
        member_table.verticalHeader().setVisible(False)
        member_table.setSelectionBehavior(QAbstractItemView.SelectionBehavior.SelectRows)
        member_table.setEditTriggers(QAbstractItemView.EditTrigger.NoEditTriggers)
        for row_index, member in enumerate(data.get_members()):
            member_table.insertRow(row_index)
            values = [
                member.member_no,
                member.name,
                member.level,
                member.phone,
                f"¥{member.balance}",
                str(member.points),
                member.status,
            ]
            for col_index, value in enumerate(values):
                member_table.setItem(row_index, col_index, QTableWidgetItem(value))
        table_layout.addWidget(member_table)

        profile_card = GlassCard(panel=True)
        profile_layout = QVBoxLayout(profile_card)
        profile_layout.setContentsMargins(18, 18, 18, 18)
        profile_layout.setSpacing(12)
        profile_layout.addWidget(QLabel("会员档案"))

        member_level_box = QComboBox()
        member_level_box.addItems(["会员等级", "黑金会员", "铂金会员", "金卡会员", "银卡会员"])
        member_tag_box = QComboBox()
        member_tag_box.addItems(["会员标签", "高频住客", "企业客户", "亲子客群", "长住客"])
        rights_box = QComboBox()
        rights_box.addItems(["权益模板", "延迟退房", "双早权益", "升级房型", "停车券"])
        phone_edit = QLineEdit()
        phone_edit.setPlaceholderText("手机号")
        save_btn = QPushButton("保存会员资料")
        save_btn.setProperty("variant", "primary")
        note = QLabel("前端已预留会员档案编辑入口，可对接会员中心、CRM 或储值系统。")
        note.setWordWrap(True)
        note.setProperty("role", "subtle")
        for widget in [phone_edit, member_level_box, member_tag_box, rights_box, save_btn, note]:
            profile_layout.addWidget(widget)
        profile_layout.addStretch()

        content.addWidget(table_card, stretch=7)
        content.addWidget(profile_card, stretch=3)

        layout.addLayout(filters)
        layout.addLayout(content)
        return page

    def _build_member_recharge_page(self) -> QWidget:
        page = QWidget()
        layout = QVBoxLayout(page)
        layout.setContentsMargins(0, 0, 0, 0)
        layout.setSpacing(18)

        layout.addWidget(SectionTitle("会员充值设置", "配置充值档位、赠送规则与适用会员范围。"))

        metrics = QGridLayout()
        metrics.setSpacing(16)
        metrics.addWidget(MetricCard("本月充值笔数", "182", "+15%"), 0, 0)
        metrics.addWidget(MetricCard("本月充值金额", "¥86,300", "+11%"), 0, 1)
        metrics.addWidget(MetricCard("赠送成本", "¥6,880", "+5%"), 1, 0)
        metrics.addWidget(MetricCard("激活率", "73%", "+8%"), 1, 1)

        content = QHBoxLayout()
        content.setSpacing(16)

        table_card = GlassCard()
        table_layout = QVBoxLayout(table_card)
        table_layout.setContentsMargins(18, 18, 18, 18)
        table_layout.setSpacing(12)

        top_filters = QHBoxLayout()
        channel_box = QComboBox()
        channel_box.addItems(["全部渠道", "前台", "小程序", "协议客户", "全渠道"])
        status_box = QComboBox()
        status_box.addItems(["全部状态", "启用", "停用"])
        level_box = QComboBox()
        level_box.addItems(["全部会员", "黑金会员", "铂金会员", "金卡及以上", "企业会员"])
        top_filters.addWidget(channel_box)
        top_filters.addWidget(status_box)
        top_filters.addWidget(level_box)
        top_filters.addStretch()

        recharge_table = QTableWidget(0, 6)
        recharge_table.setHorizontalHeaderLabels(["规则名称", "充值金额", "赠送金额", "适用渠道", "等级限制", "状态"])
        recharge_table.horizontalHeader().setStretchLastSection(True)
        recharge_table.verticalHeader().setVisible(False)
        recharge_table.setSelectionBehavior(QAbstractItemView.SelectionBehavior.SelectRows)
        recharge_table.setEditTriggers(QAbstractItemView.EditTrigger.NoEditTriggers)
        for row_index, rule in enumerate(data.get_recharge_rules()):
            recharge_table.insertRow(row_index)
            values = [
                rule.name,
                f"¥{rule.amount}",
                f"¥{rule.bonus}",
                rule.channel,
                rule.level_limit,
                rule.status,
            ]
            for col_index, value in enumerate(values):
                recharge_table.setItem(row_index, col_index, QTableWidgetItem(value))
        table_layout.addLayout(top_filters)
        table_layout.addWidget(recharge_table)

        setting_card = GlassCard(panel=True)
        setting_layout = QVBoxLayout(setting_card)
        setting_layout.setContentsMargins(18, 18, 18, 18)
        setting_layout.setSpacing(12)
        setting_layout.addWidget(QLabel("充值规则编辑"))

        rule_name = QLineEdit()
        rule_name.setPlaceholderText("规则名称")
        amount_box = QComboBox()
        amount_box.addItems(["充值档位", "¥1000", "¥2000", "¥3000", "¥5000"])
        gift_box = QComboBox()
        gift_box.addItems(["赠送金额", "¥80", "¥188", "¥388", "¥800"])
        apply_channel_box = QComboBox()
        apply_channel_box.addItems(["适用渠道", "前台", "小程序", "协议客户", "全渠道"])
        apply_level_box = QComboBox()
        apply_level_box.addItems(["适用会员", "全部会员", "金卡及以上", "企业会员"])
        validity_box = QComboBox()
        validity_box.addItems(["有效期", "长期有效", "30 天", "90 天", "180 天"])
        save_btn = QPushButton("保存充值规则")
        save_btn.setProperty("variant", "primary")
        note = QLabel("建议后续将该页面对接储值规则 API、营销活动 API 与审核流接口。")
        note.setWordWrap(True)
        note.setProperty("role", "subtle")
        for widget in [rule_name, amount_box, gift_box, apply_channel_box, apply_level_box, validity_box, save_btn, note]:
            setting_layout.addWidget(widget)
        setting_layout.addStretch()

        content.addWidget(table_card, stretch=7)
        content.addWidget(setting_card, stretch=3)

        layout.addLayout(metrics)
        layout.addLayout(content)
        return page
