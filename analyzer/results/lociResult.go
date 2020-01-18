package results

import (
	"karim/micsat_errcalc/analyzer/models"
)

// LociResult contains results from given loci
type LociResult struct {
	Name             string
	PrevalentAllele1 string
	PrevalentAllele2 string

	AmountOfAllelesForErrorCalculation int
	TotalAmountOfAlleles               int
	AmountOfErroneousAlleles           int
	AmountOfAlleleDropOuts             int
	AmountOfErroneousLoci              int
	AmountOfHomozygotes                int
	AmountOfHeterozygotes              int
	AmountOfLoci                       int
	AmountOfLociForErrorCalculation    int
	Ambiguous                          bool
}

// CreateLociResult creates LociResult from given loci
func CreateLociResult(name string, loci []models.Locus) LociResult {
	var lociResult = LociResult{}

	lociResult.Name = name
	lociResult.AmountOfAllelesForErrorCalculation = 0
	lociResult.AmountOfLoci = 0
	lociResult.TotalAmountOfAlleles = 0
	lociResult.AmountOfAlleleDropOuts = 0
	lociResult.AmountOfErroneousAlleles = 0
	lociResult.AmountOfErroneousLoci = 0
	lociResult.AmountOfHomozygotes = 0
	lociResult.AmountOfHeterozygotes = 0
	lociResult.AmountOfLociForErrorCalculation = 0
	lociResult.Ambiguous = false

	filledAlleles := 0

	for _, locus := range loci {
		if !locus.IsEmpty() {
			filledAlleles++
			lociResult.TotalAmountOfAlleles += 2
		}

		if locus.IsHomozygot() {
			lociResult.AmountOfHomozygotes++
		} else if locus.IsHeterozygot() {
			lociResult.AmountOfHeterozygotes++
		}
	}

	if filledAlleles < 2 {
		return lociResult
	}

	lociResult.AmountOfAlleleDropOuts = getAlleleDropOutAmount(loci)
	allele1PrevalentCandidates, allele2PrevalentCandidates := getPrevalentAlleleCandidates(loci)
	lociResult.AmountOfAllelesForErrorCalculation = getAmountOfAllelesForErrorCalculation(loci)
	lociResult.AmountOfLoci = len(loci)
	lociResult.AmountOfLociForErrorCalculation = getAmountOfLociForErrorCalculation(loci)

	if len(loci) < 2 {
		return lociResult
	}

	prevalentAllele1, prevalentAllele1Ambiguous, prevalentAllele1MaxCount := getPrevalentAlleleResults(allele1PrevalentCandidates)
	prevalentAllele2, prevalentAllele2Ambiguous, prevalentAllele2MaxCount := getPrevalentAlleleResults(allele2PrevalentCandidates)

	//Solve ambiguous prevalent alleles using other prevalent allele if found
	if prevalentAllele1Ambiguous && !prevalentAllele2Ambiguous {
		prevalentAllele1 = solvePrevalentAlleleUsingOtherPrevalentAllele(
			allele1PrevalentCandidates,
			prevalentAllele1MaxCount,
			prevalentAllele2)

		lociResult.Ambiguous = true
	} else if !prevalentAllele1Ambiguous && prevalentAllele2Ambiguous {
		prevalentAllele2 = solvePrevalentAlleleUsingOtherPrevalentAllele(
			allele2PrevalentCandidates,
			prevalentAllele2MaxCount,
			prevalentAllele1)

		lociResult.Ambiguous = true
	} else if prevalentAllele1Ambiguous && prevalentAllele2Ambiguous {
		lociResult.Ambiguous = true
	}

	//Count errors
	amountOfErroneousAlleles, amountOfErroneousLoci := calculateAlleleAndLocusErrorAmountFromLoci(loci, prevalentAllele1, prevalentAllele2)

	lociResult.AmountOfErroneousAlleles = amountOfErroneousAlleles
	lociResult.AmountOfErroneousLoci = amountOfErroneousLoci

	lociResult.PrevalentAllele1 = prevalentAllele1
	lociResult.PrevalentAllele2 = prevalentAllele2

	return lociResult
}

func getAmountOfAllelesForErrorCalculation(loci []models.Locus) int {
	count := 0

	for _, locus := range loci {
		if len(locus.Allele1) > 0 {
			count++
		}

		if len(locus.Allele2) > 0 {
			count++
		}
	}

	return count
}

func getAmountOfLociForErrorCalculation(loci []models.Locus) int {
	count := 0

	for _, locus := range loci {
		if len(locus.Allele1) > 0 &&
			len(locus.Allele2) > 0 &&
			locus.Allele1 != "?" &&
			locus.Allele2 != "?" {
			count++
		}
	}

	return count
}

func getPrevalentAlleleCandidates(loci []models.Locus) (map[string]int, map[string]int) {
	allele1Candidates := make(map[string]int)
	allele2Candidates := make(map[string]int)

	for _, locus := range loci {
		if len(locus.Allele1) > 0 {
			//lociResult.AmountOfAllelesForErrorCalculation++

			if locus.Allele1 != "?" {
				if _, ok := allele1Candidates[locus.Allele1]; ok {
					allele1Candidates[locus.Allele1]++
				} else {
					allele1Candidates[locus.Allele1] = 1
				}
			}
		}

		if len(locus.Allele2) > 0 {
			//lociResult.AmountOfAllelesForErrorCalculation++

			if locus.Allele2 != "?" {
				if _, ok := allele2Candidates[locus.Allele2]; ok {
					allele2Candidates[locus.Allele2]++
				} else {
					allele2Candidates[locus.Allele2] = 1
				}
			}
		}
	}

	return allele1Candidates, allele2Candidates
}

func solvePrevalentAlleleUsingOtherPrevalentAllele(
	alleleCandidates map[string]int,
	prevalentAlleleMaxCount int,
	otherSolvedPrevalentAllele string) string {
	prevalentAllele := ""

	for allelePrevalentCandidate, allelePrevalentCandidateCount := range alleleCandidates {
		if allelePrevalentCandidateCount == prevalentAlleleMaxCount &&
			allelePrevalentCandidate != otherSolvedPrevalentAllele {
			prevalentAllele = allelePrevalentCandidate
		}
	}

	return prevalentAllele
}

func calculateAlleleAndLocusErrorAmountFromLoci(loci []models.Locus, prevalentAllele1 string, prevalentAllele2 string) (int, int) {
	alleleErrorCount := 0
	locusErrorCount := 0

	for _, locus := range loci {
		locusErrorFound := false

		if len(locus.Allele1) > 0 &&
			locus.Allele1 != "?" &&
			locus.Allele1 != prevalentAllele1 {
			alleleErrorCount++

			locusErrorCount++
			locusErrorFound = true
		}

		if len(locus.Allele2) > 0 &&
			locus.Allele2 != "?" &&
			locus.Allele2 != prevalentAllele2 {
			alleleErrorCount++

			if !locusErrorFound {
				locusErrorCount++
			}
		}

		// Loci error is not calculated from single samples
		if len(loci) <= 1 {
			locusErrorCount = 0
		}
	}

	return alleleErrorCount, locusErrorCount
}

func getPrevalentAlleleResults(allelePrevalentCandidates map[string]int) (string, bool, int) {
	prevalentAllele := ""
	prevalentAlleleAmbiguous := false
	prevalentAlleleMaxCount := 0

	for allelePrevalentCandidate, allelePrevalentCandidateCount := range allelePrevalentCandidates {
		if allelePrevalentCandidateCount > prevalentAlleleMaxCount {
			prevalentAllele = allelePrevalentCandidate
			prevalentAlleleMaxCount = allelePrevalentCandidateCount
			prevalentAlleleAmbiguous = false
		} else if allelePrevalentCandidateCount == prevalentAlleleMaxCount {
			prevalentAlleleAmbiguous = true
		}
	}

	return prevalentAllele, prevalentAlleleAmbiguous, prevalentAlleleMaxCount
}

func getAlleleDropOutAmount(loci []models.Locus) int {
	amountOfHeterozygotes := 0
	amountOfHomozygotes := 0

	for _, locus := range loci {
		if locus.IsHomozygot() {
			amountOfHomozygotes++
		} else if locus.IsHeterozygot() {
			amountOfHeterozygotes++
		}
	}

	if amountOfHomozygotes > 0 &&
		amountOfHeterozygotes > 0 &&
		amountOfHomozygotes <= amountOfHeterozygotes {
		return amountOfHomozygotes
	}

	return 0
}
