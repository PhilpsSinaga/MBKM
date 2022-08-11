package model

type Class struct {
	Id           int    `JSON:"Id" form:"Id"`
	Id_pelatih   int    `JSON:"Id_pelatih" form:"Id_pelatih"`
	Id_pengguna  int    `JSON:"Id_pengguna" form:"Id_pengguna"`
	Tanggal      string `JSON:"Tanggal" form:"Tanggal"`
	NamaKelas    string `JSON:"NamaKelas" form:"NamaKelas"`
	Nama_pelatih string `JSON:"Nama_pelatih" form:"Nama_pelatih"`
}
