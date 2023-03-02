package libc

import (
	"os"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
	"modernc.org/libc/errno"
	"modernc.org/libc/fcntl"
	// "modernc.org/libc/signal"
	"modernc.org/libc/sys/types"
	"modernc.org/libc/utime"
)

// int sigaction(int signum, const struct sigaction *act, struct sigaction *oldact);
func Xsigaction(t *TLS, signum int32, act, oldact uintptr) (r int32) {
	// we don't have an implementation for SYS_SIGACTION yet
	panic(todo(""))
}

// int fcntl(int fd, int cmd, ... /* arg */ );
func Xfcntl64(t *TLS, fd, cmd int32, args uintptr) (r int32) {
	var err error
	var p uintptr
	// var i int
	switch cmd {
	case fcntl.F_GETLK, fcntl.F_SETLK:
		p = *(*uintptr)(unsafe.Pointer(args))
		err = unix.FcntlFlock(uintptr(fd), int(cmd), (*unix.Flock_t)(unsafe.Pointer(p)))
	// case fcntl.F_GETFL, fcntl.F_FULLFSYNC:
	// 	i, err = unix.FcntlInt(uintptr(fd), int(cmd), 0)
	// 	r = int32(i)
	case fcntl.F_SETFD, fcntl.F_SETFL:
		arg := *(*int32)(unsafe.Pointer(args))
		_, err = unix.FcntlInt(uintptr(fd), int(cmd), int(arg))
	default:
		panic(todo("%v: %v %v", origin(1), fd, cmd))
	}
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return r
}

// int lstat(const char *pathname, struct stat *statbuf);
func Xlstat64(t *TLS, pathname, statbuf uintptr) int32 {
	if err := unix.Lstat(GoString(pathname), (*unix.Stat_t)(unsafe.Pointer(statbuf))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int stat(const char *pathname, struct stat *statbuf);
func Xstat64(t *TLS, pathname, statbuf uintptr) int32 {
	if err := unix.Stat(GoString(pathname), (*unix.Stat_t)(unsafe.Pointer(statbuf))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int fstat(int fd, struct stat *statbuf);
func Xfstat64(t *TLS, fd int32, statbuf uintptr) int32 {
	if err := unix.Fstat(int(fd), (*unix.Stat_t)(unsafe.Pointer(statbuf))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

func Xmmap(t *TLS, addr uintptr, length types.Size_t, prot, flags, fd int32, offset types.Off_t) uintptr {
	return Xmmap64(t, addr, length, prot, flags, fd, offset)
}

// void *mmap(void *addr, size_t length, int prot, int flags, int fd, off_t offset);
func Xmmap64(t *TLS, addr uintptr, length types.Size_t, prot, flags, fd int32, offset types.Off_t) uintptr {
	// We can't throw away the address, we may need it. No choice but to use CallLeFuncByPtr
	data := leFuncCall(syscall.SYS_MMAP, []uintptr{addr, uintptr(length), uintptr(prot), uintptr(flags), uintptr(fd), uintptr(offset)})
	// on error MAP_FAILED (i.e. (void *) -1) is returned, and errno is set to indicate the cause of the error.
	if data == ^uintptr(0) {
		err := leFuncError()
		t.setErrno(err)
		return ^uintptr(0) // (void*)-1
	}

	return data
}

// x/sys/unix does not implement SYS_REMAP on zos
// void *mremap(void *old_address, size_t old_size, size_t new_size, int flags, ... /* void *new_address */);
func Xmremap(t *TLS, old_address uintptr, old_size, new_size types.Size_t, flags int32, args uintptr) uintptr {
	panic(todo(""))
	// var arg uintptr
	// if args != 0 {
	// 	arg = *(*uintptr)(unsafe.Pointer(args))
	// }
	// data, _, err := unix.Syscall_syscall6(unix.SYS_MREMAP, old_address, uintptr(old_size), uintptr(new_size), uintptr(flags), arg, 0)
	// if err != 0 {
	// 	t.setErrno(err)
	// 	return ^uintptr(0) // (void*)-1
	// }

	// return data
}

// int ftruncate(int fd, off_t length);
func Xftruncate64(t *TLS, fd int32, length types.Off_t) int32 {
	if err := unix.Ftruncate(int(fd), int64(length)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// off64_t lseek64(int fd, off64_t offset, int whence);
func Xlseek64(t *TLS, fd int32, offset types.Off_t, whence int32) types.Off_t {
	n, err := unix.Seek(int(fd), int64(offset), int(whence))
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return types.Off_t(n)
}

// int utime(const char *filename, const struct utimbuf *times);
func Xutime(t *TLS, filename, times uintptr) int32 {
	var a []unix.Timeval
	if times != 0 {
		a = make([]unix.Timeval, 2)
		a[0].Sec = (*utime.Utimbuf)(unsafe.Pointer(times)).Factime
		a[1].Sec = (*utime.Utimbuf)(unsafe.Pointer(times)).Fmodtime
	}
	if err := unix.Utimes(GoString(filename), a); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// x/sys/unix does not implement SYS_ALARM on zos
// unsigned int alarm(unsigned int seconds);
func Xalarm(t *TLS, seconds uint32) uint32 {
	panic(todo(""))
	// n, _, err := unix.Syscall(unix.SYS_ALARM, uintptr(seconds), 0, 0)
	// if err != 0 {
	// 	panic(todo(""))
	// }

	// return uint32(n)
}

// time_t time(time_t *tloc);
func Xtime(t *TLS, tloc uintptr) types.Time_t {
	n, err := unix.Time((*unix.Time_t)(unsafe.Pointer(tloc)))
	if err != nil {
		t.setErrno(err)
		return types.Time_t(-1)
	}

	return types.Time_t(n)
}

// int getrlimit(int resource, struct rlimit *rlim);
func Xgetrlimit64(t *TLS, resource int32, rlim uintptr) int32 {
	if err := unix.Getrlimit(int(resource), (*unix.Rlimit)(unsafe.Pointer(rlim))); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int mkdir(const char *path, mode_t mode);
func Xmkdir(t *TLS, path uintptr, mode types.Mode_t) int32 {
	if err := unix.Mkdir(GoString(path), uint32(mode)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int symlink(const char *target, const char *linkpath);
func Xsymlink(t *TLS, target, linkpath uintptr) int32 {
	if err := unix.Symlink(GoString(target), GoString(linkpath)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int chmod(const char *pathname, mode_t mode)
func Xchmod(t *TLS, pathname uintptr, mode types.Mode_t) int32 {
	if err := unix.Chmod(GoString(pathname), uint32(mode)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int utimes(const char *filename, const struct timeval times[2]);
func Xutimes(t *TLS, filename, times uintptr) int32 {
	var a []unix.Timeval
	if times != 0 {
		a = make([]unix.Timeval, 2)
		a[0] = *(*unix.Timeval)(unsafe.Pointer(times))
		a[1] = *(*unix.Timeval)(unsafe.Pointer(times + unsafe.Sizeof(unix.Timeval{})))
	}
	if err := unix.Utimes(GoString(filename), a); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int unlink(const char *pathname);
func Xunlink(t *TLS, pathname uintptr) int32 {
	if err := unix.Unlink(GoString(pathname)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int access(const char *pathname, int mode);
func Xaccess(t *TLS, pathname uintptr, mode int32) int32 {
	if err := unix.Access(GoString(pathname), uint32(mode)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int rmdir(const char *pathname);
func Xrmdir(t *TLS, pathname uintptr) int32 {
	if err := unix.Rmdir(GoString(pathname)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int rename(const char *oldpath, const char *newpath);
func Xrename(t *TLS, oldpath, newpath uintptr) int32 {
	if err := unix.Rename(GoString(oldpath), GoString(newpath)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int mknod(const char *pathname, mode_t mode, dev_t dev);
func Xmknod(t *TLS, pathname uintptr, mode types.Mode_t, dev types.Dev_t) int32 {
	if err := unix.Mknod(GoString(pathname), uint32(mode), int(dev)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int chown(const char *pathname, uid_t owner, gid_t group);
func Xchown(t *TLS, pathname uintptr, owner types.Uid_t, group types.Gid_t) int32 {
	if err := unix.Chown(GoString(pathname), int(owner), int(group)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int link(const char *oldpath, const char *newpath);
func Xlink(t *TLS, oldpath, newpath uintptr) int32 {
	if err := unix.Link(GoString(oldpath), GoString(newpath)); err != nil {
		t.setErrno(err)
		return -1
	}

	return 0
}

// int pipe(int pipefd[2]);
func Xpipe(t *TLS, pipefd uintptr) int32 {
	var a [2]int
	if err := unix.Pipe(a[:]); err != nil {
		t.setErrno(err)
		return -1
	}

	*(*[2]int32)(unsafe.Pointer(pipefd)) = [2]int32{int32(a[0]), int32(a[1])}
	return 0
}

// int dup2(int oldfd, int newfd);
func Xdup2(t *TLS, oldfd, newfd int32) int32 {
	err := unix.Dup2(int(oldfd), int(newfd))
	if err != nil {
		t.setErrno(err)
		return -1
	}

	// TODO - Not sure this is correct, but reading the man page for dup2 leads me to believe this is supposed to be returned
	return int32(newfd)
}

// ssize_t readlink(const char *restrict path, char *restrict buf, size_t bufsize);
func Xreadlink(t *TLS, path, buf uintptr, bufsize types.Size_t) types.Ssize_t {
	var n int
	var err error
	switch {
	case buf == 0 || bufsize == 0:
		n, err = unix.Readlink(GoString(path), nil)
	default:
		n, err = unix.Readlink(GoString(path), (*RawMem)(unsafe.Pointer(buf))[:bufsize:bufsize])
	}
	if err != nil {
		t.setErrno(err)
		return -1
	}

	return types.Ssize_t(n)
}

// FILE *fopen64(const char *pathname, const char *mode);
func Xfopen64(t *TLS, pathname, mode uintptr) uintptr {
	m := strings.ReplaceAll(GoString(mode), "b", "")
	var flags int
	switch m {
	case "r":
		flags = os.O_RDONLY
	case "r+":
		flags = os.O_RDWR
	case "w":
		flags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	case "w+":
		flags = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	case "a":
		flags = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	case "a+":
		flags = os.O_RDWR | os.O_CREATE | os.O_APPEND
	default:
		panic(m)
	}
	//TODO- flags |= fcntl.O_LARGEFILE
	fd, err := unix.Open(GoString(pathname), int(flags|unix.O_LARGEFILE), uint32(0666))
	if err != nil{
		t.setErrno(err)
		return 0
	}

	if p := newFile(t, int32(fd)); p != 0 {
		return p
	}

	Xclose(t, int32(fd))
	t.setErrno(errno.ENOMEM)
	return 0
}
