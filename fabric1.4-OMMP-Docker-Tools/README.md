1. 工具编译
   windows:
      ./build.ps1
   linux:
      $env:CGO_ENABLE=0
      $env:GOOS="linux"
      go build -o deploy-tools
