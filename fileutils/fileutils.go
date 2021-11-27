package fileutils

import (
	"fmt"
	"os"

	"../trafficdata"

)
func DumpTrafficData (filename string, data []trafficdata.PimaDiabetesRecord) int {

	f, err := os.Create(filename)
	if err != nil {
		return 0
	}

	defer f.Close ()

	count := 0
	for i := 0; i< len(data); i++ {
		rec := data[i]
		str := fmt.Sprintf ("%s,%f,%f,%f,%f\n", rec.Timestamp, rec.NorthVolume, rec.NorthAverageSpeed, rec.SouthVolume, rec.SouthAverageSpeed)

		_, err := f.WriteString(str)
		if err == nil {
			count++
		}
	}

	f.Flush ()

	return count
}