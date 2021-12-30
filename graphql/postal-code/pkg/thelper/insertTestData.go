package thelper

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"pcode/pkg/db"
	"pcode/pkg/models"
	"testing"
)

func InsertAddressData(t *testing.T) {
	insertHokkaido(t)
	insertTokyo(t)
}

func insertHokkaido(t *testing.T) {
	t.Helper()

	ctx := context.Background()
	d := db.GetDB()

	p := &models.Prefecture{Name: "北海道", NameRoma: "HOKKAIDO"}
	p.Insert(ctx, d, boil.Infer())
	p, _ = models.Prefectures(models.PrefectureWhere.Name.EQ(p.Name)).One(ctx, d)

	m := &models.Municipality{Name: "札幌市　中央区", NameRoma: "SAPPORO SHI CHUO KU", PrefectureID: p.ID}
	m.Insert(ctx, d, boil.Infer())
	m, _ = models.Municipalities(models.MunicipalityWhere.Name.EQ(m.Name)).One(ctx, d)

	town := &models.TownArea{Name: "旭ケ丘", NameRoma: "ASAHIGAOKA", MunicipalityID: m.ID}
	town.Insert(ctx, d, boil.Infer())
	town, _ = models.TownAreas(models.TownAreaWhere.Name.EQ(town.Name), models.TownAreaWhere.MunicipalityID.EQ(m.ID)).One(ctx, d)

	postalCode := models.PostalCode{Number: "0640941", PrefectureID: p.ID, MunicipalityID: m.ID, TownAreaID: town.ID}
	postalCode.Insert(ctx, d, boil.Infer())
}

func insertTokyo(t *testing.T) {
	t.Helper()

	ctx := context.Background()
	d := db.GetDB()

	p := &models.Prefecture{Name: "東京都", NameRoma: "TOKYO TO"}
	p.Insert(ctx, d, boil.Infer())
	p, _ = models.Prefectures(models.PrefectureWhere.Name.EQ(p.Name)).One(ctx, d)

	m := &models.Municipality{Name: "新宿区", NameRoma: "SHINJUKU KU", PrefectureID: p.ID}
	m.Insert(ctx, d, boil.Infer())
	m, _ = models.Municipalities(models.MunicipalityWhere.Name.EQ(m.Name)).One(ctx, d)

	town := &models.TownArea{Name: "四谷", NameRoma: "YOTSUYA", MunicipalityID: m.ID}
	town.Insert(ctx, d, boil.Infer())
	town, _ = models.TownAreas(models.TownAreaWhere.Name.EQ(town.Name), models.TownAreaWhere.MunicipalityID.EQ(m.ID)).One(ctx, d)

	postalCode := models.PostalCode{Number: "1600004", PrefectureID: p.ID, MunicipalityID: m.ID, TownAreaID: town.ID}
	postalCode.Insert(ctx, d, boil.Infer())
}
