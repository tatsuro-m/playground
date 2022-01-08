package main

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"io"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	srv, err := drive.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	r, err := srv.Files.List().PageSize(1000).
		Fields("files(id, name, mimeType, parents)").
		Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	for _, f := range r.Files {
		if f.MimeType == "application/vnd.google-apps.folder" {
			continue
		}

		fmt.Printf("%s の parent ids は、 %v\n", f.Name, f.Parents)
		if err := download(ctx, srv, f.Name, f.Id); err != nil {
			log.Fatalf("Unable to download: %v", err)
		}
	}

}

func download(ctx context.Context, srv *drive.Service, name, id string) error {
	create, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer create.Close()

	resp, err := srv.Files.Get(id).Context(ctx).Download()
	if err != nil {
		return fmt.Errorf("get drive file: %w", err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(create, resp.Body); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
