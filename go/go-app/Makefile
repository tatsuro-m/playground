build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o ./server

run:build
	./server
