package model

// RxRf configures a RF chain
type RxRf struct {
	// enable or disable that RF chain
	Enable bool

	// center frequency of the radio in Hz
	FreqHz uint32

	// Board-specific RSSI correction factor
	RssiOffset float32

	// Board-specific RSSI temperature compensation coefficients
	RssiTComp TComp

	// Radio type for that RF chain (SX1255, SX1257....)
	Type RadioType

	// enable or disable TX on that RF chain
	TxEnable bool

	// Configure the radio in single or differential input mode (SX1250 only)
	SingleInputMode bool
}

func NewRxRfConf() *RxRf {
	return &RxRf{
		Enable:          true,      // rf chain 0 needs to be enabled for calibration to work on sx1257
		FreqHz:          868500000, // dummy
		Type:            RadioTypeSX1250,
		TxEnable:        false,
		SingleInputMode: false,
	}
}

// RxIf configures an IF chain
type RxIf struct {
	// enable or disable that IF chain
	Enable bool

	// to which RF chain is that IF chain associated
	RFChain uint8

	// center frequ of the IF chain, relative to RF chain frequency
	FreqHz int32

	// RX bandwidth, 0 for default
	Bandwidth Bandwith

	// RX datarate, 0 for default
	Datarate DataRate

	// size of FSK sync word (number of bytes, 0 for default)
	SyncWordSize uint8

	// FSK sync word (ALIGN RIGHT, eg. 0xC194C1)
	SyncWord uint64

	// LoRa Service implicit header
	ImplicitHdr bool

	// LoRa Service implicit header payload length (number of bytes, 0 for default)
	ImplicitPayloadLength uint8

	// LoRa Service implicit header CRC enable
	ImplicitCrcEn bool

	// LoRa Service implicit header coding rate
	ImplicitCoderate uint8
}

// Demod configures the structure for LoRa/FSK demodulators
type Demod struct {
	// MultisfDatarate is the bitmask to enable spreading-factors for correlators (SF12 - SF5)
	MultisfDatarate uint8
}

// TXGain contains all gains of a Tx chain
type TXGain struct {
	// measured TX power at the board connector, in dBm
	RfPower int8

	// (sx125x) 2 bits: control of the digital gain of SX1302
	DigGain uint8

	// (sx125x) 2 bits: control of the external PA (SX1302 I/O) (sx1250) 1 bits: enable/disable the external PA (SX1302 I/O)
	PaGain uint8

	// (sx125x) 2 bits: control of the radio DAC
	DacGain uint8

	// (sx125x) 4 bits: control of the radio mixer
	MixGain uint8

	// (sx125x) calibrated I offset
	OffsetI int8

	// (sx125x) calibrated Q offset
	OffsetQ int8

	// (sx1250) 6 bits: control the radio power index to be used for configuration
	PwrIdx uint8
}

// NewTxGainWithDefaults creates a new TXGain object with default config
func NewTxGainWithDefaults() TXGain {
	return TXGain{
		RfPower: 14,
		DigGain: 0,
		PaGain:  2,
		DacGain: 3,
		MixGain: 10,
		OffsetI: 0,
		OffsetQ: 0,
		PwrIdx:  0,
	}
}

// TxGainLUT defines the Tx gain LUT
type TxGainLUT struct {
	// Array of Tx gain struct
	LUT []TXGain
}

// NewTxGainLUTWithDefaults creates a new TxGainLUT object with defaults
func NewTxGainLUTWithDefaults() TxGainLUT {
	return TxGainLUT{
		LUT: []TXGain{NewTxGainWithDefaults()},
	}
}

// LBTChanConf contains a Listen-Before-Talk channel configuration
type LBTChanConf struct {
	// LBT channel frequency
	FreqHz uint32

	// LBT channel bandwidth
	Bandwidth uint8

	// LBT channel carrier sense time
	ScanTimeUs ScanTime

	// LBT channel transmission duration when allowed
	TransmitTimeMs uint16
}

// LBTConf configures listen-before-talk
type LBTConf struct {
	// enable or disable LBT
	Enable bool

	// RSSI threshold to detect if channel is busy or not (dBm)
	RssiTarget int8

	// number of LBT channels
	NbChannel uint8

	// LBT channels configuration
	Channels []LBTChanConf
}
