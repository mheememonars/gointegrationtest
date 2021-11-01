package customer

import "log"

func AddCustomer(requestAddCustomer CustomerData) (ResponseAddCustomer, error) {
	log.Println("AddCustomer")
	repo := Repo{}
	_, err := repo.save(requestAddCustomer)
	if err != nil {
		return ResponseAddCustomer{}, err
	}

	return ResponseAddCustomer{Status: "Pass"}, nil
}

func FindCustomer(requestFindCustomer RequestFindCustomer) (CustomerData, error) {
	log.Println("FindCustomer")
	repo := Repo{}
	query := map[string]interface{}{
		"customer_id": requestFindCustomer.CustomerID,
	}
	customerData, err := repo.findOne(query)
	if err != nil {
		return CustomerData{}, err
	}

	return customerData, nil
}
