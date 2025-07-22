terraform {
  backend "gcs" {
    bucket = "askwise-tf-state"
    prefix = "env/default"
  }
  required_version = ">= 1.7.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 6.0"
    }
  }
}
