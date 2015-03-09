default: generate

generate-test:
	go generate ./internal/fixtures

generate:
	go generate ./aws
	go generate ./service

clean:
	find ./service -maxdepth 2 -mindepth 2 \( -type f -name "api.go" -o -name "service.go" \) -delete
	find ./service -maxdepth 1 -mindepth 1 -type d -delete
