package results

import (
	"karim/micsat_errcalc/analyzer/models"
	"karim/micsat_errcalc/utils"
	"testing"
)

func TestCreateLociResult_ToReturnEmptyLociResult_WhenOnlyOneLocusIsGiven(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 2, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnEmptyLociResult_WhenTwoLociAreGivenButOtherIsEmpty(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "",
			Allele2: "",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 2, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenThreeLociAreGivenWithSameAlleleValues(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 6, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 6, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "200", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 3, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenThreeLociAreGivenAndOneAlleleIsNotSame(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "101",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 6, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 6, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "200", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 3, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenThreeLociAreGivenAndLastLocusIsEmpty(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "",
			Allele2: "",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "200", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenFourLociAreGivenAndAllele1IsAmbiguous(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "120",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "120",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "120",
			Allele2: "120",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "120",
			Allele2: "120",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, true, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 8, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 8, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "120", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenTwoLociAreGivenAndFirstIncludesQuestionMarks(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "?",
			Allele2: "?",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "200", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenTwoLociAreGivenAndFirstIncludesQuestionMarkAndValue(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "?",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 4, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 4, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "200", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenFourLociAreGivenAndBothAllelesAreAmbiguousButFoundFromOtherPrevalentAllele(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "200",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 8, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 8, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "200", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 4, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenFiveLociAreGivenAndThreeOfThemAreHeterotsygotAndTwoOfThemAreHomotsygot(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "100",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 10, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 10, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfAlleleDropOuts, "Amount of allele drop outs was incorrect")
	test.AreEqual(t, "150", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 3, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenFiveLociAreGivenAndTwoOfThemAreHeterotsygotAndThreeOfThemAreHomotsygot(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "100",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 10, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 10, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfAlleleDropOuts, "Amount of allele drop outs was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 3, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenFourLociAreGivenAndTwoOfThemAreHeterotsygotAndTwoOfThemAreHomotsygot(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "100",
			Allele2: "100",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, true, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 8, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 8, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfAlleleDropOuts, "Amount of allele drop outs was incorrect")
	test.AreEqual(t, "150", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}

func TestCreateLociResult_ToReturnValidLociResult_WhenFourLociAreGivenAndOneOfThemHasTwoAllelleErrors(t *testing.T) {
	loci := []models.Locus{
		models.Locus{
			Name:    "Locus1",
			Allele1: "50",
			Allele2: "60",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
		models.Locus{
			Name:    "Locus1",
			Allele1: "150",
			Allele2: "100",
		},
	}

	lociResult := CreateLociResult("Locus1", loci)

	test.AreEqual(t, "Locus1", lociResult.Name, "Loci result name was incorrect")
	test.AreEqual(t, false, lociResult.Ambiguous, "Loci result ambiguous state was incorrect")
	test.AreEqual(t, 8, lociResult.AmountOfAllelesForErrorCalculation, "Amount of alleles for error calculation was incorrect")
	test.AreEqual(t, 8, lociResult.TotalAmountOfAlleles, "Total amount of alleles was incorrect")
	test.AreEqual(t, 2, lociResult.AmountOfErroneousAlleles, "Amount of erroneous alleles was incorrect")
	test.AreEqual(t, 1, lociResult.AmountOfErroneousLoci, "Amount of erroneous loci was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfAlleleDropOuts, "Amount of allele drop outs was incorrect")
	test.AreEqual(t, "150", lociResult.PrevalentAllele1, "Prevalent allele 1 was incorrect")
	test.AreEqual(t, "100", lociResult.PrevalentAllele2, "Prevalent allele 2 was incorrect")
	test.AreEqual(t, 4, lociResult.AmountOfHeterozygotes, "Amount of heterozygotes was incorrect")
	test.AreEqual(t, 0, lociResult.AmountOfHomozygotes, "Amount of homozygotes was incorrect")
}
