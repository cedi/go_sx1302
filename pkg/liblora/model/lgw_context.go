package model

// LgwContext context shared across modules
type LgwContext struct {
	// Global context
	IsStarted   bool
	BoardConfig *BoardConf

	// RX context
	RfChainCfg     []RxRf
	IfChainCfg     []RxIf
	DemodCfg       *Demod
	LoraServiceCfg *RxIf // LoRa service channel config parameters
	FSKCfg         *RxIf // FSK channel config parameters

	// TX context
	TxGainLUT []TxGainLUT

	// Misc
	FineTimestampCfg *FineTimeStampConf
	SX1261Cfg        *SX1261Conf

	// Debug
	DebugCfg *DebugConf
}

func NewLoraServiceCfg() *RxIf {
	return &RxIf{
		Enable:                false, // not used, handled by if_chain_cfg
		RFChain:               0,     // not used, handled by if_chain_cfg
		FreqHz:                0,     // not used, handled by if_chain_cfg
		Bandwidth:             Bw250KHz,
		Datarate:              DrLoraSf7,
		ImplicitHdr:           false,
		ImplicitPayloadLength: 0,
		ImplicitCrcEn:         false,
		ImplicitCoderate:      0,
	}
}

func NewFskCfg() *RxIf {
	return &RxIf{
		Enable:       false, // not used, handled by if_chain_cfg
		RFChain:      0,     // not used, handled by if_chain_cfg
		FreqHz:       0,     // not used, handled by if_chain_cfg
		Bandwidth:    Bw250KHz,
		Datarate:     DrFsk5k,
		SyncWordSize: 3,
		SyncWord:     0xC194C1,
	}
}

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
