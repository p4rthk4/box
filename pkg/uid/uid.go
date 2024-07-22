package uid

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

func GetNewId() (string, error) {
    timestamp := time.Now().Unix()

    tsString := fmt.Sprintf("%d", timestamp)
    // if len(tsString) > 8 {
    //     tsString = tsString[len(tsString)-8:]
    // }

    randString, err := generateRandomString(8)
    if err != nil {
        return "", err
    }

    uniqueID := tsString + randString

    return uniqueID, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) (string, error) {
    var result strings.Builder
    charsetLen := big.NewInt(int64(len(charset)))
    
    for i := 0; i < length; i++ {
        charIndex, err := rand.Int(rand.Reader, charsetLen)
        if err != nil {
            return "", err
        }
        result.WriteByte(charset[charIndex.Int64()])
    }

    return result.String(), nil
}
