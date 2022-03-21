# redis対応で追加
resource "google_compute_network" "vpc" {
  project = var.project
  name    = "vpc"
  depends_on = [
    google_project_service.enable_api
  ]
  mtu = 1460
  auto_create_subnetworks = false
}

# StaticIP対応で追加
resource "google_compute_subnetwork" "default" {
  ip_cidr_range = "10.8.0.0/28"
  name          = "subnetwork-serverless"
  region        = var.region
  network       = google_compute_network.vpc.id
}

resource "google_compute_router" "default" {
  name = "router-serverless"
  network = google_compute_network.vpc.name
  region = var.region
}

resource "google_compute_address" "default" {
  count        = 1
  name         = "ip-serverless"
  address_type = "EXTERNAL"
  region = var.region
}

resource "google_compute_router_nat" "default" {
  name                               = "nat-serverless"
  region = var.region
  nat_ip_allocate_option             = "MANUAL_ONLY"
  router                             = google_compute_router.default.name
  nat_ips = google_compute_address.default.*.self_link
  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  subnetwork {
    name                    = google_compute_subnetwork.default.id
    source_ip_ranges_to_nat = ["ALL_IP_RANGES"]
  }
}
