terraform {
  required_providers {
    openwebui = {
      source = "hashicorp.com/edu/openwebui"
    }
  }
}

variable "api_key" {
    description = "API key for OpenWebUI - comes from env called TF_VAR_api_key"
    type        = string
    sensitive   = true
}

provider "openwebui" {
    host    = "http://localhost:3000"
    api_key = var.api_key
}

data "openwebui_users" "default" {}

output "email" {
    value = data.openwebui_users.default
}
