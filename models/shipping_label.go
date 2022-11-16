package models

type PostShippingLabel struct {
	PalletID string `json:"pallet_id" from:"pallet_id"`
	PartID   string `json:"part_id" from:"part_id"`
	Seq      int    `json:"seq" from:"seq"`
}
