from __future__ import annotations


APP_STYLE = """
QWidget {
    background: transparent;
    color: #1E2A3B;
    font-family: "Microsoft YaHei UI", "Segoe UI";
    font-size: 14px;
}
QMainWindow, QFrame#rootFrame {
    background-color: #F4F7FB;
}
QLabel[role="eyebrow"] {
    color: #6E7D93;
    font-size: 12px;
    font-weight: 600;
    letter-spacing: 1px;
    text-transform: uppercase;
}
QLabel[role="headline"] {
    font-size: 34px;
    font-weight: 700;
    color: #142033;
}
QLabel[role="subtle"] {
    color: #71819A;
    font-size: 13px;
}
QFrame[card="true"] {
    background: #FFFFFF;
    border: 1px solid #E5EBF3;
    border-radius: 22px;
}
QFrame[panel="true"] {
    background: #223248;
    border: 1px solid #2F435D;
    border-radius: 18px;
}
QPushButton {
    background: #FFFFFF;
    border: 1px solid #D7E0EB;
    border-radius: 12px;
    padding: 10px 16px;
    color: #203047;
    font-weight: 600;
}
QPushButton:hover {
    background: #F6F9FC;
}
QPushButton[variant="primary"] {
    background: #2F7AF8;
    color: white;
    border: none;
}
QPushButton[variant="ghost"] {
    background: transparent;
    color: #E8F0FF;
    border: 1px solid transparent;
    text-align: left;
    padding: 12px 14px;
}
QPushButton[variant="ghost"]:hover {
    background: rgba(255, 255, 255, 0.08);
}
QLineEdit, QComboBox, QDateEdit, QSpinBox {
    background: #FFFFFF;
    border: 1px solid #D8E1EC;
    border-radius: 12px;
    padding: 10px 12px;
    color: #1F2E42;
    selection-background-color: #A9C9FF;
}
QComboBox QAbstractItemView {
    background: #FFFFFF;
    color: #1F2E42;
    border: 1px solid #D8E1EC;
    selection-background-color: #EAF2FF;
    selection-color: #1F2E42;
}
QLineEdit:focus, QComboBox:focus, QDateEdit:focus, QSpinBox:focus {
    border: 1px solid #5A92F2;
}
QListWidget, QTableWidget {
    background: #FFFFFF;
    border: 1px solid #E1E8F0;
    border-radius: 14px;
    gridline-color: #EDF2F7;
}
QHeaderView::section {
    background: #F6F9FC;
    color: #5C6D84;
    border: none;
    padding: 10px 8px;
    font-weight: 600;
}
QTableWidget::item {
    padding: 9px 8px;
}
QTableWidget::item:selected {
    background: #EEF1F4;
    color: #1E2A3B;
}
QListWidget::item {
    padding: 10px 12px;
    margin: 4px;
    border-radius: 10px;
}
QListWidget::item:selected {
    background: #EAF2FF;
}
QScrollBar:vertical {
    border: none;
    background: transparent;
    width: 10px;
    margin: 6px 0;
}
QScrollBar::handle:vertical {
    background: rgba(112, 132, 161, 0.28);
    border-radius: 5px;
}
QScrollBar::add-line:vertical, QScrollBar::sub-line:vertical {
    height: 0;
}
"""
