package binarysearch

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBinarySearchInt(t *testing.T) {
	var vals [100]int
	tests := make(map[int]int, 0)
	for i := 0; i < 100; i++ {
		vals[i] = i + 1
		tests[i+1] = i
	}
	tests[500] = -1
	tests[-1] = -1
	tests[-100] = -1
	tests[101] = -1
	for _, key := range tests {
		if key == 0 {
			continue
		}
		res := binarysearchInt(vals[:], key)
		if res != tests[key] {
			t.Errorf("looking for a number %d, we want an index %d, received %d", key, tests[key], res)
		}
	}
}

func TestGoLibrarySortSearchInt(t *testing.T) {
	var vals [100]int
	tests := make(map[int]int, 0)
	for i := 0; i < 100; i++ {
		vals[i] = i + 1
		tests[i+1] = i
	}
	tests[500] = 100
	tests[-1] = 0
	tests[-100] = 0
	tests[101] = 100
	for _, key := range tests {
		if key == 0 {
			continue
		}
		res := sort.Search(len(vals), func(i int) bool {
			return vals[i] >= key
		})
		if res != tests[key] {
			t.Logf("looking for a number %d, we want an index %d, received %d", key, tests[key], res)
		}
	}
}

func generateRandomPhoneNumber() string {
	return fmt.Sprintf("89%09d", rand.Intn(1_000_000_000))
}

func TestBinarySearchString(t *testing.T) {
	rand.NewSource(time.Now().Unix())
	phones := make([]string, 10)
	for i := 0; i < 10; i++ {
		phones[i] = generateRandomPhoneNumber()
	}
	sort.Strings(phones)
	fmt.Println("Сортированные номера телефонов:")
	for _, phone := range phones {
		fmt.Println(phone)
	}
	fmt.Println("Ищем...")
	fmt.Println(phones[binarySearch(len(phones), func(i int) bool {
		return phones[i] >= phones[0]
	})])
}

func TestGoLibrarySortSearchString(t *testing.T) {
	rand.NewSource(time.Now().Unix())
	phones := make([]string, 10)
	for i := 0; i < 10; i++ {
		phones[i] = generateRandomPhoneNumber()
	}
	sort.Strings(phones)
	fmt.Println("Сортированные номера телефонов:")
	for _, phone := range phones {
		fmt.Println(phone)
	}
	fmt.Println("Ищем...")
	fmt.Println(phones[sort.Search(len(phones), func(i int) bool {
		return phones[i] >= phones[0]
	})])
}
