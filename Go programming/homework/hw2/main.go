package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
type dollars float64
func (d dollars) String() string{
	return fmt.Sprintf("$%.2f",d)
}
type database map[string]dollars
func (db database) list(w http.ResponseWriter, req *http.Request){
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
} 
func (db database) price(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return 
	}
	fmt.Fprintf(w, "%s\n", price)
}
func (db database) change(item string, price dollars){
	db[item] = price
}
func (db database) create(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", p)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])
}
func (db database) fetch(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound)
		return 
	}
	fmt.Fprintf(w, "item %s has price %s\n", item, db[item])
}
func (db database) drop(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound)
		return 
	}
	delete(db, item)
	fmt.Fprintf(w, "item %s has deleted \n", item)
}
func (db database) update(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("Not found item: %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	} 
	convert_price, err := strconv.ParseFloat(price, 64)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", convert_price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	db[item]=dollars(convert_price)
	fmt.Fprintf(w, "updated %s:%s\n", item, db[item])
}

func main(){
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/fetch", db.fetch)
	http.HandleFunc("/delete", db.drop)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
		
}
