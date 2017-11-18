package space

//Planet Is the type expected by the Age function
type Planet string

//Age determines someone age as measured against earth time.
func Age(seconds float64, planet Planet) float64 {
	const stdEarthYearInSeconds float64 = 315581.4976
	switch planet {
	case "Earth":
		return float64(seconds/stdEarthYearInSeconds) / 100
	case "Mercury":
		return float64(seconds/(stdEarthYearInSeconds*0.2408467)) / 100
	case "Venus":
		return float64(seconds/(stdEarthYearInSeconds*0.61519726)) / 100
	case "Mars":
		return float64(seconds/(stdEarthYearInSeconds*1.8808158)) / 100
	case "Jupiter":
		return float64(seconds/(stdEarthYearInSeconds*11.862615)) / 100
	case "Saturn":
		return float64(seconds/(stdEarthYearInSeconds*29.447498)) / 100
	case "Uranus":
		return float64(seconds/(stdEarthYearInSeconds*84.016846)) / 100
	case "Neptune":
		return float64(seconds/(stdEarthYearInSeconds*164.79132)) / 100
	default:
		return 0.01
	}

}
