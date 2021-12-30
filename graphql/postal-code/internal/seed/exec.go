package seed

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
	"pcode/pkg/db"
	"pcode/pkg/models"
)

func Exec() error {
	p, _ := filepath.Abs("../../KEN_ALL_ROME.CSV")
	utf8F, _ := os.OpenFile(p, os.O_RDONLY, 0666)
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

		insertData(record)
	}

	return nil
}

func insertData(csvRow []string) {
	ctx := context.Background()
	d := db.GetDB()

	// ex. []string{"8180025", "福岡県", "筑紫野市", "筑紫", "FUKUOKA KEN", "CHIKUSHINO SHI", "CHIKUSHI"}
	prefecture := models.Prefecture{Name: csvRow[1], NameRoma: csvRow[4]}
	if b, _ := models.Prefectures(models.PrefectureWhere.Name.EQ(prefecture.Name)).Exists(ctx, d); !b {
		prefecture.Insert(ctx, d, boil.Infer())
	}

	p, _ := models.Prefectures(qm.Select(models.PrefectureColumns.ID), models.PrefectureWhere.Name.EQ(prefecture.Name)).One(ctx, d)
	municipality := models.Municipality{Name: csvRow[2], NameRoma: csvRow[5], PrefectureID: p.ID}
	if b, _ := models.Municipalities(models.MunicipalityWhere.Name.EQ(municipality.Name)).Exists(ctx, d); !b {
		municipality.Insert(ctx, d, boil.Infer())
	}

	m, _ := models.Municipalities(qm.Select(models.MunicipalityColumns.ID), models.MunicipalityWhere.Name.EQ(municipality.Name)).One(ctx, d)
	townArea := models.TownArea{Name: csvRow[3], NameRoma: csvRow[6], MunicipalityID: m.ID}
	townArea.Insert(ctx, d, boil.Infer())

	t, _ := models.TownAreas(qm.Select(models.TownAreaColumns.ID), models.TownAreaWhere.Name.EQ(townArea.Name)).One(ctx, d)
	postalCode := models.PostalCode{Number: csvRow[0], PrefectureID: p.ID, MunicipalityID: m.ID, TownAreaID: t.ID}
	postalCode.Insert(ctx, d, boil.Infer())
}
