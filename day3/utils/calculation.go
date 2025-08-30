package utils

func RataRataNilai(score ...int) float64{
	if len(score) == 0 {
		return 0
	}

	total := 0
	for _, value := range(score){
		total += value
	}

	average := float64(total) / float64(len(score))
	return average
	
}

func CekStatusKelulusan(num float64) (string, bool) {
	switch {
	case num > 60 && num < 70:
		return "Tidak lulus", false
	case num > 70 && num < 80:
		return "Remedial", false
	case num > 80 && num < 90:
		return "Lulus", true
	case num > 90 && num < 100:
		return "Lulus dengan pujian", true
	default:
		return ">//<", false
	}
}