terraform {
  required_version = ">= 0.12"
}

provider "appgate" {
  username = "admin"
  password = "admin"
  url      = "https://ec2-54-80-224-21.compute-1.amazonaws.com:444/admin"
  provider = "local"
  insecure = true
}


# data "appgate_appliance" "gateway_appliance" {
#   # appliance_id = "90eb0df9-14c0-45cc-bbd0-5dbd562b7d1b"
#   appliance_name = "gateway-0c8cbe4b-567d-4269-9143-7cccdd0f90ab-site1"
# }

data "appgate_site" "default_site" {
  site_name = "Default site"
}

resource "appgate_appliance" "new_gateway" {
  name     = "gateway-foo"
  hostname = "envy-10-97-168-1337.devops"

  client_interface {
    hostname       = "envy-10-97-168-1337.devops"
    proxy_protocol = true
    https_port     = 444
    dtls_port      = 445
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
    override_spa_mode = "TCP"
  }

  peer_interface {
    hostname   = "envy-10-97-168-1337.devops"
    https_port = "1337"

    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }


  admin_interface {
    hostname = "envy-10-97-168-1337.devops"
    https_ciphers = [
      "ECDHE-RSA-AES256-GCM-SHA384",
      "ECDHE-RSA-AES128-GCM-SHA256"
    ]
  }

  tags = [
    "terraform",
    "api-created"
  ]
  notes = "hello world"
  site  = data.appgate_site.default_site.id

  connection {
    type     = "ssh"
    user     = "cz"
    password = "cz"
    host     = "10.97.168.30"
  }

  networking {

    hosts {
      hostname = "bla"
      address  = "0.0.0.0"
    }
    hosts {
      hostname = "foo"
      address  = "127.0.0.1"
    }

    nics {
      enabled = true
      name    = "eth0"
      ipv4 {
        dhcp {
          enabled = true
          dns     = true
          routers = true
          ntp     = true
        }
      }
    }
    dns_servers = [
      "8.8.8.8",
      "1.1.1.1",
    ]
    dns_domains = [
      "aa.com"
    ]
    routes {
      address = "0.0.0.0"
      netmask = 24
      gateway = "1.2.3.4"
      nic = "eth0"
    }
  }


  # ntp {
  #   servers = [
  #     {
  #       hostname = "ntp.microsoft.com"
  #       key_type = "MD5"
  #       key      = "bla"
  #     }
  #   ]
  # }

  ssh_server {
    enabled                 = true
    port                    = 2222
    password_authentication = true
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
    # allow_sources {
    #   address = "0.0.0.0"
    #   netmask = 0
    #   nic     = "eth1"
    # }
  }
  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=controller-a
  controller {
    enabled = true
  }

  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=logserver-a
  log_server {
    enabled = false
    # retention_days = 2
  }
  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=gateway-a
  gateway {
    enabled = true
    vpn {
      weight = 60
      allow_destinations {
        address = "127.0.0.1"
        netmask = 0
        nic     = "eth0"
      }
    }
  }
  # Save the seed file locally in base 64 format.
  provisioner "local-exec" {
    command = "echo ${appgate_appliance.new_gateway.seed_file} > seed.b64"
  }
  # provisioner "remote-exec" {
  #   inline = [
  #     "echo ${appgate_appliance.new_gateway.seed_file}  > raw.b64",
  #     # "cat raw.b64 | base64 -d  | jq .  >> seed.json",
  #   ]
  # }

}

# output "seed_file" {
#   value = "${appgate_appliance.new_gateway.seed_file}"
# }
