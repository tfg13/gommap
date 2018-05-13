package gommap

import "golang.org/x/sys/unix"

func mmap_syscall(addr, length, prot, flags, fd uintptr, offset int64) (uintptr, error) {
	addr, _, err := unix.Syscall6(unix.SYS_MMAP, addr, length, prot, flags, fd, uintptr(offset))
	return addr, err
}
