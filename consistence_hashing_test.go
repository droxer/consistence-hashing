package consistencehashing

import (
	"strconv"
	"testing"
)

func HashFunc(data []byte) uint32 {
	i, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}
	return uint32(i)
}

func TestHashing(t *testing.T) {

}

func TestConsistence(t *testing.T) {

}

func Beachmark(b *testing.B) {

}
