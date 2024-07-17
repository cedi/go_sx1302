package model

// BoardConf contains the configuration for the sx1302 board
type BoardConf struct {
	// Enable ONLY for *public* networks using the LoRa MAC protocol
	LoRaWanPublic bool

	// Index of RF chain which provides clock to concentrator
	ClkSrc uint8

	// Indicates if the gateway operates in full duplex mode or not
	FullDuplex bool

	// The COMmunication interface (SPI/USB) to connect to the SX1302
	ComType COMType

	// Path to access the COM device to connect to the SX1302 (max 64 chars)
	ComPath string
}

// NewBoardConfig creates a new BoardConf object with default parameters
func NewBoardConfig() *BoardConf {
	return &BoardConf{
		ComType:       ComSPI,
		ComPath:       "/dev/spidev0.0",
		LoRaWanPublic: true,
		ClkSrc:        0,
		FullDuplex:    false,
	}
}

// LgwContext context shared across modules
type LgwContext struct {
	IsStarted   bool
	BoardConfig *BoardConf

	RfChainCfg     []RxRf
	IfChainCfg     []RxIf
	DemodCfg       *Demod
	LoraServiceCfg *RxIf // LoRa service channel config parameters
	FSKCfg         *RxIf // FSK channel config parameters

	TxGainLUT []TxGainLUT

	FineTimestampCfg *FineTimeStampConf
	SX1261Cfg        *SX1261Conf

	DebugCfg *DebugConf
}

// NewLoraServiceCfg creates a new RxIf object with default configuration for the LgwContext.LoraServiceCfg
func NewLoraServiceCfg() *RxIf {
	return &RxIf{
		Enable:                false, // not used, handled by if_chain_cfg
		RFChain:               0,     // not used, handled by if_chain_cfg
		FreqHz:                0,     // not used, handled by if_chain_cfg
		Bandwidth:             Bw250kHz,
		Datarate:              DrLoraSf7,
		ImplicitHdr:           false,
		ImplicitPayloadLength: 0,
		ImplicitCrcEn:         false,
		ImplicitCoderate:      0,
	}
}

// NewFskCfg creates a new RxIf object with default configuration for the LgwContext.FSKCfg
func NewFskCfg() *RxIf {
	return &RxIf{
		Enable:       false, // not used, handled by if_chain_cfg
		RFChain:      0,     // not used, handled by if_chain_cfg
		FreqHz:       0,     // not used, handled by if_chain_cfg
		Bandwidth:    Bw250kHz,
		Datarate:     5 * DrFsk1kBaud,
		SyncWordSize: 3,
		SyncWord:     0xC194C1,
	}
}

// NewLgwContextWithDefaults creates a new lora gateway context with default values
func NewLgwContextWithDefaults() *LgwContext {
	return NewLgwContext(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
}

// NewLgwContext creates a new lora gateway context which is used to store the configuration used by the board
func NewLgwContext(boardCfg *BoardConf, rfChainCfg []RxRf, ifChainCfg []RxIf, demodCfg *Demod,
	loraServiceCfg *RxIf, fskCfg *RxIf, txGainLUT []TxGainLUT, fineTimestampCfg *FineTimeStampConf,
	sx1261cfg *SX1261Conf, debugCfg *DebugConf) *LgwContext {
	if boardCfg == nil {
		boardCfg = NewBoardConfig()
	}

	if len(rfChainCfg) == 0 {
		rfChainCfg = []RxRf{{Enable: false}}
	}

	if len(ifChainCfg) == 0 {
		ifChainCfg = []RxIf{{Enable: false}}
	}

	if demodCfg == nil {
		demodCfg = &Demod{MultisfDatarate: MultiSfEn}
	}

	if loraServiceCfg == nil {
		loraServiceCfg = NewLoraServiceCfg()
	}

	if fskCfg == nil {
		fskCfg = NewFskCfg()
	}

	if len(txGainLUT) == 0 {
		txGainLUT = []TxGainLUT{
			NewTxGainLUTWithDefaults(),
			NewTxGainLUTWithDefaults(),
		}
	}

	if fineTimestampCfg == nil {
		fineTimestampCfg = NewFineTimestampConf()
	}

	if sx1261cfg == nil {
		sx1261cfg = NewSX1261Conf()
	}

	if debugCfg == nil {
		debugCfg = &DebugConf{}
	}

	return &LgwContext{
		IsStarted:        false,
		BoardConfig:      boardCfg,
		RfChainCfg:       rfChainCfg,
		IfChainCfg:       ifChainCfg,
		DemodCfg:         demodCfg,
		LoraServiceCfg:   loraServiceCfg,
		FSKCfg:           fskCfg,
		TxGainLUT:        txGainLUT,
		FineTimestampCfg: fineTimestampCfg,
		SX1261Cfg:        sx1261cfg,
		DebugCfg:         debugCfg,
	}
}

// DebugConf is debug configuration for payload
type DebugConf struct {
	RefPayload  [16]Payload
	LogFileName string
}

// NewDebugConf creates a new DebugConf object with default parameters
func NewDebugConf() *DebugConf {
	return &DebugConf{
		LogFileName: "loragw.log",
		RefPayload:  [16]Payload{},
	}
}

// FineTimeStampConf for fine timestamping
type FineTimeStampConf struct {
	Enable bool                 //Enable / Disable fine timestamping
	Mode   FineTimestampingMode //Fine timestamping mode
}

// NewFineTimestampConf creates a new FineTimeStampConf object with default configuration
func NewFineTimestampConf() *FineTimeStampConf {
	return &FineTimeStampConf{
		Enable: false,
		Mode:   FineTsModeAllSf,
	}
}

// SX1261Conf is for additional SX1261 radio used for LBT and Spectral Scan
type SX1261Conf struct {
	// enable or disable SX1261 radio
	Enable bool

	// Path to access the SPI device to connect to the SX1261 (not used for USB com type)
	SpiPath string

	// value to be applied to the sx1261 RSSI value (dBm)
	RssiOffset int8

	// listen-before-talk configuration
	LbtConf LBTConf
}

// NewSX1261Conf creates a new SX1261Conf object with default configuration
func NewSX1261Conf() *SX1261Conf {
	return &SX1261Conf{
		Enable:     false,
		SpiPath:    "/dev/spidev0.1",
		RssiOffset: 0,
		LbtConf: LBTConf{
			Enable:    false,
			NbChannel: 0,
			Channels: []LBTChanConf{
				{
					FreqHz: 0,
				},
			},
		},
	}
}

// TComp contains all coefficients necessary to compute the offset to be applied on RSSI for current temperature
type TComp struct {
	CoeffA float32
	CoeffB float32
	CoeffC float32
	CoeffD float32
	CoeffE float32
}
