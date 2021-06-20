package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/entity"
	test_helper "gin-gorm-tutorial/test-helper"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

type req struct {
	body map[string]interface{}
}
type expected struct {
	code int
	body map[string]interface{}
}

func TestController_Index(t *testing.T) {
	tests := []struct {
		name     string
		expected expected
	}{
		{
			name: "全ての user が取得できること",
			expected: expected{
				code: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test_helper.SetupTest(t)
			defer test_helper.FinalizeTest(t)

			res := httptest.NewRecorder()

			insertUser(t, 3)
			c, _ := gin.CreateTestContext(res)
			c.Request, _ = http.NewRequest(
				http.MethodGet,
				"/api/v1/users",
				nil,
			)

			// Act ---
			var ctrl Controller
			ctrl.Index(c)

			// Assert ---
			assert.Equal(t, tt.expected.code, res.Code)

			var resBody []map[string]interface{}
			_ = json.Unmarshal(res.Body.Bytes(), &resBody)

			assert.Len(t, resBody, 3)
			for i, user := range resBody {
				assert.Equal(t, "first_name"+strconv.Itoa(i), user["first_name"])
				assert.Equal(t, "last_name"+strconv.Itoa(i), user["last_name"])
				assert.Contains(t, user, "id")
				assert.Contains(t, user, "created_at")
				assert.Contains(t, user, "updated_at")
			}
		})
	}
}

func insertUser(t *testing.T, times int) {
	t.Helper()

	d := db.GetDB()
	for i := 0; i < times; i++ {
		u := entity.User{
			FirstName: "first_name" + strconv.Itoa(i),
			LastName:  "last_name" + strconv.Itoa(i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := d.Create(&u).Error; err != nil {
			t.Log(err)
		}
	}
}

func TestController_Create(t *testing.T) {
	tests := []struct {
		name     string
		req      req
		expected expected
	}{
		{
			name: "正常に作成されること",
			req: req{map[string]interface{}{
				"first_name": "test_first1",
				"last_name":  "test_last1",
			}},
			expected: expected{
				code: http.StatusCreated,
				body: map[string]interface{}{
					"first_name": "test_first1",
					"last_name":  "test_last1",
				},
			},
		},
		{
			name: "余計なフィールドがあっても作成されること",
			req: req{map[string]interface{}{
				"first_name":       "test_first1",
				"last_name":        "test_last1",
				"not_exists_field": "hogehoge",
			}},
			expected: expected{
				code: http.StatusCreated,
				body: map[string]interface{}{
					"first_name": "test_first1",
					"last_name":  "test_last1",
				},
			},
		},
		{
			name: "first_name が欠けていたらエラーになること",
			req: req{map[string]interface{}{
				"last_name": "test_last1",
			}},
			expected: expected{
				code: http.StatusBadRequest,
				body: map[string]interface{}{},
			},
		},
		{
			name: "last_name が欠けていたらエラーになること",
			req: req{map[string]interface{}{
				"first_name": "test_first1",
			}},
			expected: expected{
				code: http.StatusBadRequest,
				body: map[string]interface{}{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test_helper.SetupTest(t)
			defer test_helper.FinalizeTest(t)

			reqBody, _ := json.Marshal(tt.req.body)
			res := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(res)
			c.Request, _ = http.NewRequest(
				http.MethodPost,
				"/api/v1/users",
				bytes.NewBuffer(reqBody),
			)

			// Act ---
			var ctrl Controller
			ctrl.Create(c)

			// Assert ---
			assert.Equal(t, tt.expected.code, res.Code)

			var resBody map[string]interface{}
			_ = json.Unmarshal(res.Body.Bytes(), &resBody)

			if resBody == nil {
				assert.Nil(t, resBody)
			} else {
				for k, v := range tt.expected.body {
					assert.Equal(t, v, resBody[k])
				}
				assert.Contains(t, resBody, "id")
				assert.Contains(t, resBody, "created_at")
				assert.Contains(t, resBody, "updated_at")
			}
		})
	}
}

func TestController_Update(t *testing.T) {
	tests := []struct {
		name     string
		req      req
		expected expected
	}{
		{
			name: "正常にアップデートされること",
			req: req{map[string]interface{}{
				"first_name": "changed",
				"last_name":  "changed",
			}},
			expected: expected{
				code: http.StatusOK,
				body: map[string]interface{}{
					"first_name": "changed",
					"last_name":  "changed",
				},
			},
		},
		{
			name: "余計なリクエストボディが付いていても無視されること",
			req: req{map[string]interface{}{
				"first_name": "changed",
				"last_name":  "changed",
				"hoge":       "hoge",
			}},
			expected: expected{
				code: http.StatusOK,
				body: map[string]interface{}{
					"first_name": "changed",
					"last_name":  "changed",
				},
			},
		},
		{
			name: "リクエストボディが空の場合には更新されないこと",
			req:  req{map[string]interface{}{}},
			expected: expected{
				code: http.StatusOK,
				body: map[string]interface{}{
					"first_name": "original_first",
					"last_name":  "original_last",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test_helper.SetupTest(t)
			defer test_helper.FinalizeTest(t)

			insertUser := func() entity.User {
				d := db.GetDB()
				u := entity.User{FirstName: "original_first", LastName: "original_last"}
				d.Create(&u)
				return u
			}
			u := insertUser()
			id := fmt.Sprintf("%d", u.ID)
			reqBody, _ := json.Marshal(tt.req.body)
			res := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(res)
			c.Params = gin.Params{gin.Param{Key: "id", Value: id}}
			c.Request, _ = http.NewRequest(
				http.MethodPut,
				fmt.Sprintf("/api/v1/users/%s", id),
				bytes.NewBuffer(reqBody),
			)

			var ctrl Controller
			ctrl.Update(c)

			assert.Equal(t, tt.expected.code, res.Code)

			var resBody map[string]interface{}
			_ = json.Unmarshal(res.Body.Bytes(), &resBody)

			assert.Equal(t, tt.expected.body["first_name"], resBody["first_name"])
			assert.Equal(t, tt.expected.body["last_name"], resBody["last_name"])
			assert.Equal(t, id, fmt.Sprintf("%v", resBody["id"]))

			format := func(t time.Time) string {
				l := "2006-01-02"
				return t.Format(l)
			}
			assert.Contains(t, resBody["created_at"], format(u.CreatedAt))
			assert.Contains(t, resBody["updated_at"], format(u.UpdatedAt))
		})
	}
}
