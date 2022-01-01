package seed

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
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

type Record struct {
	postalCode           string
	prefectureName       string
	municipalityName     string
	townAreaName         string
	prefectureNameRome   string
	municipalityNameRome string
	townAreaNameRome     string
}

var record Record

var rowPrefecture *models.Prefecture
var rowMunicipality *models.Municipality
var rowTownArea *models.TownArea

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

		assignCSVDataToVar(record)
		insertData()
	}

	return nil
}

var q = `
SELECT postal_codes.id AS 'postal_code.id',postal_codes.code AS 'postal_code.code', p.name AS 'prefecture.name', m.name AS 'municipality.name', t.name AS 'town_area.name'
FROM postal_codes
         LEFT JOIN prefectures p on p.id = postal_codes.prefecture_id
         LEFT JOIN municipalities m on m.id = postal_codes.municipality_id
         LEFT JOIN town_areas t on t.id = postal_codes.town_area_id
WHERE code = ? AND p.name = ? AND m.name = ? AND t.name = ?;
`

func insertData() {
	if checkAlreadyExits() {
		return
	}

	err := insertPrefecture()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = insertMunicipality()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = insertTownArea()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = insertPostalCode()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func insertPrefecture() error {
	ctx := context.Background()
	d := db.GetDB()

	prefecture, err := models.Prefectures(models.PrefectureWhere.Name.EQ(record.prefectureName)).One(ctx, d)
	if err != nil {
		p := models.Prefecture{Name: record.prefectureName, NameRoma: record.prefectureNameRome}
		err = p.Insert(context.Background(), db.GetDB(), boil.Infer())
		if err != nil {
			return err
		}

		rowPrefecture = &p
	} else {
		rowPrefecture = prefecture
	}

	return nil
}

func insertMunicipality() error {
	ctx := context.Background()
	d := db.GetDB()

	municipality, err := models.Municipalities(models.MunicipalityWhere.Name.EQ(record.municipalityName), models.MunicipalityWhere.PrefectureID.EQ(rowPrefecture.ID)).One(ctx, d)
	if err != nil {
		m := models.Municipality{Name: record.municipalityName, NameRoma: record.municipalityNameRome, PrefectureID: rowPrefecture.ID}
		err = m.Insert(ctx, d, boil.Infer())
		if err != nil {
			return err
		}

		rowMunicipality = &m
	} else {
		rowMunicipality = municipality
	}

	return nil
}

func insertTownArea() error {
	ctx := context.Background()
	d := db.GetDB()

	townArea := models.TownArea{Name: record.townAreaName, NameRoma: record.townAreaNameRome, MunicipalityID: rowMunicipality.ID}
	err := townArea.Insert(ctx, d, boil.Infer())
	if err != nil {
		t, err := models.TownAreas(models.TownAreaWhere.Name.EQ(record.townAreaName), models.TownAreaWhere.Name.EQ(record.townAreaNameRome), models.TownAreaWhere.MunicipalityID.EQ(rowMunicipality.ID)).One(ctx, d)
		if err != nil {
			return err
		}
		rowTownArea = t
	} else {
		rowTownArea = &townArea
	}

	return nil
}

func insertPostalCode() error {
	ctx := context.Background()
	d := db.GetDB()

	pCode := models.PostalCode{Code: record.postalCode, PrefectureID: rowPrefecture.ID, MunicipalityID: rowMunicipality.ID, TownAreaID: rowTownArea.ID}
	err := pCode.Insert(ctx, d, boil.Infer())
	if err != nil {
		_, err := models.PostalCodes(models.PostalCodeWhere.Code.EQ(record.postalCode),
			models.PostalCodeWhere.PrefectureID.EQ(rowPrefecture.ID),
			models.PostalCodeWhere.MunicipalityID.EQ(rowMunicipality.ID),
			models.PostalCodeWhere.TownAreaID.EQ(rowTownArea.ID),
		).One(ctx, d)

		if err != nil {
			return err
		}
	}

	return nil
}

func checkAlreadyExits() bool {
	ctx := context.Background()
	d := db.GetDB()

	var a address.Address
	queries.Raw(q, record.postalCode, record.prefectureName, record.municipalityName, record.townAreaName).Bind(ctx, d, &a)
	return a.PostalCode.Code == record.postalCode && a.Prefecture.Name == record.prefectureName && a.Municipality.Name == record.municipalityName && a.TownArea.Name == record.townAreaName
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

func assignCSVDataToVar(csvRow []string) {
	// ex. []string{"8180025", "福岡県", "筑紫野市", "筑紫", "FUKUOKA KEN", "CHIKUSHINO SHI", "CHIKUSHI"}
	record.postalCode = csvRow[0]
	record.prefectureName = csvRow[1]
	record.municipalityName = csvRow[2]
	record.townAreaName = csvRow[3]
	record.prefectureNameRome = csvRow[4]
	record.municipalityNameRome = csvRow[5]
	record.townAreaNameRome = csvRow[6]
}
