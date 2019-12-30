package space

var conversion map[Planet]float64 = map[Planet]float64{
	"Mercury": 7_600_543.81992,
	"Venus":   19_414_149.052176,
	"Earth":   31_557_600.0,
	"Mars":    59_354_032.690079994,
	"Jupiter": 374_355_659.124,
	"Saturn":  929_292_362.8848,
	"Uranus":  2_651_370_019.3296,
	"Neptune": 5_200_418_560.032001,
}

//Planet The name of a particular planet
type Planet string

// Age returns how old a person is given a planet and number of seconds
func Age(seconds float64, planet Planet) float64 {
	return seconds / conversion[planet]
}
