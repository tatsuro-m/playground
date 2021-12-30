package graph

import "github.com/vektah/gqlparser/v2/gqlerror"

func NewGqlError(msg string, code string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}
