package hash

import "strconv"

func HashWithByteShift(input string) string {
	bytes := []byte(input)
	var hash string

	for i, b := range bytes {
		// Сдвигаем каждый байт влево на i позиций. Используем операцию & 0xFF для обеспечения того,
		// чтобы результат оставался в пределах байта после сдвига.
		shifted := (b << i) & 0xFF
		// Конвертируем результат сдвига в строку в шестнадцатеричном формате и добавляем к итоговому хешу
		hash += strconv.FormatInt(int64(shifted), 16)
	}

	return hash
}
