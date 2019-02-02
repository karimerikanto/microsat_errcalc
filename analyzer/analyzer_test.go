package analyzer

import (
	"karim/microsatellite_analyzer/analyzer/models"
	"karim/microsatellite_analyzer/utils"
	"testing"
)

//TestAnalyzeWithAnEmptyArray tests that an empty result is returned if there is no samples given
func TestGetResultFromImportData_ToReturnNothing_WhenImportDataIsEmpty(t *testing.T) {
	importData := models.NewImportData()

	result, err := GetResultFromImportData(importData)

	test.AreEqual(t, nil, err, "Error was not null")
	test.AreEqual(t, 0, result.ReplicateAmount, "Replicate amount was not correct")
	test.AreEqual(t, 0, result.SampleAmount, "Sample amount was not correct")
	test.AreEqual(t, 0, result.SingleSampleAmount, "Single sample amount was not correct")
	test.AreEqual(t, 0, len(result.SampleResults), "Sample results amount was not correct")
}
