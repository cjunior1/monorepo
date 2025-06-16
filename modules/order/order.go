package order

import "fmt"

func CreateOrder(id string) string {
	return fmt.Sprintf("Order %s created successfull", id)

}
