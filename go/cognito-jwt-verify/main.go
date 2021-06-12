package main

import (
	"cognito-jwt-verify/lib"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	resp, err := http.Get(fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", lib.Region, lib.UserPoolId))
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(err)
	}

	body := resp.Body
	defer body.Close()

	var jwk map[string][]map[string]string
	if err := json.NewDecoder(body).Decode(&jwk); err != nil {
		fmt.Println("デコードに失敗したらしい")
		return
	}

	keys := jwk["keys"]
	fmt.Println(keys)

	src := lib.SampleJwt
	token, err := jwt.Parse(src, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}

		var targetKey map[string]string
		for i, key := range keys {
			if key["kid"] == kid {
				targetKey = keys[i]
			}
		}

		if _, ok := targetKey["kid"]; !ok {
			return nil, fmt.Errorf("key with specified kid is not present in jwks")
		}

		fmt.Printf("cognito が公開している鍵で使うべき方: %v\n", targetKey)
		return "", nil
	})

	fmt.Printf("最終的な token: %v\n", token)
}
