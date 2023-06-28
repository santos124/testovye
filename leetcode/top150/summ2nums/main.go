package main

func main() {
	
}
// Мы хотим складывать очень большие числа, которые превышают емкость
// базовых типов, поэтому мы храним их в виде массива неотрицательных чисел.
// Нужно написать функцию, которая примет на вход два таких массива, 
// вычислит сумму чисел, представленных массивами, и вернет результат
// в виде такого же массива.

// # Пример 1
// # ввод
// arr1 = [1, 2, 3] # число 123
// arr2 = [4, 5, 6] # число 456
// # вывод
// res = [5, 7, 9] # число 579. Допустим ответ с первым незначимым нулем [0, 5, 7, 9]


// arr1 = [1, 2, 3] # число 123
// arr2 = [4, 5, 6] # число 456
//lr = 9, o = 0, r = [9, 7, 5]


// arr1 = [7, 2, 3] # число 123
// arr2 = [4, 5, 6] # число 456
//[9, 7, ]


// arr1 =  [7, 2, 3] # число 123
// arr2 =  [4, 5, 6] # число 456
//[9,7,1]

func calcArrs(arr1 []int, arr2 []int) []int {
    result := []int{}
    upper := []int{}
    lower := []int{}
    if len(arr1) > len(arr2) {
        upper = arr1
        lower = arr2
    } else {
        upper = arr2
        lower = arr1
    }

    ostatok := 0
    cnt := 0
    i2 := len(lower) - 1
    for i := len(upper) - 1; i >= 0; i-- {
        localResult := 0
        if cnt <= len(lower) {
            localResult = upper[i] + lower[i2] + ostatok / 10
            ostatok = 0
            if localResult > 9 {
                ostatok = localResult - localResult % 10
                localResult = localResult % 10
            }
            i2--
        } else {
            localResult = upper[i] + ostatok / 10
            ostatok = 0
            if localResult > 9 {
                ostatok = localResult - localResult % 10
                localResult = localResult % 10
            }
        }
        result = append(result, localResult)
    }
    if ostatok > 0 {
        result = append(result, ostatok / 10)
    }
    for i := 0; i < len(result) / 2; i++ {
        result[i], result[len(result - 1 - i)] = result[len(result - 1 - i)], result[i]
    }
    return result
}

// Дан массив целых чисел nums и целое число k. Нужно написать функцию, которая
// вынимает из массива k наиболее часто встречающихся элементов.
// Пример
// # ввод
// nums = [1,1,1,2,2,3]
// k = 2
// # вывод (в любом порядке)
// [1, 2]


// []
// [1]
// [1, 2]
// [1, 1, 1, 2, 3, 3, 3]
// []


func getMostWanted(nums []int,k int) []int{
    
    result := []int{}
    mapa := map[int]int{}
    maxCnt := 0
    mostWantedNum := 0
    //заполние мапы
    for _, num := range nums {
        mapa[num]++
    }
    // map[int]int{1:3, 2:2, 3:1} 

    //определение самого часто встречаемого числа
    
    for i := 0; i < k; i++ {
        for num, cnt := range mapa {
            if cnt > maxCnt {
                maxCnt = cnt
                mostWantedNum = num
            }   
        }
        delete(mapa, mostWantedNum)
        result = append(result, mostWantedNum)     
    }
    return result
}

