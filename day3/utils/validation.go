package utils

import "errors"

func ValidasiNilai(score []int) error {
	for _, value := range score {
		if value < 0 || value > 100 {
			return errors.New("nilai ga valid")
		}
	}
	return nil
}

func CekNoHP(num *string) error {

	if num == nil{
		return errors.New("no hp gaboleh kosong")
	}

	length := len(*num)

	if length < 10 || length > 13 {
		return errors.New("nomor hp ga valid")
	}
	return nil
}