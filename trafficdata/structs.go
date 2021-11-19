package trafficdata

//London Road traffic data
type PimaDiabetesRecord struct {

	Timestamp string
	NorthVolume,
	NorthAverageSpeed,
	SouthVolume,
	SouthAverageSpeed float64

	Outcome int // maybe should be a bool buit stored in file as int
}
// end of file
