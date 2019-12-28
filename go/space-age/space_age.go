package space

const earthYearInSeconds float64 = 31557600

func toEarthYears(sec float64) float64 {
	return sec / earthYearInSeconds
}

//Planet The name of a particular planet
type Planet string

func convertFromEarth(years *float64, planet Planet) {
	convertion := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	*years /= convertion[planet]
}

// Age returns how old a person is given a planet and number of seconds
func Age(seconds float64, planet Planet) float64 {
	age := toEarthYears(seconds)
	if planet != "Earth" {
		convertFromEarth(&age, planet)
	}
	return age
}
