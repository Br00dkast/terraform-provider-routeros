package routeros

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/*
{
    ".id": "*4",
    "action": "echo",
    "default": "true",
    "disabled": "false",
    "invalid": "false",
    "prefix": "",
    "topics": "critical"
}
*/

var validTopics = []string{
	"account", "async", "backup", "bfd", "bgp", "bridge", "calc", "caps", "certificate", "clock", "container", "critical",
	"ddns", "debug", "dhcp", "disk", "dns", "dot1x", "dude", "e-mail", "error", "event", "fetch", "firewall", "gps", "gsm",
	"health", "hotspot", "igmp-proxy", "info", "interface", "ipsec", "iscsi", "isdn", "isis", "kvm", "l2tp", "ldp",
	"lora", "lte", "manager", "mme", "mpls", "mqtt", "natpmp", "netinstall", "netwatch", "ntp", "ospf", "ovpn",
	"packet", "pim", "poe-out", "ppp", "pppoe", "pptp", "queue", "radius", "radvd", "raw", "read", "rip", "route",
	"rpki", "rsvp", "script", "sertcp", "simulator", "smb", "snmp", "ssh", "sstp", "state", "store", "stp", "system",
	"telephony", "tftp", "timer", "tr069", "update", "upnp", "ups", "vpls", "vrrp", "warning", "watchdog", "web-proxy",
	"wireguard", "wireless", "write",
}

// ResourceSystemLogging defines the resource for configuring logging rules
// https://help.mikrotik.com/docs/display/ROS/Log
func ResourceSystemLogging() *schema.Resource {
	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/system/logging"),
		MetaId:           PropId(Id),

		"action": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "specifies one of the system default actions or user specified action listed in actions menu",
		},
		KeyDefault:  PropDefaultRo,
		KeyDisabled: PropDisabledRw,
		KeyInvalid:  PropInvalidRo,
		"prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "prefix added at the beginning of log messages",
			Default:     "",
		},
		"regex": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Regex which will be used in order to match or not match message. If the regex is not matched, " +
				"then even if topic is configured to be logged, but log message does not match regex, action will not " +
				"be performed.",
		},
		"topics": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: ValidationValInSlice(validTopics, false, true),
			},
			Optional: true,
			Description: `log all messages that falls into specified topic or list of topics.
						  '!' character can be used before topic to exclude messages falling under this topic. For example, we want to log NTP debug info without too much details:
						  /system logging add topics=ntp,debug,!packet`,
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: ImportStateCustomContext(resSchema),
		},

		Schema: resSchema,
	}
}
