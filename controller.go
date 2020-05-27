package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func (customer *Customer_Info) GetCust() (customers []Customer_Info, err error) {
	if err = db.Find(&customers).Error; err != nil {
		return
	}
	return
}

func (customer Customer_Info) Insert() (id string, err error) {
	result := db.Create(&customer)
	id = customer.Cifno
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (customer *Customer_Info) Update(id string) (updateCustomer Customer_Info, err error) {

	if err = db.Select([]string{"id", "username"}).First(&updateCustomer, id).Error; err != nil {
		return
	}
	if err = db.Model(&updateCustomer).Updates(&customer).Error; err != nil {
		return
	}
	return
}

// ====================================================================================
// func createCustomer(c *gin.Context) {
// 	var user Customer_Info
// 	user.Cifno = c.Request.FormValue("cifno")
// 	user.Nama_Sesuai_ID = c.Request.FormValue("nama_sesuai_id")
// 	id, err := user.Insert()

// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    -1,
// 			"message": "添加失败",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":    1,
// 		"message": "添加成功",
// 		"data":    id,
// 	})

// 	// var customer Customer_Info
// 	// var custID []Customer_Info
// 	// c.BindJSON(&customer)
// 	// db.Find(&custID)
// 	// for _, item := range custID {
// 	// 	if customer.Cifno == item.Cifno {
// 	// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Data dengan CIFNO tersebut sudah tersedia"})
// 	// 		return
// 	// 	}
// 	// }
// 	// db.Create(&customer)
// 	// c.JSON(200, customer)
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

// func getCustomer(c *gin.Context) {
// 	var user Customer_Info
// 	user.Cifno = c.Request.FormValue("cifno")
// 	user.Nama_Sesuai_ID = c.Request.FormValue("nama_sesuai_id")
// 	result, err := user.Users()

// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    -1,
// 			"message": "抱歉未找到相关信息",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 1,
// 		"data": result,
// 	})
// }

// func (cust *Customer_Info) Customer() (customers []Customer_Info, err error) {
// var db *sql.DB

// rows, err := db.Query("SELECT cifno, nama_sesuai_id, kewarganegaraan FROM customer_infos")
// if err != nil {
// 	return
// }
// for rows.Next() {
// 	var customer Customer_Info
// 	rows.Scan(&customer.Cifno, &customer.Nama_Sesuai_ID, &customer.Kewarganegaraan)
// 	customers = append(customers, customer)
// }
// defer rows.Close()
// return

// var customers Customer_Info
// var customer []Customer_Info = []Customer_Info{}
// var result ResultCustomer = ResultCustomer{}
// c.BindJSON(&request)
// var db, err = sql.Open("mysql", "root:A123b456c@tcp(localhost:3306)/mydb")

// if err != nil {
// 	fmt.Println(err.Error())
// 	return result, err
// }
// err = db.Ping()

// //Get Data From Phoenix Table
// rows, err := db.Query("Select * from customer_infos")

// if err != nil {
// 	fmt.Println(err.Error())
// 	return result, err
// }

// for rows.Next() {
// 	if err := rows.Scan(&customers.Cifno); err != nil {
// 		log.Fatal(err.Error())

// 	} else {
// 		customer = append(customer, customers)
// 	}
// }

// result.Data = customer

// return result, nil

// var customer []Customer_Info
// // geallcustomer(customer)
// // a := c.BindJSON(&request)
// db.Find(&customer)
// // if len(customer) <= 0 {
// // 	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Data tidak Ada!"})
// // 	return
// // }
// // c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": customer})
// return customers

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
