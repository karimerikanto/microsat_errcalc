package results

import (
	"karim/microsatellite_analyzer/analyzer/models"
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
	Ambiguous                          bool
}

// CreateLociResult creates LociResult from given loci
func CreateLociResult(name string, loci []models.Locus) LociResult {
	var lociResult = LociResult{}

	lociResult.Name = name
	lociResult.AmountOfAllelesForErrorCalculation = 0
	lociResult.TotalAmountOfAlleles = 0
	lociResult.AmountOfAlleleDropOuts = 0
	lociResult.AmountOfErroneousAlleles = 0
	lociResult.Ambiguous = false

	filledAlleles := 0

	for _, locus := range loci {
		if !locus.IsEmpty() {
			filledAlleles++
			lociResult.TotalAmountOfAlleles += 2
		}
	}

	if filledAlleles < 2 {
		return lociResult
	}

	lociResult.AmountOfAlleleDropOuts = getAlleleDropOutAmount(loci)

	//Find the most prevalent alleles
	allele1PrevalentCandidates := make(map[string]int)
	allele2PrevalentCandidates := make(map[string]int)

	for _, locus := range loci {
		if len(locus.Allele1) > 0 {
			lociResult.AmountOfAllelesForErrorCalculation++

			if locus.Allele1 != "?" {
				if _, ok := allele1PrevalentCandidates[locus.Allele1]; ok {
					allele1PrevalentCandidates[locus.Allele1]++
				} else {
					allele1PrevalentCandidates[locus.Allele1] = 1
				}
			}
		}

		if len(locus.Allele2) > 0 {
			lociResult.AmountOfAllelesForErrorCalculation++

			if locus.Allele2 != "?" {
				if _, ok := allele2PrevalentCandidates[locus.Allele2]; ok {
					allele2PrevalentCandidates[locus.Allele2]++
				} else {
					allele2PrevalentCandidates[locus.Allele2] = 1
				}
			}
		}
	}

	if len(loci) < 2 {
		return lociResult
	}

	prevalentAllele1, prevalentAllele1Ambiguous, prevalentAllele1MaxCount := getPrevalentAlleleResults(allele1PrevalentCandidates)
	prevalentAllele2, prevalentAllele2Ambiguous, prevalentAllele2MaxCount := getPrevalentAlleleResults(allele2PrevalentCandidates)

	//Solve ambiguous prevalent alleles using other prevalent allele if found
	if prevalentAllele1Ambiguous && !prevalentAllele2Ambiguous {
		for allele1PrevalentCandidate, allele1PrevalentCandidateCount := range allele1PrevalentCandidates {
			if allele1PrevalentCandidateCount == prevalentAllele1MaxCount &&
				allele1PrevalentCandidate != prevalentAllele2 {
				prevalentAllele1 = allele1PrevalentCandidate
			}
		}

		lociResult.Ambiguous = true
	} else if !prevalentAllele1Ambiguous && prevalentAllele2Ambiguous {
		for allele2PrevalentCandidate, allele2PrevalentCandidateCount := range allele2PrevalentCandidates {
			if allele2PrevalentCandidateCount == prevalentAllele2MaxCount &&
				allele2PrevalentCandidate != prevalentAllele1 {
				prevalentAllele2 = allele2PrevalentCandidate
			}
		}

		lociResult.Ambiguous = true
	} else if prevalentAllele1Ambiguous && prevalentAllele2Ambiguous {
		lociResult.Ambiguous = true
	}

	//Count errors
	for _, locus := range loci {
		if len(locus.Allele1) > 0 &&
			locus.Allele1 != "?" &&
			locus.Allele1 != prevalentAllele1 {
			lociResult.AmountOfErroneousAlleles++
		}

		if len(locus.Allele2) > 0 &&
			locus.Allele2 != "?" &&
			locus.Allele2 != prevalentAllele2 {
			lociResult.AmountOfErroneousAlleles++
		}

	}

	lociResult.PrevalentAllele1 = prevalentAllele1
	lociResult.PrevalentAllele2 = prevalentAllele2

	return lociResult
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
	amountOfHeterozygots := 0
	amountOfHomozygots := 0

	for _, locus := range loci {
		if locus.IsHomozygot() {
			amountOfHomozygots++
		} else if locus.IsHeterozygot() {
			amountOfHeterozygots++
		}
	}

	if amountOfHomozygots > 0 &&
		amountOfHeterozygots > 0 &&
		amountOfHomozygots <= amountOfHeterozygots {
		return amountOfHomozygots
	}

	return 0
}
