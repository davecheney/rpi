package rpi

import (
	"log"
	"syscall"
	"unsafe"
)

var (
	gpfsel, gpset, gpclr, gplev []*uint32
)

func initGPIO(memfd int) {
	buf, err := syscall.Mmap(memfd, BCM2835_GPIO_BASE, BCM2835_BLOCK_SIZE, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("rpi: unable to mmap GPIO page: %v", err)
	}
	gpfsel = []*uint32{
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPFSEL0])),
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPFSEL1])),
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPFSEL2])),
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPFSEL3])),
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPFSEL4])),
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPFSEL5])),
	}
	gpset = []*uint32{
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPSET0])),
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPSET1])),
	}
	gpclr = []*uint32{
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPCLR0])),
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPCLR1])),
	}
	gplev = []*uint32{
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPLEV0])),
		(*uint32)(unsafe.Pointer(&buf[BCM2835_GPLEV1])),
	}
}

func GPIOSet(pin uint8) {
	offset := pin / 32
	shift := pin % 32
	*gpset[offset] = (1 << shift)
}

func GPIOClear(pin uint8) {
	offset := pin / 32
	shift := pin % 32
	*gpclr[offset] = (1 << shift)
}

func GPIOGet(pin uint8) bool {
	offset := pin / 32
	shift := pin % 32
	return *gplev[offset]&(1<<shift) == (1 << shift)
}

func GPIOFSel(pin, mode uint8) {
	offset := pin / 10
	shift := (pin % 10) * 3
	value := *gpfsel[offset]
	mask := BCM2835_GPIO_FSEL_MASK << shift
	value &= ^uint32(mask)
	value |= uint32(mode) << shift
	*gpfsel[offset] = value & mask
}
