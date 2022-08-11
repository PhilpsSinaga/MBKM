package model

type Admins struct {
	Id            int    `json:"Id" form:"Id"`
	Id_pelatih    int    `json:"Id_pelatih" form:"Id_pelatih"`
	Id_pengguna   int    `json:"Id_pengguna" form:"Id_pengguna"`
	Peran         string `json:"Peran" form:"Peran"`
	Nama_petugas  string `json:"Nama_petugas" form:"Nama_petugas"`
	Jenis_kelamin string `json:"Jenis_kelamin" form:"Jenis_kelamin"`
	Alamat        string `json:"Alamat" form:"Alamat"`
	Email         string `json:"Email" form:"Email"`
	Password      string `json:"Password" form:"Password"`
	Tgl_lahir     string `json:"Tgl_lahir" form:"Tgl_lahir"`
}
