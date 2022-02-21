terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.4.0"
    }

    random = {
      source  = "registry.terraform.io/hashicorp/random"
      version = "3.1.0"
    }
  }
  #  required_version = "= 0.15.0"
  required_version = "~> 1.1.2"

#  backend "remote" {
#    organization = "hayashiki"
#
#    workspaces {
#      name = "chat-boiler"
#    }
#  }
}
