package algorithms

import (
	"../trafficdata"
	"../support"
	"fmt"
	"../logging"
)

// Algo=2
func replaceMissingValuesWithMean (dataset []trafficdata.PimaDiabetesRecord) ([]trafficdata.PimaDiabetesRecord, error) {

	numberOfFields := support.SizeOfPimaDiabetesRecord () - 1
	numberOfRecords := len(dataset)

	var resultSet = make([]trafficdata.PimaDiabetesRecord, numberOfRecords)

	// loop through and replace all missing elements with mean for the column

	// must be a simpler way to do this?????
	var columnTotal = make([]float64, numberOfFields)
	var columnMean = make([]float64, numberOfFields)

	for index := 0; index < numberOfRecords; index++ {
		columnTotal[0] += float64(dataset[index].NorthVolume)
		columnTotal[1] += float64(dataset[index].NorthAverageSpeed)
		columnTotal[2] += float64(dataset[index].SouthVolume)
		columnTotal[3] += float64(dataset[index].SouthAverageSpeed)
	
	}

	// work out means
	for index := 0; index < numberOfFields; index++ {
		// round up mean to n2 dp.
		columnMean[index] = support.RoundFloat64 (float64(columnTotal[index])/float64(numberOfRecords), 2)
	}

	// Dump all the column means
	for index := 0; index < numberOfFields; index++ {
		str := fmt.Sprintf ("Mean (%s) = %0.2f\n", textNameforColumn(index), columnMean[index])
		logging.DoWriteString (str, true, true)
	}

	// now sycle through the record and replace missing data with the mean for that column
	for index := 0; index < numberOfRecords; index++ {

		if dataset[index].NorthVolume == 0 {
			resultSet[index].NorthVolume = columnMean[0]
		} else {
			resultSet[index].NorthVolume = dataset[index].NorthVolume
		}
	
		if dataset[index].NorthAverageSpeed == 0 {
			resultSet[index].NorthAverageSpeed = columnMean[1]
		} else {
			resultSet[index].NorthAverageSpeed = dataset[index].NorthAverageSpeed
		}
	
		if dataset[index].SouthVolume == 0 {
			resultSet[index].SouthVolume = columnMean[2]
		} else {
			resultSet[index].SouthVolume = dataset[index].SouthVolume
		}

		if dataset[index].SouthAverageSpeed == 0 {
			resultSet[index].SouthAverageSpeed = columnMean[3]
		} else {
			resultSet[index].SouthAverageSpeed = dataset[index].SouthAverageSpeed
		}

		

		// TestedPositive field could actually be zero
		resultSet[index].Outcome = dataset[index].Outcome

	}

	return resultSet, nil
}

