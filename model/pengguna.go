package model

type User struct {
	Id            int    `json:"Id" form:"Id"`
	Nama          string `json:"Nama" form:"Nama"`
	Jenis_kelamin string `json:"Jenis_kelamin" form:"Jenis_kelamin"`
	Alamat        string `json:"Alamat" form:"Alamat"`
	Tgl_lahir     string `json:"Tgl_lahir" form:"Tgl_lahir"`
	Jenis_member  string `json:"Jenis_member" form:"Jenis_member"`
	Status        string `json:"Status" form:"Status"`
	Email         string `json:"Email" form:"Email"`
	Username      string `json:"Username" form:"Username"`
	Password      string `json:"Password" form:"Password"`
	No_hp         string `json:"No_hp" form:"No_hp"`
}
