package buf

import (
	"bytes"
	"fmt"
	"generator/internal/sagenerator/models"
	"text/template"
)

func TmplExec(sa *models.Sa) (*bytes.Buffer, error) {
	// プログラム起動時のワーキングディレクトリから見ての相対パス
	t, err := template.ParseFiles("./internal/sagenerator/tf-tmpl/sa.tf.tmpl")
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
