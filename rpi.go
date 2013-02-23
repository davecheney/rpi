package rpi

import (
	"log"
	"os"
	"syscall"
)

const (
	// Physical addresses for various peripheral register sets

	// Base Physical Address of the BCM 2835 peripheral registers
	BCM2835_PERI_BASE = 0x20000000

	// Base Physical Address of the System Timer registers
	BCM2835_ST_BASE = BCM2835_PERI_BASE + 0x3000

	// Base Physical Address of the Pads registers
	BCM2835_GPIO_PADS = BCM2835_PERI_BASE + 0x100000

	// Base Physical Address of the Clock/timer registers

	BCM2835_CLOCK_BASE = BCM2835_PERI_BASE + 0x101000

	// Base Physical Address of the GPIO registers
	BCM2835_GPIO_BASE = BCM2835_PERI_BASE + 0x20000

	// Base Physical Address of the SPI0 registers
	BCM2835_SPI0_BASE = BCM2835_PERI_BASE + 0x204000

	// Base Physical Address of the BSC0 registers
	BCM2835_BSC0_BASE = BCM2835_PERI_BASE + 0x205000

	// Base Physical Address of the PWM registers
	BCM2835_GPIO_PWM = BCM2835_PERI_BASE + 0x20C000

	// Base Physical Address of the BSC1 registers
	BCM2835_BSC1_BASE = BCM2835_PERI_BASE + 0x804000

	// Size of memory page on RPi
	BCM2835_PAGE_SIZE = 4 * 1024

	// Size of memory block on RPi
	BCM2835_BLOCK_SIZE = 4 * 1024
)

func init() {
	memfd, err := os.OpenFile("/dev/mem", os.O_RDWR|os.O_SYNC, 0)
	if err != nil {
		log.Fatalf("rpi: unable to open /dev/mem: %v", err)
	}
	buf, err := syscall.Mmap(int(memfd.Fd()), BCM2835_GPIO_BASE, BCM2835_BLOCK_SIZE, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("rpi: unable to mmap GPIO page: %v", err)
	}
	log.Println(buf)
}
