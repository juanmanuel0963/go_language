package main

import (
	"errors"
	"fmt"
)

/*
/// DESCRIPTION

Julius Caesar protected his confidential information by encrypting it using a cipher.
Caesar's cipher shifts each letter by a number of letters. For example, in the case of
a rotation by 3, w, x, y and z would map to z, a, b and c.

Original alphabet:      abcdefghijklmnopqrstuvwxyz
Alphabet rotated +3:    defghijklmnopqrstuvwxyzabc

/// REQUIREMENTS.

1. If the shift takes you past the end of the alphabet, just rotate back to the front of the alphabet.
2. If the shift number is a negative number, log an error.
3. The cipher only encrypts letters; all other characters remain unencrypted.
4. Build a function CaesarCipher(s string, n int) (string, error) to handle your implementation.
5. Create unit tests to validate previous requirements.

/// EXAMPLE.

s := "uno-square *es* %una gran empresa."
n := 3

The alphabet is rotated by 3, matching the mapping above.

The encrypted string is: "xqr-vtxduh *hv* %xqd judq hpsuhvd."
						  xqr-vtxduh *hv* %xqd judq hpsuhvd.

/// BONUS

Implement a concurrent mechanism to handle the execution of Caesar's Cipher over multiple strings
*/

func GetCipheredChar(s string, rotationMap map[string]string) string {

	ciphered := rotationMap[s]

	if len(ciphered) == 0 {
		ciphered = s
	}

	return ciphered
}

// func CaesarCipher(s string, n int) (string, error) {
func CaesarCipher(s string, n int, responseChannel chan<- string, errorChannel chan<- error) {

	if n < 0 {
		errorChannel <- errors.New("shift can't be negative")
		//return "", errors.New("shift can't be negative")
	} else {

		alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		rotation := make([]string, 0)

		//Finding new corresponding chars by rotating n positions the alphabet
		for i := 0; i < len(alphabet); i++ {
			if i >= n && i < len(alphabet) {
				rotation = append(rotation, alphabet[i])
			}
		}
		//Adding missing chars at end of the rotation array
		if len(rotation) < len(alphabet) {
			for i := 0; i < len(alphabet); i++ {
				if len(rotation) < len(alphabet) {
					rotation = append(rotation, alphabet[i])
				}
			}
		}

		var rotationMap = make(map[string]string, 0)

		//Creating the rotated alphabet into a Map for easy to use
		for i := 0; i < len(rotation); i++ {
			rotationMap[alphabet[i]] = rotation[i]
		}

		response := ""

		for i := 0; i < len(s); i++ {
			response += GetCipheredChar(string(s[i]), rotationMap)
		}

		responseChannel <- response
	}
	//return response, nil
}

func main() {

	fmt.Println("Welcome to the Code Challenge")

	tests := []string{"uno-square *es* %una gran empresa.", "*** welcome to the code challenge ***"}
	shift := []int{3, -1}

	for i := 0; i < len(tests); i++ {

		responseChannel := make(chan string)
		errorChannel := make(chan error)

		go CaesarCipher(tests[i], shift[i], responseChannel, errorChannel)

		defer close(errorChannel)
		defer close(responseChannel)

		// Wait for the cipher to be created and sent through the channel
		select {
		case response := <-responseChannel:
			fmt.Println(response)
		case err := <-errorChannel:
			fmt.Println(err.Error())
		}

	}
}
