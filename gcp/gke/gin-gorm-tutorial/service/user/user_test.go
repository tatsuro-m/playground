package user_test

import (
	"gin-gorm-tutorial/db"
	test_helper "gin-gorm-tutorial/test-helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForeignKeyBehavior(t *testing.T) {
	t.Run("post から参照されている user は削除できないこと", func(t *testing.T) {
		test_helper.SetupTest(t)
		defer test_helper.FinalizeTest(t)

		u := test_helper.InsertUser(t, 1)[0]
		test_helper.InsertPost(t, 5, u)

		d := db.GetDB()
		if err := d.Delete(&u).Error; err != nil {
			expected := "ERROR: update or delete on table \"users\" violates foreign key constraint \"posts_user_id_fkey\" on table \"posts\""
			assert.Contains(t, err.Error(), expected)
		}
	})
}
