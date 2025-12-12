// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"time"
// 	// IMPORT YOUR LOCAL PACKAGE
// 	// Format: <module-name>/<path-to-package>
// 	"my-monitor/internal/checker"
// 	"my-monitor/internal/storage"
// )
// type Server struct{
// 	db *storage.DB
// }
// func (s *Server) handleCheck(w http.ResponseWriter, r *http.Request){
// 	var input struct {
// 		URL string `json:"url"`
// 	}
// 	//Decode Json 
// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
// 	defer cancel()
// 	status := checker.CheckURL(ctx, input.URL)
// 	if err := s.db.Save(input.URL, status); err != nil{
// 		log.Println("Database error:", err)
// 		http.Error(w, "Failed to save result", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"url": input.URL,
// 		"status": status,
// 	})

// }
// func (s *Server) handleList(w http.ResponseWriter, r *http.Request){
// 	records, err := s.db.GetAll()
// 	if err != nil {
// 		http.Error(w, "Database Error", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(records)
// }
// func main() {
// 	db, err := storage.NewSQLite("./data/monitor.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	server := &Server{db: db}
// 	// Setup Router
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("POST /check", server.handleCheck)
// 	mux.HandleFunc("GET /history", server.handleList)
// 	log.Println("Server starting on :8880")
// 	if err := http.ListenAndServe(":8880", mux); err != nil {
// 		log.Fatal(err)
// 	}
// }
