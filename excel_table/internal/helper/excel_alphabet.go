package helper

var (
	alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func GenExcelColumnAlphabet(lColumn int) []string {
	if lColumn < 1 {
		return nil
	}
	excelAlph := make([]string, 0, lColumn+len(alphabet))
	n := lColumn / len(alphabet)
	// first just append
	for n >= 0 {
		excelAlph = append(excelAlph, alphabet...)
		n--
	}

	// then cut extra
	excelAlph = excelAlph[:lColumn]

	// generate extra letters
	j := 0
	count := 0
	for i := len(alphabet); i < lColumn; i++ {
		excelAlph[i] = alphabet[j] + excelAlph[i]
		count++
		if count == len(alphabet) {
			count = 0
			j++ // next letter in alphabet
		}
	}

	return excelAlph
}
