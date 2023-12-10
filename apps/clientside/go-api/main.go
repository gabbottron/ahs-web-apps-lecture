package main

import (
	"database/sql"
	"encoding/json"

	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // must be included this way for postgres functionality on standard driver
)

// Item represents an item in the database
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

var db *sql.DB

func initDB() {
	var err error
	// Replace with your PostgreSQL connection details
	connStr := "user=postgres dbname=webapp password=password host=postgres port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/health", healthCheck).Methods("GET")
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{item_id}", getItem).Methods("GET")

	// Enable CORS for all routes
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}), // Replace with your allowed origins
	)

	// Use the CORS middleware with your router
	http.Handle("/", corsHandler(router))

	log.Fatal(http.ListenAndServe(":8080", nil))
	// Without CORS, use below
	//log.Fatal(http.ListenAndServe(":8080", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func getItems(w http.ResponseWriter, r *http.Request) {
	log.Println("-->getItems")
	items := []Item{}
	rows, err := db.Query("SELECT id, name, quantity FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Quantity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	log.Println("-->getItem")
	params := mux.Vars(r)
	itemID, err := strconv.Atoi(params["item_id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item Item
	err = db.QueryRow("SELECT id, name, quantity, description FROM items WHERE id = $1", itemID).Scan(&item.ID, &item.Name, &item.Quantity, &item.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Item not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}
