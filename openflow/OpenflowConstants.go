package openflow

//ofp_version:
const (
	OFP10_VERSION = 0x01
	OFP11_VERSION = 0x02
	OFP12_VERSION = 0x03
	OFP13_VERSION = 0x04
)

//ofp_config_flags:
const (
	OFPC_FRAG_NORMAL               = 0
	OFPC_FRAG_DROP                 = 1
	OFPC_FRAG_REASM                = 2
	OFPC_FRAG_NX_MATCH             = 3
	OFPC_FRAG_MASK                 = 3
	OFPC_INVALID_TTL_TO_CONTROLLER = 1 << 2
)

//ofp_port_config:
const (
	OFPPC_PORT_DOWN    = 1 << 0
	OFPPC_NO_RECV      = 1 << 2
	OFPPC_NO_FWD       = 1 << 5
	OFPPC_NO_PACKET_IN = 1 << 6
)

//ofp_port_state:
const (
	OFPPS_LINK_DOWN = 1 << 0
)

//ofp_port_features:
const (
	OFPPF_10MB_HD  = 1 << 0
	OFPPF_10MB_FD  = 1 << 1
	OFPPF_100MB_HD = 1 << 2
	OFPPF_100MB_FD = 1 << 3
	OFPPF_1GB_HD   = 1 << 4
	OFPPF_1GB_FD   = 1 << 5
	OFPPF_10GB_FD  = 1 << 6
)

//ofp_queue_properties:
const (
	OFPQT_MIN_RATE     = 1
	OFPQT_MAX_RATE     = 2
	OFPQT_EXPERIMENTER = 0xffff
)

//ofp_capabilities:
const (
	OFPC_FLOW_STATS   = 1 << 0
	OFPC_TABLE_STATS  = 1 << 1
	OFPC_PORT_STATS   = 1 << 2
	OFPC_IP_REASM     = 1 << 5
	OFPC_QUEUE_STATS  = 1 << 6
	OFPC_ARP_MATCH_IP = 1 << 7
)

//ofp_packet_in_reason:
const (
	OFPR_NO_MATCH    = 0
	OFPR_ACTION      = 1
	OFPR_INVALID_TTL = 2
	OFPR_N_REASONS   = 3
)

//ofp_flow_mod_command:
const (
	OFPFC_ADD           = 0
	OFPFC_MODIFY        = 1
	OFPFC_MODIFY_STRICT = 2
	OFPFC_DELETE        = 3
	OFPFC_DELETE_STRICT = 4
)

//ofp_flow_mod_flags:
const (
	OFPFF_SEND_FLOW_REM = 1 << 0
	OFPFF_CHECK_OVERLAP = 1 << 1
)

//ofp_flow_removed_reason:
const (
	OFPRR_IDLE_TIMEOUT = 0
	OFPRR_HARD_TIMEOUT = 1
	OFPRR_DELETE       = 2
	OFPRR_GROUP_DELETE = 3
	OFPRR_METER_DELETE = 4
	OFPRR_EVICTION     = 5
)

//ofp_port_reason:
const (
	OFPPR_ADD    = 0
	OFPPR_DELETE = 1
	OFPPR_MODIFY = 2
)

//ofp_stats_reply_flags:
const (
	OFPSF_REPLY_MORE = 1 << 0
)

//ofp_match_type:
const (
	OFPMT_STANDARD = 0
	OFPMT_OXM      = 1
)

//ofp_group:
const (
	OFPG_MAX = 0xffffff00
	OFPG_ALL = 0xfffffffc
	OFPG_ANY = 0xffffffff
)

//ofp_group_capabilities:
const (
	OFPGFC_SELECT_WEIGHT   = 1 << 0
	OFPGFC_SELECT_LIVENESS = 1 << 1
	OFPGFC_CHAINING        = 1 << 2
	OFPGFC_CHAINING_CHECKS = 1 << 3
)

//ofp_hello_elem_type:
const (
	OFPHET_VERSIONBITMAP = 1
)

//ofp_table:
const (
	OFPTT_MAX = 0xfe
	OFPTT_ALL = 0xff
)

//ofp_table_config:
const (
	OFPTC11_TABLE_MISS_CONTROLLER = 0 << 0
	OFPTC11_TABLE_MISS_CONTINUE   = 1 << 0
	OFPTC11_TABLE_MISS_DROP       = 2 << 0
	OFPTC11_TABLE_MISS_MASK       = 3 << 0
	OFPTC14_EVICTION              = 1 << 2
	OFPTC14_VACANCY_EVENTS        = 1 << 3
)

//ofp10_capabilities:
const (
	OFPC10_STP      = 1 << 3
	OFPC10_RESERVED = 1 << 4
)

//ofp10_port_config:
const (
	OFPPC10_NO_STP      = 1 << 1
	OFPPC10_NO_RECV_STP = 1 << 3
	OFPPC10_NO_FLOOD    = 1 << 4
)

//ofp10_port_state:
const (
	OFPPS10_STP_LISTEN  = 0 << 8
	OFPPS10_STP_LEARN   = 1 << 8
	OFPPS10_STP_FORWARD = 2 << 8
	OFPPS10_STP_BLOCK   = 3 << 8
	OFPPS10_STP_MASK    = 3 << 8
)

//ofp10_port_features:
const (
	OFPPF10_COPPER     = 1 << 7
	OFPPF10_FIBER      = 1 << 8
	OFPPF10_AUTONEG    = 1 << 9
	OFPPF10_PAUSE      = 1 << 10
	OFPPF10_PAUSE_ASYM = 1 << 11
)

//ofp10_action_type:
const (
	OFPAT10_OUTPUT       = 0
	OFPAT10_SET_VLAN_VID = 1
	OFPAT10_SET_VLAN_PCP = 2
	OFPAT10_STRIP_VLAN   = 3
	OFPAT10_SET_DL_SRC   = 4
	OFPAT10_SET_DL_DST   = 5
	OFPAT10_SET_NW_SRC   = 6
	OFPAT10_SET_NW_DST   = 7
	OFPAT10_SET_NW_TOS   = 8
	OFPAT10_SET_TP_SRC   = 9
	OFPAT10_SET_TP_DST   = 10
	OFPAT10_ENQUEUE      = 11
	OFPAT10_VENDOR       = 0xffff
)

//ofp10_flow_wildcards:
const (
	OFPFW10_IN_PORT      = 1 << 0
	OFPFW10_DL_VLAN      = 1 << 1
	OFPFW10_DL_SRC       = 1 << 2
	OFPFW10_DL_DST       = 1 << 3
	OFPFW10_DL_TYPE      = 1 << 4
	OFPFW10_NW_PROTO     = 1 << 5
	OFPFW10_TP_SRC       = 1 << 6
	OFPFW10_TP_DST       = 1 << 7
	OFPFW10_NW_SRC_SHIFT = 8
	OFPFW10_NW_SRC_BITS  = 6
	OFPFW10_NW_SRC_MASK  = 1<<OFPFW10_NW_SRC_BITS - 1<<OFPFW10_NW_SRC_SHIFT
	OFPFW10_NW_SRC_ALL   = 32 << OFPFW10_NW_SRC_SHIFT
	OFPFW10_NW_DST_SHIFT = 14
	OFPFW10_NW_DST_BITS  = 6
	OFPFW10_NW_DST_MASK  = 1<<OFPFW10_NW_DST_BITS - 1<<OFPFW10_NW_DST_SHIFT
	OFPFW10_NW_DST_ALL   = 32 << OFPFW10_NW_DST_SHIFT
	OFPFW10_DL_VLAN_PCP  = 1 << 20
	OFPFW10_NW_TOS       = 1 << 21
	OFPFW10_ALL          = 1<<22 - 1
)

//ofp10_flow_mod_flags:
const (
	OFPFF10_EMERG = 1 << 2
)

//ofp11_port_state:
const (
	OFPPS11_BLOCKED = 1 << 1
	OFPPS11_LIVE    = 1 << 2
)

//ofp11_port_features:
const (
	OFPPF11_40GB_FD    = 1 << 7
	OFPPF11_100GB_FD   = 1 << 8
	OFPPF11_1TB_FD     = 1 << 9
	OFPPF11_OTHER      = 1 << 10
	OFPPF11_COPPER     = 1 << 11
	OFPPF11_FIBER      = 1 << 12
	OFPPF11_AUTONEG    = 1 << 13
	OFPPF11_PAUSE      = 1 << 14
	OFPPF11_PAUSE_ASYM = 1 << 15
)

//ofp11_group_mod_command:
const (
	OFPGC11_ADD    = 0
	OFPGC11_MODIFY = 1
	OFPGC11_DELETE = 2
)

//ofp11_capabilities:
const (
	OFPC11_GROUP_STATS = 1 << 3
)

//ofp11_action_type:
const (
	OFPAT11_OUTPUT         = 0
	OFPAT11_SET_VLAN_VID   = 1
	OFPAT11_SET_VLAN_PCP   = 2
	OFPAT11_SET_DL_SRC     = 3
	OFPAT11_SET_DL_DST     = 4
	OFPAT11_SET_NW_SRC     = 5
	OFPAT11_SET_NW_DST     = 6
	OFPAT11_SET_NW_TOS     = 7
	OFPAT11_SET_NW_ECN     = 8
	OFPAT11_SET_TP_SRC     = 9
	OFPAT11_SET_TP_DST     = 10
	OFPAT11_COPY_TTL_OUT   = 11
	OFPAT11_COPY_TTL_IN    = 12
	OFPAT11_SET_MPLS_LABEL = 13
	OFPAT11_SET_MPLS_TC    = 14
	OFPAT11_SET_MPLS_TTL   = 15
	OFPAT11_DEC_MPLS_TTL   = 16
	OFPAT11_PUSH_VLAN      = 17
	OFPAT11_POP_VLAN       = 18
	OFPAT11_PUSH_MPLS      = 19
	OFPAT11_POP_MPLS       = 20
	OFPAT11_SET_QUEUE      = 21
	OFPAT11_GROUP          = 22
	OFPAT11_SET_NW_TTL     = 23
	OFPAT11_DEC_NW_TTL     = 24
	OFPAT11_EXPERIMENTER   = 0xffff
)

//ofp11_flow_wildcards:
const (
	OFPFW11_IN_PORT     = 1 << 0
	OFPFW11_DL_VLAN     = 1 << 1
	OFPFW11_DL_VLAN_PCP = 1 << 2
	OFPFW11_DL_TYPE     = 1 << 3
	OFPFW11_NW_TOS      = 1 << 4
	OFPFW11_NW_PROTO    = 1 << 5
	OFPFW11_TP_SRC      = 1 << 6
	OFPFW11_TP_DST      = 1 << 7
	OFPFW11_MPLS_LABEL  = 1 << 8
	OFPFW11_MPLS_TC     = 1 << 9
	OFPFW11_ALL         = 1<<10 - 1
)

//ofp11_vlan_id:
const (
	OFPVID11_ANY  = 0xfffe
	OFPVID11_NONE = 0xffff
)

//ofp11_instruction_type:
const (
	OFPIT11_GOTO_TABLE     = 1
	OFPIT11_WRITE_METADATA = 2
	OFPIT11_WRITE_ACTIONS  = 3
	OFPIT11_APPLY_ACTIONS  = 4
	OFPIT11_CLEAR_ACTIONS  = 5
	OFPIT11_EXPERIMENTER   = 0xFFFF
)

//ofp11_group_type:
const (
	OFPGT11_ALL      = 0
	OFPGT11_SELECT   = 1
	OFPGT11_INDIRECT = 2
	OFPGT11_FF       = 3
)

//ofp11_group:
const (
	OFPG11_MAX = 0xffffff00
	OFPG11_ALL = 0xfffffffc
	OFPG11_ANY = 0xffffffff
)

//ofp11_flow_match_fields:
const (
	OFPFMF11_IN_PORT     = 1 << 0
	OFPFMF11_DL_VLAN     = 1 << 1
	OFPFMF11_DL_VLAN_PCP = 1 << 2
	OFPFMF11_DL_TYPE     = 1 << 3
	OFPFMF11_NW_TOS      = 1 << 4
	OFPFMF11_NW_PROTO    = 1 << 5
	OFPFMF11_TP_SRC      = 1 << 6
	OFPFMF11_TP_DST      = 1 << 7
	OFPFMF11_MPLS_LABEL  = 1 << 8
	OFPFMF11_MPLS_TC     = 1 << 9
	OFPFMF11_TYPE        = 1 << 10
	OFPFMF11_DL_SRC      = 1 << 11
	OFPFMF11_DL_DST      = 1 << 12
	OFPFMF11_NW_SRC      = 1 << 13
	OFPFMF11_NW_DST      = 1 << 14
	OFPFMF11_METADATA    = 1 << 15
)

//ofp12_oxm_class:
const (
	OFPXMC12_NXM_0          = 0x0000
	OFPXMC12_NXM_1          = 0x0001
	OFPXMC12_OPENFLOW_BASIC = 0x8000
	OFPXMC12_EXPERIMENTER   = 0xffff
)

//ofp12_vlan_id:
const (
	OFPVID12_PRESENT = 0x1000
	OFPVID12_NONE    = 0x0000
)

//ofp12_ipv6exthdr_flags:
const (
	OFPIEH12_NONEXT = 1 << 0
	OFPIEH12_ESP    = 1 << 1
	OFPIEH12_AUTH   = 1 << 2
	OFPIEH12_DEST   = 1 << 3
	OFPIEH12_FRAG   = 1 << 4
	OFPIEH12_ROUTER = 1 << 5
	OFPIEH12_HOP    = 1 << 6
	OFPIEH12_UNREP  = 1 << 7
	OFPIEH12_UNSEQ  = 1 << 8
)

//ofp12_action_type:
const (
	OFPAT12_SET_FIELD = 25
)

//ofp12_controller_max_len:
const (
	OFPCML12_MAX       = 0xffe5
	OFPCML12_NO_BUFFER = 0xffff
)

//ofp12_flow_mod_flags:
const (
	OFPFF12_RESET_COUNTS = 1 << 2
)

//ofp12_capabilities:
const (
	OFPC12_PORT_BLOCKED = 1 << 8
)

//ofp12_group_capabilities:
const (
	OFPGFC12_SELECT_WEIGHT   = 1 << 0
	OFPGFC12_SELECT_LIVENESS = 1 << 1
	OFPGFC12_CHAINING        = 1 << 2
	OFPGFC12_CHAINING_CHECKS = 1 << 3
)

//ofp12_controller_role:
const (
	OFPCR12_ROLE_NOCHANGE = 0
	OFPCR12_ROLE_EQUAL    = 1
	OFPCR12_ROLE_MASTER   = 2
	OFPCR12_ROLE_SLAVE    = 3
)

//ofp13_instruction_type:
const (
	OFPIT13_METER = 6
)

//ofp13_action_type:
const (
	OFPAT13_PUSH_PBB = 26
	OFPAT13_POP_PBB  = 27
)

//ofp13_table_config:
const (
	OFPTC13_DEPRECATED_MASK = 3
)

//ofp13_flow_mod_flags:
const (
	OFPFF13_NO_PKT_COUNTS = 1 << 3
	OFPFF13_NO_BYT_COUNTS = 1 << 4
)

//ofp13_meter:
const (
	OFPM13_MAX        = 0xffff0000
	OFPM13_SLOWPATH   = 0xfffffffd
	OFPM13_CONTROLLER = 0xfffffffe
	OFPM13_ALL        = 0xffffffff
)

//ofp13_meter_mod_command:
const (
	OFPMC13_ADD    = 0
	OFPMC13_MODIFY = 1
	OFPMC13_DELETE = 2
)

//ofp13_meter_flags:
const (
	OFPMF13_KBPS  = 1 << 0
	OFPMF13_PKTPS = 1 << 1
	OFPMF13_BURST = 1 << 2
	OFPMF13_STATS = 1 << 3
)

//ofp13_meter_band_type:
const (
	OFPMBT13_DROP         = 1
	OFPMBT13_DSCP_REMARK  = 2
	OFPMBT13_EXPERIMENTER = 0xFFFF
)

//ofp13_multipart_request_flags:
const (
	OFPMPF13_REQ_MORE = 1 << 0
)

//ofp13_table_feature_prop_type:
const (
	OFPTFPT13_INSTRUCTIONS        = 0
	OFPTFPT13_INSTRUCTIONS_MISS   = 1
	OFPTFPT13_NEXT_TABLES         = 2
	OFPTFPT13_NEXT_TABLES_MISS    = 3
	OFPTFPT13_WRITE_ACTIONS       = 4
	OFPTFPT13_WRITE_ACTIONS_MISS  = 5
	OFPTFPT13_APPLY_ACTIONS       = 6
	OFPTFPT13_APPLY_ACTIONS_MISS  = 7
	OFPTFPT13_MATCH               = 8
	OFPTFPT13_WILDCARDS           = 10
	OFPTFPT13_WRITE_SETFIELD      = 12
	OFPTFPT13_WRITE_SETFIELD_MISS = 13
	OFPTFPT13_APPLY_SETFIELD      = 14
	OFPTFPT13_APPLY_SETFIELD_MISS = 15
	OFPTFPT13_EXPERIMENTER        = 0xFFFE
	OFPTFPT13_EXPERIMENTER_MISS   = 0xFFFF
)

//ofp14_async_config_prop_type:
const (
	OFPACPT_PACKET_IN_SLAVE       = 0
	OFPACPT_PACKET_IN_MASTER      = 1
	OFPACPT_PORT_STATUS_SLAVE     = 2
	OFPACPT_PORT_STATUS_MASTER    = 3
	OFPACPT_FLOW_REMOVED_SLAVE    = 4
	OFPACPT_FLOW_REMOVED_MASTER   = 5
	OFPACPT_ROLE_STATUS_SLAVE     = 6
	OFPACPT_ROLE_STATUS_MASTER    = 7
	OFPACPT_TABLE_STATUS_SLAVE    = 8
	OFPACPT_TABLE_STATUS_MASTER   = 9
	OFPACPT_REQUESTFORWARD_SLAVE  = 10
	OFPACPT_REQUESTFORWARD_MASTER = 11
	OFPTFPT_EXPERIMENTER_SLAVE    = 0xFFFE
	OFPTFPT_EXPERIMENTER_MASTER   = 0xFFFF
)

//ofp14_controller_role_reason:
const (
	OFPCRR_MASTER_REQUEST = 0
	OFPCRR_CONFIG         = 1
	OFPCRR_EXPERIMENTER   = 2
)

//ofp14_role_prop_type:
const (
	OFPRPT_EXPERIMENTER = 0xFFFF
)
