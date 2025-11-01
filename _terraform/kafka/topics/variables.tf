variable "kafka_hostname" {
  type        = string
  description = "Kafka host name"
  default     = "localhost"
}

variable "kafka_hostport" {
  type        = string
  description = "Kafka host port"
  default     = "9092"
}

variable "kafka_tls_enabled" {
  type        = bool
  description = "Kafka enable tls for terraform provider"
  default     = false
}
