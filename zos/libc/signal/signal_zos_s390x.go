// Code created manually. Would like to be generated in the future.

package signal

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
	MINSIGSTKSZ   = 2105344
	SA_NODEFER    = 0x01000000
	SA_NOCLDWAIT  = 0x02000000
	SA_SIGINFO    = 0x04000000
	SA_RESTART    = 0x08000000
	SA_RESETHAND  = 0x10000000
	SA_ONSTACK    = 0x20000000
	SA_NOCLDSTOP  = 0x80000000
	SIGHUP        = 1
	SIGINT        = 2
	SIGABRT       = 3
	SIGILL        = 4
	SIGPOLL       = 5
	SIGURG        = 6
	SIGSTOP       = 7
	SIGFPE        = 8
	SIGKILL       = 9
	SIGBUS        = 10
	SIGSEGV       = 11
	SIGSYS        = 12
	SIGPIPE       = 13
	SIGALRM       = 14
	SIGTERM       = 15
	SIGUSR1       = 16
	SIGUSR2       = 17
	SIGABND       = 18
	SIGCONT       = 19
	SIGCHLD       = 20
	SIGTTIN       = 21
	SIGTTOU       = 22
	SIGIO         = 23
	SIGQUIT       = 24
	SIGTSTP       = 25
	SIGTRAP       = 26
	SIGIOERR      = 27
	SIGWINCH      = 28
	SIGXCPU       = 29
	SIGXFSZ       = 30
	SIGVTALRM     = 31
	SIGPROF       = 32
	SIGDANGER     = 33
	SIGTHSTOP     = 34
	SIGTHCONT     = 35
	SIGTRACE      = 37
	SIGDCE        = 38
	SIGDUMP       = 39
	X_LP64        = 1
	X__features_h = 1
	X__types      = 1
)

// `sigev_notify` values.
const (
	SIGEV_NONE   = 1
	SIGEV_SIGNAL = 0
)

// `si_code` values for SIGSEGV signal
const (
	SEGV_MAPERR  = 51 // Address not mapped to an object
	SEGV_ACCERR  = 52 // Invalid permissions for a mapped object
	SEGV_PROTECT = 53 // Protection exception
	SEGV_ADDRESS = 54 // Addressing exception
)

// `si_code` values for SIGBUS signal
const (
	BUS_ADRALN = 71 // Invalid address alignment
	BUS_ADRERR = 72 // Non-existent physical address
	BUS_OBJERR = 73 // Object specific hardware error
)

// `si_code` values for SIGCHLD signal
const (
	CLD_EXITED    = 101 // Child process exited
	CLD_KILLED    = 102 // Child process killed
	CLD_DUMPED    = 103 // Child process terminated abnormally
	CLD_TRAPPED   = 104 // Traced child process encountered a trap
	CLD_STOPPED   = 105 // Child process stopped
	CLD_CONTINUED = 106 // Child process continued after being stopped
)

// `si_code` values for SIGPOLL signal
const (
	POLL_IN  = 111 // Input data is available
	POLL_OUT = 112 // Output buffers are available
	POLL_MSG = 113 // Input message is available
	POLL_ERR = 114 // I/O error
	POLL_PRI = 115 // High priority input is available
	POLL_HUP = 116 // Device was disconnected
)

// Values for `si_code`. Positive values are reserved for kernel-generatd signals.
const (
	SI_ASYNCIO = 175
	SI_QUEUE   = 176
)

// `si_code` values for SIGILL signal
const (
	ILL_ILLOPC  = 11 // Illegal opcode
	ILL_ILLOPN  = 12 // Illegal operand
	ILL_ILLADR  = 13 // Illegal addressing mode
	ILL_ILLTRP  = 14 // Illegal trap
	ILL_PRVOPC  = 15 // Priveleged opcode
	ILL_PRVREG  = 16 // Priveleged register
	ILL_COPROC  = 17 // Coprocessor error
	ILL_BADSTK  = 18 // Internal stack error
	ILL_EXECUTE = 19 // Execute exception
	ILL_ILLSPEC = 20 // Specification exception
)

// `si_code` values for SIGFPE signal
const (
	FPE_INTDIV  = 31 // Integer divide by zero
	FPE_INTOVF  = 32 // Integer overflow
	FPE_FLTDIV  = 33 // Float divide by zero
	FPE_FLTOVF  = 34 // Float overflow
	FPE_FLTUND  = 35 // Float underflow
	FPE_FLTRES  = 36 // Float inexact result
	FPE_FLTINV  = 37 // Invalid float operation
	FPE_FLTSUB  = 38 // Subscript out of range
	FPE_FLTSIG  = 39 // Significance exception
	FPE_DECDATA = 40 // Data exception
	FPE_DECDIV  = 41 // Decimal divide exception
	FPE_DECOVF  = 42 // Decimal overflow exception
	FPE_UNKWN   = 43 // Undetermined exception
	FPE_CTDXC   = 44 // Compare and Trap exception
)

// Flags used in the ss_flags member of a stack_t
const (
	SS_ONSTACK = 1
	SS_DISABLE = 2
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

type X__dev_t = uint32 /* sys/types.h:66 */
type X__uid_t = int32  /* sys/types.h:162 */ // uid_t is uint32 if __VM__ is defined
type X__gid_t = int32  /* sys/types.h:77 */  // gid_t is uint32 if __VM__ is defined
type X__ino_t = uint32 /* sys/types.h:86 */
// type X__ino64_t = uint64
type X__mode_t = int32  /* sys/types.h:93 */
type X__nlink_t = int32 /* sys/types.h:101 */
type X__off_t = int64   /* sys/types.h:113 */
type X__off64_t = int64 /* sys/types.h:123 */
type X__pid_t = int32   /* sys/types.h:129 */
// type X__fsid_t = struct{ F__val [2]int32 }
type X__clock_t = uint32 /* sys/types.h:129 */
// type X__rlim_t = uint64
// type X__rlim64_t = uint64
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

type Pid_t = X__pid_t
type Uid_t = X__uid_t

// Adjusted for zos
type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec int64
}

// TODO (ZOS) - May require adjustments.
// Type for data associated with a signal.
type Sigval = struct {
	_          [0]uint64
	Fsival_int int32
	_          [4]byte
} /* signal.h:512 */

type X__sigval_t = Sigval
type Sigval_t = X__sigval_t

// (ZOS) - Adjusted
type Siginfo_t = struct {
	Fsi_signo         int32
	Fsi_errno         int32
	Fsi_code          int32
	Fsi_pid           X__pid_t
	Fsi_uid           X__uid_t
	Fsi_addr_31       [4]uint8
	Fsi_status        int32
	Fsi_band_value_31 [8]uint8
	Fsi_resv          [4]uint8
	Fsi_addr          uintptr
	Fsi_value         X__sigval_t
	Fsi_band          int64
	// These are not present on zos, but will be kept
	F__pad0    int32
	F_sifields struct {
		_     [0]uint64
		F_pad [28]int32
	}
} /* signal.h:526 */

// TODO (ZOS) - Not sure about this struct.
type Pthread_attr_t1 = struct {
	__ [13]float64
} /* signal.h:645 */

// (ZOS) - Adjusted
// Structure to transport application-defined values with signals.
type Sigevent = struct {
	Fsigev_value             X__sigval_t
	Fsigev_signo             int32
	Fsigev_notify            int32
	Fsigev_notify_attributes *Pthread_attr_t
} /* signal.h:656 */

// Structure to transport application-defined values with signals.
type Sigevent_t = Sigevent

// TODO - Doesn't seem to be implemented on zos?
type X__sighandler_t = uintptr
type Sig_t = X__sighandler_t

// (ZOS) - Adjusted
// Structure describing the action to be taken when a signal arrives.
type Sigaction = struct {
	Fsa_handler uintptr
	Fsa_flags   int32
	Fsa_mask    X__sigset_t
} /* signal.h:183 */

type Stack_t = struct {
	Fss_sp    uintptr
	Fss_size  Size_t
	Fss_flags int32
}

// TODO (ZOS) - May require further change
// Context to describe whole processor state.
type X__mcontext_t = [40]int64 /* sys/types.h:462 */

type Mcontext_t = X__mcontext_t

// (ZOS) - Adjusted
// Userlevel context.
type Ucontext_t1 = struct {
	Fuc_link     uintptr
	Fuc_stack    Stack_t
	Fuc_mcontext Mcontext_t
	Fuc_sigmask  Sigset_t
} /* signal.h:582 */

// Userlevel context.
type Ucontext_t = Ucontext_t1

// Structure describing a signal stack (obsolete).
type Sigstack = struct {
	Fss_sp      uintptr
	Fss_onstack int32
} /* signal.h:594 */

// TODO (ZOS) - May need adjustment
// Thread identifiers.  The structure of the attribute type is not
//    exposed on purpose.
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
