# routeros_ip_firewall_raw (Resource)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `action` (String) Action to take if a packet is matched by the rule
- `chain` (String) Specifies to which chain rule will be added. If the input does not match the name of an already defined chain, a new chain will be created.

### Optional

- `address_list` (String) Name of the address list used in 'add-dst-to-address-list' and 'add-src-to-address-list' actions.
- `address_list_timeout` (String) Time interval after which the address will be removed from the address list specified by address-list parameter. Used in conjunction with add-dst-to-address-list or add-src-to-address-list actions.
- `comment` (String)
- `content` (String) Match packets that contain specified text.
- `disabled` (Boolean)
- `dscp` (Number) Matches DSCP IP header field.
- `dst_address` (String) Matches packets which destination is equal to specified IP or falls into specified IP range.
- `dst_address_list` (String) Matches destination address of a packet against user-defined address list.
- `dst_address_type` (String) Matches destination address type.
- `dst_limit` (String) Matches packets until a given rate is exceeded.
- `dst_port` (String) List of destination port numbers or port number ranges.
- `fragment` (Boolean) Matches fragmented packets. First (starting) fragment does not count. If connection tracking is enabled there will be no fragments as system automatically assembles every packet
- `hotspot` (String) Matches packets received from HotSpot clients against various HotSpot matchers.
- `icmp_options` (String) Matches ICMP type: code fields.
- `in_bridge_port` (String) Actual interface the packet has entered the router if the incoming interface is a bridge. Works only if use-ip-firewall is enabled in bridge settings.
- `in_bridge_port_list` (String) Set of interfaces defined in interface list. Works the same as in-bridge-port.
- `in_interface` (String) Interface the packet has entered the router.
- `in_interface_list` (String) Set of interfaces defined in interface list. Works the same as in-interface.
- `ingress_priority` (Number) Matches the priority of an ingress packet. Priority may be derived from VLAN, WMM, DSCP, or MPLS EXP bit.
- `ipsec_policy` (String) Matches the policy used by IPsec. Value is written in the following format: direction, policy.
- `ipv4_options` (String) Matches IPv4 header options.
- `jump_target` (String) Name of the target chain to jump to. Applicable only if action=jump.
- `limit` (String) Matches packets up to a limited rate (packet rate or bit rate). A rule using this matcher will match until this limit is reached. Parameters are written in the following format: rate[/time],burst:mode.
- `log` (Boolean) Add a message to the system log.
- `log_prefix` (String) Adds specified text at the beginning of every log message. Applicable if action=log or log=yes configured.
- `nth` (String) Matches every nth packet: nth=2,1 rule will match every first packet of 2, hence, 50% of all the traffic that is matched by the rule
- `out_bridge_port` (String) Actual interface the packet is leaving the router if the outgoing interface is a bridge. Works only if use-ip-firewall is enabled in bridge settings.
- `out_bridge_port_list` (String) Set of interfaces defined in interface list. Works the same as out-bridge-port.
- `out_interface` (String) Interface the packet is leaving the router.
- `out_interface_list` (String) Set of interfaces defined in interface list. Works the same as out-interface.
- `packet_mark` (String) Matches packets marked via mangle facility with particular packet mark. If no-mark is set, the rule will match any unmarked packet.
- `packet_size` (String) Matches packets of specified size or size range in bytes.
- `per_connection_classifier` (String) PCC matcher allows dividing traffic into equal streams with the ability to keep packets with a specific set of options in one particular stream.
- `place_before` (String) Before which position the rule will be inserted.  
	> Please check the effect of this option, as it does not work as you think!  
	> Best way to use in conjunction with a data source. See [example](../data-sources/ip_firewall.md#example-usage).
- `port` (String) Matches if any (source or destination) port matches the specified list of ports or port ranges. Applicable only if protocol is TCP or UDP
- `priority` (Number) Matches the packet's priority after a new priority has been set. Priority may be derived from VLAN, WMM, DSCP, MPLS EXP bit, or from the priority that has been set using the set-priority action.
- `protocol` (String) Matches particular IP protocol specified by protocol name or number.
- `psd` (String) Attempts to detect TCP and UDP scans. Parameters are in the following format WeightThreshold, DelayThreshold, LowPortWeight, HighPortWeight.
- `random` (Number) Matches packets randomly with a given probability.
- `src_address` (String) Matches packets which source is equal to specified IP or falls into a specified IP range.
- `src_address_list` (String) Matches source address of a packet against user-defined address list.
- `src_address_type` (String) Matches source address type.
- `src_mac_address` (String) Matches source MAC address of the packet.
- `src_port` (String) List of source ports and ranges of source ports. Applicable only if a protocol is TCP or UDP.
- `tcp_flags` (String) Matches specified TCP flags.
- `tcp_mss` (String) Matches TCP MSS value of an IP packet.
- `time` (String) Allows to create a filter based on the packets' arrival time and date or, for locally generated packets, departure time and date.
- `tls_host` (String) Allows matching HTTPS traffic based on TLS SNI hostname.
- `ttl` (String) Matches packets TTL value.

### Read-Only

- `dynamic` (Boolean) Configuration item created by software, not by management interface. It is not exported, and cannot be directly modified.
- `id` (String) The ID of this resource.
- `invalid` (Boolean)


