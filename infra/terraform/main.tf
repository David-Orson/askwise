provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

resource "google_compute_address" "static_ip" {
  name = "askwise-static-ip"
}

resource "google_compute_instance" "askwise_vm" {
  name         = "askwise-k3s-vm"
  machine_type = "e2-small"
  zone         = var.zone

  tags = ["k3s"]

  boot_disk {
    initialize_params {
      image = "ubuntu-2204-jammy-v20240612"
    }
  }

  network_interface {
    network       = "default"
    access_config {
      nat_ip = google_compute_address.static_ip.address
    }
  }

  metadata_startup_script = file("startup.sh")

  service_account {
    email  = var.service_account_email
    scopes = ["cloud-platform"]
  }
}

resource "google_compute_firewall" "default" {
  name    = "allow-k3s-traffic"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["22", "80", "443", "8080", "3000", "9090", "9091"]
  }

  target_tags = ["k3s"]
}

output "vm_ip" {
  value = google_compute_address.static_ip.address
}

metadata = {
  ssh-keys = "ubuntu:${file(var.ssh_public_key_path)}"
}


