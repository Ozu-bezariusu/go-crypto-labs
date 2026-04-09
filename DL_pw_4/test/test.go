package test

import (
	randomlesstests "dl_pw_4/randomless_tests"
	"fmt"
	"math/rand"
	"time"
)

func MonobitTesting() {
	rand.Seed(time.Now().UnixNano())

	data := make([]byte, 2500)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}

	if randomlesstests.MonobitTest(data) {
		fmt.Println("Послідовність випадкова за монобітним тестом \u2714")
	} else {
		fmt.Println("Послідовність не випадкова за монобітним тестом \u2716")
	}
}

func SeriesTesting() {
	rand.Seed(time.Now().UnixNano())

	data := make([]byte, 2500)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}

	if randomlesstests.SeriesTest(data) {
		fmt.Println("Послідовність випадкова за тестом максимальної довжини \u2714")
	} else {
		fmt.Println("Послідовність не випадкова за тестом максимальної довжини \u2716")
	}
}

func PokerTesting() {
	rand.Seed(time.Now().UnixNano())

	data := make([]byte, 2500)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}

	if randomlesstests.PokerTest(data) {
		fmt.Println("Послідовність випадкова за тестом Поккера \u2714")
	} else {
		fmt.Println("Послідовність не випадкова за тестом Поккера \u2716")
	}
}

func SeriesLengthTesting() {
	rand.Seed(time.Now().UnixNano())

	data := make([]byte, 2500)
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}

	if randomlesstests.SeriesLengthTest(data) {
		fmt.Println("Послідовність випадкова за тестом довжин серій \u2714")
	} else {
		fmt.Println("Послідовність не випадкова за тестом довжин серій \u2716")
	}

}
