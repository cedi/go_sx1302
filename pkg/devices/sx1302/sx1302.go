package sx1302

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/spi"

	"github.com/cedi/go_sx1302/pkg/devices/sx1302/commands"
	"github.com/cedi/go_sx1302/pkg/devices/sx1302/model"
)

// Dev is an handle to an sx1302 LoRa HAT.
type Dev struct {
	context  model.LgwContext
	LowLevel *commands.LowLevel
}

// SX1302Config is the function option for the Options pattern
type SX1302Config func(*Dev)

func NewSX1302Device(opts ...SX1302Config) *Dev {
	d := &Dev{
		context: *model.NewLgwContextWithDefaults(),
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(d)
	}

	return d
}

// WithBoardConfig specify board configuration
func WithBoardConfig(conf *model.BoardConf) SX1302Config {
	return func(d *Dev) {
		if d.context.IsStarted {
			log.Fatal("gateway is already running. Please stop it before changing configuration")
		}

		d.context.BoardConfig = conf

		log.WithFields(log.Fields{
			"com_type":       conf.ComType,
			"com_path":       conf.ComPath,
			"lorawan_public": conf.LoRaWanPublic,
			"clksrc":         conf.ClkSrc,
			"full_duplex":    conf.FullDuplex,
		}).Info("Board configuration loaded")
	}
}

// WithRfRxConfig configures the RF-Chain
func WithRfRxConfig(rfChain uint8, conf *model.RxRf) SX1302Config {
	return func(d *Dev) {
		if d.context.IsStarted {
			log.Fatal("gateway is already running. Please stop it before changing configuration")
		}

		// if radio is disabled -> nothing to do
		if !conf.Enable {
			log.WithFields(log.Fields{
				"rf_chain": rfChain,
			}).Info("RF-Chain disabled")
			return
		}

		if rfChain >= model.MaxRfChains {
			log.Fatalf("rf-chain %d is not a valid rf-chain number", rfChain)
		}

		// check if frequency is valid
		if conf.FreqHz < model.RfRxFreqMin || conf.FreqHz > model.RfRxFreqMax {
			log.Fatal("invalid radio center frequency. Please check the if it has been given in Hz")
		}

		if len(d.context.RfChainCfg) < int(rfChain) {
			d.context.RfChainCfg[rfChain] = *conf
		} else {
			d.context.RfChainCfg = append(d.context.RfChainCfg, *conf)
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
	}
}

func WithSPIPort(spiPort spi.Port, resetPin gpio.PinOut, irqPin gpio.PinIn) SX1302Config {
	return func(d *Dev) {
		raw, err := commands.NewLowLevelSPI(spiPort, resetPin, irqPin)
		if err != nil {
			log.WithError(err).Fatal("failed to create low_level spi driver")
		}

		if err := raw.Init(); err != nil {
			log.WithError(err).Fatal("failed to initialize low_level spi driver")
		}

		d.LowLevel = raw
	}
}

// Start starts the sx1302 board
func (d *Dev) Start() error {
	if d.context.IsStarted {
		return errors.New("gateway is already running. Please stop it before changing configuration")
	}

	// ToDo: Support USB as well
	if d.context.BoardConfig.ComType != model.ComSPI {
		log.WithField("com_type", d.context.BoardConfig.ComType).Fatal("only SPI is supported as of now")
	}

	return nil
}
