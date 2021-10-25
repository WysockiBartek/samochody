package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type car struct {
	grpNum int
	carNum string
	place  int
}

var cars = make([]car, 0, 25)
var gr1, gr2, gr3, gr4, gr5, gr6 []car
var podium []car
var indx int = 0
var debug bool = false

//=======================
const iloscTestow int = 10

//=======================
func main() {
	for i := 0; i < iloscTestow; i++ {
		time.Sleep(10 * time.Millisecond)
		doIt()
		indx = 0
		cars = cars[:0]
		gr1 = gr1[:0]
		gr2 = gr2[:0]
		gr3 = gr3[:0]
		gr4 = gr4[:0]
		gr5 = gr5[:0]
		gr6 = gr6[:0]
		podium = podium[:0]
	}
}
func doIt() {

	rand.Seed(time.Now().UnixNano())
	var firstExpected, secondExpected, thirdExpected string
	createCarsAndGroups()
	// get expected result
	for _, v := range cars {
		switch v.place {
		case 3:
			thirdExpected = v.carNum
		case 2:
			secondExpected = v.carNum
		case 1:
			firstExpected = v.carNum
		}
	}
	// pierwsze 5 okrazen
	race(gr1)
	race(gr2)
	race(gr3)
	race(gr4)
	race(gr5)
	// stwoz gr6
	gr6 = append(gr6, gr1[0], gr2[0], gr3[0], gr4[0], gr5[0])
	// okrazenie nr 6
	race(gr6)
	if debug {
		fmt.Printf("---- Groups sorted:\n%v\n%v\n%v\n%v\n%v\n", gr1, gr2, gr3, gr4, gr5)
		fmt.Println("---- Group 6:\n", gr6)
	}
	// dodaj naszybszy samochod do podium
	podium = append(podium, gr6[0])
	//---------------------------
	var grs [][]car
	grs = append(grs, gr1, gr2, gr3, gr4, gr5)
	var gr7 []car

	for i, v := range gr6 {
		switch i {
		case 0:
			gr7 = append(gr7, grs[v.grpNum-1][1], grs[v.grpNum-1][2])

		case 1:
			gr7 = append(gr7, grs[v.grpNum-1][0], grs[v.grpNum-1][1])
		case 2:
			gr7 = append(gr7, grs[v.grpNum-1][0])

		}

	}

	race(gr7)
	podium = append(podium, gr7[0], gr7[1])

	fmt.Println("---- Expected:\n", firstExpected, secondExpected, thirdExpected)
	fmt.Printf("----Result:\n %v %v %v\n", podium[0].carNum, podium[1].carNum, podium[2].carNum)
	fmt.Println("==============================")
	// fmt.Println(grs)
}

func createCarsAndGroups() {
	ranNum := rand.Perm(25)
	for i := 0; i < 25; i++ {
		s := "s" + strconv.Itoa(i)
		switch {
		case i <= 4:
			gr1 = append(gr1, car{1, s, ranNum[i] + 1})
		case i <= 9:
			gr2 = append(gr2, car{2, s, ranNum[i] + 1})
		case i <= 14:
			gr3 = append(gr3, car{3, s, ranNum[i] + 1})
		case i <= 19:
			gr4 = append(gr4, car{4, s, ranNum[i] + 1})
		case i <= 24:
			gr5 = append(gr5, car{5, s, ranNum[i] + 1})
		}
		cars = append(cars, car{1, s, ranNum[i] + 1})
	}
}
func race(grp []car) {
	sort.Slice(grp, func(i, j int) bool {
		return grp[i].place < grp[j].place
	})
}
