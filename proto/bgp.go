package packets

// https://github.com/bodea/l2tpns/blob/master/bgp.h
const (
	BGP_MAX_PACKET_SIZE = 4096
	BGP_HOLD_TIME       = 180 /* seconds before peer times us out */
	BGP_KEEPALIVE_TIME  = 60  /* seconds between messages */
	BGP_STATE_TIME      = 60  /* state transition timeout in seconds */
	BGP_MAX_RETRY       = 42  /* maximum number of times to retry */
	BGP_RETRY_BACKOFF   = 60  /* number of seconds between retries, cumulative */

	BGP_METRIC     = 1   /* multi_exit_disc */
	BGP_LOCAL_PREF = 100 /* local preference value */

	BGP_MARKER_SIZE = 16
)

type BgpHtype uint8

const (
	BGP_MSG_OPEN         = BgpHtype(1)
	BGP_MSG_UPDATE       = 2
	BGP_MSG_NOTIFICATION = 3
	BGP_MSG_KEEPALIVE    = 4
)

type BgpHeader struct {
	// marker = 16,
	// len = 2,
	// type = 1,
}

type BgpOpen struct {
	// version 8
	// as 2
	// hold_time 2
	// identifier 4
	// opt_len 4
}

/* bgp_path_attr.flags (bitfields) */
const (
	BGP_PATH_ATTR_FLAG_OPTIONAL = (1 << 7)
	BGP_PATH_ATTR_FLAG_TRANS    = (1 << 6)
	BGP_PATH_ATTR_FLAG_PARTIAL  = (1 << 5)
	BGP_PATH_ATTR_FLAG_EXTLEN   = (1 << 4)
)

/* bgp_path_attr.code, ...value */
const (
	BGP_PATH_ATTR_CODE_ORIGIN            = 1 /* well-known, mandatory */
	BGP_PATH_ATTR_CODE_ORIGIN_IGP        = 0
	BGP_PATH_ATTR_CODE_ORIGIN_EGP        = 1
	BGP_PATH_ATTR_CODE_ORIGIN_INCOMPLETE = 2

	BGP_PATH_ATTR_CODE_AS_PATH             = 2 /* well-known, mandatory */
	BGP_PATH_ATTR_CODE_AS_PATH_AS_SET      = 1
	BGP_PATH_ATTR_CODE_AS_PATH_AS_SEQUENCE = 2

	BGP_PATH_ATTR_CODE_NEXT_HOP         = 3 /* well-known, mandatory */
	BGP_PATH_ATTR_CODE_MULTI_EXIT_DISC  = 4 /* optional, non-transitive */
	BGP_PATH_ATTR_CODE_LOCAL_PREF       = 5 /* well-known, discretionary */
	BGP_PATH_ATTR_CODE_ATOMIC_AGGREGATE = 6 /* well-known, discretionary */
	BGP_PATH_ATTR_CODE_AGGREGATOR       = 7 /* optional, transitive */
	BGP_PATH_ATTR_CODE_COMMUNITIES      = 8 /* optional, transitive (RFC1997) */
)

/* well known COMMUNITIES */
const (
	BGP_COMMUNITY_NO_EXPORT           = 0xffffff01 /* don't advertise outside confederation */
	BGP_COMMUNITY_NO_ADVERTISE        = 0xffffff02 /* don't advertise to any peer */
	BGP_COMMUNITY_NO_EXPORT_SUBCONFED = 0xffffff03 /* don't advertise to any other AS */
)

/* bgp_data_notification.error_code, .error_subcode */
const (
	BGP_ERR_HEADER       = 1
	BGP_ERR_HDR_NOT_SYNC = 1
	BGP_ERR_HDR_BAD_LEN  = 2
	BGP_ERR_HDR_BAD_TYPE = 3

	BGP_ERR_OPEN             = 2
	BGP_ERR_OPN_VERSION      = 1
	BGP_ERR_OPN_BAD_AS       = 2
	BGP_ERR_OPN_BAD_IDENT    = 3
	BGP_ERR_OPN_UNSUP_PARAM  = 4
	BGP_ERR_OPN_AUTH_FAILURE = 5
	BGP_ERR_OPN_HOLD_TIME    = 6

	BGP_ERR_UPDATE            = 3
	BGP_ERR_UPD_BAD_ATTR_LIST = 1
	BGP_ERR_UPD_UNKN_WK_ATTR  = 2
	BGP_ERR_UPD_MISS_WK_ATTR  = 3
	BGP_ERR_UPD_BAD_ATTR_FLAG = 4
	BGP_ERR_UPD_BAD_ATTR_LEN  = 5
	BGP_ERR_UPD_BAD_ORIGIN    = 6
	BGP_ERR_UPD_ROUTING_LOOP  = 7
	BGP_ERR_UPD_BAD_NEXT_HOP  = 8
	BGP_ERR_UPD_BAD_OPT_ATTR  = 9
	BGP_ERR_UPD_BAD_NETWORK   = 10
	BGP_ERR_UPD_BAD_AS_PATH   = 11

	BGP_ERR_HOLD_TIMER_EXP = 4
	BGP_ERR_FSM            = 5
	BGP_ERR_CEASE          = 6
)

const (
	/* bgp_peer.cli_flag */
	BGP_CLI_SUSPEND = 1
	BGP_CLI_ENABLE  = 2
	BGP_CLI_RESTART = 3
)

// states
const (
	Disabled    /* initial, or failed */
	Idle        /* trying to connect */
	Connect     /* connect issued */
	Active      /* connected, waiting to send OPEN */
	OpenSent    /* OPEN sent, waiting for peer OPEN */
	OpenConfirm /* KEEPALIVE sent, waiting for peer KEEPALIVE */
	Established /* established */
)
