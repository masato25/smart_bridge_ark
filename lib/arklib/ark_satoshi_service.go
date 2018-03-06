package arklib

const SatoshiPerArk = float64(100000000)

func ToSatoshi(arkMount int64) {

}

func ToArk(satoshis int64) float64 {
	return float64(satoshis) / SatoshiPerArk
}

func ToArkFalt(satoshis float64) float64 {
	return satoshis / SatoshiPerArk
}
