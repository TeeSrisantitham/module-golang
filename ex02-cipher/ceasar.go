package cipher

import "sync"

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

	process := func (source string, mulShift int) string {
		plainText := convertToPlainText(source)
		v := len(plainText)
		result := make ([]byte, v)
		var wg sync.WaitGroup
		wg.Add(v)
		for i:=0; i < v; i++ {
			go vigenereCipher(key, plainText, mulShift, i, result, &wg)
		}
		wg.Wait()
		return string(result)
	}

	t := struct { caesar }{}

	t.encode = func (source string) string {
		return process(source, 1)
	}
	t.decode = func (source string) string {
		return process(source, -1)
	}

	return t
}

func vigenereCipher(key string, plainText string, mulShift int, i int, result []byte, wg *sync.WaitGroup) {
	keyLen := len(key)
	code := (i - ((i / keyLen) * keyLen))
	shift := key[code] - byte(97)
	result[i] = shiftByte(plainText[i], int(shift) * mulShift)
	wg.Done()
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
	plainText := convertToPlainText(source)
	v := len(plainText)
	result := make ([]byte, v)
	var wg sync.WaitGroup
	wg.Add(v)
	for i:=0; i < v; i++ {
		go func (plainText string, i int, result []byte, wg *sync.WaitGroup) {
			result[i] = shiftByte(plainText[i], shift)
			wg.Done()
		} (plainText, i, result, &wg)
	}
	wg.Wait()
	return string(result)
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