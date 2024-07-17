package model

const (
	// COMSuccess indicates success on COM
	COMSuccess int = 0
	// ComError indicates failure on COM
	ComError int = -1
	// SpiMuxTargetSX1302 define SX1302 target
	SpiMuxTargetSX1302 int = 0x00
	// SpiMuxTargetRadioA define generic RadioA
	SpiMuxTargetRadioA int = 0x01
	// SpiMuxTargetRadioB defines generic RadioB
	SpiMuxTargetRadioB int = 0x02
)

// Radio-specific parameters
const (
	// number of RF chains
	MaxRfChains uint8 = 2

	// Maximum size of Tx gain LUT
	MaxTxGainLutSize int = 16
)

// Listen-Before-Talk
const (
	// Maximum number of LBT channels
	LBTChannelCountMax int = 16
)

type COMType int

const (
	COMUnknown COMType = iota
	ComSPI
	ComUSB
)

func (c COMType) String() string {
	switch c {
	case ComSPI:
		return "SPI"
	case ComUSB:
		return "USB"
	}

	return "Unknown"
}

type COMWriteMode int

const (
	ComWriteModeUnknown COMWriteMode = iota
	ComWriteModeSingle
	ComWriteModeBulk
)

type RadioType int

const (
	RadioTypeNone RadioType = iota
	RadioTypeSX1255
	RadioTypeSX1257
	RadioTypeSX1272
	RadioTypeSX1276
	RadioTypeSX1250
)

func (r RadioType) String() string {
	switch r {
	case RadioTypeSX1255:
		return "SX1255"
	case RadioTypeSX1257:
		return "SX1257"
	case RadioTypeSX1272:
		return "SX1272"
	case RadioTypeSX1276:
		return "SX1276"
	case RadioTypeSX1250:
		return "SX1250"
	}

	return "Unknown"
}

// Fine timestamping modes
type FineTimestampingMode int

const (
	// fine timestamps for SF5 -> SF10
	FineTsModeHighCap FineTimestampingMode = iota

	// fine timestamps for SF5 -> SF12
	FineTsModeAllSf
)

// Radio types that can be found on the LoRa Gateway
type ScanTime int

const (
	ScanTime12Us   ScanTime = 128
	ScanTime5000Us ScanTime = 5000
)

// concentrator chipset-specific parameters
const (
	// number of IF+modem RX chains
	IFChains int = 10
	// bitmask to enable/disable SF for multi-sf correlators  (12 11 10 9 8 7 6 5)
	MultiSfEn uint8 = 0xFF
)

type SpectralScanStatus int

const (
	SpectralScanStatusUnknown SpectralScanStatus = iota
	SpectralScanStatusNone
	SpectralScanStatusOngoing
	SpectralScanStatusAborted
	SpectralScanStatusCompleted
)

func (s SpectralScanStatus) String() string {
	switch s {
	case SpectralScanStatusNone:
		return "None"
	case SpectralScanStatusOngoing:
		return "Ongoing"
	case SpectralScanStatusAborted:
		return "Aborted"
	case SpectralScanStatusCompleted:
		return "Completed"
	}

	return "Unknown"
}

// values available for the 'bandwidth' parameters (LoRa & FSK)
// NOTE: directly encode FSK RX bandwidth, do not change

type Bandwith uint8

const (
	BwUndefined Bandwith = 0
	Bw500KHz    Bandwith = 0x06
	Bw250KHz    Bandwith = 0x05
	Bw125KHz    Bandwith = 0x04
)

/* values available for the 'datarate' parameters */
/* NOTE: LoRa values used directly to code SF bitmask in 'multi' modem, do not change */
type DataRate uint32

const (
	DrUndefined DataRate = 0
	DrLoraSf5   DataRate = 5
	DrLoraSf6   DataRate = 6
	DrLoraSf7   DataRate = 7
	DrLoraSf8   DataRate = 8
	DrLoraSf9   DataRate = 9
	DrLoraSf10  DataRate = 10
	DrLoraSf11  DataRate = 11
	DrLoraSf12  DataRate = 12
	// NOTE: for FSK directly use baudrate between 500 bauds and 250 kbauds
	DrFskMin DataRate = 500
	DrFsk5k  DataRate = 50000
	DrFskMax DataRate = 250000
)

// Valid frequency range
const (
	RfRxFreqMin uint32 = 100e6
	RfRxFreqMax uint32 = 1e9
)
