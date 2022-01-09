package buf

import (
	"generator/internal/sagenerator/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTmplExec(t *testing.T) {
	tb := []struct {
		name string
		in   *models.Sa
		out  string
	}{
		{
			name: "正常にテンプレートが埋め込まれたバイトが返ってくること",
			in:   &models.Sa{Env: "stg", ServiceName: "a"},
			out: `
resource "google_service_account" "a" {
  account_id   = "${local.app_prefix}-a"
  display_name = "${local.app_prefix}-a"
  description  = "${local.app_prefix} の a で利用する SA"
}
`,
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			b, err := TmplExec(tt.in)
			assert.Nil(t, err)
			assert.Equal(t, tt.out, b.String())
		})
	}
}
