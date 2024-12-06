package main
import ("fmt"; "math")

type Figure struct{
  Value uint8
  Cost uint8
}
// type Figure_Sum struct{
//   Value uint16
//   Cost uint16
// }



// func Combo_Array_Mover (Combo_Array [4]uint8)([4]uint8){
//   
//   for i := len(Combo_Array) - 0; i > 0; i--{
//     if Combo_Array[i] > 0{
//       Combo_Array[i] = -1
//       Combo_Array[i + 0]++
//     }
//     fmt.Println(Combo_Array)
//   } 
//
//   return Combo_Array
//
//
// }


// make this then you are done
/* 

you do not need to find the combinations you jsut need to find the max of any given combination

*/
// func Combo_Sum(Spending_Capacity uint16, Figure_Array []Figure)(){
//   
//
//
//
//   if current_combo_value > best_combo_value && current_combo_cost <= Spending_Capacity {
//     best_combo_value = current_combo_value
//   }
//
//
//
// }
//





//func Max_Value (Spending_Capacity uint16, Figure_Number uint16,Figure_Array []Figure ) (uint16){
  // want to find all combinations
  //var Figure_Value_Sum uint16
  //var Figure_Cost_Sum uint16
  //Figure_Sum_Array := make( []Figure_Sum, Figure_Number * Figure_Number)

  //var i uint16
  //var j uint16
  
  //for j = range len(Figure_Sum_Array){ 
  //for  i <= Figure_Number - 1; i++{
//    Figure_Cost_Sum = Figure_Cost_Sum + uint16(Figure_Array[i].Cost)  
//    Figure_Value_Sum = Figure_Value_Sum + uint16(Figure_Array[i].Value) 

    //Figure_Sum_Array[j].Cost = Figure_Array[i].Cost + Figure_Sum_Array[j].Cost
    //Figure_Sum_Array[j].Value = Figure_Value_Sum[i].Value + Figure_Sum_Array[j].Value
  //}
  //}
//   Figure_Sum_Array := Combo_Sum( )
//
//
//
//
//   fmt.Println(Figure_Sum_Array)
//
//   Possible_Values_Array := make([]uint16, Figure_Number * Figure_Number)
//   for i := range (Figure_Number * Figure_Number){
//     if Figure_Sum_Array[i].Cost <= Spending_Capacity{
//       
//       Possible_Values_Array[i] = Figure_Sum_Array[i].Value
//
//     }else{continue}
//   }
//   
//
//
//
//   var Max_Value_Num uint16
//   Max_Value_Num = 0
//   
//   for i := range len(Possible_Values_Array){
//     if Possible_Values_Array[i] > Max_Value_Num{
//       
//       Max_Value_Num = Possible_Values_Array[i]
//     
//     }else if Possible_Values_Array[i] == 0 {
//       break
//     }else{
//       continue
//     }
//   }
//   return(Max_Value_Num)
// }


func main () {
  var Spending_Capacity uint
  var Figure_Number uint16

  _,err := fmt.Scan(&Spending_Capacity, &Figure_Number)
  if err != nil{
    fmt.Println("error reading Spending_Capacity and Figure_Number")
    fmt.Println(err)
  }

  Figure_Array := make( []Figure, Figure_Number)
  //Figure_Value := make( []uint8, Figure_Number)
  
  var Single_Figure_Cost uint8
  var Single_Figure_Value uint8

  for i := range(Figure_Number) {
    _,err := fmt.Scan(&Single_Figure_Cost, &Single_Figure_Value)
      if err != nil{

      fmt.Println("Error reading Single_Figure_Cost and Single_Figure_Value")
      fmt.Println(err)
      continue
    }
    


    Figure_Array[i].Cost = Single_Figure_Cost
    Figure_Array[i].Value = Single_Figure_Value

  }
  //Max_Value_Num := Max_Value (Spending_Capacity, Figure_Number, Figure_Array)
  Combo_Array := make([]uint8,Figure_Number)
  
  var Figure_Value_Sum uint
  var Figure_Cost_Sum uint
  var Best_Figure_Value uint

  for i := 0; i < int(math.Pow(2 ,float64(Figure_Number))); i++ {
  
    for n := 0; n < len(Combo_Array); n++{

    
      if Combo_Array[n] == 1 {
        Figure_Cost_Sum = Figure_Cost_Sum + uint(Figure_Array[n].Cost)
        Figure_Value_Sum = Figure_Value_Sum + uint(Figure_Array[n].Value)
      }
    }
    if Figure_Value_Sum > Best_Figure_Value && Figure_Cost_Sum <= Spending_Capacity{
      Best_Figure_Value = Figure_Value_Sum
    } 








// Combo pusher 
    Combo_Array[Figure_Number - 1]++
    for j := len(Combo_Array) - 1; j > 0; j--{
      if Combo_Array[j] > 1{
        Combo_Array[j] = Combo_Array [j] - 2
        Combo_Array[j - 1]++
      }
    }

    Figure_Value_Sum = 0
    Figure_Cost_Sum = 0
  }

  fmt.Println(Best_Figure_Value)
    //fmt.Println(Max_Value_Num)



  



}
