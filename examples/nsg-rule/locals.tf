locals {
  network_security_group_name = module.resource_names["network_security_group"].standard
  resource_group_name         = module.resource_names["resource_group"].standard

  security_rules = {
    for k, v in var.security_rules : k => {
      name                                       = v.name
      resource_group_name                        = local.resource_group_name
      network_security_group_name                = local.network_security_group_name
      description                                = v.description
      protocol                                   = v.protocol
      source_port_range                          = v.source_port_range
      source_port_ranges                         = v.source_port_ranges
      destination_port_range                     = v.destination_port_range
      destination_port_ranges                    = v.destination_port_ranges
      source_address_prefix                      = v.source_address_prefix
      source_address_prefixes                    = v.source_address_prefixes
      destination_address_prefix                 = v.destination_address_prefix
      destination_address_prefixes               = v.destination_address_prefixes
      destination_application_security_group_ids = v.destination_application_security_group_ids
      access                                     = v.access
      priority                                   = v.priority
      direction                                  = v.direction
    }
  }
}
