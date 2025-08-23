/*
İstersen ben sana şimdi bunun üstüne tüm fonksiyonları kapsayan bir mini CRUD uygulaması (Create-Read-Update-Delete) örneği yazabilirim. 
İster MySQL, ister PostgreSQL üzerinde çalışacak şekilde gösterebilirim. Hangisini tercih edersin?

Süper 👍 Şimdi sana **MySQL üzerinde çalışan, `database/sql` paketinin neredeyse tüm fonksiyonlarını kapsayan bir mini CRUD uygulaması** yazacağım. Bu uygulama:

* Veritabanına bağlanacak
* `users` tablosu oluşturacak
* **C**reate → kullanıcı ekleyecek
* **R**ead → kullanıcıları listeleyecek ve tek kullanıcıyı çekecek
* **U**pdate → kullanıcı güncelleyecek
* **D**elete → kullanıcı silecek
* Transaction ve prepared statement kullanacak
* NULL değerler ile çalışmayı gösterecek

---

## 📌 MySQL `users` Tablosu

Öncelikle şu tabloyu oluşturmanı öneririm:
*/
``sql
CREATE DATABASE testdb;
USE testdb;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    age INT,
    email VARCHAR(100) NULL
);
``
/*
---

## 📌 Go CRUD Uygulaması (MySQL + database/sql)
*/
``
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type User struct {
	ID    int
	Name  string
	Age   int
	Email sql.NullString // NULL destekli tip
}

func main() {
	// MySQL bağlantısı
	dsn := "root:1234@tcp(127.0.0.1:3306)/testdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Bağlantı açılırken hata:", err)
	}
	defer db.Close()

	// Bağlantıyı test et
	if err := db.Ping(); err != nil {
		log.Fatal("Ping hatası:", err)
	}
	fmt.Println("✅ Veritabanına bağlanıldı")

	// Bağlantı havuzu ayarları
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	// ---- CREATE ----
	userID := createUser(db, "Ahmet", 25, sql.NullString{String: "ahmet@example.com", Valid: true})
	fmt.Println("Yeni kullanıcı eklendi, ID:", userID)

	// ---- READ ALL ----
	fmt.Println("\n--- Tüm Kullanıcılar ---")
	users := getAllUsers(db)
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

	// ---- READ ONE ----
	fmt.Println("\n--- Tek Kullanıcı ---")
	user := getUserByID(db, userID)
	fmt.Printf("%+v\n", user)

	// ---- UPDATE ----
	fmt.Println("\n--- Kullanıcı Güncelle ---")
	updateUser(db, userID, "Ahmet Yılmaz", 30)
	fmt.Println("Kullanıcı güncellendi")

	// ---- DELETE ----
	fmt.Println("\n--- Kullanıcı Sil ---")
	deleteUser(db, userID)
	fmt.Println("Kullanıcı silindi")

	// ---- TRANSACTION ÖRNEĞİ ----
	fmt.Println("\n--- Transaction ---")
	transactionExample(db)
}

// CREATE
func createUser(db *sql.DB, name string, age int, email sql.NullString) int64 {
	stmt, err := db.Prepare("INSERT INTO users(name, age, email) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, age, email)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := res.LastInsertId()
	return id
}

// READ ALL
func getAllUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT id, name, age, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

// READ ONE
func getUserByID(db *sql.DB, id int64) User {
	var u User
	err := db.QueryRow("SELECT id, name, age, email FROM users WHERE id = ?", id).
		Scan(&u.ID, &u.Name, &u.Age, &u.Email)

	if err == sql.ErrNoRows {
		fmt.Println("Kullanıcı bulunamadı")
	} else if err != nil {
		log.Fatal(err)
	}
	return u
}

// UPDATE
func updateUser(db *sql.DB, id int64, newName string, newAge int) {
	_, err := db.Exec("UPDATE users SET name=?, age=? WHERE id=?", newName, newAge, id)
	if err != nil {
		log.Fatal(err)
	}
}

// DELETE
func deleteUser(db *sql.DB, id int64) {
	_, err := db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
}

// TRANSACTION ÖRNEĞİ
func transactionExample(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// 1. kullanıcı ekle
	_, err = tx.Exec("INSERT INTO users(name, age, email) VALUES(?, ?, ?)", "Mehmet", 40, nil)
	if err != nil {
		tx.Rollback()
		log.Fatal("Insert hatası:", err)
	}

	// 2. yaş güncelle
	_, err = tx.Exec("UPDATE users SET age=? WHERE name=?", 41, "Mehmet")
	if err != nil {
		tx.Rollback()
		log.Fatal("Update hatası:", err)
	}

	// Commit işlemi
	if err := tx.Commit(); err != nil {
		log.Fatal("Commit hatası:", err)
	}
	fmt.Println("Transaction başarıyla tamamlandı")
}
``
/*
---

## 📌 Bu Kodda Kullanılan `database/sql` Özellikleri

* `sql.Open`, `db.Ping` → bağlantı
* `Exec`, `Query`, `QueryRow` → sorgular
* `Stmt` → prepared statement
* `Rows`, `Row` → veri çekme
* `Result.LastInsertId`, `RowsAffected` → işlem sonuçları
* `NullString` → NULL değer desteği
* `Begin`, `Commit`, `Rollback` → transaction
* `SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime` → connection pool yönetimi

---

👉 Bu uygulama `database/sql` paketinin **tam kapsamlı bir CRUD örneği** oldu.

İstersen ben bunu sana daha da **katmanlı bir yapı (repository pattern + service layer)** ile yazabilirim. Böylece gerçek bir projeye daha yakın olur. Bunu ister misin?
*/
/*
Harika 👍 O zaman şimdi aynı **MySQL CRUD uygulamasını** bu kez **katmanlı mimari (repository + service + main)** ile yazalım.

Bu yaklaşım, gerçek dünyadaki projelerde **temiz kod** ve **sürdürülebilirlik** için tercih edilir.

---

# 📌 Mimari Katmanlar

1. **Repository Layer**

   * Veritabanı işlemlerini yapar (`Create`, `GetAll`, `GetByID`, `Update`, `Delete`)
   * Yalnızca `database/sql` ile uğraşır

2. **Service Layer**

   * İş mantığını barındırır
   * Repository’yi kullanarak işlemleri gerçekleştirir

3. **Main (Uygulama Katmanı)**

   * Kullanıcıdan giriş alır (veya sabit değerlerle test eder)
   * Service katmanını çağırır

---

# 📌 Proje Yapısı
*/
``
go-crud/
│── main.go
│── user.go          (entity/model)
│── user_repository.go
│── user_service.go
``
/*
---

## 📌 `user.go` (Entity / Model)
*/
``
package main

import "database/sql"

// User struct'ı veritabanındaki users tablosunu temsil eder
type User struct {
	ID    int
	Name  string
	Age   int
	Email sql.NullString
}
``
/*
---

## 📌 `user_repository.go` (Repository Layer)
*/
``
package main

import (
	"database/sql"
	"log"
)

// UserRepository veritabanı işlemlerini içerir
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository yeni repository döner
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CREATE
func (r *UserRepository) Create(user User) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO users(name, age, email) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Age, user.Email)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// READ ALL
func (r *UserRepository) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name, age, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users, nil
}

// READ ONE
func (r *UserRepository) GetByID(id int) (*User, error) {
	var u User
	err := r.db.QueryRow("SELECT id, name, age, email FROM users WHERE id=?", id).
		Scan(&u.ID, &u.Name, &u.Age, &u.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// UPDATE
func (r *UserRepository) Update(user User) error {
	_, err := r.db.Exec("UPDATE users SET name=?, age=?, email=? WHERE id=?",
		user.Name, user.Age, user.Email, user.ID)
	return err
}

// DELETE
func (r *UserRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
``
/*
---

## 📌 `user_service.go` (Service Layer)
*/
``
package main

import "fmt"

// UserService iş mantığını içerir
type UserService struct {
	repo *UserRepository
}

// NewUserService yeni service döner
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Kullanıcı ekleme
func (s *UserService) CreateUser(name string, age int, email string) {
	u := User{Name: name, Age: age}
	if email != "" {
		u.Email.String = email
		u.Email.Valid = true
	}
	id, err := s.repo.Create(u)
	if err != nil {
		fmt.Println("Kullanıcı eklenemedi:", err)
		return
	}
	fmt.Println("Yeni kullanıcı eklendi. ID:", id)
}

// Tüm kullanıcıları listeleme
func (s *UserService) ListUsers() {
	users, err := s.repo.GetAll()
	if err != nil {
		fmt.Println("Kullanıcılar alınamadı:", err)
		return
	}
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

// Tek kullanıcı
func (s *UserService) GetUser(id int) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	if user == nil {
		fmt.Println("Kullanıcı bulunamadı")
		return
	}
	fmt.Printf("Bulunan kullanıcı: %+v\n", *user)
}

// Kullanıcı güncelleme
func (s *UserService) UpdateUser(id int, name string, age int, email string) {
	u := User{ID: id, Name: name, Age: age}
	if email != "" {
		u.Email.String = email
		u.Email.Valid = true
	}
	err := s.repo.Update(u)
	if err != nil {
		fmt.Println("Güncelleme hatası:", err)
		return
	}
	fmt.Println("Kullanıcı güncellendi.")
}

// Kullanıcı silme
func (s *UserService) DeleteUser(id int) {
	err := s.repo.Delete(id)
	if err != nil {
		fmt.Println("Silme hatası:", err)
		return
	}
	fmt.Println("Kullanıcı silindi.")
}
``
/*
---

## 📌 `main.go` (Uygulama Katmanı)
*/
``
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL bağlantısı
	dsn := "root:1234@tcp(127.0.0.1:3306)/testdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Bağlantı hatası:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Ping hatası:", err)
	}

	// Bağlantı havuzu ayarları
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	fmt.Println("✅ MySQL'e bağlanıldı.")

	// Katmanları bağla
	repo := NewUserRepository(db)
	service := NewUserService(repo)

	// CRUD İşlemleri
	service.CreateUser("Ahmet", 25, "ahmet@example.com")
	service.CreateUser("Mehmet", 30, "")

	fmt.Println("\n--- Kullanıcı Listesi ---")
	service.ListUsers()

	fmt.Println("\n--- Tek Kullanıcı ---")
	service.GetUser(1)

	fmt.Println("\n--- Kullanıcı Güncelle ---")
	service.UpdateUser(1, "Ahmet Yılmaz", 28, "ahmetyilmaz@example.com")

	fmt.Println("\n--- Kullanıcı Sil ---")
	service.DeleteUser(2)

	fmt.Println("\n--- Son Liste ---")
	service.ListUsers()
}
``
/*
---

## 📌 Bu Yapıda Kazandıklarımız

* **Repository** → sadece SQL ile ilgilenir
* **Service** → iş mantığını barındırır (hata kontrolü, iş kuralları)
* **Main** → uygulamayı başlatır, servisleri çağırır
* Test yazmak kolaylaşır: Repository’yi mock’layıp servisleri test edebilirsin.

---
*/
/*
👉 Eğer istersen, bunun üstüne bir de **REST API (Gin veya Echo ile HTTP endpoint’leri)** yazabiliriz.
İstersen bu kodu **REST API** olarak geliştirmemi ister misin?EVET

Harika 👍 O zaman şimdi bu **katmanlı CRUD uygulamasını REST API** haline getirelim.

Framework olarak **Gin** kullanacağım çünkü çok popüler, hızlı ve öğrenmesi kolay.

---

# 📌 Proje Yapısı (REST API)
*/
``
go-crud-api/
│── main.go             (API server)
│── user.go             (model/entity)
│── user_repository.go  (repository)
│── user_service.go     (service)
│── user_handler.go     (HTTP handlers - REST API)
``
/*
---

## 📌 `user.go` (Model / Entity)
*/
``
package main

import "database/sql"

type User struct {
	ID    int            `json:"id"`
	Name  string         `json:"name"`
	Age   int            `json:"age"`
	Email sql.NullString `json:"email"`
}
``
/*
---

## 📌 `user_repository.go` (Repository)
*/
``
package main

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user User) (int64, error) {
	res, err := r.db.Exec("INSERT INTO users(name, age, email) VALUES(?, ?, ?)",
		user.Name, user.Age, user.Email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *UserRepository) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name, age, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*User, error) {
	var u User
	err := r.db.QueryRow("SELECT id, name, age, email FROM users WHERE id=?", id).
		Scan(&u.ID, &u.Name, &u.Age, &u.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Update(user User) error {
	_, err := r.db.Exec("UPDATE users SET name=?, age=?, email=? WHERE id=?",
		user.Name, user.Age, user.Email, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
``
/*
---

## 📌 `user_service.go` (Service Layer)
*/
``
package main

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(u User) (int64, error) {
	return s.repo.Create(u)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(u User) error {
	return s.repo.Update(u)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
``
/*
---

## 📌 `user_handler.go` (REST API – Handlers)
*/
``
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "kullanıcı eklenemedi"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GET /users
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "kullanıcılar alınamadı"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "kullanıcı alınamadı"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "kullanıcı bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// PUT /users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.ID = id

	if err := h.service.UpdateUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "güncelleme hatası"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "kullanıcı güncellendi"})
}

// DELETE /users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "silme hatası"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "kullanıcı silindi"})
}
``
/*
---

## 📌 `main.go` (API Server)
*/
``
package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL bağlantısı
	dsn := "root:1234@tcp(127.0.0.1:3306)/testdb?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Bağlantı hatası:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Ping hatası:", err)
	}

	// Connection pool ayarları
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Katmanları bağla
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	// Gin router
	r := gin.Default()

	// REST API endpointleri
	r.POST("/users", handler.CreateUser)
	r.GET("/users", handler.GetAllUsers)
	r.GET("/users/:id", handler.GetUser)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	// Server başlat
	log.Println("🚀 Server http://localhost:8080 üzerinde çalışıyor")
	r.Run(":8080")
}
``
/*
---

# 📌 REST API Endpointleri

* `POST   /users` → Yeni kullanıcı ekle
*/
  ``json
  {
    "name": "Ahmet",
    "age": 25,
    "email": "ahmet@example.com"
  }
  ``
/*
* `GET    /users` → Tüm kullanıcıları listele

* `GET    /users/:id` → Tek kullanıcı getir

* `PUT    /users/:id` → Kullanıcı güncelle

  ```json
  {
    "name": "Ahmet Yılmaz",
    "age": 30,
    "email": "ahmetyilmaz@example.com"
  }
  ```

* `DELETE /users/:id` → Kullanıcı sil

---
*/