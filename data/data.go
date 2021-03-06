package data

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// creating a random UUID from RFC 4122.
func CreateUUID() (uuid string){
	u := new([16]byte)
	_,err := rand.Read(u[:])

	if err != nil {
		log.Fatalln("can't generate uuid",err)
		return
	}

	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hashing plain text with SHA-1.
func Encrypt(plaintext string)(cryptext string){
	cryptext = fmt.Sprintf("%x",sha1.Sum([]byte(plaintext)))
	return
}