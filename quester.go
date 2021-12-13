package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "sort"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type food struct {
    Name string `json:"name"`
    Qty  int    `json:"qty"`
}

func displayDetails(w http.ResponseWriter, r *http.Request) {
    vars:= mux.vars(r)
    miniQty,_ :=strconv.Atoi(vars["quantity"])
    flag :=false
    var sum []food
    fruits, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/fruits")
    vegetables, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/vegetables")
    grains, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/grains")
    response, _ := ioutil.ReadAll(fruits.Body)
    var fetch []food
    json.Unmarshal(response, &fetch)
    for i := 0; i < len(fetch); i++ {
             sum=append(sum,fetch[i])
    }
    response1, _ := ioutil.ReadAll(vegetables.Body)
    var fetch1 []food
    json.Unmarshal(response1, &fetch1)
    for i := 0; i < len(fetch1); i++ {
            sum=append(sum,fetch1[i])
    }
    response2, _ := ioutil.ReadAll(grains.Body)
    var fetch2 []food
    json.Unmarshal(response2, &fetch2)
    for i := 0; i <len(fetch2); i++ {
             sum=append(sum,fetch2[i])
    }
    sort.Slice(sum, func(i, j int) bool {
        return sum[i].Name < sum[j].Name
     })
     for _,item :=range sum{
         if item.Qty <= miniQty{
             flag = true
             fmt.Fprintln(w,item)
         }
     }
     if flag == false{
        fmt.Fprintln(w,"Not_Found")
     }
}
func handleRequests() {
    myRouter :=mux.NewRouter()
    myRouter.HandleFunc("/quest/{quantity}".displayDetails).Methods(http.MethodGet)
    log.Fatal(http.ListenAndServe(":8091", nil))
}
func main() {
    handleRequests()
   
}