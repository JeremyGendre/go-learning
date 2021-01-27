package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("=== FOR ===")
    forFunc()
    fmt.Println("=== If-ELSE ===")
    ifElse()
    fmt.Println("=== SWITCH ===")
    switchFunc()
    fmt.Println("=== ARRAY ===")
    arrayFunc()
    fmt.Println("=== SLICE ===")
    slicesFunc()
    fmt.Println("=== MAP ===")
    mapFunc()
    fmt.Println("=== RANGE ===")
    rangeFunc()
    fmt.Println("=== FUNCTIONS ===")
    funcUsage()
}

func plus(a int, b int) int {
    return a + b
}

func plusPlus (a, b, c int) int {
    return a + b + c
}

func funcUsage() {
    res := plus(2,4)
    fmt.Println("2 + 4 =", res)

    res = plusPlus(3,1,7)
    fmt.Println("3 + 1 + 7 =", res)
}

func rangeFunc() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, value := range nums {
        sum += value
    }
    fmt.Println("sum:", sum)

    for i, value := range nums {
        if value == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}

func mapFunc() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string][]int{"foo": { 1, 6, 8 }, "bar": { 5, 2, 7 }}
    fmt.Println("map:", n)
}

func slicesFunc() {

    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

func arrayFunc() {

    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

func forFunc() {

    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        i = i + 1
        if i == 12 {
            break
        }
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}

func ifElse() {

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if num := 10; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

func switchFunc() {

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
        case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
    }

    switch time.Now().Weekday() {
        case time.Saturday, time.Sunday:
            fmt.Println("It's the weekend")
        default:
            fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
        case t.Hour() < 12:
            fmt.Println("It's before noon")
        default:
            fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
            case bool:
                fmt.Println("I'm a bool")
            case int:
                fmt.Println("I'm an int")
            case string:
                fmt.Println("I'm an string")
            default:
                fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
    whatAmI(ifElse)
}