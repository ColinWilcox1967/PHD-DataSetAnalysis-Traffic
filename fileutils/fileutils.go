package fileutils

import (
	"fmt"
	"os"
	"time"
	"strings"

	"../trafficdata"

)

func CreateDumpFileName (algotype string) string {

	str :="algo_"+algotype+"_"
	str += time.Now().Format("2006-01-02-150405")
	str += ".csv"
	return strings.ToLower(str)
}

func DumpTrafficData (filename string, data []trafficdata.PimaDiabetesRecord) int {

	f, err := os.Create(filename)
	if err != nil {
		return 0
	}

	defer f.Close ()

	count := 0
	for i := 0; i< len(data); i++ {
		rec := data[i]
		str := fmt.Sprintf ("%s,%.3f,%.3f,%.3f,%.3f\n", rec.Timestamp, rec.NorthVolume, rec.NorthAverageSpeed, rec.SouthVolume, rec.SouthAverageSpeed)

		_, err := f.WriteString(str)
		if err == nil {
			count++
		}
	}

	

	return count
}