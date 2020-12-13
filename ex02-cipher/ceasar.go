package cipher

type caesar struct { 
	encode func(string) string
	decode func(string) string
}

func (t caesar) Encode(source string) string { return t.encode(source) }
func (t caesar) Decode(source string) string { return t.decode(source) }

func NewCaesar() Cipher {
	t := struct { caesar }{}

	t.encode = func (source string) string {
		return caesarCipher(source, 3)
	}
	t.decode = func (source string) string {
		return caesarCipher(source, -3)
	}

	return t
}

func NewShift(shift int) Cipher {
	if shift < -25 || shift > 25 || shift == 0 {
		return nil
	}

	t := struct { caesar }{}

	t.encode = func (source string) string {
		return caesarCipher(source, shift)
	}
	t.decode = func (source string) string {
		return caesarCipher(source, shift * -1)
	}

	return t
}

func NewVigenere(key string) Cipher {
	if !isValidKey(key) {
		return nil
	}

	t := struct { caesar }{}

	t.encode = func (source string) string {
		plainText := convertToPlainText(source)
		keyLen := len(key)

		var resultByte []byte 
	  var	first int = 97
		for i:=0; i < len(plainText); i++ {
			code := (i - ((i / keyLen) * keyLen))
			shift := key[code] - byte(first)
			resultByte = append(resultByte, shiftByte(plainText[i], int(shift)))
		}
		return string(resultByte)
	}
	t.decode = func (source string) string {
		plainText := convertToPlainText(source)
		keyLen := len(key)

		var resultByte []byte 
	  var	first int = 97
		for i:=0; i < len(plainText); i++ {
			code := (i - ((i / keyLen) * keyLen))
			shift := key[code] - byte(first)
			resultByte = append(resultByte, shiftByte(plainText[i], int(shift) * -1))
		}
		return string(resultByte)
	}

	return t
}

func isValidKey(key string) bool {
	isValid := false
	for i:=0; i < len(key); i++ {
		current := key[i]
		if !isValid && current != 97 {
			isValid = true
		} 
		if current < 97 || current > 122 {
			return false
		}
	}
	return isValid
}

func caesarCipher(source string, shift int) string { 
	var resultByte []byte
	plainText := convertToPlainText(source)
	for i:=0; i < len(plainText); i++ {
		resultByte = append(resultByte, shiftByte(plainText[i], shift))
	}
	return string(resultByte)
}

func shiftByte(b byte, shift int) byte {
	r := b + byte(shift)
	if shift >= 0 {
		if r > 122 { 
			return r - 26 
		} else {
			return r
		}
	} else {
		if r < 97 { 
			return r + 26 
		} else {
			return r
		}
	}
}