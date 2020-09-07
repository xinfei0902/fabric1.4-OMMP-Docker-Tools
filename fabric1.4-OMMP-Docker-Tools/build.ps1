$env:CGO_ENABLE=0
$env:GOOS="linux"
go build -o deploy-tools