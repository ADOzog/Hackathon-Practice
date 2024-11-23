package main
import "fmt"
func DNFsort (arry [16]int) ([16]int){

 var Zero_count int
 var One_count int
 var Two_count int
 
 for i:= 0; i < 16; i++ { 
  switch arry[i]{
   case 0:
    Zero_count++
   case 1:
    One_count++
   case 2:
    Two_count++
  }// switch ends here
 }// for loop ends here 
 var sorted_array [16]int
 
 // this loop counts the 0s
 for n := 0; n< Zero_count; n++ {
 sorted_array[n] = 0 
 // fmt.Println("0")
 }


 // this loop counts the 1s
 for l := Zero_count; l< One_count + Zero_count; l++ {
 sorted_array[l] = 1   
 // fmt.Println("1")
 }
 
 // this loop counts the 2s
 for m := Zero_count + One_count ; m< Zero_count + One_count + Two_count; m++ {
 sorted_array[m] = 2 
 // fmt.Println("2")
 }
 
 // var sorted_array [16]int
 // sorted_array = append(Zero_Array, One_Array,Two_Array...)
 
 return sorted_array

}// function ends here



func main () {
 original_array := [16]int{ 1, 1, 0, 1, 2, 2, 1, 0, 0, 0, 0, 1, 1, 1, 1 , 0}
 var sorted_array [16]int
 sorted_array = DNFsort(original_array)
 fmt.Println("The original array",original_array)
 fmt.Println("The sorted array",sorted_array)




}
