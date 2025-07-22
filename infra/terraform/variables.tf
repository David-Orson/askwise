variable "project_id" {
  default = "askwise-465909"
}

variable "region" {
  default = "europe-west2"
}

variable "zone" {
  default = "europe-west2-a"
}

variable "service_account_email" {}

variable "ssh_public_key_path" {
  default = "~/.ssh/id_rsa.pub"
}
