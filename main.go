package main
import "fmt"
func main() {
    sliceMakeTest()
    sliceChangeTest()
    testAnd()
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
    fmt.Println("old array value:", array)
    bytes := array[:4]
    bytes[0] = 0
    fmt.Println("array value:", array)
    fmt.Println("bytes value:", bytes)

    temp := make([]int, 5)
    i := 0
    for i < len(temp) {
        temp[i] = i
        i++
    }
    fmt.Println("temp value:", temp)
    bytes = append(bytes, temp...)
    fmt.Println("after append bytes value:", bytes)
    bytes[0] = 100
    fmt.Println("after modify bytes value:", bytes)
    fmt.Println("after modify array value:", array)

    newBytes := bytes[:]

    newBytes[0]=1000
    fmt.Println("after modify newBytes, bytes value:", bytes)
    fmt.Println("after modify newBytes, newBytes value:", newBytes)

    newBytes = append(newBytes, 10000)
    fmt.Println("after modify newBytes, bytes value:", bytes)
    fmt.Println("after modify newBytes, newBytes value:", newBytes)

    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes = append(newBytes, 10000)
    newBytes[0]=888
    fmt.Println("after modify newBytes, bytes value:", bytes)
    fmt.Println("afger modify newBytes, newBytes value:", newBytes)

}

func testAnd()  {
    i := 6
    slice := make([]byte, 0)
    slice = append(slice, 0)
    slice = append(slice, 0)
    slice = append(slice, 0)
    newSlice := slice[1:1]
    fmt.Println("len of newSlice is: ", len(newSlice))
    if i == 5 && slice[3]==5 {
    } else {
        fmt.Println("go test && only exec successed condition")
    }

}