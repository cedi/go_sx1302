package model

import "fmt"

const (
	// SpiMuxTargetSX1302 define SX1302 target
	SpiMuxTargetSX1302 uint8 = 0x00

	// SpiMuxTargetRadioA define generic RadioA
	SpiMuxTargetRadioA uint8 = 0x01

	// SpiMuxTargetRadioB defines generic RadioB
	SpiMuxTargetRadioB uint8 = 0x02

	// MaxIFChains is the number of IF+modem RX chains
	MaxIFChains int = 10

	// MaxRfChains is the max number of RF chains
	MaxRfChains uint8 = 2

	// MaxTxGainLutSize is the maximum size of Tx gain LUT
	MaxTxGainLutSize int = 16

	// LBTChannelCountMax is the maximum number of Listen-Before-Talk channels
	LBTChannelCountMax int = 16

	// MultiSfEn is the bitmask to enable/disable SF for multi-sf correlators (12 11 10 9 8 7 6 5)
	MultiSfEn uint8 = 0xFF

	// RfRxFreqMin is the lower bound of valid frequency range for LoRa
	RfRxFreqMin uint32 = 100e6

	// RfRxFreqMax is the upper bound of valid frequency range for LoRa
	RfRxFreqMax uint32 = 1e9
)

// COMType is the type which shall be used for communication with the board
type COMType int

const (
	// ComSPI is to use SPI to communicate with the board
	ComSPI COMType = iota
	// ComUSB is to use USB to communicate with the board. NOTE: This is currently unsupported
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

// COMWriteMode configures the write-mode for communication with the board
type COMWriteMode int

const (
	// ComWriteModeSingle single writes
	ComWriteModeSingle COMWriteMode = iota
	// ComWriteModeBulk bulk mode
	ComWriteModeBulk
)

func (c COMWriteMode) String() string {
	switch c {
	case ComWriteModeSingle:
		return "Single"
	case ComWriteModeBulk:
		return "Bulk"
	}

	return "Unknown"
}

// RadioType is to set which radio is being used
type RadioType int

const (
	// RadioTypeSX1255 specifies a SX1255 baord
	RadioTypeSX1255 RadioType = iota

	// RadioTypeSX1257 specifies a SX1257 baord
	RadioTypeSX1257

	// RadioTypeSX1272 specifies a SX1272 baord
	RadioTypeSX1272

	// RadioTypeSX1276 specifies a SX1276 baord
	RadioTypeSX1276

	// RadioTypeSX1250 specifies a SX1250 baord
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

// FineTimestampingMode configures timestamping
type FineTimestampingMode int

const (
	// FineTsModeHighCap is fine timestamps for SF5 -> SF10
	FineTsModeHighCap FineTimestampingMode = iota

	// FineTsModeAllSf is fine timestamps for SF5 -> SF12
	FineTsModeAllSf
)

func (f FineTimestampingMode) String() string {
	switch f {
	case FineTsModeHighCap:
		return "fine timestamps for SF5 -> SF10"

	case FineTsModeAllSf:
		return "fine timestamps for SF5 -> SF12"
	}

	return "Unknown"
}

// ScanTime is the channel carrier sense time
type ScanTime uint16

const (
	// ScanTime12Us is 12Us
	ScanTime12Us ScanTime = 128

	// ScanTime5000Us is 5000Us / 5ms
	ScanTime5000Us ScanTime = 5000
)

func (s ScanTime) String() string {
	switch s {
	case ScanTime12Us:
		return "12Us"
	case ScanTime5000Us:
		return "5000Us"
	}

	return fmt.Sprintf("%dus", uint16(s))
}

// SpectralScanStatus is to capture the current status of a spectral scan
type SpectralScanStatus int

const (
	// SpectralScanStatusNone means the scan was not started before
	SpectralScanStatusNone SpectralScanStatus = iota
	// SpectralScanStatusOngoing means the scan is ongoing
	SpectralScanStatusOngoing
	// SpectralScanStatusAborted means the scan is aborted
	SpectralScanStatusAborted
	// SpectralScanStatusCompleted means the scan is completed
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

// Bandwith is the values available for the 'bandwidth' parameters (LoRa & FSK)
// NOTE: directly encode FSK RX bandwidth, do not change
type Bandwith uint8

const (
	// Bw500kHz means 500kHz
	Bw500kHz Bandwith = 0x06
	// Bw250kHz means 250kHz
	Bw250kHz Bandwith = 0x05
	// Bw125kHz means 125kHz
	Bw125kHz Bandwith = 0x04
)

func (b Bandwith) String() string {
	switch b {
	case Bw500kHz:
		return "500kHz"
	case Bw250kHz:
		return "250kHz"
	case Bw125kHz:
		return "125kHz"
	}

	return "Undefined"
}

// DataRate is the values available for the 'datarate' parameters
// NOTE: LoRa values used directly to code SF bitmask in 'multi' modem, do not change
type DataRate uint32

const (
	// DrLoraSf5 means a spreadfactor of 5
	DrLoraSf5 DataRate = 5
	// DrLoraSf6 means a spreadfactor of 6
	DrLoraSf6 DataRate = 6
	// DrLoraSf7 means a spreadfactor of 7
	DrLoraSf7 DataRate = 7
	// DrLoraSf8 means a spreadfactor of 8
	DrLoraSf8 DataRate = 8
	// DrLoraSf9 means a spreadfactor of 9
	DrLoraSf9 DataRate = 9
	// DrLoraSf10 means a spreadfactor of 10
	DrLoraSf10 DataRate = 10
	// DrLoraSf11 means a spreadfactor of 11
	DrLoraSf11 DataRate = 11
	// DrLoraSf12 means a spreadfactor of 12
	DrLoraSf12 DataRate = 12
	// DrFskMin means a spreadfactor of 500 bauds. NOTE: for FSK directly use baudrate between 500 bauds and 250 kbauds
	DrFskMin DataRate = 500
	// DrFsk1Baud specifies the baud rate of 1, used to create all baud-rates in between
	DrFsk1Baud DataRate = 1
	// DrFsk1kBaud specifies the baud rate of 1000, used to create all baud-rates in between
	DrFsk1kBaud DataRate = 1000 * DrFsk1Baud
	// DrFskMax means a spreadfactor of 250 kilo bauds. NOTE: for FSK directly use baudrate between 500 bauds and 250 kbauds
	DrFskMax DataRate = 250000
)

func (d DataRate) String() string {
	switch d {
	case DrLoraSf5:
		return "SF_5"
	case DrLoraSf6:
		return "SF_6"
	case DrLoraSf7:
		return "SF_7"
	case DrLoraSf8:
		return "SF_8"
	case DrLoraSf9:
		return "SF_9"
	case DrLoraSf10:
		return "SF_10"
	case DrLoraSf11:
		return "SF_11"
	case DrLoraSf12:
		return "SF_12"
	}

	baudRate := fmt.Sprintf("%d Bd", uint32(d))

	switch d {
	case DrFskMin:
		baudRate += " (FSK_MIN)"
	case DrFskMax:
		baudRate += " (FSK_MAX)"
	}

	return baudRate
}
