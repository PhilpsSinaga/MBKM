package model

import "time"

type Membership struct {
	Id                 int    `JSON:"Id" form:"Id"`
	Id_pengguna        int    `JSON:"Id_pengguna" form:"Id_pengguna"`
	Kategori_langganan string `JSON:"Kategori_langganan" form:"Kategori_langganan"`
	Harga              string `JSON:"Harga" form:"Harga"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	ExpiredAt          time.Time
}
