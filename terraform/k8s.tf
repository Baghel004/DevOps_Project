resource "kubernetes_deployment" "name"{
    metadata {
        name = "goappdeployment"
        labels = {
            "type" = "backend"
            "app" = "goapp"
        }
    }
    spec {
        replicas = 1
        selector {
            match_labels = {
                "type" = "backend"
                "app" = "goapp"
            }
        }
        template {
            metadata {
                name = "goapppod"
                labels = {
                    "type" = "backend"
                    "app" = "goapp"
                } 
            }
            spec {
                container {
                    name = "goappcontainer"
                    image = var.container_image
                    port {
                        container_port = 80
                    }
                }
            }
        }
    }
}

resource "google_compute_address" "default" {
    name = "ipforservice"
    region = var.region
}
resource "kubernetes_service" "appservice" {
    metadata {
        name = "goapp-lb-service"
    }
    spec {
        type = "LoadBalancer"
        load_balancer_ip = google_compute_address.default.address
        port {
            port = 80
            target_port = 80
        }
        selector = {
             "type" = "backend"
             "app" = "goapp"
        }
    }
}