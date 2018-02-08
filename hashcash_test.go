package hc

import (
	"bytes"
	"testing"
)

func TestKeyGeneration(t *testing.T) {
	keyGen := KeyGenerator()
	key := keyGen(4)
	if key == nil {
		t.Error("Nil was returned from the GetKey func")
	}
	if len(key) != 4 {
		t.Error("The key generated was not the requested length")
	}
	if bytes.Equal(key, make([]byte, 4)) {
		t.Error("The key array was not initialized with random values")
	}
	if bytes.Equal(key, keyGen(4)) {
		t.Error("GetKey is not generating unique keys")
	}
}

func TestHash(t *testing.T) {
	keyGen := KeyGenerator()
	key := keyGen(4)
	hcash1, cnt1 := Hash("Message for you sir...", 12, key)
	hcash2, _ := Hash("Message for you sir... But a different one!", 12, key)
	hcash3, cnt3 := Hash("Message for you sir...", 12, key)

	if bytes.Equal(hcash1, hcash2) {
		t.Error("Two different Messages hashed to the same value")
	}
	if bytes.Equal(hcash1, hcash3) == false {
		t.Error("The same message hashed to a different value")
	}
	if cnt1 != cnt3 {
		t.Error("The same message returned different counts")
	}
}

func TestHashValidation(t *testing.T) {
	keyGen := KeyGenerator()
	key := keyGen(4)
	hcash1, cnt1 := Hash("Message for you sir...", 12, key)
	hcash2, _ := Hash("Message for you sir... But a different one!", 12, key)

	v, err := Validate("Message for you sir...", cnt1, hcash1, 12, key)
	if v == false {
		t.Error("The provided hash did not validate")
	}
	v, err = Validate("Message for you sir... But a different one!", cnt1, hcash1, 12, key)
	if v {
		t.Error("A different message validated with the same hash")
	} else if err == "" {
		t.Error("Invalid hash returned no err")
	}
	v, err = Validate("Message for you sir...", cnt1, hcash2, 12, key)
	if v {
		t.Error("A different hash validated with the same message")
	} else if err == "" {
		t.Error("Invalid hash returned no err")
	}
	v, err = Validate("Message for you sir...", 426356, hcash2, 12, key)
	if v {
		t.Error("A different count validated with the same message")
	} else if err == "" {
		t.Error("Invalid hash returned no err")
	}
	v, err = Validate("Message for you sir...", cnt1, hcash2, 24, key)
	if v {
		t.Error("A hash validated with less zeros than required")
	} else if err == "" {
		t.Error("Invalid hash returned no err")
	}
}
