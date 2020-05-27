package main

type GeneralRequest struct {
	Cifno string `json:"cifno"`
}

type ResultCustomer struct {
	Data []Customer_Info `json:"data"`
}

type Customer_Info struct {
	Cifno                 string `json:"cifno" gorm:"column:Cifno;primary_key"`
	Nama_Sesuai_ID        string `json:"nama_sesuai_id" gorm:"column:Nama_Sesuai_ID"`
	Tipe_Nasabah          string `json:"tipe_nasabah" gorm:"column:Tipe_Nasabah"`
	Tipe_Nasabah_Desc     string `json:"tipe_nasabah_desc" gorm:"column:Tipe_Nasabah_Desc"`
	Jenis_Kelamin         string `json:"jenis_kelamin" gorm:"column:Jenis_Kelamin"`
	Kewarganegaraan       string `json:"kewarganegaraan" gorm:"column:Kewarganegaraan"`
	Kode_Pendidikan       string `json:"kode_pendidikan" gorm:"column:Kode_Pendidikan"`
	Pendidikan            string `json:"pendidikan" gorm:"column:Pendidikan"`
	Kode_Jenis_Pekerjaan  string `json:"kode_jenis_pekerjaan" gorm:"column:Kode_Jenis_Pekerjaan"`
	Jenis_Pekerjaan       string `json:"jenis_pekerjaan" gorm:"column:Jenis_Pekerjaan"`
	Nama_Kantor           string `json:"nama_kantor" gorm:"column:Nama_Kantor"`
	Kode_Bidang_Pekerjaan string `json:"kode_bidang_pekerjaan" gorm:"column:Kode_Bidang_Pekerjaan"`
	Bidang_Pekerjaan      string `json:"bidang_pekerjaan" gorm:"column:Bidang_Pekerjaan"`
	Kode_Jabatan          string `json:"kode_jabatan" gorm:"column:Kode_Jabatan"`
	Jabatan               string `json:"jabatan" gorm:"column:Jabatan"`
}
