package algorithms

import (
	"../metrics"
	"../trafficdata"
	"../datasets"

	"os"

)

// algo=5

// using plain nearest neighbour removing incomplete data from the set of possible donors
func replaceNearestNeighbours (dataset []trafficdata.PimaDiabetesRecord) ([]trafficdata.PimaDiabetesRecord, error) {

	numberOfRecords := len(dataset)
	
	var resultSet = make([]trafficdata.PimaDiabetesRecord, numberOfRecords)
	var completeRecords = make([]trafficdata.PimaDiabetesRecord, numberOfRecords)

	// remove all incomplete records
	for index := 0; index < numberOfRecords; index++ {
		if !metrics.HasMissingElements (dataset[index]) {
			completeRecords = append (completeRecords, dataset[index])
		} 
	}

	// find nearest match for each test record
	numberOfTestRecords := len(datasets.PimaTestData)
	for index := 0; index < numberOfTestRecords; index++ {
		BuildSimilarityTable (datasets.PimaTestData[index])

		// take the closest neighbours from this table, calculate mean and replace

	}

	os.Exit(0)

	return resultSet, nil
}
