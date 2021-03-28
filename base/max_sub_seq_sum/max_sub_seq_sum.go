package main

import(
    "os"
    "bufio"
    "strings"
    "strconv"
    "fmt"
)

func input()(res []int) {
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
    s, _:= reader.ReadString('\n')
    resStr := strings.Split(strings.TrimRight(s, "\n"), " ")
    for _, str := range resStr {
        item, _ := strconv.Atoi(str)
        res = append(res, item)
    }
    return res
}

func main() {
    input := input()
    //MaxSubSeqSum1(input)        // O(N^3)
    //MaxSubSeqSum2(input)        // O(N^2)
    //MaxSubSeqSum3(input)        // O(N*lngN)
    //MaxSubSeqSum4(input)        // O(N)
    MaxSubSeqSum5(input)          // 同时输出left and right value
}

func MaxSubSeqSum1(input []int) {
    var maxsum, thissum int
    for i := 0;i < len(input);i++ {
        for j := i;j < len(input);j++{
            thissum = 0
            for k := i;k <= j;k++ {
                thissum = thissum + input[k]
            }
            if thissum > maxsum {
                maxsum = thissum
            }
        }
    }
    fmt.Printf("%d\n", maxsum)
}

func MaxSubSeqSum2(input []int) {
    var maxsum, thissum int
    for i := 0;i < len(input);i++ {
        thissum = 0
        for j := i;j < len(input);j++{
            thissum = thissum + input[j]
            if thissum > maxsum {
                maxsum = thissum
            }
        }
    }
    fmt.Printf("%d\n", maxsum)
}

func MaxSubSeqSum3(input []int) {
    maxsum := reverse(input)
    fmt.Printf("%d\n", maxsum)
}

func reverse(input []int) (res int) {
    if len(input) == 1 {
        if input[0] > 0{
            return input[0]
        } else {
            return 0
        }
    }
    leftSum := reverse(input[0:len(input)/2])
    rightSum := reverse(input[len(input)/2:len(input)])
    middleSum := middle(input)

    if leftSum > rightSum && leftSum > middleSum {
        return leftSum
    } else if rightSum > leftSum && rightSum > middleSum {
        return rightSum
    } else {
        return middleSum
    }
}

func middle(input []int) (sum int) {
    var left, right int
    var tempLeft, tempRight int

    for i := len(input)/2 - 1;i >= 0;i-- {
        tempLeft = tempLeft + input[i]
        if tempLeft > left {
            left = tempLeft
        }
    }

    for i := len(input)/2;i < len(input);i++ {
        tempRight = tempRight + input[i]
        if tempRight > right {
            right = tempRight
        }
    }
    return left + right
}

func MaxSubSeqSum4(input []int) {
    var maxsum, thissum int
    for _, in := range input {
        thissum = thissum + in
        if thissum > maxsum {
            maxsum = thissum
        } else if thissum < 0 {
            thissum = 0
        }
    }
    fmt.Printf("%d\n", maxsum)
}

func IsNotPos(input []int) bool {
    for _, v := range input{
        if v > 0 {
            return false
        }
    }
    return true
}

func IsContainZero(input []int) bool {
    for _, v := range input{
        if v == 0 {
            return true
        }
    }
    return false
}

func MaxSubSeqSum5(input []int) {
    if IsNotPos(input) {
        if IsContainZero(input) {
            fmt.Printf("0 0 0\n")
        } else {
            fmt.Printf("0 %d %d\n", input[0], input[len(input) - 1])
        }
        return
    }
    var maxsum, thissum int
    var left, right, tempLeft, tempRight int
    for index, in := range input {
        thissum = thissum + in
        tempRight = index
        if thissum > maxsum {
            maxsum = thissum
            left = tempLeft
            right = tempRight
        } else if thissum < 0 {
            thissum = 0
            tempLeft = index + 1
        }
    }
    fmt.Printf("%d %d %d\n", maxsum, input[left], input[right])
}
