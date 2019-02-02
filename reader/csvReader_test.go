package reader

import (
	"karim/microsatellite_analyzer/utils"
	"testing"
)

func TestReadCsvDataMatrixToImportData_ToReturnNothing_WhenInputMatrixIsEmpty(t *testing.T) {
	var dataLines [][]string

	importData, err := readCsvDataMatrixToImportData(dataLines)

	test.AreEqual(t, nil, err, "Error was not null")
	test.AreEqual(t, 0, len(importData.Headers), "Wrong amount of headers")
	test.AreEqual(t, 0, len(importData.Samples), "Wrong amount of samples")
	test.AreEqual(t, 0, len(importData.LociGroups), "Wrong amount of loci groups")
}

func TestReadCsvDataMatrixToImportData_ToReturnOnlyHeaders_WhenInputMatrixContainsOnlyHeaders(t *testing.T) {
	dataLines := [][]string{
		{"Header 1", "continues"},
		{"Header 2", "continues"},
	}

	importData, err := readCsvDataMatrixToImportData(dataLines)

	test.AreEqual(t, nil, err, "Error was not null")

	test.AreEqual(t, 2, len(importData.Headers), "Wrong amount of headers")
	test.AreEqual(t, "Header 1", importData.Headers[0][0], "Header 1 (cell 1) was invalid")
	test.AreEqual(t, "continues", importData.Headers[0][1], "Header 1 (cell 2) was invalid")
	test.AreEqual(t, "Header 2", importData.Headers[1][0], "Header 2 (cell 1) was invalid")
	test.AreEqual(t, "continues", importData.Headers[1][1], "Header 2 (cell 2) was invalid")

	test.AreEqual(t, 0, len(importData.Samples), "Wrong amount of samples")
	test.AreEqual(t, 0, len(importData.LociGroups), "Wrong amount of loci groups")
}

func TestReadCsvDataMatrixToImportData_ToReturnHeadersAndTwoSamples_WhenInputMatrixContainsHeadersAndFiveReplicaRowsDividedByEmptyLine(t *testing.T) {
	dataLines := [][]string{
		{"Header 1", "continues"},
		{"Header 2", "continues"},
		{""},
		{"", "Locus 1", "", "Locus 2"},
		{"Replica 1", "100", "120", "200", "210"},
		{"Replica 2", "100", "130", "?", "240"},
		{""},
		{"Replica 1", "100", "120", "200", "210"},
		{"Replica 2", "100", "130", "", "240"},
		{"Replica 3", "120", "130", "?", "210"},
	}

	importData, err := readCsvDataMatrixToImportData(dataLines)

	test.AreEqual(t, nil, err, "Error was not null")

	test.AreEqual(t, 2, len(importData.Headers), "Wrong amount of headers")
	test.AreEqual(t, "Header 1", importData.Headers[0][0], "Header 1 (cell 1) was invalid")
	test.AreEqual(t, "continues", importData.Headers[0][1], "Header 1 (cell 2) was invalid")
	test.AreEqual(t, "Header 2", importData.Headers[1][0], "Header 2 (cell 1) was invalid")
	test.AreEqual(t, "continues", importData.Headers[1][1], "Header 2 (cell 2) was invalid")

	test.AreEqual(t, 2, len(importData.Samples), "Wrong amount of samples")

	//Sample 1
	sample := importData.Samples[0]

	test.AreEqual(t, 2, len(sample.ReplicaArray), "Wrong amount of replicas in sample 1")

	//Replica 1
	replica1 := sample.ReplicaArray[0]

	test.AreEqual(t, "Replica 1", replica1.Name, "Wrong replica name for replica 1 (sample 1)")
	test.AreEqual(t, 2, len(replica1.LocusArray), "Wrong amount of loci in replica 1 (sample 1)")

	test.AreEqual(t, "Locus 1", replica1.LocusArray[0].Name, "Wrong name for first locus in replica 1 (sample 1)")
	test.AreEqual(t, "100", replica1.LocusArray[0].Allele1, "Wrong allele 1 value for first locus in replica 1 (sample 1)")
	test.AreEqual(t, "120", replica1.LocusArray[0].Allele2, "Wrong allele 2 value for first locus in replica 1 (sample 1)")

	test.AreEqual(t, "Locus 2", replica1.LocusArray[1].Name, "Wrong name for second locus in replica 1 (sample 1)")
	test.AreEqual(t, "200", replica1.LocusArray[1].Allele1, "Wrong allele 1 value for second locus in replica 1 (sample 1)")
	test.AreEqual(t, "210", replica1.LocusArray[1].Allele2, "Wrong allele 2 value for second locus in replica 1 (sample 1)")

	//Replica 2
	replica2 := sample.ReplicaArray[1]

	test.AreEqual(t, "Replica 2", replica2.Name, "Wrong replica name for replica 2 (sample 1)")
	test.AreEqual(t, 2, len(replica2.LocusArray), "Wrong amount of loci in replica 2 (sample 1)")

	test.AreEqual(t, "Locus 1", replica2.LocusArray[0].Name, "Wrong name for first locus in replica 2 (sample 1)")
	test.AreEqual(t, "100", replica2.LocusArray[0].Allele1, "Wrong allele 1 value for first locus in replica 2 (sample 1)")
	test.AreEqual(t, "130", replica2.LocusArray[0].Allele2, "Wrong allele 2 value for first locus in replica 2 (sample 1)")

	test.AreEqual(t, "Locus 2", replica2.LocusArray[1].Name, "Wrong name for second locus in replica 2 (sample 1)")
	test.AreEqual(t, "?", replica2.LocusArray[1].Allele1, "Wrong allele 1 value for second locus in replica 2 (sample 1)")
	test.AreEqual(t, "240", replica2.LocusArray[1].Allele2, "Wrong allele 2 value for second locus in replica 2 (sample 1)")

	//Sample 2
	sample = importData.Samples[1]

	test.AreEqual(t, 3, len(sample.ReplicaArray), "Wrong amount of replicas in sample 2")

	//Replica 1
	replica1 = sample.ReplicaArray[0]

	test.AreEqual(t, "Replica 1", replica1.Name, "Wrong replica name for replica 1 (sample 2)")
	test.AreEqual(t, 2, len(replica1.LocusArray), "Wrong amount of loci in replica 1 (sample 2)")

	test.AreEqual(t, "Locus 1", replica1.LocusArray[0].Name, "Wrong name for first locus in replica 1 (sample 2)")
	test.AreEqual(t, "100", replica1.LocusArray[0].Allele1, "Wrong allele 1 value for first locus in replica 1 (sample 2)")
	test.AreEqual(t, "120", replica1.LocusArray[0].Allele2, "Wrong allele 2 value for first locus in replica 1 (sample 2)")

	test.AreEqual(t, "Locus 2", replica1.LocusArray[1].Name, "Wrong name for second locus in replica 1 (sample 2)")
	test.AreEqual(t, "200", replica1.LocusArray[1].Allele1, "Wrong allele 1 value for second locus in replica 1 (sample 2)")
	test.AreEqual(t, "210", replica1.LocusArray[1].Allele2, "Wrong allele 2 value for second locus in replica 1 (sample 2)")

	//Replica 2
	replica2 = sample.ReplicaArray[1]

	test.AreEqual(t, "Replica 2", replica2.Name, "Wrong replica name for replica 2 (sample 2)")
	test.AreEqual(t, 2, len(replica2.LocusArray), "Wrong amount of loci in replica 2 (sample 2)")

	test.AreEqual(t, "Locus 1", replica2.LocusArray[0].Name, "Wrong name for first locus in replica 2 (sample 2)")
	test.AreEqual(t, "100", replica2.LocusArray[0].Allele1, "Wrong allele 1 value for first locus in replica 2 (sample 2)")
	test.AreEqual(t, "130", replica2.LocusArray[0].Allele2, "Wrong allele 2 value for first locus in replica 2 (sample 2)")

	test.AreEqual(t, "Locus 2", replica2.LocusArray[1].Name, "Wrong name for second locus in replica 2 (sample 2)")
	test.AreEqual(t, "", replica2.LocusArray[1].Allele1, "Wrong allele 1 value for second locus in replica 2 (sample 2)")
	test.AreEqual(t, "240", replica2.LocusArray[1].Allele2, "Wrong allele 2 value for second locus in replica 2 (sample 2)")

	//Replica 3
	replica3 := sample.ReplicaArray[2]

	test.AreEqual(t, "Replica 3", replica3.Name, "Wrong replica name for replica 3 (sample 2)")
	test.AreEqual(t, 2, len(replica3.LocusArray), "Wrong amount of loci in replica 3 (sample 2)")

	test.AreEqual(t, "Locus 1", replica3.LocusArray[0].Name, "Wrong name for first locus in replica 3 (sample 2)")
	test.AreEqual(t, "120", replica3.LocusArray[0].Allele1, "Wrong allele 1 value for first locus in replica 3 (sample 2)")
	test.AreEqual(t, "130", replica3.LocusArray[0].Allele2, "Wrong allele 2 value for first locus in replica 3 (sample 2)")

	test.AreEqual(t, "Locus 2", replica3.LocusArray[1].Name, "Wrong name for second locus in replica 3 (sample 2)")
	test.AreEqual(t, "?", replica3.LocusArray[1].Allele1, "Wrong allele 1 value for second locus in replica 3 (sample 2)")
	test.AreEqual(t, "210", replica3.LocusArray[1].Allele2, "Wrong allele 2 value for second locus in replica 3 (sample 2)")

	// Loci groups
	test.AreEqual(t, 2, len(importData.LociGroups), "Wrong amount of loci groups")
	test.AreEqual(t, 5, len(importData.LociGroups["Locus 1"]), "Wrong amount loci in loci group 'Locus 1'")
	test.AreEqual(t, 5, len(importData.LociGroups["Locus 2"]), "Wrong amount loci in loci group 'Locus 2'")
}

func TestReadCsvDataMatrixToImportData_ToReturnNoHeadersAndOneSample_WhenInputMatrixContainsNoHeadersAndTwoReplicaRows(t *testing.T) {
	dataLines := [][]string{
		{"", "Locus 1", "", "Locus 2"},
		{"Replica 1", "100", "120", "200", "210"},
		{"Replica 2", "100", "130", "?", "240"},
	}

	importData, err := readCsvDataMatrixToImportData(dataLines)

	test.AreEqual(t, nil, err, "Error was not null")

	test.AreEqual(t, 0, len(importData.Headers), "Wrong amount of headers")
	test.AreEqual(t, 1, len(importData.Samples), "Wrong amount of samples")

	sample := importData.Samples[0]

	test.AreEqual(t, 2, len(sample.ReplicaArray), "Wrong amount of replicas in sample")

	//Replica 1
	replica1 := sample.ReplicaArray[0]

	test.AreEqual(t, "Replica 1", replica1.Name, "Wrong replica name for replica 1")
	test.AreEqual(t, 2, len(replica1.LocusArray), "Wrong amount of loci in replica 1")

	test.AreEqual(t, "Locus 1", replica1.LocusArray[0].Name, "Wrong name for first locus in replica 1")
	test.AreEqual(t, "100", replica1.LocusArray[0].Allele1, "Wrong allele 1 value for first locus in replica 1")
	test.AreEqual(t, "120", replica1.LocusArray[0].Allele2, "Wrong allele 2 value for first locus in replica 1")

	test.AreEqual(t, "Locus 2", replica1.LocusArray[1].Name, "Wrong name for second locus in replica 1")
	test.AreEqual(t, "200", replica1.LocusArray[1].Allele1, "Wrong allele 1 value for second locus in replica 1")
	test.AreEqual(t, "210", replica1.LocusArray[1].Allele2, "Wrong allele 2 value for second locus in replica 1")

	//Replica 2
	replica2 := sample.ReplicaArray[1]

	test.AreEqual(t, "Replica 2", replica2.Name, "Wrong replica name for replica 2")
	test.AreEqual(t, 2, len(replica2.LocusArray), "Wrong amount of loci in replica 2")

	test.AreEqual(t, "Locus 1", replica2.LocusArray[0].Name, "Wrong name for first locus in replica 2")
	test.AreEqual(t, "100", replica2.LocusArray[0].Allele1, "Wrong allele 1 value for first locus in replica 2")
	test.AreEqual(t, "130", replica2.LocusArray[0].Allele2, "Wrong allele 2 value for first locus in replica 2")

	test.AreEqual(t, "Locus 2", replica2.LocusArray[1].Name, "Wrong name for second locus in replica 2")
	test.AreEqual(t, "?", replica2.LocusArray[1].Allele1, "Wrong allele 1 value for second locus in replica 2")
	test.AreEqual(t, "240", replica2.LocusArray[1].Allele2, "Wrong allele 2 value for second locus in replica 2")
}

func TestReadCsvDataMatrixToImportData_ToReturnError_WhenInputMatrixContainsWrongAmountOfSampleColumnsAndLociNames(t *testing.T) {
	dataLines := [][]string{
		{"", "Locus 1", "", "Locus 2"},
		{"Replica 1", "100", "120", "200", "210", "100", "100"},
	}

	importData, err := readCsvDataMatrixToImportData(dataLines)

	test.AreEqual(t, "Locus index isn't matching with the locus name indexes on line 1", err.Error(), "Error messagewas incorrect")

	test.AreEqual(t, 0, len(importData.Headers), "Wrong amount of headers")
	test.AreEqual(t, 0, len(importData.Samples), "Wrong amount of samples")
}
