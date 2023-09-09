build_wasm:
			GOOS=js GOARCH=wasm go build -o build/main.wasm ./cmd/wasm/main.go
build_darwin:
			GOOS=darwin GOARCH=arm64 go build -o build/main_darwin ./cmd/binary/main.go