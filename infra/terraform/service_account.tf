resource "google_service_account" "ci" {
  account_id   = "askwise-ci"
  display_name = "AskWise CI/CD Service Account"
}

# Roles — start with these; tighten later.
locals {
  ci_roles = [
    "roles/compute.admin",            # create VM, IPs, firewall
    "roles/iam.serviceAccountUser",   # actAs when needed
    "roles/artifactregistry.admin",   # push images (or writer)
    "roles/storage.admin",            # manage Terraform state bucket (or objectAdmin)
    "roles/dns.admin",                # only if you do DNS later (optional)
  ]
}

resource "google_project_iam_member" "ci_bindings" {
  for_each = toset(local.ci_roles)
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.ci.email}"
}

output "ci_service_account_email" {
  value = google_service_account.ci.email
}
