
gofmt:
	go fmt ./...

gotest:
	go test -v  ./...

golint:
	golint $$(go list -f="{{.Dir}}" ./...)


go: gofmt golint gotest