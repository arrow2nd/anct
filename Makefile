.PHONY: generate
generate:
	go get github.com/Yamashou/gqlgenc
	go run github.com/Yamashou/gqlgenc
	go mod tidy -v
