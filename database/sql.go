/*
Go’da **`sql` paketi** diye ayrı bir paket **yoktur**. Veritabanı ile ilgili asıl kullanılan standart paket:

👉 **`database/sql`**

Ama `sql` adıyla gördüğün şey genelde şu anlamlara gelir:

1. `import "database/sql"` yaptığında bu paketi **`sql`** adıyla kullanırız.
   Yani `sql.DB`, `sql.Row`, `sql.Rows` gibi tipler aslında `database/sql` içindedir.
2. Bazı durumlarda `sql` kelimesi, Go’nun içinde hazır gelen **nullable tipleri** (`sql.NullString`, `sql.NullInt64` gibi) ifade eder.

Dolayısıyla senin sorduğun “sql paketi” aslında **`database/sql` paketidir.**
Ben şimdi sana bu paketin **tamamını, tiplerini, fonksiyonlarını, örnekleriyle birlikte** anlatacağım.

---

# 📌 `database/sql` Paketinin Amacı

* Farklı veritabanlarına (MySQL, PostgreSQL, SQLite, MSSQL) **ortak bir arayüz (interface)** sunar.
* Doğrudan veritabanı motorunu içermez, sürücü (driver) gerekir.
* Örneğin:

  * MySQL için: `github.com/go-sql-driver/mysql`
  * PostgreSQL için: `github.com/lib/pq`
  * SQLite için: `modernc.org/sqlite`

---

# 📌 Temel Tipler (Types)

| Tip                                                                  | Açıklama                                                      |
| -------------------------------------------------------------------- | ------------------------------------------------------------- |
| **`DB`**                                                             | Veritabanı bağlantısı ve connection pool’u temsil eder.       |
| **`Tx`**                                                             | Transaction (işlem) yönetimi sağlar.                          |
| **`Rows`**                                                           | Çoklu satır sonucu (`SELECT`) temsil eder.                    |
| **`Row`**                                                            | Tek satır sonucu temsil eder (`QueryRow`).                    |
| **`Stmt`**                                                           | Hazırlanmış SQL ifadelerini (prepared statement) temsil eder. |
| **`Result`**                                                         | `Exec` sonucunu tutar (`LastInsertId`, `RowsAffected`).       |
| **`NullString`, `NullInt64`, `NullFloat64`, `NullBool`, `NullTime`** | NULL değerleri Go tipleriyle uyumlu hale getirir.             |

---

# 📌 DB (Veritabanı Bağlantısı)
*/
``go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Kullanıcı:Şifre@tcp(HOST:PORT)/Veritabanı
    db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Bağlantıyı test et
    if err := db.Ping(); err != nil {
        panic(err)
    }
}
``
/*
### DB Fonksiyonları

* `db.Ping()` → Bağlantıyı test eder
* `db.Query()` → Çoklu satır sorgular
* `db.QueryRow()` → Tek satır sorgular
* `db.Exec()` → INSERT, UPDATE, DELETE çalıştırır
* `db.Prepare()` → Hazırlanmış sorgu döner (`Stmt`)
* `db.Begin()` → Transaction başlatır
* `db.Close()` → Bağlantıyı kapatır

---

# 📌 Veri Ekleme (INSERT)
/*
``go
stmt, _ := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
res, _ := stmt.Exec("Ahmet", 25)

id, _ := res.LastInsertId()
fmt.Println("Son eklenen ID:", id)
``
/*
---

# 📌 Veri Okuma (SELECT)

### Çoklu Satır (`Rows`)
*/
``go
rows, _ := db.Query("SELECT id, name, age FROM users")
defer rows.Close()

for rows.Next() {
    var id, age int
    var name string
    rows.Scan(&id, &name, &age)
    fmt.Println(id, name, age)
}
``
/*
### Tek Satır (`Row`)
*/
``go
var name string
var age int
err := db.QueryRow("SELECT name, age FROM users WHERE id=?", 1).Scan(&name, &age)

if err == sql.ErrNoRows {
    fmt.Println("Kayıt bulunamadı")
}
``
/*
---

# 📌 Güncelleme (UPDATE)
*/
``go
res, _ := db.Exec("UPDATE users SET age=? WHERE id=?", 30, 1)
count, _ := res.RowsAffected()
fmt.Println("Güncellenen satır sayısı:", count)
```
/*
---

# 📌 Silme (DELETE)
*/
``go
res, _ := db.Exec("DELETE FROM users WHERE id=?", 2)
count, _ := res.RowsAffected()
fmt.Println("Silinen satır sayısı:", count)
``
/*
---

# 📌 Transaction Kullanımı
*/
``go
tx, _ := db.Begin()

_, err := tx.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Mehmet", 40)
if err != nil {
    tx.Rollback()
    return
}

_, err = tx.Exec("UPDATE users SET age=? WHERE name=?", 41, "Mehmet")
if err != nil {
    tx.Rollback()
    return
}

tx.Commit()
``
/*
---

# 📌 NULL Değerler (`sql.Null*`)
*/
``go
var email sql.NullString
err := db.QueryRow("SELECT email FROM users WHERE id=?", 1).Scan(&email)
if err != nil {
    panic(err)
}

if email.Valid {
    fmt.Println("Email:", email.String)
} else {
    fmt.Println("Email NULL")
}
``
/*
---

# 📌 Özel Hatalar

| Hata            | Açıklama                                                |
| --------------- | ------------------------------------------------------- |
| `sql.ErrNoRows` | Sorgu sonucunda satır bulunamadı                        |
| `sql.ErrTxDone` | Transaction commit/rollback sonrası tekrar kullanılırsa |

---

# 📌 Özet

* **`database/sql`** Go’nun standart veritabanı paketi
* Kendisi motor içermez → sürücü gerekir
* Temel fonksiyonlar: `Exec`, `Query`, `QueryRow`, `Prepare`, `Begin`, `Commit`, `Rollback`
* Null değerler için `sql.Null*` tipleri
* Transaction desteği mevcut
* Connection pool ayarları (`SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`) yapılabilir

---
*/

✅ Böylece `sql` paketinin (aslında `database/sql`) **tam kapsamlı bir açıklamasını** ve örneklerini vermiş olduk.
/*
İstersen ben sana bunun üstüne **`database/sql` ile farklı veritabanı sürücülerini (MySQL, PostgreSQL, SQLite) karşılaştırmalı olarak** gösterebilirim. İstiyor musun?
EEVET

Harika 👍 o zaman sana **tek bir proje** yazayım:

Bu uygulama:

* `database/sql` kullanacak
* MySQL, PostgreSQL, SQLite için **driver seçilebilir** olacak
* Tek bir `users` tablosu üzerinde **CRUD işlemleri** yapacak
* Config’ten hangi DB kullanılacaksa ona göre bağlanacak

---

# 📌 Ortak CRUD Uygulaması
*/
``
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"    // MySQL
	_ "github.com/lib/pq"                 // PostgreSQL
	_ "modernc.org/sqlite"                // SQLite
)

var db *sql.DB

// Kullanıcı modeli
type User struct {
	ID   int
	Name string
	Age  int
}

// Veritabanına bağlan
func connectDB(driver, dsn string) {
	var err error
	db, err = sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("Bağlantı hatası:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Ping hatası:", err)
	}
	fmt.Println(driver, "bağlantısı başarılı!")
}

// Tabloyu oluştur
func createTable() {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100),
		age INT
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Tablo oluşturulamadı:", err)
	}
}

// Kullanıcı ekle
func insertUser(name string, age int) {
	// PostgreSQL için $1,$2 - MySQL/SQLite için ?
	var query string
	if os.Getenv("DB_DRIVER") == "postgres" {
		query = "INSERT INTO users(name, age) VALUES($1, $2)"
	} else {
		query = "INSERT INTO users(name, age) VALUES(?, ?)"
	}
	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Println("Ekleme hatası:", err)
	} else {
		fmt.Println("Kullanıcı eklendi:", name)
	}
}

// Tüm kullanıcıları listele
func listUsers() {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("Sorgu hatası:", err)
	}
	defer rows.Close()

	fmt.Println("Kullanıcılar:")
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Age)
		fmt.Printf("%d - %s (%d)\n", u.ID, u.Name, u.Age)
	}
}

// Kullanıcı güncelle
func updateUser(id, age int) {
	var query string
	if os.Getenv("DB_DRIVER") == "postgres" {
		query = "UPDATE users SET age=$1 WHERE id=$2"
	} else {
		query = "UPDATE users SET age=? WHERE id=?"
	}
	_, err := db.Exec(query, age, id)
	if err != nil {
		log.Println("Güncelleme hatası:", err)
	} else {
		fmt.Println("Kullanıcı güncellendi, ID:", id)
	}
}

// Kullanıcı sil
func deleteUser(id int) {
	var query string
	if os.Getenv("DB_DRIVER") == "postgres" {
		query = "DELETE FROM users WHERE id=$1"
	} else {
		query = "DELETE FROM users WHERE id=?"
	}
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println("Silme hatası:", err)
	} else {
		fmt.Println("Kullanıcı silindi, ID:", id)
	}
}

func main() {
	// Ortam değişkeninden driver seç (mysql, postgres, sqlite)
	driver := os.Getenv("DB_DRIVER")
	var dsn string

	switch driver {
	case "mysql":
		dsn = "root:1234@tcp(127.0.0.1:3306)/testdb"
	case "postgres":
		dsn = "user=postgres password=1234 host=localhost port=5432 dbname=testdb sslmode=disable"
	case "sqlite":
		dsn = "test.db"
	default:
		log.Fatal("Lütfen DB_DRIVER ortam değişkenini ayarla: mysql | postgres | sqlite")
	}

	// Bağlantı
	connectDB(driver, dsn)
	defer db.Close()

	// Tablo
	createTable()

	// CRUD işlemleri
	insertUser("Ahmet", 25)
	insertUser("Mehmet", 30)

	listUsers()

	updateUser(1, 28)
	deleteUser(2)

	listUsers()
}
``
/*
---

# 📌 Çalıştırma
*/
//### MySQL için:

``sh
export DB_DRIVER=mysql
go run main.go
``

//### PostgreSQL için:

``sh
export DB_DRIVER=postgres
go run main.go
``

//### SQLite için:

``sh
export DB_DRIVER=sqlite
go run main.go
``
/*
---

# ✅ Özellikler

* Tek kod → üç veritabanı ile uyumlu
* CRUD işlemleri (`INSERT`, `SELECT`, `UPDATE`, `DELETE`)
* `os.Getenv("DB_DRIVER")` ile hangi sürücü kullanılacağı seçiliyor
* PostgreSQL `($1, $2)` parametre stilleri → ayrı kontrol

---

İstersen ben bu kodu biraz daha geliştirip **repository katmanı** ekleyeyim, böylece kodun içinde SQL sorguları kalmaz, daha temiz olur. Bunu ister misin?
*/