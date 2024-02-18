package main

import "fmt"


func linearSearch(arr []int, num int)(int, bool) {
    for i := 0; i < len(arr); i++ {
        if arr[i] == num {
            return i, true
        }
    }
    return -1, false
}

func main() {

    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

    index, isExist := linearSearch(arr, 5)
    if isExist {
        fmt.Printf("index : %d\n", index)
    }else {
        fmt.Println("Does not exist")
    }

}
