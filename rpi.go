package rpi

import (
	"log"
	"os"
	"syscall"
	"unsafe"
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
	BCM2835_GPIO_BASE = BCM2835_PERI_BASE + 0x200000

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

        BCM2835_GPFSEL0   = 0x0000 // GPIO Function Select 0
        BCM2835_GPFSEL1   = 0x0004 // GPIO Function Select 1
        BCM2835_GPFSEL2   = 0x0008 // GPIO Function Select 2
        BCM2835_GPFSEL3   = 0x000c // GPIO Function Select 3
        BCM2835_GPFSEL4   = 0x0010 // GPIO Function Select 4
        BCM2835_GPFSEL5   = 0x0014 // GPIO Function Select 5
        BCM2835_GPSET0    = 0x001c // GPIO Pin Output Set 0
        BCM2835_GPSET1    = 0x0020 // GPIO Pin Output Set 1
        BCM2835_GPCLR0    = 0x0028 // GPIO Pin Output Clear 0
        BCM2835_GPCLR1    = 0x002c // GPIO Pin Output Clear 1
        BCM2835_GPLEV0    = 0x0034 // GPIO Pin Level 0
        BCM2835_GPLEV1    = 0x0038 // GPIO Pin Level 1
        BCM2835_GPEDS0    = 0x0040 // GPIO Pin Event Detect Status 0
        BCM2835_GPEDS1    = 0x0044 // GPIO Pin Event Detect Status 1
        BCM2835_GPREN0    = 0x004c // GPIO Pin Rising Edge Detect Enable 0
        BCM2835_GPREN1    = 0x0050 // GPIO Pin Rising Edge Detect Enable 1
        BCM2835_GPFEN0    = 0x0048 // GPIO Pin Falling Edge Detect Enable 0
        BCM2835_GPFEN1    = 0x005c // GPIO Pin Falling Edge Detect Enable 1
        BCM2835_GPHEN0    = 0x0064 // GPIO Pin High Detect Enable 0
        BCM2835_GPHEN1    = 0x0068 // GPIO Pin High Detect Enable 1
        BCM2835_GPLEN0    = 0x0070 // GPIO Pin Low Detect Enable 0
        BCM2835_GPLEN1    = 0x0074 // GPIO Pin Low Detect Enable 1
        BCM2835_GPAREN0   = 0x007c // GPIO Pin Async. Rising Edge Detect 0
        BCM2835_GPAREN1   = 0x0080 // GPIO Pin Async. Rising Edge Detect 1
        BCM2835_GPAFEN0   = 0x0088 // GPIO Pin Async. Falling Edge Detect 0
        BCM2835_GPAFEN1   = 0x008c // GPIO Pin Async. Falling Edge Detect 1
        BCM2835_GPPUD     = 0x0094 // GPIO Pin Pull-up/down Enable
        BCM2835_GPPUDCLK0 = 0x0098 // GPIO Pin Pull-up/down Enable Clock 0
        BCM2835_GPPUDCLK1 = 0x009c // GPIO Pin Pull-up/down Enable Clock 1

	GPIO_P1_22 = 25
)

var GPIO gpio

func init() {
	memfd, err := os.OpenFile("/dev/mem", os.O_RDWR|os.O_SYNC, 0)
	if err != nil {
		log.Fatalf("rpi: unable to open /dev/mem: %v", err)
	}
	buf, err := syscall.Mmap(int(memfd.Fd()), BCM2835_GPIO_BASE, BCM2835_BLOCK_SIZE, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("rpi: unable to mmap GPIO page: %v", err)
	}
        GPIO = gpio{
		set0: (*uint32)(unsafe.Pointer(&buf[BCM2835_GPSET0])),
		clr0: (*uint32)(unsafe.Pointer(&buf[BCM2835_GPCLR0])),
		lev0: (*uint32)(unsafe.Pointer(&buf[BCM2835_GPLEV0])),
	}
}

type gpio struct {
        set0, clr0, lev0 *uint32
}

func (g *gpio) Set(pin uint8) {
        shift := pin % 32
	*g.set0 = (1 << shift)
}

func (g *gpio) Clear(pin uint8) {
        shift := pin % 32
	*g.clr0 = (1 << shift)
}

func (g *gpio) Get(pin uint8) bool {
        shift := pin % 32
	return *g.lev0 & (1 << shift) == (1 << shift)
}
