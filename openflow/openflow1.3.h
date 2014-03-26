













       

       















#pragma pack(push,_CRT_PACKING)

#pragma pack(pop)

#define __extension__ 

#pragma pack(push,_CRT_PACKING)

__extension__ typedef unsigned long long size_t;

__extension__ typedef long long ssize_t;

__extension__ typedef long long intptr_t;

__extension__ typedef unsigned long long uintptr_t;

__extension__ typedef long long ptrdiff_t;

typedef unsigned short wchar_t;







typedef unsigned short wint_t;
typedef unsigned short wctype_t;

typedef int errno_t;




typedef long __time32_t;




__extension__ typedef long long __time64_t;







typedef __time64_t time_t;

#define __attribute__(A)
#define __inline__
#define __asm__




const char *__mingw_get_crt_info (void);






#pragma pack(pop)


typedef size_t rsize_t;

struct threadlocaleinfostruct;
struct threadmbcinfostruct;
typedef struct threadlocaleinfostruct *pthreadlocinfo;
typedef struct threadmbcinfostruct *pthreadmbcinfo;
struct __lc_time_data;

typedef struct localeinfo_struct {
  pthreadlocinfo locinfo;
  pthreadmbcinfo mbcinfo;
} _locale_tstruct,*_locale_t;



typedef struct tagLC_ID {
  unsigned short wLanguage;
  unsigned short wCountry;
  unsigned short wCodePage;
} LC_ID,*LPLC_ID;




typedef struct threadlocaleinfostruct {
  int refcount;
  unsigned int lc_codepage;
  unsigned int lc_collate_cp;
  unsigned long lc_handle[6];
  LC_ID lc_id[6];
  struct {
    char *locale;
    wchar_t *wlocale;
    int *refcount;
    int *wrefcount;
  } lc_category[6];
  int lc_clike;
  int mb_cur_max;
  int *lconv_intl_refcount;
  int *lconv_num_refcount;
  int *lconv_mon_refcount;
  struct lconv *lconv;
  int *ctype1_refcount;
  unsigned short *ctype1;
  const unsigned short *pctype;
  const unsigned char *pclmap;
  const unsigned char *pcumap;
  struct __lc_time_data *lc_time_curr;
} threadlocinfo;







  __attribute__ ((__dllimport__)) extern int *__attribute__((__cdecl__)) _errno(void);

  errno_t __attribute__((__cdecl__)) _set_errno(int _Value);
  errno_t __attribute__((__cdecl__)) _get_errno(int *_Value);


  __attribute__ ((__dllimport__)) extern unsigned long __attribute__((__cdecl__)) __threadid(void);

  __attribute__ ((__dllimport__)) extern uintptr_t __attribute__((__cdecl__)) __threadhandle(void);




typedef signed char int8_t;
typedef unsigned char uint8_t;
typedef short int16_t;
typedef unsigned short uint16_t;
typedef int int32_t;
typedef unsigned uint32_t;
__extension__ typedef long long int64_t;
__extension__ typedef unsigned long long uint64_t;


typedef signed char int_least8_t;
typedef unsigned char uint_least8_t;
typedef short int_least16_t;
typedef unsigned short uint_least16_t;
typedef int int_least32_t;
typedef unsigned uint_least32_t;
__extension__ typedef long long int_least64_t;
__extension__ typedef unsigned long long uint_least64_t;





typedef signed char int_fast8_t;
typedef unsigned char uint_fast8_t;
typedef short int_fast16_t;
typedef unsigned short uint_fast16_t;
typedef int int_fast32_t;
typedef unsigned int uint_fast32_t;
__extension__ typedef long long int_fast64_t;
__extension__ typedef unsigned long long uint_fast64_t;


__extension__ typedef long long intmax_t;
__extension__ typedef unsigned long long uintmax_t;



enum ofp_port_no {

    OFPP_MAX = 0xffffff00,


    OFPP_IN_PORT = 0xfffffff8,



    OFPP_TABLE = 0xfffffff9,


    OFPP_NORMAL = 0xfffffffa,
    OFPP_FLOOD = 0xfffffffb,

    OFPP_ALL = 0xfffffffc,
    OFPP_CONTROLLER = 0xfffffffd,
    OFPP_LOCAL = 0xfffffffe,
    OFPP_ANY = 0xffffffff



};

enum ofp_type {

    OFPT_HELLO = 0,
    OFPT_ERROR = 1,
    OFPT_ECHO_REQUEST = 2,
    OFPT_ECHO_REPLY = 3,
    OFPT_EXPERIMENTER = 4,


    OFPT_FEATURES_REQUEST = 5,
    OFPT_FEATURES_REPLY = 6,
    OFPT_GET_CONFIG_REQUEST = 7,
    OFPT_GET_CONFIG_REPLY = 8,
    OFPT_SET_CONFIG = 9,


    OFPT_PACKET_IN = 10,
    OFPT_FLOW_REMOVED = 11,
    OFPT_PORT_STATUS = 12,


    OFPT_PACKET_OUT = 13,
    OFPT_FLOW_MOD = 14,
    OFPT_GROUP_MOD = 15,
    OFPT_PORT_MOD = 16,
    OFPT_TABLE_MOD = 17,


    OFPT_MULTIPART_REQUEST = 18,
    OFPT_MULTIPART_REPLY = 19,


    OFPT_BARRIER_REQUEST = 20,
    OFPT_BARRIER_REPLY = 21,


    OFPT_QUEUE_GET_CONFIG_REQUEST = 22,
    OFPT_QUEUE_GET_CONFIG_REPLY = 23,


    OFPT_ROLE_REQUEST = 24,
    OFPT_ROLE_REPLY = 25,


    OFPT_GET_ASYNC_REQUEST = 26,
    OFPT_GET_ASYNC_REPLY = 27,
    OFPT_SET_ASYNC = 28,


    OFPT_METER_MOD = 29,
};


struct ofp_header {
    uint8_t version;
    uint8_t type;
    uint16_t length;
    uint32_t xid;


};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_header) == 8) ? 1 : -1; })];



enum ofp_hello_elem_type {
    OFPHET_VERSIONBITMAP = 1,
};


struct ofp_hello_elem_header {
    uint16_t type;
    uint16_t length;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_hello_elem_header) == 4) ? 1 : -1; })];


struct ofp_hello_elem_versionbitmap {
    uint16_t type;
    uint16_t length;




    uint32_t bitmaps[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_hello_elem_versionbitmap) == 4) ? 1 : -1; })];




struct ofp_hello {
    struct ofp_header header;


    struct ofp_hello_elem_header elements[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_hello) == 8) ? 1 : -1; })];



enum ofp_config_flags {

    OFPC_FRAG_NORMAL = 0,
    OFPC_FRAG_DROP = 1 << 0,
    OFPC_FRAG_REASM = 1 << 1,
    OFPC_FRAG_MASK = 3,
};


struct ofp_switch_config {
    struct ofp_header header;
    uint16_t flags;
    uint16_t miss_send_len;



};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_switch_config) == 12) ? 1 : -1; })];


enum ofp_table_config {
    OFPTC_DEPRECATED_MASK = 3,
};


enum ofp_table {

    OFPTT_MAX = 0xfe,


    OFPTT_ALL = 0xff

};



struct ofp_table_mod {
    struct ofp_header header;
    uint8_t table_id;
    uint8_t pad[3];
    uint32_t config;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_mod) == 16) ? 1 : -1; })];


enum ofp_capabilities {
    OFPC_FLOW_STATS = 1 << 0,
    OFPC_TABLE_STATS = 1 << 1,
    OFPC_PORT_STATS = 1 << 2,
    OFPC_GROUP_STATS = 1 << 3,
    OFPC_IP_REASM = 1 << 5,
    OFPC_QUEUE_STATS = 1 << 6,
    OFPC_PORT_BLOCKED = 1 << 8
};





enum ofp_port_config {
    OFPPC_PORT_DOWN = 1 << 0,

    OFPPC_NO_RECV = 1 << 2,
    OFPPC_NO_FWD = 1 << 5,
    OFPPC_NO_PACKET_IN = 1 << 6
};




enum ofp_port_state {
    OFPPS_LINK_DOWN = 1 << 0,
    OFPPS_BLOCKED = 1 << 1,
    OFPPS_LIVE = 1 << 2,
};


enum ofp_port_features {
    OFPPF_10MB_HD = 1 << 0,
    OFPPF_10MB_FD = 1 << 1,
    OFPPF_100MB_HD = 1 << 2,
    OFPPF_100MB_FD = 1 << 3,
    OFPPF_1GB_HD = 1 << 4,
    OFPPF_1GB_FD = 1 << 5,
    OFPPF_10GB_FD = 1 << 6,
    OFPPF_40GB_FD = 1 << 7,
    OFPPF_100GB_FD = 1 << 8,
    OFPPF_1TB_FD = 1 << 9,
    OFPPF_OTHER = 1 << 10,

    OFPPF_COPPER = 1 << 11,
    OFPPF_FIBER = 1 << 12,
    OFPPF_AUTONEG = 1 << 13,
    OFPPF_PAUSE = 1 << 14,
    OFPPF_PAUSE_ASYM = 1 << 15
};


struct ofp_port {
    uint32_t port_no;
    uint8_t pad[4];
    uint8_t hw_addr[6];
    uint8_t pad2[2];
    char name[16];

    uint32_t config;
    uint32_t state;



    uint32_t curr;
    uint32_t advertised;
    uint32_t supported;
    uint32_t peer;

    uint32_t curr_speed;
    uint32_t max_speed;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_port) == 64) ? 1 : -1; })];


struct ofp_switch_features {
    struct ofp_header header;
    uint64_t datapath_id;



    uint32_t n_buffers;

    uint8_t n_tables;
    uint8_t auxiliary_id;
    uint8_t pad[2];


    uint32_t capabilities;
    uint32_t reserved;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_switch_features) == 32) ? 1 : -1; })];


enum ofp_port_reason {
    OFPPR_ADD = 0,
    OFPPR_DELETE = 1,
    OFPPR_MODIFY = 2,
};


struct ofp_port_status {
    struct ofp_header header;
    uint8_t reason;
    uint8_t pad[7];
    struct ofp_port desc;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_port_status) == 80) ? 1 : -1; })];


struct ofp_port_mod {
    struct ofp_header header;
    uint32_t port_no;
    uint8_t pad[4];
    uint8_t hw_addr[6];




    uint8_t pad2[2];
    uint32_t config;
    uint32_t mask;

    uint32_t advertise;

    uint8_t pad3[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_port_mod) == 40) ? 1 : -1; })];

enum ofp_match_type {
    OFPMT_STANDARD = 0,
    OFPMT_OXM = 1,
};


struct ofp_match {
    uint16_t type;
    uint16_t length;







    uint8_t oxm_fields[0];
    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_match) == 8) ? 1 : -1; })];

enum ofp_oxm_class {
    OFPXMC_NXM_0 = 0x0000,
    OFPXMC_NXM_1 = 0x0001,
    OFPXMC_OPENFLOW_BASIC = 0x8000,
    OFPXMC_EXPERIMENTER = 0xFFFF,
};


enum oxm_ofb_match_fields {
    OFPXMT_OFB_IN_PORT = 0,
    OFPXMT_OFB_IN_PHY_PORT = 1,
    OFPXMT_OFB_METADATA = 2,
    OFPXMT_OFB_ETH_DST = 3,
    OFPXMT_OFB_ETH_SRC = 4,
    OFPXMT_OFB_ETH_TYPE = 5,
    OFPXMT_OFB_VLAN_VID = 6,
    OFPXMT_OFB_VLAN_PCP = 7,
    OFPXMT_OFB_IP_DSCP = 8,
    OFPXMT_OFB_IP_ECN = 9,
    OFPXMT_OFB_IP_PROTO = 10,
    OFPXMT_OFB_IPV4_SRC = 11,
    OFPXMT_OFB_IPV4_DST = 12,
    OFPXMT_OFB_TCP_SRC = 13,
    OFPXMT_OFB_TCP_DST = 14,
    OFPXMT_OFB_UDP_SRC = 15,
    OFPXMT_OFB_UDP_DST = 16,
    OFPXMT_OFB_SCTP_SRC = 17,
    OFPXMT_OFB_SCTP_DST = 18,
    OFPXMT_OFB_ICMPV4_TYPE = 19,
    OFPXMT_OFB_ICMPV4_CODE = 20,
    OFPXMT_OFB_ARP_OP = 21,
    OFPXMT_OFB_ARP_SPA = 22,
    OFPXMT_OFB_ARP_TPA = 23,
    OFPXMT_OFB_ARP_SHA = 24,
    OFPXMT_OFB_ARP_THA = 25,
    OFPXMT_OFB_IPV6_SRC = 26,
    OFPXMT_OFB_IPV6_DST = 27,
    OFPXMT_OFB_IPV6_FLABEL = 28,
    OFPXMT_OFB_ICMPV6_TYPE = 29,
    OFPXMT_OFB_ICMPV6_CODE = 30,
    OFPXMT_OFB_IPV6_ND_TARGET = 31,
    OFPXMT_OFB_IPV6_ND_SLL = 32,
    OFPXMT_OFB_IPV6_ND_TLL = 33,
    OFPXMT_OFB_MPLS_LABEL = 34,
    OFPXMT_OFB_MPLS_TC = 35,
    OFPXMT_OFP_MPLS_BOS = 36,
    OFPXMT_OFB_PBB_ISID = 37,
    OFPXMT_OFB_TUNNEL_ID = 38,
    OFPXMT_OFB_IPV6_EXTHDR = 39,
};

enum ofp_vlan_id {
    OFPVID_PRESENT = 0x1000,
    OFPVID_NONE = 0x0000,
};

enum ofp_ipv6exthdr_flags {
    OFPIEH_NONEXT = 1 << 0,
    OFPIEH_ESP = 1 << 1,
    OFPIEH_AUTH = 1 << 2,
    OFPIEH_DEST = 1 << 3,
    OFPIEH_FRAG = 1 << 4,
    OFPIEH_ROUTER = 1 << 5,
    OFPIEH_HOP = 1 << 6,
    OFPIEH_UNREP = 1 << 7,
    OFPIEH_UNSEQ = 1 << 8,
};


struct ofp_oxm_experimenter_header {
    uint32_t oxm_header;
    uint32_t experimenter;

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_oxm_experimenter_header) == 8) ? 1 : -1; })];





enum ofp_action_type {
    OFPAT_OUTPUT = 0,
    OFPAT_COPY_TTL_OUT = 11,

    OFPAT_COPY_TTL_IN = 12,

    OFPAT_SET_MPLS_TTL = 15,
    OFPAT_DEC_MPLS_TTL = 16,

    OFPAT_PUSH_VLAN = 17,
    OFPAT_POP_VLAN = 18,
    OFPAT_PUSH_MPLS = 19,
    OFPAT_POP_MPLS = 20,
    OFPAT_SET_QUEUE = 21,
    OFPAT_GROUP = 22,
    OFPAT_SET_NW_TTL = 23,
    OFPAT_DEC_NW_TTL = 24,
    OFPAT_SET_FIELD = 25,
    OFPAT_PUSH_PBB = 26,
    OFPAT_POP_PBB = 27,
    OFPAT_EXPERIMENTER = 0xffff
};




struct ofp_action_header {
    uint16_t type;
    uint16_t len;



    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_header) == 8) ? 1 : -1; })];

enum ofp_controller_max_len {
 OFPCML_MAX = 0xffe5,

 OFPCML_NO_BUFFER = 0xffff


};







struct ofp_action_output {
    uint16_t type;
    uint16_t len;
    uint32_t port;
    uint16_t max_len;
    uint8_t pad[6];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_output) == 16) ? 1 : -1; })];


struct ofp_action_mpls_ttl {
    uint16_t type;
    uint16_t len;
    uint8_t mpls_ttl;
    uint8_t pad[3];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_mpls_ttl) == 8) ? 1 : -1; })];


struct ofp_action_push {
    uint16_t type;
    uint16_t len;
    uint16_t ethertype;
    uint8_t pad[2];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_push) == 8) ? 1 : -1; })];


struct ofp_action_pop_mpls {
    uint16_t type;
    uint16_t len;
    uint16_t ethertype;
    uint8_t pad[2];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_pop_mpls) == 8) ? 1 : -1; })];


struct ofp_action_group {
    uint16_t type;
    uint16_t len;
    uint32_t group_id;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_group) == 8) ? 1 : -1; })];


struct ofp_action_nw_ttl {
    uint16_t type;
    uint16_t len;
    uint8_t nw_ttl;
    uint8_t pad[3];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_nw_ttl) == 8) ? 1 : -1; })];


struct ofp_action_set_field {
    uint16_t type;
    uint16_t len;





    uint8_t field[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_set_field) == 8) ? 1 : -1; })];



struct ofp_action_experimenter_header {
    uint16_t type;
    uint16_t len;
    uint32_t experimenter;


};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_experimenter_header) == 8) ? 1 : -1; })];





enum ofp_instruction_type {
    OFPIT_GOTO_TABLE = 1,

    OFPIT_WRITE_METADATA = 2,

    OFPIT_WRITE_ACTIONS = 3,

    OFPIT_APPLY_ACTIONS = 4,
    OFPIT_CLEAR_ACTIONS = 5,

    OFPIT_METER = 6,

    OFPIT_EXPERIMENTER = 0xFFFF
};




struct ofp_instruction {
    uint16_t type;
    uint16_t len;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_instruction) == 4) ? 1 : -1; })];


struct ofp_instruction_goto_table {
    uint16_t type;
    uint16_t len;
    uint8_t table_id;
    uint8_t pad[3];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_instruction_goto_table) == 8) ? 1 : -1; })];


struct ofp_instruction_write_metadata {
    uint16_t type;
    uint16_t len;
    uint8_t pad[4];
    uint64_t metadata;
    uint64_t metadata_mask;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_instruction_write_metadata) == 24) ? 1 : -1; })];


struct ofp_instruction_actions {
    uint16_t type;
    uint16_t len;
    uint8_t pad[4];
    struct ofp_action_header actions[0];


};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_instruction_actions) == 8) ? 1 : -1; })];


struct ofp_instruction_meter {
    uint16_t type;
    uint16_t len;
    uint32_t meter_id;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_instruction_meter) == 8) ? 1 : -1; })];


struct ofp_instruction_experimenter {
    uint16_t type;
    uint16_t len;
    uint32_t experimenter;


};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_instruction_experimenter) == 8) ? 1 : -1; })];





enum ofp_flow_mod_command {
    OFPFC_ADD = 0,
    OFPFC_MODIFY = 1,
    OFPFC_MODIFY_STRICT = 2,

    OFPFC_DELETE = 3,
    OFPFC_DELETE_STRICT = 4,

};

enum ofp_flow_mod_flags {
    OFPFF_SEND_FLOW_REM = 1 << 0,

    OFPFF_CHECK_OVERLAP = 1 << 1,
    OFPFF_RESET_COUNTS = 1 << 2,
    OFPFF_NO_PKT_COUNTS = 1 << 3,
    OFPFF_NO_BYT_COUNTS = 1 << 4,
};


struct ofp_flow_mod {
    struct ofp_header header;
    uint64_t cookie;
    uint64_t cookie_mask;





    uint8_t table_id;



    uint8_t command;
    uint16_t idle_timeout;
    uint16_t hard_timeout;
    uint16_t priority;
    uint32_t buffer_id;


    uint32_t out_port;



    uint32_t out_group;



    uint16_t flags;
    uint8_t pad[2];
    struct ofp_match match;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_flow_mod) == 56) ? 1 : -1; })];


enum ofp_group {

    OFPG_MAX = 0xffffff00,


    OFPG_ALL = 0xfffffffc,

    OFPG_ANY = 0xffffffff



};


enum ofp_group_mod_command {
    OFPGC_ADD = 0,
    OFPGC_MODIFY = 1,
    OFPGC_DELETE = 2,
};


struct ofp_bucket {
    uint16_t len;


    uint16_t weight;

    uint32_t watch_port;


    uint32_t watch_group;


    uint8_t pad[4];
    struct ofp_action_header actions[0];



};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_bucket) == 16) ? 1 : -1; })];


struct ofp_group_mod {
    struct ofp_header header;
    uint16_t command;
    uint8_t type;
    uint8_t pad;
    uint32_t group_id;
    struct ofp_bucket buckets[0];

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_group_mod) == 16) ? 1 : -1; })];



enum ofp_group_type {
    OFPGT_ALL = 0,
    OFPGT_SELECT = 1,
    OFPGT_INDIRECT = 2,
    OFPGT_FF = 3,
};





struct ofp_packet_out {
    struct ofp_header header;
    uint32_t buffer_id;

    uint32_t in_port;
    uint16_t actions_len;
    uint8_t pad[6];
    struct ofp_action_header actions[0];




};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_packet_out) == 24) ? 1 : -1; })];


enum ofp_packet_in_reason {
    OFPR_NO_MATCH = 0,
    OFPR_ACTION = 1,
    OFPR_INVALID_TTL = 2,
};


struct ofp_packet_in {
    struct ofp_header header;
    uint32_t buffer_id;
    uint16_t total_len;
    uint8_t reason;
    uint8_t table_id;
    uint64_t cookie;
    struct ofp_match match;

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_packet_in) == 32) ? 1 : -1; })];


enum ofp_flow_removed_reason {
    OFPRR_IDLE_TIMEOUT = 0,
    OFPRR_HARD_TIMEOUT = 1,
    OFPRR_DELETE = 2,
    OFPRR_GROUP_DELETE = 3,
};


struct ofp_flow_removed {
    struct ofp_header header;
    uint64_t cookie;

    uint16_t priority;
    uint8_t reason;
    uint8_t table_id;

    uint32_t duration_sec;
    uint32_t duration_nsec;

    uint16_t idle_timeout;
    uint16_t hard_timeout;
    uint64_t packet_count;
    uint64_t byte_count;
    struct ofp_match match;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_flow_removed) == 56) ? 1 : -1; })];


enum ofp_meter {

    OFPM_MAX = 0xffff0000,


    OFPM_SLOWPATH = 0xfffffffd,
    OFPM_CONTROLLER = 0xfffffffe,
    OFPM_ALL = 0xffffffff,

};


enum ofp_meter_band_type {
    OFPMBT_DROP = 1,
    OFPMBT_DSCP_REMARK = 2,
    OFPMBT_EXPERIMENTER = 0xFFFF
};


struct ofp_meter_band_header {
    uint16_t type;
    uint16_t len;
    uint32_t rate;
    uint32_t burst_size;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_band_header) == 12) ? 1 : -1; })];


struct ofp_meter_band_drop {
    uint16_t type;
    uint16_t len;
    uint32_t rate;
    uint32_t burst_size;
    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_band_drop) == 16) ? 1 : -1; })];


struct ofp_meter_band_dscp_remark {
    uint16_t type;
    uint16_t len;
    uint32_t rate;
    uint32_t burst_size;
    uint8_t prec_level;
    uint8_t pad[3];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_band_dscp_remark) == 16) ? 1 : -1; })];


struct ofp_meter_band_experimenter {
    uint16_t type;
    uint16_t len;
    uint32_t rate;
    uint32_t burst_size;
    uint32_t experimenter;


};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_band_experimenter) == 16) ? 1 : -1; })];


enum ofp_meter_mod_command {
    OFPMC_ADD,
    OFPMC_MODIFY,
    OFPMC_DELETE,
};


enum ofp_meter_flags {
    OFPMF_KBPS = 1 << 0,
    OFPMF_PKTPS = 1 << 1,
    OFPMF_BURST = 1 << 2,
    OFPMF_STATS = 1 << 3,
};


struct ofp_meter_mod {
    struct ofp_header header;
    uint16_t command;
    uint16_t flags;
    uint32_t meter_id;
    struct ofp_meter_band_header bands[0];


};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_mod) == 16) ? 1 : -1; })];




enum ofp_error_type {
    OFPET_HELLO_FAILED = 0,
    OFPET_BAD_REQUEST = 1,
    OFPET_BAD_ACTION = 2,
    OFPET_BAD_INSTRUCTION = 3,
    OFPET_BAD_MATCH = 4,
    OFPET_FLOW_MOD_FAILED = 5,
    OFPET_GROUP_MOD_FAILED = 6,
    OFPET_PORT_MOD_FAILED = 7,
    OFPET_TABLE_MOD_FAILED = 8,
    OFPET_QUEUE_OP_FAILED = 9,
    OFPET_SWITCH_CONFIG_FAILED = 10,
    OFPET_ROLE_REQUEST_FAILED = 11,
    OFPET_METER_MOD_FAILED = 12,
    OFPET_TABLE_FEATURES_FAILED = 13,
    OFPET_EXPERIMENTER = 0xffff
};



enum ofp_hello_failed_code {
    OFPHFC_INCOMPATIBLE = 0,
    OFPHFC_EPERM = 1,
};



enum ofp_bad_request_code {
    OFPBRC_BAD_VERSION = 0,
    OFPBRC_BAD_TYPE = 1,
    OFPBRC_BAD_MULTIPART = 2,
    OFPBRC_BAD_EXPERIMENTER = 3,



    OFPBRC_BAD_EXP_TYPE = 4,
    OFPBRC_EPERM = 5,
    OFPBRC_BAD_LEN = 6,
    OFPBRC_BUFFER_EMPTY = 7,
    OFPBRC_BUFFER_UNKNOWN = 8,
    OFPBRC_BAD_TABLE_ID = 9,

    OFPBRC_IS_SLAVE = 10,
    OFPBRC_BAD_PORT = 11,
    OFPBRC_BAD_PACKET = 12,
    OFPBRC_MULTIPART_BUFFER_OVERFLOW = 13,

};



enum ofp_bad_action_code {
    OFPBAC_BAD_TYPE = 0,
    OFPBAC_BAD_LEN = 1,
    OFPBAC_BAD_EXPERIMENTER = 2,
    OFPBAC_BAD_EXP_TYPE = 3,
    OFPBAC_BAD_OUT_PORT = 4,
    OFPBAC_BAD_ARGUMENT = 5,
    OFPBAC_EPERM = 6,
    OFPBAC_TOO_MANY = 7,
    OFPBAC_BAD_QUEUE = 8,
    OFPBAC_BAD_OUT_GROUP = 9,
    OFPBAC_MATCH_INCONSISTENT = 10,

    OFPBAC_UNSUPPORTED_ORDER = 11,

    OFPBAC_BAD_TAG = 12,

    OFPBAC_BAD_SET_TYPE = 13,
    OFPBAC_BAD_SET_LEN = 14,
    OFPBAC_BAD_SET_ARGUMENT = 15,
};



enum ofp_bad_instruction_code {
    OFPBIC_UNKNOWN_INST = 0,
    OFPBIC_UNSUP_INST = 1,

    OFPBIC_BAD_TABLE_ID = 2,
    OFPBIC_UNSUP_METADATA = 3,
    OFPBIC_UNSUP_METADATA_MASK = 4,

    OFPBIC_BAD_EXPERIMENTER = 5,
    OFPBIC_BAD_EXP_TYPE = 6,
    OFPBIC_BAD_LEN = 7,
    OFPBIC_EPERM = 8,
};



enum ofp_bad_match_code {
    OFPBMC_BAD_TYPE = 0,

    OFPBMC_BAD_LEN = 1,
    OFPBMC_BAD_TAG = 2,
    OFPBMC_BAD_DL_ADDR_MASK = 3,


    OFPBMC_BAD_NW_ADDR_MASK = 4,


    OFPBMC_BAD_WILDCARDS = 5,

    OFPBMC_BAD_FIELD = 6,
    OFPBMC_BAD_VALUE = 7,
    OFPBMC_BAD_MASK = 8,

    OFPBMC_BAD_PREREQ = 9,
    OFPBMC_DUP_FIELD = 10,
    OFPBMC_EPERM = 11,
};



enum ofp_flow_mod_failed_code {
    OFPFMFC_UNKNOWN = 0,
    OFPFMFC_TABLE_FULL = 1,
    OFPFMFC_BAD_TABLE_ID = 2,
    OFPFMFC_OVERLAP = 3,

    OFPFMFC_EPERM = 4,
    OFPFMFC_BAD_TIMEOUT = 5,

    OFPFMFC_BAD_COMMAND = 6,
    OFPFMFC_BAD_FLAGS = 7,
};



enum ofp_group_mod_failed_code {
    OFPGMFC_GROUP_EXISTS = 0,


    OFPGMFC_INVALID_GROUP = 1,

    OFPGMFC_WEIGHT_UNSUPPORTED = 2,

    OFPGMFC_OUT_OF_GROUPS = 3,
    OFPGMFC_OUT_OF_BUCKETS = 4,

    OFPGMFC_CHAINING_UNSUPPORTED = 5,

    OFPGMFC_WATCH_UNSUPPORTED = 6,

    OFPGMFC_LOOP = 7,
    OFPGMFC_UNKNOWN_GROUP = 8,


    OFPGMFC_CHAINED_GROUP = 9,

    OFPGMFC_BAD_TYPE = 10,
    OFPGMFC_BAD_COMMAND = 11,
    OFPGMFC_BAD_BUCKET = 12,
    OFPGMFC_BAD_WATCH = 13,
    OFPGMFC_EPERM = 14,
};



enum ofp_port_mod_failed_code {
    OFPPMFC_BAD_PORT = 0,
    OFPPMFC_BAD_HW_ADDR = 1,

    OFPPMFC_BAD_CONFIG = 2,
    OFPPMFC_BAD_ADVERTISE = 3,
    OFPPMFC_EPERM = 4,
};



enum ofp_table_mod_failed_code {
    OFPTMFC_BAD_TABLE = 0,
    OFPTMFC_BAD_CONFIG = 1,
    OFPTMFC_EPERM = 2,
};



enum ofp_queue_op_failed_code {
    OFPQOFC_BAD_PORT = 0,
    OFPQOFC_BAD_QUEUE = 1,
    OFPQOFC_EPERM = 2,
};



enum ofp_switch_config_failed_code {
    OFPSCFC_BAD_FLAGS = 0,
    OFPSCFC_BAD_LEN = 1,
    OFPSCFC_EPERM = 2,
};



enum ofp_role_request_failed_code {
    OFPRRFC_STALE = 0,
    OFPRRFC_UNSUP = 1,
    OFPRRFC_BAD_ROLE = 2,
};



enum ofp_meter_mod_failed_code {
    OFPMMFC_UNKNOWN = 0,
    OFPMMFC_METER_EXISTS = 1,

    OFPMMFC_INVALID_METER = 2,

    OFPMMFC_UNKNOWN_METER = 3,


    OFPMMFC_BAD_COMMAND = 4,
    OFPMMFC_BAD_FLAGS = 5,
    OFPMMFC_BAD_RATE = 6,
    OFPMMFC_BAD_BURST = 7,
    OFPMMFC_BAD_BAND = 8,
    OFPMMFC_BAD_BAND_VALUE = 9,
    OFPMMFC_OUT_OF_METERS = 10,
    OFPMMFC_OUT_OF_BANDS = 11,

};



enum ofp_table_features_failed_code {
    OFPTFFC_BAD_TABLE = 0,
    OFPTFFC_BAD_METADATA = 1,
    OFPTFFC_BAD_TYPE = 2,
    OFPTFFC_BAD_LEN = 3,
    OFPTFFC_BAD_ARGUMENT = 4,
    OFPTFFC_EPERM = 5,
};


struct ofp_error_msg {
    struct ofp_header header;

    uint16_t type;
    uint16_t code;
    uint8_t data[0];

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_error_msg) == 12) ? 1 : -1; })];


struct ofp_error_experimenter_msg {
    struct ofp_header header;

    uint16_t type;
    uint16_t exp_type;
    uint32_t experimenter;

    uint8_t data[0];

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_error_experimenter_msg) == 16) ? 1 : -1; })];

enum ofp_multipart_type {



    OFPMP_DESC = 0,




    OFPMP_FLOW = 1,




    OFPMP_AGGREGATE = 2,




    OFPMP_TABLE = 3,




    OFPMP_PORT_STATS = 4,




    OFPMP_QUEUE = 5,




    OFPMP_GROUP = 6,




    OFPMP_GROUP_DESC = 7,




    OFPMP_GROUP_FEATURES = 8,




    OFPMP_METER = 9,




    OFPMP_METER_CONFIG = 10,




    OFPMP_METER_FEATURES = 11,







    OFPMP_TABLE_FEATURES = 12,




    OFPMP_PORT_DESC = 13,





    OFPMP_EXPERIMENTER = 0xffff
};




enum ofp_multipart_request_flags {
    OFPMPF_REQ_MORE = 1 << 0
};

struct ofp_multipart_request {
    struct ofp_header header;
    uint16_t type;
    uint16_t flags;
    uint8_t pad[4];
    uint8_t body[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_multipart_request) == 16) ? 1 : -1; })];

enum ofp_multipart_reply_flags {
    OFPMPF_REPLY_MORE = 1 << 0
};

struct ofp_multipart_reply {
    struct ofp_header header;
    uint16_t type;
    uint16_t flags;
    uint8_t pad[4];
    uint8_t body[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_multipart_reply) == 16) ? 1 : -1; })];





struct ofp_desc {
    char mfr_desc[256];
    char hw_desc[256];
    char sw_desc[256];
    char serial_num[32];
    char dp_desc[256];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_desc) == 1056) ? 1 : -1; })];


struct ofp_flow_stats_request {
    uint8_t table_id;

    uint8_t pad[3];
    uint32_t out_port;


    uint32_t out_group;


    uint8_t pad2[4];
    uint64_t cookie;

    uint64_t cookie_mask;


    struct ofp_match match;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_flow_stats_request) == 40) ? 1 : -1; })];


struct ofp_flow_stats {
    uint16_t length;
    uint8_t table_id;
    uint8_t pad;
    uint32_t duration_sec;
    uint32_t duration_nsec;

    uint16_t priority;
    uint16_t idle_timeout;
    uint16_t hard_timeout;
    uint16_t flags;
    uint8_t pad2[4];
    uint64_t cookie;
    uint64_t packet_count;
    uint64_t byte_count;
    struct ofp_match match;


};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_flow_stats) == 56) ? 1 : -1; })];


struct ofp_aggregate_stats_request {
    uint8_t table_id;

    uint8_t pad[3];
    uint32_t out_port;


    uint32_t out_group;


    uint8_t pad2[4];
    uint64_t cookie;

    uint64_t cookie_mask;


    struct ofp_match match;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_aggregate_stats_request) == 40) ? 1 : -1; })];


struct ofp_aggregate_stats_reply {
    uint64_t packet_count;
    uint64_t byte_count;
    uint32_t flow_count;
    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_aggregate_stats_reply) == 24) ? 1 : -1; })];





enum ofp_table_feature_prop_type {
    OFPTFPT_INSTRUCTIONS = 0,
    OFPTFPT_INSTRUCTIONS_MISS = 1,
    OFPTFPT_NEXT_TABLES = 2,
    OFPTFPT_NEXT_TABLES_MISS = 3,
    OFPTFPT_WRITE_ACTIONS = 4,
    OFPTFPT_WRITE_ACTIONS_MISS = 5,
    OFPTFPT_APPLY_ACTIONS = 6,
    OFPTFPT_APPLY_ACTIONS_MISS = 7,
    OFPTFPT_MATCH = 8,
    OFPTFPT_WILDCARDS = 10,
    OFPTFPT_WRITE_SETFIELD = 12,
    OFPTFPT_WRITE_SETFIELD_MISS = 13,
    OFPTFPT_APPLY_SETFIELD = 14,
    OFPTFPT_APPLY_SETFIELD_MISS = 15,
    OFPTFPT_EXPERIMENTER = 0xFFFE,
    OFPTFPT_EXPERIMENTER_MISS = 0xFFFF,
};


struct ofp_table_feature_prop_header {
    uint16_t type;
    uint16_t length;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_feature_prop_header) == 4) ? 1 : -1; })];


struct ofp_table_feature_prop_instructions {
    uint16_t type;

    uint16_t length;




    struct ofp_instruction instruction_ids[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_feature_prop_instructions) == 4) ? 1 : -1; })];


struct ofp_table_feature_prop_next_tables {
    uint16_t type;

    uint16_t length;




    uint8_t next_table_ids[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_feature_prop_next_tables) == 4) ? 1 : -1; })];


struct ofp_table_feature_prop_actions {
    uint16_t type;



    uint16_t length;




    struct ofp_action_header action_ids[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_feature_prop_actions) == 4) ? 1 : -1; })];


struct ofp_table_feature_prop_oxm {
    uint16_t type;





    uint16_t length;




    uint32_t oxm_ids[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_feature_prop_oxm) == 4) ? 1 : -1; })];


struct ofp_table_feature_prop_experimenter {
    uint16_t type;

    uint16_t length;
    uint32_t experimenter;


    uint32_t exp_type;




    uint32_t experimenter_data[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_feature_prop_experimenter) == 12) ? 1 : -1; })];



struct ofp_table_features {
    uint16_t length;
    uint8_t table_id;

    uint8_t pad[5];
    char name[32];
    uint64_t metadata_match;
    uint64_t metadata_write;
    uint32_t config;
    uint32_t max_entries;


    struct ofp_table_feature_prop_header properties[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_features) == 64) ? 1 : -1; })];


struct ofp_table_stats {
    uint8_t table_id;

    uint8_t pad[3];
    uint32_t active_count;
    uint64_t lookup_count;
    uint64_t matched_count;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_table_stats) == 24) ? 1 : -1; })];


struct ofp_port_stats_request {
    uint32_t port_no;



    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_port_stats_request) == 8) ? 1 : -1; })];



struct ofp_port_stats {
    uint32_t port_no;
    uint8_t pad[4];
    uint64_t rx_packets;
    uint64_t tx_packets;
    uint64_t rx_bytes;
    uint64_t tx_bytes;
    uint64_t rx_dropped;
    uint64_t tx_dropped;
    uint64_t rx_errors;



    uint64_t tx_errors;



    uint64_t rx_frame_err;
    uint64_t rx_over_err;
    uint64_t rx_crc_err;
    uint64_t collisions;
    uint32_t duration_sec;
    uint32_t duration_nsec;

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_port_stats) == 112) ? 1 : -1; })];


struct ofp_group_stats_request {
    uint32_t group_id;
    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_group_stats_request) == 8) ? 1 : -1; })];


struct ofp_bucket_counter {
    uint64_t packet_count;
    uint64_t byte_count;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_bucket_counter) == 16) ? 1 : -1; })];


struct ofp_group_stats {
    uint16_t length;
    uint8_t pad[2];
    uint32_t group_id;
    uint32_t ref_count;

    uint8_t pad2[4];
    uint64_t packet_count;
    uint64_t byte_count;
    uint32_t duration_sec;
    uint32_t duration_nsec;

    struct ofp_bucket_counter bucket_stats[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_group_stats) == 40) ? 1 : -1; })];


struct ofp_group_desc {
    uint16_t length;
    uint8_t type;
    uint8_t pad;
    uint32_t group_id;
    struct ofp_bucket buckets[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_group_desc) == 8) ? 1 : -1; })];





enum ofp_group_capabilities {
    OFPGFC_SELECT_WEIGHT = 1 << 0,
    OFPGFC_SELECT_LIVENESS = 1 << 1,
    OFPGFC_CHAINING = 1 << 2,
    OFPGFC_CHAINING_CHECKS = 1 << 3,
};


struct ofp_group_features {
    uint32_t types;
    uint32_t capabilities;
    uint32_t max_groups[4];
    uint32_t actions[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_group_features) == 40) ? 1 : -1; })];


struct ofp_meter_multipart_request {
    uint32_t meter_id;
    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_multipart_request) == 8) ? 1 : -1; })];


struct ofp_meter_band_stats {
    uint64_t packet_band_count;
    uint64_t byte_band_count;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_band_stats) == 16) ? 1 : -1; })];


struct ofp_meter_stats {
    uint32_t meter_id;
    uint16_t len;
    uint8_t pad[6];
    uint32_t flow_count;
    uint64_t packet_in_count;
    uint64_t byte_in_count;
    uint32_t duration_sec;
    uint32_t duration_nsec;

    struct ofp_meter_band_stats band_stats[0];

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_stats) == 40) ? 1 : -1; })];


struct ofp_meter_config {
    uint16_t length;
    uint16_t flags;
    uint32_t meter_id;
    struct ofp_meter_band_header bands[0];

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_config) == 8) ? 1 : -1; })];


struct ofp_meter_features {
    uint32_t max_meter;
    uint32_t band_types;
    uint32_t capabilities;
    uint8_t max_bands;
    uint8_t max_color;
    uint8_t pad[2];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_meter_features) == 16) ? 1 : -1; })];


struct ofp_experimenter_multipart_header {
    uint32_t experimenter;

    uint32_t exp_type;

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_experimenter_multipart_header) == 8) ? 1 : -1; })];


struct ofp_experimenter_header {
    struct ofp_header header;
    uint32_t experimenter;


    uint32_t exp_type;

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_experimenter_header) == 16) ? 1 : -1; })];

enum ofp_queue_properties {
    OFPQT_MIN_RATE = 1,
    OFPQT_MAX_RATE = 2,
    OFPQT_EXPERIMENTER = 0xffff
};


struct ofp_queue_prop_header {
    uint16_t property;
    uint16_t len;
    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_queue_prop_header) == 8) ? 1 : -1; })];


struct ofp_queue_prop_min_rate {
    struct ofp_queue_prop_header prop_header;
    uint16_t rate;
    uint8_t pad[6];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_queue_prop_min_rate) == 16) ? 1 : -1; })];


struct ofp_queue_prop_max_rate {
    struct ofp_queue_prop_header prop_header;
    uint16_t rate;
    uint8_t pad[6];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_queue_prop_max_rate) == 16) ? 1 : -1; })];


struct ofp_queue_prop_experimenter {
    struct ofp_queue_prop_header prop_header;
    uint32_t experimenter;


    uint8_t pad[4];
    uint8_t data[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_queue_prop_experimenter) == 16) ? 1 : -1; })];


struct ofp_packet_queue {
    uint32_t queue_id;
    uint32_t port;
    uint16_t len;
    uint8_t pad[6];
    struct ofp_queue_prop_header properties[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_packet_queue) == 16) ? 1 : -1; })];


struct ofp_queue_get_config_request {
    struct ofp_header header;
    uint32_t port;



    uint8_t pad[4];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_queue_get_config_request) == 16) ? 1 : -1; })];


struct ofp_queue_get_config_reply {
    struct ofp_header header;
    uint32_t port;
    uint8_t pad[4];
    struct ofp_packet_queue queues[0];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_queue_get_config_reply) == 16) ? 1 : -1; })];


struct ofp_action_set_queue {
    uint16_t type;
    uint16_t len;
    uint32_t queue_id;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_action_set_queue) == 8) ? 1 : -1; })];

struct ofp_queue_stats_request {
    uint32_t port_no;
    uint32_t queue_id;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_queue_stats_request) == 8) ? 1 : -1; })];

struct ofp_queue_stats {
    uint32_t port_no;
    uint32_t queue_id;
    uint64_t tx_bytes;
    uint64_t tx_packets;
    uint64_t tx_errors;
    uint32_t duration_sec;
    uint32_t duration_nsec;

};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_queue_stats) == 40) ? 1 : -1; })];

enum ofp_controller_role {
    OFPCR_ROLE_NOCHANGE = 0,
    OFPCR_ROLE_EQUAL = 1,
    OFPCR_ROLE_MASTER = 2,
    OFPCR_ROLE_SLAVE = 3,
};


struct ofp_role_request {
    struct ofp_header header;
    uint32_t role;
    uint8_t pad[4];
    uint64_t generation_id;
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_role_request) == 24) ? 1 : -1; })];


struct ofp_async_config {
    struct ofp_header header;
    uint32_t packet_in_mask[2];
    uint32_t port_status_mask[2];
    uint32_t flow_removed_mask[2];
};
extern int (*build_assert(void))[ sizeof(struct { unsigned int build_assert_failed : (sizeof(struct ofp_async_config) == 32) ? 1 : -1; })];
