package libc // import "modernc.org/libc"

// zos does not have header files for fts
import (
	// "bufio"
	// "encoding/hex"
	"fmt"
	// "math"
	// "math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	// "strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"

	// guuid "github.com/google/uuid"
	"golang.org/x/sys/unix"
	"modernc.org/libc/errno"
	"modernc.org/libc/fcntl"
	// "modernc.org/libc/fts"
	// "modernc.org/libc/grp"
	// gonetdb "modernc.org/libc/honnef.co/go/netdb"
	"modernc.org/libc/langinfo"
	"modernc.org/libc/limits"
	// "modernc.org/libc/netdb"
	"modernc.org/libc/netinet/in"
	"modernc.org/libc/pwd"
	// "modernc.org/libc/signal"
	"modernc.org/libc/stdio"
	// "modernc.org/libc/stdlib"
	"modernc.org/libc/sys/socket"
	"modernc.org/libc/sys/stat"
	"modernc.org/libc/sys/types"
	"modernc.org/libc/termios"
	ctime "modernc.org/libc/time"
	"modernc.org/libc/unistd"
	// "modernc.org/libc/uuid/uuid"
)

var (
	in6_addr_any in.In6_addr
	// _            = X__ctype_b_loc //X__ctype_b_loc is a musl function
)

// ZOS sys/types doesn't have X__syscall_slong_t or X__syscall_ulong_t
type (
	long  = int64
	ulong = uint64
)

// Copying the approach that the windows port took
var zos_fdLock sync.Mutex
var zos_fd_to_file = map[int32]*file{}

type file struct {
	_fd    int32
	hadErr bool
	t      uintptr // See token() in etc.go
}

// (TODO) This may be redundant
func fdToFile(fd int32) (*file, bool) {
	zos_fdLock.Lock()
	defer zos_fdLock.Unlock()
	f, ok := zos_fd_to_file[fd]
	return f, ok
}

func remFile(f *file) {
	removeObject(f.t)

	zos_fdLock.Lock()
	defer zos_fdLock.Unlock()
	delete(zos_fd_to_file, f._fd)
}

func (f *file) err() bool {
	return f.hadErr
}

func (f *file) setErr() {
	f.hadErr = true
}

func newFile(t *TLS, fd int32) uintptr {
	var f = file{_fd: fd}

	zos_fdLock.Lock()
	defer zos_fdLock.Unlock()
	zos_fd_to_file[fd] = &f
	f.t = addObject(&f)
	return f.t
}

func (f *file) close(t *TLS) int32 {
	remFile(f)

	r := Xclose(t, f._fd)
	if r < 0 {
		return stdio.EOF
	}

	return 0
}

func fwrite(fd int32, b []byte) (int, error) {
	if fd == unistd.STDOUT_FILENO {
		return write(b)
	}

	// if dmesgs {
	// 	dmesg("%v: fd %v: %s", origin(1), fd, hex.Dump(b))
	// }
	return unix.Write(int(fd), b)
}

// int fprintf(FILE *stream, const char *format, ...);
func Xfprintf(t *TLS, stream, format, args uintptr) int32 {
	f, ok := getObject(stream).(*file)
	if !ok {
		t.setErrno(errno.EBADF)
		return -1
	}
	n, _ := fwrite(f._fd, printf(format, args))
	return int32(n)
}

// int usleep(useconds_t usec);
func Xusleep(t *TLS, usec types.X__useconds_t) int32 {
	time.Sleep(time.Microsecond * time.Duration(usec))
	return 0
}

// int getrusage(int who, struct rusage *usage);
func Xgetrusage(t *TLS, who int32, usage uintptr) int32 {
	if err := unix.Getrusage(int(who), (*unix.Rusage)(unsafe.Pointer(usage))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// char *fgets(char *s, int size, FILE *stream);
func Xfgets(t *TLS, s uintptr, size int32, stream uintptr) uintptr {

	f, ok := getObject(stream).(*file)
	if !ok {
		t.setErrno(errno.EBADF)
		return 0
	}
	var b []byte
	buf := [1]byte{}
	for ; size > 0; size-- {
		n, err := unix.Read(int(f._fd), buf[:])
		if n != 0 {
			b = append(b, buf[0])
			if buf[0] == '\n' {
				b = append(b, 0)
				copy((*RawMem)(unsafe.Pointer(s))[:len(b):len(b)], b)
				return s
			}

			continue
		}

		switch {
		case n == 0 && err == nil && len(b) == 0:
			return 0
		default:
			panic(todo(""))
		}

		// if err == nil {
		// 	panic("internal error")
		// }

		// if len(b) != 0 {
		// 		b = append(b, 0)
		// 		copy((*RawMem)(unsafe.Pointer(s)[:len(b)]), b)
		// 		return s
		// }

		// t.setErrno(err)
	}
	panic(todo(""))
}

// int lstat(const char *pathname, struct stat *statbuf);
func Xlstat(t *TLS, pathname, statbuf uintptr) int32 {
	return Xlstat64(t, pathname, statbuf)
}

// int stat(const char *pathname, struct stat *statbuf);
func Xstat(t *TLS, pathname, statbuf uintptr) int32 {
	return Xstat64(t, pathname, statbuf)
}

// int chdir(const char *path);
func Xchdir(t *TLS, path uintptr) int32 {
	if err := unix.Chdir(GoString(path)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

var localtime ctime.Tm

// struct tm *localtime(const time_t *timep);
func Xlocaltime(_ *TLS, timep uintptr) uintptr {
	loc := time.Local
	if r := getenv(Environ(), "TZ"); r != 0 {
		zone, off := parseZone(GoString(r))
		loc = time.FixedZone(zone, -off)
	}
	ut := *(*unix.Time_t)(unsafe.Pointer(timep))
	t := time.Unix(int64(ut), 0).In(loc)
	localtime.Ftm_sec = int32(t.Second())
	localtime.Ftm_min = int32(t.Minute())
	localtime.Ftm_hour = int32(t.Hour())
	localtime.Ftm_mday = int32(t.Day())
	localtime.Ftm_mon = int32(t.Month() - 1)
	localtime.Ftm_year = int32(t.Year() - 1900)
	localtime.Ftm_wday = int32(t.Weekday())
	localtime.Ftm_yday = int32(t.YearDay())
	localtime.Ftm_isdst = Bool32(isTimeDST(t))
	return uintptr(unsafe.Pointer(&localtime))
}

// struct tm *localtime_r(const time_t *timep, struct tm *result);
func Xlocaltime_r(_ *TLS, timep, result uintptr) uintptr {
	loc := time.Local
	if r := getenv(Environ(), "TZ"); r != 0 {
		zone, off := parseZone(GoString(r))
		loc = time.FixedZone(zone, -off)
	}
	ut := *(*unix.Time_t)(unsafe.Pointer(timep))
	t := time.Unix(int64(ut), 0).In(loc)
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_sec = int32(t.Second())
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_min = int32(t.Minute())
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_hour = int32(t.Hour())
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_mday = int32(t.Day())
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_mon = int32(t.Month() - 1)
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_year = int32(t.Year() - 1900)
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_wday = int32(t.Weekday())
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_yday = int32(t.YearDay())
	(*ctime.Tm)(unsafe.Pointer(result)).Ftm_isdst = Bool32(isTimeDST(t))
	return result
}

// int open(const char *pathname, int flags, ...);
func Xopen(t *TLS, pathname uintptr, flags int32, args uintptr) int32 {
	return Xopen64(t, pathname, flags, args)
}

// int open(const char *pathname, int flags, ...);
func Xopen64(t *TLS, pathname uintptr, flags int32, args uintptr) int32 {
	var mode types.Mode_t
	if args != 0 {
		mode = (types.Mode_t)(VaUint32(&args))
	}
	fd, err := unix.Open(GoString(pathname), int(flags), uint32(mode))
	if err != nil {
		t.setErrno(err)
		return -1
	}

	newFile(t, int32(fd))

	return int32(fd)
}

// off_t lseek(int fd, off_t offset, int whence);
func Xlseek(t *TLS, fd int32, offset types.Off_t, whence int32) types.Off_t {
	return types.Off_t(Xlseek64(t, fd, offset, whence))
}

func whenceStr(whence int32) string {
	switch whence {
	case fcntl.SEEK_CUR:
		return "SEEK_CUR"
	case fcntl.SEEK_END:
		return "SEEK_END"
	case fcntl.SEEK_SET:
		return "SEEK_SET"
	default:
		return fmt.Sprintf("whence(%d)", whence)
	}
}

var fsyncStatbuf stat.Stat

// int fsync(int fd);
func Xfsync(t *TLS, fd int32) int32 {
	if noFsync {
		// Simulate -DSQLITE_NO_SYNC for sqlite3 testfixture, see function full_sync in sqlite3.c
		return Xfstat(t, fd, uintptr(unsafe.Pointer(&fsyncStatbuf)))
	}

	if err := unix.Fsync(int(fd)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// long sysconf(int name);
func Xsysconf(t *TLS, name int32) long {
	switch name {
	case unistd.X_SC_PAGESIZE:
		return long(unix.Getpagesize())
	case unistd.X_SC_GETPW_R_SIZE_MAX:
		return -1
	case unistd.X_SC_GETGR_R_SIZE_MAX:
		return -1
	}

	panic(todo("", name))
	switch name {
	case unistd.X_SC_PAGESIZE:
		return long(unix.Getpagesize())
	case unistd.X_SC_GETPW_R_SIZE_MAX:
		return -1
	case unistd.X_SC_GETGR_R_SIZE_MAX:
		return -1
	}

	panic(todo("", name))
}

// int close(int fd);
func Xclose(t *TLS, fd int32) int32 {
	if err := unix.Close(int(fd)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// char *getcwd(char *buf, size_t size);
func Xgetcwd(t *TLS, buf uintptr, size types.Size_t) uintptr {
	_, err := unix.Getcwd((*RawMem)(unsafe.Pointer(buf))[:size:size])
	if err != nil {
		t.setErrno(err)
		return 0
	}

	return buf
}

// int fstat(int fd, struct stat *statbuf);
func Xfstat(t *TLS, fd int32, statbuf uintptr) int32 {
	return Xfstat64(t, fd, statbuf)
}

// int ftruncate(int fd, off_t length);
func Xftruncate(t *TLS, fd int32, length types.Off_t) int32 {
	return Xftruncate64(t, fd, length)
}

func Xclosedir(t *TLS, dir uintptr) int32 {
	err := unix.Closedir(dir)
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// DIR *opendir(const char *name)
func Xopendir(t *TLS, name uintptr) uintptr {
	dir, err := unix.Opendir(GoString(name))
	if err != nil {
		t.setErrno(err)
		return 0
	}

	return dir
}

// struct dirent *readdir(DIR *dirp);
func Xreaddir(t *TLS, dir uintptr) uintptr {
	var de uintptr

	dirent, err := unix.Readdir(dir)
	if err != nil {
		t.setErrno(err)
		return 0
	}

	de = uintptr(unsafe.Pointer(dirent))
	return de
}

// int fcntl(int fd, int cmd, ... /* arg */ );
func Xfcntl(t *TLS, fd, cmd int32, args uintptr) int32 {
	return Xfcntl64(t, fd, cmd, args)
}

// ssize_t read(int fd, void *buf, size_t count);
func Xread(t *TLS, fd int32, buf uintptr, count types.Size_t) types.Ssize_t {
	var n int
	var err error
	switch {
	case count == 0:
		n, err = unix.Read(int(fd), nil)
	default:
		n, err = unix.Read(int(fd), (*RawMem)(unsafe.Pointer(buf))[:count:count])
	}
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return types.Ssize_t(n)
}

// ssize_t write(int fd, const void *buf, size_t count);
func Xwrite(t *TLS, fd int32, buf uintptr, count types.Size_t) types.Ssize_t {
	var n int
	var err error
	switch {
	case count == 0:
		n, err = unix.Write(int(fd), nil)
	default:
		n, err = unix.Write(int(fd), (*RawMem)(unsafe.Pointer(buf))[:count:count])
	}
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return types.Ssize_t(n)
}

// int fchmod(int fd, mode_t mode);
func Xfchmod(t *TLS, fd int32, mode types.Mode_t) int32 {
	if err := unix.Fchmod(int(fd), uint32(mode)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int fchown(int fd, uid_t owner, gid_t group);
func Xfchown(t *TLS, fd int32, owner types.Uid_t, group types.Gid_t) int32 {
	if err := unix.Fchown(int(fd), int(owner), int(group)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// uid_t geteuid(void);
func Xgeteuid(t *TLS) types.Uid_t {
	return types.Uid_t(unix.Getuid())
}

// int munmap(void *addr, size_t length);
func Xmunmap(t *TLS, addr uintptr, length types.Size_t) int32 {
	// TODO - Darwin implementation says we can't avoid using syscall. This is so that it matches Xmmap.
	if err := leFuncCall(syscall.SYS_MUNMAP, []uintptr{addr, uintptr(length)}); err != 0 {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int gettimeofday(struct timeval *tv, struct timezone *tz);
func Xgettimeofday(t *TLS, tv, tz uintptr) int32 {
	if tz != 0 {
		panic(todo(""))
	}

	var tvs unix.Timeval
	err := unix.Gettimeofday(&tvs)
	if err != nil {
		t.setErrno(err)
		return -1
	}

	*(*unix.Timeval)(unsafe.Pointer(tv)) = tvs
	return 0
}

// int getsockopt(int sockfd, int level, int optname, void *optval, socklen_t *optlen);
func Xgetsockopt(t *TLS, sockfd, level, optname int32, optval, optlen uintptr) int32 {
	// if _, _, err := unix.Syscall_syscall6(unix.SYS_GETSOCKOPT, uintptr(sockfd), uintptr(level), uintptr(optname), optval, optlen, 0); err != 0 {
	// 	t.setErrno(err)
	// 	return -1
	// }

	// return 0

	panic(todo(""))
}

// int setsockopt(int sockfd, int level, int optname, const void *optval, socklen_t optlen);
func Xsetsockopt(t *TLS, sockfd, level, optname int32, optval uintptr, optlen socket.Socklen_t) int32 {
	// if _, _, err := unix.Syscall_syscall6(unix.SYS_SETSOCKOPT, uintptr(sockfd), uintptr(level), uintptr(optname), optval, uintptr(optlen), 0); err != 0 {
	// 	t.setErrno(err)
	// 	return -1
	// }

	// return 0

	panic(todo(""))
}

// int ioctl(int fd, unsigned long request, ...);
func Xioctl(t *TLS, fd int32, request ulong, va uintptr) int32 {
	var argp uintptr
	if va != 0 {
		argp = VaUintptr(&va)
	}
	n := leFuncCall(syscall.SYS_IOCTL, []uintptr{uintptr(fd), uintptr(request), argp})
	if n < 0 {
		err := leFuncError()
		t.setErrno(err)
		return -1
	}

	return int32(n)
}

// int getsockname(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
func Xgetsockname(t *TLS, sockfd int32, addr, addrlen uintptr) int32 {
	if err := leFuncCall(syscall.SYS___GETSOCKNAME_A, []uintptr{uintptr(sockfd), addr, addrlen}); err != 0 {
		t.setErrno(leFuncError())
		return -1
	}

	return 0
}

// int select(int nfds, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, struct timeval *timeout);
func Xselect(t *TLS, nfds int32, readfds, writefds, exceptfds, timeout uintptr) int32 {
	n, err := unix.Select(
		int(nfds),
		(*unix.FdSet)(unsafe.Pointer(readfds)),
		(*unix.FdSet)(unsafe.Pointer(writefds)),
		(*unix.FdSet)(unsafe.Pointer(exceptfds)),
		(*unix.Timeval)(unsafe.Pointer(timeout)),
	)
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return int32(n)
}

// int mkfifo(const char *pathname, mode_t mode);
func Xmkfifo(t *TLS, pathname uintptr, mode types.Mode_t) int32 {
	if err := unix.Mkfifo(GoString(pathname), uint32(mode)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// mode_t umask(mode_t mask);
func Xumask(t *TLS, mask types.Mode_t) types.Mode_t {
	return types.Mode_t(unix.Umask(int(mask)))
}

// ZOS - Not implemented in sys/unix
// int execvp(const char *file, char *const argv[]);

// pid_t waitpid(pid_t pid, int *wstatus, int options);
func Xwaitpid(t *TLS, pid types.Pid_t, wstatus uintptr, optname int32) types.Pid_t {
	n, err := unix.Wait4(int(pid), (*unix.WaitStatus)(unsafe.Pointer(wstatus)), int(optname), nil)
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return types.Pid_t(n)
}

// int uname(struct utsname *buf);
func Xuname(t *TLS, buf uintptr) int32 {
	if err := unix.Uname((*unix.Utsname)(unsafe.Pointer(buf))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// ssize_t recv(int sockfd, void *buf, size_t len, int flags);
func Xrecv(t *TLS, sockfd int32, buf uintptr, len types.Size_t, flags int32) types.Ssize_t {
	// n, _, err := unix.Syscall6(unix.SYS_RECVFROM, uintptr(sockfd), buf, uintptr(len), uintptr(flags), 0, 0)

	// ZOS - will need to investigate if this works, the file descriptor situation needs more attention
	n, _, err := unix.Recvfrom(int(sockfd), *(*[]byte)(unsafe.Pointer(buf)), int(flags))
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return types.Ssize_t(n)
}

// ssize_t send(int sockfd, const void *buf, size_t len, int flags);
func Xsend(t *TLS, sockfd int32, buf uintptr, len types.Size_t, flags int32) types.Ssize_t {
	// n, _, err := unix.Syscall6(unix.SYS_SENDTO, uintptr(sockfd), buf, uintptr(len), uintptr(flags), 0, 0)
	// if err != 0 {
	// 	t.setErrno(err)
	// 	return -1
	// }

	// return types.Ssize_t(n)
	panic(todo(""))
}

// int shutdown(int sockfd, int how);
func Xshutdown(t *TLS, sockfd, how int32) int32 {
	// TODO - may need adjustment, fd situation needs to be investigated
	if err := unix.Shutdown(int(sockfd), int(how)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int getpeername(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
func Xgetpeername(t *TLS, sockfd int32, addr uintptr, addrlen uintptr) int32 {
	if err := leFuncCall(syscall.SYS___GETPEERNAME_A, []uintptr{uintptr(sockfd), addr, uintptr(addrlen)}); err != 0 {
	t.setErrno(leFuncError())
		return -1
	}

	return 0
}

// int socket(int domain, int type, int protocol);
func Xsocket(t *TLS, domain, type1, protocol int32) int32 {
	fd, err := unix.Socket(int(domain), int(type1), int(protocol))
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return int32(fd)
}

// int bind(int sockfd, const struct sockaddr *addr, socklen_t addrlen);
func Xbind(t *TLS, sockfd int32, addr uintptr, addrlen uint32) int32 {
	n := leFuncCall(syscall.SYS___BIND_A, []uintptr{uintptr(sockfd), addr, uintptr(addrlen)})
	if n != 0 {
		err := leFuncError()
		t.setErrno(err)
		return -1
	}

	return int32(n)
}

// int connect(int sockfd, const struct sockaddr *addr, socklen_t addrlen);
func Xconnect(t *TLS, sockfd int32, addr uintptr, addrlen uint32) int32 {
	if err := leFuncCall(syscall.SYS___CONNECT_A, []uintptr{uintptr(sockfd), addr, uintptr(addrlen)}); err != 0 {
		t.setErrno(leFuncError())
		return -1
	}

	return 0
}

// int listen(int sockfd, int backlog);
func Xlisten(t *TLS, sockfd, backlog int32) int32 {
	if err := leFuncCall(syscall.SYS_LISTEN, []uintptr{uintptr(sockfd), uintptr(backlog)}); err != 0 {
		t.setErrno(leFuncError())
		return -1
	}

	return 0
}

// int accept(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
func Xaccept(t *TLS, sockfd int32, addr uintptr, addrlen uintptr) int32 {
	n, _, err := unix.Accept(int(sockfd))
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return int32(n)
}

// int getrlimit(int resource, struct rlimit *rlim);
func Xgetrlimit(t *TLS, resource int32, rlim uintptr) int32 {
	return Xgetrlimit64(t, resource, rlim)
}

// int setrlimit(int resource, const struct rlimit *rlim);
func Xsetrlimit(t *TLS, resource int32, rlim uintptr) int32 {
	return Xsetrlimit64(t, resource, rlim)
}

// int setrlimit(int resource, const struct rlimit *rlim);
func Xsetrlimit64(t *TLS, resource int32, rlim uintptr) int32 {
	if err := unix.Setrlimit(int(resource), (*unix.Rlimit)(unsafe.Pointer(rlim))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// uid_t getuid(void);
func Xgetuid(t *TLS) types.Uid_t {
	return types.Uid_t(os.Getuid())
}

// pid_t getpid(void);
func Xgetpid(t *TLS) int32 {
	return int32(os.Getpid())
}

// int system(const char *command);
func Xsystem(t *TLS, command uintptr) int32 {
	s := GoString(command)
	if command == 0 {
		panic(todo(""))
	}

	cmd := exec.Command("sh", "-c", s)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		ps := err.(*exec.ExitError)
		return int32(ps.ExitCode())
	}

	return 0
}

/*
TODO (ZOS) - pwd.Passwd is not the same on zos, it is missing a few properties
that are on the linux implementation. Will probably follow the approach that the
darwin implementation took and just leave related functions as todos.
*/

var staticGetpwuid pwd.Passwd

func init() {
	atExit = append(atExit, func() { closePasswd(&staticGetpwuid) })
}

// (ZOS) - copied over and adjusted from the linux implementation
func closePasswd(p *pwd.Passwd) {
	Xfree(nil, p.Fpw_name)
	Xfree(nil, p.Fpw_dir)
	Xfree(nil, p.Fpw_shell)
	*p = pwd.Passwd{}
}

// struct passwd *getpwuid(uid_t uid);

// (ZOS) - copied over and adjusted from the linux implementation
func initPasswd(t *TLS, p *pwd.Passwd, name string, uid, gid uint32, dir, shell string) {
	p.Fpw_name = cString(t, name)
	p.Fpw_uid = int32(uid)
	p.Fpw_gid = int32(gid)
	p.Fpw_dir = cString(t, dir)
	p.Fpw_shell = cString(t, shell)
}

// (ZOS) - copied over and adjusted from the linux implementation
func initPasswd2(t *TLS, buf uintptr, buflen types.Size_t, p *pwd.Passwd, name string, uid, gid uint32, dir, shell string) bool {
	p.Fpw_name, buf, buflen = bufString(buf, buflen, name)
	if buf == 0 {
		return false
	}

	p.Fpw_uid = int32(uid)
	p.Fpw_gid = int32(gid)
	if buf == 0 {
		return false
	}

	p.Fpw_dir, buf, buflen = bufString(buf, buflen, dir)
	if buf == 0 {
		return false
	}

	p.Fpw_shell, buf, buflen = bufString(buf, buflen, shell)
	if buf == 0 {
		return false
	}

	return true
}

func bufString(buf uintptr, buflen types.Size_t, s string) (uintptr, uintptr, types.Size_t) {
	buf0 := buf
	rq := len(s) + 1
	if rq > int(buflen) {
		return 0, 0, 0
	}

	copy((*RawMem)(unsafe.Pointer(buf))[:len(s):len(s)], s)
	buf += uintptr(len(s))
	*(*byte)(unsafe.Pointer(buf)) = 0
	return buf0, buf + 1, buflen - types.Size_t(rq)
}

// int getpwuid_r(uid_t uid, struct passwd *pwd, char *buf, size_t buflen, struct passwd **result);

// struct passwd *getpwnam(const char *name);

// int getpwnam_r(char *name, struct passwd *pwd, char *buf, size_t buflen, struct passwd **result);

// int setvbuf(FILE *stream, char *buf, int mode, size_t size);
func Xsetvbuf(t *TLS, stream, buf uintptr, mode int32, size types.Size_t) int32 {
	return 0 //TODO
}

// int raise(int sig);
func Xraise(t *TLS, sig int32) int32 {
	panic(todo(""))
}

// int backtrace(void **buffer, int size);
func Xbacktrace(t *TLS, buf uintptr, size int32) int32 {
	panic(todo(""))
}

// void backtrace_symbols_fd(void *const *buffer, int size, int fd);
func Xbacktrace_symbols_fd(t *TLS, buffer uintptr, size, fd int32) {
	panic(todo(""))
}

// int fileno(FILE *stream);
func Xfileno(t *TLS, stream uintptr) int32 {
	if stream == 0 {
		t.setErrno(errno.EBADF)
		return -1
	}

	f, ok := getObject(stream).(*file)
	if !ok {
		t.setErrno(errno.EBADF)
		return -1
	}
	return f._fd
}

var staticGetpwnam pwd.Passwd

func init() {
	atExit = append(atExit, func() { closePasswd(&staticGetpwnam) })
}

// var staticGetgrnam grp.Group

// func init() {
//	atExit = append(atExit, func() { closeGroup(&staticGetgrnam) })
// }

// struct group *getgrnam(const char *name);

// int getgrnam_r(const char *name, struct group *grp, char *buf, size_t buflen, struct group **result);

// struct group *getgrgid(gid_t gid);

// int getgrgid_r(gid_t gid, struct group *grp, char *buf, size_t buflen, struct group **result);

// int mkstemps(char *template, int suffixlen);

// int mkstemps(char *template, int suffixlen);

// int mkostemp(char *template, int flags);

// int mkstemp(char *template);

// int mkstemp(char *template);

// FTS *fts_open(char * const *path_argv, int options, int (*compar)(const FTSENT **, const FTSENT **));

// FTS *fts_open(char * const *path_argv, int options, int (*compar)(const FTSENT **, const FTSENT **));

// FTSENT *fts_read(FTS *ftsp);

// FTSENT *fts_read(FTS *ftsp);

// int fts_close(FTS *ftsp);

// int fts_close(FTS *ftsp);

// void tzset (void);
func Xtzset(t *TLS) {
	//TODO
}

var strerrorBuf [100]byte

// char *strerror(int errnum);
func Xstrerror(t *TLS, errnum int32) uintptr {
	if dmesgs {
		dmesg("%v: %v\n%s", origin(1), errnum, debug.Stack())
	}
	copy(strerrorBuf[:], fmt.Sprintf("strerror(%d)\x00", errnum))
	return uintptr(unsafe.Pointer(&strerrorBuf[0]))
}

// int strerror_r(int errnum, char *buf, size_t buflen);
func Xstrerror_r(t *TLS, errnum int32, buf uintptr, buflen types.Size_t) int32 {
	panic(todo(""))
}

// void *dlopen(const char *filename, int flags);
func Xdlopen(t *TLS, filename uintptr, flags int32) uintptr {
	panic(todo("%q", GoString(filename)))
}

// char *dlerror(void);
func Xdlerror(t *TLS) uintptr {
	panic(todo(""))
}

// int dlclose(void *handle);
func Xdlclose(t *TLS, handle uintptr) int32 {
	panic(todo(""))
}

// void *dlsym(void *handle, const char *symbol);
func Xdlsym(t *TLS, handle, symbol uintptr) uintptr {
	panic(todo(""))
}

// void perror(const char *s);
func Xperror(t *TLS, s uintptr) {
	panic(todo(""))
}

// int pclose(FILE *stream);
func Xpclose(t *TLS, stream uintptr) int32 {
	panic(todo(""))
}

var gai_strerrorBuf [100]byte

// const char *gai_strerror(int errcode);
func Xgai_strerror(t *TLS, errcode int32) uintptr {
	copy(gai_strerrorBuf[:], fmt.Sprintf("gai error %d\x00", errcode))
	return uintptr(unsafe.Pointer(&gai_strerrorBuf))
}

// int tcgetattr(int fd, struct termios *termios_p);
func Xtcgetattr(t *TLS, fd int32, termios_p uintptr) int32 {
	panic(todo(""))
}

// int tcsetattr(int fd, int optional_actions, const struct termios *termios_p);
func Xtcsetattr(t *TLS, fd, optional_actions int32, termios_p uintptr) int32 {
	panic(todo(""))
}

// speed_t cfgetospeed(const struct termios *termios_p);
func Xcfgetospeed(t *TLS, termios_p uintptr) termios.Speed_t {
	panic(todo(""))
}

// int cfsetospeed(struct termios *termios_p, speed_t speed);
func Xcfsetospeed(t *TLS, termios_p uintptr, speed uint32) int32 {
	panic(todo(""))
}

// int cfsetispeed(struct termios *termios_p, speed_t speed);
func Xcfsetispeed(t *TLS, termios_p uintptr, speed uint32) int32 {
	panic(todo(""))
}

// pid_t fork(void);
func Xfork(t *TLS) int32 {
	t.setErrno(errno.ENOSYS)
	return -1
}

var emptyStr = [1]byte{}

// char *setlocale(int category, const char *locale);
func Xsetlocale(t *TLS, category int32, locale uintptr) uintptr {
	return uintptr(unsafe.Pointer(&emptyStr)) //TODO
}

// char *nl_langinfo(nl_item item);
func Xnl_langinfo(t *TLS, item langinfo.Nl_item) uintptr {
	return uintptr(unsafe.Pointer(&emptyStr)) //TODO
}

// FILE *popen(const char *command, const char *type);
func Xpopen(t *TLS, command, type1 uintptr) uintptr {
	panic(todo(""))
}

// char *realpath(const char *path, char *resolved_path);
func Xrealpath(t *TLS, path, resolved_path uintptr) uintptr {
	s, err := filepath.EvalSymlinks(GoString(path))
	if err != nil {
		if os.IsNotExist(err) {
			t.setErrno(errno.ENOENT)
			return 0
		}

		panic(todo("", err))
	}

	if resolved_path == 0 {
		panic(todo(""))
	}

	if len(s) >= limits.PATH_MAX {
		s = s[:limits.PATH_MAX-1]
	}

	copy((*RawMem)(unsafe.Pointer(resolved_path))[:len(s):len(s)], s)
	(*RawMem)(unsafe.Pointer(resolved_path))[len(s)] = 0
	return resolved_path
}

// struct tm *gmtime_r(const time_t *timep, struct tm *result);
func Xgmtime_r(t *TLS, timep, result uintptr) uintptr {
	panic(todo(""))
}

// char *inet_ntoa(struct in_addr in);
func Xinet_ntoa(t *TLS, in1 in.In_addr) uintptr {
	panic(todo(""))
}

func X__ccgo_in6addr_anyp(t *TLS) uintptr {
	// Can't seem to find the definition for in6_addr_any
	return uintptr(unsafe.Pointer(&in6_addr_any))
}

func Xabort(t *TLS) {
	// zos is missing SYS_SIGACTION which is used by Xsigaction
	/*
		p := Xmalloc(t, types.Size_t(unsafe.Sizeof(signal.Sigaction{})))
		if p == 0 {
			panic("OOM")
		}

		*(*signal.Sigaction)(unsafe.Pointer(p)) = signal.Sigaction{
			Fsa_handler: signal.SIG_DFL
		}
		Xsigaction(t, signal.SIGABRT, p, 0)
		Xfree(t, p)
		unix.Kill(unix.Getpid(), syscall.Signal(signal.SIGABRT))
		panic(todo("unrechable"))
	*/
	panic(todo(""))
}

// int fflush(FILE *stream);
func Xfflush(t *TLS, stream uintptr) int32 {
	return 0 //TODO
}

// size_t fread(void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfread(t *TLS, ptr uintptr, size, nmemb types.Size_t, stream uintptr) types.Size_t {
	f, ok := getObject(stream).(*file)
	if !ok {
		t.setErrno(errno.EBADF)
		return 0
	}
	count := size * nmemb

	n, err := unix.Read(int(f._fd), (*RawMem)(unsafe.Pointer(ptr))[:count:count])
	if err != nil {
		f.setErr()
		return 0
	}

	return types.Size_t(n)
}

// size_t fwrite(const void *ptr, size_t size, size_t nmemb, FILE *stream);

// int fclose(FILE *stream);
func Xfclose(t *TLS, stream uintptr) int32 {
	f, ok := getObject(stream).(*file)
	if !ok {
		t.setErrno(errno.EBADF)
		return 0
	}

	return f.close(t)
}

// int fputc(int c, FILE *stream);
func Xfputc(t *TLS, c int32, stream uintptr) int32 {
	// TODO (ZOS) - need to figure out the fd situation first
	panic(todo(""))
}

// int fseek(FILE *stream, long offset, int whence);
func Xfseek(t *TLS, stream uintptr, offset long, whence int32) int32 {
	f, ok := getObject(stream).(*file)
	if !ok {
		t.setErrno(errno.EBADF)
		return -1
	}

	if n := Xlseek(t, f._fd, types.Off_t(offset), whence); n < 0 {
		f.setErr()
		return -1
	}

	return 0
}

// long ftell(FILE *stream);
func Xftell(t *TLS, stream uintptr) long {
	f, ok := getObject(stream).(*file)
	if !ok {
		t.setErrno(errno.EBADF)
		return -1
	}

	n := Xlseek(t, f._fd, 0, stdio.SEEK_CUR)
	if n < 0 {
		f.setErr()
		return -1
	}

	return long(n)
}

// int ferror(FILE *stream);
func Xferror(t *TLS, stream uintptr) int32 {
	f, ok := getObject(stream).(*file)
	if !ok {
		t.setErrno(errno.EBADF)
		return -1
	}

	return Bool32(f.err())
}

// int fgetc(FILE *stream);
func Xfgetc(t *TLS, stream uintptr) int32 {
	panic(todo(""))
}

// int getc(FILE *stream);
func Xgetc(t *TLS, stream uintptr) int32 {
	return Xfgetc(t, stream)
}

// int ungetc(int c, FILE *stream);
func Xungetc(t *TLS, c int32, stream uintptr) int32 {
	panic(todo(""))
}

// int fscanf(FILE *stream, const char *format, ...);
func Xfscanf(t *TLS, stream, format, va uintptr) int32 {
	panic(todo(""))
}

// int fputs(const char *s, FILE *stream);

// struct servent *getservbyname(const char *name, const char *proto);

// TODO - Uses musl?
// func Xreaddir64(t *TLS, dir uintptr) uintptr {
// 	return Xreaddir(t, dir)
// }

// func __syscall(r, _ uintptr, errno syscall.Errno) long {
// 	if errno != 0 {
// 		return long(-errno)
// 	}

// 	return long(r)
// }

// func X__syscall1(t *TLS, trap, p1 long) long {
// 	return __syscall(unix.Syscall(uintptr(trap), uintptr(p1), 0, 0))
// }

// func X__syscall3(t *TLS, trap, p1, p2, p3 long) long {
// 	return __syscall(unix.Syscall(uintptr(trap), uintptr(p1), uintptr(p2), uintptr(p3)))
// }

// func X__syscall4(t *TLS, trap, p1, p2, p3, p4 long) long {
// 	return __syscall(unix.Syscall6(uintptr(trap), uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), 0, 0))
// }

func fcntlCmdStr(cmd int32) string {
	switch cmd {
	case fcntl.F_GETOWN:
		return "F_GETOWN"
	case fcntl.F_SETLK:
		return "F_SETLK"
	case fcntl.F_GETLK:
		return "F_GETLK"
	case fcntl.F_SETFD:
		return "F_SETFD"
	case fcntl.F_GETFD:
		return "F_GETFD"
	default:
		return fmt.Sprintf("cmd(%d)", cmd)
	}
}

// int setenv(const char *name, const char *value, int overwrite);
func Xsetenv(t *TLS, name, value uintptr, overwrite int32) int32 {
	panic(todo(""))
}

// int unsetenv(const char *name);
func Xunsetenv(t *TLS, name uintptr) int32 {
	panic(todo(""))
}

// int pause(void);

// ssize_t writev(int fd, const struct iovec *iov, int iovcnt);

// void endpwent(void);
func Xendpwent(t *TLS) {
	// nop
}

// int __isoc99_sscanf(const char *str, const char *format, ...);
func X__isoc99_sscanf(t *TLS, str, format, va uintptr) int32 {
	r := scanf(strings.NewReader(GoString(str)), format, va)

	return r
}

// int sscanf(const char *str, const char *format, ...);
func Xsscanf(t *TLS, str, format, va uintptr) int32 {
	r := scanf(strings.NewReader(GoString(str)), format, va)

	return r
}

var ctimeStaticBuf [32]byte

// char *ctime(const time_t *timep);

// char *ctime_r(const time_t *timep, char *buf);

// ssize_t pwrite(int fd, const void *buf, size_t count, off_t offset);

// int fstatfs(int fd, struct statfs *buf);

// ssize_t getrandom(void *buf, size_t buflen, unsigned int flags);

// int posix_fadvise(int fd, off_t offset, off_t len, int advice);

// void uuid_generate_random(uuid_t out);
func Xuuid_generate_random(t *TLS, out uintptr) {
	panic(todo(""))
}

// void uuid_unparse(uuid_t uu, char *out);

// int uuid_parse( char *in, uuid_t uu);

// The initstate_r() function is like initstate(3) except that it initializes
// the state in the object pointed to by buf, rather than initializing the
// global state  variable.   Before  calling this function, the buf.state field
// must be initialized to NULL.  The initstate_r() function records a pointer
// to the statebuf argument inside the structure pointed to by buf.  Thus,
// state‐ buf should not be deallocated so long as buf is still in use.  (So,
// statebuf should typically be allocated as a static variable, or allocated on
// the heap using malloc(3) or similar.)
//
// char *initstate_r(unsigned int seed, char *statebuf, size_t statelen, struct random_data *buf);

// int random_r(struct random_data *buf, int32_t *result);

// void uuid_copy(uuid_t dst, uuid_t src);

// void rewind(FILE *stream);
func Xrewind(t *TLS, stream uintptr) {
	Xfseek(t, stream, 0, stdio.SEEK_SET)
}

// this is a musl function, it does not vary across ports
func Xstrdup(tls *TLS, s uintptr) uintptr { /* strdup.c:4:6: */
	var l types.Size_t = Xstrlen(tls, s)
	var d uintptr = Xmalloc(tls, (l + uint64(1)))
	if !(d != 0) {
		return uintptr(0)
	}
	return Xmemcpy(tls, d, s, (l + uint64(1)))
}

func leFuncCall(sysnum uintptr, args []uintptr) uintptr {
	return runtime.CallLeFuncByPtr(runtime.XplinkLibvec + sysnum<<4, args)
}

func leFuncError() syscall.Errno {
	err := *(*int32)(unsafe.Pointer(leFuncCall(syscall.SYS___ERRNO, []uintptr{})))
	return syscall.Errno(err)
}
