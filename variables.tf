// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

variable "security_rules" {
  description = "A map of security rules to create"
  type = map(object({
    name                                       = string
    resource_group_name                        = string
    network_security_group_name                = string
    description                                = optional(string)
    protocol                                   = string
    source_port_range                          = optional(number)
    source_port_ranges                         = optional(list(number))
    destination_port_range                     = optional(number)
    destination_port_ranges                    = optional(list(number))
    source_address_prefix                      = optional(string)
    source_address_prefixes                    = optional(list(string))
    source_application_security_group_ids      = optional(list(string))
    destination_address_prefix                 = optional(string)
    destination_address_prefixes               = optional(list(string))
    destination_application_security_group_ids = optional(list(string))
    access                                     = string
    priority                                   = number
    direction                                  = string
  }))
  default = {}
}
