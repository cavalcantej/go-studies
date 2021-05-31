package main

import "fmt"

var x [5]float64

func main() {

ARRAYS	
	x[0] = 10
	x[1] = 5
	x[2] = 3 + 6
	x[3] = x[1] + x[2]
	x[4] = 100

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))

	y := [5]float64{ 78, 98, 99, 87, 65 }

	var total2 float64 = 0
	for _, value := range y {
		total2 += value
	}
	fmt.Println(total2 / float64(len(y)))

SLICES
	var s []float64
	
	fmt.Printf("%T", s)

	slice := make([]float64, 5, 10)
	fmt.Printf("%T", slice)

	[ low : high ] expression
	slice2 := arr[0:5]
	arr[0:len(arr)]
	arr[0:]
	arr[:5]
	arr[:]
	arr[0:len(arr)]
	
append & copy
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	fmt.Println("////////////////////////////")
	slice3 := make([]int, 2)
	slice3[0] = 10
	slice3[1] = 101
	copy(slice1, slice3)
	fmt.Println(slice1, slice3)


Maps

	x := make(map[string]int)
	x["key"] = 10
	
	fmt.Println(x["key"])

	elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

	elements2 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}


    fmt.Println(elements["Li"])
	name, ok := elements["Un"]

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	fmt.Println(name, ok)

	elements3 := map[string]map[string]string{
        "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium",
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium",
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium",
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
    }

    if el, ok := elements["Li"]; ok {
        fmt.Println(el["name"], el["state"])
    }


	y := make([]int, 3, 9)

	fmt.Printf("%v",y)


	x := [6]string{"a","b","c","d","e","f"}
	fmt.Println(x[2:5])

	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	var smaller int = x[0]
	for _, value := range x {
		if value < smaller{
			smaller = value
		} 
	}
	fmt.Println(smaller)
	
}

