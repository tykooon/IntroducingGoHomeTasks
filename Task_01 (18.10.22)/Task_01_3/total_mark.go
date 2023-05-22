/*Задача 3.
Система оценок в США (в зависимости от процента от максимального 
количества баллов): A – 90-100%, B – 80-89%, C – 70-79%, D – 65-69%,
F – 64% и ниже. Программа запрашивает максимальное количество баллов
за одно задание, количество заданий, а затем – баллы за каждое задание. 
(Имеется в виду, что некий студент выполнил ряд заданий, и мы хотим 
оценить успеваемость этого студента). Программа выводит оценку в виде буквы.
*/

package main

import (
	"fmt"
)

func main() {

	var maxMark, markAmount int

	fmt.Println("Enter maximal mark value:")
	fmt.Scanln(&maxMark)
	if maxMark <= 0 {
		fmt.Println("Incorrect input")
		return
	}

	fmt.Println("Enter amount of marks:")
	fmt.Scanln(&markAmount)
	if markAmount <= 0 {	
		fmt.Println("Incorrect input")
		return
	}

	var sum int
	for mark, i:= 0, 0; i< markAmount; i++ {
		for {
			fmt.Printf("Enter mark %d\n", i+1)
			fmt.Scanln(&mark)
			if mark >=0 && mark <= maxMark {
				sum += mark	
				break
			} else {
				fmt.Println("Incorrect input.")
				fmt.Println("Mark should be positive integer less or equal than ", maxMark)				
			}
		}
		
	}

	percent:= float64(sum)/float64(markAmount)/float64(maxMark)
	mark:= ""

	switch {
	case percent <= .64:
		mark = "F"	
	case percent <= .69:
		mark = "D"
	case percent <= .79:
		mark = "C"
	case percent <= .89:
		mark = "B"
	default:
		mark = "A"
	}
	fmt.Println("Overall mark: ", mark)

}
