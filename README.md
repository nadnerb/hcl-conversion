# HCL Configuration

This project was created to be able to read HCL variable files to be used in other projects.

For example, I would like to pass configuration from a Terraform tfvars file and pass them to ansible.

### Using

Example:

`./hcl-configuration filename.tfvars`

### Building

``go build -ldflags "-X main.GitCommit `git rev-parse HEAD`"``

### Caveats

I don't actually know Go..
