package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func getKey(len int) []byte {
	token := make([]byte, 4)
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Read(token)
	return token
}

func hash(message string, zeros int, key []byte) ([]byte, uint64) {
	var cnt uint64
	cnt, mac := 0, hmac.New(sha256.New, key)
	mac.Write([]byte(message + ":" + string(cnt)))
	var sum = mac.Sum(nil)
	for checkZeros(sum, zeros) == false {
		if cnt%1000000 == 0 {
			fmt.Print(".")
		}
		cnt++
		mac.Reset()
		mac.Write([]byte(message + ":" + strconv.FormatUint(cnt, 10)))
		sum = mac.Sum(nil)
	}
	fmt.Println()
	return sum, cnt
}

func validate(message string, cnt int, sum []byte, zeros int, key []byte) (valid bool, err string) {
	valid, mac := true, hmac.New(sha256.New, key)
	if checkZeros(sum, zeros) == false {
		valid = false
		err = "Error: The provided sum failed to prove sufficient work."
	} else {
		mac.Write([]byte(message + ":" + string(cnt)))
		valid = bytes.Equal(mac.Sum(nil), sum)
		if valid == false {
			err = "Error: The provided and calculated sums did not match."
		}
	}
	if valid == false {
		println(err)
	} else {
		println("Message is VALID")
	}
	return
}

func checkZeros(a []byte, bits int) bool {
	valid, zB, zb := true, bits/8, bits%8
	for i := 0; i < zB; i++ {
		valid = valid && a[i] == 0
	}
	return valid && int(a[zB]) < int(math.Pow(2, float64(8-zb)))
}
