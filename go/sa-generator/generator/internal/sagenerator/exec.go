package sagenerator

import (
	"fmt"
	"generator/internal/sagenerator/buf"
	"generator/internal/sagenerator/config"
	"generator/internal/sagenerator/models"
	"os"
)

func Exec() {
	sas := models.GetMatrix(config.GetConfig())
	for _, sa := range sas {
		err := writeTfFile(sa)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func writeTfFile(sa *models.Sa) error {
	b, _ := buf.TmplExec(sa)
	f, err := os.Create(getFilePath(sa))
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(b.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func getFilePath(sa *models.Sa) string {
	fPath := fmt.Sprintf("../infrastructure/terraform/%s/%s", sa.Env, sa.ServiceName+"-sa.tf")
	return fPath
}
