// Code created manually. Would like to be generated in the future.
package netdb

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	AF_APPLETALK              = 16
	AF_CCITT                  = 10
	AF_CHAOS                  = 5
	AF_DATAKIT                = 9
	AF_DECnet                 = 12
	AF_DLI                    = 13
	AF_ECMA                   = 8
	AF_HYLINK                 = 15
	AF_IMPLINK                = 3
	AF_INET                   = 2
	AF_INET6                  = 19
	AF_INTF                   = 20
	AF_ISO                    = AF_NBS
	AF_IUCV                   = 17
	AF_LAT                    = 14
	AF_LINK                   = 18
	AF_MAX                    = 30
	AF_NBS                    = 7
	AF_NDD                    = 23
	AF_NETWARE                = 22
	AF_NS                     = 6
	AF_PUP                    = 4
	AF_RIF                    = 21
	AF_ROUTE                  = 20
	AF_SNA                    = 11
	AF_UNIX                   = 1
	AF_UNSPEC                 = 0
	AI_ADDRCONFIG             = 0x40
	AI_ALL                    = 0x20
	AI_CANONNAME              = 0x02
	AI_EXTFLAGS               = 0x80
	AI_NUMERICHOST            = 0x04
	AI_NUMERICSERV            = 0x08
	AI_PASSIVE                = 0x01
	AI_V4MAPPED               = 0x10
	EAI_AGAIN                 = 2
	EAI_BADEXTFLAGS           = 11
	EAI_BADFLAGS              = 7
	EAI_FAIL                  = 3
	EAI_FAMILY                = 5
	EAI_MAX                   = 10
	EAI_MEMORY                = 6
	EAI_NONAME                = 1
	EAI_OVERFLOW              = 4
	EAI_SERVICE               = 8
	EAI_SOCKTYPE              = 9
	EAI_SYSTEM                = 10
	HOST_NOT_FOUND            = 1
	IBMTCP_IMAGE              = 1
	IMPLINK_HIGHEXPER         = 158
	IMPLINK_IP                = 155
	IMPLINK_LOWEXPER          = 156
	INADDR_ANY                = (In_addr_t)(0x00000000)
	INADDR_BROADCAST          = (In_addr_t)(0xffffffff)
	INADDR_LOOPBACK           = (In_addr_t)(0x7f000001)
	INADDR_NONE               = (In_addr_t)(0xffffffff)
	INET6_ADDRSTRLEN          = 46
	INET_ADDRSTRLEN           = 16
	IN_CLASSA_HOST            = 0x00ffffff
	IN_CLASSA_MAX             = 128
	IN_CLASSA_NET             = 0xff000000
	IN_CLASSA_NSHIFT          = 24
	IN_CLASSB_HOST            = 0x0000ffff
	IN_CLASSB_MAX             = 65536
	IN_CLASSB_NET             = 0xffff0000
	IN_CLASSB_NSHIFT          = 16
	IN_CLASSC_HOST            = 0x000000ff
	IN_CLASSC_NET             = 0xffffff00
	IN_CLASSC_NSHIFT          = 8
	IN_CLASSD_HOST            = 0x0fffffff
	IN_CLASSD_NET             = 0xf0000000
	IN_CLASSD_NSHIFT          = 28
	IPV6_ADDR_PREFERENCES     = 32
	IPV6_CHECKSUM             = 19
	IPV6_DONTFRAG             = 29
	IPV6_DSTOPTS              = 23
	IPV6_HOPLIMIT             = 11
	IPV6_HOPOPTS              = 22
	IPV6_JOIN_GROUP           = 5
	IPV6_LEAVE_GROUP          = 6
	IPV6_MULTICAST_HOPS       = 9
	IPV6_MULTICAST_IF         = 7
	IPV6_MULTICAST_LOOP       = 4
	IPV6_NEXTHOP              = 20
	IPV6_PATHMTU              = 12
	IPV6_PKTINFO              = 13
	IPV6_PREFER_SRC_CGA       = 0x10
	IPV6_PREFER_SRC_COA       = 0x02
	IPV6_PREFER_SRC_HOME      = 0x01
	IPV6_PREFER_SRC_NONCGA    = 0x20
	IPV6_PREFER_SRC_PUBLIC    = 0x08
	IPV6_PREFER_SRC_TMP       = 0x04
	IPV6_RECVDSTOPTS          = 28
	IPV6_RECVHOPLIMIT         = 14
	IPV6_RECVHOPOPTS          = 26
	IPV6_RECVPATHMTU          = 16
	IPV6_RECVPKTINFO          = 15
	IPV6_RECVRTHDR            = 25
	IPV6_RECVTCLASS           = 31
	IPV6_RTHDR                = 21
	IPV6_RTHDRDSTOPTS         = 24
	IPV6_RTHDR_TYPE_0         = 0
	IPV6_TCLASS               = 30
	IPV6_UNICAST_HOPS         = 3
	IPV6_USE_MIN_MTU          = 18
	IPV6_V6ONLY               = 10
	IP_ADD_MEMBERSHIP         = 5
	IP_DEFAULT_MULTICAST_LOOP = 1
	IP_DEFAULT_MULTICAST_TTL  = 1
	IP_DROP_MEMBERSHIP        = 6
	IP_MAX_MEMBERSHIPS        = 20
	IP_MULTICAST_IF           = 7
	IP_MULTICAST_LOOP         = 4
	IP_MULTICAST_TTL          = 3
	IP_OPTIONS                = 1
	IP_TOS                    = 2
	NI_DGRAM                  = 0x10
	NI_NAMEREQD               = 0x04
	NI_NOFQDN                 = 0x01
	NI_NUMERICHOST            = 0x02
	NI_NUMERICSCOPE           = 0x20
	NI_NUMERICSERV            = 0x08
	NO_ADDRESS                = NO_DATA
	NO_DATA                   = 4
	NO_RECOVERY               = 3
	PATH_MAX                  = 1023
	PF_APPLETALK              = AF_APPLETALK
	PF_CCITT                  = AF_CCITT
	PF_CHAOS                  = AF_CHAOS
	PF_DATAKIT                = AF_DATAKIT
	PF_DECnet                 = AF_DECnet
	PF_DLI                    = AF_DLI
	PF_ECMA                   = AF_ECMA
	PF_HYLINK                 = AF_HYLINK
	PF_IMPLINK                = AF_IMPLINK
	PF_INET                   = AF_INET
	PF_INET6                  = AF_INET6
	PF_IUCV                   = AF_IUCV
	PF_LAT                    = AF_LAT
	PF_MAX                    = AF_MAX
	PF_NBS                    = AF_NBS
	PF_NS                     = AF_NS
	PF_PUP                    = AF_PUP
	PF_ROUTE                  = AF_ROUTE
	PF_SNA                    = AF_SNA
	PF_UNIX                   = AF_UNIX
	PF_UNSPEC                 = AF_UNSPEC

	SF_CLOSE                        = 0x00000002
	SF_REUSE                        = 0x00000001
	SOL_SOCKET                      = 0xffff
	SOMAXCONN                       = 10
	SO_ACCEPTCONN                   = 0x0002
	SO_ACCEPTECONNABORTED           = 0x0006
	SO_ACKNOW                       = 0x7700
	SO_BROADCAST                    = 0x0020
	SO_BULKMODE                     = 0x8000
	SO_CKSUMRECV                    = 0x0800
	SO_CLOSE                        = 0x01
	SO_CLUSTERCONNTYPE              = 0x00004001
	SO_CLUSTERCONNTYPE_INTERNAL     = 8
	SO_CLUSTERCONNTYPE_NOCONN       = 0
	SO_CLUSTERCONNTYPE_NONE         = 1
	SO_CLUSTERCONNTYPE_SAME_CLUSTER = 2
	SO_CLUSTERCONNTYPE_SAME_IMAGE   = 4
	SO_DEBUG                        = 0x0001
	SO_DONTROUTE                    = 0x0010
	SO_ERROR                        = 0x1007
	SO_EioIfNewTP                   = 5
	SO_IGNOREINCOMINGPUSH           = 0x1
	SO_IGNORESOURCEVIPA             = 0x0002
	SO_KEEPALIVE                    = 0x0008
	SO_LINGER                       = 0x0080
	SO_NONBLOCKLOCAL                = 0x8001
	SO_NOREUSEADDR                  = 0x1000
	SO_OOBINLINE                    = 0x0100
	SO_OPTACK                       = 0x8004
	SO_OPTMSS                       = 0x8003
	SO_RCVBUF                       = 0x1002
	SO_RCVLOWAT                     = 0x1004
	SO_RCVTIMEO                     = 0x1006
	SO_REUSEADDR                    = 0x0004
	SO_REUSEPORT                    = 0x0200
	SO_SECINFO                      = 0x00004002
	SO_SET                          = 0x0200
	SO_SNDBUF                       = 0x1001
	SO_SNDLOWAT                     = 0x1003
	SO_SNDTIMEO                     = 0x1005
	SO_TYPE                         = 0x1008
	SO_UNSET                        = 0x0400
	SO_USELOOPBACK                  = 0x0040
	SO_USE_IFBUFS                   = 0x0400
	TRY_AGAIN                       = 2
	X_AE_BIMODAL                    = 1
	X_ALL_SOURCE                    = 1
	X_BIG_ENDIAN                    = 1
	X_CHAR_UNSIGNED                 = 1
	X_ENHANCED_ASCII_EXT            = 0x410A0000
	X_EXT                           = 1
	X_ISOC99_SOURCE                 = 1
	X_LARGE_TIME_API                = 1
	X_LONG_LONG                     = 1
	X_LP64                          = 1
	X_MADDR6_SCP_GLO                = 14
	X_MADDR6_SCP_LINK               = 2
	X_MADDR6_SCP_NODE               = 1
	X_MADDR6_SCP_ORG                = 8
	X_MADDR6_SCP_SITE               = 5
	X_MI_BUILTIN                    = 1
	X_OPEN_DEFAULT                  = 1
	X_OPEN_MSGQ_EXT                 = 1
	X_OPEN_SYS                      = 1
	X_OPEN_SYS_FILE_EXT             = 1
	X_OPEN_SYS_SOCK_IPV6            = 1
	X_PATH_HEQUIV                   = "/etc/hosts.equiv"
	X_PATH_HOSTS                    = "/etc/hosts"
	X_PATH_IRS_CONF                 = "/etc/irs.conf"
	X_PATH_NETWORKS                 = "/etc/networks"
	X_PATH_PROTOCOLS                = "/etc/protocols"
	X_PATH_RPCDB                    = "/etc/rpc"
	X_PATH_SERVCONF                 = "/etc/netsvc.conf"
	X_PATH_SERVICES                 = "/etc/services"
	X_POE_ACTION_READ               = 0x00000008
	X_POE_ACTION_SETGET             = 0x00000020
	X_POE_ACTION_WRITE              = 0x00000010
	X_POE_FILE                      = 2
	X_POE_FILE_LEN                  = 4
	X_POE_PROCESS                   = 0x00000002
	X_POE_SCOPE_PROCESS             = 0x00000002
	X_POE_SCOPE_SOCKET              = 0x00000004
	X_POE_SCOPE_THREAD              = 0x00000001
	X_POE_SOCKET                    = 1
	X_POE_SOCKET_LEN                = 4
	X_POE_THREAD                    = 0x00000001
	X_SO_PROPAGATEUSERID            = 0x4000
	X_SO_SELECT                     = 0x02
	X_SS_ALIGNSIZE                  = 8
	X_SS_MAXSIZE                    = 128
	X_UNIX03_SOURCE                 = 1
	X_UNIX03_THREADS                = 1
	X_UNIX03_WITHDRAWN              = 1
	X_XOPEN_CRYPT                   = 1
	X_XOPEN_LEGACY                  = 1
	X_XOPEN_REALTIME                = -1
	X_XOPEN_REALTIME_THREADS        = -1
	X_XOPEN_SOURCE                  = 600
	X_XOPEN_SOURCE_EXTENDED         = 1
	X__features_h                   = 1
	X__inttypes                     = 1
	X__netinet_in                   = 1
	X__netdb                        = 1
	X__types                        = 1
	X__socket_h                     = 1
)

// Bits in the FLAGS argument to `send', `recv', et al.
const (
	MSG_OOB          = 0x1
	MSG_PEEK         = 0x2
	MSG_DONTROUTE    = 0x4
	MSG_CTRUNC       = 0x20
	MSG_TRUNC        = 0x10
	MSG_EOR          = 0x8
	MSG_ACK_EXPECTED = 0x10
	MSG_ACK_GEN      = 0x40
	MSG_ACK_TIMEOUT  = 0x20
	MSG_CONNTERM     = 0x80
	MSG_EOF          = 0x8000
	MSG_MAXIOVLEN    = 16
	MSG_NONBLOCK     = 0x4000
	MSG_WAITALL      = 0x40
)

// Socket level message types.  This must match the definitions in
//    <linux/socket.h>.
const (
	SCM_RIGHTS = 0x01
)

// Types of sockets.
const (
	SOCK_CONN_DGRAM = 6
	SOCK_DGRAM      = 2
	SOCK_RAW        = 3
	SOCK_RDM        = 4
	SOCK_SEQPACKET  = 5
	SOCK_STREAM     = 1
)

// Standard well-known ports.
const (
	IPPORT_RESERVED     = 1024
	IPPORT_USERRESERVED = 5000
)

// Standard well-defined IP protocols.
const (
	IPPROTO_IP   = 0
	IPPROTO_ICMP = 1
	IPPROTO_GGP  = 2
	IPPROTO_TCP  = 6
	IPPROTO_EGP  = 8
	IPPROTO_PUP  = 12
	IPPROTO_UDP  = 17
	IPPROTO_IDP  = 22
	IPPROTO_IPV6 = 41
	IPPROTO_ESP  = 50
	IPPROTO_AH   = 51
	IPPROTO_RAW  = 255
	IPPROTO_MAX  = 256
)

// If __USE_KERNEL_IPV6_DEFS is 1 then the user has included the kernel
//    network headers first and we should use those ABI-identical definitions
//    instead of our own, otherwise 0.
const (
	IPPROTO_HOPOPTS  = 0
	IPPROTO_ROUTING  = 43
	IPPROTO_FRAGMENT = 44
	IPPROTO_ICMPV6   = 58
	IPPROTO_NONE     = 59
	IPPROTO_DSTOPTS  = 60
)

// The following constants should be used for the second parameter of
//    `shutdown'.
const (
	SHUT_RD   = 0
	SHUT_WR   = 1
	SHUT_RDWR = 2
)

var (
	IN6ADDR_ANY_INIT      = getIn6Addr_Any_Init()
	IN6ADDR_LOOPBACK_INIT = getIn6Addr_Loopback_Init()
)

func getIn6Addr_Any_Init() [16]uint8 {
	return [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func getIn6Addr_Loopback_Init() [16]uint8 {
	return [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
}

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

// Convenience types.
type X__u_char = uint8   /* sys/types.h:481 */
type X__u_short = uint16 /* sys/types.h:496 */
type X__u_int = uint32   /* sys/types.h:486 */
type X__u_long = uint64  /* sys/types.h:501 */

// Fixed-size types, underlying types depend on word size and compiler.
type X__int8_t = int8
type X__int16_t = int16
type X__int32_t = int32
type X__int64_t = int64
type X__uint8_t = uint8
type X__uint16_t = uint16
type X__uint32_t = uint32
type X__uint64_t = uint64

// Smallest types with at least a given width.
type X__int_least8_t = X__int8_t
type X__int_least16_t = X__int16_t
type X__int_least32_t = X__int32_t
type X__int_least64_t = X__int64_t
type X__uint_least8_t = X__uint8_t
type X__uint_least16_t = X__uint16_t
type X__uint_least32_t = X__uint32_t
type X__uint_least64_t = X__uint64_t

type X__dev_t = uint32      /* sys/types.h:66 */
type X__uid_t = int32       /* sys/types.h:162 */
type X__gid_t = int32       /* sys/types.h:77 */
type X__ino_t = uint32      /* sys/types.h:86 */
type X__mode_t = int32      /* sys/types.h:93 */
type X__nlink_t = int32     /* sys/types.h:101 */
type X__off_t = int64       /* sys/types.h:113 */
type X__off64_t = int64     /* sys/types.h:123 */
type X__pid_t = int32       /* sys/types.h:129 */
type X__clock_t = uint32    /* sys/types.h:129 */
type X__id_t = int32        /* sys/types.h:430 */
type X__time_t = int64      /* sys/types.h:182 */
type X__time64_t = uint64   /* sys/types.h:570 */
type X__useconds_t = uint32 /* sys/types.h:437 */
type X__suseconds_t = int32 /* sys/types.h:559 */

type X__key_t = int32 /* sys/types.h:407 */

// Type to represent block size
type X__blksize_t = int32 /* sys/types.h:550 */

// Type to count number of disk blocks
type X__blkcnt_t = int64 /* sys/types.h:547 */

// Type to count file system blocks.
type X__fsblkcnt_t = uint32 /* sys/types.h:552 */

// Type to count file system nodes.
type X__fsfilcnt_t = uint32 /* sys/types.h:555 */

type X__ssize_t = int32 /* sys/types.h:148 */

type X__caddr_t = uintptr /* sys/types.h:529 */

// Duplicates info from stdint.h but this is used in unistd.h.
type X__intptr_t = int64

// Duplicate info from sys/socket.h.
type X__socklen_t = uint32 /* socket.h:115 */

type X__sig_atomic_t = int32 /* signal.h:52 */

type U_char = X__u_char
type U_short = X__u_short
type U_int = X__u_int
type U_long = X__u_long

type Ino_t = X__ino_t
type Dev_t = X__dev_t
type Gid_t = X__gid_t
type Mode_t = X__mode_t
type Nlink_t = X__nlink_t
type Uid_t = X__uid_t
type Off_t = X__off64_t
type Pid_t = X__pid_t
type Id_t = X__id_t
type Ssize_t = X__ssize_t
type Caddr_t = X__caddr_t
type Key_t = X__key_t

type Clock_t = X__clock_t

type Time_t = X__time_t

type Ulong = uint64
type Ushort = uint16
type Uint = uint32

type Int8_t = X__int8_t
type Int16_t = X__int16_t
type Int32_t = X__int32_t
type Int64_t = X__int64_t

// These were defined by ISO C without the first `_'.
type U_int8_t = X__uint8_t
type U_int16_t = X__uint16_t
type U_int32_t = X__uint32_t
type U_int64_t = X__uint64_t

// Adjusted for zos
type X__sigset_t = struct {
	F__sigs_0 uint32
	F__sigs_1 uint32
} /* sys/types.h:200 */

type Sigset_t = X__sigset_t

// Adjusted for zos
type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec int64
} /* time.h:64 */

type Useconds_t = X__useconds_t
type Suseconds_t = X__suseconds_t

type Blksize_t = X__blksize_t

// TODO - These three types have been adjusted, but may need more work
type Blkcnt_t = X__blkcnt_t
type Fsblkcnt_t = X__fsblkcnt_t
type Fsfilcnt_t = X__fsfilcnt_t

type Pthread_t = struct {
	__ [0x08]uint8
} /* sys/types.h:282 */

// Data structures for mutex handling.  The structure of the attribute
//    type is not exposed on purpose.
type Pthread_mutexattr_t = struct {
	__ [0x04]uint8
} /* sys/types.h:314 */

// Data structure for condition variable handling.  The structure of
//    the attribute type is not exposed on purpose.
type Pthread_condattr_t = struct {
	__ [0x08]uint8
} /* sys/types.h:343 */

type Pthread_key_t = struct {
	__ [0x08]uint8
} /* sys/types.h:354 */

type Pthread_once_t = int32 /* sys/types.h:365 */

type Pthread_attr_t = struct {
	__ [13]float64
} /* sys/types.h:289 */

type Pthread_mutex_t = struct {
	F__m uint64
	F__d [8]float64
} /* sys/types:305 */

type Pthread_cond_t = struct {
	F__m int64
	F__d [8]float64
} /* sys/types:334 */

type Pthread_rwlock_t = struct {
	__ [0x04]uint8
} /* sys/types.h:382 */

type Pthread_rwlockattr_t = struct {
	__ [0x04]uint8
} /* sys/types.h:371 */

type Socklen_t = X__socklen_t

type Sa_family_t = uint8 /* socket.h:460 */

type Sockaddr = struct {
	Fsa_len    uint8
	Fsa_family Sa_family_t
	Fda_data   [14]uint8
} /* socket.h:467 */

type Sockaddr_storage = struct {
	Fss_len     U_int8_t
	Fsa_family  Sa_family_t
	F_ss_pad1   [6]uint8
	F__ss_align Int64_t
	F__ss_pad2  [112]uint8
} /* socket.h:553 */

// Structure describing messages sent by
//    `sendmsg' and received by `recvmsg'.
type Msghdr = struct {
	Fmsg_name       uintptr
	Fmsg_iov        uintptr
	Fmsg_control    uintptr
	Fmsg_flags      int32
	Fmsg_namelen    Socklen_t
	Fmsg_iovlen     int32
	Fmsg_controllen Socklen_t
} /* socket.h:597 */

// Structure used for storage of ancillary data object information.
type Cmsghdr = struct {
	Fcmsg_len   Socklen_t
	Fcmsg_level int32
	Fcmsg_type  int32
} /* socket.h:626 */

// Structure used to manipulate the SO_LINGER option.
type Linger = struct {
	Fl_onoff  int32
	Fl_linger int32
} /* socket.h:380 */

type In_addr_t = uint32
type In_addr = struct{ Fs_addr In_addr_t }

type In_pktinfo = struct {
	Fipi_addr    struct{ Fs_addr In_addr_t }
	Fipi_ifindex uint32
}

type In_port_t = uint16

type in6_addr = struct {
	F_S6_un struct {
		F_S6_u8 [16]U_int8_t
		_       [4]U_int32_t
	}
}

type Sockaddr_in = struct {
	Fsin_len    uint8
	Fsin_family Sa_family_t
	Fsin_port   In_port_t
	Fsin_addr   struct{ Fs_addr In_addr_t }
	Fsin_zero   [8]uint8
}

type Sockaddr_in6 = struct {
	Fsin6_len      U_int8_t
	Fsin6_family   Sa_family_t
	Fsin6_port     In_port_t
	Fsin6_flowinfo U_int32_t
	Fsin6_addr     struct {
		F_S6_un struct {
			F_S6_u8 [16]U_int8_t
			_       [4]U_int32_t
		}
	}
	Fsin6_scope_id U_int32_t
}

type Ip_mreq = struct {
	Fimr_multiaddr struct{ Fs_addr In_addr_t }
	Fimr_interface struct{ Fs_addr In_addr_t }
}

type Ip_mreq_source = struct {
	Fimr_multiaddr  struct{ Fs_addr In_addr_t }
	Fimr_interface  struct{ Fs_addr In_addr_t }
	Fimr_sourceaddr struct{ Fs_addr In_addr_t }
}

type Ipv6_mreq = struct {
	Fipv6_multiaddr struct {
		F_S6_un struct {
			F_S6_u8 [16]U_int8_t
			_       [4]U_int32_t
		}
	}
	Fipv6mr_interface uint32
}

type Group_req = struct {
	Fgr_interface U_int32_t
	Fgr_group     Sockaddr_storage
	F__gr_01      U_int32_t
}

type Group_source_req = struct {
	Fgsr_interface U_int32_t
	F__gsr_01      U_int32_t
	Fgsr_group     Sockaddr_storage
	Fgsr_source    Sockaddr_storage
}

type Rpcent = struct {
	Fr_name    uintptr
	Fr_aliases uintptr
	Fr_number  int32
} /* rpc/netdb.h:40 */

// Description of data base entry for a single network.  NOTE: here a
//    poor assumption is made.  The network number is expected to fit
//    into an unsigned long int variable.
type Netent = struct {
	Fn_name     uintptr
	Fn_aliases  uintptr
	Fn_addrtype int32
	Fn_net      U_int32_t
} /* netdb.h:104 */

// Description of data base entry for a single host.
type Hostent = struct {
	Fh_name      uintptr
	Fh_aliases   uintptr
	Fh_addrtype  int32
	Fh_length    int32
	Fh_addr_list uintptr
} /* netdb.h:96 */

// Description of data base entry for a single service.
type Servent = struct {
	Fs_name    uintptr
	Fs_aliases uintptr
	Fs_port    int32
	Fs_proto   uintptr
} /* netdb.h:115 */

// Description of data base entry for a single service.
type Protoent = struct {
	Fp_name    uintptr
	Fp_aliases uintptr
	Fp_proto   int32
} /* netdb.h:122 */

// Extension from POSIX.1:2001.
// Structure to contain information about address of a service provider.
type Addrinfo = struct {
	Fai_flags     int32
	Fai_family    int32
	Fai_socktype  int32
	Fai_protocol  int32
	Fai_addrlen   Socklen_t
	Fai_reserved  int32
	Fai_canonname uintptr
	Fai_addr      uintptr
	Fai_next      uintptr
} /* netdb.h:565 */

var _ uint8
