resource "google_project_service" "enable_api" {
  for_each = toset([
    "iam.googleapis.com",
    "containerregistry.googleapis.com",
    "run.googleapis.com",
    "compute.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "iamcredentials.googleapis.com",

    # if use redis
    "redis.googleapis.com",

    # if use vpc access
    "vpcaccess.googleapis.com"
  ])
  project = var.project
  service = each.value

  disable_dependent_services = true
}
