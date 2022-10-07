package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/riteshgaigawali/go-microservices.git/details"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Checking the application health.....")
	response := map[string]string{
		"status":       "UP",
		"current time": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Serving homepage.....")
	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprintf(w, "Application is up and running....")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Fetching details.....")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token : %s\n", r.URL.Path, r.URL.Query().Get("token"))

// }
// func main() {
// 	http.HandleFunc("/", rootHandler)

// fs := http.FileServer(http.Dir("static/"))
// http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("Web server has started listning.......")
// 	http.ListenAndServe(":80", nil)
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"

// 	"github.com/riteshgaigawali/go-microservices.git/geometry" //importing geometry package
// 	"rsc.io/quote"
// )

// //Simple function declaration
// func rectProps(length, width float32) (float32, float32) {
// 	area := length * width
// 	perimeter := (length + width) * 2
// 	return area, perimeter
// }

// //Named return
// func rectProp(length, width float32) (area, perimeter float32) {
// 	area = length * width
// 	perimeter = (length + width) * 2
// 	return
// }

// func main() {

// 	//Importing packages
// 	fmt.Println(quote.Go())

// 	//Variables & Constants
// 	var x int32 = 10               //variable declaration with type
// 	var y, z, str = 2, 3, "Ritesh" //multiple declarations in one line without specifying type
// 	var name string                //tow line declaration
// 	name = "DevOps"                //assignment of the above declared variable "name"
// 	isWorking := false             //shorthand declaration
// 	const PI float64 = 3.14        //Constant declaration
// 	fmt.Printf("Hello, %s!\n", str)
// 	fmt.Println(x, y, z, name, isWorking, PI)
// 	fmt.Printf("The type of variable x is %T and size is %d.\n", x, unsafe.Sizeof(x)) //formated printing

// 	//Functions and Named return
// 	area, perimeter := rectProps(10, 2) //calling a funtion
// 	fmt.Printf("Area is %f and Perimeter is %f.\n", area, perimeter)
// 	area, perimeter = rectProp(12, 2) //calling another function *named return*
// 	fmt.Printf("Area is %f and Perimeter is %f.\n", area, perimeter)

// 	// //Maps (Dictionary like python)
// 	// var dayOfMonth map[string]int //tow line declaration
// 	// dayOfMonth["April"] = 30
// 	// fmt.Println(dayOfMonth)
// 	var daysOfMonth = map[string]int{"Jan": 31, "Feb": 28} //multiple values at once
// 	fmt.Println(daysOfMonth)
// 	daysOfMonths := map[string]int{"Mar": 31, "April": 30} //shorthand
// 	fmt.Println(daysOfMonths)

// 	//Importing local packages
// 	fmt.Print(geometry.Area(2, 3)) //call to method from external local package

// }
