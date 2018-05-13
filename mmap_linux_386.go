package gommap

import "golang.org/x/sys/unix"

func mmap_syscall(addr, length, prot, flags, fd uintptr, offset int64) (uintptr, error) {
	page := uintptr(offset / 4096)
	if offset != int64(page)*4096 {
		return 0, unix.EINVAL
	}
	addr, _, err := unix.Syscall6(unix.SYS_MMAP2, addr, length, prot, flags, fd, page)
	return addr, err
}
