# GophKeeper

## Start Server
```bash
go run .\cmd\server\
```

## Start Client
```bash
go run .\cmd\client\
go run -ldflags "-X main.BuildVersion=v1.0.1 -X main.BuildDate=01.01.2024" .\cmd\client\  # To see build version
```