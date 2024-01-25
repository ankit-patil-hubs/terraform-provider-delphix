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
  key               = "1.pvR2JlMe9MEWHHV38yhNPbGMOHV9W1R2iiGYguXXSgskSIlAlyeNxiDmESFGNBLC"
  host              = "dct101.dlpxdc.co"
}



# resource "delphix_appdata_dsource" "test_app_data_dsource" {
#   source_value                  = "1-APPDATA_STAGED_SOURCE_CONFIG-6"
#   group_id                   = "1-GROUP-1"
#   log_sync_enabled           = false
#   make_current_account_owner = true
#   link_type                  = "AppDataStaged"
#   name                       = "appdata_dsource"
#   staging_mount_base         = ""
#   environment_user           = "HOST_USER-2"
#   staging_environment        = "1-UNIX_HOST_ENVIRONMENT-2"
#   parameters = jsonencode({
#     externalBackup : [],
#     delphixInitiatedBackupFlag : true,
#     delphixInitiatedBackup : [
#       {
#         userName : "XXXX",
#         postgresSourcePort : XXXX,
#         userPass : "XXXX",
#         sourceHostAddress : "HOSTNAME"
#       }
#     ],
#     singleDatabaseIngestionFlag : false,
#     singleDatabaseIngestion : [],
#     stagingPushFlag : false,
#     postgresPort : XXXX,
#     configSettingsStg : [],
#     mountLocation : "/tmp/delphix_mnt"
#   })
#   sync_parameters = jsonencode({
#     resync = true
#   })
# }

#  "id": "11-APPDATA_STAGED_SOURCE_CONFIG-2",
#   "database_type": "postgres-vsdk",
#   "name": "ankit",
#   "is_replica": false,
#   "environment_id": "11-UNIX_HOST_ENVIRONMENT-1",
#   "ip_address": "10.110.254.226",
#   "fqdn": "rhel-86-boph-qar-113675-27a4593a.dlpxdc.co",
#   "plugin_version": "4.1.1",
#   "is_dsource": false,
#   "repository": "11-APPDATA_REPOSITORY-2"


# {
#   "source_id": "12-APPDATA_STAGED_SOURCE_CONFIG-7",
#   "group_id": "12-GROUP-1",
#   "log_sync_enabled": false,
#   "make_current_account_owner": true,
#   "link_type": "AppDataStaged",
#   "staging_mount_base": "",
#   "environment_user": "HOST_USER-8",
# "staging_environment": "12-UNIX_HOST_ENVIRONMENT-7",
#   "parameters": {
# "mount_location": "/abc"
# },
#   "sync_parameters": {
#     "resync": true
#   }
# }




resource "delphix_appdata_dsource" "test_app_data_dsource_second" {
  source_value               = "7-APPDATA_STAGED_SOURCE_CONFIG-9"
  group_id                   = "7-GROUP-2"
  log_sync_enabled           = false
  make_current_account_owner = true
  link_type                  = "AppDataStaged"
  name                       = "appdata_dsource_second_new"
  staging_mount_base         = ""
  environment_user           = "HOST_USER-4"
  staging_environment        = "7-UNIX_HOST_ENVIRONMENT-4"
  parameters = jsonencode({
    delphixInitiatedBackupFlag : true,
    delphixInitiatedBackup : [
      {
        userName : "delphix",
        postgresSourcePort : 5432,
        userPass : "delphix",
        sourceHostAddress : "postgressrc.dlpxdc.co"
      }
    ],
    postgresPort : 5433,
    mountLocation : "/datadrive1/provision/ds-assetmanagement-neur-tntnpk-dev-rocsexecution-1"
  })
  sync_parameters = jsonencode({
    resync = true
  })
}


# resource "delphix_vdb" "example" {
#   name                   = "vdb_to_be_created"
#   source_data_id         = delphix_appdata_dsource.test_app_data_dsource_second.id
#   vdb_restart            = true
#   auto_select_repository = true
#   appdata_source_params = jsonencode({
#     mountLocation     = "/mnt/GAT"
#     postgresPort      = 5434
#     configSettingsStg = [{ propertyName : "timezone", value : "GMT", commentProperty : false }]
#   })
#   make_current_account_owner = true
# }


# "appdata_source_params": {
# "configSettingsStg": [],
# "mountLocation": "/snowvdbmount05",
# "postgresPort": 5434
# }
# Below are the 3 ways to link dsource with params , use any one of them
#  externalBackup: [
#             {
#                 keepStagingInSync: false,
#                 backupPath: "/var/tmp/backup",
#                 walLogPath: "/var/tmp/backup"
#             }
# ]

# singleDatabaseIngestion: [
#             {
#                 databaseUserName: "postgres",
#                 sourcePort: 5432,
#                 dumpJobs: 2,
#                 restoreJobs: 2,
#                 databaseName: "abcd",
#                 databaseUserPassword: "xxxx",
#                 dumpDir: "abcd",
#                 sourceHost: "abcd",
#                 postgresqlFile: "abcd"
#             }
#         ]

# delphixInitiatedBackup : [
#   {
#     userName : "XXXX",
#     postgresSourcePort : XXXX,
#     userPass : "XXXX",
#     sourceHostAddress : "HOSTNAME"
#   }
# ]
