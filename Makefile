build:
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build  -o dist/Tech_Radar-linux src/main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/Tech_Radar-darwin src/main.go
	GOOS=windows GOARCH=amd64 go build -o dist/Tech_Radar-windows.exe src/main.go

clean_build:
	rm -rf ./dist/*
