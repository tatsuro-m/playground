package postalcode

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"pcode/pkg/db"
	"pcode/pkg/models"
	"pcode/pkg/util"
)

type Service struct{}

func (s Service) GetOne(address string) (*models.PostalCode, error) {
	p, err := models.PostalCodes(
		qm.LeftOuterJoin("prefectures p on p.id = postal_codes.prefecture_id"),
		qm.LeftOuterJoin("municipalities m on m.id = postal_codes.municipality_id"),
		qm.LeftOuterJoin("town_areas t on t.id = postal_codes.town_area_id"),
		qm.Where("CONCAT(p.name,m.name,t.name) = ?", util.TrimFullWidthSpace(address)),
	).One(context.Background(), db.GetDB())

	if err != nil {
		return nil, err
	}

	return p, nil
}
