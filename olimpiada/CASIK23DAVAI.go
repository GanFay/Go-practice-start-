package main

import (
	"fmt"
	"sort"
)

type casik struct {
	number int
	amount int
}

func main() {
	m := 1
	masive := []int{}
	casino := []casik{}
	for i := 0; i < 10; i++ {
		fmt.Scanln(&m)
		masive = append(masive, m)
		fmt.Println("Kalibrovka", masive)
	}
	raschet(&masive, &casino, false)
	sortcasik(&casino)
	fmt.Println("Kalibrovka", casino)
	raschet(&masive, &casino, true)
}

func raschet(m *[]int, n *[]casik, kalibrovka bool) {
	ll := 0
	i := 0
	if kalibrovka == false {
		i = 0
	} else {
		i = len(*m)
	}
	for {
		if kalibrovka == true {
			fmt.Scanln(&ll)
			*m = append(*m, ll)
		}
		for ; i < len(*m); i++ {
			num := (*m)[i]
			if len(*n) == 0 {
				*n = append(*n, casik{num, 1})
				continue
			}
			flag := false
			for k := 0; k < len(*n); k++ {
				if flag == true {
					break
				} else {
					if (*n)[k].number == num {
						(*n)[k].amount++
						flag = true
					}
				}
			}
			if flag == false {
				*n = append(*n, casik{num, 1})
				continue
			} else {
				continue
			}
		}
		if kalibrovka == true {
			sortcasik(n)
			fmt.Println(n)
			amount(n)

		}
		if kalibrovka == false {
			return
		}
	}
}

func amount(n *[]casik) {
	num := (*n)[0].amount
	num2 := (*n)[0].number
	for i := 0; i < len(*n); i++ {
		if (*n)[i].amount > num {
			num = (*n)[i].amount
			num2 = (*n)[i].number
		}
	}
	fmt.Println("Наиболее выпадаемое число:", num2, "Кол-во выпадений", num)
}

func sortcasik(n *[]casik) {
	sort.Slice(*n, func(i, j int) bool {
		return (*n)[i].amount > (*n)[j].amount
	})
}
