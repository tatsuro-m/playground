package sagenerator

import (
	"fmt"
	"generator/internal/sagenerator/config"
	"generator/internal/sagenerator/models"
)

func Exec() {
	sas := models.GetMatrix(config.GetConfig())
	fmt.Println(sas)
}
