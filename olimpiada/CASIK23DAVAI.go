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
	var allnum []int
	var numsortall []casik

	for i := 0; i < 8; i++ {
		_, err := fmt.Scanln(&m)
		if err != nil {
			return
		}
		allnum = append(allnum, m)
		fmt.Println("Kalibrovka: [ПОСЛЕДНИЕ 8:]", allnum)
	}
	lastten := allnum
	raschet(&allnum, &numsortall, false, lastten)
	sortcasik(&numsortall)
	fmt.Println("Kalibrovka [ВИПАВШЫЕ ЧИСЛА: {ЧИСЛО; КОЛ-ВО_ВЫПАДЕНИЙ}:", numsortall)
	raschet(&allnum, &numsortall, true, lastten)
}

func raschet(m *[]int, n *[]casik, kalibrovka bool, lastten []int) {
	amountst := 8
	innum := 0
	i := 0
	if kalibrovka == false {
		i = 0
	} else {
		i = len(*m)
	}
	for {
		if kalibrovka == true {
			_, err := fmt.Scanln(&innum)
			if err != nil {
				return
			}

			*m = append(*m, innum)

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
			n2 := *n
			sortcasik2(&n2)
			lastten = append(lastten[1:], innum)
			fmt.Println("Все(sort am):", *n)
			fmt.Println("Все(sort num):", n2)
			fmt.Println("Последние 8 выпавших:", lastten)
			amountst++
			fmt.Println("Число ставок:", amountst)
			flagyard(lastten)
			amount(n)

		}
		if kalibrovka == false {
			return
		}
	}
}

func flagyard(lastten []int) {
	twoto1up := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36}
	twoto1middle := []int{2, 5, 8, 11, 14, 17, 20, 23, 26, 29, 32, 35}
	twoto1down := []int{1, 4, 7, 10, 13, 16, 19, 22, 25, 28, 31, 34}

	first := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	second := []int{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	rd3 := []int{25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}

	flag1st := false
	flag2nd := false
	flag3rd := false

	flag2to1up := false
	flag2to1middle := false
	flag2to1down := false

	for j := 0; j < len(lastten); j++ {
		for l := 0; l < 12; l++ {
			if lastten[j] == first[l] {
				flag1st = true
			}
			if lastten[j] == second[l] {
				flag2nd = true
			}
			if lastten[j] == rd3[l] {
				flag3rd = true
			}
			if lastten[j] == twoto1up[l] {
				flag2to1up = true
			}
			if lastten[j] == twoto1middle[l] {
				flag2to1middle = true
			}
			if lastten[j] == twoto1down[l] {
				flag2to1down = true
			}
		}
	}

	if flag1st == false {
		fmt.Println("ПЕРВЫЙ ЯРД НЕ ПАДАЛ ДОЛГО БРАТАНЧИК")
	}
	if flag2nd == false {
		fmt.Println("ВТОРОЙ ЯРД НЕ ПАДАЛ ДОЛГО БРАТАНЧИК")
	}
	if flag3rd == false {
		fmt.Println("ТРЕТИЙ ЯРД НЕ ПАДАЛ ДОЛГО БРАТАНЧИК")
	}

	if flag2to1up == false {
		fmt.Println("БОКОВОЙ ВЕРХНИЙ НЕ ПАДАЛ ДАВНО СУБО")
	}
	if flag2to1middle == false {
		fmt.Println("БОКОВОЙ ЦЕНТРАЛЬНЫЙ НЕ ПАДАЛ ДАВНО СУБО")
	}
	if flag2to1down == false {
		fmt.Println("БОКОВОЙ НИЖНИЙ НЕ ПАДАЛ ДАВНО СУБО")
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

func sortcasik2(n *[]casik) {
	sort.Slice(*n, func(i, j int) bool {
		return (*n)[i].number < (*n)[j].number
	})
}
