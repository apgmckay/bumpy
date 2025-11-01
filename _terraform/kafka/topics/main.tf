terraform {
  required_providers {
    kafka = {
      source  = "Mongey/kafka"
      version = "~> 0.7"
    }
  }
}

provider "kafka" {
  bootstrap_servers = [format("%s:%s", var.hostname, var.hostport)]
  tls_enabled       = false
}

resource "kafka_topic" "bumpy_send" {
  name               = "bumpy_send"
  replication_factor = 1
  partitions         = 10
}

resource "kafka_topic" "bumpy_recieve" {
  name               = "bumpy_recieve"
  replication_factor = 1
  partitions         = 10
}
