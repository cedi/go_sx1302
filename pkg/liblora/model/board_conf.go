package model

// BoardConf structure for board specificities
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

func NewBoardConfig() *BoardConf {
	return &BoardConf{
		ComType:       ComSPI,
		ComPath:       "/dev/spidev0.0",
		LoRaWanPublic: true,
		ClkSrc:        0,
		FullDuplex:    false,
	}
}

// DebugConf is debug configuration for payload
type DebugConf struct {
	RefPayload  [16]Payload
	LogFileName string
}

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
