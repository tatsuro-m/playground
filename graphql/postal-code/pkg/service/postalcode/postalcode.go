package postalcode

import "pcode/pkg/models"

type Service struct{}

func (s Service) GetOne(address string) (*models.PostalCode, error) {
	return nil, nil
}
