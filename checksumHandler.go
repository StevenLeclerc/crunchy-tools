package crunchyTools

import "crypto/md5"

func GenMd5FromString(errorMsg string) [16]byte {
	//create an MD5 from errorMsg
	return md5.Sum([]byte(errorMsg))
}

func CheckMd5FromStringInSlice(data [][16]byte, stringToFound string) bool {
	md5ToFound := GenMd5FromString(stringToFound)
	for _, m := range data {
		if m == md5ToFound {
			return true
		}
	}
	return false
}
