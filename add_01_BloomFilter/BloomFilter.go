package main

//package add_01

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

//
// https://cipepser.hatenablog.com/entry/2017/02/04/090629

const size = 64
const n = 10 // expected number of elements

var (
	// k is the number of hash functions.
	k int = int(math.Log(2) * float64(size) / float64(n))
)

type BloomFilter struct {
	BloomFilter [size]bool
}

func GetMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}

func DoubleHashing(hashA, hashB int64, n int) (hash int64) {
	// h = hashA + n * hashBの計算
	h := new(big.Int).Mul(big.NewInt(int64(n)), big.NewInt(hashB))
	h = new(big.Int).Add(big.NewInt(hashA), h)
	h = new(big.Int).Rem(h, big.NewInt(int64(size)))

	// 余りが負の数になったときは正の余りにする
	hash = h.Int64()
	if hash < 0 {
		hash += int64(size)
	}
	return
}

func Add(bf *BloomFilter, element string) {
	hash := GetMD5Hash(element)
	hashA := hash[:int(len(hash)/2)] // 前半
	hashB := hash[int(len(hash)/2):] // 後半

	i64_hashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64_hashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		bf.BloomFilter[DoubleHashing(i64_hashA, i64_hashB, i)] = true
	}
}

func Exists(bf *BloomFilter, element string) (exists bool) {
	hash := GetMD5Hash(element)
	hashA := hash[:int(len(hash)/2)] // 前半
	hashB := hash[int(len(hash)/2):] // 後半

	i64_hashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64_hashB, _ := strconv.ParseInt(hashB, 16, 64)

	exists = true
	for i := 0; i < k; i++ {
		if bf.BloomFilter[DoubleHashing(i64_hashA, i64_hashB, i)] == false {
			exists = false
			break
		}
	}
	return
}

func main() {
	var bf BloomFilter

	// 要素の追加
	Add(&bf, "1")
	Add(&bf, "2")
	Add(&bf, "3")
	Add(&bf, "4")
	Add(&bf, "5")

	// 要素が含まれているか検証
	fmt.Println(Exists(&bf, "1")) // true
	fmt.Println(Exists(&bf, "2")) // true
	fmt.Println(Exists(&bf, "3")) // true
	fmt.Println(Exists(&bf, "4")) // true
	fmt.Println(Exists(&bf, "5")) // true

	fmt.Println(Exists(&bf, "10")) // false

	fmt.Println(Exists(&bf, "999")) // true(偽陽性)
}
