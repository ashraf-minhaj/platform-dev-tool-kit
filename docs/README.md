# Run locally

```bash
cd platform
go mod init platform

# run
go run .

go run -C platform/ . start ticket EPD-300
```

# make it binary
```bash
go build -o platform

GOOS=windows GOARCH=amd64 go build -o platform.exe
C:\Users\<username>\bin\

GOOS=darwin GOARCH=arm64 go build -o platform
sudo mv platform /opt/homebrew/bin/platform

GOOS=linux GOARCH=amd64 go build -o platform
sudo mv platform /usr/local/bin/platform
```
