package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	src := sampleJwt()
	claims := strings.Split(src, ".")[1]

	claimsDec, err := base64.RawURLEncoding.DecodeString(claims)

	var claimMap map[string]interface{}
	json.Unmarshal(claimsDec, &claimMap)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(claimsDec))
	fmt.Println(claimMap["email"])
}
