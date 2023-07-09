package request

// struct for api request to process recived message
type ProcessMessageRequest struct {
	Message string `json:"message"`
	Category string `json:"category"`
	AppName string `json:"appName"`
	TransactionDate string `json:"transactionDate"`	
}