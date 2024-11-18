build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o install-comfyui-custom-node main.go
