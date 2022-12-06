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
	ID        int64     `json:"id,omitempty"`
	Tagrp     string    `json:"tagrp,omitempty"`
	Slug      string    `json:"slug,omitempty"`
	PartNo    string    `json:"part_no,omitempty"`
	PartName  string    `json:"part_name,omitempty"`
	LotNo     string    `json:"lot_no,omitempty"`
	LineNo    string    `json:"die_no,omitempty"`
	ReviseNo  string    `json:"revise_no,omitempty"`
	Shelve    string    `json:"shelve,omitempty"`
	PalletNo  string    `json:"pallet_no,omitempty"`
	Qty       float64   `json:"qty,omitempty"`
	Ctn       float64   `json:"ctn,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
