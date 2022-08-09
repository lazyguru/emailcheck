clean:
	rm -rf bin
	mkdir bin

build: clean fmt
	go build -o bin/emailcheck cmd/emailcheck/main.go

fmt:
	go fmt github.com/lazyguru/emailcheck/cmd/... github.com/lazyguru/emailcheck/internal/...

