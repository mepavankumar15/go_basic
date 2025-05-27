package main

import (
	"fmt"
)

/*
owner struct {
  name string 
  age int64
}
type wifiSpeed struct {
  payment int64
  mbps int64
  ownerInfo owner // or owner 
}

// ^ where the items can be called in wifi.name and wifi.age


type detailsId struct {
  name string
  age int64
  College string
}

func (e detailsId) details() string {
  e.College  = " aditya"
  e.name = "pavankumar"

  dels := fmt.Sprintf("the name is %v and age is %v and college is %v", e.name, e.age, e.College)
  return dels
} 

//INTERFACES
type lan interface{
  details() string
}

func printDetails(l lan){
  if l != nil {
    fmt.Println(l.details())
  }else {
    fmt.Println("the details are nil")
  }
}
*/

// var Nullify int = 5252


func main(){
 /* var intSlice = []int{1,2,3}
  var floatSlice = []float64{1.2,2.3,3.4}
 var stringSlice = []string{"hello", "world"}

  fmt.Println("the sum of the int val :",sumslice[int](intSlice))
  fmt.Println("the sum of the float val :",sumslice[float64](floatSlice))
  fmt.Println("the sum of the string val :",sumslice[string](stringSlice))


*/
    
}


/*
func sumslice[T int | float64 | string](slice []T) T {
  var sum T
  for _, v := range slice {
    sum += v
  }
  return sum
} 

*/

  // POINTERS
  /* var p *int32
  var address int32 = 10
  p = &address
  val := *p
  fmt.Println("the value of p is : ", p, val)
  */

  //GO ROUTINES or concurrency
  /* for _, v := range dbData {
    wg.Add(1)
    go saveDb(v)
  }
  wg.Wait()
  fmt.Println("the results are : ", results) */
// CHANNELS
  
  /* var Nullfee = make(chan string)
  var inputNull = []string {"svethakesh", "Kala", "aprimana", "shunya bhiti"}
  for i := range inputNull {
    go checkInput(inputNull[i], Nullfee)
  }
  for i := range Nullfee{
    fmt.Println(i)
  } 
  */
 // the main end 2 }

/*
func checkInput(s string, Nullfee chan string){
  for {
    time.Sleep(time.Second*1)
    var Inputval = rand.Intn(10000)
    if Inputval >= Nullify{
      fmt.Println(" the input is : ", Inputval)
      fmt.Println("the input is : ", s)
      Nullfee <- s
      break
    }
  }
  close(Nullfee)
} 
*/

/*  func saveDb(s string) {
  time.Sleep(2 * time.Second)
  w.Lock()
  results = append(results, s)
  w.Unlock()
  wg.Done()
}
  */

  // STRUCTS and Embedded structs
  /*
var myWifi wifiSpeed = wifiSpeed{1000,100, owner{"pavan", 23}}
  fmt.Println("the paymnet: ", myWifi.payment," the mbps: ", myWifi.mbps)
  myWifi.payment = 1000
  myWifi.mbps = 100 
  fmt.Printf("the owner %v with age %v had paid about %v for %v mbps \n", myWifi.ownerInfo.name, myWifi.ownerInfo.age, myWifi.payment, myWifi.mbps)

  var myDetails detailsId = detailsId{"pavan", 23, "aditya"}
  fmt.Println(" the details are : ", myDetails.details())
*/
  //Interfaces

 /* var myEntity detailsId = detailsId{"pavan", 24, "aditya"}
  printDetails(myEntity)
  */

  /* var (
    c float64 = 0.3
    b bool = true
    s string = "hello"
  )
  var a int = 10 */

 /* a:= 2323
  b:= 23.2
  c:=true
  d:= "bye"

  fmt.Printf(" the values are \n %T %T %T %T \n",a,b,c,d) */

  // ARRAYS
  
 /* var arr = [3]string{"hell" , "55.2", "55.3" }
  for i:= 0; i < len(arr); i++ {
    fmt.Println("arr[i]", arr[i])
  } 
  */
  
  /* var intArr [3]int32 
  intArr[1] = 10
  intArr[0] = 12
  intArr[2] = 13
  fmt.Println("intArr", intArr[0]) */

  
  // ARRAYS OPtional
  /* Deploy := [...]string{"hell" , "55.2", "55.3" }
    for index , option := range Deploy {
      fmt.Println(index , option)
    } */

  // Slices

 /* languages := [9]string{
    "C", "Lisp", "C++", "Java", "Python",
    "JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
  } 
  */
  
  // slice 
  
  /* 
  intslice  := []int32 { 1,2,3}
  fmt.Println("sliced : ", intslice)
  intslice = append(intslice, 23)
  fmt.Print("slice is : ", intslice , "\n")

  var intslice2  = []int32 { 24,25}
  intslice = append(intslice, intslice2...)
  fmt.Println("slice is : ", intslice) */
  
  /* Define slices */
  
  /* classics := languages[0:3]
  modern := languages[3:7]
  new := languages[7:9]

  fmt.Println("Languages classics:", classics)
  fmt.Println("Languages modern:", modern)
  fmt.Println("Languages new:", new) 
  */
  // MAPS
  /*
  var map1 = map[string]int{ "Pavan" : 5 , "Avyukth": 7 , "Sahaha" : 2}
  fmt.Println("map1", map1["Pavan"])
  delete(map1, "Sahaha") 
  var age ,ok = map1["Fuck"]
  if ok {
    fmt.Print(" the value is ok ")
  }else {
    fmt.Println("the age is not there: ", age)
  }

  for name, age := range map1{
    fmt.Println("name: ", name, "\t","age : ", age)
  } */

//STRINGS

/* var myString = "resume"
  fmt.Print(myString)

  var mystring = []string { "hello", "world" , "fuck" ,"you"}
  var catStr = ""
  for i:= range mystring {
    catStr += mystring[i]
  }
  fmt.Println("catStr", catStr)
   */
  //function call
  
  // printMe("ehllo")
  
  // var result , rem, err = intDiv(10,2)
  
 // control if else statement 
  /* if err!=nil{
    fmt.Printf(err.Error())
  }else if rem ==0 {
    fmt.Println("the result is ok")
  }else {
  fmt.Println("the fcking answeer is :",result, rem) 
    } 
    */ 
  /*
  switch rem {
    case 0: 
    fmt.Println(" the rem is zero")
    case 1,2:
    fmt.Println("the rem is 1 or 2")
    case 3: 
    fmt.Println("the rem is 3")
  default: 
    fmt.Println("it is not true ", result,err)
  }
  */

  
// main func ends here }


//FUNCTIONs
/*
func printMe(printvalue string){
  fmt.Println(printvalue)
}

func intDiv(num int , denum int ) (int,int,error) {
 var result int = num%denum
  var err error 
  if denum == 0 {
    err = errors.New("cannot divide it by zero")
    return 0 , 0 , err
  }


  return num/denum ,result ,err

} */
