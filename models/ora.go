package models

import "time"

type OraResponseStock struct {
	Message string       `json:"message,omitempty"`
	Data    []StockCheck `json:"data,omitempty"`
	Code    int          `json:"code,omitempty"`
	Success bool         `json:"success,omitempty"`
	Error   string       `json:"error,omitempty"`
}

type StockCheck struct {
	Tagrp      string    `json:"tagrp"` // TAGRP
	Slug       string    `json:"slug"`
	PartNo     string    `json:"partno"`      // PARTNO
	PartName   string    `json:"partname"`    // PARTNAME
	Total      int64     `json:"total"`       // TOTAL
	Checked    int64     `json:"checked"`     // CHECKED
	NotCheck   int64     `json:"notcheck"`    // NOTCHECK
	LastUpdate time.Time `json:"last_update"` // LASTUPDATE
}

type OraResponseStockDetail struct {
	Message string             `json:"message,omitempty"`
	Data    []StockCheckDetail `json:"data,omitempty"`
	Code    int                `json:"code,omitempty"`
	Success bool               `json:"success,omitempty"`
	Error   string             `json:"error,omitempty"`
}

type StockCheckDetail struct {
	Tagrp      string    `json:"tagrp"` //TAGRP
	Slug       string    `json:"slug"`
	PartNo     string    `json:"partno"` //PARTNO
	PartName   string    `json:"partname"`
	LotNo      string    `json:"lot_no"`    // LOTNO
	SerialNo   string    `json:"serial_no"` // RUNNINGNO
	Qty        int64     `json:"qty"`       // STOCKQUANTITY
	Shelve     string    `json:"shelve"`    // SHELVE
	PalletNo   string    `json:"pallet_no"` // PALLETKEY
	LineNo     string    `json:"line_no"`
	ReviseNo   string    `json:"revise_no"`
	Checked    int64     `json:"checked_flg"` // STKTAKECHKFLG
	LastUpdate time.Time `json:"last_update"` // UPDDTE
}
type OraResponse struct {
	Message string     `json:"message,omitempty"`
	Data    []OraStock `json:"data,omitempty"`
	Code    int        `json:"code,omitempty"`
	Success bool       `json:"success,omitempty"`
	Error   string     `json:"error,omitempty"`
}

type OraStock struct {
	ID        int64     `json:"id"`
	Tagrp     string    `json:"tagrp"`
	Slug      string    `json:"slug"`
	PartNo    string    `json:"part_no"`
	PartName  string    `json:"part_name"`
	SerialNo  string    `json:"serial_no"`
	LotNo     string    `json:"lot_no"`
	LineNo    string    `json:"die_no"`
	ReviseNo  string    `json:"revise_no"`
	Shelve    string    `json:"shelve"`
	PalletNo  string    `json:"pallet_no"`
	Qty       float64   `json:"qty"`
	Ctn       float64   `json:"ctn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FrmUpdateStock struct {
	Shelve string `form:"shelve" json:"shelve"`
	Ctn    int64  `form:"ctn" json:"ctn"`
}

type UpdateStockData struct {
	Data interface{} `json:"data"`
}
