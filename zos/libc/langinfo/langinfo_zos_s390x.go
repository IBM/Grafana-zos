// Code created manually. Would like to be generated in the future.

package langinfo

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
	X_EXT                    = 1
	X_ISOC99_SOURCE          = 1
	X_LARGE_TIME_API         = 1
	X_LONG_LONG              = 1
	X_LP64                   = 1
	X_AE_BIMODAL             = 1
	X_ALL_SOURCE             = 1
	X_BIG_ENDIAN             = 1
	X_CHAR_UNSIGNED          = 1
	X_ENHANCED_ASCII_EXT     = 0x410A0000
	X_UNIX03_SOURCE          = 1
	X_UNIX03_THREADS         = 1
	X_UNIX03_WITHDRAWN       = 1
	X_XOPEN_CRYPT            = 1
	X_XOPEN_LEGACY           = 1
	X_XOPEN_REALTIME         = -1
	X_XOPEN_REALTIME_THREADS = -1
	X_XOPEN_SOURCE           = 600
	X_XOPEN_SOURCE_EXTENDED  = 1
	X_OPEN_DEFAULT           = 1
	X_OPEN_MSGQ_EXT          = 1
	X_OPEN_SYS               = 1
	X_OPEN_SYS_FILE_EXT      = 1
	X_OPEN_SYS_SOCK_IPV6     = 1
	X__features_h            = 1
	X__nl_types              = 1
)

const (
	ABDAY_1 = X_ABDAY_1_A
	ABDAY_2 = X_ABDAY_2_A
	ABDAY_3 = X_ABDAY_3_A
	ABDAY_4 = X_ABDAY_4_A
	ABDAY_5 = X_ABDAY_5_A
	ABDAY_6 = X_ABDAY_6_A
	ABDAY_7 = X_ABDAY_7_A

	DAY_1 = X_DAY_1_A
	DAY_2 = X_DAY_2_A
	DAY_3 = X_DAY_3_A
	DAY_4 = X_DAY_4_A
	DAY_5 = X_DAY_5_A
	DAY_6 = X_DAY_6_A
	DAY_7 = X_DAY_7_A

	ABMON_1  = X_ABMON_1_A
	ABMON_2  = X_ABMON_2_A
	ABMON_3  = X_ABMON_3_A
	ABMON_4  = X_ABMON_4_A
	ABMON_5  = X_ABMON_5_A
	ABMON_6  = X_ABMON_6_A
	ABMON_7  = X_ABMON_7_A
	ABMON_8  = X_ABMON_8_A
	ABMON_9  = X_ABMON_9_A
	ABMON_10 = X_ABMON_10_A
	ABMON_11 = X_ABMON_11_A
	ABMON_12 = X_ABMON_12_A

	MON_1  = X_MON_1_A
	MON_2  = X_MON_2_A
	MON_3  = X_MON_3_A
	MON_4  = X_MON_4_A
	MON_5  = X_MON_5_A
	MON_6  = X_MON_6_A
	MON_7  = X_MON_7_A
	MON_8  = X_MON_8_A
	MON_9  = X_MON_9_A
	MON_10 = X_MON_10_A
	MON_11 = X_MON_11_A
	MON_12 = X_MON_12_A

	AM_STR = X_AM_STR_A
	PM_STR = X_PM_STR_A

	D_FMT      = X_D_FMT_A
	D_T_FMT    = X_D_T_FMT_A
	T_FMT      = X_T_FMT_A
	T_FMT_AMPM = X_T_FMT_AMPM_A

	ERA         = X_ERA_A
	ERA_D_FMT   = X_ERA_D_FMT_A
	ALT_DIGITS  = X_ALT_DIGITS_A
	ERA_D_T_FMT = X_ERA_D_T_FMT_A
	ERA_T_FMT   = X_ERA_T_FMT_A

	CODESET       = X_CODESET_A
	CRNCYSTR      = X_CRNCYSTR_A
	NL_CAT_LOCALE = 1
	NL_SETD       = 1
	NOEXPR        = X_NOEXPR_A
	NOSTR         = X_NOSTR_A
	RADIXCHAR     = X_RADIXCHAR_A
	THOUSEP       = X_THOUSEP_A
	YESEXPR       = X_YESEXPR_A
	YESSTR        = X_YESSTR_A

	X_ABDAY_1_A  = 6
	X_ABDAY_1_E  = 8
	X_ABDAY_2_A  = 7
	X_ABDAY_2_E  = 9
	X_ABDAY_3_A  = 8
	X_ABDAY_3_E  = 10
	X_ABDAY_4_A  = 9
	X_ABDAY_4_E  = 11
	X_ABDAY_5_A  = 10
	X_ABDAY_5_E  = 12
	X_ABDAY_6_A  = 11
	X_ABDAY_6_E  = 13
	X_ABDAY_7_A  = 12
	X_ABDAY_7_E  = 14
	X_ABMON_10_A = 29
	X_ABMON_10_E = 31
	X_ABMON_11_A = 30
	X_ABMON_11_E = 32
	X_ABMON_12_A = 31
	X_ABMON_12_E = 33
	X_ABMON_1_A  = 20
	X_ABMON_1_E  = 22
	X_ABMON_2_A  = 21
	X_ABMON_2_E  = 23
	X_ABMON_3_A  = 22
	X_ABMON_3_E  = 24
	X_ABMON_4_A  = 23
	X_ABMON_4_E  = 25
	X_ABMON_5_A  = 24
	X_ABMON_5_E  = 26
	X_ABMON_6_A  = 25
	X_ABMON_6_E  = 27
	X_ABMON_7_A  = 26
	X_ABMON_7_E  = 28
	X_ABMON_8_A  = 27
	X_ABMON_8_E  = 29
	X_ABMON_9_A  = 28
	X_ABMON_9_E  = 30

	X_ALT_DIGITS_A = 55
	X_ALT_DIGITS_E = 55
	X_AM_STR_A     = 4
	X_AM_STR_E     = 6
	X_CODESET_A    = 49
	X_CODESET_E    = 1
	X_CRNCYSTR_A   = 48
	X_CRNCYSTR_E   = 50

	X_DAY_1_A = 13
	X_DAY_1_E = 15
	X_DAY_2_A = 14
	X_DAY_2_E = 16
	X_DAY_3_A = 15
	X_DAY_3_E = 17
	X_DAY_4_A = 16
	X_DAY_4_E = 18
	X_DAY_5_A = 17
	X_DAY_5_E = 19
	X_DAY_6_A = 18
	X_DAY_6_E = 20
	X_DAY_7_A = 19
	X_DAY_7_E = 21

	X_D_FMT_A       = 2
	X_D_FMT_E       = 3
	X_D_T_FMT_A     = 1
	X_D_T_FMT_E     = 2
	X_ERA_A         = 51
	X_ERA_D_FMT_A   = 52
	X_ERA_D_FMT_E   = 52
	X_ERA_D_T_FMT_A = 53
	X_ERA_D_T_FMT_E = 53
	X_ERA_E         = 51
	X_ERA_T_FMT_A   = 54
	X_ERA_T_FMT_E   = 54

	X_MI_BUILTIN  = 1
	X_MON_10_A    = 41
	X_MON_10_E    = 43
	X_MON_11_A    = 42
	X_MON_11_E    = 44
	X_MON_12_A    = 43
	X_MON_12_E    = 45
	X_MON_1_A     = 32
	X_MON_1_E     = 34
	X_MON_2_A     = 33
	X_MON_2_E     = 35
	X_MON_3_A     = 34
	X_MON_3_E     = 36
	X_MON_4_A     = 35
	X_MON_4_E     = 37
	X_MON_5_A     = 36
	X_MON_5_E     = 38
	X_MON_6_A     = 37
	X_MON_6_E     = 39
	X_MON_7_A     = 38
	X_MON_7_E     = 40
	X_MON_8_A     = 39
	X_MON_8_E     = 41
	X_MON_9_A     = 40
	X_MON_9_E     = 42
	X_PM_STR_A    = 5
	X_PM_STR_E    = 7
	X_RADIXCHAR_A = 44
	X_RADIXCHAR_E = 46
	X_THOUSEP_A   = 45
	X_THOUSEP_E   = 47

	X_T_FMT_A      = 3
	X_T_FMT_AMPM_A = 50
	X_T_FMT_AMPM_E = 5
	X_T_FMT_E      = 4

	X_YESEXPR_A = 56
	X_YESEXPR_E = 48
	X_YESSTR_A  = 46
	X_YESSTR_E  = 56
	X_NOEXPR_A  = 57
	X_NOEXPR_E  = 49
	X_NOSTR_A   = 47
	X_NOSTR_E   = 57
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

type Nl_catd = uintptr
type Nl_item = int32

var _ uint8
