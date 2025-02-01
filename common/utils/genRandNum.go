package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenRandNums(length int) string {
	if length <= 0 {
		return ""
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Первую цифру делаем ненулевой, чтобы число не начиналось с 0
	number := fmt.Sprintf("%d", rand.Intn(9)+1)

	// Генерируем оставшиеся цифры
	for i := 1; i < length; i++ {
		number += fmt.Sprintf("%d", rng.Intn(10))
	}

	return number
}
