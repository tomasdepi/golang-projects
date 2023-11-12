package main

// type Product struct {
// 	Id       int
// 	Name     string
// 	Quantity int
// 	Price    float64
// }

// var Products []Product

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to homepage")
// 	log.Println("Endpoint hit: homepage")
// }

// func returnAllProducts(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Endpoint hit: returnAllProducts")
// 	json.NewEncoder(w).Encode(Products)
// }

// func getProduct(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Endpoint hit: getProduct, ", r.URL.Path)
// 	vars := mux.Vars(r)
// 	productId, _ := strconv.Atoi(vars["id"])

// 	json.NewEncoder(w).Encode(Products[productId-1])
// }

// func handleRequests() {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/product/{id}", getProduct)
// 	myRouter.HandleFunc("/products", returnAllProducts)
// 	myRouter.HandleFunc("/", homePage)
// 	http.ListenAndServe("127.0.0.1:8080", myRouter)
// }

func main() {

	// Products = []Product{
	// 	{Id: 1, Name: "Chair", Quantity: 100, Price: 100.00},
	// 	{Id: 2, Name: "Desk", Quantity: 50, Price: 200.00},
	// }

	// handleRequests()
	app := App{}
	app.Initialize()
	app.Run("localhost:8080")
}
