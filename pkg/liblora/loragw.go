package liblora

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"

	"github.com/cedi/go_sx1302/pkg/liblora/model"
)

type LoraGw struct {
	context model.LgwContext
}

func (l *LoraGw) SetBoardConf(conf *model.BoardConf) error {
	if l.context.IsStarted {
		return errors.New("gateway is already running. Please stop it before changing configuration")
	}

	if conf.ComType == model.COMUnknown {
		return errors.New("COM type is unknown")
	}

	l.context.BoardConfig = conf

	log.WithFields(log.Fields{
		"com_type":       conf.ComType,
		"com_path":       conf.ComPath,
		"lorawan_public": conf.LoRaWanPublic,
		"clksrc":         conf.ClkSrc,
		"full_duplex":    conf.FullDuplex,
	}).Info("Board configuration loaded")

	return nil
}

func (l *LoraGw) SetRfRxConf(rfChain uint8, conf *model.RxRf) error {
	if l.context.IsStarted {
		return errors.New("gateway is already running. Please stop it before changing configuration")
	}

	// if radio is disabled -> nothing to do
	if !conf.Enable {
		log.WithFields(log.Fields{
			"rf_chain": rfChain,
		}).Info("RF-Chain disabled")
		return nil
	}

	if rfChain >= model.MaxRfChains {
		return errors.New(fmt.Sprintf("rf-chain %d is not a valid rf-chain number", rfChain))
	}

	// check if the radio-type is supported
	if conf.Type == model.RadioTypeNone {
		return errors.New("not a valid radio-type")
	}

	// check if frequency is valid
	if conf.FreqHz < model.RfRxFreqMin || conf.FreqHz > model.RfRxFreqMax {
		return errors.New("invalid radio center frequency. Please check the if it has been given in Hz!")
	}

	if len(l.context.RfChainCfg) < int(rfChain) {
		l.context.RfChainCfg[rfChain] = *conf
	} else {
		l.context.RfChainCfg = append(l.context.RfChainCfg, *conf)
	}

	log.WithFields(log.Fields{
		"rf_chain":          rfChain,
		"enable":            conf.Enable,
		"freq_hz":           conf.FreqHz,
		"rssi_offset":       conf.RssiOffset,
		"rssi_t_comp":       conf.RssiTComp,
		"type":              conf.Type,
		"tx_enable":         conf.TxEnable,
		"single_input_mode": conf.SingleInputMode,
	}).Info("RF configuration loaded")

	return nil
}

func (l *LoraGw) Start() error {
	if l.context.IsStarted {
		return errors.New("gateway is already running. Please stop it before changing configuration")
	}

	// ToDo: Support USB as well
	if l.context.BoardConfig.ComType != model.ComSPI {
		return errors.New("Only SPI is supported as of now")
	}

	l.connectSPI()

	return nil
}

func (l *LoraGw) connectSPI() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	if _, err := driverreg.Init(); err != nil {
		log.Fatal(err)
	}

	// Use spireg SPI port registry to find the first available SPI bus.
	p, err := spireg.Open(l.context.BoardConfig.ComPath)
	if err != nil {
		log.WithFields(log.Fields{
			"com_path": l.context.BoardConfig.ComPath,
		}).Fatal(err)
	}
	defer p.Close()

	// Convert the spi.Port into a spi.Conn so it can be used for communication.
	c, err := p.Connect(2*physic.MegaHertz, spi.Mode3, 8)
	if err != nil {
		log.WithFields(log.Fields{
			"frequency": 2 * physic.MegaHertz,
			"mode":      spi.Mode3,
			"bits":      8,
		}).Fatal(err)
	}

	// Prints out the gpio pin used.
	if p, ok := c.(spi.Pins); ok {
		log.WithFields(log.Fields{
			"CLK":  p.CLK(),
			"MOSI": p.MOSI(),
			"MISO": p.MISO(),
			"CS":   p.CS(),
		}).Info("GPIO PIN used")
	}
}
