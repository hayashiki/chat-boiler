variable "env" {
  type = string
}
variable "project" {
  type = string
}

variable "location" {
  type    = string
  default = "asia-northeast1"
}

# cloudrun service name
variable "name" {
  type    = string
  default = "chat-boiler-api"
}

variable "gar_repository" {
  type    = string
  default = "hayashiki"
}

variable "image_name" {
  type    = string
  default = "go-boiler-api"
}

variable "image_tag" {
  type    = string
  default = "latest"
}
