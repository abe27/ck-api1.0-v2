package models

import "time"

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
