test:
	GO111MODULE=on go test -v ./...
build-ui:
	cd ui; go run gen.go
	cd ui; sh ./build-macos.sh
