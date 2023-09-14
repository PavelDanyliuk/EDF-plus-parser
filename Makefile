build_wasm:
			GOOS=js GOARCH=wasm go build -o build/main.wasm ./cmd/wasm/main.go
build_darwin:
			GOOS=darwin GOARCH=arm64 go build -o build/main_darwin ./cmd/binary/main.go
build_cmodule:
		    go build -buildmode=c-shared -o build/main.so ./cmd/cmodule/main.go