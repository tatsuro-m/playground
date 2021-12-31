package address

import (
	"context"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"pcode/pkg/db"
	"pcode/pkg/models"
	"pcode/pkg/util"
	"strconv"
)

type Service struct{}
type Address struct {
	ID                  string
	Name                string
	models.PostalCode   `boil:",bind"`
	models.Prefecture   `boil:",bind"`
	models.Municipality `boil:",bind"`
	models.TownArea     `boil:",bind"`
}

func (s Service) GetAddress(postalCode string) (*Address, error) {
	q := `
SELECT postal_codes.id AS 'postal_code.id' ,postal_codes.number AS 'postal_code.number', p.name AS 'prefecture.name', m.name AS 'municipality.name', t.name AS 'town_area.name'
FROM postal_codes
         LEFT JOIN prefectures p on p.id = postal_codes.prefecture_id
         LEFT JOIN municipalities m on m.id = postal_codes.municipality_id
         LEFT JOIN town_areas t on t.id = postal_codes.town_area_id
WHERE number = ?;
`
	ctx := context.Background()
	d := db.GetDB()
	var a Address
	err := queries.Raw(q, postalCode).Bind(ctx, d, &a)
	if err != nil {
		return nil, err
	}
	a.ID = strconv.Itoa(a.PostalCode.ID)
	a.Name = fmt.Sprintf("%s%s%s", a.Prefecture.Name, a.Municipality.Name, a.TownArea.Name)
	a.Name = util.TrimFullWidthSpace(a.Name)

	return &a, nil
}
