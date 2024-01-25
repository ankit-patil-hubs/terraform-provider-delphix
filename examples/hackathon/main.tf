terraform {
  required_version = ">= 1.2.0"
  required_providers {
    delphix = {
      version = "3.1.0"
      source  = "delphix-integrations/delphix"
    }
    time = {
      source  = "hashicorp/time"
      version = "0.9.2"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "1.pvR2JlMe9MEWHHV38yhNPbGMOHV9W1R2iiGYguXXSgskSIlAlyeNxiDmESFGNBLC"
  host              = "dct101.dlpxdc.co"
}

provider "time" {
  # Configuration options
}

resource "time_sleep" "wait_30_seconds" {
  destroy_duration = "30s"
}

resource "null_resource" "register_engine" {
  provisioner "local-exec" {
    command = "/bin/bash reg_engg.sh"
  }
  depends_on = [time_sleep.wait_30_seconds]
}

# output "engine_id" {
#   value = jsondecode(file("output.json")).id
# }

locals {
  host_name = "psgrs-rhel.dlpxdc.co"
}

resource "null_resource" "cloudability-setup" {
  provisioner "local-exec" {
    command = <<EOT
        curl -s -X POST https://api.cloudability.com/v3/vendors/aws/accounts \
             -H 'Content-Type: application/json' \
             -u "$${CldAbltyAPIToken:?Missing Cloudability API Token Env Variable}:" \
             -d '{"vendorAccountId": "${data.aws_caller_identity.current.account_id}", "type": "aws_role" }'
EOT
  }
}

# resource "time_sleep" "wait_30_seconds" {
#   create_duration = "30s"
# }

# resource "delphix_environment" "unixtgt" {
#   depends_on   = [time_sleep.wait_30_seconds]
#   engine_id    = jsondecode(file("output.json")).id
#   os_name      = "UNIX"
#   username     = "postgres"
#   password     = "postgres"
#   hostname     = local.host_name
#   toolkit_path = "/home/delphix_os/toolkit"
#   name         = "unixtgt"
#   description  = "This is a unix target."
# }
