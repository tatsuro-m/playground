package address

import "fmt"

type Service struct{}
type Address struct {
	ID   int
	Name string
}

func (s Service) GetAddress(postalCode string) (*Address, error) {
	fmt.Println(postalCode)
	return &Address{}, nil
}
