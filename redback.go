package radius

// VendorRedback Redback vendor ID
const VendorRedback = 2352

func init() {
	builtinOnce.Do(initDictionary)
	Builtin.VsaMustRegister(VendorRedback, "Client-DNS-Pri", 1, AttributeAddress)
	Builtin.VsaMustRegister(VendorRedback, "Client-DNS-Sec", 2, AttributeAddress)
	Builtin.VsaMustRegister(VendorRedback, "DHCP-Max-Leases", 3, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Context-Name", 4, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bridge-Group", 5, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "BG-Aging-Time", 6, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "BG-Path-Cost", 7, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "BG-Span-Dis", 8, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "BG-Trans-BPDU", 9, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Rate-Limit-Rate", 10, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Rate-Limit-Burst", 11, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Police-Rate", 12, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Police-Burst", 13, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Source-Validation", 14, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Domain", 15, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Local-Name", 16, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Remote-Name", 17, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Function", 18, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Flow-Control", 19, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Static", 20, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Max-Sessions", 21, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Max-Tunnels", 22, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Session-Auth", 23, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Window", 24, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Retransmit", 25, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Cmd-Timeout", 26, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "PPPOE-URL", 27, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "PPPOE-MOTM", 28, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Group", 29, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Context", 30, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Algorithm", 31, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Deadtime", 32, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Mcast-Send", 33, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Mcast-Receive", 34, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Mcast-MaxGroups", 35, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Ip-Address-Pool-Name", 36, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-DNIS", 37, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Medium-Type", 38, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "PVC-Encapsulation-Type", 39, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "PVC-Profile-Name", 40, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "PVC-Circuit-Padding", 41, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Type", 42, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Auth-Protocol", 43, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Auth-Max-Sessions", 44, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Bypass-Bypass", 45, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Auth-Context", 46, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Auth-Service-Grp", 47, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Bypass-Context", 48, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Int-Context", 49, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Tun-Context", 50, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Ses-Context", 51, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Dot1q-Slot", 52, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Dot1q-Port", 53, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Dot1q-Vlan-Tag-Id", 54, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Int-Interface-Name", 55, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-L2TP-Tunnel-Name", 56, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-L2TP-Flow-Control", 57, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Sub-User-At-Context", 58, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Sub-Password", 59, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Ip-Host-Addr", 60, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "IP-TOS-Field", 61, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "NAS-Real-Port", 62, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Session-Auth-Ctx", 63, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Session-Auth-Service-Grp", 64, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Rate-Limit-Rate", 65, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Rate-Limit-Burst", 66, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Police-Rate", 67, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Police-Burst", 68, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-L2F-Second-Password", 69, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "ACL-Definition", 70, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "PPPoE-IP-Route-Add", 71, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "TTY-Level-Max", 72, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "TTY-Level-Start", 73, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Checksum", 74, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Profile", 75, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Client-VPN", 78, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Server-VPN", 79, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Client-Rhost", 80, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Server-Rhost", 81, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Client-Int-Addr", 82, AttributeAddress)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Server-Int-Addr", 83, AttributeAddress)
	Builtin.VsaMustRegister(VendorRedback, "PPP-Compression", 84, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Hello-Timer", 85, AttributeInteger) // has_tag
	Builtin.VsaMustRegister(VendorRedback, "Redback-Reason", 86, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Qos-Policing-Profile-Name", 87, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Qos-Metering-Profile-Name", 88, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Qos-Policy-Queuing", 89, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "IGMP-Service-Profile-Name", 90, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Subscriber-Profile-Name", 91, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Forward-Policy", 92, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Remote-Port", 93, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Reauth", 94, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Reauth-More", 95, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Agent-Remote-Id", 96, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Agent-Circuit-Id", 97, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Platform-Type", 98, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Client-NBNS-Pri", 99, AttributeAddress)
	Builtin.VsaMustRegister(VendorRedback, "Client-NBNS-Sec", 100, AttributeAddress)
	Builtin.VsaMustRegister(VendorRedback, "Shaping-Profile-Name", 101, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "BG-Cct-Addr-Max", 103, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "IP-Interface-Name", 104, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "NAT-Policy-Name", 105, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "RB-NPM-Service-Id", 106, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "HTTP-Redirect-Profile-Name", 107, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Auto-Sub-User", 108, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Auto-Sub-Context", 109, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Bind-Auto-Sub-Password", 110, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Circuit-Protocol-Encap", 111, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "OS-Version", 112, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Session-Traffic-Limit", 113, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "QOS-Reference", 114, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Rate-Limit-Excess-Burst", 121, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Police-Excess-Burst", 122, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Rate-Limit-Excess-Burst", 123, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Tunnel-Police-Excess-Burst", 124, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "DHCP-Vendor-Class-ID", 125, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Qos-Rate", 126, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "DHCP-Vendor-Encap-Option", 127, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Input-AttributeString-64", 128, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Output-AttributeString-64", 129, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Input-Packets-64", 130, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Output-Packets-64", 131, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Assigned-IP-Address", 132, AttributeAddress)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Mcast-In-AttributeString-64", 133, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Mcast-Out-AttributeString-64", 134, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Mcast-In-Packets-64", 135, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Mcast-Out-Packets-64", 136, AttributeString)
	Builtin.VsaMustRegister(VendorRedback, "LAC-Port", 137, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "LAC-Real-Port", 138, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "LAC-Port-Type", 139, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "LAC-Real-Port-Type", 140, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Dyn-Ac-Ent", 141, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Session-Error-Code", 142, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Session-Error-Msg", 143, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Update-Reason", 144, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Mac-Addr", 145, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Vlan-Source-Info", 146, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Mcast-In-AttributeString", 147, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Mcast-Out-AttributeString", 148, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Mcast-In-Packets", 149, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Mcast-Out-Packets", 150, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Reauth-Session-Id", 151, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "QOS-Rate-Inbound", 156, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "QOS-Rate-Outbound", 157, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Route-Tag", 158, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "LI-Id", 159, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "LI-Md-Address", 160, AttributeAddress)
	Builtin.VsaMustRegister(VendorRedback, "LI-Md-Port", 161, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "LI-Action", 162, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "LI-Profile", 163, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Dynamic-Policy-Filter", 164, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "HTTP-Redirect-URL", 165, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Actual-Rate-Up", 166, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Actual-Rate-Down", 167, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Min-Rate-Up", 168, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Min-Rate-Down", 169, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Attainable-Rate-Up", 170, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Attainable-Rate-Down", 171, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Max-Rate-Up", 172, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Max-Rate-Down", 173, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Min-Low-Power-Rate-Up", 174, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Min-Low-Power-Rate-Down", 175, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Max-Inter-Delay-Up", 176, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Actual-Inter-Delay-Up", 177, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Max-Inter-Delay-Down", 178, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Actual-Inter-Delay-Down", 179, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Line-State", 180, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-L2-Encapsulation", 181, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Transmission-System", 182, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-PPPOA-PPPOE-Inter-Work-Flag", 183, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Actual-Rate-Down-Factor", 185, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "DSL-Combined-Line-Info", 184, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Class-Volume-limit", 186, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Class-Volume-In-Counter", 187, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Class-Volume-Out-Counter", 188, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Flow-FAC-Profile", 189, AttributeText)
	Builtin.VsaMustRegisterTag(VendorRedback, "Service-Name", 190, AttributeText)            //has_tag
	Builtin.VsaMustRegisterTag(VendorRedback, "Service-Action", 191, AttributeInteger)       //has_tag
	Builtin.VsaMustRegisterTag(VendorRedback, "Service-Parameter", 192, AttributeText)       //has_tag
	Builtin.VsaMustRegisterTag(VendorRedback, "Service-Error-Cause", 193, AttributeInteger)  //has_tag
	Builtin.VsaMustRegisterTag(VendorRedback, "Deactivate-Service-Name", 194, AttributeText) //has_tag
	Builtin.VsaMustRegister(VendorRedback, "Qos-Profile-Overhead", 195, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Dynamic-QoS-Param", 196, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Acct-Alt-Session-ID", 197, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Idle-Timeout-Threshold", 198, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "Double-Authentication", 199, AttributeInteger)
	Builtin.VsaMustRegister(VendorRedback, "SBC-Adjacency", 200, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "DHCP-Field", 201, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "DHCP-Option", 202, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Security-Service", 203, AttributeText)
	Builtin.VsaMustRegisterTag(VendorRedback, "Reauth-Service-Name", 204, AttributeText) //has_tag
	Builtin.VsaMustRegister(VendorRedback, "Flow-IP-Profile", 205, AttributeText)
	Builtin.VsaMustRegister(VendorRedback, "Radius-Throttle-Watermark", 206, AttributeInteger)
}

/*
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-RAW",1
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-ROUTE1483 2
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-AUTO1483,3
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-MULTI",4
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-BRIDGE1483 5
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-PPP",6
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-PPP-SERIAL 7
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-PPP-NLPID 8
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-PPP-AUTO",9
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-PPPOE",10
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-L2TP",11
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-PPP-LLC",12
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-FRAME-AUTO1490 13
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-FRAME-MULTI",14
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-FRAME-BRIDGE1490 15
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-FRAME-PPP",16
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-FRAME-PPP-AUTO 17
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-FRAME-PPPOE",18
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-FRAME-ROUTE1490 19
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-FRAME-L2TP",20
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-L2TP-VC-MUXED 21
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ETH",22
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ETH-PPPOE",23
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ETH-MULTI",24
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ETH-DOT1Q,25
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ETH-DOT1Q-PPPOE 26
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-MULTI-PPPOE 27
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-MULTI-IPV6OE 28
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ATM-MULTI-PPPOE-N-IPV6OE 29
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ETH-DOT1Q-TUNNEL 30
VALUE",PVC-Encapsulation-Type",AAA-ENCAPS-ETH-DOT1Q-TUNNEL-PPPOE 31

VALUE",PVC-Circuit-Padding",AAA-CIRCUIT-PADDING",1
VALUE",PVC-Circuit-Padding",AAA-CIRCUIT-NO-PADDING",2
VALUE",Bind-Type",AAA-AUTH-BIND",1
VALUE",Bind-Type",AAA-BYPASS-BIND",2
VALUE",Bind-Type",AAA-INTERFACE-BIND",3
VALUE",Bind-Type",AAA-SUBSCRIBE-BIND",4
VALUE",Bind-Type",AAA-TUNNEL-BIND",5
VALUE",Bind-Type",AAA-SESSION-BIND",6
VALUE",Bind-Type",AAA-Q8021-BIND",7
VALUE",Bind-Type",AAA-MULTI-BIND",8
VALUE",Bind-Type",AAA-DHCP-BIND",9
VALUE",Bind-Type",AAA-MULTI-BIND-SUB",10
VALUE",Bind-Type",AAA-BRIDGE-GROUP-BIND",11
VALUE",Bind-Type",AAA-VLAN-BIND",12
VALUE",Bind-Type",AAA-VLAN-GROUP-BIND",13
VALUE",Bind-Type",AAA-AUTO-SUBSCRIBER-BIND 14
VALUE",Bind-Auth-Protocol",AAA-PPP-PAP",1
VALUE",Bind-Auth-Protocol",AAA-PPP-CHAP",2
VALUE",Bind-Auth-Protocol",AAA-PPP-CHAP-WAIT",3
VALUE",Bind-Auth-Protocol",AAA-PPP-CHAP-PAP",4
VALUE",Bind-Auth-Protocol",AAA-PPP-CHAP-WAIT-PAP",5
VALUE",Bind-Auth-Protocol",AAA-PPP-EAP",6
VALUE",Bind-Auth-Protocol",AAA-PPP-PAP-CHAP",7
VALUE",Bind-Auth-Protocol",AAA-PPP-PAP-CHAP-WAIT",8

VALUE",Source-Validation",Enabled",1
VALUE",Source-Validation",Disabled",2
VALUE",Tunnel-Domain",Enabled",1
VALUE",Tunnel-Domain",Disabled",2
VALUE",Tunnel-Function",LAC-Only",1
VALUE",Tunnel-Function",LNS-Only",2
VALUE",Tunnel-Function",LAC-LNS",3
VALUE",Tunnel-Session-Auth",CHAP",1
VALUE",Tunnel-Session-Auth",PAP",2
VALUE",Tunnel-Session-Auth",CHAP-PAP",3
VALUE",Tunnel-Group",Enabled",1
VALUE",Tunnel-Group",Disabled",2
VALUE",Tunnel-Algorithm",First",1
VALUE",Tunnel-Algorithm",Load-Balance",2
VALUE",Tunnel-Algorithm",WRR",3
VALUE",Mcast-Send",NO-SEND",1
VALUE",Mcast-Send",SEND",2
VALUE",Mcast-Send",UNSOLICITED-SEND",3
VALUE",Mcast-Receive",NO-RECEIVE",1
VALUE",Mcast-Receive",RECEIVE",2

VALUE",Tunnel-DNIS",DNIS",1
VALUE",Tunnel-DNIS",DNIS-Only",2

VALUE",Platform-Type",SMS",1
VALUE",Platform-Type",SmartEdge-800,2
VALUE",Platform-Type",SE-400,3
VALUE",Platform-Type",SE-100,4

VALUE",Circuit-Protocol-Encap",ENCAPS-PPPOE",27

VALUE",Medium-Type",DSL",11
VALUE",Medium-Type",Cable",12
VALUE",Medium-Type",Wireless",13
VALUE",Medium-Type",Satellite",14

VALUE",IP-TOS-Field",normal",0
VALUE",IP-TOS-Field",min-cost-only",1
VALUE",IP-TOS-Field",max-reliability-only",2
VALUE",IP-TOS-Field",max-reliability-plus-min-cost 3
VALUE",IP-TOS-Field",max-throughput-only",4
VALUE",IP-TOS-Field",max-throughput-plus-min-cost 5
VALUE",IP-TOS-Field",max-throughput-plus-max-reliability 6
VALUE",IP-TOS-Field",max-throughput-plus-max-reliability-plus-min-cost 7
VALUE",IP-TOS-Field",min-delay-only",8
VALUE",IP-TOS-Field",min-delay-plus-min-cost",9
VALUE",IP-TOS-Field",min-delay-plus-max-reliability 10
VALUE",IP-TOS-Field",min-delay-plus-max-reliability-plus-min-cost 11
VALUE",IP-TOS-Field",min-delay-plus-max-throughput 12
VALUE",IP-TOS-Field",min-delay-plus-max-throughput-plus-min-cost 13
VALUE",IP-TOS-Field",min-delay-plus-max-throughput-plus-max-reliability 14
VALUE",IP-TOS-Field",min-delay-plus-max-throughput-plus-max-reliability-plus-min-cost 15

VALUE",LAC-Port-Type",NAS-PORT-TYPE-10BT",40
VALUE",LAC-Port-Type",NAS-PORT-TYPE-100BT",41
VALUE",LAC-Port-Type",NAS-PORT-TYPE-DS3-FR",42
VALUE",LAC-Port-Type",NAS-PORT-TYPE-DS3-ATM",43
VALUE",LAC-Port-Type",NAS-PORT-TYPE-OC3",44
VALUE",LAC-Port-Type",NAS-PORT-TYPE-HSSI",45
VALUE",LAC-Port-Type",NAS-PORT-TYPE-EIA530,46
VALUE",LAC-Port-Type",NAS-PORT-TYPE-T1",47
VALUE",LAC-Port-Type",NAS-PORT-TYPE-CHAN-T3",48
VALUE",LAC-Port-Type",NAS-PORT-TYPE-DS1-FR",49
VALUE",LAC-Port-Type",NAS-PORT-TYPE-E3-ATM",50
VALUE",LAC-Port-Type",NAS-PORT-TYPE-IMA-ATM",51
VALUE",LAC-Port-Type",NAS-PORT-TYPE-DS3-ATM-2,52
VALUE",LAC-Port-Type",NAS-PORT-TYPE-OC3-ATM-2,53
VALUE",LAC-Port-Type",NAS-PORT-TYPE-1000BSX",54
VALUE",LAC-Port-Type",NAS-PORT-TYPE-E1-FR",55
VALUE",LAC-Port-Type",NAS-PORT-TYPE-E1-ATM",56
VALUE",LAC-Port-Type",NAS-PORT-TYPE-E3-FR",57
VALUE",LAC-Port-Type",NAS-PORT-TYPE-OC3-POS",58
VALUE",LAC-Port-Type",NAS-PORT-TYPE-OC12-POS",59
VALUE",LAC-Port-Type",NAS-PORT-TYPE-PPPOE",60
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-10BT",40
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-100BT",41
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-DS3-FR",42
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-DS3-ATM",43
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-OC3",44
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-HSSI",45
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-EIA530,46
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-T1",47
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-CHAN-T3",48
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-DS1-FR",49
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-E3-ATM",50
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-IMA-ATM",51
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-DS3-ATM-2,52
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-OC3-ATM-2,53
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-1000BSX",54
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-E1-FR",55
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-E1-ATM",56
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-E3-FR",57
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-OC3-POS",58
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-OC12-POS",59
VALUE",LAC-Real-Port-Type",NAS-PORT-TYPE-PPPOE",60


VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_SESSION_UP 1
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_SESSION_DOWN 2
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_PERIODIC",3
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_DYN_AC_ENT_START 4
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_DYN_AC_ENT_STOP 5
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_DYN_AC_ENT_TIMEOUT 6
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_SUBSCRIBER_REAUTHOR 7
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_PPP_IPCP_UP 8
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_PPP_MP_LINK_UP 9
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_DHCP_IP_ADDR_GRANTED 10
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_DHCP_IP_ADDR_RELEASED 11
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_ACL_TIMERED_ACTION 12
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_ACL_ACTION 13
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_CMD",14
VALUE",Acct-Update-Reason",AAA_LOAD_ACCT_TEST",15

VALUE",DSL-Line-State",Showtime",1
VALUE",DSL-Line-State",Idle",2
VALUE",DSL-Line-State",Silent",3

VALUE",DSL-Transmission-System",ADSL1",1
VALUE",DSL-Transmission-System",ADSL2",2
VALUE",DSL-Transmission-System",ADSL2+,3
VALUE",DSL-Transmission-System",VDSL1",4
VALUE",DSL-Transmission-System",VDSL2",5
VALUE",DSL-Transmission-System",SDSL",6
VALUE",DSL-Transmission-System",UNKNOWN",7

VALUE",Service-Action",DE-ACTIVATE",0
VALUE",Service-Action",ACTIVATE-WITH-ACCT",1
VALUE",Service-Action",ACTIVATE-WITHOUT-ACCT",2

VALUE",Service-Error-Cause",Service-success",0
VALUE",Service-Error-Cause",Unsupported-Builtin.VsaMustRegister(VendorRedback,"401
VALUE",Service-Error-Cause",Missing-Builtin.VsaMustRegister(VendorRedback,"402
VALUE",Service-Error-Cause",Invalid-request",404
VALUE",Service-Error-Cause",Resource-unavailable",506
VALUE",Service-Error-Cause",Generic-service-error",550
VALUE",Service-Error-Cause",Service-not-found",551
VALUE",Service-Error-Cause",Service-already-active",552
VALUE",Service-Error-Cause",Service-accounting-disabled 553
VALUE",Service-Error-Cause",Service-duplicate-parameter 554

END-VENDOR",Redback)
*/
