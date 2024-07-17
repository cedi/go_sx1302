package commands

import (
	"fmt"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
)

// LowLevel is a low-level handler of a MFRC522 RFID reader.
type LowLevel struct {
	resetPin    gpio.PinOut
	irqPin      gpio.PinIn
	spiDev      spi.Conn
	antennaGain int
	stop        chan struct{}
}

// NewLowLevelSPI creates and initializes the RFID card reader attached to SPI.
//
//	spiPort - the SPI device to use.
//	resetPin - reset GPIO pin.
//	irqPin - irq GPIO pin.
func NewLowLevelSPI(spiPort spi.Port, resetPin gpio.PinOut, irqPin gpio.PinIn) (*LowLevel, error) {
	if resetPin == nil {
		return nil, wrapf("reset pin is not set")
	}
	spiDev, err := spiPort.Connect(10*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		return nil, err
	}
	if err := resetPin.Out(gpio.High); err != nil {
		return nil, err
	}
	if irqPin != nil {
		if err := irqPin.In(gpio.PullUp, gpio.FallingEdge); err != nil {
			return nil, err
		}
	}

	dev := &LowLevel{
		spiDev:      spiDev,
		irqPin:      irqPin,
		resetPin:    resetPin,
		antennaGain: 4,
		stop:        make(chan struct{}, 1),
	}

	return dev, nil
}

// Init initializes the RFID chip.
func (r *LowLevel) Init() error {
	return nil
}

// DevWrite sends data to a device.
func (r *LowLevel) DevWrite(address int, data byte) error {
	newData := []byte{(byte(address) << 1) & 0x7E, data}
	return r.spiDev.Tx(newData, nil)
}

func wrapf(format string, a ...interface{}) error {
	return fmt.Errorf("mfrc522 lowlevel: "+format, a...)
}

func (r *LowLevel) writeCommandSequence(commands [][]byte) error {
	for _, cmdData := range commands {
		if err := r.DevWrite(int(cmdData[0]), cmdData[1]); err != nil {
			return err
		}
	}
	return nil
}
