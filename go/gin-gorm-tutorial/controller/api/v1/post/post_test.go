package post_test

import (
	"bytes"
	"encoding/json"
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/entity"
	"gin-gorm-tutorial/server"
	test_helper "gin-gorm-tutorial/test-helper"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController_Index(t *testing.T) {
	test_helper.SetupTest(t)
	defer test_helper.FinalizeTest(t)

	insertPost := func() {
		d := db.GetDB()
		u := entity.User{FirstName: "first", LastName: "last"}
		d.Create(&u)

		for i := 0; i < 5; i++ {
			p := entity.Post{Title: "title" + strconv.Itoa(i), Content: "content" + strconv.Itoa(i), UserID: u.ID}
			d.Create(&p)
		}
	}
	insertPost()

	router := server.Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/posts", nil)
	router.ServeHTTP(w, req)

	var reqBody []map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &reqBody)
	for i, b := range reqBody {
		assert.Equal(t, "title"+strconv.Itoa(i), b["title"])
		assert.Equal(t, "content"+strconv.Itoa(i), b["content"])
	}
}

func TestController_Create(t *testing.T) {
	test_helper.SetupTest(t)
	defer test_helper.FinalizeTest(t)

	router := server.Router()
	w := httptest.NewRecorder()

	d := db.GetDB()
	insertUser := func() entity.User {
		u := entity.User{FirstName: "first", LastName: "last"}
		d.Create(&u)
		return u
	}

	u := insertUser()
	body := map[string]interface{}{"title": "title1", "content": "content1", "user_id": u.ID, "User": u}
	reqBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/posts", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resBody map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &resBody)
	assert.Equal(t, "title1", resBody["title"])
	assert.Equal(t, "content1", resBody["content"])
}
