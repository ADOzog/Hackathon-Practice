package main
import ("fmt" ; "math")



func Num_Len (C_num float64) (int , bool){
  var n int 
  for C_num > 0 {
    if C_num < 1 {
      return n , false 
    }else {
      n++
      C_num = C_num / 10
    }
  }
  return -1 , true
}

func Array_make (C_num float64, C_num_len int) ([]int){
  C_num_arry := make([]int, C_num_len)
  // has to index backwards here 
  for i := C_num_len - 1; i >= 0; i-- {
    if C_num < 1 {
      return C_num_arry
      break
    }else {
      Num_in := int(C_num) % 10
      C_num_arry[i] = Num_in
      C_num = math.Floor(C_num / 10)
    }
  }
  return C_num_arry
}
func Repeat_Check  (C_num_arry []int , C_num_len int, i int) (int){
  var counter int
  if i > 0 {
    if C_num_arry[i] == C_num_arry[i - 1] {
      i--
      counter = Repeat_Check(C_num_arry , C_num_len , i ) + 1
    }
  }
  return counter 
}





func Is_Tidy (C_num_arry []int , C_num_len int) (bool, int, int){
  var counter int
  for i := 0 ; i + 1< C_num_len; i++ {
    if C_num_arry[i + 1] < C_num_arry[i] {
      counter = Repeat_Check (C_num_arry, C_num_len , i)

      return false, i + 1, counter
    }
  }//for loop end
  counter = 0
  return true, C_num_len, counter
}// Is_Tidy end 




// if i check for the frist how many digits are tidy in my Is_Tidy function i can make this part go faster
func Find_Tidy (C_num_arry []int , C_num_len int, Num_tidy_digits int, Tidy_Repeats int) ([]int) {
  if Tidy_Repeats != 1 {
    // this is for the repeating digits cases
    // fmt.Println("had repeating digits before untidy")
    Num_tidy_digits = Num_tidy_digits - Tidy_Repeats + 1
    C_num_arry[Num_tidy_digits - 1]-- 
    for i := Num_tidy_digits; i < C_num_len ; i++{
      C_num_arry[i] = 9
    }
  }else{
    C_num_arry[Num_tidy_digits - 1]-- 
    for i := Num_tidy_digits; i < C_num_len ; i++{
      C_num_arry[i] = 9
    }
  } 
  return C_num_arry
}

func Array_unmake (Array []int, Array_len int)(int){
  var F_num int
  for i:=0; i < Array_len; i++{
    F_num = Array[i] * int( math.Pow(10,float64(Array_len - i - 1))) + F_num
  }
  return F_num
}



func main () {
  var C_num float64
  var C_num_len int
  var Err bool
  

  fmt.Println("What number do you want to check")
  fmt.Scan(&C_num)
  C_num_len , Err = Num_Len(C_num)
  _ = Err
//  fmt.Println("The number of digits is -",C_num_len)
//  fmt.Println("Was there an error? -", Err)

  C_num_arry := make([]int, C_num_len)
  C_num_arry = Array_make(C_num, C_num_len)
 // fmt.Println(C_num_arry)

  Tidy_Val, Num_tidy_digits, Tidy_Repeats := Is_Tidy(C_num_arry, C_num_len)
  Tidy_Repeats++
  fmt.Println ("The first", Num_tidy_digits, "digits were tidy of", C_num_len, "digits")
 // fmt.Println("Repeated digits before untidy", Tidy_Repeats)
  if Tidy_Val {
    fmt.Println("The number is tidy" )
  }else{
    Prev_tidy_array := make([]int, C_num_len)
    Prev_tidy_array = Find_Tidy(C_num_arry, C_num_len, Num_tidy_digits, Tidy_Repeats)
// fmt.Println(Prev_tidy_array)
    Prev_tidy := Array_unmake(Prev_tidy_array, C_num_len)
    fmt.Println("Previous tidy number was", Prev_tidy)


  }
}

