package cipher

type caesar struct { 
	encode func(string) string
	decode func(string) string
}

func (t caesar) Encode(source string) string { return t.encode(source) }
func (t caesar) Decode(source string) string { return t.decode(source) }

func NewCaesar() Cipher {
	t := struct { caesar }{}

	encrypt := func(b byte) byte {
		r := b + 3
		if r > 122 { 
			return r - 26 
		} else {
			return r
		}
	}

	decrypt := func(b byte) byte {
		r := b - 3
		if r < 97 { 
			return r + 26 
		} else {
			return r
		}
	}

	t.encode = func(source string) string { 
		var resultByte []byte
		plainText := convertToPlainText(source)
		for i:=0; i < len(plainText); i++ {
			resultByte = append(resultByte, encrypt(plainText[i]))
		}
		return string(resultByte)
	}
	t.decode = func(source string) string { 
		var resultByte []byte
		plainText := convertToPlainText(source)
		for i:=0; i < len(plainText); i++ {
			resultByte = append(resultByte, decrypt(plainText[i]))
		}
		return string(resultByte)
	}

	return t
}