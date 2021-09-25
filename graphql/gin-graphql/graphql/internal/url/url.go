package url

import (
	"fmt"
	"os"
)

func GetAPIPath(path string) string {
	port := os.Getenv("GIN_API_PORT")
	if port == "" {
		port = "8080"
	}

	return "http://" + os.Getenv("GIN_API_HOST") + ":" + port + addAPIV1Path(path)
}

func addAPIV1Path(path string) string {
	return fmt.Sprintf("/api/v1%s", path)
}
