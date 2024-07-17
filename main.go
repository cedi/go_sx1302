package main

import (
	log "github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"

	"github.com/cedi/go_sx1302/pkg/devices/sx1302"
	"github.com/cedi/go_sx1302/pkg/devices/sx1302/model"
)

var (
	version string
	commit  string
	date    string
	builtBy string
)

func main() {
	var boardConf model.BoardConf
	boardConf.LoRaWanPublic = true
	boardConf.ClkSrc = 0
	boardConf.FullDuplex = false
	boardConf.ComType = model.ComSPI
	boardConf.ComPath = "/dev/spidev0.0"

	var rfConf model.RxRf
	rfConf.Enable = true      // rf chain 0 needs to be enabled for calibration to work on sx1257
	rfConf.FreqHz = 868500000 // dummy
	rfConf.Type = model.RadioTypeSX1250
	rfConf.TxEnable = false
	rfConf.SingleInputMode = false

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	// Use spireg SPI port registry to find the first available SPI bus.
	port, err := spireg.Open(boardConf.ComPath)
	if err != nil {
		log.WithFields(log.Fields{
			"com_path": boardConf.ComPath,
		}).Fatal(err)
	}
	defer port.Close()

	lora := sx1302.NewSX1302Device(
		sx1302.WithBoardConfig(&boardConf),
		sx1302.WithRfRxConfig(0, &rfConf),
		sx1302.WithSPIPort(port, rpi.P1_13, rpi.P1_11),
	)

	if err := lora.Start(); err != nil {
		log.Fatal(err)
	}
}
