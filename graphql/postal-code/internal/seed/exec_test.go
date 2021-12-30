package seed

import (
	"context"
	"github.com/stretchr/testify/assert"
	"pcode/pkg/db"
	"pcode/pkg/models"
	"pcode/pkg/thelper"
	"testing"
)

func TestExec(t *testing.T) {
	t.Run("seed が登録されること", func(t *testing.T) {
		thelper.SetupTest(t)
		defer thelper.FinalizeTest(t)

		Exec()

		ctx := context.Background()
		d := db.GetDB()

		prefecture, _ := models.Prefectures(models.PrefectureWhere.Name.EQ("東京都")).One(ctx, d)
		assert.Equal(t, "TOKYO TO", prefecture.NameRoma)
		municipality, _ := models.Municipalities(models.MunicipalityWhere.Name.EQ("新宿区")).One(ctx, d)
		assert.Equal(t, "SHINJUKU KU", municipality.NameRoma)
		townArea, _ := models.TownAreas(models.TownAreaWhere.Name.EQ("四谷")).One(ctx, d)
		assert.Equal(t, "YOTSUYA", townArea.NameRoma)

		postalCode, _ := models.PostalCodes(models.PostalCodeWhere.Number.EQ("1600004")).One(ctx, d)
		assert.Equal(t, prefecture.ID, postalCode.PrefectureID)
		assert.Equal(t, municipality.ID, postalCode.MunicipalityID)
		assert.Equal(t, townArea.ID, postalCode.TownAreaID)

		count, _ := models.PostalCodes().Count(ctx, d)
		testDataRow := int64(34)
		assert.Equal(t, testDataRow, count)
	})
}
