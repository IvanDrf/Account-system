package encryption

func Encrypt(str *string) {
	bytes := []byte(*str)

	for i, element := range bytes {
		bytes[i] = element * element
	}

	*str = string(bytes)
}
