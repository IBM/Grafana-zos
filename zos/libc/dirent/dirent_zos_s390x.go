// Code created manually. Would like to be generated in the future.

package dirent

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
	AUDTEXECFAIL               = 0x00000200
	AUDTEXECSUCC               = 0x00000100
	AUDTREADFAIL               = 0x02000000
	AUDTREADSUCC               = 0x01000000
	AUDTWRITEFAIL              = 0x00020000
	AUDTWRITESUCC              = 0x00010000
	AUDT_AUDITOR               = 0x01
	AUDT_USER                  = 0x00
	FT_BINARY                  = 0xFFFF
	FT_UNTAGGED                = 0x0000
	MNT2_SYSLIST_EXCLUDE       = 0x0001
	MNT2_SYSLIST_INCLUDE       = 0x0000
	MNTE2H_ID                  = "\xD4\xD5\xE3\xF2"
	MNTE3H_ID                  = "\xD4\xD5\xE3\xF3"
	MNT_FILE_ACTIVE            = 0x00
	MNT_FILE_ASYNCH_MOUNT      = 0x82
	MNT_FILE_DEAD              = 0x01
	MNT_FILE_DRAIN             = 0x04
	MNT_FILE_FORCE             = 0x08
	MNT_FILE_IMMED             = 0x10
	MNT_FILE_IMMED_TRIED       = 0x40
	MNT_FILE_MOUNT_IN_PROGRESS = 0x81
	MNT_FILE_NORM              = 0x20
	MNT_FILE_QUIESCED          = 0x80
	MNT_FILE_RECYCLEMOUNTING   = 0x10
	MNT_FILE_RECYCLENOTACTIVE  = 0x20
	MNT_FILE_RECYCLESTARTED    = 0x08
	MNT_FILE_RESET             = 0x02
	MNT_FSTYPE_MVS             = 1
	MNT_FSTYPE_NFS             = 2
	MNT_FSTYPE_PIPE            = 3
	MNT_FSTYPE_SOCKET          = 4
	MNT_MODE_EXPORT            = 0x00000004
	MNT_MODE_NOSECURITY        = 0x00000008
	MNT_MODE_NOSUID            = 0x00000002
	MNT_MODE_RDONLY            = 0x00000001
	MNT_MODE_RDWR              = 0x00000000
	MTM_DRAIN                  = 0x02000000
	MTM_FORCE                  = 0x04000000
	MTM_IMMED                  = 0x08000000
	MTM_NOSECURITY             = 0x00000080
	MTM_NOSUID                 = 0x00000400
	MTM_RDONLY                 = 0x80000000
	MTM_RDWR                   = 0x40000000
	MTM_REMOUNT                = 0x00000100
	MTM_RESET                  = 0x01000000
	MTM_SAMEMODE               = 0x00100000
	MTM_SYNCHONLY              = 0x00000200
	MTM_UMOUNT                 = 0x10000000
	MTM_UNQSEFORCE             = 0x00040000
	NULL                       = (uintptr)(0)
	PATH_MAX                   = 1023
	S_ACCESSACL                = 0x80
	S_DMODELACL                = 0x20
	S_FFBINARY                 = 1
	S_FFCR                     = 3
	S_FFCRLF                   = 5
	S_FFCRNL                   = 7
	S_FFLF                     = 4
	S_FFLFCR                   = 6
	S_FFNA                     = 0
	S_FFNL                     = 2
	S_FFRECORD                 = 8
	S_FMODELACL                = 0x40
	S_IEXEC                    = S_IXUSR
	S_IFAPFCTL                 = 0x00000004
	S_IFBLK                    = 0x06000000
	S_IFCHR                    = 0x02000000
	S_IFDIR                    = 0x01000000
	S_IFEXTL                   = 0x00000001
	S_IFFIFO                   = 0x04000000
	S_IFIFO                    = 0x04000000
	S_IFLNK                    = 0x05000000
	S_IFMST                    = 0x00FF0000
	S_IFMT                     = 0xFF000000
	S_IFNOSHARE                = 0x00000008
	S_IFPROGCTL                = 0x00000002
	S_IFREG                    = 0x03000000
	S_IFSHARELIB               = 0x00000010
	S_IFSOCK                   = 0x07000000
	S_IFVMEXTL                 = 0xFE000000
	S_IFVMEXTL_DATA            = 0x00020000
	S_IFVMEXTL_EXEC            = 0x00010000
	S_IFVMEXTL_MEL             = 0x00030000
	S_IREAD                    = S_IRUSR
	S_IRGRP                    = 0x0020
	S_IROTH                    = 0x0004
	S_IRUSR                    = 0x0100
	S_IRWXG                    = 0x0038
	S_IRWXO                    = 0x0007
	S_IRWXU                    = 0x01C0
	S_ISGID                    = 0x0400
	S_ISUID                    = 0x0800
	S_ISVTX                    = 0x0200
	S_IWGRP                    = 0x0010
	S_IWOTH                    = 0x0002
	S_IWRITE                   = S_IWUSR
	S_IWUSR                    = 0x0080
	S_IXGRP                    = 0x0008
	S_IXOTH                    = 0x0001
	S_IXUSR                    = 0x0040
	S_TYPEISMQ                 = (0)
	S_TYPEISSEM                = (0)
	S_TYPEISSHM                = (0)
	X_AE_BIMODAL               = 1
	X_ALL_SOURCE               = 1
	X_BIG_ENDIAN               = 1
	X_CHAR_UNSIGNED            = 1
	X_ENHANCED_ASCII_EXT       = 0x410A0000
	X_EXT                      = 1
	X_ISOC99_SOURCE            = 1
	X_LARGE_TIME_API           = 1
	X_LONG_LONG                = 1
	X_LP64                     = 1
	X_MI_BUILTIN               = 1
	X_OPEN_DEFAULT             = 1
	X_OPEN_MSGQ_EXT            = 1
	X_OPEN_SYS                 = 1
	X_OPEN_SYS_FILE_EXT        = 1
	X_OPEN_SYS_SOCK_IPV6       = 1
	X_UNIX03_SOURCE            = 1
	X_UNIX03_THREADS           = 1
	X_UNIX03_WITHDRAWN         = 1
	X_XOPEN_CRYPT              = 1
	X_XOPEN_LEGACY             = 1
	X_XOPEN_REALTIME           = -1
	X_XOPEN_REALTIME_THREADS   = -1
	X_XOPEN_SOURCE             = 600
	X_XOPEN_SOURCE_EXTENDED    = 1
	X__dirent                  = 1
	X__features_h              = 1
	X__stat                    = 1
	X__types                   = 1
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

type X__didesc = struct {
	Fdd_opaque1 [40]uint8
	Fdd_fd      int32
} /* dirent.h:61 */

// TODO (ZOS)
type Dirent = struct {
	Fd_ino    X__ino_t
	Fd_reclen uint16
	Fd_namlen uint16
	Fd_name   [256]uint8
	Fd_extra  uintptr
} /* dirent.h:73 */

var _ uint8
