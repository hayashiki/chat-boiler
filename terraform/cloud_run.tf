data "google_cloud_run_service" "default" {
  name     = var.image_name
  location = var.region

  depends_on = [
    google_project_service.enable_api
  ]
}

locals {
  current_image = data.google_cloud_run_service.default.template != null ? data.google_cloud_run_service.default.template.0.spec.0.containers.0.image : null
  timestamp = formatdate("YYYY-MM-DD-hh:mm:ss", timestamp())
}

resource "google_cloud_run_service" "default" {
  name     = var.image_name
  location = var.region

  depends_on = [
    google_project_service.enable_api
  ]

  template {
    spec {
      service_account_name = google_service_account.run_sa.email

      containers {
        image = local.current_image

        resources {
          limits = {
            cpu    = "1000m"
            memory = "128Mi"
          }
        }

        # if use redis
        env {
          name  = "REDIS_URL"
          value = "redis://${google_redis_instance.cache.host}:6379"
        }
        # Hack to force terraform to re-deploy this service (e.g. update latest image)
        env {
          name = "TERRAFORM_UPDATED_AT"
          value = local.timestamp
        }
      }
    }

    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "1"
        "autoscaling.knative.dev/minScale" : "0",
        # if use vpc connector
        "run.googleapis.com/vpc-access-connector" : google_vpc_access_connector.default.name
        "run.googleapis.com/vpc-access-egress" : "all-traffic" #"private-ranges-only"
      }

      labels = {
        service = var.image_name
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  autogenerate_revision_name = true
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.default.location
  project  = google_cloud_run_service.default.project
  service  = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}

output "run_urls" {
  value = google_cloud_run_service.default.status.0.url
}
