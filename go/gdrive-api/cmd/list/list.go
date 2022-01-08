package main

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"log"
)

func main() {
	ctx := context.Background()

	srv, err := drive.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	r, err := srv.Files.List().PageSize(1000).
		Fields("files(id, name, createdTime, kind, mimeType)").
		Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	for _, f := range r.Files {
		fmt.Printf("file name: %s\n", f.Name)
		fmt.Printf("file id: %s\n", f.Id)
		fmt.Printf("file kind: %s\n", f.Kind)
		fmt.Printf("file mime type: %s\n", f.MimeType)
		fmt.Printf("file created time: %s\n", f.CreatedTime)
		fmt.Println("--------------")
	}
}
