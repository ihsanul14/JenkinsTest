package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCustomer(c *gin.Context) {
	var customer Customer_Info
	customer.Cifno = c.Request.FormValue("cifno")
	customer.Nama_Sesuai_ID = c.Request.FormValue("nama_sesuai_id")
	result, err := customer.GetCust()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Data Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func createCustomer(c *gin.Context) {
	var customer Customer_Info
	customers := c.BindJSON(&customer)
	id, err := customer.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": customers,
		"data":    id,
	})

}

func getCustomerDetail(c *gin.Context) {

	var customer Customer_Info
	customer.Cifno = c.Request.FormValue("cifno")
	customer.Nama_Sesuai_ID = c.Request.FormValue("nama_sesuai_id")
	result, err := customer.GetCust()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Cifno not Found",
		})
		return
	}
	for _, item := range result {
		if item.Cifno == c.Param("cifno") {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusAccepted, "message": item})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func editCustomer(c *gin.Context) {
	var customer Customer_Info
	id := c.Param("cifno")
	// user. = c.Request.FormValue("password")
	result, err := customer.Update(id)
	if err != nil || result.Cifno == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Check cifno",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "Record was updated",
	})
}
