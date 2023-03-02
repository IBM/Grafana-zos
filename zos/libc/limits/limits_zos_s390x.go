// Code created manually. Would like to be generated in the future.

package limits

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
	ATEXIT_MAX                           = 2048
	BC_BASE_MAX                          = SHRT_MAX
	BC_DIM_MAX                           = (SHRT_MAX + 1)
	BC_SCALE_MAX                         = SHRT_MAX
	BC_STRING_MAX                        = LINE_MAX
	CHARCLASS_NAME_MAX                   = 14
	CHAR_BIT                             = 8
	CHAR_MAX                             = UCHAR_MAX
	CHAR_MIN                             = 0
	COLL_WEIGHTS_MAX                     = X_POSIX2_COLL_WEIGHTS_MAX
	DBL_DIG                              = 15
	DBL_MAX                              = (float64)(1.79769313486231570814527423731704356e+308)
	EXPR_NEST_MAX                        = X_POSIX2_EXPR_NEST_MAX
	FLT_DIG                              = 6
	FLT_MAX                              = (float32)(3.40282346638528859811704183484516925e+38)
	INT_MAX                              = 2147483647
	INT_MIN                              = (-2147483647 - 1)
	IOV_MAX                              = 120
	LINE_MAX                             = X_POSIX2_LINE_MAX
	LLONG_MAX                            = (int64)(9223372036854775807)
	LLONG_MIN                            = (-LLONG_MAX - 1)
	LOGIN_NAME_MAX                       = X_POSIX_LOGIN_NAME_MAX
	LONGLONG_MAX                         = (int64)(9223372036854775807)
	LONGLONG_MIN                         = (-LONGLONG_MAX - 1)
	LONG_BIT                             = 64
	LONG_MAX                             = (int64)(9223372036854775807)
	LONG_MIN                             = (-LONG_MAX - 1)
	MB_LEN_MAX                           = 4
	NGROUPS_MAX                          = 300
	NL_ARGMAX                            = 9
	NL_LANGMAX                           = 14
	NL_MSGMAX                            = 32767
	NL_NMAX                              = 10
	NL_SETMAX                            = 2048
	NL_TEXTMAX                           = 8192
	NZERO                                = 20
	PAGESIZE                             = 4096
	PAGE_SIZE                            = 4096
	PASS_MAX                             = 255
	PATH_MAX                             = 1023
	POSIX_SYMLOOP                        = 24
	PTHREAD_STACK_MIN                    = 1048576
	RE_DUP_MAX                           = X_POSIX2_RE_DUP_MAX
	SCHAR_MAX                            = 127
	SCHAR_MIN                            = (-128)
	SHRT_MAX                             = 32767
	SHRT_MIN                             = (-32768)
	SSIZE_MAX                            = LONG_MAX
	TMP_MAX                              = 10000
	TTY_NAME_MAX                         = X_POSIX_TTY_NAME_MAX
	UCHAR_MAX                            = 255
	UINT_MAX                             = (uint32)(4294967295)
	ULLONG_MAX                           = (uint64)(18446744073709551615)
	ULONGLONG_MAX                        = (uint64)(18446744073709551615)
	ULONG_MAX                            = (uint64)(18446744073709551615)
	USHRT_MAX                            = 65535
	WORD_BIT                             = 32
	X_POSIX2_BC_BASE_MAX                 = 99
	X_POSIX2_BC_DIM_MAX                  = 2048
	X_POSIX2_BC_SCALE_MAX                = 99
	X_POSIX2_BC_STRING_MAX               = 1000
	X_POSIX2_CHARCLASS_NAME_MAX          = 14
	X_POSIX2_COLL_WEIGHTS_MAX            = 4
	X_POSIX2_COLL_WEIGHTS_MAX_A          = 4
	X_POSIX2_COLL_WEIGHTS_MAX_E          = 2
	X_POSIX2_EXPR_NEST_MAX               = 32
	X_POSIX2_LINE_MAX                    = 2048
	X_POSIX2_RE_DUP_MAX                  = 255
	X_POSIX_ARG_MAX                      = 4096
	X_POSIX_CHILD_MAX                    = 25
	X_POSIX_HOST_NAME_MAX                = 255
	X_POSIX_LINK_MAX                     = 8
	X_POSIX_LOGIN_NAME_MAX               = 9
	X_POSIX_MAX_CANON                    = 255
	X_POSIX_MAX_INPUT                    = 255
	X_POSIX_NAME_MAX                     = 14
	X_POSIX_NGROUPS_MAX                  = 8
	X_POSIX_OPEN_MAX                     = 20
	X_POSIX_PATH_MAX                     = 256
	X_POSIX_PIPE_BUF                     = 512
	X_POSIX_RE_DUP_MAX                   = X_POSIX2_RE_DUP_MAX
	X_POSIX_SSIZE_MAX                    = 32767
	X_POSIX_STREAM_MAX                   = 8
	X_POSIX_SYMLINK_MAX                  = 255
	X_POSIX_SYMLOOP_MAX                  = 8
	X_POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
	X_POSIX_THREAD_KEYS_MAX              = 128
	X_POSIX_THREAD_THREADS_MAX           = 64
	X_POSIX_TTY_NAME_MAX                 = 9
	X_POSIX_TZNAME_MAX                   = 6
	X_AE_BIMODAL                         = 1
	X_ALL_SOURCE                         = 1
	X_BIG_ENDIAN                         = 1
	X_CHAR_UNSIGNED                      = 1
	X_DBL_DIG_I                          = 6
	X_ENHANCED_ASCII_EXT                 = 0x410A0000
	X_EXT                                = 1
	X_FLT_DIG_I                          = 5
	X_ISOC99_SOURCE                      = 1
	X_LARGE_TIME_API                     = 1
	X_LONG_LONG                          = 1
	X_LP64                               = 1
	X_MI_BUILTIN                         = 1
	X_OPEN_DEFAULT                       = 1
	X_OPEN_MSGQ_EXT                      = 1
	X_OPEN_SYS                           = 1
	X_OPEN_SYS_FILE_EXT                  = 1
	X_OPEN_SYS_SOCK_IPV6                 = 1
	X_UNIX03_SOURCE                      = 1
	X_UNIX03_THREADS                     = 1
	X_UNIX03_WITHDRAWN                   = 1
	X_XOPEN_CRYPT                        = 1
	X_XOPEN_IOV_MAX                      = 16
	X_XOPEN_LEGACY                       = 1
	X_XOPEN_NAME_MAX                     = 255
	X_XOPEN_PATH_MAX                     = 1024
	X_XOPEN_REALTIME                     = -1
	X_XOPEN_REALTIME_THREADS             = -1
	X_XOPEN_SOURCE                       = 600
	X_XOPEN_SOURCE_EXTENDED              = 1
	X__limits                            = 1
	X__features_h                        = 1
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

var _ uint8
