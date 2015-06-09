# HCL Configuration

### Using

./hcl-configuration filename

### Building
go build -ldflags "-X main.GitCommit `git rev-parse HEAD`"
