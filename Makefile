run:
	@./bin/GoChatter

build:
	@go build -o bin/GoChatter ./cmd/main.go

tailwind:
	@tailwindcss -i ./ui/css/styles.css -o ./ui/public/styles.css --watch

templ:
	@templ generate --watch --proxy=http://localhost:6969

test:
	@go test -v ./...

