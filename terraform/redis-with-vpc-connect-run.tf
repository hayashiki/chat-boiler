

resource "google_redis_instance" "cache" {
  depends_on = [
    google_project_service.enable_api,
    google_project_iam_member.github_actions_default
  ]
  project            = var.project
  name               = "cache"
  region             = var.region
  memory_size_gb     = 1
  authorized_network = google_compute_network.serverless-vpc.name
}

resource "google_vpc_access_connector" "default" {
  depends_on = [
    google_project_service.enable_api
  ]
  provider      = google-beta # allows us to configure machine_type for now
  name          = "vpc-connector"
  region        = var.region
  project       = var.project
#  subnet側で対応したので削除
#  network       = google_compute_network.vpc.name
  machine_type  = "f1-micro"
  subnet {
    name = google_compute_subnetwork.default.name
  }
#  subnet側で対応したので削除
#  ip_cidr_range = "10.8.0.0/28"
}
