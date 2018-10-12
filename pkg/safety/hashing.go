package safety

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

// AutoPasswordHash is the same as HashPassword but generates a hash for you
func AutoPasswordHash(password string) (hash string, salt string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	salt = strconv.FormatInt(r.Int63(), 10)
	hash = HashPassword(password, salt)
	return hash, salt
}

// HashPassword hashes a password
// This hashes the password using pbkdf2 with sha512
func HashPassword(password string, salt string) string {
	return string(pbkdf2.Key([]byte(password), []byte(salt), 4096, 512, sha1.New))
}

// ValidHashTypes are valid hashing types
var ValidHashTypes = []string{
	"md5",
	"sha512",
	"sha256",
	"pbkdf2",
}

// Hash a string of text with a algorithm
// Options: md5, sha512, sha256
func Hash(algorithm string, toHash string) string {
	data := []byte(toHash)
	switch algorithm {
	case "md5":
		return fmt.Sprintf("%x", md5.Sum(data))
	case "sha512":
		return fmt.Sprintf("%x", sha512.Sum512(data))
	case "sha256":
		return fmt.Sprintf("%x", sha256.Sum256(data))
	case "pbkdf2":
		return HashPassword(toHash, "")
	default:
		// if a algorithm is not a valid option fallback on sha256
		return fmt.Sprintf("%x", sha256.Sum256(data))
	}
}
