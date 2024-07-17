package model

// Structure containing all coefficients necessary to compute the offset to be applied on RSSI for current temperature
type TComp struct {
	CoeffA float32
	CoeffB float32
	CoeffC float32
	CoeffD float32
	CoeffE float32
}
