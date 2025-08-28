/*
Golang’da aslında **“database” diye tek başına bir paket yoktur**. Go’nun standart kütüphanesinde veritabanı ile çalışmak için kullanılan asıl paket:

👉 **`database/sql`** paketidir.

Bu paket, **veritabanlarına erişim için genel bir arayüz** (interface) sunar. Yani MySQL, PostgreSQL, SQLite, MSSQL gibi farklı veritabanlarına aynı API üzerinden erişebilmeni sağlar.
Ama `database/sql` doğrudan bir veritabanı motorunu içermez, onun yerine **sürücü (driver)** kullanır.

Örneğin:

* MySQL için: `github.com/go-sql-driver/mysql`
* PostgreSQL için: `github.com/lib/pq` veya `github.com/jackc/pgx`
* SQLite için: `modernc.org/sqlite`

---

## 📌 `database/sql` Paketinin Temel Kavramları

1. **`sql.DB`**

   * Bir veritabanına bağlantıyı temsil eder.
   * Gerçekte tek bir bağlantı değil, bağlantı havuzu (connection pool) yönetir.
   * `sql.Open(driverName, dataSourceName)` ile oluşturulur.

2. **`sql.Stmt` (Statement)**

   * Tekrar tekrar çalıştırılabilen, önceden hazırlanmış SQL ifadesidir (prepared statement).

3. **`sql.Rows`**

   * `SELECT` sorgularının döndürdüğü satırlar üzerinde gezinmek için kullanılır.

4. **`sql.Row`**

   * Tek bir satır döndüren sorgularda kullanılır (`QueryRow`).

5. **`sql.Tx` (Transaction)**

   * İşlemleri bir bütün olarak yönetmek için kullanılır (BEGIN, COMMIT, ROLLBACK).

6. **Hatalar**

   * Paket, `ErrNoRows` gibi bazı özel hatalar döndürebilir.

---

## 📌 Temel Fonksiyonlar ve Kullanımları

### 1. Veritabanına Bağlanma
*/
``
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL sürücüsü
)

func main() {
	// Kullanıcı:şifre@tcp(host:port)/veritabani
	dsn := "root:1234@tcp(127.0.0.1:3306)/testdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Bağlantıyı test et
	if err := db.Ping(); err != nil {
		log.Fatal("Bağlantı hatası:", err)
	}

	fmt.Println("Veritabanına bağlanıldı!")
}
``
/*
---

### 2. Veri Ekleme (`INSERT`)
*/
``
stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
if err != nil {
	log.Fatal(err)
}
defer stmt.Close()

res, err := stmt.Exec("Ahmet", 25)
if err != nil {
	log.Fatal(err)
}

lastID, _ := res.LastInsertId()
rowsAffected, _ := res.RowsAffected()

fmt.Println("Son eklenen ID:", lastID)
fmt.Println("Etkilenen satır sayısı:", rowsAffected)
``
/*
---

### 3. Veri Sorgulama (`SELECT`)
*/
``
rows, err := db.Query("SELECT id, name, age FROM users")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
	var id int
	var name string
	var age int
	err := rows.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %d, İsim: %s, Yaş: %d\n", id, name, age)
}
``

//👉 Tek satır sorgusu (`QueryRow`):

``
var name string
var age int
err := db.QueryRow("SELECT name, age FROM users WHERE id = ?", 1).Scan(&name, &age)
if err == sql.ErrNoRows {
	fmt.Println("Kayıt bulunamadı")
} else if err != nil {
	log.Fatal(err)
}
fmt.Println("Kullanıcı:", name, "Yaş:", age)
``
/*
---

### 4. Veri Güncelleme (`UPDATE`)
*/
``
res, err := db.Exec("UPDATE users SET age=? WHERE id=?", 30, 1)
if err != nil {
	log.Fatal(err)
}
count, _ := res.RowsAffected()
fmt.Println("Güncellenen kayıt sayısı:", count)
``
/*
---

### 5. Veri Silme (`DELETE`)
*/
``
res, err := db.Exec("DELETE FROM users WHERE id=?", 2)
if err != nil {
	log.Fatal(err)
}
count, _ := res.RowsAffected()
fmt.Println("Silinen kayıt sayısı:", count)
``
/*
---

### 6. Transaction Kullanımı
*/
``
tx, err := db.Begin()
if err != nil {
	log.Fatal(err)
}

_, err = tx.Exec("INSERT INTO users(name, age) VALUES(?, ?)", "Mehmet", 40)
if err != nil {
	tx.Rollback()
	log.Fatal(err)
}

_, err = tx.Exec("UPDATE users SET age=? WHERE name=?", 41, "Mehmet")
if err != nil {
	tx.Rollback()
	log.Fatal(err)
}

err = tx.Commit()
if err != nil {
	log.Fatal(err)
}

fmt.Println("Transaction başarılı!")
``
/*
---

## 📌 Özet

* `database/sql`, Go’nun **veritabanı ile çalışma için standart paketi**dir.
* Kendisi veritabanı motoru içermez, **driver** gerekir.
* `sql.DB`, `sql.Row`, `sql.Rows`, `sql.Tx` en önemli yapı taşlarıdır.
* Temel fonksiyonlar:

  * `Open`, `Ping` → Bağlantı
  * `Query`, `QueryRow` → Sorgu
  * `Exec` → Insert/Update/Delete
  * `Begin`, `Commit`, `Rollback` → Transaction

---

istersen ben sana `database/sql` paketinin **tüm fonksiyonlarını ve tiplerini tablo halinde** çıkarabilirim, ne dersin?
//Evet
*/

Harika 👍 Şimdi sana **`database/sql` paketinin tüm önemli tiplerini, fonksiyonlarını ve sabitlerini** tablo halinde çıkarıyorum. Böylece paketin tamamını kolayca görebilirsin.

---

# 📌 `database/sql` Paketinin Yapısı

## 1. Temel Tipler (Types)

| Tip                                                                  | Açıklama                                                                                  |
| -------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| **`DB`**                                                             | Veritabanı bağlantı havuzunu (connection pool) temsil eder. `sql.Open()` ile oluşturulur. |
| **`Tx`**                                                             | Transaction (işlem) yönetimini sağlar. `db.Begin()` ile başlar.                           |
| **`Rows`**                                                           | `SELECT` sorgularının döndürdüğü çoklu satırları temsil eder.                             |
| **`Row`**                                                            | Tek bir satırı temsil eder (`QueryRow`).                                                  |
| **`Stmt`**                                                           | Hazırlanmış SQL ifadelerini (prepared statement) temsil eder.                             |
| **`Result`**                                                         | `Exec` sonucu dönen etkilenen satır sayısı ve son eklenen ID bilgilerini tutar.           |
| **`NullString`, `NullInt64`, `NullBool`, `NullFloat64`, `NullTime`** | NULL değerler ile çalışmak için özel tiplerdir.                                           |

---

## 2. DB (Veritabanı Nesnesi) Fonksiyonları

| Fonksiyon                                  | Açıklama                                             |
| ------------------------------------------ | ---------------------------------------------------- |
| **`sql.Open(driverName, dataSourceName)`** | Yeni bir veritabanı bağlantısı açar.                 |
| **`db.Ping()` / `db.PingContext(ctx)`**    | Bağlantıyı test eder.                                |
| **`db.Close()`**                           | Bağlantıyı kapatır.                                  |
| **`db.Exec(query, args...)`**              | SQL komutlarını çalıştırır (INSERT, UPDATE, DELETE). |
| **`db.Query(query, args...)`**             | Sorgu çalıştırır, `Rows` döner.                      |
| **`db.QueryRow(query, args...)`**          | Tek satır dönen sorgular için kullanılır.            |
| **`db.Begin()` / `db.BeginTx(ctx, opts)`** | Transaction başlatır.                                |
| **`db.Prepare(query)`**                    | Hazırlanmış SQL ifadesi döner (`Stmt`).              |
| **`db.SetMaxOpenConns(n)`**                | Maksimum açık bağlantı sayısını belirler.            |
| **`db.SetMaxIdleConns(n)`**                | Kullanılmayan bağlantıların sayısını belirler.       |
| **`db.SetConnMaxLifetime(d)`**             | Bağlantının maksimum ömrünü belirler.                |

---

## 3. Rows (Çoklu Satır Sorguları)

| Fonksiyon                 | Açıklama                                         |
| ------------------------- | ------------------------------------------------ |
| **`rows.Next()`**         | Sonraki satıra geçer.                            |
| **`rows.Scan(&dest...)`** | Mevcut satırdaki değerleri değişkenlere aktarır. |
| **`rows.Err()`**          | Okuma sırasında hata olup olmadığını döner.      |
| **`rows.Close()`**        | `rows` nesnesini kapatır.                        |

---

## 4. Row (Tek Satır Sorguları)

| Fonksiyon                | Açıklama                                                        |
| ------------------------ | --------------------------------------------------------------- |
| **`row.Scan(&dest...)`** | Tek satırdan veri alır. Eğer satır yoksa `sql.ErrNoRows` döner. |

---

## 5. Stmt (Prepared Statement)

| Fonksiyon                    | Açıklama                        |
| ---------------------------- | ------------------------------- |
| **`stmt.Exec(args...)`**     | Hazırlanmış sorguyu çalıştırır. |
| **`stmt.Query(args...)`**    | Çoklu satır döner.              |
| **`stmt.QueryRow(args...)`** | Tek satır döner.                |
| **`stmt.Close()`**           | Statement’i kapatır.            |

---

## 6. Tx (Transaction)

| Fonksiyon                         | Açıklama                                        |
| --------------------------------- | ----------------------------------------------- |
| **`tx.Exec(query, args...)`**     | Transaction içinde SQL çalıştırır.              |
| **`tx.Query(query, args...)`**    | Transaction içinde sorgu yapar.                 |
| **`tx.QueryRow(query, args...)`** | Transaction içinde tek satır sorgular.          |
| **`tx.Prepare(query)`**           | Transaction içinde prepared statement hazırlar. |
| **`tx.Commit()`**                 | Transaction’ı tamamlar.                         |
| **`tx.Rollback()`**               | Transaction’ı geri alır.                        |

---

## 7. Result (Exec Sonuçları)

| Fonksiyon                | Açıklama                           |
| ------------------------ | ---------------------------------- |
| **`res.LastInsertId()`** | Son eklenen satırın ID’sini döner. |
| **`res.RowsAffected()`** | Etkilenen satır sayısını döner.    |

---

## 8. Hata (Errors)

| Hata                | Açıklama                                            |
| ------------------- | --------------------------------------------------- |
| **`sql.ErrNoRows`** | `QueryRow.Scan()` sırasında satır bulunmazsa döner. |
| **`sql.ErrTxDone`** | Transaction zaten commit/rollback edilmişse döner.  |

---

## 9. Null Tipleri

`database/sql` NULL değerleri doğrudan `string`, `int`, `bool` gibi tiplere aktaramaz. Bunun için **nullable tipler** vardır:
*/
``
var name sql.NullString
var age sql.NullInt64

err := db.QueryRow("SELECT name, age FROM users WHERE id=?", 1).Scan(&name, &age)
if err != nil {
	log.Fatal(err)
}

if name.Valid {
	fmt.Println("İsim:", name.String)
} else {
	fmt.Println("İsim NULL")
}
``

---

✅ Böylece `database/sql` paketinin **tam fonksiyonlarını, tiplerini ve kullanım alanlarını** sana özetlemiş oldum.

İstersen ben sana şimdi bunun üstüne **tüm fonksiyonları kapsayan bir mini CRUD uygulaması (Create-Read-Update-Delete)** örneği yazabilirim. İster MySQL, ister PostgreSQL üzerinde çalışacak şekilde gösterebilirim. Hangisini tercih edersin?



