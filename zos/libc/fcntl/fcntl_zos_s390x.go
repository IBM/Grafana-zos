// Code created manually. Would like to be generated in the future.

package fcntl

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
	FD_CLOEXEC               = 0x01
	FD_CLOFORK               = 0x02
	FNDELAY                  = 0x04
	F_CLOSFD                 = 9
	F_CONTROL_CVT            = 13
	F_DUPFD                  = 0
	F_DUPFD2                 = 8
	F_GETFD                  = 1
	F_GETFL                  = 259
	F_GETLK                  = 5
	F_GETOWN                 = 10
	F_RDLCK                  = 1
	F_SETFD                  = 2
	F_SETFL                  = 4
	F_SETLK                  = 6
	F_SETLKW                 = 7
	F_SETOWN                 = 11
	F_SETTAG                 = 12
	F_UNLCK                  = 3
	F_WRLCK                  = 2
	O_ACCMODE                = 0x03
	O_APPEND                 = 0x08
	O_ASYNCSIG               = 0x0200
	O_CREAT                  = 0x80
	O_EXCL                   = 0x40
	O_GETFL                  = 0x0F
	O_LARGEFILE              = 0x0400
	O_NDELAY                 = 0x04
	O_NOCTTY                 = 0x20
	O_NONBLOCK               = 0x04
	O_RDONLY                 = 0x02
	O_RDWR                   = 0x03
	O_SYNC                   = 0x0100
	O_TRUNC                  = 0x10
	O_WRONLY                 = 0x01
	PATH_MAX                 = 1023
	QUERYCVT                 = 3
	SEEK_CUR                 = 1
	SEEK_END                 = 2
	SEEK_SET                 = 0
	SETAUTOCVTALL            = 5
	SETAUTOCVTON             = 2
	SETCVTALL                = 4
	SETCVTOFF                = 0
	SETCVTON                 = 1
	S_IEXEC                  = 0x0040
	S_IFAPFCTL               = 0x00000004
	S_IFBLK                  = 0x06000000
	S_IFCHR                  = 0x02000000
	S_IFDIR                  = 0x01000000
	S_IFEXTL                 = 0x00000001
	S_IFFIFO                 = 0x04000000
	S_IFIFO                  = 0x04000000
	S_IFLNK                  = 0x05000000
	S_IFMST                  = 0x00FF0000
	S_IFMT                   = 0xFF000000
	S_IFNOSHARE              = 0x00000008
	S_IFPROGCTL              = 0x00000002
	S_IFREG                  = 0x03000000
	S_IFSHARELIB             = 0x00000010
	S_IFSOCK                 = 0x07000000
	S_IFVMEXTL               = 0xFE000000
	S_IFVMEXTL_DATA          = 0x00020000
	S_IFVMEXTL_EXEC          = 0x00010000
	S_IFVMEXTL_MEL           = 0x00030000
	S_IREAD                  = 0x0100
	S_IRGRP                  = 0x0020
	S_IROTH                  = 0x0004
	S_IRUSR                  = 0x0100
	S_IRWXG                  = 0x0038
	S_IRWXO                  = 0x0007
	S_IRWXU                  = 0x01C0
	S_ISVTX                  = 0x0200
	S_IWGRP                  = 0x0010
	S_IWOTH                  = 0x0002
	S_IWRITE                 = 0x0080
	S_IWUSR                  = 0x0080
	S_IXGRP                  = 0x0008
	S_IXOTH                  = 0x0001
	S_IXUSR                  = 0x0040
	S_TYPEISMQ               = 0
	S_TYPEISSEM              = 0
	S_TYPEISSHM              = 0
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
	X__fcntl                 = 1
	X__features_h            = 1
	X__modes                 = 1
	X__types                 = 1
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

type Flock = struct {
	Fl_type   int16
	Fl_whence int16
	Fl_start  X__off_t
	Fl_len    X__off_t
	Fl_pid    X__pid_t
} /* fcntl.h:158 */

type Mode_t = X__mode_t
type Off_t = X__off64_t
type Pid_t = X__pid_t

// Adjusted for zos. Does not appear in fcntl.h
type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec int64
}

type File_Tag = struct {
	Fft_ccsid    uint16
	Fft_txtflag  uint32
	Fft_deferred uint32
	Fft_rsvflags uint32
} /* sys/stat.h:83 */

type Stat = struct {
	Fst_eye     [4]uint8   // Constant eye catcher
	Fst_length  uint16     // Length of this struct
	Fst_version uint16     // Version number
	Fst_mode    X__mode_t  // File mode - see <modes.h>
	Fst_ino     X__ino_t   // File serial number
	Fst_dev     X__dev_t   // File devide ID
	Fst_nlink   X__nlink_t // Number of links
	Fst_uid     X__uid_t   // User ID of the owner
	Fst_gid     X__gid_t   // Group ID of the group
	Fst_size    X__off_t   // Size of the file

	// #ifdef __LP64
	Fst_atime31 [4]uint8 // Time last access
	Fst_mtime31 [4]uint8 // Time last data modification
	Fst_ctime31 [4]uint8 // Time last file status change

	// #ifdef __XPG4
	Fst_rdev X__dev_t // Major and minor numbers for character special file

	Fst_auditoraudit uint32 // Auditor audit info
	Fst_useraudit    uint32 // User audit info

	// #if (defined(_LP64) || defined(__U98) || defined(__SUSV3_POSIX))
	Fst_blksize X__blksize_t // File block size

	// #ifdef __LP64
	Fst_createtime_31 [4]uint8 // File creation time

	Fst_auditid [16]uint8 // File id for auditing
	Fst_rsrvd1  [4]uint8  // Reserved for expansion

	// #ifdef __FX
	Fst_tag       File_Tag // File tag
	Fst_charsetid [8]uint8 // Obsolete code char set ID

	// #ifdef __LF
	Fst_blocks X__blkcnt_t // Number of I/O blocks

	Fst_genvalue uint32 // General attribute values

	// #ifdef __LP64
	Fst_reftime_31 [4]uint8

	Fst_fid     [4]uint8 // File id
	Fst_filefmt uint8    // File format

	// #if EDC_TARGET >= 0x41030000
	Fst_fspflag2 uint8 // ACL support
	Fst_rsrvd2   uint8 // Reserved for expansion

	// #ifdef __LP64
	Fst_rsrvd4     [4]uint8  // Alignment
	Fst_atime      X__time_t // Time last access
	Fst_mtime      X__time_t // Time last data modification
	Fst_ctime      X__time_t // Time last file status change
	Fst_createtime X__time_t // File creation time
	Fst_reftime    X__time_t // Reference time
	Fst_rsrvd5     [24]uint8
} /* sys/stat.h:340 */

var _ uint8
