// Code created manually. Would like to be generated in the future.

package errno

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
	EDOM                     = 1
	ERANGE                   = 2
	EACCES                   = 111
	EAGAIN                   = 112
	EBADF                    = 113
	EBUSY                    = 114
	ECHILD                   = 115
	EDEADLK                  = 116
	EEXIST                   = 117
	EFAULT                   = 118
	EFBIG                    = 119
	EINTR                    = 120
	EINVAL                   = 121
	EIO                      = 122
	EISDIR                   = 123
	EMFILE                   = 124
	EMLINK                   = 125
	ENAMETOOLONG             = 126
	ENFILE                   = 127
	ENODEV                   = 128
	ENOENT                   = 129
	ENOEXEC                  = 130
	ENOLCK                   = 131
	ENOMEM                   = 132
	ENOSPC                   = 133
	ENOSYS                   = 134
	ENOTDIR                  = 135
	ENOTEMPTY                = 136
	ENOTTY                   = 137
	ENXIO                    = 138
	EPERM                    = 139
	EPIPE                    = 140
	EROFS                    = 141
	ESPIPE                   = 142
	ESRCH                    = 143
	EXDEV                    = 144
	E2BIG                    = 145
	ELOOP                    = 146
	EILSEQ                   = 147
	ENODATA                  = 148
	EOVERFLOW                = 149
	EMVSNOTUP                = 150
	ECMSSTORAGE              = 151
	EMVSDYNALC               = 151
	EMVSCVAF                 = 152
	EMVSCATLG                = 153
	ECMSINITIAL              = 156
	EMVSINITIAL              = 156
	ECMSERR                  = 157
	EMVSERR                  = 157
	EMVSPARM                 = 158
	ECMSPFSFILE              = 159
	EMVSPFSFILE              = 159
	EMVSBADCHAR              = 160
	ECMSPFSPERM              = 162
	EMVSPFSPERM              = 162
	EMVSSAFEXTRERR           = 163
	EMVSSAF2ERR              = 164
	EMVSTODNOTSET            = 165
	EMVSPATHOPTS             = 166
	EMVSNORTL                = 167
	EMVSEXPIRE               = 168
	EMVSPASSWORD             = 169
	EMVSWLMERROR             = 170
	EMVSCPLERROR             = 171
	EMVSARMERROR             = 172
	ELENOFORK                = 200
	ELEMSGERR                = 201
	EBUFLEN                  = 227
	EEXTLINK                 = 228
	ENODD                    = 229
	ECMSESMERR               = 230
	ECPERR                   = 231
	ELEMULTITHREAD           = 232
	ELEFENCE                 = 244
	ENOTSUP                  = 247
	ELEMULTITHREADFORK       = 257
	ECUNNOENV                = 258
	ECUNNOCONV               = 259
	ECUNNOTALIGNED           = 260
	ECUNERR                  = 262
	EIBMBADCALL              = 1000
	EIBMBADPARM              = 1001
	EIBMSOCKOUTOFRANGE       = 1002
	EIBMSOCKINUSE            = 1003
	EIBMIUCVERR              = 1004
	EOFFLOADboxERROR         = 1005
	EOFFLOADboxRESTART       = 1006
	EOFFLOADboxDOWN          = 1007
	EIBMCONFLICT             = 1008
	EIBMCANCELLED            = 1009
	EIBMBADTCPNAME           = 1011
	ENOTBLK                  = 1100
	ETXTBSY                  = 1101
	EWOULDBLOCK              = 1102
	EINPROGRESS              = 1103
	EALREADY                 = 1104
	ENOTSOCK                 = 1105
	EDESTADDRREQ             = 1106
	EMSGSIZE                 = 1107
	EPROTOTYPE               = 1108
	ENOPROTOOPT              = 1109
	EPROTONOSUPPORT          = 1110
	ESOCKTNOSUPPORT          = 1111
	EOPNOTSUPP               = 1112
	EPFNOSUPPORT             = 1113
	EAFNOSUPPORT             = 1114
	EADDRINUSE               = 1115
	EADDRNOTAVAIL            = 1116
	ENETDOWN                 = 1117
	ENETUNREACH              = 1118
	ENETRESET                = 1119
	ECONNABORTED             = 1120
	ECONNRESET               = 1121
	ENOBUFS                  = 1122
	EISCONN                  = 1123
	ENOTCONN                 = 1124
	ESHUTDOWN                = 1125
	ETOOMANYREFS             = 1126
	ETIMEDOUT                = 1127
	ECONNREFUSED             = 1128
	EHOSTDOWN                = 1129
	EHOSTUNREACH             = 1130
	EPROCLIM                 = 1131
	EUSERS                   = 1132
	EDQUOT                   = 1133
	ESTALE                   = 1134
	EREMOTE                  = 1135
	ENOSTR                   = 1136
	ETIME                    = 1137
	ENOSR                    = 1138
	ENOMSG                   = 1139
	EBADMSG                  = 1140
	EIDRM                    = 1141
	ENONET                   = 1142
	ERREMOTE                 = 1143
	ENOLINK                  = 1144
	EADV                     = 1145
	ESRMNT                   = 1146
	ECOMM                    = 1147
	EPROTO                   = 1148
	EMULTIHOP                = 1149
	EDOTDOT                  = 1150
	EREMCHG                  = 1151
	ECANCELED                = 1152
	EINTRNODATA              = 1159
	ENOREUSE                 = 1160
	ENOMOVE                  = 1161
	PATH_MAX                 = 1023
	X_AE_BIMODAL             = 1
	X_ALL_SOURCE             = 1
	X_BIG_ENDIAN             = 1
	X_CHAR_UNSIGNED          = 1
	X_ENHANCED_ASCII_EXT     = 0x410A0000
	X_EXT                    = 1
	X_ISOC99_SOURCE          = 1
	X_LARGE_TIME_API         = 1
	X_LONG_LONG              = 1
	X_LP64                   = 1
	X_MI_BUILTIN             = 1
	X_OPEN_DEFAULT           = 1
	X_OPEN_MSGQ_EXT          = 1
	X_OPEN_SYS               = 1
	X_OPEN_SYS_FILE_EXT      = 1
	X_OPEN_SYS_SOCK_IPV6     = 1
	X_UNIX03_SOURCE          = 1
	X_UNIX03_THREADS         = 1
	X_UNIX03_WITHDRAWN       = 1
	X_XOPEN_CRYPT            = 1
	X_XOPEN_LEGACY           = 1
	X_XOPEN_REALTIME         = -1
	X_XOPEN_REALTIME_THREADS = -1
	X_XOPEN_SOURCE           = 600
	X_XOPEN_SOURCE_EXTENDED  = 1
	X__errnoh                = 1
	X__features_h            = 1
)

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

var _ int8
