package Routes

import (
	"Day4-5/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/")
	{

		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user", Controllers.GetAllUsers)

		grp1.POST("product", Controllers.AddProduct)
		grp1.POST("product", Controllers.UpdateProduct)
		grp1.GET("product", Controllers.GetAllProducts)

		grp1.POST("order", Controllers.PlaceOrder)
		grp1.GET("order", Controllers.GetAllOrders)
		grp1.GET("order/:customer_id", Controllers.GetUserOrders)

		grp1.POST("retailer", Controllers.CreateRetailer)
		grp1.GET("retailer/:retailer_id", Controllers.GetRetailerOrders)
	}
	return r
}
