package downcase

import (
	"sync"
)

func Downcase(arg string) (string, error) {
	length := len(arg)
	result := make([]byte, length)

	var wg sync.WaitGroup
  wg.Add(length)

	for i := 0; i < length; i++ {
		go downcase(arg[i], result, i, &wg)
	}
	wg.Wait()
	return string(result), nil
}

func downcase(arg byte, result []byte, index int, wg *sync.WaitGroup) {
	if arg >= 65 && arg <= 90 {
		result[index] = arg + 32
	} else {
		result[index] = arg
	}
	wg.Done()
}

// func Downcase(arg string) (string, error) {
// 	var result []byte
// 	for i := 0; i < len(arg); i++ {
// 		result = append(result, downcase(arg[i]))
// 	}
// 	return string(result), nil
// }

// func downcase(arg byte) byte {
// 	fmt.Println(string(arg))
// 	if arg >= 65 && arg <= 90 {
// 		return arg + 32
// 	} else {
// 		return arg
// 	}
// }