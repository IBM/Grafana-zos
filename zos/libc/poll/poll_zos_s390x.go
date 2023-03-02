// Code created manually. Would like to be generated in the future.

package poll

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
	PATH_MAX                 = 1023
	POLLERR                  = 0x0020
	POLLHUP                  = 0x0040
	POLLIN                   = (POLLRDNORM | POLLRDBAND)
	POLLNVAL                 = 0x0080
	POLLOUT                  = POLLWRNORM
	POLLPRI                  = 0x0010
	POLLRDBAND               = 0x0002
	POLLRDNORM               = 0x0001
	POLLWRBAND               = 0x0008
	POLLWRNORM               = 0x0004
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
	X__poll                  = 1
	X__features_h            = 1
)

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

// Type used for the number of file descriptors.
type X__nfds_t = uint32 /* poll.h:68 */
type Nfds_t = X__nfds_t
type Nmsgsfds_t = X__nfds_t /* poll.h:69 */

// Data structure describing a polling request.
type Pollfd = struct {
	Ffd      int32
	Fevents  int16
	Frevents int16
} /* poll.h:71 */

type Pollmsg = struct {
	Fmsgid   int32
	Fevents  int16
	Frevents int16
}

var _ uint8
