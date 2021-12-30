package address

import (
	"github.com/volatiletech/sqlboiler/v4/queries"
	"pcode/pkg/db"
	"pcode/pkg/models"
)

type Service struct{}
type Address struct {
	Name                string
	models.PostalCode   `boil:",bind"`
	models.Prefecture   `boil:",bind"`
	models.Municipality `boil:",bind"`
	models.TownArea     `boil:",bind"`
}

func (s Service) GetAddress(postalCode string) (*Address, error) {
	q := `
SELECT postal_codes.number, p.name, m.name, t.name
FROM postal_codes
         LEFT JOIN prefectures p on p.id = postal_codes.prefecture_id
         LEFT JOIN municipalities m on m.id = postal_codes.municipality_id
         LEFT JOIN town_areas t on t.id = postal_codes.town_area_id
WHERE number = ?;
`
	queries.Raw(q, postalCode).Exec(db.GetDB())
	return &Address{}, nil
}
