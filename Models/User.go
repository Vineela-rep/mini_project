package Models

import (
	"Day4-5/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(product *Product)(err error){
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func PlaceOrder(order *Order) (err error) {
	var product Product
	var prevorder Order

	Config.DB.Where("id = ?", order.CustomerId).Last(&prevorder)

	if prevorder.id != 0{
		if int(time.Now().Unix())- prevorder.OrderTime < 300 {
			err := fmt.Errorf("please try again after %d minutes", 5 - (int(time.Now().Unix()) - prevorder.OrderTime)/60)
			return err
		}
	}

	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	fmt.Println(order)
	Config.DB.Where("product_id = ?", order.ProductId).First(product)
	fmt.Println(product)
	if err = Config.DB.Where("product_id = ?", order.ProductId).First(product).Error; err != nil {
		return err
	}
	//fmt.Println(product)
	if product.Quantity < order.Quantity {
		Config.DB.Model(order).Updates(Order{Status: "Order Failed (Currently out of stock)", Amount: product.Price * float64(order.Quantity)})
		return nil
	}
	Config.DB.Model(product).Update("Quantity", product.Quantity-order.Quantity)
	Config.DB.Model(order).Update(Order{Status: "Order Placed Successfully", Amount: product.Price * float64(order.Quantity)})

	if int(time.Now().UnixNano())- order.OrderTime<300 {
		Config.DB.Model(order).Updates(Order{Status: "Can't Place your order. Please try again after some time", Amount: product.Price * float64(order.Quantity)})
		return nil
	}
	return nil
}

func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func GetUserOrders(orders *[]Order, CustomerId string) (err error){
	if err := Config.DB.Find(&orders,"id = ?", CustomerId).Error; err != nil{
		return err
	}
	return nil
}

func CreateRetailer(retailer *Retailer) (err error) {
	if err = Config.DB.Create(retailer).Error; err != nil {
		return err
	}
	return nil
}

func GetRetailerOrders(orders *[]Order, RetailerId string) (err error){
	if err := Config.DB.Find(&orders,"id = ?", RetailerId).Error; err != nil{
		return err
	}
	return nil
}


