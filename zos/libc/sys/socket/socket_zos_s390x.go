// Code created manually. Would like to be generated in the future.

package socket

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
	AF_APPLETALK                    = 16
	AF_CCITT                        = 10
	AF_CHAOS                        = 5
	AF_DATAKIT                      = 9
	AF_DECnet                       = 12
	AF_DLI                          = 13
	AF_ECMA                         = 8
	AF_HYLINK                       = 15
	AF_IMPLINK                      = 3
	AF_INET                         = 2
	AF_INET6                        = 19
	AF_INTF                         = 20
	AF_ISO                          = AF_NBS
	AF_IUCV                         = 17
	AF_LAT                          = 14
	AF_LINK                         = 18
	AF_MAX                          = 30
	AF_NBS                          = 7
	AF_NDD                          = 23
	AF_NETWARE                      = 22
	AF_NS                           = 6
	AF_PUP                          = 4
	AF_RIF                          = 21
	AF_ROUTE                        = 20
	AF_SNA                          = 11
	AF_UNIX                         = 1
	AF_UNSPEC                       = 0
	IBMTCP_IMAGE                    = 1
	IMPLINK_HIGHEXPER               = 158
	IMPLINK_IP                      = 155
	IMPLINK_LOWEXPER                = 156
	INADDR_ANY                      = (X__in_addr_t)(0x00000000)
	INADDR_BROADCAST                = (X__in_addr_t)(0xffffffff)
	INADDR_LOOPBACK                 = (X__in_addr_t)(0x7f000001)
	INADDR_NONE                     = (X__in_addr_t)(0xffffffff)
	INET6_ADDRSTRLEN                = 46
	INET_ADDRSTRLEN                 = 16
	IN_CLASSA_HOST                  = 0x00ffffff
	IN_CLASSA_MAX                   = 128
	IN_CLASSA_NET                   = 0xff000000
	IN_CLASSA_NSHIFT                = 24
	IN_CLASSB_HOST                  = 0x0000ffff
	IN_CLASSB_MAX                   = 65536
	IN_CLASSB_NET                   = 0xffff0000
	IN_CLASSB_NSHIFT                = 16
	IN_CLASSC_HOST                  = 0x000000ff
	IN_CLASSC_NET                   = 0xffffff00
	IN_CLASSC_NSHIFT                = 8
	IN_CLASSD_HOST                  = 0x0fffffff
	IN_CLASSD_NET                   = 0xf0000000
	IN_CLASSD_NSHIFT                = 28
	IPPORT_RESERVED                 = 1024
	IPPORT_USERRESERVED             = 5000
	IPPROTO_AH                      = 51
	IPPROTO_DSTOPTS                 = 60
	IPPROTO_EGP                     = 8
	IPPROTO_ESP                     = 50
	IPPROTO_FRAGMENT                = 44
	IPPROTO_GGP                     = 2
	IPPROTO_HOPOPTS                 = 0
	IPPROTO_ICMP                    = 1
	IPPROTO_ICMPV6                  = 58
	IPPROTO_IDP                     = 22
	IPPROTO_IP                      = 0
	IPPROTO_IPV6                    = 41
	IPPROTO_MAX                     = 256
	IPPROTO_NONE                    = 59
	IPPROTO_PUP                     = 12
	IPPROTO_RAW                     = 255
	IPPROTO_ROUTING                 = 43
	IPPROTO_TCP                     = 6
	IPPROTO_UDP                     = 17
	IPV6_ADDR_PREFERENCES           = 32
	IPV6_CHECKSUM                   = 19
	IPV6_DONTFRAG                   = 29
	IPV6_DSTOPTS                    = 23
	IPV6_HOPLIMIT                   = 11
	IPV6_HOPOPTS                    = 22
	IPV6_JOIN_GROUP                 = 5
	IPV6_LEAVE_GROUP                = 6
	IPV6_MULTICAST_HOPS             = 9
	IPV6_MULTICAST_IF               = 7
	IPV6_MULTICAST_LOOP             = 4
	IPV6_NEXTHOP                    = 20
	IPV6_PATHMTU                    = 12
	IPV6_PKTINFO                    = 13
	IPV6_PREFER_SRC_CGA             = 0x10
	IPV6_PREFER_SRC_COA             = 0x02
	IPV6_PREFER_SRC_HOME            = 0x01
	IPV6_PREFER_SRC_NONCGA          = 0x20
	IPV6_PREFER_SRC_PUBLIC          = 0x08
	IPV6_PREFER_SRC_TMP             = 0x04
	IPV6_RECVDSTOPTS                = 28
	IPV6_RECVHOPLIMIT               = 14
	IPV6_RECVHOPOPTS                = 26
	IPV6_RECVPATHMTU                = 16
	IPV6_RECVPKTINFO                = 15
	IPV6_RECVRTHDR                  = 25
	IPV6_RECVTCLASS                 = 31
	IPV6_RTHDR                      = 21
	IPV6_RTHDRDSTOPTS               = 24
	IPV6_RTHDR_TYPE_0               = 0
	IPV6_TCLASS                     = 30
	IPV6_UNICAST_HOPS               = 3
	IPV6_USE_MIN_MTU                = 18
	IPV6_V6ONLY                     = 10
	IP_ADD_MEMBERSHIP               = 5
	IP_DEFAULT_MULTICAST_LOOP       = 1
	IP_DEFAULT_MULTICAST_TTL        = 1
	IP_DROP_MEMBERSHIP              = 6
	IP_MAX_MEMBERSHIPS              = 20
	IP_MULTICAST_IF                 = 7
	IP_MULTICAST_LOOP               = 4
	IP_MULTICAST_TTL                = 3
	IP_OPTIONS                      = 1
	IP_TOS                          = 2
	PF_APPLETALK                    = AF_APPLETALK
	PF_CCITT                        = AF_CCITT
	PF_CHAOS                        = AF_CHAOS
	PF_DATAKIT                      = AF_DATAKIT
	PF_DECnet                       = AF_DECnet
	PF_DLI                          = AF_DLI
	PF_ECMA                         = AF_ECMA
	PF_HYLINK                       = AF_HYLINK
	PF_IMPLINK                      = AF_IMPLINK
	PF_INET                         = AF_INET
	PF_INET6                        = AF_INET6
	PF_IUCV                         = AF_IUCV
	PF_LAT                          = AF_LAT
	PF_MAX                          = AF_MAX
	PF_NBS                          = AF_NBS
	PF_NS                           = AF_NS
	PF_PUP                          = AF_PUP
	PF_ROUTE                        = AF_ROUTE
	PF_SNA                          = AF_SNA
	PF_UNIX                         = AF_UNIX
	PF_UNSPEC                       = AF_UNSPEC
	PRIX16                          = "hX"
	PRIX32                          = "X"
	PRIX64                          = "lX"
	PRIX8                           = "hhX"
	PRIXFAST16                      = "X"
	PRIXFAST32                      = "X"
	PRIXFAST64                      = "lX"
	PRIXFAST8                       = "hhX"
	PRIXLEAST16                     = "hX"
	PRIXLEAST32                     = "X"
	PRIXLEAST64                     = "lX"
	PRIXLEAST8                      = "hhX"
	PRIXMAX                         = "jX"
	PRIXPTR                         = "lX"
	PRId16                          = "hd"
	PRId32                          = "d"
	PRId64                          = "ld"
	PRId8                           = "hhd"
	PRIdFAST16                      = "d"
	PRIdFAST32                      = "d"
	PRIdFAST64                      = "ld"
	PRIdFAST8                       = "hhd"
	PRIdLEAST16                     = "hd"
	PRIdLEAST32                     = "d"
	PRIdLEAST64                     = "ld"
	PRIdLEAST8                      = "hhd"
	PRIdMAX                         = "jd"
	PRIdPTR                         = "ld"
	PRIi16                          = "hi"
	PRIi32                          = "i"
	PRIi64                          = "li"
	PRIi8                           = "hhi"
	PRIiFAST16                      = "i"
	PRIiFAST32                      = "i"
	PRIiFAST64                      = "li"
	PRIiFAST8                       = "hhi"
	PRIiLEAST16                     = "hi"
	PRIiLEAST32                     = "i"
	PRIiLEAST64                     = "li"
	PRIiLEAST8                      = "hhi"
	PRIiMAX                         = "ji"
	PRIiPTR                         = "li"
	PRIo16                          = "ho"
	PRIo32                          = "o"
	PRIo64                          = "lo"
	PRIo8                           = "hho"
	PRIoFAST16                      = "o"
	PRIoFAST32                      = "o"
	PRIoFAST64                      = "lo"
	PRIoFAST8                       = "hho"
	PRIoLEAST16                     = "ho"
	PRIoLEAST32                     = "o"
	PRIoLEAST64                     = "lo"
	PRIoLEAST8                      = "hho"
	PRIoMAX                         = "jo"
	PRIoPTR                         = "lo"
	PRIu16                          = "hu"
	PRIu32                          = "u"
	PRIu64                          = "lu"
	PRIu8                           = "hhu"
	PRIuFAST16                      = "u"
	PRIuFAST32                      = "u"
	PRIuFAST64                      = "lu"
	PRIuFAST8                       = "hhu"
	PRIuLEAST16                     = "hu"
	PRIuLEAST32                     = "u"
	PRIuLEAST64                     = "lu"
	PRIuLEAST8                      = "hhu"
	PRIuMAX                         = "ju"
	PRIuPTR                         = "lu"
	PRIx16                          = "hx"
	PRIx32                          = "x"
	PRIx64                          = "lx"
	PRIx8                           = "hhx"
	PRIxFAST16                      = "x"
	PRIxFAST32                      = "x"
	PRIxFAST64                      = "lx"
	PRIxFAST8                       = "hhx"
	PRIxLEAST16                     = "hx"
	PRIxLEAST32                     = "x"
	PRIxLEAST64                     = "lx"
	PRIxLEAST8                      = "hhx"
	PRIxMAX                         = "jx"
	PRIxPTR                         = "lx"
	SCNd16                          = "hd"
	SCNd32                          = "d"
	SCNd64                          = "ld"
	SCNd8                           = "hhd"
	SCNdFAST16                      = "d"
	SCNdFAST32                      = "d"
	SCNdFAST64                      = "ld"
	SCNdFAST8                       = "hhd"
	SCNdLEAST16                     = "hd"
	SCNdLEAST32                     = "d"
	SCNdLEAST64                     = "ld"
	SCNdLEAST8                      = "hhd"
	SCNdMAX                         = "jd"
	SCNdPTR                         = "ld"
	SCNi16                          = "hi"
	SCNi32                          = "i"
	SCNi64                          = "li"
	SCNi8                           = "hhi"
	SCNiFAST16                      = "i"
	SCNiFAST32                      = "i"
	SCNiFAST64                      = "li"
	SCNiFAST8                       = "hhi"
	SCNiLEAST16                     = "hi"
	SCNiLEAST32                     = "i"
	SCNiLEAST64                     = "li"
	SCNiLEAST8                      = "hhi"
	SCNiMAX                         = "ji"
	SCNiPTR                         = "li"
	SCNo16                          = "ho"
	SCNo32                          = "o"
	SCNo64                          = "lo"
	SCNo8                           = "hho"
	SCNoFAST16                      = "o"
	SCNoFAST32                      = "o"
	SCNoFAST64                      = "lo"
	SCNoFAST8                       = "hho"
	SCNoLEAST16                     = "ho"
	SCNoLEAST32                     = "o"
	SCNoLEAST64                     = "lo"
	SCNoLEAST8                      = "hho"
	SCNoMAX                         = "jo"
	SCNoPTR                         = "lo"
	SCNu16                          = "hu"
	SCNu32                          = "u"
	SCNu64                          = "lu"
	SCNu8                           = "hhu"
	SCNuFAST16                      = "u"
	SCNuFAST32                      = "u"
	SCNuFAST64                      = "lu"
	SCNuFAST8                       = "hhu"
	SCNuLEAST16                     = "hu"
	SCNuLEAST32                     = "u"
	SCNuLEAST64                     = "lu"
	SCNuLEAST8                      = "hhu"
	SCNuMAX                         = "ju"
	SCNuPTR                         = "lu"
	SCNx16                          = "hx"
	SCNx32                          = "x"
	SCNx64                          = "lx"
	SCNx8                           = "hhx"
	SCNxFAST16                      = "x"
	SCNxFAST32                      = "x"
	SCNxFAST64                      = "lx"
	SCNxFAST8                       = "hhx"
	SCNxLEAST16                     = "hx"
	SCNxLEAST32                     = "x"
	SCNxLEAST64                     = "lx"
	SCNxLEAST8                      = "hhx"
	SCNxMAX                         = "jx"
	SCNxPTR                         = "lx"
	SF_REUSE                        = 0x00000001
	SF_CLOSE                        = 0x00000002
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
)

const (
	MSG_ACK_EXPECTED = 0x10
	MSG_ACK_GEN      = 0x40
	MSG_ACK_TIMEOUT  = 0x20
	MSG_CONNTERM     = 0x80
	MSG_CTRUNC       = 0x20
	MSG_DONTROUTE    = 0x4
	MSG_EOF          = 0x8000
	MSG_EOR          = 0x8
	MSG_MAXIOVLEN    = 16
	MSG_NONBLOCK     = 0x4000
	MSG_OOB          = 0x1
	MSG_PEEK         = 0x2
	MSG_TRUNC        = 0x10
	MSG_WAITALL      = 0x40
)

const (
	SCM_RIGHTS = 0x01
)

const (
	SOCK_STREAM     = 1
	SOCK_DGRAM      = 2
	SOCK_RAW        = 3
	SOCK_RDM        = 4
	SOCK_SEQPACKET  = 5
	SOCK_CONN_DGRAM = 6
)

const (
	SHUT_RD   = 0
	SHUT_RDWR = 2
	SHUT_WR   = 1
)

// Since go doesn't allow constant arrays
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

type Ptrdiff_t = int64

type Size_t = uint64

type Wchar_t = uint32

// Couldn't find references to these on zos, but all other implementations have them
type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr
type X__float128 = float64

// Structure for scatter/gather I/O.
type Iovec = struct {
	Fiov_base uintptr
	Fiov_len  Size_t
} /* sys/uio.h:100 */

// Convenience types.
type X__u_char = uint8
type X__u_short = uint16
type X__u_int = uint32
type X__u_long = uint64

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
type X__uid_t = int32       /* sys/types.h:162 */ // uid_t is uint32 if __VM__ is defined
type X__gid_t = int32       /* sys/types.h:77 */  // gid_t is uint32 if __VM__ is defined
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

type U_int8_t = X__uint8_t
type U_int16_t = X__uint16_t
type U_int32_t = X__uint32_t
type U_int64_t = X__uint64_t

// (ZOS) - Adjusted
type X__sigset_t = struct {
	X__sigs_0 uint32
	X__sigs_1 uint32
} /* sys/types.h:200 */

// A set of signals to be blocked, unblocked, or waited for.
type Sigset_t = X__sigset_t

// An integral type that can be modified atomically, without the
//    possibility of a signal arriving in the middle of the operation.
type Sig_atomic_t = X__sig_atomic_t

// Adjusted for zos
type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec int64
}

type Useconds_t = X__useconds_t
type Suseconds_t = X__suseconds_t

type Blksize_t = X__blksize_t

type Blkcnt_t = X__blkcnt_t
type Fsblkcnt_t = X__fsblkcnt_t
type Fsfilcnt_t = X__fsfilcnt_t

type X__in_addr_t = uint32

// TODO (ZOS) - May need adjustment
// Thread identifiers.  The structure of the attribute type is not
//    exposed on purpose.
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

var _ uint8
