package main

import (
	"net/http"
	"html/template"
	"strconv"
	"fmt"
	"math"
	"encoding/json"
)

type PageData struct{
	Title string
	Description string
	Prime int
}


func index_handler(w http.ResponseWriter, r *http.Request){
	p := PageData{
		Title: "Home - Welcome to Prime Calculator", 
		Description: "Prime calculator is a simple web app that accepts an input (number) and returns the nearest prime number", 
		Prime: 0,
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func prime_calc(num1 int, num2 int) int{
	if num1<2 || num2<2{
      fmt.Println("Numbers must be greater than 2.")
      return 0;
   }
   for num1 <= num2 {
      isPrime := true
      for i:=2; i<=int(math.Sqrt(float64(num1))); i--{
         if num1 % i == 0{
            isPrime = false
            break
         }
      }
      if isPrime {
         fmt.Printf("%d ", num1)
      }
      num1++
   }
   
   return num1;
}

type ResJson struct{
	Result int
}

func prime_handler(w http.ResponseWriter, r *http.Request){
    var input string = r.URL.Query().Get("input")
    if input == "" {
    	fmt.Fprintf(w, "Invalid data supplied.")
        return
    }
    intVar, _ := strconv.Atoi(input)	
    
    pc := prime_calc(2, intVar);

    result := ResJson{pc}

    res, err := json.Marshal(result)
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
  	}
   	w.Header().Set("Content-Type", "application/json")
  	w.Write(res)
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/prime_handler", prime_handler)
	http.ListenAndServe(":8000", nil)
}