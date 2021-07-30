package post_test

import (
	"bytes"
	"encoding/json"
	test_helper "gin-gorm-tutorial/test-helper"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController_Index(t *testing.T) {
	test_helper.SetupTest(t)
	defer test_helper.FinalizeTest(t)

	u := test_helper.InsertUser(t, 1)[0]
	test_helper.InsertPost(t, 5, u)

	w, _ := test_helper.SendHttpRequest(t, http.MethodGet, "/api/v1/posts", nil)

	var resBody []map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &resBody)
	assert.Len(t, resBody, 5)
	for i, b := range resBody {
		assert.Equal(t, "title"+strconv.Itoa(i), b["title"])
		assert.Equal(t, "content"+strconv.Itoa(i), b["content"])
	}
}

func TestController_Create(t *testing.T) {
	test_helper.SetupTest(t)
	defer test_helper.FinalizeTest(t)

	u := test_helper.InsertUser(t, 1)[0]

	body := map[string]interface{}{"title": "title1", "content": "content1", "user_id": u.ID, "User": u}
	reqBody, _ := json.Marshal(body)

	w, _ := test_helper.SendHttpRequest(t, http.MethodPost, "/api/v1/posts", bytes.NewBuffer(reqBody))
	assert.Equal(t, http.StatusCreated, w.Code)

	var resBody map[string]interface{}
	_ = json.Unmarshal([]byte(w.Body.String()), &resBody)
	assert.Equal(t, "title1", resBody["title"])
	assert.Equal(t, "content1", resBody["content"])
}
