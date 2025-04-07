terraform {
  required_providers {
    openwebui = {
      source = "hashicorp.com/edu/openwebui"
    }
  }
}

variable "api_key" {
  description = "API key for the OpenWebUI provider"
  type        = string
  sensitive   = true
}

provider "openwebui" {
    host = "http://localhost:3000"
    api_key = var.api_key
}

locals {
    groups = {
        "noam" = {
            name = "noam"
            description = "This is an example group"
            permissions = {}
        },
        "imported" = {
            name = "imported"
            description = "test"
            permissions = {
                workspace = {
                    models     = true
                    knowledge  = false
                    prompts    = false
                    tools      = false
                }
                chat = {
                    controls     = true
                    delete       = true
                    edit         = true
                    temporary    = true
                    file_upload  = true
                }
                features = {
                    web_search        = true
                    image_generation  = true
                    code_interpreter  = true
                }
            }
        }
    }
}

resource "openwebui_group" "noam" {
    for_each = local.groups
    name = each.key
    description = each.value.description
    permissions = each.value.permissions
}
