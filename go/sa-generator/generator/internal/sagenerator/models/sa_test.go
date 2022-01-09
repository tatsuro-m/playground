package models

import (
	"generator/internal/sagenerator/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMatrix(t *testing.T) {
	tb := []struct {
		name string
		in   *config.Config
		out  []*Sa
	}{
		{
			name: "dev と stg で正常に返ってくること",
			in:   &config.Config{TargetEnvs: []string{"dev", "stg"}, TargetServices: []string{"a", "b"}},
			out:  []*Sa{{Env: "dev", ServiceName: "a"}, {Env: "dev", ServiceName: "b"}, {Env: "stg", ServiceName: "a"}, {Env: "stg", ServiceName: "b"}},
		},
		{
			name: "dev と stg と prod で正常に返ってくること",
			in:   &config.Config{TargetEnvs: []string{"dev", "stg", "prod"}, TargetServices: []string{"a", "b"}},
			out:  []*Sa{{Env: "dev", ServiceName: "a"}, {Env: "dev", ServiceName: "b"}, {Env: "stg", ServiceName: "a"}, {Env: "stg", ServiceName: "b"}, {Env: "prod", ServiceName: "a"}, {Env: "prod", ServiceName: "b"}},
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetMatrix(tt.in)
			assert.Equal(t, tt.out, actual)
		})
	}
}
