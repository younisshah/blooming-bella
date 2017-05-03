package blooming_bella

import (
	"errors"
	"math"
	"github.com/spaolacci/murmur3"
	"encoding/binary"
	"reflect"
	"log"
)

/**
*
* author: Younis Shah
*
* An implementation of bloom filters
**/

type bella struct {
	probability    float64 // desired probability
	items          int     // number of items in collection
	bitArrayLength int     // bit array length
	bitArray       []int8
}

func NewBella(items int, probability float64) (bella, error) {
	if probability < 0.0 && probability > 1.0 {
		return bella{}, errors.New("Desired probability cannot be < 0.0 and > 1.0")
	} else {
		b := int(math.Floor(-((float64(items) * math.Log(probability))) / (math.Pow(math.Log(2), 2))))
		return bella{probability: probability, items: items, bitArrayLength: b, bitArray: make([]int8, b)}, nil
	}
}

func (b bella) Add(item int) {
	if marvin, err := marvin32(item, b.items); err != nil {
		log.Fatalln(err)
	} else {
		murmur := murmurHash(item, b.items)
		fastHash, err := sfh(item, b.items)
		if err != nil {
			log.Fatalln(err)
		}
		b.bitArray[murmur] = 1
		b.bitArray[marvin] = 1
		b.bitArray[fastHash] = 1
	}
}

func (b bella) Test(item int) (bool) {
	if marvin, err := marvin32(item, b.items); err != nil {
		log.Fatalln(err)
		return false
	} else {
		murmur := murmurHash(item, b.items)
		fastHash, err := sfh(item, b.items)
		if err != nil {
			log.Fatalln(err)
		}
		if b.bitArray[murmur] == 1 && b.bitArray[marvin] == 1 && b.bitArray[fastHash] == 1 {
			return true
		} else {
			return false
		}
	}
}

func murmurHash(item, n int) uint16 {
	murmur := murmur3.Sum32(getBytes(item))
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, murmur)
	return binary.LittleEndian.Uint16(b) % uint16(n)
}

func marvin32(item, n int) (uint16, error) {
	marvin := NewMarvin32(uint64(item))
	if _, err := marvin.Write(getBytes(item)); err != nil {
		return 0, err
	} else {
		hash := marvin.Sum32()
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, hash)
		return binary.LittleEndian.Uint16(b) % uint16(n), err
	}
}

func sfh(item, n int) (uint16, error) {
	fastHash := NewSuperFastHash()
	if _, err := fastHash.Write(getBytes(item)); err != nil {
		return 0, err
	} else {
		hash := fastHash.Sum32()
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, hash)
		return binary.LittleEndian.Uint16(b) % uint16(n), err
	}
}

func getBytes(i int) []byte {
	size := reflect.TypeOf(i).Size()
	var b []byte
	switch size {
	case 2:
		b = make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(i))
	case 4:
		b = make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32(i))
	case 8:
		b = make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(i))
	}
	return b
}
