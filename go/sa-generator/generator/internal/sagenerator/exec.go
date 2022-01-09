package sagenerator

import (
	"fmt"
	"generator/internal/sagenerator/models"
)

func Exec() {
	fmt.Println("hello generator!")
	models.GetMatrix(getConfig())
}
