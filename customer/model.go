package customer

type CustomerData struct {
	CustomerID string `json:"customer_id" bson:"customer_id"`
	Name       string `json:"name"  bson:"name"`
	Lastname   string `json:"lastname"  bson:"lastname"`
}

type ResponseAddCustomer struct {
	Status string `json:"status"`
}

type RequestFindCustomer struct {
	CustomerID string `json:"customer_id"`
}
