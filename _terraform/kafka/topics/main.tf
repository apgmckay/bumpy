terraform {
  required_providers {
    kafka = {
      source  = "Mongey/kafka"
      version = "~> 0.7"
    }
  }
}

provider "kafka" {
  bootstrap_servers = [format("%s:%s", var.kafka_hostname, var.kafka_hostport)]
  tls_enabled       = var.kafka_tls_enabled
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
