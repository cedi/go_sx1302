package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/cedi/go_sx1302/pkg/liblora"
	"github.com/cedi/go_sx1302/pkg/liblora/model"
)

var (
	version string
	commit  string
	date    string
	builtBy string
)

func main() {
	var lora liblora.LoraGw

	var boardConf model.BoardConf
	boardConf.LoRaWanPublic = true
	boardConf.ClkSrc = 0
	boardConf.FullDuplex = false
	boardConf.ComType = model.ComSPI
	boardConf.ComPath = "/dev/spidev0.0"

	if err := lora.SetBoardConf(&boardConf); err != nil {
		log.Fatal(err)
	}

	var rfConf model.RxRf
	rfConf.Enable = true      // rf chain 0 needs to be enabled for calibration to work on sx1257
	rfConf.FreqHz = 868500000 // dummy
	rfConf.Type = model.RadioTypeSX1250
	rfConf.TxEnable = false
	rfConf.SingleInputMode = false

	if err := lora.SetRfRxConf(0, &rfConf); err != nil {
		log.Fatal(err)
	}

	if err := lora.Start(); err != nil {
		log.Fatal(err)
	}
}
