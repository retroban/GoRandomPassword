package main

// Program for generating passwords using cryptographically secure random number generator
// and Fisher-Yates algorithm for mixing characters. Passwords contain lowercase and uppercase letters, numbers and special characters.
// The programme also removes ambiguous characters such as "l", "1", "O", "0" to avoid confusion when entering the password.
// Generate 12 passwords that are 12 characters long and contain lowercase and uppercase letters, numbers, and special characters.
// The programme uses a cryptographically secure random number generator to generate the passwords and the Fisher-Yates algorithm to mix the characters.
// The programme is written in Go language and uses standard libraries for generating random numbers and working with strings.

import (
    "crypto/rand"
    "fmt"
    "math/big"
    "strings"
)

const passwordLength = 12 // Password length
const numberOfPasswords = 12 // Number of passwords to generate
const ambiguousChars = "B8G6I1l|0OQDS5Z2" // Symbols that may be difficult to read

func generatePassword(length int) string {
	// Generate a password of the specified length
    if length < 4 {
        panic("Password length must be at least 4 to include all character types")
    }

    // Character sets for password generation
    lowercase := removeAmbiguousChars("abcdefghijklmnopqrstuvwxyz")
    uppercase := removeAmbiguousChars("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    digits := removeAmbiguousChars("0123456789")
    special := removeAmbiguousChars("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
    allCharset := lowercase + uppercase + digits + special

    // Ensure that the password contains at least one character from each set.
    password := []byte{
        lowercase[randomIndex(len(lowercase))],
        uppercase[randomIndex(len(uppercase))],
        digits[randomIndex(len(digits))],
        special[randomIndex(len(special))],
    }

    // Fill the remaining part of the password with random characters from the full set
// Use a cryptographically secure random number generator.
    for i := 4; i < length; i++ {
        password = append(password, allCharset[randomIndex(len(allCharset))])
    }

    // Shuffle the password to randomly distribute the characters
// Use the Fisher-Yates algorithm for shuffling.
    shuffle(password)

    return string(password)
}

func removeAmbiguousChars(charset string) string {
	// Remove ambiguous characters from the set
	    for _, char := range ambiguousChars {
        charset = strings.ReplaceAll(charset, string(char), "")
    }
    return charset
}

func randomIndex(max int) int {
	// Generate a random index
// Use a cryptographically secure random number generator
    index, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
    return int(index.Int64())
}

func shuffle(data []byte) {
	// Shuffle the character array
// Use the Fisher-Yates algorithm for shuffling.
    for i := range data {
        j := randomIndex(len(data))
        data[i], data[j] = data[j], data[i]
    }
}

func main() {
	// Generate and output passwords
    for i := 0; i < numberOfPasswords; i++ {
        password := generatePassword(passwordLength)
        fmt.Println(password)
    }
 // Waiting for user input so that the console window does not close
    fmt.Println("Press Enter to exit...")
    fmt.Scanln()
}
