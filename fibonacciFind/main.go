package main

import ("fmt" ; "os"; "strconv")

func findFib(n int) int {
    if n == 0 || n == 1 {
        return 1
    } 

    return findFib(n - 1) + findFib(n - 2)
}

func main() {

    if len(os.Args) != 2 {
        fmt.Println("Please enter the number of fibbonacci numbers to return")
        return
    }

    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Please enter a number")
        return
    }
    
    for i := 0; i < n; i++ {
        fmt.Print(findFib(i), " ")
    }
}
