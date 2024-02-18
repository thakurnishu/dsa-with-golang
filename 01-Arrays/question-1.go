package main

import (
	"fmt"
)

/*
   WAP to Read 10 no in array
       a) Display sum of all
       b) Find occurance of no 'x'
       c) Count the frequency of no 'x'
       d) Count frequency of each no in array
           Note One no if counted once must not be repeated in frequency
*/

func main() {
    array := []int{}

    fmt.Println("Enter 5 Number")
    for i := 0; i < 5; i++{
        var val int
        fmt.Scan(&val)
        array = append(array, val)
    }
    fmt.Println()

    fmt.Printf("Sum of All: %d\n", sumOfAll(array))

    fmt.Printf("Occurrence of %d: %v\n", 5, occurrenceOfNum(array, 5))

    fmt.Printf("Frequency of %d: %d\n", 5, frequencyOfNum(array, 5))

    fmt.Println("Frequency of Every Number: ")
    fmt.Println("-----------------------------")
    fmt.Printf("Number\t:\tFrequency\n")
    fmt.Println("-----------------------------")
    for key, val := range frequencyOfEveryNum(array) {
        fmt.Printf("%d\t:\t%d\n", key, val)
    }

}

func sumOfAll(arr []int) int {
    var sum int
    for _, val := range arr {
        sum = sum + val
    }
    return sum
}

func occurrenceOfNum(arr []int, num int) bool {
    for _, val := range arr {
        if val == num {
            return true
        }
    }
    return false
}

func frequencyOfNum(arr []int, num int) int {
    var count int
    for _, val := range arr {
        if val == num {
            count = count + 1
        }
    }
    return count
}

func frequencyOfEveryNum(arr []int) map[int]int {
    freqOfNum := make(map[int]int)
    for _, val := range arr {
            freqOfNum[val]++
    }

    return freqOfNum
}
