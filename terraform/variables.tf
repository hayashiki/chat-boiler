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

variable "repo" {
  type    = string
  default = "hayashiki/chat-boiler"
}

variable "image_name" {
  type    = string
  default = "chat-boiler-api"
}

variable "image_tag" {
  type    = string
  default = "latest"
}
