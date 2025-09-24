package main

import (
	"fmt"
	"strconv"
	"time"
)

type PaymentInfo struct {
	ID          int
	Description string
	Usd         int
	Cancelled   bool
}

type PaymentModulInfoSlice struct {
	sl []PaymentInfo
}

func (p *PaymentModulInfoSlice) addSlice(pinf PaymentInfo) {
	p.sl = append(p.sl, pinf)
}

func (p *PaymentModulInfoSlice) GetSlice(id int) PaymentInfo {
	for _, val := range p.sl {
		if val.ID == id {
			return val
		}
	}
	return PaymentInfo{}
}

//-----------------------------------------------------------------------

type PaymentModulInfoMap struct {
	mp map[int]PaymentInfo
}

func (p *PaymentModulInfoMap) addMap(pinf PaymentInfo) {
	p.mp[pinf.ID] = pinf
}

func (p *PaymentModulInfoMap) getMap(id int) PaymentInfo {
	info, ok := p.mp[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

func main() {
	pSlice := PaymentModulInfoSlice{}
	
	pMAp := PaymentModulInfoMap{
		mp: make(map[int]PaymentInfo),
	}

	iteration := 100000

	before := time.Now()

	for i := 0; i < iteration; i++ {
		info := PaymentInfo{
			ID:          i,
			Description: strconv.Itoa(i),
		}

		pSlice.addSlice(info)
	}

	fmt.Println("slise add", time.Since(before))

	before = time.Now()

	for i := 0; i < iteration; i++ {
		info := PaymentInfo{
			ID:          i,
			Description: strconv.Itoa(i),
		}

		pMAp.addMap(info)
	}

	fmt.Println("map add", time.Since(before))

	//------------------------------------------------

	before = time.Now()

	for i := 0; i < iteration; i++ {
		inf := pSlice.GetSlice(i)
		_ = inf
	}

	fmt.Println("slise get", time.Since(before))

	before = time.Now()

	for i := 0; i < iteration; i++ {

		inf := pMAp.getMap(i)
		_ = inf
	}

	fmt.Println("map get", time.Since(before))

	// payInf := PaymentInfo{
	// 	ID:          10,
	// 	Description: "BURGER",
	// 	Usd:         5,
	// 	Cancelled:   false,
	// }

	// pSlice.addSlice(payInf)
	// pMAp.addMap(payInf)

	// pp.Println(pSlice)
	// pp.Println(pMAp)

	// i1 := pSlice.GetSlice(10)
	// i2 := pMAp.getMap(10)

	// fmt.Println(i1, i2)
}
