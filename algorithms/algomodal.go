package algorithms

import (
	"sort"
	"fmt"

	"../trafficdata"
	"../support"
	"../logging"
)

type valueCount struct {
	Value float64
	Count int
}

// just checks if value already exists in the list for this feature
func valueExistsForFeature (list []valueCount, value float64) (bool, int) {
	for i := 0; i < len(list); i++ {
		if list[i].Value == value {
			return true, i
		}
	}

	return false, -1
}

//algo=3
func replaceMissingValuesWithModal (dataset []trafficdata.PimaDiabetesRecord) ([]trafficdata.PimaDiabetesRecord, error) {
	numberOfFields := support.SizeOfPimaDiabetesRecord () - 1
	numberOfRecords := len(dataset)

	var resultSet = make([]trafficdata.PimaDiabetesRecord, numberOfRecords)

	columnCount := make([][]valueCount, numberOfFields)
	columnModal := make([]valueCount, numberOfFields)

	for index := 0; index < numberOfRecords; index++ {
		r := dataset[index]

		var v valueCount
		var pos int
		var exists bool
		var value float64
		
		for field := 0; field < numberOfFields; field++ {

			switch field {
				case 0: value = r.NorthVolume
				case 1: value = r.NorthAverageSpeed
				case 2: value = r.SouthVolume
				case 3: value = r.SouthAverageSpeed
				
			}

			exists, pos = valueExistsForFeature (columnCount[field], value)
		
			if !exists {
				v.Count = 1
				v.Value = value
				columnCount[field] = append(columnCount[field], v)
			} else {
				columnCount[field][pos].Count++
			}
		}
	}

	// done all the counts. need to find modal value for each column
	for field := 0; field < numberOfFields; field++ {
		sort.Slice(columnCount[field][:], 
					func(i, j int) bool {
					return columnCount[field][i].Count > columnCount[field][j].Count})
		
		// select first non missing value for mode

		if columnCount[field][0].Value == 0 { // can used a gap as modal value
			columnModal[field].Value = columnCount[field][1].Value
		} else {
			columnModal[field].Value = columnCount[field][0].Value
		}
	}

	// Dump all the column modal values
	for index := 0; index < numberOfFields; index++ {
		str := fmt.Sprintf ("Modal (%s) = %0.2f\n", textNameforColumn(index), columnModal[index].Value)
	
		logging.DoWriteString (str, true, true)
	}
	// now we have the modal for each columm run through and process the data set
	
	for index:= 0; index < numberOfRecords; index++ {
		if dataset[index].NorthVolume == 0 {
			resultSet[index].NorthVolume = support.RoundFloat64 (columnModal[0].Value,2)
		} else {
			resultSet[index].NorthVolume = support.RoundFloat64 (dataset[index].NorthVolume, 2)
		}
	
		if dataset[index].NorthAverageSpeed == 0 {
			resultSet[index].NorthAverageSpeed = support.RoundFloat64 (columnModal[1].Value, 2)
		} else {
			resultSet[index].NorthAverageSpeed = support.RoundFloat64(dataset[index].NorthAverageSpeed, 2)
		}
	
		if dataset[index].SouthVolume == 0 {
			resultSet[index].SouthVolume = support.RoundFloat64(columnModal[2].Value, 2)
		} else {
			resultSet[index].SouthVolume = support.RoundFloat64(dataset[index].SouthVolume, 2)
		}

		if dataset[index].SouthAverageSpeed == 0 {
			resultSet[index].SouthAverageSpeed = support.RoundFloat64(columnModal[3].Value, 2)
		} else {
			resultSet[index].SouthAverageSpeed = support.RoundFloat64(dataset[index].SouthAverageSpeed, 2)
		}

		resultSet[index].Timestamp = dataset[index].Timestamp


		
		// TestedPositive field may actually be zero
		resultSet[index].Outcome = dataset[index].Outcome
	}

	return resultSet,nil
}
