locals {
  greeting = "hello from terraform-tools"
}

output "greeting" {
  value = local.greeting
}
