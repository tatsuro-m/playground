package seed

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
	"pcode/pkg/db"
	"pcode/pkg/models"
	"pcode/pkg/service/address"
	"pcode/pkg/util"
)

var postalCode string
var prefectureName string
var municipalityName string
var townAreaName string
var prefectureNameRome string
var municipalityNameRome string
var townAreaNameRome string

func Exec() error {
	utf8F, _ := os.OpenFile(getCSVPath(), os.O_RDONLY, 0666)
	defer utf8F.Close()
	r := csv.NewReader(transform.NewReader(utf8F, japanese.ShiftJIS.NewDecoder()))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		assign(record)
		insertData()
	}

	return nil
}

var q = `
SELECT postal_codes.id AS 'postal_code.id' ,postal_codes.code AS 'postal_code.code', p.name AS 'prefecture.name', m.name AS 'municipality.name', t.name AS 'town_area.name'
FROM postal_codes
         LEFT JOIN prefectures p on p.id = postal_codes.prefecture_id
         LEFT JOIN municipalities m on m.id = postal_codes.municipality_id
         LEFT JOIN town_areas t on t.id = postal_codes.town_area_id
WHERE code = ? AND p.name = ? AND m.name = ? AND t.name = ?;
`

func insertData() {
	ctx := context.Background()
	d := db.GetDB()

	if checkAlreadyExits() {
		return
	}

	prefecture := models.Prefecture{Name: prefectureName, NameRoma: prefectureNameRome}
	if b, _ := models.Prefectures(models.PrefectureWhere.Name.EQ(prefecture.Name)).Exists(ctx, d); !b {
		prefecture.Insert(ctx, d, boil.Infer())
	}

	p, _ := models.Prefectures(qm.Select(models.PrefectureColumns.ID), models.PrefectureWhere.Name.EQ(prefecture.Name)).One(ctx, d)
	municipality := models.Municipality{Name: municipalityName, NameRoma: municipalityNameRome, PrefectureID: p.ID}
	if b, _ := models.Municipalities(models.MunicipalityWhere.Name.EQ(municipality.Name), models.MunicipalityWhere.PrefectureID.EQ(p.ID)).Exists(ctx, d); !b {
		municipality.Insert(ctx, d, boil.Infer())
	}

	m, _ := models.Municipalities(qm.Select(models.MunicipalityColumns.ID), models.MunicipalityWhere.Name.EQ(municipality.Name), models.MunicipalityWhere.PrefectureID.EQ(p.ID)).One(ctx, d)
	townArea := models.TownArea{Name: townAreaName, NameRoma: townAreaNameRome, MunicipalityID: m.ID}
	townArea.Insert(ctx, d, boil.Infer())

	t, _ := models.TownAreas(qm.Select(models.TownAreaColumns.ID), models.TownAreaWhere.Name.EQ(townArea.Name), models.TownAreaWhere.MunicipalityID.EQ(m.ID)).One(ctx, d)
	pCode := models.PostalCode{Code: postalCode, PrefectureID: p.ID, MunicipalityID: m.ID, TownAreaID: t.ID}
	pCode.Insert(ctx, d, boil.Infer())
}

func checkAlreadyExits() bool {
	ctx := context.Background()
	d := db.GetDB()

	var a address.Address
	queries.Raw(q, postalCode, prefectureName, municipalityName, townAreaName).Bind(ctx, d, &a)
	return a.PostalCode.Code == postalCode && a.Prefecture.Name == prefectureName && a.Municipality.Name == municipalityName && a.TownArea.Name == townAreaName
}

func getCSVPath() string {
	var p string
	if util.IsDev() {
		p, _ = filepath.Abs("./KEN_ALL_ROME.CSV")
		return p
	} else if util.IsTest() {
		p, _ = filepath.Abs("../../internal/seed/testdata/KEN_ALL_ROME_TEST.CSV")
		return p
	}

	return p
}

func assign(csvRow []string) {
	// ex. []string{"8180025", "福岡県", "筑紫野市", "筑紫", "FUKUOKA KEN", "CHIKUSHINO SHI", "CHIKUSHI"}
	postalCode = csvRow[0]
	prefectureName = csvRow[1]
	municipalityName = csvRow[2]
	townAreaName = csvRow[3]
	prefectureNameRome = csvRow[4]
	municipalityNameRome = csvRow[5]
	townAreaNameRome = csvRow[6]
}
