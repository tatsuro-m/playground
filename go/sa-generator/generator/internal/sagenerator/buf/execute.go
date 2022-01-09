package buf

import (
	"bytes"
	"fmt"
	"generator/internal/sagenerator/models"
	"text/template"
)

func TmplExec(sa *models.Sa) (*bytes.Buffer, error) {
	t, err := template.ParseFiles("./tf-tmpl/sa.tf.tmpl")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	b := &bytes.Buffer{}
	err = t.Execute(b, sa)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return b, nil
}
