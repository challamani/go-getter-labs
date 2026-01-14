package fundamentals

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
}

type Customer struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Address []Address `json:"address"`
}

func JsonParse() {
	jsonData := `{
		"id": 1,
		"name": "John Doe",
		"email": "test@example.com",
		"address": [
			{
				"street": "123 Main St",
				"city": "Anytown",
				"state": "CA",
				"zipcode": "12345"
			}
		]
	}`

	var customer Customer
	err := json.Unmarshal([]byte(jsonData), &customer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Customer ID: %d\n", customer.Id)
	fmt.Printf("Customer Name: %s\n", customer.Name)
	fmt.Printf("Customer Email: %s\n", customer.Email)

	for i, addr := range customer.Address {

		fmt.Printf("Address: %d, %s, %s, %s %s\n",
			i,
			addr.Street,
			addr.City,
			addr.State,
			addr.Zipcode)

	}
}

//func main() {
//	JsonParse()
//}
