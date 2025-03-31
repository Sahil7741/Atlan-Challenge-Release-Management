package main
import (
    "database/sql"
    "testing"
    _ "github.com/mattn/go-sqlite3"
)
type Repository struct { db *sql.DB }
func (r *Repository) GetUser(id int) (string, error) {
    var name string
    err := r.db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)
    return name, err
}
func TestGetUser(t *testing.T) {
    db, _ := sql.Open("sqlite3", ":memory:")
    defer db.Close()
    db.Exec("CREATE TABLE users (id INTEGER, name TEXT)")
    db.Exec("INSERT INTO users (id, name) VALUES (1, 'Ankit')")
    repo := &Repository{db: db}
    name, err := repo.GetUser(1)
    if err != nil || name != "Ankit" {
        t.Errorf("Expected 'Ankit', got '%s'", name)
    }
}