package metrics

import(
	"math"
	"fmt"
	"../support"
	"../trafficdata"
	"../logging"
)

func HasMissingElements (record trafficdata.PimaDiabetesRecord) bool {
	return countMissingElements (record) > 0
}

func countMissingElements (record trafficdata.PimaDiabetesRecord) int {

	missingFieldCount := 0

	// for each zero value increment count
	if record.NorthVolume == 0{
		missingFieldCount++
	}
	if record.NorthAverageSpeed == 0 {
		missingFieldCount++
	}
	if record.SouthVolume == 0 {
		missingFieldCount++
	}
	if record.SouthAverageSpeed == 0 {
		missingFieldCount++
	}
	
	
	return missingFieldCount
}

func ShowDataSetStatistics (displayName string, metrics DataSetMetrics) {

	str := fmt.Sprintf ("%s : ", displayName)
	
	sparsity := float64(100.0*metrics.NumberOfMissingElements)/float64(metrics.Size)

	str += fmt.Sprintf ("Sparsity = %.2f%% (%d out of %d elements)\n", sparsity, metrics.NumberOfMissingElements, metrics.Size)
	logging.DoWriteString (str, true,true)
}

func SourceDataSetMetricsByFeature (dataset []trafficdata.PimaDiabetesRecord) []DataSetStatisticsRecord {
	recordCount := len(dataset)
	numberOfFields := 7 // hardcode for now

	statistics := make([]DataSetStatisticsRecord, numberOfFields) // Use same indexes as per Pima Diabetes record definition

	// Field : RecordCount
	for feature := 0; feature < numberOfFields; feature++ {
		statistics[feature].RecordCount = recordCount
	}

	// Pima Diabetes Database, all fields are numeric
	//0. Number of times pregnant.
	//1. Plasma glucose concentration a 2 hours in an oral glucose tolerance test.
	//2. Diastolic blood pressure (mm Hg).
	//3. Triceps skinfold thickness (mm).
	//4. 2-Hour serum insulin (mu U/ml).
	//5. Body mass index (weight in kg/(height in m)^2).
	//6. Diabetes pedigree function.
	//7. Age (years).
	//8. Class variable (0 or 1).
	
	totalFeatureCount := make([]float64, numberOfFields) // used for SD and mean values

	// Upper and Lower feature values
	for index := 0; index < recordCount;index++ {
	
		for feature := 0; feature < numberOfFields; feature++ {

			var value float64

			totalFeatureCount[feature] += value // add cell value to appropriate feature count

			switch (feature) {
				case 0: // Volume north
					value = float64(dataset[index].NorthVolume)
				case 1: //Average speed north
					value = float64(dataset[index].NorthAverageSpeed)
				case 2: // Volume south
					value = float64(dataset[index].SouthVolume)
				case 3: //Average speed south
					value = float64(dataset[index].SouthAverageSpeed)
			
			}

			if index == 0 {
				statistics[feature].Lowest = value
				statistics[feature].Highest = value
			} else {
				if value < statistics[feature].Lowest {
					statistics[feature].Lowest = value
				}
				if value > statistics[feature].Highest {
					statistics[feature].Highest = value
				}
			}
		}
	}

	// Mean 
	for feature := 0; feature < numberOfFields; feature++ {
		statistics[feature].Mean = totalFeatureCount[feature]/float64(recordCount)
	}

	// Standard Deviation
	for feature := 0; feature < numberOfFields; feature++ {
		mean := statistics[feature].Mean
		totalDeviation := 0.0

		var value float64
		for index := 0; index < recordCount; index++ {
			switch (feature) {
				case 0: // North volume
					value = float64(dataset[index].NorthVolume)
				case 1: //average speed north
					value = float64(dataset[index].NorthAverageSpeed)
				case 2: // south volume
					value = float64(dataset[index].SouthVolume)
				case 3: // Taverage speed south
					value = float64(dataset[index].SouthAverageSpeed)
				
			}

			totalDeviation += (value-mean)*(value-mean)
		}

		statistics[feature].StandardDeviation = math.Sqrt(totalDeviation/float64(recordCount))
	}

	return statistics
}

func GetDataSetMetrics (dataset []trafficdata.PimaDiabetesRecord) DataSetMetrics {

	var metrics DataSetMetrics
	
	numberOfFields := support.SizeOfPimaDiabetesRecord ()

	metrics.Size = len(dataset) * numberOfFields 
	metrics.NumberOfMissingElements = 0

	// loop round finding missing elements for each record. fixed for now
	for index := 0; index < len(dataset); index++ {
		metrics.NumberOfMissingElements += countMissingElements (dataset[index])
	}

	return (metrics)
}

// end of file
