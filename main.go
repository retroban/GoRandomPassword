package main

// Программа для генерации паролей с использованием криптографически безопасного генератора случайных чисел
// и алгоритма Фишера-Йетса для перемешивания символов. Пароли содержат строчные и прописные буквы, цифры и специальные символы.
// Программа также удаляет неоднозначные символы, такие как "l", "1", "O", "0", чтобы избежать путаницы при вводе пароля.
// Генерируем 12 паролей длиной 12 символов, которые содержат строчные и прописные буквы, цифры и специальные символы.
// Программа использует криптографически безопасный генератор случайных чисел для генерации паролей и алгоритм Фишера-Йетса для перемешивания символов.
// Программа написана на языке Go и использует стандартные библиотеки для генерации случайных чисел и работы со строками.

import (
    "crypto/rand"
    "fmt"
    "math/big"
    "strings"
)

const passwordLength = 12 // Длина пароля
const numberOfPasswords = 12 // Количество паролей для генерации
const ambiguousChars = "B8G6I1l|0OQDS5Z2" // Символы, которые могут быть затруднительны для чтения

func generatePassword(length int) string {
	// Генерируем пароль заданной длины
    if length < 4 {
        panic("Password length must be at least 4 to include all character types")
    }

    // Наборы символов для генерации пароля
    lowercase := removeAmbiguousChars("abcdefghijklmnopqrstuvwxyz")
    uppercase := removeAmbiguousChars("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    digits := removeAmbiguousChars("0123456789")
    special := removeAmbiguousChars("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
    allCharset := lowercase + uppercase + digits + special

    // Гарантируем, что в пароле будет хотя бы один символ из каждого набора
    password := []byte{
        lowercase[randomIndex(len(lowercase))],
        uppercase[randomIndex(len(uppercase))],
        digits[randomIndex(len(digits))],
        special[randomIndex(len(special))],
    }

    // Заполняем оставшуюся часть пароля случайными символами из полного набора
	// Используем криптографически безопасный генератор случайных чисел
    for i := 4; i < length; i++ {
        password = append(password, allCharset[randomIndex(len(allCharset))])
    }

    // Перемешиваем пароль, чтобы случайно распределить символы
	// Используем алгоритм Фишера-Йетса для перемешивания
    shuffle(password)

    return string(password)
}

func removeAmbiguousChars(charset string) string {
	// Удаляем неоднозначные символы из набора
	    for _, char := range ambiguousChars {
        charset = strings.ReplaceAll(charset, string(char), "")
    }
    return charset
}

func randomIndex(max int) int {
	// Генерируем случайный индекс
	// Используем криптографически безопасный генератор случайных чисел
    index, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
    return int(index.Int64())
}

func shuffle(data []byte) {
	// Перемешиваем массив символов
	// Используем алгоритм Фишера-Йетса для перемешивания
    for i := range data {
        j := randomIndex(len(data))
        data[i], data[j] = data[j], data[i]
    }
}

func main() {
	// Генерируем и выводим пароли
    for i := 0; i < numberOfPasswords; i++ {
        password := generatePassword(passwordLength)
        fmt.Println(password)
    }
    // Ожидание ввода от пользователя, чтобы окно консоли не закрылось
    fmt.Println("Нажмите Enter, чтобы выйти...")
    fmt.Scanln()
}
