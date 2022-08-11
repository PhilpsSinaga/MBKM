package model

type Trainer struct {
	Id            int    `JSON:"Id" form:"Id"`
	Nama          string `JSON:"Nama" form:"Nama"`
	Tgl_lahir     string `JSON:"Tgl_lahir" form:"Tgl_lahir"`
	Jenis_kelamin string `JSON:"Jenis_kelamin" form:"Jenis_kelamin"`
	Alamat        string `JSON:"Alamat" form:"Alamat"`
	Email         string `JSON:"Email" form:"Email"`
	Keahlihan     string `JSON:"Keahlihan" form:"Keahlihan"`
}
