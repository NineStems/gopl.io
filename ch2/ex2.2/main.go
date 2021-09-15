package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		fund := Funt(t)
		kg := KG(t)

		fmt.Printf("%v = %v, %v = %v\n",
			kg, ToFunt(kg), fund, ToKG(fund))

		funt := Futes(t)
		metrs := Metrs(t)

		fmt.Printf("%v = %v, %s = %v\n",
			funt, ToFutes(metrs), metrs, ToMetrs(funt))


	}
}

type Funt float64

func ToFunt(v KG) Funt {
	f := v * 2.20462
	return Funt(f)
}

func (f Funt) String() string {
	res :=  fmt.Sprintf("%g funts", f)
	return res
}

type KG float64

func ToKG(v Funt) KG {
	kg := v/2.20462
	return KG(kg)
}

func (kg KG) String() string {
	return fmt.Sprintf("%g kg", kg)
}

type Futes float64

func ToFutes(v Metrs) Futes {
	f := v * 3.28084
	return Futes(f)
}

func (f Futes) String() string {
	return fmt.Sprintf("%g futes", f)
}

type Metrs float64

func ToMetrs(v Futes) Metrs {
	m := v / 3.28084
	return Metrs(m)
}

func (m Metrs) String() string {
	return fmt.Sprintf("%g metr", m)
}
