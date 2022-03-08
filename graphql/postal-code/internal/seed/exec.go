package seed

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
	"pcode/pkg/db"
	"pcode/pkg/models"
	"pcode/pkg/util"
)

type Row struct {
	postalCode           string
	prefectureName       string
	municipalityName     string
	townAreaName         string
	prefectureNameRome   string
	municipalityNameRome string
	townAreaNameRome     string
}

// DB PK を保存する構造体
type Pk struct {
	prefecture   int
	municipality int
	townArea     int
}

var previousRow Row
var currentRow Row
var pk Pk

var rowPrefecture *models.Prefecture
var rowMunicipality *models.Municipality
var rowTownArea *models.TownArea

func Exec() error {
	utf8F, _ := os.Open(getCSVPath())
	defer utf8F.Close()
	r := csv.NewReader(transform.NewReader(utf8F, japanese.ShiftJIS.NewDecoder()))

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		assignCSVDataToCurrentRow(row)
		insertData()
		assignCSVDataToPreviousRow(row)
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
	id, err := insertPrefecture()
	if err != nil {
		fmt.Println(err)
		return
	}
	pk.prefecture = id

	id, err = insertMunicipality()
	if err != nil {
		fmt.Println(err)
		return
	}
	pk.municipality = id

	id, err = insertTownArea()
	if err != nil {
		fmt.Println(err)
		return
	}
	pk.townArea = id

	err = insertPostalCode()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("pk: ", pk)
}

func insertPrefecture() (int, error) {
	ctx := context.Background()
	d := db.GetDB()

	id := 0
	if previousRow.prefectureName != currentRow.prefectureName {
		// 一致していない時だけ INSERT する
		p := models.Prefecture{Name: currentRow.prefectureName, NameRoma: currentRow.prefectureNameRome}
		err := p.Insert(context.Background(), db.GetDB(), boil.Infer())
		if err != nil {
			return 0, err
		}

		id = p.ID
		rowPrefecture = &p
	}

	prefecture, err := models.Prefectures(models.PrefectureWhere.Name.EQ(currentRow.prefectureName)).One(ctx, d)
	if err != nil {
		p := models.Prefecture{Name: currentRow.prefectureName, NameRoma: currentRow.prefectureNameRome}
		err = p.Insert(context.Background(), db.GetDB(), boil.Infer())
		if err != nil {
			return 0, err
		}

		id = p.ID
		rowPrefecture = &p
	} else {
		id = prefecture.ID
		rowPrefecture = prefecture
	}

	return id, nil
}

func insertMunicipality() (int, error) {
	ctx := context.Background()
	d := db.GetDB()

	id := 0
	municipality, err := models.Municipalities(models.MunicipalityWhere.Name.EQ(currentRow.municipalityName), models.MunicipalityWhere.PrefectureID.EQ(rowPrefecture.ID)).One(ctx, d)
	if err != nil {
		m := models.Municipality{Name: currentRow.municipalityName, NameRoma: currentRow.municipalityNameRome, PrefectureID: rowPrefecture.ID}
		err = m.Insert(ctx, d, boil.Infer())
		if err != nil {
			return 0, err
		}

		id = m.ID
		rowMunicipality = &m
	} else {
		id = municipality.ID
		rowMunicipality = municipality
	}

	return id, nil
}

func insertTownArea() (int, error) {
	ctx := context.Background()
	d := db.GetDB()

	id := 0
	townArea := models.TownArea{Name: currentRow.townAreaName, NameRoma: currentRow.townAreaNameRome, MunicipalityID: rowMunicipality.ID}
	err := townArea.Insert(ctx, d, boil.Infer())
	if err != nil {
		t, err := models.TownAreas(models.TownAreaWhere.Name.EQ(currentRow.townAreaName), models.TownAreaWhere.NameRoma.EQ(currentRow.townAreaNameRome), models.TownAreaWhere.MunicipalityID.EQ(rowMunicipality.ID)).One(ctx, d)
		if err != nil {
			return 0, err
		}

		id = t.ID
		rowTownArea = t
	} else {
		id = townArea.ID
		rowTownArea = &townArea
	}

	return id, nil
}

func insertPostalCode() error {
	ctx := context.Background()
	d := db.GetDB()

	pCode := models.PostalCode{Code: currentRow.postalCode, PrefectureID: rowPrefecture.ID, MunicipalityID: rowMunicipality.ID, TownAreaID: rowTownArea.ID}
	err := pCode.Insert(ctx, d, boil.Infer())
	if err != nil {
		_, err := models.PostalCodes(models.PostalCodeWhere.Code.EQ(currentRow.postalCode),
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

func assignCSVDataToCurrentRow(csvRow []string) {
	// ex. []string{"8180025", "福岡県", "筑紫野市", "筑紫", "FUKUOKA KEN", "CHIKUSHINO SHI", "CHIKUSHI"}
	currentRow.postalCode = csvRow[0]
	currentRow.prefectureName = csvRow[1]
	currentRow.municipalityName = csvRow[2]
	currentRow.townAreaName = csvRow[3]
	currentRow.prefectureNameRome = csvRow[4]
	currentRow.municipalityNameRome = csvRow[5]
	currentRow.townAreaNameRome = csvRow[6]
}

func assignCSVDataToPreviousRow(csvRow []string) {
	// ex. []string{"8180025", "福岡県", "筑紫野市", "筑紫", "FUKUOKA KEN", "CHIKUSHINO SHI", "CHIKUSHI"}
	previousRow.postalCode = csvRow[0]
	previousRow.prefectureName = csvRow[1]
	previousRow.municipalityName = csvRow[2]
	previousRow.townAreaName = csvRow[3]
	previousRow.prefectureNameRome = csvRow[4]
	previousRow.municipalityNameRome = csvRow[5]
	previousRow.townAreaNameRome = csvRow[6]
}
