package main

import (
    "fmt"
    "strconv"
)
func main() {
    sliceMakeTest()
    sliceChangeTest()
    testAnd()
    testStrconvAtoi()
}

func sliceMakeTest() {
    vchxxx := []byte{0}
    vchxx1 := vchxxx[1:]
    fmt.Println("when a slice is made from an old slice and the start index is len of old slice, " +
        "the len of new slice is: ", len(vchxx1))
    // test construct slice with equal start and end  position of original slice
    testBytes := make([]byte, 2)
    testNewBytes := testBytes[1:len(testBytes)-1]
    fmt.Println("when a slice is made from an old slice and the start index is equal to end index, " +
        "the len of new slice is: ", len(testNewBytes))
}

func sliceChangeTest() {
    // test construct slice with array/slice, changing new slice will also have efect to original array/slice
    array := [5]int{1,2,3,4,5}
    bytes := array[:4]
    bytes[0] = 0
    if array[0] != 0 {
        fmt.Println("when modify a slice which is made from old slice, the old slice is not changed" +
            " if the new slice is not realloc memory")
    } else {
        fmt.Println("when modify a slice which is made from old slice, the old slice is also changed" +
            " if the new slice is not realloc memory")
    }

    temp := make([]int, 5)
    i := 0
    for i < len(temp) {
        temp[i] = i
        i++
    }
    bytes = append(bytes, temp...)
    bytes[0] = 100
    if array[0] != 100 {
        fmt.Println("when modify a slice which is made from old slice, the old slice is not changed" +
            " if the new slice is realloced memory")
    }

}

func testAnd()  {
    i := 6
    slice := make([]byte, 0)
    slice = append(slice, 0)
    slice = append(slice, 0)
    slice = append(slice, 0)

    if i == 5 && slice[3]==5 {
    } else {
        fmt.Println("if a condition inculde '&&' opcode, the condition after '&&' exec only " +
            "if the condition before '&&' exec successed")
    }

}

func testStrconvAtoi() {
    var strTest string
    _, err := strconv.Atoi(strTest)
    if err == nil {
        fmt.Println("strconv.Atoi a empty string should get an error")
    }
    result, err := strconv.Atoi("0")
    if result != 0 {
        fmt.Println("strconv.Atoi a \"0\" string should get an zero value")
    }
}