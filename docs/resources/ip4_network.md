# bluecat\_ip4\_network Resource

Use this resource to create an IPv4 network.

## Example Usage

```hcl
resource "bluecat_ip4_network" "network" {
    parent_id = data.bluecat_ip4_network-block-range.block.id
    name = "New Network"
    size = 256
}

output "bluecat_ip4_network_cidr" {
    value = bluecat_ip4_network.network.cidr
}
```

## Argument Reference

* `parent_id` - (Required) The object ID of the parent object that will contain the new IPv4 network.

* `name` - (Required) The display name of the IPv4 network.

* `size` - (Required) The size of the IPv4 network expressed as a power of 2.
  For example, 256 would create a /24.

* `is_larger_allowed` - (Optional) Is it ok to return a network that is larger than the size specified?
  Default is false.

* `traversal_method` - (Optional) The traversal method used to find the range to allocate the network.
  Must be one of "NO_TRAVERSAL", "DEPTH_FIRST", or "BREADTH_FIRST". Defaults to "NO_TRAVERSAL"

## Attribute Reference

The atributes returned will vary based on the object returned.

* `properties` -  The properties of the resource as returned by the API (pipe delimited).

* `type` - The type of the resource.

* `cidr` - The CIDR address of the IPv4 network.

* `allow_duplicate_host` - Duplicate host names check.

* `inherit_allow_duplicate_host` -  Duplicate host names check is inherited.

* `ping_before_assign` - The network pings an address before assignment.

* `inherit_ping_before_assign` - The network pings an address before assignment is inherited.

* `gateway` - The gateway of the IPv4 network.

* `inherit_default_domains` - Default domains are inherited.

* `default_view` - The object id of the default DNS View for the network.

* `inherit_default_view` - The default DNS Viewis inherited.

* `inherit_dns_restrictions` - DNS restrictions are inherited.

* `addresses_in_use` - The number of addresses allocated/in use on the network.

* `addresses_free` - The number of addresses unallocated/free on the network.
