package arrays

func Sum(numbers [] int) int{
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}

func SumAll(numbers ... [] int) (sum [] int){
    var sums []int

    for _, number := range numbers {
        sums = append(sums, Sum(number))
    }
    return sums
}

func SumAllTails(numbers ... [] int) (sum [] int){
    var sums []int

    for _, number := range numbers {
        if len(number) == 0 {
            sums = append(sums, 0)
        } else {
            tail := number[1:]
            sums = append(sums, Sum(tail))
        }
    }
    return sums
}
