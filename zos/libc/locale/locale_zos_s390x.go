// Code created manually. Would like to be generated in the future.

package locale

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
	LC_ALL                   = (-1)
	LC_COLLATE               = 0
	LC_CTYPE                 = 1
	LC_MESSAGES              = 5
	LC_MONETARY              = 2
	LC_NUMERIC               = 3
	LC_SYNTAX                = 7
	LC_TIME                  = 4
	LC_TOD                   = 6
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
	X__locale                = 1
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

// Structure giving information about numeric and monetary notation.
type X__lconv_a = struct { /* ASCII category */
	Fdecimal_point      uintptr
	Fthousands_sep      uintptr
	Fgrouping           uintptr
	Fint_curr_symbol    uintptr
	Fcurrency_symbol    uintptr
	Fmon_decimal_point  uintptr
	Fmon_thousands_sep  uintptr
	Fmon_grouping       uintptr
	Fpositive_sign      uintptr
	Fnegative_sign      uintptr
	Fint_frac_digits    uint8
	Ffrac_digits        uint8
	Fp_cs_precedes      uint8
	Fp_sep_by_space     uint8
	Fn_cs_precedes      uint8
	Fn_sep_by_space     uint8
	Fp_sign_posn        uint8
	Fn_sign_posn        uint8
	Fint_p_cs_precedes  uint8
	Fint_p_sep_by_space uint8
	Fint_n_cs_precedes  uint8
	Fint_n_sep_by_space uint8
	Fint_p_sign_posn    uint8
	Fint_n_sign_posn    uint8
	Fleft_paranthesis   uint8
	Fright_paranthesis  uint8
} /* locale.h:67 */

type X__lconv_e = struct { /* EBCDIC category */
	Fdecimal_point       uintptr
	Fthousands_sep       uintptr
	Fgrouping            uintptr
	Fint_curr_symbol     uintptr
	Fcurrency_symbol     uintptr
	Fmon_decimal_point   uintptr
	Fmon_thousands_sep   uintptr
	Fmon_grouping        uintptr
	Fpositive_sign       uintptr
	Fnegative_sign       uintptr
	Fint_frac_digits     uint8
	Ffrac_digits         uint8
	Fp_cs_precedes       uint8
	Fp_sep_by_space      uint8
	Fn_cs_precedes       uint8
	Fn_sep_by_space      uint8
	Fp_sign_posn         uint8
	Fn_sign_posn         uint8
	Fint_p_cs_precedes   uint8
	Fint_p_sep_by_space  uint8
	Fint_n_cs_precedes   uint8
	Fint_n_sep_by_space  uint8
	Fint_p_sign_posn     uint8
	Fint_n_sign_posn     uint8
	F__left_paranthesis  uintptr
	F__right_paranthesis uintptr
	F__debit_sign        uintptr
	F__credit_sign       uintptr
	Fint_p_cs_precedes   uint8
	Fint_p_sep_by_space  uint8
	Fint_n_cs_precedes   uint8
	Fint_n_sep_by_space  uint8
	Fint_p_sign_posn     uint8
	Fint_n_sign_posn     uint8
} /* locale.h:118 */

type Lconv = struct {
	Fdecimal_point      uintptr
	Fthousands_sep      uintptr
	Fgrouping           uintptr
	Fint_curr_symbol    uintptr
	Fcurrency_symbol    uintptr
	Fmon_decimal_point  uintptr
	Fmon_thousands_sep  uintptr
	Fmon_grouping       uintptr
	Fpositive_sign      uintptr
	Fnegative_sign      uintptr
	Fint_frac_digits    uint8
	Ffrac_digits        uint8
	Fp_cs_precedes      uint8
	Fp_sep_by_space     uint8
	Fn_cs_precedes      uint8
	Fn_sep_by_space     uint8
	Fp_sign_posn        uint8
	Fn_sign_posn        uint8
	Fint_p_cs_precedes  uint8
	Fint_p_sep_by_space uint8
	Fint_n_cs_precedes  uint8
	Fint_n_sep_by_space uint8
	Fint_p_sign_posn    uint8
	Fint_n_sign_posn    uint8
	Fleft_paranthesis   uintptr
	Fright_paranthesis  uintptr
	Fint_p_cs_precedes  uint8
	Fint_p_sep_by_space uint8
	Fint_n_cs_precedes  uint8
	Fint_n_sep_by_space uint8
	Fint_p_sign_posn    uint8
	Fint_n_sign_posn    uint8
} /* locale.h:189 */

var _ uint8
