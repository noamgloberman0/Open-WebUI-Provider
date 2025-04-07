# Terraform Provider for Open WebUI

This Terraform provider allows you to manage resources within your Open WebUI instance.

## Requirements

* Terraform 1.0+
* Go 1.20+ (for building the provider)
* An accessible Open WebUI instance.

## Installation

To use this provider, you'll need to either download a pre-built binary (if available) or build it yourself.

### Using Pre-built Binaries (If Available)

Instructions for downloading and installing pre-built binaries will be provided here if releases are available. Typically, this involves downloading the binary for your platform and placing it in your Terraform plugins directory.

### Building the Provider from Source

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd terraform-provider-openwebui
    ```

    Replace `<repository_url>` with the actual URL of this provider's repository.

2.  **Build the provider:**

    ```bash
    go build -o terraform-provider-openwebui
    ```

    This command will build the provider executable.

3.  **Install the provider:**

    Terraform discovers plugins in specific directories. You'll need to place the built executable in one of these directories. The recommended location is within your user's plugin cache:

    ```bash
    mkdir -p ~/.terraform.d/plugins/registry.terraform.io/your-organization/openwebui/your-version/$(go env GOOS)_$(go env GOARCH)
    mv terraform-provider-openwebui ~/.terraform.d/plugins/registry.terraform.io/your-organization/openwebui/your-version/$(go env GOOS)_$(go env GOARCH)/
    ```

    * Replace `your-organization` with your organization name or a placeholder.
    * Replace `your-version` with the version of the provider you are building (e.g., `0.1.0` or `dev`).

    Alternatively, for local development, you can configure Terraform to load plugins from a specific development directory. Refer to Terraform's plugin development documentation for more details.

## Configuration

To configure the provider, you need to specify the Open WebUI API host and an API key.

```terraform
terraform {
  required_providers {
    openwebui = {
      source = "your-organization/openwebui" # Replace with your organization/namespace
      version = "your-version"             # Replace with your provider version
    }
  }
}

provider "openwebui" {
  host    = "http://your-openwebui-instance:port" # Replace with your Open WebUI host and port
  api_key = "your_api_key"                       # Replace with your Open WebUI API key
}
```
