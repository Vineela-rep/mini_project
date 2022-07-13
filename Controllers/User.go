package Controllers

import (
	"Day4-5/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


//GetUsers ... Get all users
func GetAllUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Enter valid details",
		})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetAllProducts ... Get all the products that are available
func GetAllProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//AddProduct ... Adding new product
func AddProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.AddProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Enter valid details",
		})
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//UpdateProduct ... Update the details of a product with given product_id
func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//PlaceOrder ...
func PlaceOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.PlaceOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//GetAllOrders ... Get all orders that are placed
func GetAllOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

//GetUserOrders ... Get all orders of a customer with given customer_id
func GetUserOrders(c *gin.Context){
	var orders []Models.Order
	customerId := c.Params.ByName("customer_id")
	err := Models.GetUserOrders(&orders, customerId )
	if err != nil{
		//fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		c.AbortWithStatus(http.StatusNotFound)
	} else{
		c.JSON(http.StatusOK, orders)
	}
}

func CreateRetailer(c *gin.Context) {
	var retailer Models.Retailer
	c.BindJSON(&retailer)
	err := Models.CreateRetailer(&retailer)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Enter valid details",
		})
	} else {
		c.JSON(http.StatusOK, retailer)
	}
}

func GetRetailerOrders(c *gin.Context){
	var orders []Models.Order
	retailerId := c.Params.ByName("customer_id")
	err := Models.GetRetailerOrders(&orders, retailerId )
	if err != nil{
		//fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Retailer not found"})
		c.AbortWithStatus(http.StatusNotFound)
	} else{
		c.JSON(http.StatusOK, orders)
	}
}