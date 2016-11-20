package utils

import (

)
import (
	"fmt"
	"log"
	"crypto/aes"
	"encoding/base64"
	//"io"
	//"crypto/rand"
	"crypto/cipher"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Hash a text using bcrypt.
func HashString(text []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(text, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

// Compare a string with a hashed text.
func HashMatch(hashedText []byte, text []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedText, text)
	if err != nil {
		return false
	}
	return true
}

func Cypter() {
	key := []byte("key3456789012345") // Must be 16, 24 or 32 bytes

	sourceText := [] byte("Ok encript this.")
	fmt.Printf("%s\n", sourceText)

	// Encrypt the message
	cyphertext, err := encrypt(key, sourceText)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(cyphertext))

	// Now decrypt it again
	uncrypted, err := decrypt(key, cyphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", uncrypted)

}

// See alternate IV creation from ciphertext below
//var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// Perform the encryption of a message.
func encrypt(key, text []byte) ([]byte, error){
	// First hash the

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("New cipher erroe.")
		return  nil, err
	}
	log.Printf("Blocksize %s", block.BlockSize())

	b := base64.StdEncoding.EncodeToString(text)

	var dst  = make([]byte, 16)
	var src = []byte(b)

	block.Encrypt(dst, src)
	return dst, nil
	//ciphertext := make([]byte, aes.BlockSize+len(b))
	//
	//iv := ciphertext[:aes.BlockSize] // Should be generated separately in production
	//if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	//	return nil, err
	//}
	//
	//cfb := cipher.NewCBCEncrypter(block, iv)
	//cfb.CryptBlocks(ciphertext[aes.BlockSize:], []byte(b)) //XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	//return ciphertext, nil
}

// Perform the decryption of a message.
func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext is too short.")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCBCDecrypter(block, iv)
	cfb.CryptBlocks(text, text)//.XORKeyStream(test, test)

	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
