package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	InitDb()
	router := InitRouter()
	router.Run(":8091")

}

// =========================== all code one page =========================

// var db *gorm.DB

// func init() {
// 	var err error
// 	db, err = gorm.Open("mysql", "root:12345@tcp(localhost:3306)/mydb")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	db.AutoMigrate(&Customer_Info{})
// }

// router := gin.Default()
// router.Use(cors.Default())
// router.Use(apmgin.Middleware(router))
// router.GET("/customerinfo", getCustomer)
// router.POST("/customerinfo", createCustomer)
// router.GET("/customerinfo/:cifno", getCustomerDetail)
// router.PUT("/customerinfo/:cifno", editCustomer)
// router.Run(":8090")

// type Customer_Info struct {
// 	Cifno                 string `json:"cifno" gorm:"column:Cifno;primary_key"`
// 	Nama_Sesuai_ID        string `json:"nama_sesuai_id" gorm:"column:Nama_Sesuai_ID"`
// 	Tipe_Nasabah          string `json:"tipe_nasabah" gorm:"column:Tipe_Nasabah"`
// 	Tipe_Nasabah_Desc     string `json:"tipe_nasabah_desc" gorm:"column:Tipe_Nasabah_Desc"`
// 	Jenis_Kelamin         string `json:"jenis_kelamin" gorm:"column:Jenis_Kelamin"`
// 	Kewarganegaraan       string `json:"kewarganegaraan" gorm:"column:Kewarganegaraan"`
// 	Kode_Pendidikan       string `json:"kode_pendidikan" gorm:"column:Kode_Pendidikan"`
// 	Pendidikan            string `json:"pendidikan" gorm:"column:Pendidikan"`
// 	Kode_Jenis_Pekerjaan  string `json:"kode_jenis_pekerjaan" gorm:"column:Kode_Jenis_Pekerjaan"`
// 	Jenis_Pekerjaan       string `json:"jenis_pekerjaan" gorm:"column:Jenis_Pekerjaan"`
// 	Nama_Kantor           string `json:"nama_kantor" gorm:"column:Nama_Kantor"`
// 	Kode_Bidang_Pekerjaan string `json:"kode_bidang_pekerjaan" gorm:"column:Kode_Bidang_Pekerjaan"`
// 	Bidang_Pekerjaan      string `json:"bidang_pekerjaan" gorm:"column:Bidang_Pekerjaan"`
// 	Kode_Jabatan          string `json:"kode_jabatan" gorm:"column:Kode_Jabatan"`
// 	Jabatan               string `json:"jabatan" gorm:"column:Jabatan"`
// }

// func createCustomer(c *gin.Context) {
// 	var customer Customer_Info
// 	var custID []Customer_Info
// 	c.BindJSON(&customer)
// 	db.Find(&custID)
// 	for _, item := range custID {
// 		if customer.Cifno == item.Cifno {
// 			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Data dengan CIFNO tersebut sudah tersedia"})
// 			return
// 		}
// 	}
// 	db.Create(&customer)
// 	c.JSON(200, customer)
// }

// func getCustomer(c *gin.Context) {
// 	var customer []Customer_Info
// 	db.Find(&customer)

// 	if len(customer) <= 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Data tidak Ada!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": customer})
// }

// func getCustomerDetail(c *gin.Context) {
// 	var customer Customer_Info
// 	var customerID []Customer_Info
// 	// customerID := c.Param("cifno")
// 	var custID = c.Param("cifno")
// 	db.Find(&customerID)
// 	c.BindJSON(&customer)
// 	for _, item := range customerID {
// 		if item.Cifno == custID {
// 			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": item})
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": customer})
// }

// func editCustomer(c *gin.Context) {
// 	var customer Customer_Info
// 	var customerID []Customer_Info
// 	id := c.Params.ByName("cifno")
// 	db.Find(&customerID)
// 	if err := db.Where("cifno = ?", id).First(&customerID).Error; err != nil {
// 		c.AbortWithStatus(404)
// 	}
// 	c.BindJSON(&customer)
// 	db.Save(&customer)
// 	c.JSON(200, customer)
// }
