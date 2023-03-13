package math_calc


func Add(numbers ...int) (sum int) {
    for _, v := range numbers {
        sum += v
    }
    return sum
}
