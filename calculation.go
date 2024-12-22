package calculation

import ("errors"
    "strconv"
    "strings")

func PairInArray(pair []int, array [][]int) bool {
  for _, brackets := range array {
    if brackets[0] == pair[0] && brackets[1] == pair[1] {
      return true
    }
  }

  return false
}

func SymbolInString(el rune, source []rune) bool {
  for _, x := range source {
    if x == el {
      return true
    }
  }
  return false
}

func LastElement(array []string) string {
  return array[len(array)-1]
}

func IsBracketsCorrect(source string) bool {
  count := 0
  for _, el := range source {
    if count < 0 {
      return false
    }
    if el == '(' {
      count++
    } else if el == ')' {
      count--
    }
  }
  return count == 0
}

func IsAllowedSymboles(symbol rune, array []rune) bool {
  flag := false
  for _, el := range array {
    if symbol == el {
      flag = true
      break
    }
  }
  return flag
}

func IsStringCorrect(source string, array []rune) bool {
  for _, el := range source {
    if !IsAllowedSymboles(el, array) {
      return false
    }
  }
  return IsBracketsCorrect(source)
}

func OpenBracketsIndexes(source string) ([]int, error) {
  x := make([]int, 0)
  for index, el := range source {
    if el == '(' {
      x = append(x, index)
    }
  }
  if len(x) != 0 {
    return x, nil
  }
  return x, errors.New("error")
}

func CloseBracketsIndexes(source string) ([]int, error) {
  x := make([]int, 0)
  for index, el := range source {
    if el == ')' {
      x = append(x, index)
    }
  }
  if len(x) != 0 {
    return x, nil
  }
  return x, errors.New("error")
}

func MinLength(element, start int, open []int) int {
  count := 0
  for _, open_br := range open[start:] {
    if element > open_br {
      count += 1
    } else {
      break
    }
  }
  return count
}

func PairsBracketsIndexes(source string, open []int) ([][]int, error) {
  result := make([][]int, 0)
  if len(open) == 0 {
    return nil, errors.New("error")
  }
  for _, open_br := range open {
    close_br := open_br + 1
    count := 1
    for count != 0 {
      if source[close_br] == '(' {
        count++
      } else if source[close_br] == ')' {
        count--
      }
      close_br++
    }
    close_br--
    pair := []int{open_br, close_br}
    result = append(result, pair)
  }
  return result, nil
}

func IsSubstringHaveBrackets(source string, pair []int) bool {
  start_i := pair[0] + 1
  end_i := pair[1]
  for _, symbol := range source[start_i:end_i] {
    if symbol == '(' || symbol == ')' {
      return false
    }
  }
  return true
}

func OnlySimpleBreakets(source string, array [][]int) [][]int {
  result := make([][]int, 0)
  for _, pair := range array {
    if IsSubstringHaveBrackets(source, pair) {
      if !PairInArray(pair, result) {
        result = append(result, pair)
      }
    }
  }
  return result
}

func Operation(array []float64, operation rune) (float64, error) {
  if len(array) < 2 {
    return 0, errors.New("error")
  }
  number1 := array[len(array)-2]
  number2 := array[len(array)-1]
  if operation == '+' {
    return number1 + number2, nil
  } else if operation == '-' {
    return number1 - number2, nil
  } else if operation == '*' {
    return number1 * number2, nil
  } else if operation == '/' {
    if number2 != 0 {
      return number1 / number2, nil
    }
    return 0, errors.New("error")
  }
  return 0, errors.New("error")
}

func IsSymbolDigit(symbol rune) bool {
  return symbol >= '0' && symbol <= '9'
}

func CalculateExpression(source string) (float64, error) {
  numbers := make([]float64, 0)
  operations := make([]rune, 0)
  signs := []rune{
    '+',
    '-',
    '*',
    '/',
  }
  rating := map[rune]int{
    '+': 1,
    '-': 1,
    '/': 2,
    '*': 2,
  }
  var last_digit string = "_"
  var last_sign rune = '_'
  var last_symbol rune = '_'
  count_numbers := 0
  count_operations := 0
  for _, symbol := range source {
    if IsSymbolDigit(symbol) || symbol == '.' {
      if last_digit == "_" {
        numbers = append(numbers, float64(symbol-'0'))
        count_numbers++
        last_digit = string(symbol)
      } else if IsSymbolDigit(last_symbol) || last_symbol == '.' {
        x := last_digit + string(symbol)
        n, _ := strconv.ParseFloat(x, 64)
        count_numbers--
        numbers[count_numbers] = float64(n)
        count_numbers++
        last_digit = x
      } else if symbol == '.' {
        last_digit = last_digit + string(symbol)
      } else {
        numbers = append(numbers, float64(symbol-'0'))
        last_digit = string(symbol)
        count_numbers++
      }
    } else if SymbolInString(symbol, signs) {
      if len(numbers) < 1 {
        return numbers[0], errors.New("error")
      } else if len(operations) == 0 {
        operations = append(operations, symbol)
        count_operations++
        last_sign = symbol
      } else {
        if rating[last_sign] >= rating[symbol] {
          if len(numbers) < 2 {
            return 0, errors.New("error")
          }
          result, err := Operation(numbers, last_sign)
          if err != nil {
            return 0, errors.New("error")
          }
          count_numbers -= 2
          numbers = numbers[:count_numbers+1]
          numbers[count_numbers] = result
          count_numbers++
          count_operations--
          operations = operations[:count_operations]
          operations = append(operations, symbol)
          if len(numbers) > 1 && len(operations) > 1 {
            if rating[last_sign] >= rating[symbol] {
              result2, _ := Operation(numbers, operations[len(operations)-2])
              count_numbers -= 2
              numbers = numbers[:count_numbers+1]
              numbers[count_numbers] = result2
              count_numbers++
              operations = operations[count_operations:]
              count_operations--
            }
          }
          last_sign = symbol
          count_operations++
        } else {
          operations = append(operations, symbol)
          count_operations++
          last_sign = symbol
        }
      }
    }
    last_symbol = symbol
  }
  j := len(numbers) - 1
  for i := len(operations) - 1; i >= 0; i-- {
    result, err := Operation(numbers, operations[i])
    if err != nil {
      return 0, errors.New("error")
    }
    numbers = numbers[:j]
    j--
    numbers[j] = result
    operations = operations[:i]
  }
  return numbers[0], nil
}

func Calc(source string) (float64, error) {
  array := []rune{'+', '-', '*', '/', '(', ')', ' '}
  for _, el := range "0123456789" {
    array = append(array, el)
  }
  if !IsStringCorrect(source, array) {
    return 0, errors.New("error")
  }
  source = strings.ReplaceAll(source, " ", "")
    if len(source) == 0 {
        return 0, errors.New("error")
    }
  for {
    x, _ := OpenBracketsIndexes(source)
    z, _ := PairsBracketsIndexes(source, x)
    arr := OnlySimpleBreakets(source, z)
    if len(arr) != 0 {
      for i := len(arr) - 1; i >= 0; i-- {
        start_br := arr[i][0]
        end_br := arr[i][1]
        substring := source[start_br+1 : end_br]
        result, _ := CalculateExpression(substring)
        substitution := strconv.FormatFloat(result, 'f', 2, 64)
        source = source[:start_br] + substitution + source[end_br+1:]
      }
    } else {
      result, err := CalculateExpression(source)
      if err != nil {
        return 0, errors.New("error")
      }
      return result, err
    }
  }
}
