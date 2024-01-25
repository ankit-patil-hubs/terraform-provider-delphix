/**
* Summary: This template showcases the properties available when creating an app data dsource.
*/

# Configure the connection to Data Control Tower
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

resource "delphix_source_config" "test_source_config" {
  engine_ip = "http://enginetf.dlpxdc.co/resources/json/delphix"
  user_name = "DBOMSRB331B3"
  password  = "CREATED_VIA_API"
}

