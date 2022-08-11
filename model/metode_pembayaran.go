package model

type Payment_method struct {
	Id                int    `json:"Id" form:"Id"`
	Id_pelanggan      int    `json:"Id_pelanggan" form:"Id_pelanggan"`
	Nama_pembayaran   string `json:"Nama_pembayaran" form:"Nama_pembayaran"`
	Metode_pembayaran string `json:"Metode_pembayaran" form:"Metode_pembayaran"`
}
