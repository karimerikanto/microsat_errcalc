package analyzer

import (
	"karim/microsat_errcalc/analyzer/models"
	"karim/microsat_errcalc/utils"
	"testing"
)

// clear; go test ./... -v

func TestGetResultFromImportData_ToReturnNothing_WhenImportDataIsEmpty(t *testing.T) {
	importData := models.NewImportData()

	result, err := GetResultFromImportData(importData)

	test.AreEqual(t, nil, err, "Error was not null")
	test.AreEqual(t, 0, result.ReplicateAmount, "Replicate amount was incorrect")
	test.AreEqual(t, 0, result.SampleAmount, "Sample amount was incorrect")
	test.AreEqual(t, 0, result.SingleSampleAmount, "Single sample amount was incorrect")
	test.AreEqual(t, 0, result.AmountOfAlleles, "Amount of alleles was incorrect")
	test.AreEqual(t, 0, result.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 0, result.AmountOfLoci, "Amount of loci was incorrect")
	test.AreEqual(t, 0, result.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 0, len(result.LociResults), "Amount of loci results was incorrect")
	test.AreEqual(t, 0, len(result.SampleResults), "Sample results amount was incorrect")
	test.AreEqual(t, 0, len(result.LociNamesInOrder), "Amount of loci orders was incorrect")
}

func TestGetResultFromImportData_ToReturnValidData_WhenImportDataContainsTwoSampleIncludingSevenReplicas(t *testing.T) {
	importData := models.NewImportData()

	importData.Samples = []models.Sample{
		models.Sample{
			ReplicaArray: []models.Replica{
				models.Replica{
					Name: "Replica 1",
					LocusArray: []models.Locus{
						models.Locus{
							Name:    "Afa05",
							Allele1: "182", Allele2: "194",
						},
						models.Locus{
							Name:    "Afa13",
							Allele1: "219", Allele2: "227",
						},
						models.Locus{
							Name:    "Afa15",
							Allele1: "239", Allele2: "243",
						},
					},
				},
				models.Replica{
					Name: "Replica 2",
					LocusArray: []models.Locus{
						models.Locus{
							Name:    "Afa05",
							Allele1: "182", Allele2: "194",
						},
						models.Locus{
							Name:    "Afa13",
							Allele1: "219", Allele2: "227",
						},
						models.Locus{
							Name:    "Afa15",
							Allele1: "239", Allele2: "243",
						},
					},
				},
				models.Replica{
					Name: "Replica 3",
					LocusArray: []models.Locus{
						models.Locus{
							Name:    "Afa05",
							Allele1: "182", Allele2: "194",
						},
						models.Locus{
							Name:    "Afa13",
							Allele1: "219", Allele2: "222",
						},
						models.Locus{
							Name:    "Afa15",
							Allele1: "239", Allele2: "243",
						},
					},
				},
			},
		},
		models.Sample{
			ReplicaArray: []models.Replica{
				models.Replica{
					Name: "Replica 1",
					LocusArray: []models.Locus{
						models.Locus{
							Name:    "Afa05",
							Allele1: "182", Allele2: "194",
						},
						models.Locus{
							Name:    "Afa13",
							Allele1: "219", Allele2: "227",
						},
						models.Locus{
							Name:    "Afa15",
							Allele1: "239", Allele2: "243",
						},
					},
				},
				models.Replica{
					Name: "Replica 2",
					LocusArray: []models.Locus{
						models.Locus{
							Name:    "Afa05",
							Allele1: "182", Allele2: "194",
						},
						models.Locus{
							Name:    "Afa13",
							Allele1: "219", Allele2: "219",
						},
						models.Locus{
							Name:    "Afa15",
							Allele1: "239", Allele2: "243",
						},
					},
				},
			},
		},
	}

	result, err := GetResultFromImportData(importData)

	test.AreEqual(t, nil, err, "Error was not null")
	test.AreEqual(t, 5, result.ReplicateAmount, "Replicate amount was incorrect")
	test.AreEqual(t, 2, result.SampleAmount, "Sample amount was incorrect")
	test.AreEqual(t, 0, result.SingleSampleAmount, "Single sample amount was incorrect")
	test.AreEqual(t, 30, result.AmountOfAlleles, "Amount of alleles was incorrect")
	test.AreEqual(t, 2, result.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 15, result.AmountOfLoci, "Amount of loci was incorrect")
	test.AreEqual(t, 30, result.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")

	//Loci orders
	test.AreEqual(t, 3, len(result.LociNamesInOrder), "Amount of loci orders was incorrect")
	test.AreEqual(t, "Afa05", result.LociNamesInOrder[0], "First loci order name was incorrect")
	test.AreEqual(t, "Afa13", result.LociNamesInOrder[1], "Second loci order name was incorrect")
	test.AreEqual(t, "Afa15", result.LociNamesInOrder[2], "Thrid loci order name was incorrect")

	//Loci results
	test.AreEqual(t, 3, len(result.LociResults), "Amount of loci results was incorrect")

	//First loci results
	lociResults := result.LociResults["Afa05"]
	test.AreEqual(t, 2, len(lociResults), "First loci result group amount was incorrect")

	lociResult := lociResults[0]
	test.AreEqual(t, "Afa05", lociResult.Name, "First loci group first result name was incorrect")
	test.AreEqual(t, 6, lociResult.AmountOfAllelesForErrorCalculation, "First loci group first result amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 6, lociResult.TotalAmountOfAlleles, "First loci group first result total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "First loci group first result amount of erroneous alleles was incorrect")
	test.AreEqual(t, "182", lociResult.PrevalentAllele1, "First loci group first result first prevalent allele was incorrect")
	test.AreEqual(t, "194", lociResult.PrevalentAllele2, "First loci group first result second prevalent allele was incorrect")

	lociResult = lociResults[1]
	test.AreEqual(t, "Afa05", lociResult.Name, "First loci group second result name was incorrect")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "First loci group second result amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "First loci group second result total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "First loci group second result amount of erroneous alleles was incorrect")
	test.AreEqual(t, "182", lociResult.PrevalentAllele1, "First loci group second result first prevalent allele was incorrect")
	test.AreEqual(t, "194", lociResult.PrevalentAllele2, "First loci group second result second prevalent allele was incorrect")

	//Second loci results
	lociResults = result.LociResults["Afa13"]
	test.AreEqual(t, 2, len(lociResults), "Second loci result group amount was incorrect")

	lociResult = lociResults[0]
	test.AreEqual(t, "Afa13", lociResult.Name, "Second loci group first result name was incorrect")
	test.AreEqual(t, 6, lociResult.AmountOfAllelesForErrorCalculation, "Second loci group first result amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 6, lociResult.TotalAmountOfAlleles, "Second loci group first result total amount of alleles was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfErroneousAlleles, "Second loci group first result amount of erroneous alleles was incorrect")
	test.AreEqual(t, "219", lociResult.PrevalentAllele1, "Second loci group first result first prevalent allele was incorrect")
	test.AreEqual(t, "227", lociResult.PrevalentAllele2, "Second loci group first result second prevalent allele was incorrect")

	lociResult = lociResults[1]
	test.AreEqual(t, "Afa13", lociResult.Name, "Second loci group second result name was incorrect")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "Second loci group second result amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "Second loci group second result total amount of alleles was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfErroneousAlleles, "Second loci group second result amount of erroneous alleles was incorrect")
	test.AreEqual(t, "219", lociResult.PrevalentAllele1, "Second loci group second result first prevalent allele was incorrect")
	test.AreEqual(t, "227", lociResult.PrevalentAllele2, "Second loci group second result second prevalent allele was incorrect")

	//Third loci results
	lociResults = result.LociResults["Afa15"]
	test.AreEqual(t, 2, len(lociResults), "Third loci result group amount was incorrect")

	lociResult = lociResults[0]
	test.AreEqual(t, "Afa15", lociResult.Name, "Third loci group first result name was incorrect")
	test.AreEqual(t, 6, lociResult.AmountOfAllelesForErrorCalculation, "Third loci group first result amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 6, lociResult.TotalAmountOfAlleles, "Third loci group first result total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Third loci group first result amount of erroneous alleles was incorrect")
	test.AreEqual(t, "239", lociResult.PrevalentAllele1, "Third loci group first result first prevalent allele was incorrect")
	test.AreEqual(t, "243", lociResult.PrevalentAllele2, "Third loci group first result second prevalent allele was incorrect")

	lociResult = lociResults[1]
	test.AreEqual(t, "Afa15", lociResult.Name, "Third loci group second result name was incorrect")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "Third loci group second result amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "Third loci group second result total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Third loci group second result amount of erroneous alleles was incorrect")
	test.AreEqual(t, "239", lociResult.PrevalentAllele1, "Third loci group second result first prevalent allele was incorrect")
	test.AreEqual(t, "243", lociResult.PrevalentAllele2, "Third loci group second result second prevalent allele was incorrect")

	//Sample results
	test.AreEqual(t, 2, len(result.SampleResults), "Sample results amount was incorrect")

	//First sample result
	sampleResult := result.SampleResults[0]

	test.AreEqual(t, "(Replica 1, Replica 2, Replica 3)", sampleResult.Name, "First sample result name was incorrect")
	test.AreEqual(t, false, sampleResult.Single, "First sample single state was incorrect")
	test.AreEqual(t, 3, len(sampleResult.LociResults), "First sample result loci result group amount was incorrect")

	//First sample result loci result
	lociResult = sampleResult.LociResults[0]
	test.AreEqual(t, "Afa05", lociResult.Name, "First loci result name was incorrect in first sample result")
	test.AreEqual(t, 6, lociResult.AmountOfAllelesForErrorCalculation, "First loci result amount of alleles for error calculation was incorrect in first sample result")
	test.AreEqual(t, 6, lociResult.TotalAmountOfAlleles, "First loci result total amount of alleles was incorrect in first sample result")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "First loci result amount of erroneous alleles was incorrect in first sample result")
	test.AreEqual(t, "182", lociResult.PrevalentAllele1, "First loci result first prevalent allele was incorrect in first sample result")
	test.AreEqual(t, "194", lociResult.PrevalentAllele2, "First loci result second prevalent allele was incorrect in first sample result")

	//Second sample result loci result
	lociResult = sampleResult.LociResults[1]
	test.AreEqual(t, "Afa13", lociResult.Name, "Second loci result name was incorrect in first sample result")
	test.AreEqual(t, 6, lociResult.AmountOfAllelesForErrorCalculation, "Second loci result amount of alleles for error calculation was incorrect in first sample result")
	test.AreEqual(t, 6, lociResult.TotalAmountOfAlleles, "Second loci result total amount of alleles was incorrect in first sample result")
	test.AreEqual(t, 1, lociResult.AmountOfErroneousAlleles, "Second loci result amount of erroneous alleles was incorrect in first sample result")
	test.AreEqual(t, "219", lociResult.PrevalentAllele1, "Second loci result first prevalent allele was incorrect in first sample result")
	test.AreEqual(t, "227", lociResult.PrevalentAllele2, "Second loci result second prevalent allele was incorrect in first sample result")

	//Third sample result loci result
	lociResult = sampleResult.LociResults[2]
	test.AreEqual(t, "Afa15", lociResult.Name, "Third loci result name was incorrect in first sample result")
	test.AreEqual(t, 6, lociResult.AmountOfAllelesForErrorCalculation, "Third loci result amount of alleles for error calculation was incorrect in first sample result")
	test.AreEqual(t, 6, lociResult.TotalAmountOfAlleles, "Third loci result total amount of alleles was incorrect in first sample result")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Third loci result amount of erroneous alleles was incorrect in first sample result")
	test.AreEqual(t, "239", lociResult.PrevalentAllele1, "Third loci result first prevalent allele was incorrect in first sample result")
	test.AreEqual(t, "243", lociResult.PrevalentAllele2, "Third loci result second prevalent allele was incorrect in first sample result")

	//Second sample result
	sampleResult = result.SampleResults[1]

	test.AreEqual(t, "(Replica 1, Replica 2)", sampleResult.Name, "Second sample result name was incorrect")
	test.AreEqual(t, false, sampleResult.Single, "Second sample single state was incorrect")
	test.AreEqual(t, 3, len(sampleResult.LociResults), "Second sample result loci result group amount was incorrect")

	//First sample result loci result
	lociResult = sampleResult.LociResults[0]
	test.AreEqual(t, "Afa05", lociResult.Name, "First loci result name was incorrect in second sample result")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "First loci result amount of alleles for error calculation was incorrect in second sample result")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "First loci result total amount of alleles was incorrect in second sample result")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "First loci result amount of erroneous alleles was incorrect in second sample result")
	test.AreEqual(t, "182", lociResult.PrevalentAllele1, "First loci result first prevalent allele was incorrect in second sample result")
	test.AreEqual(t, "194", lociResult.PrevalentAllele2, "First loci result second prevalent allele was incorrect in second sample result")

	//Second sample result loci result
	lociResult = sampleResult.LociResults[1]
	test.AreEqual(t, "Afa13", lociResult.Name, "Second loci result name was incorrect in second sample result")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "Second loci result amount of alleles for error calculation was incorrect in second sample result")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "Second loci result total amount of alleles was incorrect in second sample result")
	test.AreEqual(t, 1, lociResult.AmountOfErroneousAlleles, "Second loci result amount of erroneous alleles was incorrect in second sample result")
	test.AreEqual(t, "219", lociResult.PrevalentAllele1, "Second loci result first prevalent allele was incorrect in second sample result")
	test.AreEqual(t, "227", lociResult.PrevalentAllele2, "Second loci result second prevalent allele was incorrect in second sample result")

	//Third sample result loci result
	lociResult = sampleResult.LociResults[2]
	test.AreEqual(t, "Afa15", lociResult.Name, "Third loci result name was incorrect in second sample result")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "Third loci result amount of alleles for error calculation was incorrect in second sample result")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "Third loci result total amount of alleles was incorrect in second sample result")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Third loci result amount of erroneous alleles was incorrect in second sample result")
	test.AreEqual(t, "239", lociResult.PrevalentAllele1, "Third loci result first prevalent allele was incorrect in second sample result")
	test.AreEqual(t, "243", lociResult.PrevalentAllele2, "Third loci result second prevalent allele was incorrect in second sample result")
}

func TestGetResultFromImportData_ToReturnValidData_WhenImportDataContainsOneSingleSample(t *testing.T) {
	importData := models.NewImportData()

	importData.Samples = []models.Sample{
		models.Sample{
			ReplicaArray: []models.Replica{
				models.Replica{
					Name: "Replica 1",
					LocusArray: []models.Locus{
						models.Locus{
							Name:    "Afa05",
							Allele1: "182", Allele2: "194",
						},
						models.Locus{
							Name:    "Afa13",
							Allele1: "219", Allele2: "227",
						},
						models.Locus{
							Name:    "Afa15",
							Allele1: "239", Allele2: "243",
						},
					},
				},
			},
		},
	}

	result, err := GetResultFromImportData(importData)

	test.AreEqual(t, nil, err, "Error was not null")
	test.AreEqual(t, 1, result.ReplicateAmount, "Replicate amount was incorrect")
	test.AreEqual(t, 1, result.SampleAmount, "Sample amount was incorrect")
	test.AreEqual(t, 1, result.SingleSampleAmount, "Single sample amount was incorrect")
	test.AreEqual(t, 6, result.AmountOfAlleles, "Amount of alleles was incorrect")
	test.AreEqual(t, 0, result.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 3, result.AmountOfLoci, "Amount of loci was incorrect")
	test.AreEqual(t, 0, result.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")

	//Loci orders
	test.AreEqual(t, 3, len(result.LociNamesInOrder), "Amount of loci orders was incorrect")
	test.AreEqual(t, "Afa05", result.LociNamesInOrder[0], "First loci order name was incorrect")
	test.AreEqual(t, "Afa13", result.LociNamesInOrder[1], "Second loci order name was incorrect")
	test.AreEqual(t, "Afa15", result.LociNamesInOrder[2], "Thrid loci order name was incorrect")
}

func TestGetResultFromImportData_ToReturnError_WhenImportDataContainsSameLocusNameTwice(t *testing.T) {
	importData := models.NewImportData()

	importData.Samples = []models.Sample{
		models.Sample{
			ReplicaArray: []models.Replica{
				models.Replica{
					Name: "Replica 1",
					LocusArray: []models.Locus{
						models.Locus{
							Name:    "Afa05",
							Allele1: "182", Allele2: "194",
						},
						models.Locus{
							Name:    "Afa05",
							Allele1: "219", Allele2: "227",
						},
						models.Locus{
							Name:    "Afa15",
							Allele1: "239", Allele2: "243",
						},
					},
				},
			},
		},
	}

	result, err := GetResultFromImportData(importData)

	test.AreEqual(t, "Multiple same loci names found. Loci names must be unique", err.Error(), "Error message was incorrect")
	test.AreEqual(t, 0, result.ReplicateAmount, "Replicate amount was incorrect")
	test.AreEqual(t, 0, result.SampleAmount, "Sample amount was incorrect")
	test.AreEqual(t, 0, result.SingleSampleAmount, "Single sample amount was incorrect")
	test.AreEqual(t, 0, result.AmountOfAlleles, "Amount of alleles was incorrect")
	test.AreEqual(t, 0, result.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 0, result.AmountOfLoci, "Amount of loci was incorrect")
	test.AreEqual(t, 0, result.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
}
