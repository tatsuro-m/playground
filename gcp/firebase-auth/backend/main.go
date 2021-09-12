package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
)

var app *firebase.App

func main() {
	fmt.Println("Hello Firebase!!")
	app, _ = firebaseInit()

	http.HandleFunc("/ping", pingHandlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pingHandlerFunc(w http.ResponseWriter, r *http.Request) {
	jwt := r.Header["Authorization"]
	if len(jwt) != 1 {
		fmt.Fprintf(w, "Authorization ヘッダーがセットされていません。")
	}

	token := jwt[0]
	token = strings.Replace(token, "Bearer ", "", 1)
	fmt.Println("送られてきたそのままの jwt !!")
	fmt.Println(token)

	decodedToken, err := verifyIDToken(context.Background(), app, token)
	if err != nil {
		fmt.Fprintf(w, "token が有効ではありません")
		return
	}

	fmt.Println("デコードした後のトークン↓")
	fmt.Println(decodedToken)
	fmt.Println("ユーザの email")
	fmt.Println(decodedToken.Claims["email"])
	fmt.Println("ユーザの user_id")
	fmt.Println(decodedToken.Claims["user_id"])

	fmt.Fprintf(w, "トークンは有効でした！")
}

func firebaseInit() (*firebase.App, error) {
	opt := option.WithCredentialsFile("./sa_key_file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return app, nil
}

func verifyIDToken(ctx context.Context, app *firebase.App, idToken string) (*auth.Token, error) {
	// [START verify_id_token_golang]
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}

	log.Printf("Verified ID token: %v\n", token)
	// [END verify_id_token_golang]

	return token, nil
}
