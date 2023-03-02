// Code created manually. Would like to be generated in the future.

package termios

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
	B0                       = 0
	B110                     = 3
	B1200                    = 9
	B134                     = 4
	B150                     = 5
	B1800                    = 10
	B19200                   = 14
	B200                     = 6
	B2400                    = 11
	B300                     = 7
	B38400                   = 15
	B4800                    = 12
	B50                      = 1
	B600                     = 8
	B75                      = 2
	B9600                    = 13
	BRKINT                   = 0x0001
	BS0                      = 0x0000
	BS1                      = 0x8000
	BSDLY                    = 0x8000
	CLOCAL                   = 0x0001
	CR0                      = 0x0000
	CR1                      = 0x1000
	CR2                      = 0x2000
	CR3                      = 0x3000
	CRDLY                    = 0x3000
	CREAD                    = 0x0002
	CS5                      = 0x0000
	CS6                      = 0x0010
	CS7                      = 0x0020
	CS8                      = 0x0030
	CSIZE                    = 0x0030
	CSTOPB                   = 0x0080
	ECHO                     = 0x00000008
	ECHOE                    = 0x0002
	ECHOK                    = 0x0004
	ECHONL                   = 0x00000001
	FF0                      = 0x0000
	FF1                      = 0x4000
	FFDLY                    = 0x4000
	HUPCL                    = 0x0100
	ICANON                   = 0x0010
	ICRNL                    = 0x0002
	IEXTEN                   = 0x0020
	IGNBRK                   = 0x0004
	IGNCR                    = 0x0008
	IGNPAR                   = 0x0010
	INLCR                    = 0x0020
	INPCK                    = 0x0040
	IOCGSMAXSTKS             = 32
	IOCRTRERROR              = 0
	IOCUO_CS                 = 3
	IOCUO_READ               = 1
	IOCUO_WRITE              = 2
	ISIG                     = 0x0040
	ISTRIP                   = 0x0080
	IUCLC                    = 0x0800
	IXANY                    = 0x1000
	IXOFF                    = 0x0100
	IXON                     = 0x0200
	NCCS                     = 11
	NL0                      = 0x0000
	NL1                      = 0x0100
	NLDLY                    = 0x0100
	NOFLSH                   = 0x80000000
	OCRNL                    = 0x0008
	OFDEL                    = 0x0080
	OFILL                    = 0x0040
	OLCUC                    = 0x0002
	ONLCR                    = 0x0004
	ONLRET                   = 0x0020
	ONOCR                    = 0x0010
	OPOST                    = 0x0001
	PACKET                   = 0x0800
	PARENB                   = 0x0200
	PARMRK                   = 0x0400
	PARODD                   = 0x0400
	PATH_MAX                 = 1023
	PKT3270                  = 0x1000
	PKTXTND                  = 0x4000
	PTU3270                  = 0x2000
	SIOCGFDPOEATTRS          = 0x4000D306
	SIOCGPARTNERINFO         = 0xC000F612
	SIOCGSOCKMLSINFO         = 0x4000D307
	SIOCGSOCKPOEATTRS        = 0x4000D305
	TAB0                     = 0x0000
	TAB1                     = 0x0400
	TAB2                     = 0x0800
	TAB3                     = 0x0C00
	TABDLY                   = 0x0C00
	TCIFLUSH                 = 0
	TCIOFF                   = 2
	TCIOFLUSH                = 2
	TCION                    = 3
	TCOFLUSH                 = 1
	TCOOFF                   = 0
	TCOON                    = 1
	TCSADRAIN                = 1
	TCSAFLUSH                = 2
	TCSANOW                  = 0
	TOSTOP                   = 0x00400000
	VEOF                     = 4
	VEOL                     = 5
	VERASE                   = 2
	VINTR                    = 0
	VKILL                    = 3
	VMIN                     = 6
	VQUIT                    = 1
	VSTART                   = 7
	VSTOP                    = 8
	VSUSP                    = 9
	VT0                      = 0x00000000
	VT1                      = 0x00010000
	VTDLY                    = 0x00010000
	VTIME                    = 10
	XCASE                    = 0x0080
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
	X__termios               = 1
	X__types                 = 1
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

type Pid_t = X__pid_t

type Cc_t = uint8
type Speed_t = uint32
type Tcflag_t = uint32

type Termios = struct {
	Fc_iflag Tcflag_t
	Fc_oflag Tcflag_t
	Fc_cflag Tcflag_t
	Fc_lflag Tcflag_t
	Fc_cc    [11]Cc_t
} /* termios.h:345 */

var _ uint8
