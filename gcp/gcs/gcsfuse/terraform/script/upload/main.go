package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
)

var projectID = "playground-318023"

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name for the new bucket.
	bucketName := getBucketName(1)

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	for i := 0; i < 20; i++ {
		f, err := os.Create(fmt.Sprintf("sample%d.txt", i))
		if err != nil {
			log.Fatal(err)
		}

		objectPath := fmt.Sprintf("sample%d.txt", i)
		obj := bucket.Object(objectPath)
		r, err := obj.NewReader(ctx)
		if err != nil {
			log.Fatal(err)
		}

		tee := io.TeeReader(r, f)
		s := bufio.NewScanner(tee)
		for s.Scan() {
		}
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func getBucketName(i int) string {
	n := strconv.Itoa(i)
	return "stg-gcsfuse-test" + n
}
