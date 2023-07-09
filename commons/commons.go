package commons

// Define struct for message
type KafkaMessage struct {
	Category string `json:"category"`
	AppName  string `json:"appName"`
	Message  string `json:"message"`
}

// Category enum
const (
	Restaurant     = "RESTAURANT"
	GroceryStore   = "GROCERYSTORE"
	ShoppingMall   = "SHOPPINGMALL"
	OnlineShopping = "ONLINESHOPPING"
)

// App enum
const (
	PhonePe = "PHONEPE"
	Paytm   = "PAYTM"
	GPay    = "GPAY"
	Card    = "CARD"
)
