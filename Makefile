.PHONY: lint

lint:
	golangci-lint run

b:
	env GOOS=linux GOARCH=amd64 go build -o ./build/incomer .

build_linux_amd64:
	env GOOS=linux GOARCH=amd64 go build -o ./build/incomer .

build_linux_arm:
	env GOOS=linux GOARCH=arm go build -o ./build/incomer .

build_linux_386:
	env GOOS=linux GOARCH=386 go build -o ./build/incomer .

build_darwin_amd64:
	env GOOS=darwin GOARCH=amd64 go build -o ./build/incomer .

build_darwin_arm:
	env GOOS=darwin GOARCH=arm go build -o ./build/incomer .

build_windows_amd64:
	env GOOS=windows GOARCH=amd64 go build -o ./build/incomer .



