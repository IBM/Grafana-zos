// Code created manually. Would like to be generated in the future.
package stdio

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
	BUFSIZ                               = 512
	CHARCLASS_NAME_MAX                   = 14
	CHAR_BIT                             = 8
	CHAR_MAX                             = UCHAR_MAX
	CHAR_MIN                             = 0
	COLL_WEIGHTS_MAX                     = X_POSIX2_COLL_WEIGHTS_MAX
	DBL_DIG                              = 15
	DBL_MAX                              = 1.79769313486231570814527423731704356e+308
	EOF                                  = (-1)
	EXPR_NEST_MAX                        = X_POSIX2_EXPR_NEST_MAX
	FILENAME_MAX                         = 1024
	FLT_DIG                              = 6
	FLT_MAX                              = 3.40282346638528859811704183484516925e+38
	FOPEN_MAX                            = 64
	INT_MAX                              = 2147483647
	INT_MIN                              = (-2147483647 - 1)
	IOV_MAX                              = 120
	LC_ALL                               = (-1)
	LC_COLLATE                           = 0
	LC_CTYPE                             = 1
	LC_MESSAGES                          = 5
	LC_MONETARY                          = 2
	LC_NUMERIC                           = 3
	LC_SYNTAX                            = 7
	LC_TIME                              = 4
	LC_TOD                               = 6
	LINE_MAX                             = X_POSIX2_LINE_MAX
	LLONG_MAX                            = (9223372036854775807)
	LLONG_MIN                            = (-LLONG_MAX - 1)
	LOGIN_NAME_MAX                       = X_POSIX_LOGIN_NAME_MAX
	LONGLONG_MAX                         = (9223372036854775807)
	LONGLONG_MIN                         = (-LONGLONG_MAX - 1)
	LONG_BIT                             = 64
	LONG_MAX                             = (9223372036854775807)
	LONG_MIN                             = (-LONG_MAX - 1)
	L_ctermid                            = 1024
	L_cuserid                            = 9
	L_tmpnam                             = 1024
	MB_LEN_MAX                           = 4
	NGROUPS_MAX                          = 300
	NL_ARGMAX                            = 9
	NL_LANGMAX                           = 14
	NL_MSGMAX                            = 32767
	NL_NMAX                              = 10
	NL_SETMAX                            = 2048
	NL_TEXTMAX                           = 8192
	NULL                                 = (uintptr)(0)
	NZERO                                = 20
	PAGESIZE                             = 4096
	PAGE_SIZE                            = 4096
	PASS_MAX                             = 255
	PATH_MAX                             = 1023
	POSIX_SYMLOOP                        = 24
	PTHREAD_STACK_MIN                    = 1048576
	P_tmpdir                             = "/tmp"
	RE_DUP_MAX                           = X_POSIX2_RE_DUP_MAX
	SCHAR_MAX                            = 127
	SCHAR_MIN                            = (-128)
	SEEK_CUR                             = 1
	SEEK_END                             = 2
	SEEK_SET                             = 0
	SHRT_MAX                             = 32767
	SHRT_MIN                             = (-32768)
	SSIZE_MAX                            = LONG_MAX
	TMP_MAX                              = 10000
	TTY_NAME_MAX                         = X_POSIX_TTY_NAME_MAX
	UCHAR_MAX                            = 255
	UINT_MAX                             = 4294967295
	ULLONG_MAX                           = (18446744073709551615)
	ULONGLONG_MAX                        = (18446744073709551615)
	ULONG_MAX                            = (18446744073709551615)
	USHRT_MAX                            = 65535
	WORD_BIT                             = 32
	X_COLL_BACKWARD_MASK                 = 2
	X_COLL_FORWARD_MASK                  = 1
	X_COLL_NOSUBS_MASK                   = 4
	X_COLL_POSITION_MASK                 = 8
	X_COLL_WEIGHTS_MAX                   = COLL_WEIGHTS_MAX
	X_COLL_WEIGHTS_MAX_A                 = 4
	X_COLL_WEIGHTS_MAX_E                 = 6
	X_DFLT_LOC_PATH                      = "/usr/lib/nls/loc/"
	X_ENHANCED_ASCII_EXT                 = 0x410A0000
	X_IOFBF                              = 1
	X_IOLBF                              = 2
	X_IONBF                              = 3
	X_ISALNUM_A                          = 0x002
	X_ISALPHA_A                          = 0x001
	X_ISBLANK_A                          = 0x004
	X_ISCNTRL_A                          = 0x008
	X_ISDIGIT_A                          = 0x010
	X_ISGRAPH_A                          = 0x020
	X_ISLOWER_A                          = 0x040
	X_ISOC99_SOURCE                      = 1
	X_ISPRINT_A                          = 0x080
	X_ISPUNCT_A                          = 0x100
	X_ISSPACE_A                          = 0x200
	X_ISUPPER_A                          = 0x400
	X_ISXDIGIT_A                         = 0x800
	X_LARGE_TIME_API                     = 1
	X_LC_ASCII_TYPE                      = 256
	X_LC_CAR                             = 1
	X_LC_CAR_A                           = 1 + X_LC_ASCII_TYPE
	X_LC_CHARMAP                         = 3
	X_LC_CHARMAP_A                       = 3 + X_LC_ASCII_TYPE
	X_LC_COLLATE                         = 5
	X_LC_COLLATE_A                       = 5 + X_LC_ASCII_TYPE
	X_LC_CTYPE                           = 4
	X_LC_CTYPE_A                         = 4 + X_LC_ASCII_TYPE
	X_LC_LOCALE                          = 2
	X_LC_LOCALE_A                        = 2 + X_LC_ASCII_TYPE
	X_LC_MAGIC                           = 0xD3C3
	X_LC_MAGIC_A                         = 0xD3C3
	X_LC_MAX_OBJECTS                     = 256
	X_LC_MONETARY                        = 7
	X_LC_MONETARY_A                      = 7 + X_LC_ASCII_TYPE
	X_LC_NUMERIC                         = 6
	X_LC_NUMERIC_A                       = 6 + X_LC_ASCII_TYPE
	X_LC_OBJ_HDL                         = 12 + X_LC_ASCII_TYPE
	X_LC_RESP                            = 9
	X_LC_RESP_A                          = 9 + X_LC_ASCII_TYPE
	X_LC_STANDARD_TYPE                   = 1000000
	X_LC_SUSV3                           = 200201
	X_LC_SYNTAX                          = 10
	X_LC_SYNTAX_A                        = 10 + X_LC_ASCII_TYPE
	X_LC_TIME                            = 8
	X_LC_TIME_A                          = 8 + X_LC_ASCII_TYPE
	X_LC_TOD                             = 11
	X_LC_TOD_A                           = 11 + X_LC_ASCII_TYPE
	X_LC_VERSION                         = 0x00011000
	X_LC_VERSION_A                       = 0x40000000
	X_LC_VERSION_SUSV3                   = (X_LC_STANDARD_TYPE + X_LC_SUSV3)
	X_LC_VERSION_SUSV3_A                 = (X_LC_VERSION_A + X_LC_SUSV3)
	X_LC_VERSION_XPG4_VERSION2           = 1199409
	X_LC_VERSION_XPG4_VERSION2C          = 1199506
	X_LOC_HAS_MCCE                       = 1
	X_LONG_LONG                          = 1
	X_LP64                               = 1
	X_MI_BUILTIN                         = 1
	X_NL_NUM_ITEMS                       = 58
	X_NL_NUM_ITEMS1                      = 51
	X_NL_NUM_ITEMS2                      = 7
	X_NL_NUM_ITEMS_A                     = 58
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
	X__stdio                             = 1
	X__features_h                        = 1
	X__types                             = 1
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

type X__valist = [2]uintptr

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

type X__mbstate_t = uint16

type X__fpos_t = struct {
	F__fpos_elem [8]int64
} /* stdio.h:103 */

// (ZOS) Adjusted
type X__file = struct {
	F__bufPtr          uintptr
	F__countIn         int64
	F__countOut        int64
	F__fcbgetc         uintptr
	F__fcbputc         uintptr
	F__cntlinterpret   uint32
	F__fcb_ascii       uint32
	F__fcb_orientation uint32
} /* stdio.h:76 */

type X__ffile = struct {
	F__fp uintptr
}

type X__FILEP = uintptr

// TODO (ZOS) - this type is used on other OSs, however it doesn't exist on zos
type X_IO_FILE = struct {
	F_flags          int32
	F__ccgo_pad1     [4]byte
	F_IO_read_ptr    uintptr
	F_IO_read_end    uintptr
	F_IO_read_base   uintptr
	F_IO_write_base  uintptr
	F_IO_write_ptr   uintptr
	F_IO_write_end   uintptr
	F_IO_buf_base    uintptr
	F_IO_buf_end     uintptr
	F_IO_save_base   uintptr
	F_IO_backup_base uintptr
	F_IO_save_end    uintptr
	F_markers        uintptr
	F_chain          uintptr
	F_fileno         int32
	F_flags2         int32
	F_old_offset     X__off_t
	F_cur_column     uint16
	F_vtable_offset  int8
	F_shortbuf       [1]uint8
	F__ccgo_pad2     [4]byte
	F_lock           uintptr
	F_offset         X__off64_t
	F_codecvt        uintptr
	F_wide_data      uintptr
	F_freeres_list   uintptr
	F_freeres_buf    uintptr
	F__pad5          Size_t
	F_mode           int32
	F_unused2        [20]uint8
} /* __FILE.h:4:1 */

type X__FILE = X_IO_FILE /* __FILE.h:5:25 */

type FILE = X__ffile /* stdio.h:99 */

type Fpos_t = X__fpos_t /* stdio.h:107 */

type Va_list = X__valist

type Off_t = X__off64_t

type Ssize_t = X__ssize_t

var _ uint8
