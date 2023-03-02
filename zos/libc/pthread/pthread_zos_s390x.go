// Code created manually. Would like to be generated in the future.
package pthread

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
	CLOCKS_PER_SEC = 1000000

	SCHED_OTHER = 1
	SCHED_FIFO  = 2
	SCHED_RR    = 3

	NULL                     = (uintptr)(0)
	PTHREAD_CANCELED         = 0xFFFFFFFFFFFFFFFF
	PTHREAD_COND_INITIALIZER = [1]uint64{0x00000000FBC3C9D6}

	PTHREAD_CREATE_DETACHED = 1
	PTHREAD_CREATE_JOINABLE = 0

	// Scheduler inheritance.
	PTHREAD_EXPLICIT_SCHED = 0
	PTHREAD_INHERIT_SCHED  = 1

	// Mutex types
	PTHREAD_MUTEX_DEFAULT     = 0
	PTHREAD_MUTEX_ERRORCHECK  = 0
	PTHREAD_MUTEX_INITIALIZER = [1]uint64{0x00000000FBD7C9D6}
	PTHREAD_MUTEX_NORMAL      = 4
	PTHREAD_MUTEX_RECURSIVE   = 1

	PTHREAD_NONVALID_ID_INITIALIZER_NP = [8]uint8{0x00, 0x00, 0x00, 0x01, 0xEE, 0xEE, 0xEE, 0xEE}
	PTHREAD_ONCE_INIT                  = 0

	PTHREAD_PROCESS_PRIVATE = 0
	PTHREAD_PROCESS_SHARED  = 8

	PTHREAD_RWLOCK_INITIALIZER_NP = [8]uint8{0, 0, 0, 0, 0xFB, 0xD7, 0xC9, 0xD6}

	// Cancellation
	PTHREAD_CANCEL_ENABLE  = 0
	PTHREAD_CANCEL_DISABLE = 1

	PTHREAD_CANCEL_DEFERRED     = 0
	PTHREAD_CANCEL_ASYNCHRONOUS = 1

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
	X__pthread               = 1
	X__features_h            = 1
	X__sched_h               = 1
	X__time                  = 1
	X__types                 = 1
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

type Clock_t = X__clock_t

type Time_t = X__time_t

// Adjusted for zos
type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec int64
} /* time.h:64 */

type Pid_t = X__pid_t

// Data structure to describe a process' schedulability.
type Sched_param = struct{ Fsched_priority int32 }

// ISO C `broken-down time' structure.
type Tm = struct {
	Ftm_sec   int32
	Ftm_min   int32
	Ftm_hour  int32
	Ftm_mday  int32
	Ftm_mon   int32
	Ftm_year  int32
	Ftm_wday  int32
	Ftm_yday  int32
	Ftm_isdst int32
}

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

var _ uint8
