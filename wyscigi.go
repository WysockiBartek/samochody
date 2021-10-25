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
const iloscTestow int = 20

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
	// znajdz drugi i trzeci najszybszy samochod poprzez podmienianie [0] samochodu w grupie
	// samochodem gr6[i] == v
	// for i, v := range gr6 {
	// 	//pomiń pierwszy smochod poniewaz znajduje sie na podium
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	// jezeli podium ma 3 samochody przerwij
	// 	if len(podium) == 3 {
	// 		break
	// 	}
	// 	// jezeli podium ma 2 samochody i najszybszy samochod w func 'final()' to ten sam co w gr6[2]
	// 	// zrob kolejne okrazenie, tym razem podmien pierwszy samochod w grupie ktorej znajduje sie zamochod
	// 	// z podium[1]
	// 	if indx != 0 {
	// 		v = gr6[i-1]
	// 		switch {
	// 		case indx == 1:
	// 			indx = final(gr1, v)

	// 		case indx == 2:
	// 			indx = final(gr2, v)

	// 		case indx == 3:
	// 			indx = final(gr3, v)

	// 		case indx == 4:
	// 			indx = final(gr4, v)

	// 		case indx == 5:
	// 			indx = final(gr5, v)

	// 		}
	// 	} else {
	// 		//normalne poszukiwanie samochodow
	// 		switch {
	// 		case gr6[0].grpNum == 1:
	// 			indx = final(gr1, v)

	// 		case gr6[0].grpNum == 2:
	// 			indx = final(gr2, v)

	// 		case gr6[0].grpNum == 3:
	// 			indx = final(gr3, v)

	// 		case gr6[0].grpNum == 4:
	// 			indx = final(gr4, v)

	// 		case gr6[0].grpNum == 5:
	// 			indx = final(gr5, v)

	// 		}
	// 	}
	// }
	fmt.Println("---- Expected:\n", firstExpected, secondExpected, thirdExpected)
	fmt.Printf("----Result:\n %v %v %v\n", podium[0].carNum, podium[1].carNum, podium[2].carNum)
	fmt.Println("==============================")
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

// jak to działa? są tylko dwa racjonalne wytłumaczenia
// jeden fizyka czarnej dziury...
// dwa czary
func final(gr []car, c car) int {
	grFastest := c
	gr = append(gr[:0], gr[1:]...)
	gr = append(gr, c)
	race(gr)
	// PS pewnie mogl bym to zrobic w switch-u ale nie chce mi sie przepisywac
	if debug {
		fmt.Println("========\n", gr)
	}
	if len(podium) == 2 && gr[0] == grFastest && indx == 0 {
		if debug {
			fmt.Println("Switch to group of 2-nd fastest car and run it again...")
		}
		return podium[1].grpNum
	} else if grFastest != gr[0] && len(podium) == 2 {

		podium = append(podium, gr[0])
	} else if grFastest == gr[0] {
		podium = append(podium, grFastest)
		if debug {
			fmt.Println("else if: +1")
		}
	} else if c == gr[0] || c == gr[1] && len(podium) < 2 {
		if debug {
			fmt.Println("else if: +2")
		}
		podium = append(podium, gr[0], gr[1])

	} else {
		if debug {
			fmt.Println("else: +1")
		}
		podium = append(podium, gr[0])

	}
	return 0
}
