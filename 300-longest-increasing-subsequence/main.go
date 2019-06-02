package main

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func main() {
    s := []int{}

    println(lengthOfLIS(s))
}

func lengthOfLIS(nums []int) int {
    var m map[int]int = map[int]int{}

    return f(m, MinInt, 0, nums)
}

func f(m map[int]int, num int, index int, slice []int) int {
    if val, ok := m[index]; ok {
        return val
    }

    min := MaxInt
    s := []int{}

    for i := index; i < len(slice); i++ {
        if slice[i] > num && min > slice[i]{
            min = slice[i]
            s = append(s, 1 + f(m, slice[i], i+1, slice))
        }
    }

    var result int
    if len(s) == 0 {
        result = 0
    }
    result = getMaxInSlice(s)

    m[index] = result

    return result
}

func getMaxInSlice(slice []int) int {
    var m int
    for i, e := range slice {
        if i==0 || e > m {
            m = e
        }
    }
    return m
}