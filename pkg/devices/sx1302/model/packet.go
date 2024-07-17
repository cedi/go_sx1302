package model

// Structure containing the metadata of a packet that was received and a pointer to the payload
type PktRx struct {
	FreqHz        uint32     // central frequency of the IF chain
	FreqOffset    int32      // frequency offset
	IfChain       uint8      // by which IF chain was packet received
	Status        uint8      // status of the received packet
	CountUs       uint32     // internal concentrator counter for timestamping, 1 microsecond resolution
	RfChain       uint8      // through which RF chain the packet was received
	ModemID       uint8      // Modem ID
	Modulation    uint8      // modulation used by the packet
	Bandwidth     uint8      // modulation bandwidth (LoRa only)
	Datarate      uint32     // RX datarate of the packet (SF for LoRa)
	Coderate      uint8      // error-correcting code of the packet (LoRa only)
	Rssic         float32    // average RSSI of the channel in dB
	Rssis         float32    // average RSSI of the signal in dB
	Snr           float32    // average packet SNR, in dB (LoRa only)
	SnrMin        float32    // minimum packet SNR, in dB (LoRa only)
	SnrMax        float32    // maximum packet SNR, in dB (LoRa only)
	Crc           uint16     // CRC that was received in the payload
	Size          uint16     // payload size in bytes
	Payload       [256]uint8 // buffer containing the payload
	FtimeReceived bool       // a fine timestamp has been received
	Ftime         uint32     // packet fine timestamp (nanoseconds since last PPS)
}

// Structure containing the configuration of a packet to send and a pointer to the payload
type PktTx struct {
	FreqHz     uint32     // center frequency of TX
	TxMode     uint8      // select on what event/time the TX is triggered
	CountUs    uint32     // timestamp or delay in microseconds for TX trigger
	RfChain    uint8      // through which RF chain will the packet be sent
	RfPower    int8       // TX power, in dBm
	Modulation uint8      // modulation to use for the packet
	FreqOffset int8       // frequency offset from Radio Tx frequency (CW mode)
	Bandwidth  uint8      // modulation bandwidth (LoRa only)
	Datarate   uint32     // TX datarate (baudrate for FSK, SF for LoRa)
	Coderate   uint8      // error-correcting code of the packet (LoRa only)
	InvertPol  bool       // invert signal polarity, for orthogonal downlinks (LoRa only)
	FDev       uint8      // frequency deviation, in kHz (FSK only)
	Preamble   uint16     // set the preamble length, 0 for default
	NoCrc      bool       // if true, do not send a CRC in the packet
	NoHeader   bool       // if true, enable implicit header mode (LoRa), fixed length (FSK)
	Size       uint16     // payload size in bytes
	Payload    [256]uint8 // buffer containing the payload
}

// Configuration structure for debug
type Payload struct {
	ID      uint32
	Payload [255]uint8
	PrevCnt uint32
}
