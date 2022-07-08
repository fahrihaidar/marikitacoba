job "fahrihaidar" {
  datacenters = ["dc1"]

  group "cache" {
    network {
      port "db" {
        to = 5432
      }
    }

    task "redis" {
      driver = "docker"

      config {
        image = "fahrihaidar/marikitacoba"

        ports = ["db"]
      }

      resources {
        cpu    = 500
        memory = 256
      }
    }
  }
}
