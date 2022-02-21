terraform {
  backend "gcs" {
    bucket = "chat-boiler-t1-tf-state"
  }
}
