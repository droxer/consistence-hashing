package ch

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := NewMap(3, func(data []byte) uint32 {
		i, err := strconv.Atoi(string(data))
		if err != nil {
			panic(err)
		}
		return uint32(i)
	})

	//2, 12, 22, 4, 14, 24, 9, 19, 29
	hash.Add("2", "4", "9")

	fixtures := map[string]string{
		"2":  "2",
		"29": "9",
		"13": "4",
	}

	for k, v := range fixtures {
		if hash.Get(k) != v {
			t.Fatalf("expected %s, actual is %s", v, hash.Get(k))
		}
	}

	// 8, 18, 28
	hash.Add("8")

	fixtures["26"] = "8"

	for k, v := range fixtures {
		if hash.Get(k) != v {
			t.Fatalf("expected %s, actual is %s", v, hash.Get(k))
		}
	}

	// 4, 14, 24
	hash.Remove("4")
	delete(fixtures, "13")

	for k, v := range fixtures {
		if hash.Get(k) != v {
			t.Fatalf("expected %s, actual is %s", v, hash.Get(k))
		}
	}
}
