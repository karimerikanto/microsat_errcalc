package models

// ImportData includes all the data for analysis (headers, loci and samples)
type ImportData struct {
	Headers    [][]string
	Samples    []Sample
	LociGroups map[string][]Locus
}

// NewImportData creates new ImportData object
func NewImportData() *ImportData {
	importData := new(ImportData)
	importData.LociGroups = make(map[string][]Locus)

	return importData
}

// AppendLoci appends new loci to the existing loci group
func (importData ImportData) AppendLoci(loci []Locus) {
	for _, locus := range loci {
		if _, ok := importData.LociGroups[locus.Name]; ok {
			importData.LociGroups[locus.Name] = append(importData.LociGroups[locus.Name], locus)
		} else {
			importData.LociGroups[locus.Name] = []Locus{locus}
		}
	}
}
