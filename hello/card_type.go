package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type ListCard []string

func (d ListCard) printAll() {
	for i, card := range d {
		fmt.Println(card, i)
	}
}

func newCard() ListCard {
	newCard := []string{}

	cardTypes := [5]string{"sức mạnh", "dũng cảm", "tự tin", "ma lực", "khoai bé"}

	point := [5]float32{
		float32(rand.Float32()) * 100,
		float32(rand.Float32()) * 100,
		float32(rand.Float32()) * 100,
		float32(rand.Float32()) * 100,
		float32(rand.Float32()) * 100,
	}

	num := 1
	for _, t := range cardTypes {
		for _, v := range point {
			newCard = append(newCard, t+" "+strconv.FormatFloat(float64(v), 'f', 2, 32)+" điểm"+" và vị trí thứ"+" "+strconv.Itoa(num))
			num += 1
		}
	}
	return newCard
}

func slice(start, end int, value ListCard) (ListCard, ListCard) {

	if len(value) == 0 {
		return ListCard{}, ListCard{}
	}

	if start < 0 || end < 0 {
		return ListCard{}, ListCard{}
	}

	if end > len(value) {
		end = len(value)
	}

	take := ListCard{}

	remain := ListCard{}

	for idx, t := range value {
		if idx >= start && idx < end {
			take = append(take, t)
		}
		remain = append(remain, t)
	}

	return take, remain
}

func deal(array ListCard, value int) (ListCard, ListCard) {

	return array[:value], array[value:]
}

func (d ListCard) toString() string {
	return strings.Join(d, ",")
}

func (d ListCard) saveToFile(name string) error {
	return os.WriteFile(name, []byte(d.toString()), 0666)
}

func readTofFile(name string) ListCard {
	bs, err := os.ReadFile(name)

	if err != nil {
		fmt.Println("Oh no error: ", err)
		os.Exit((1))
	}
	return ListCard(strings.Split(string(bs), ","))
}

func (d ListCard) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// rand.Shuffle(len(d), func(i, j int) {
	// 	d[i], d[j] = d[j], d[i]
	// })

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
