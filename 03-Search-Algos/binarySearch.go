package main

import "fmt"

func binarySearch(arr []int, num int) (int, bool) {

    start := 0
    end := len(arr) - 1

    for start <= end {
        mid := (start+end)/2

        if arr[mid] == num {
            return mid, true
        }
        if arr[mid] > num {
            end = mid - 1
        }
        if arr[mid] < num {
            start = mid + 1
        }
    }
    return -1, false
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

    index, isExist := binarySearch(arr, 5)
    if isExist {
        fmt.Printf("index : %d\n", index)
    }else {
        fmt.Println("Does not exist")
    }
}
