// Code created manually. Would like to be generated in the future.

package unistd

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
	F_LOCK                              = 1
	F_OK                                = 0x08
	F_TEST                              = 3
	F_TLOCK                             = 2
	F_ULOCK                             = 0
	INT16_MAX                           = (32767)
	INT16_MIN                           = (-INT16_MAX - 1)
	INT32_MAX                           = (2147483647)
	INT32_MIN                           = (-INT32_MAX - 1)
	INT64_MAX                           = (9223372036854775807)
	INT64_MIN                           = (-INT64_MAX - 1)
	INT8_MAX                            = (127)
	INT8_MIN                            = (-INT8_MAX - 1)
	INTMAX_MAX                          = (9223372036854775807)
	INTMAX_MIN                          = (-INTMAX_MAX - 1)
	INTPTR_MAX                          = (INT64_MAX)
	INTPTR_MIN                          = (INT64_MIN)
	INT_FAST16_MAX                      = (INT32_MAX)
	INT_FAST16_MIN                      = (INT32_MIN)
	INT_FAST32_MAX                      = (INT32_MAX)
	INT_FAST32_MIN                      = (INT32_MIN)
	INT_FAST64_MAX                      = (9223372036854775807)
	INT_FAST64_MIN                      = (-INT_FAST64_MAX - 1)
	INT_FAST8_MAX                       = (INT8_MAX)
	INT_FAST8_MIN                       = (INT8_MIN)
	INT_LEAST16_MAX                     = (INT16_MAX)
	INT_LEAST16_MIN                     = (INT16_MIN)
	INT_LEAST32_MAX                     = (INT32_MAX)
	INT_LEAST32_MIN                     = (INT32_MIN)
	INT_LEAST64_MAX                     = (9223372036854775807)
	INT_LEAST64_MIN                     = (-INT_LEAST64_MAX - 1)
	INT_LEAST8_MAX                      = (INT8_MAX)
	INT_LEAST8_MIN                      = (INT8_MIN)
	NULL                                = (uintptr)(0)
	PATH_MAX                            = 1023
	PRIX16                              = "hX"
	PRIX32                              = "X"
	PRIX64                              = "lX"
	PRIX8                               = "hhX"
	PRIXFAST16                          = "X"
	PRIXFAST32                          = "X"
	PRIXFAST64                          = "lX"
	PRIXFAST8                           = "hhX"
	PRIXLEAST16                         = "hX"
	PRIXLEAST32                         = "X"
	PRIXLEAST64                         = "lX"
	PRIXLEAST8                          = "hhX"
	PRIXMAX                             = "jX"
	PRIXPTR                             = "lX"
	PRId16                              = "hd"
	PRId32                              = "d"
	PRId64                              = "ld"
	PRId8                               = "hhd"
	PRIdFAST16                          = "d"
	PRIdFAST32                          = "d"
	PRIdFAST64                          = "ld"
	PRIdFAST8                           = "hhd"
	PRIdLEAST16                         = "hd"
	PRIdLEAST32                         = "d"
	PRIdLEAST64                         = "ld"
	PRIdLEAST8                          = "hhd"
	PRIdMAX                             = "jd"
	PRIdPTR                             = "ld"
	PRIi16                              = "hi"
	PRIi32                              = "i"
	PRIi64                              = "li"
	PRIi8                               = "hhi"
	PRIiFAST16                          = "i"
	PRIiFAST32                          = "i"
	PRIiFAST64                          = "li"
	PRIiFAST8                           = "hhi"
	PRIiLEAST16                         = "hi"
	PRIiLEAST32                         = "i"
	PRIiLEAST64                         = "li"
	PRIiLEAST8                          = "hhi"
	PRIiMAX                             = "ji"
	PRIiPTR                             = "li"
	PRIo16                              = "ho"
	PRIo32                              = "o"
	PRIo64                              = "lo"
	PRIo8                               = "hho"
	PRIoFAST16                          = "o"
	PRIoFAST32                          = "o"
	PRIoFAST64                          = "lo"
	PRIoFAST8                           = "hho"
	PRIoLEAST16                         = "ho"
	PRIoLEAST32                         = "o"
	PRIoLEAST64                         = "lo"
	PRIoLEAST8                          = "hho"
	PRIoMAX                             = "jo"
	PRIoPTR                             = "lo"
	PRIu16                              = "hu"
	PRIu32                              = "u"
	PRIu64                              = "lu"
	PRIu8                               = "hhu"
	PRIuFAST16                          = "u"
	PRIuFAST32                          = "u"
	PRIuFAST64                          = "lu"
	PRIuFAST8                           = "hhu"
	PRIuLEAST16                         = "hu"
	PRIuLEAST32                         = "u"
	PRIuLEAST64                         = "lu"
	PRIuLEAST8                          = "hhu"
	PRIuMAX                             = "ju"
	PRIuPTR                             = "lu"
	PRIx16                              = "hx"
	PRIx32                              = "x"
	PRIx64                              = "lx"
	PRIx8                               = "hhx"
	PRIxFAST16                          = "x"
	PRIxFAST32                          = "x"
	PRIxFAST64                          = "lx"
	PRIxFAST8                           = "hhx"
	PRIxLEAST16                         = "hx"
	PRIxLEAST32                         = "x"
	PRIxLEAST64                         = "lx"
	PRIxLEAST8                          = "hhx"
	PRIxMAX                             = "jx"
	PRIxPTR                             = "lx"
	PTRDIFF_MAX                         = (INT64_MAX)
	PTRDIFF_MIN                         = (INT64_MIN)
	R_OK                                = 0x04
	SCNd16                              = "hd"
	SCNd32                              = "d"
	SCNd64                              = "ld"
	SCNd8                               = "hhd"
	SCNdFAST16                          = "d"
	SCNdFAST32                          = "d"
	SCNdFAST64                          = "ld"
	SCNdFAST8                           = "hhd"
	SCNdLEAST16                         = "hd"
	SCNdLEAST32                         = "d"
	SCNdLEAST64                         = "ld"
	SCNdLEAST8                          = "hhd"
	SCNdMAX                             = "jd"
	SCNdPTR                             = "ld"
	SCNi16                              = "hi"
	SCNi32                              = "i"
	SCNi64                              = "li"
	SCNi8                               = "hhi"
	SCNiFAST16                          = "i"
	SCNiFAST32                          = "i"
	SCNiFAST64                          = "li"
	SCNiFAST8                           = "hhi"
	SCNiLEAST16                         = "hi"
	SCNiLEAST32                         = "i"
	SCNiLEAST64                         = "li"
	SCNiLEAST8                          = "hhi"
	SCNiMAX                             = "ji"
	SCNiPTR                             = "li"
	SCNo16                              = "ho"
	SCNo32                              = "o"
	SCNo64                              = "lo"
	SCNo8                               = "hho"
	SCNoFAST16                          = "o"
	SCNoFAST32                          = "o"
	SCNoFAST64                          = "lo"
	SCNoFAST8                           = "hho"
	SCNoLEAST16                         = "ho"
	SCNoLEAST32                         = "o"
	SCNoLEAST64                         = "lo"
	SCNoLEAST8                          = "hho"
	SCNoMAX                             = "jo"
	SCNoPTR                             = "lo"
	SCNu16                              = "hu"
	SCNu32                              = "u"
	SCNu64                              = "lu"
	SCNu8                               = "hhu"
	SCNuFAST16                          = "u"
	SCNuFAST32                          = "u"
	SCNuFAST64                          = "lu"
	SCNuFAST8                           = "hhu"
	SCNuLEAST16                         = "hu"
	SCNuLEAST32                         = "u"
	SCNuLEAST64                         = "lu"
	SCNuLEAST8                          = "hhu"
	SCNuMAX                             = "ju"
	SCNuPTR                             = "lu"
	SCNx16                              = "hx"
	SCNx32                              = "x"
	SCNx64                              = "lx"
	SCNx8                               = "hhx"
	SCNxFAST16                          = "x"
	SCNxFAST32                          = "x"
	SCNxFAST64                          = "lx"
	SCNxFAST8                           = "hhx"
	SCNxLEAST16                         = "hx"
	SCNxLEAST32                         = "x"
	SCNxLEAST64                         = "lx"
	SCNxLEAST8                          = "hhx"
	SCNxMAX                             = "jx"
	SCNxPTR                             = "lx"
	SEEK_CUR                            = 1
	SEEK_END                            = 2
	SEEK_SET                            = 0
	SIG_ATOMIC_MAX                      = (INT64_MAX)
	SIG_ATOMIC_MIN                      = (INT64_MIN)
	SIZE_MAX                            = (UINT64_MAX)
	STDERR_FILENO                       = 2
	STDIN_FILENO                        = 0
	STDOUT_FILENO                       = 1
	UINT16_MAX                          = (uint16)(65535)
	UINT32_MAX                          = (uint32)(4294967295)
	UINT64_MAX                          = (uint64)(18446744073709551615)
	UINT8_MAX                           = (uint8)(255)
	UINTMAX_MAX                         = (uint64)(18446744073709551615)
	UINTPTR_MAX                         = (UINT64_MAX)
	UINT_FAST16_MAX                     = (UINT32_MAX)
	UINT_FAST32_MAX                     = (UINT32_MAX)
	UINT_FAST64_MAX                     = (uint64)(18446744073709551615)
	UINT_FAST8_MAX                      = (UINT8_MAX)
	UINT_LEAST16_MAX                    = (UINT16_MAX)
	UINT_LEAST32_MAX                    = (UINT32_MAX)
	UINT_LEAST64_MAX                    = (uint64)(18446744073709551615)
	UINT_LEAST8_MAX                     = (UINT8_MAX)
	WCHAR_MAX                           = UINT32_MAX
	WCHAR_MIN                           = 0
	WINT_MAX                            = (INT32_MAX)
	WINT_MIN                            = (INT32_MIN)
	W_OK                                = 0x02
	X_OK                                = 0x01
	X_AE_BIMODAL                        = 1
	X_ALL_SOURCE                        = 1
	X_BIG_ENDIAN                        = 1
	X_CHAR_UNSIGNED                     = 1
	X_CS_PATH                           = 1
	X_CS_POSIX_V6_ILP32_OFF32_CFLAGS    = 3
	X_CS_POSIX_V6_ILP32_OFF32_LDFLAGS   = 4
	X_CS_POSIX_V6_ILP32_OFF32_LIBS      = 5
	X_CS_POSIX_V6_ILP32_OFFBIG_CFLAGS   = 6
	X_CS_POSIX_V6_ILP32_OFFBIG_LDFLAGS  = 7
	X_CS_POSIX_V6_ILP32_OFFBIG_LIBS     = 8
	X_CS_POSIX_V6_LP64_OFF64_CFLAGS     = 9
	X_CS_POSIX_V6_LP64_OFF64_LDFLAGS    = 10
	X_CS_POSIX_V6_LP64_OFF64_LIBS       = 11
	X_CS_POSIX_V6_LPBIG_OFFBIG_CFLAGS   = 12
	X_CS_POSIX_V6_LPBIG_OFFBIG_LDFLAGS  = 13
	X_CS_POSIX_V6_LPBIG_OFFBIG_LIBS     = 14
	X_CS_POSIX_V6_WIDTH_RESTRICTED_ENVS = 15
	X_CS_SHELL                          = 2
	X_CS_XBS5_ILP32_OFF32_CFLAGS        = 16
	X_CS_XBS5_ILP32_OFF32_LDFLAGS       = 17
	X_CS_XBS5_ILP32_OFF32_LIBS          = 18
	X_CS_XBS5_ILP32_OFF32_LINTFLAGS     = 19
	X_CS_XBS5_ILP32_OFFBIG_CFLAGS       = 20
	X_CS_XBS5_ILP32_OFFBIG_LDFLAGS      = 21
	X_CS_XBS5_ILP32_OFFBIG_LIBS         = 22
	X_CS_XBS5_ILP32_OFFBIG_LINTFLAGS    = 23
	X_CS_XBS5_LP64_OFF64_CFLAGS         = 24
	X_CS_XBS5_LP64_OFF64_LDFLAGS        = 25
	X_CS_XBS5_LP64_OFF64_LIBS           = 26
	X_CS_XBS5_LP64_OFF64_LINTFLAGS      = 27
	X_CS_XBS5_LPBIG_OFFBIG_CFLAGS       = 28
	X_CS_XBS5_LPBIG_OFFBIG_LDFLAGS      = 29
	X_CS_XBS5_LPBIG_OFFBIG_LIBS         = 30
	X_CS_XBS5_LPBIG_OFFBIG_LINTFLAGS    = 31
	X_ENHANCED_ASCII_EXT                = 0x410A0000
	X_EXT                               = 1
	X_ISOC99_SOURCE                     = 1
	X_LARGE_TIME_API                    = 1
	X_LONG_LONG                         = 1
	X_LP64                              = 1
	X_MI_BUILTIN                        = 1
	X_MSC_ENABLE                        = 1
	X_MSC_ENABLED                       = 1
	X_MSC_ENABLED_COND                  = 2
	X_MSC_FAILED                        = -1
	X_MSC_NOT_ENABLED                   = 0
	X_MSC_QUERY                         = 0
	X_OPEN_DEFAULT                      = 1
	X_OPEN_MSGQ_EXT                     = 1
	X_OPEN_SYS                          = 1
	X_OPEN_SYS_FILE_EXT                 = 1
	X_OPEN_SYS_SOCK_IPV6                = 1
	X_OSENV_GET                         = 1
	X_OSENV_PERSIST                     = 8
	X_OSENV_SECURITY                    = 2
	X_OSENV_SET                         = 2
	X_OSENV_UNPERSIST                   = 16
	X_OSENV_UNSET                       = 4
	X_OSENV_WLM                         = 1
	X_PC_ACL                            = 10
	X_PC_ACL_ENTRIES_MAX                = 11
	X_PC_CHOWN_RESTRICTED               = 1
	X_PC_LINK_MAX                       = 2
	X_PC_MAX_CANON                      = 3
	X_PC_MAX_INPUT                      = 4
	X_PC_NAME_MAX                       = 5
	X_PC_NO_TRUNC                       = 6
	X_PC_PATH_MAX                       = 7
	X_PC_PIPE_BUF                       = 8
	X_PC_VDISABLE                       = 9
	X_POSIX2_CHAR_TERM                  = 1
	X_POSIX2_C_BIND                     = (int32)(200112)
	X_POSIX2_C_DEV                      = (int32)(200112)
	X_POSIX2_C_VERSION                  = (int32)(199209)
	X_POSIX2_LOCALEDEF                  = (int32)(200112)
	X_POSIX2_SW_DEV                     = 1
	X_POSIX2_UPE                        = (int32)(200112)
	X_POSIX2_VERSION                    = (int32)(199209)
	X_POSIX_IPV6                        = (int32)(200112)
	X_POSIX_JOB_CONTROL                 = 1
	X_POSIX_READER_WRITER_LOCKS         = (int32)(200112)
	X_POSIX_REGEXP                      = 1
	X_POSIX_SAVED_IDS                   = 1
	X_POSIX_SHELL                       = 1
	X_POSIX_THREADS                     = (int32)(200112)
	X_POSIX_THREAD_ATTR_STACKSIZE       = (int32)(200112)
	X_POSIX_THREAD_PROCESS_SHARED       = (int32)(200112)
	X_POSIX_THREAD_SAFE_FUNCTIONS       = (int32)(200112)
	X_POSIX_V6_ILP32_OFF32              = 1
	X_POSIX_V6_ILP32_OFFBIG             = 1
	X_POSIX_V6_LP64_OFF64               = 1
	X_POSIX_V6_LPBIG_OFFBIG             = 1
	X_POSIX_VERSION                     = (int32)(199209)
	X_SC_2_CHAR_TERM                    = 12
	X_SC_2_C_BIND                       = 115
	X_SC_2_C_DEV                        = 111
	X_SC_2_C_VERSION                    = 116
	X_SC_2_FORT_DEV                     = 112
	X_SC_2_FORT_RUN                     = 110
	X_SC_2_LOCALEDEF                    = 114
	X_SC_2_PBS                          = 200
	X_SC_2_PBS_ACCOUNTING               = 201
	X_SC_2_PBS_CHECKPOINT               = 202
	X_SC_2_PBS_LOCATE                   = 203
	X_SC_2_PBS_MESSAGE                  = 204
	X_SC_2_PBS_TRACK                    = 205
	X_SC_2_SW_DEV                       = 113
	X_SC_2_UPE                          = 117
	X_SC_2_VERSION                      = 109
	X_SC_ADVISORY_INFO                  = 148
	X_SC_AIO_LISTIO_MAX                 = 140
	X_SC_AIO_MAX                        = 141
	X_SC_AIO_PRIO_DELTA_MAX             = 142
	X_SC_ARG_MAX                        = 1
	X_SC_ASYNCHRONOUS_IO                = 150
	X_SC_ATEXIT_MAX                     = 133
	X_SC_BARRIERS                       = 149
	X_SC_BC_BASE_MAX                    = 101
	X_SC_BC_DIM_MAX                     = 102
	X_SC_BC_SCALE_MAX                   = 103
	X_SC_BC_STRING_MAX                  = 104
	X_SC_CHILD_MAX                      = 2
	X_SC_CLK_TCK                        = 3
	X_SC_CLOCK_SELECTION                = 151
	X_SC_COLL_WEIGHTS_MAX               = 105
	X_SC_CPUTIME                        = 152
	X_SC_DELAYTIMER_MAX                 = 143
	X_SC_EXPR_NEST_MAX                  = 106
	X_SC_FSYNC                          = 153
	X_SC_GETGR_R_SIZE_MAX               = 139
	X_SC_GETPW_R_SIZE_MAX               = 138
	X_SC_HOST_NAME_MAX                  = 144
	X_SC_IOV_MAX                        = 134
	X_SC_IPV6                           = 154
	X_SC_JOB_CONTROL                    = 4
	X_SC_LINE_MAX                       = 107
	X_SC_LOGIN_NAME_MAX                 = 145
	X_SC_MAPPED_FILES                   = 155
	X_SC_MEMLOCK                        = 156
	X_SC_MEMLOCK_RANGE                  = 157
	X_SC_MEMORY_PROTECTION              = 158
	X_SC_MESSAGE_PASSING                = 159
	X_SC_MMAP_MEM_MAX_NP                = 14
	X_SC_MONOTONIC_CLOCK                = 160
	X_SC_MQ_OPEN_MAX                    = 146
	X_SC_MQ_PRIO_MAX                    = 147
	X_SC_NGROUPS_MAX                    = 5
	X_SC_OPEN_MAX                       = 6
	X_SC_PAGESIZE                       = 135
	X_SC_PAGE_SIZE                      = 136
	X_SC_PASS_MAX                       = 132
	X_SC_PRIORITIZED_IO                 = 161
	X_SC_PRIORITY_SCHEDULING            = 162
	X_SC_RAW_SOCKETS                    = 163
	X_SC_READER_WRITER_LOCKS            = 164
	X_SC_REALTIME_SIGNALS               = 165
	X_SC_REGEXP                         = 166
	X_SC_RE_DUP_MAX                     = 108
	X_SC_RTSIG_MAX                      = 209
	X_SC_SAVED_IDS                      = 7
	X_SC_SEMAPHORES                     = 167
	X_SC_SEM_NSEMS_MAX                  = 210
	X_SC_SEM_VALUE_MAX                  = 211
	X_SC_SHARED_MEMORY_OBJECTS          = 168
	X_SC_SHELL                          = 169
	X_SC_SIGQUEUE_MAX                   = 212
	X_SC_SPAWN                          = 170
	X_SC_SPIN_LOCKS                     = 171
	X_SC_SPORADIC_SERVER                = 172
	X_SC_SS_REPL_MAX                    = 173
	X_SC_STREAM_MAX                     = 118
	X_SC_SYMLOOP_MAX                    = 213
	X_SC_SYNCHRONIZED_IO                = 174
	X_SC_THREADS                        = 184
	X_SC_THREADS_MAX_NP                 = 13
	X_SC_THREAD_ATTR_STACKADDR          = 175
	X_SC_THREAD_ATTR_STACKSIZE          = 176
	X_SC_THREAD_CPUTIME                 = 177
	X_SC_THREAD_DESTRUCTOR_ITERATIONS   = 206
	X_SC_THREAD_KEYS_MAX                = 207
	X_SC_THREAD_PRIORITY_SCHEDULING     = 180
	X_SC_THREAD_PRIO_INHERIT            = 178
	X_SC_THREAD_PRIO_PROTECT            = 179
	X_SC_THREAD_PROCESS_SHARED          = 181
	X_SC_THREAD_SAFE_FUNCTIONS          = 182
	X_SC_THREAD_SPORADIC_SERVER         = 183
	X_SC_THREAD_STACK_MIN               = 208
	X_SC_THREAD_TASKS_MAX_NP            = 11
	X_SC_THREAD_THREADS_MAX             = X_SC_THREADS_MAX_NP
	X_SC_TIMEOUTS                       = 185
	X_SC_TIMERS                         = 186
	X_SC_TIMER_MAX                      = 214
	X_SC_TRACE                          = 187
	X_SC_TRACE_EVENT_FILTER             = 188
	X_SC_TRACE_EVENT_NAME_MAX           = 189
	X_SC_TRACE_INHERIT                  = 190
	X_SC_TRACE_LOG                      = 191
	X_SC_TRACE_NAME_MAX                 = 192
	X_SC_TRACE_SYS_MAX                  = 193
	X_SC_TRACE_USER_EVENT_MAX           = 194
	X_SC_TTY_GROUP                      = 15
	X_SC_TTY_NAME_MAX                   = 215
	X_SC_TYPED_MEMORY_OBJECTS           = 195
	X_SC_TZNAME_MAX                     = 9
	X_SC_V6_ILP32_OFF32                 = 196
	X_SC_V6_ILP32_OFFBIG                = 197
	X_SC_V6_LP64_OFF64                  = 198
	X_SC_V6_LPBIG_OFFBIG                = 199
	X_SC_VERSION                        = 10
	X_SC_XBS5_ILP32_OFF32               = 216
	X_SC_XBS5_ILP32_OFFBIG              = 217
	X_SC_XBS5_LP64_OFF64                = 218
	X_SC_XBS5_LPBIG_OFFBIG              = 219
	X_SC_XOPEN_CRYPT                    = 127
	X_SC_XOPEN_ENH_I18N                 = 128
	X_SC_XOPEN_LEGACY                   = 220
	X_SC_XOPEN_REALTIME                 = 221
	X_SC_XOPEN_REALTIME_THREADS         = 222
	X_SC_XOPEN_SHM                      = 129
	X_SC_XOPEN_STREAMS                  = 223
	X_SC_XOPEN_UNIX                     = 137
	X_SC_XOPEN_VERSION                  = 130
	X_SC_XOPEN_XCU_VERSION              = 131
	X_SMF_IEFU83                        = 0x00000001
	X_SMF_IEFU84                        = 0x00000002
	X_UNIX03_SOURCE                     = 1
	X_UNIX03_THREADS                    = 1
	X_UNIX03_WITHDRAWN                  = 1
	X_XBS5_ILP32_OFF32                  = 1
	X_XBS5_ILP32_OFFBIG                 = 1
	X_XBS5_LP64_OFF64                   = 1
	X_XBS5_LPBIG_OFFBIG                 = 1
	X_XOPEN_CRYPT                       = 1
	X_XOPEN_ENH_I18N                    = 1
	X_XOPEN_LEGACY                      = 1
	X_XOPEN_REALTIME                    = -1
	X_XOPEN_REALTIME_THREADS            = -1
	X_XOPEN_SOURCE                      = 600
	X_XOPEN_SOURCE_EXTENDED             = 1
	X_XOPEN_STREAMS                     = -1
	X_XOPEN_UNIX                        = 1
	X_XOPEN_VERSION                     = 420
	X_XOPEN_XCU_VERSION                 = 4
	X_XOPEN_XPG4                        = 1
	X__features_h                       = 1
	X__inttypes                         = 1
	X__types                            = 1
	X__unistd                           = 1
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

// (ZOS) - Adjusted
type X__sigset_t = struct {
	X__sigs_0 uint32
	X__sigs_1 uint32
} /* sys/types.h:200 */

// A set of signals to be blocked, unblocked, or waited for.
type Sigset_t = X__sigset_t

// An integral type that can be modified atomically, without the
//    possibility of a signal arriving in the middle of the operation.
type Sig_atomic_t = X__sig_atomic_t

// Adjusted for zos
type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec int64
}

type Ssize_t = X__ssize_t
type Gid_t = X__gid_t
type Uid_t = X__uid_t
type Off_t = X__off64_t
type Useconds_t = X__useconds_t
type Pid_t = X__pid_t
type Intptr_t = X__intptr_t
type Socklen_t = X__socklen_t

var _ uint8
