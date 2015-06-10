# HCL Conversion

This project was created to enable reading HCL configuration files to be used by other tools.

For example, I would like to pass configuration from a Terraform tfvars file and use them in ansible.

### Using

Example:

`./hcl-configuration example.tfvars`

With ansible

``ansible-playbook playbook.yml --extra-vars "`$HCLC_PATH/hcl-conversion ~/.terraform/example.tfvars`"``

### Building

``go build -ldflags "-X main.GitCommit `git rev-parse HEAD`"``

### Caveats

I don't know Go.
