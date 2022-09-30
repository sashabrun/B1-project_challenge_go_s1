package main

import "os"

func BasicAtoi(s string) int {
	j := 0
	for i := 0; i < len(s); i++ {
		j = j*10 + (int(s[i]) - 48)
	}
	return j
}

func exam(long int, larg int) {
	counter := 0
	espace := " "
	col2 := counter + long - 2
	for i := 0; i < long; i++ {
		for j := 0; j < larg; j++ {
			if i == 0 && j == 0 {
				print("X")
			} else if i == 0 && j == larg-1 {
				print("X")
			} else if i == long-1 && j == 0 {
				print("X")
			} else if i == long-1 && j == larg-1 {
				print("X")
			} else if i == 0 || i == long-1 {
				print("-")
			} else if j == 0 {
				for x := 0; x < 1; x++ {
					if counter > 9 {
						counter = 0
					}
					print(counter)
					counter++
				}
				for z := 0; z < 1; z++ {
					for y := 0; y < larg-2; y++ {
						print(espace)
					}
					if col2 > 9 {
						col2 = 0
					}
					print(col2)
					col2++
				}
			} else {
				print(" ")
			}
		}
		print("\n")
	}
}
func main() {
	arg := os.Args[1:]
	if len(arg) == 0 || len(arg) == 1 {
		print("incomplete...")
		return
	} else if (BasicAtoi(arg[0]) == 0 || BasicAtoi(arg[1]) == 0) || (BasicAtoi(arg[0]) < 0 || BasicAtoi(arg[1]) < 0) {
		print("not possible...")
		return
	} else {
		exam(BasicAtoi(arg[0]), BasicAtoi(arg[1]))
	}
}
