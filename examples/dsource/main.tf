/**
* Summary: This template showcases the properties available when creating an app data dsource.
*/

terraform {
  required_providers {
    delphix = {
      version = "2.0.4-beta"
      source  = "delphix.com/local/delphix"
    }
  }
}

provider "delphix" {
  tls_insecure_skip = true
  key               = "1.DRzz1XQVNDX36biMJ7mzVv5ANPu2QipYFF3MA6nDP0y9swwIpoxZaOkafZpkSYlS"
  host              = "dct10.dlpxdc.co"
}

resource "delphix_appdata_dsource" "test_app_data_dsource" {
  source_id                  = "2-APPDATA_STAGED_SOURCE_CONFIG-2"
  group_id                   = "2-GROUP-1"
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = "AppDataStaged"
  name                       = "test_app_data_dsource_ankit_patil"
  staging_mount_base         = "200"
  environment_user           = "HOST_USER-2"
  staging_environment        = "2-UNIX_HOST_ENVIRONMENT-2"
  parameters = jsonencode({
    postgresPort = 5432
  })
  sync_parameters = jsonencode({
    resync = true
  })
}

