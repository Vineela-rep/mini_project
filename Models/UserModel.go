package Models

type User struct {
	id int `json:"id"`
	CustomerId      string   `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}
type Product struct {
	id int `json:"id"`
	ProductId      string   `json:"product_id"`
	ProductName    string `json:"Product_name"`
	Price   float64 `json:"price"`
	Quantity   int `json:"quantity"`
	Message string `json:"message"`
	RetailerId      string   `json:"retailer_id"`
}
type Order struct {
	id int `json:"id"`
	OrderId      string   `json:"order_id"`
	CustomerId      string   `json:"customer_id"`
	ProductId      string   `json:"product_id"`
	Quantity   int `json:"quantity"`
	Amount float64 `json:"amount"`
	Status string `json:"status"`
	OrderTime int `json:"order_time"`
	RetailerId      string   `json:"retailer_id"`
}

type Retailer struct {
	id      int   `json:"id"`
	RetailerId      string   `json:"retailer_id"`
	RetailerName    string `json:"retailer_name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`

}
func (u *User) TableName() string {
	return "user"
}
func (p *Product) TableName() string {
	return "product"
}
func (o *Order) TableName() string {
	return "order"
}
func (o *Retailer) TableName() string {
	return "retailer"
}
