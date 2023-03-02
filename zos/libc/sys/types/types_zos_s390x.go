// Code created manually. Would like to be generated in the future.
package types

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
	X__ORDER_BIG_ENDIAN__    = 4321
	X__ORDER_PDP_ENDIAN__    = 3412
	X__ORDER_LITTLE_ENDIAN__ = 1234
	X__features_h            = 1
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

type U_int8_t = X__uint8_t
type U_int16_t = X__uint16_t
type U_int32_t = X__uint32_t
type U_int64_t = X__uint64_t

// Adjusted for zos
type X__sigset_t = struct {
	F__sigs_0 uint32
	F__sigs_1 uint32
} /* sys/types.h:200 */

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

// Duplicates info from stdint.h but this is used in unistd.h.
type X__intptr_t = int64

// Duplicate info from sys/socket.h.
type X__socklen_t = uint32 /* socket.h:115 */

type X__sig_atomic_t = int32 /* signal.h:52 */

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
