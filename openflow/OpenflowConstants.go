package openflow

type OfpPortNo uint32

const (
	OFPP_MAX        OfpPortNo = 0xffffff00
	OFPP_IN_PORT              = 0xfffffff8
	OFPP_TABLE                = 0xfffffff9
	OFPP_NORMAL               = 0xfffffffa
	OFPP_FLOOD                = 0xfffffffb
	OFPP_ALL                  = 0xfffffffc
	OFPP_CONTROLLER           = 0xfffffffd
	OFPP_LOCAL                = 0xfffffffe
	OFPP_ANY                  = 0xffffffff
)

type OfpType uint32

const (
	OFPT_HELLO                    OfpType = 0
	OFPT_ERROR                            = 1
	OFPT_ECHO_REQUEST                     = 2
	OFPT_ECHO_REPLY                       = 3
	OFPT_EXPERIMENTER                     = 4
	OFPT_FEATURES_REQUEST                 = 5
	OFPT_FEATURES_REPLY                   = 6
	OFPT_GET_CONFIG_REQUEST               = 7
	OFPT_GET_CONFIG_REPLY                 = 8
	OFPT_SET_CONFIG                       = 9
	OFPT_PACKET_IN                        = 10
	OFPT_FLOW_REMOVED                     = 11
	OFPT_PORT_STATUS                      = 12
	OFPT_PACKET_OUT                       = 13
	OFPT_FLOW_MOD                         = 14
	OFPT_GROUP_MOD                        = 15
	OFPT_PORT_MOD                         = 16
	OFPT_TABLE_MOD                        = 17
	OFPT_MULTIPART_REQUEST                = 18
	OFPT_MULTIPART_REPLY                  = 19
	OFPT_BARRIER_REQUEST                  = 20
	OFPT_BARRIER_REPLY                    = 21
	OFPT_QUEUE_GET_CONFIG_REQUEST         = 22
	OFPT_QUEUE_GET_CONFIG_REPLY           = 23
	OFPT_ROLE_REQUEST                     = 24
	OFPT_ROLE_REPLY                       = 25
	OFPT_GET_ASYNC_REQUEST                = 26
	OFPT_GET_ASYNC_REPLY                  = 27
	OFPT_SET_ASYNC                        = 28
	OFPT_METER_MOD                        = 29
)

type OfpHelloElemType uint32

const (
	OFPHET_VERSIONBITMAP OfpHelloElemType = 1
)

type OfpConfigFlags uint32

const (
	OFPC_FRAG_NORMAL OfpConfigFlags = 0
	OFPC_FRAG_DROP                  = 1 << 0
	OFPC_FRAG_REASM                 = 1 << 1
	OFPC_FRAG_MASK                  = 3
)

type OfpTableConfig uint32

const (
	OFPTC_DEPRECATED_MASK OfpTableConfig = 3
)

type OfpTable uint32

const (
	OFPTT_MAX OfpTable = 0xfe
	OFPTT_ALL          = 0xff
)

type OfpCapabilities uint32

const (
	OFPC_FLOW_STATS   OfpCapabilities = 1 << 0
	OFPC_TABLE_STATS                  = 1 << 1
	OFPC_PORT_STATS                   = 1 << 2
	OFPC_GROUP_STATS                  = 1 << 3
	OFPC_IP_REASM                     = 1 << 5
	OFPC_QUEUE_STATS                  = 1 << 6
	OFPC_PORT_BLOCKED                 = 1 << 8
)

type OfpPortConfig uint32

const (
	OFPPC_PORT_DOWN    OfpPortConfig = 1 << 0
	OFPPC_NO_RECV                    = 1 << 2
	OFPPC_NO_FWD                     = 1 << 5
	OFPPC_NO_PACKET_IN               = 1 << 6
)

type OfpPortState uint32

const (
	OFPPS_LINK_DOWN OfpPortState = 1 << 0
	OFPPS_BLOCKED                = 1 << 1
	OFPPS_LIVE                   = 1 << 2
)

type OfpPortFeatures uint32

const (
	OFPPF_10MB_HD    OfpPortFeatures = 1 << 0
	OFPPF_10MB_FD                    = 1 << 1
	OFPPF_100MB_HD                   = 1 << 2
	OFPPF_100MB_FD                   = 1 << 3
	OFPPF_1GB_HD                     = 1 << 4
	OFPPF_1GB_FD                     = 1 << 5
	OFPPF_10GB_FD                    = 1 << 6
	OFPPF_40GB_FD                    = 1 << 7
	OFPPF_100GB_FD                   = 1 << 8
	OFPPF_1TB_FD                     = 1 << 9
	OFPPF_OTHER                      = 1 << 10
	OFPPF_COPPER                     = 1 << 11
	OFPPF_FIBER                      = 1 << 12
	OFPPF_AUTONEG                    = 1 << 13
	OFPPF_PAUSE                      = 1 << 14
	OFPPF_PAUSE_ASYM                 = 1 << 15
)

type OfpPortReason uint32

const (
	OFPPR_ADD    OfpPortReason = 0
	OFPPR_DELETE               = 1
	OFPPR_MODIFY               = 2
)

type OfpMatchType uint32

const (
	OFPMT_STANDARD OfpMatchType = 0
	OFPMT_OXM                   = 1
)

type OfpOxmClass uint32

const (
	OFPXMC_NXM_0          OfpOxmClass = 0x0000
	OFPXMC_NXM_1                      = 0x0001
	OFPXMC_OPENFLOW_BASIC             = 0x8000
	OFPXMC_EXPERIMENTER               = 0xFFFF
)

type OfpVlanId uint32

const (
	OFPVID_PRESENT OfpVlanId = 0x1000
	OFPVID_NONE              = 0x0000
)

type OfpIpv6exthdrFlags uint32

const (
	OFPIEH_NONEXT OfpIpv6exthdrFlags = 1 << 0
	OFPIEH_ESP                       = 1 << 1
	OFPIEH_AUTH                      = 1 << 2
	OFPIEH_DEST                      = 1 << 3
	OFPIEH_FRAG                      = 1 << 4
	OFPIEH_ROUTER                    = 1 << 5
	OFPIEH_HOP                       = 1 << 6
	OFPIEH_UNREP                     = 1 << 7
	OFPIEH_UNSEQ                     = 1 << 8
)

type OfpActionType uint32

const (
	OFPAT_OUTPUT       OfpActionType = 0
	OFPAT_COPY_TTL_OUT               = 11
	OFPAT_COPY_TTL_IN                = 12
	OFPAT_SET_MPLS_TTL               = 15
	OFPAT_DEC_MPLS_TTL               = 16
	OFPAT_PUSH_VLAN                  = 17
	OFPAT_POP_VLAN                   = 18
	OFPAT_PUSH_MPLS                  = 19
	OFPAT_POP_MPLS                   = 20
	OFPAT_SET_QUEUE                  = 21
	OFPAT_GROUP                      = 22
	OFPAT_SET_NW_TTL                 = 23
	OFPAT_DEC_NW_TTL                 = 24
	OFPAT_SET_FIELD                  = 25
	OFPAT_PUSH_PBB                   = 26
	OFPAT_POP_PBB                    = 27
	OFPAT_EXPERIMENTER               = 0xffff
)

type OfpControllerMaxLen uint32

const (
	OFPCML_MAX       OfpControllerMaxLen = 0xffe5
	OFPCML_NO_BUFFER                     = 0xffff
)

type OfpInstructionType uint32

const (
	OFPIT_GOTO_TABLE     OfpInstructionType = 1
	OFPIT_WRITE_METADATA                    = 2
	OFPIT_WRITE_ACTIONS                     = 3
	OFPIT_APPLY_ACTIONS                     = 4
	OFPIT_CLEAR_ACTIONS                     = 5
	OFPIT_METER                             = 6
	OFPIT_EXPERIMENTER                      = 0xFFFF
)

type OfpFlowModCommand uint32

const (
	OFPFC_ADD           OfpFlowModCommand = 0
	OFPFC_MODIFY                          = 1
	OFPFC_MODIFY_STRICT                   = 2
	OFPFC_DELETE                          = 3
	OFPFC_DELETE_STRICT                   = 4
)

type OfpFlowModFlags uint32

const (
	OFPFF_SEND_FLOW_REM OfpFlowModFlags = 1 << 0
	OFPFF_CHECK_OVERLAP                 = 1 << 1
	OFPFF_RESET_COUNTS                  = 1 << 2
	OFPFF_NO_PKT_COUNTS                 = 1 << 3
	OFPFF_NO_BYT_COUNTS                 = 1 << 4
)

type OfpGroup uint32

const (
	OFPG_MAX OfpGroup = 0xffffff00
	OFPG_ALL          = 0xfffffffc
	OFPG_ANY          = 0xffffffff
)

type OfpGroupModCommand uint32

const (
	OFPGC_ADD    OfpGroupModCommand = 0
	OFPGC_MODIFY                    = 1
	OFPGC_DELETE                    = 2
)

type OfpGroupType uint32

const (
	OFPGT_ALL      OfpGroupType = 0
	OFPGT_SELECT                = 1
	OFPGT_INDIRECT              = 2
	OFPGT_FF                    = 3
)

type OfpPacketInReason uint32

const (
	OFPR_NO_MATCH    OfpPacketInReason = 0
	OFPR_ACTION                        = 1
	OFPR_INVALID_TTL                   = 2
)

type OfpFlowRemovedReason uint32

const (
	OFPRR_IDLE_TIMEOUT OfpFlowRemovedReason = 0
	OFPRR_HARD_TIMEOUT                      = 1
	OFPRR_DELETE                            = 2
	OFPRR_GROUP_DELETE                      = 3
)

type OfpMeter uint32

const (
	OFPM_MAX        OfpMeter = 0xffff0000
	OFPM_SLOWPATH            = 0xfffffffd
	OFPM_CONTROLLER          = 0xfffffffe
	OFPM_ALL                 = 0xffffffff
)

type OfpMeterBandType uint32

const (
	OFPMBT_DROP         OfpMeterBandType = 1
	OFPMBT_DSCP_REMARK                   = 2
	OFPMBT_EXPERIMENTER                  = 0xFFFF
)

type OfpMeterModCommand uint32

const (
	OFPMC_ADD    OfpMeterModCommand = 0
	OFPMC_MODIFY                    = 1
	OFPMC_DELETE                    = 2
)

type OfpMeterFlags uint32

const (
	OFPMF_KBPS  OfpMeterFlags = 1 << 0
	OFPMF_PKTPS               = 1 << 1
	OFPMF_BURST               = 1 << 2
	OFPMF_STATS               = 1 << 3
)

type OfpErrorType uint32

const (
	OFPET_HELLO_FAILED          OfpErrorType = 0
	OFPET_BAD_REQUEST                        = 1
	OFPET_BAD_ACTION                         = 2
	OFPET_BAD_INSTRUCTION                    = 3
	OFPET_BAD_MATCH                          = 4
	OFPET_FLOW_MOD_FAILED                    = 5
	OFPET_GROUP_MOD_FAILED                   = 6
	OFPET_PORT_MOD_FAILED                    = 7
	OFPET_TABLE_MOD_FAILED                   = 8
	OFPET_QUEUE_OP_FAILED                    = 9
	OFPET_SWITCH_CONFIG_FAILED               = 10
	OFPET_ROLE_REQUEST_FAILED                = 11
	OFPET_METER_MOD_FAILED                   = 12
	OFPET_TABLE_FEATURES_FAILED              = 13
	OFPET_EXPERIMENTER                       = 0xffff
)

type OfpHelloFailedCode uint32

const (
	OFPHFC_INCOMPATIBLE OfpHelloFailedCode = 0
	OFPHFC_EPERM                           = 1
)

type OfpBadRequestCode uint32

const (
	OFPBRC_BAD_VERSION               OfpBadRequestCode = 0
	OFPBRC_BAD_TYPE                                    = 1
	OFPBRC_BAD_MULTIPART                               = 2
	OFPBRC_BAD_EXPERIMENTER                            = 3
	OFPBRC_BAD_EXP_TYPE                                = 4
	OFPBRC_EPERM                                       = 5
	OFPBRC_BAD_LEN                                     = 6
	OFPBRC_BUFFER_EMPTY                                = 7
	OFPBRC_BUFFER_UNKNOWN                              = 8
	OFPBRC_BAD_TABLE_ID                                = 9
	OFPBRC_IS_SLAVE                                    = 10
	OFPBRC_BAD_PORT                                    = 11
	OFPBRC_BAD_PACKET                                  = 12
	OFPBRC_MULTIPART_BUFFER_OVERFLOW                   = 13
)

type OfpBadActionCode uint32

const (
	OFPBAC_BAD_TYPE           OfpBadActionCode = 0
	OFPBAC_BAD_LEN                             = 1
	OFPBAC_BAD_EXPERIMENTER                    = 2
	OFPBAC_BAD_EXP_TYPE                        = 3
	OFPBAC_BAD_OUT_PORT                        = 4
	OFPBAC_BAD_ARGUMENT                        = 5
	OFPBAC_EPERM                               = 6
	OFPBAC_TOO_MANY                            = 7
	OFPBAC_BAD_QUEUE                           = 8
	OFPBAC_BAD_OUT_GROUP                       = 9
	OFPBAC_MATCH_INCONSISTENT                  = 10
	OFPBAC_UNSUPPORTED_ORDER                   = 11
	OFPBAC_BAD_TAG                             = 12
	OFPBAC_BAD_SET_TYPE                        = 13
	OFPBAC_BAD_SET_LEN                         = 14
	OFPBAC_BAD_SET_ARGUMENT                    = 15
)

type OfpBadInstructionCode uint32

const (
	OFPBIC_UNKNOWN_INST        OfpBadInstructionCode = 0
	OFPBIC_UNSUP_INST                                = 1
	OFPBIC_BAD_TABLE_ID                              = 2
	OFPBIC_UNSUP_METADATA                            = 3
	OFPBIC_UNSUP_METADATA_MASK                       = 4
	OFPBIC_BAD_EXPERIMENTER                          = 5
	OFPBIC_BAD_EXP_TYPE                              = 6
	OFPBIC_BAD_LEN                                   = 7
	OFPBIC_EPERM                                     = 8
)

type OfpBadMatchCode uint32

const (
	OFPBMC_BAD_TYPE         OfpBadMatchCode = 0
	OFPBMC_BAD_LEN                          = 1
	OFPBMC_BAD_TAG                          = 2
	OFPBMC_BAD_DL_ADDR_MASK                 = 3
	OFPBMC_BAD_NW_ADDR_MASK                 = 4
	OFPBMC_BAD_WILDCARDS                    = 5
	OFPBMC_BAD_FIELD                        = 6
	OFPBMC_BAD_VALUE                        = 7
	OFPBMC_BAD_MASK                         = 8
	OFPBMC_BAD_PREREQ                       = 9
	OFPBMC_DUP_FIELD                        = 10
	OFPBMC_EPERM                            = 11
)

type OfpFlowModFailedCode uint32

const (
	OFPFMFC_UNKNOWN      OfpFlowModFailedCode = 0
	OFPFMFC_TABLE_FULL                        = 1
	OFPFMFC_BAD_TABLE_ID                      = 2
	OFPFMFC_OVERLAP                           = 3
	OFPFMFC_EPERM                             = 4
	OFPFMFC_BAD_TIMEOUT                       = 5
	OFPFMFC_BAD_COMMAND                       = 6
	OFPFMFC_BAD_FLAGS                         = 7
)

type OfpGroupModFailedCode uint32

const (
	OFPGMFC_GROUP_EXISTS         OfpGroupModFailedCode = 0
	OFPGMFC_INVALID_GROUP                              = 1
	OFPGMFC_WEIGHT_UNSUPPORTED                         = 2
	OFPGMFC_OUT_OF_GROUPS                              = 3
	OFPGMFC_OUT_OF_BUCKETS                             = 4
	OFPGMFC_CHAINING_UNSUPPORTED                       = 5
	OFPGMFC_WATCH_UNSUPPORTED                          = 6
	OFPGMFC_LOOP                                       = 7
	OFPGMFC_UNKNOWN_GROUP                              = 8
	OFPGMFC_CHAINED_GROUP                              = 9
	OFPGMFC_BAD_TYPE                                   = 10
	OFPGMFC_BAD_COMMAND                                = 11
	OFPGMFC_BAD_BUCKET                                 = 12
	OFPGMFC_BAD_WATCH                                  = 13
	OFPGMFC_EPERM                                      = 14
)

type OfpPortModFailedCode uint32

const (
	OFPPMFC_BAD_PORT      OfpPortModFailedCode = 0
	OFPPMFC_BAD_HW_ADDR                        = 1
	OFPPMFC_BAD_CONFIG                         = 2
	OFPPMFC_BAD_ADVERTISE                      = 3
	OFPPMFC_EPERM                              = 4
)

type OfpTableModFailedCode uint32

const (
	OFPTMFC_BAD_TABLE  OfpTableModFailedCode = 0
	OFPTMFC_BAD_CONFIG                       = 1
	OFPTMFC_EPERM                            = 2
)

type OfpQueueOpFailedCode uint32

const (
	OFPQOFC_BAD_PORT  OfpQueueOpFailedCode = 0
	OFPQOFC_BAD_QUEUE                      = 1
	OFPQOFC_EPERM                          = 2
)

type OfpSwitchConfigFailedCode uint32

const (
	OFPSCFC_BAD_FLAGS OfpSwitchConfigFailedCode = 0
	OFPSCFC_BAD_LEN                             = 1
	OFPSCFC_EPERM                               = 2
)

type OfpRoleRequestFailedCode uint32

const (
	OFPRRFC_STALE    OfpRoleRequestFailedCode = 0
	OFPRRFC_UNSUP                             = 1
	OFPRRFC_BAD_ROLE                          = 2
)

type OfpMeterModFailedCode uint32

const (
	OFPMMFC_UNKNOWN        OfpMeterModFailedCode = 0
	OFPMMFC_METER_EXISTS                         = 1
	OFPMMFC_INVALID_METER                        = 2
	OFPMMFC_UNKNOWN_METER                        = 3
	OFPMMFC_BAD_COMMAND                          = 4
	OFPMMFC_BAD_FLAGS                            = 5
	OFPMMFC_BAD_RATE                             = 6
	OFPMMFC_BAD_BURST                            = 7
	OFPMMFC_BAD_BAND                             = 8
	OFPMMFC_BAD_BAND_VALUE                       = 9
	OFPMMFC_OUT_OF_METERS                        = 10
	OFPMMFC_OUT_OF_BANDS                         = 11
)

type OfpTableFeaturesFailedCode uint32

const (
	OFPTFFC_BAD_TABLE    OfpTableFeaturesFailedCode = 0
	OFPTFFC_BAD_METADATA                            = 1
	OFPTFFC_BAD_TYPE                                = 2
	OFPTFFC_BAD_LEN                                 = 3
	OFPTFFC_BAD_ARGUMENT                            = 4
	OFPTFFC_EPERM                                   = 5
)

type OfpMultipartType uint32

const (
	OFPMP_DESC           OfpMultipartType = 0
	OFPMP_FLOW                            = 1
	OFPMP_AGGREGATE                       = 2
	OFPMP_TABLE                           = 3
	OFPMP_PORT_STATS                      = 4
	OFPMP_QUEUE                           = 5
	OFPMP_GROUP                           = 6
	OFPMP_GROUP_DESC                      = 7
	OFPMP_GROUP_FEATURES                  = 8
	OFPMP_METER                           = 9
	OFPMP_METER_CONFIG                    = 10
	OFPMP_METER_FEATURES                  = 11
	OFPMP_TABLE_FEATURES                  = 12
	OFPMP_PORT_DESC                       = 13
	OFPMP_EXPERIMENTER                    = 0xffff
)

type OfpMultipartRequestFlags uint32

const (
	OFPMPF_REQ_MORE OfpMultipartRequestFlags = 1 << 0
)

type OfpMultipartReplyFlags uint32

const (
	OFPMPF_REPLY_MORE OfpMultipartReplyFlags = 1 << 0
)

type OfpTableFeaturePropType uint32

const (
	OFPTFPT_INSTRUCTIONS        OfpTableFeaturePropType = 0
	OFPTFPT_INSTRUCTIONS_MISS                           = 1
	OFPTFPT_NEXT_TABLES                                 = 2
	OFPTFPT_NEXT_TABLES_MISS                            = 3
	OFPTFPT_WRITE_ACTIONS                               = 4
	OFPTFPT_WRITE_ACTIONS_MISS                          = 5
	OFPTFPT_APPLY_ACTIONS                               = 6
	OFPTFPT_APPLY_ACTIONS_MISS                          = 7
	OFPTFPT_MATCH                                       = 8
	OFPTFPT_WILDCARDS                                   = 10
	OFPTFPT_WRITE_SETFIELD                              = 12
	OFPTFPT_WRITE_SETFIELD_MISS                         = 13
	OFPTFPT_APPLY_SETFIELD                              = 14
	OFPTFPT_APPLY_SETFIELD_MISS                         = 15
	OFPTFPT_EXPERIMENTER                                = 0xFFFE
	OFPTFPT_EXPERIMENTER_MISS                           = 0xFFFF
)

type OfpGroupCapabilities uint32

const (
	OFPGFC_SELECT_WEIGHT   OfpGroupCapabilities = 1 << 0
	OFPGFC_SELECT_LIVENESS                      = 1 << 1
	OFPGFC_CHAINING                             = 1 << 2
	OFPGFC_CHAINING_CHECKS                      = 1 << 3
)

type OfpQueueProperties uint32

const (
	OFPQT_MIN_RATE     OfpQueueProperties = 1
	OFPQT_MAX_RATE                        = 2
	OFPQT_EXPERIMENTER                    = 0xffff
)

type OfpControllerRole uint32

const (
	OFPCR_ROLE_NOCHANGE OfpControllerRole = 0
	OFPCR_ROLE_EQUAL                      = 1
	OFPCR_ROLE_MASTER                     = 2
	OFPCR_ROLE_SLAVE                      = 3
)
