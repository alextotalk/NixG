package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println("____________")

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println("y =", total/float64(len(y)))
	fmt.Println("____________")
	var z [5]float64
	z[0] = 98
	z[1] = 93
	z[2] = 77
	z[3] = 82
	z[4] = 83
	var total2 float64 = 0
	for _, value := range z {
		total2 += value
	}
	fmt.Println("z =", total2/float64(len(z)))
	fmt.Println("____________")
	s := [5]float64{98, 93, 77, 82, 83}
	fmt.Println("s =", s)

	fmt.Println("____________")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	fmt.Println("____________")
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	fmt.Println("____________")
	var m = make(map[string]int)
	m["key"] = 10
	fmt.Println(m, "m[\"key\"]reset jetbrains ide evals=", m["key"])

	fmt.Println("____________")
	f := make(map[string]int)
	f["key"] = 10
	fmt.Println("f==", f["key"])

	fmt.Println("____________")
	d := make(map[int]int)
	d[1] = 10
	fmt.Println("d==", d[1])
	delete(d, 1)
	fmt.Println("d==", d[1])

	fmt.Println("____________")
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

	fmt.Println(elements["Li"])
	fmt.Println(elements["UK"])
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	if name, ok := elements["H"]; ok {
		fmt.Println(name, ok)
	}
	if _, ok := elements["uk"]; ok == false {
		fmt.Println("ok==false;", ok)
	}
	fmt.Println("____________")

	//shot   declaration
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
	fmt.Println(elements2)

	fmt.Println("____________")

	elements3 := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements3["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	fmt.Println("____________")

	l := make([]int, 3, 9)
	fmt.Println(l, len(l))

	fmt.Println("____________")

	w := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(w[2:5])

	fmt.Println("____________")

	//Напишите программу, которая находит самый
	//наименьший элемент в этом списке:

	els := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	minEl := els[0]
	for _, vol := range els {
		if minEl > vol {
			minEl = vol
		}
	}
	fmt.Println(minEl)
}
