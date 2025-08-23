/*
Go’da `database/sql` paketini anlattık. Şimdi onun **alt seviyesindeki** pakete geldik:

---

# 📌 `database/sql/driver` Paketi

* `database/sql` yüksek seviyeli bir **database API** sağlar.
* Ama **veritabanı sürücüleri** (MySQL, Postgres, SQLite driver’ları vs.) aslında `database/sql/driver` paketine dayanır.
* Yani:

  * Uygulama → `database/sql` kullanır
  * Driver (örn: `go-sql-driver/mysql`) → `database/sql/driver` arayüzlerini implement eder

👉 Sen normalde `sql/driver` paketini **direkt kullanmazsın**.
Onu genellikle **kendi veritabanı sürücünü yazmak** istiyorsan kullanırsın.

---

# 📌 Temel Yapılar ve Arayüzler

`sql/driver` paketinde en önemli arayüzler şunlardır:

| Arayüz             | Açıklama                                                                                                                 |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| **Driver**         | Yeni bir veritabanı bağlantısı açar (`Open` fonksiyonu ile).                                                             |
| **Conn**           | Açılan veritabanı bağlantısını temsil eder (query/exec burada yapılır).                                                  |
| **Stmt**           | Hazırlanmış SQL ifadelerini (prepared statement) temsil eder.                                                            |
| **Tx**             | Transaction yönetimi için kullanılır.                                                                                    |
| **Rows**           | `SELECT` sonucunu temsil eder (satırları okur).                                                                          |
| **Result**         | `INSERT/UPDATE/DELETE` sonrası etkilenen satır sayısı ve son eklenen ID bilgisini taşır.                                 |
| **Value**          | Go tarafındaki bir veritabanı değerini temsil eder (`int64`, `float64`, `bool`, `[]byte`, `string`, `time.Time`, `nil`). |
| **ValueConverter** | Go → DB tipi dönüşümlerini yapar.                                                                                        |

---

# 📌 Örnek: Basit "MemoryDB" Driver Yazalım

Aşağıdaki örnekte kendi **hafıza tabanlı (in-memory)** veritabanı sürücümüzü yazıyoruz.
Bu driver `database/sql` tarafından kullanılabilecek.
*/
``go
package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
)

// ---------------- DRIVER ----------------

// MemoryDriver basit driver
type MemoryDriver struct{}

func (d *MemoryDriver) Open(name string) (driver.Conn, error) {
	return &MemoryConn{}, nil
}

// ---------------- CONNECTION ----------------

type MemoryConn struct{}

func (c *MemoryConn) Prepare(query string) (driver.Stmt, error) {
	return &MemoryStmt{query: query}, nil
}
func (c *MemoryConn) Close() error { return nil }
func (c *MemoryConn) Begin() (driver.Tx, error) {
	return &MemoryTx{}, nil
}

// ---------------- STATEMENT ----------------

type MemoryStmt struct {
	query string
}

func (s *MemoryStmt) Close() error { return nil }
func (s *MemoryStmt) NumInput() int {
	return -1 // parametre sayısı bilinmiyor
}
func (s *MemoryStmt) Exec(args []driver.Value) (driver.Result, error) {
	fmt.Println("Exec:", s.query, args)
	return driver.RowsAffected(1), nil
}
func (s *MemoryStmt) Query(args []driver.Value) (driver.Rows, error) {
	fmt.Println("Query:", s.query, args)
	return &MemoryRows{}, nil
}

// ---------------- TRANSACTION ----------------

type MemoryTx struct{}

func (t *MemoryTx) Commit() error   { return nil }
func (t *MemoryTx) Rollback() error { return nil }

// ---------------- ROWS ----------------

type MemoryRows struct{}

func (r *MemoryRows) Columns() []string {
	return []string{"id", "name"}
}
func (r *MemoryRows) Close() error { return nil }
func (r *MemoryRows) Next(dest []driver.Value) error {
	// Tek satır örnek
	dest[0] = int64(1)
	dest[1] = "Ahmet"
	return io.EOF // başka satır yok
}

// ---------------- MAIN ----------------

func main() {
	// Driver’ı register et
	sql.Register("memorydb", &MemoryDriver{})

	// Veritabanına bağlan
	db, err := sql.Open("memorydb", "")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Exec denemesi
	db.Exec("INSERT INTO users VALUES(?, ?)", 1, "Ahmet")

	// Query denemesi
	rows, _ := db.Query("SELECT * FROM users")
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println("Sonuç:", id, name)
	}
}
``
/*
---

# 📌 Örnekte Neler Oldu?

* `sql.Register("memorydb", &MemoryDriver{})` → kendi sürücümüzü `memorydb` adıyla kaydettik.
* `Open` → yeni bağlantı açıyor (`MemoryConn`).
* `Prepare` → SQL cümlesi hazırlıyor (`MemoryStmt`).
* `Exec` → `INSERT/UPDATE` gibi sorguları çalıştırıyor.
* `Query` → `SELECT` sonucu döndürüyor (`MemoryRows`).
* `Next` → satır satır veri sağlıyor.

---

# 📊 Özet

* `database/sql` → Uygulamanın kullandığı **yüksek seviye API**
* `database/sql/driver` → Veritabanı sürücülerinin implement ettiği **düşük seviye API**
* Sen genellikle `sql/driver` ile uğraşmazsın, ama **kendi driver’ını yazmak** istersen buradaki arayüzleri implement etmen gerekir.

---

İstersen ben sana `sql/driver` paketinde bulunan **tüm interface’leri tablo halinde** çıkarıp, her birine minik örnek fonksiyonlar da yazabilirim. Bunu ister misin?
EVET
*/
/*
O zaman şimdi **`database/sql/driver` paketindeki tüm interface’leri** tek tek açıklayıp, her birinin **ne işe yaradığını** ve **küçük örneklerini** göstereyim.

---

# 📌 `database/sql/driver` Paketindeki Tüm Interface’ler

## 1. **Driver**

* Yeni bir veritabanı bağlantısı açar.
* Bütün sürücülerin uygulaması gerekir.
*/
``go
type Driver interface {
    Open(name string) (Conn, error)
}
``

//**Örnek:**

``go
type MyDriver struct{}

func (d *MyDriver) Open(name string) (driver.Conn, error) {
    return &MyConn{}, nil
}
``
/*
---

## 2. **Conn**

* Açılmış bir bağlantıyı temsil eder.
* `Prepare`, `Close`, `Begin` metodları var.
*/
``go
type Conn interface {
    Prepare(query string) (Stmt, error)
    Close() error
    Begin() (Tx, error)
}
``

**Örnek:**

``go
type MyConn struct{}

func (c *MyConn) Prepare(query string) (driver.Stmt, error) {
    return &MyStmt{query: query}, nil
}
func (c *MyConn) Close() error { return nil }
func (c *MyConn) Begin() (driver.Tx, error) { return &MyTx{}, nil }
``

---

## 3. **Stmt (Statement)**

* Prepared statement’ı temsil eder.
* SQL cümlesi çalıştırılır (`Exec`, `Query`).

``go
type Stmt interface {
    Close() error
    NumInput() int
    Exec(args []Value) (Result, error)
    Query(args []Value) (Rows, error)
}
``

//**Örnek:**

``go
type MyStmt struct{ query string }

func (s *MyStmt) Close() error { return nil }
func (s *MyStmt) NumInput() int { return -1 }
func (s *MyStmt) Exec(args []driver.Value) (driver.Result, error) {
    return driver.RowsAffected(1), nil
}
func (s *MyStmt) Query(args []driver.Value) (driver.Rows, error) {
    return &MyRows{}, nil
}
``
/*
---

## 4. **Tx (Transaction)**

* Transaction (BEGIN/COMMIT/ROLLBACK) yönetimi.
*/
``go
type Tx interface {
    Commit() error
    Rollback() error
}
``

**Örnek:**

``go
type MyTx struct{}

func (t *MyTx) Commit() error   { return nil }
func (t *MyTx) Rollback() error { return nil }
``
/*
---

## 5. **Result**

* `INSERT/UPDATE/DELETE` sonrası dönen sonuç.
* `LastInsertId`, `RowsAffected` bilgilerini döner.
*/
``go
type Result interface {
    LastInsertId() (int64, error)
    RowsAffected() (int64, error)
}
``

//**Örnek:**

``go
type MyResult struct{}

func (r MyResult) LastInsertId() (int64, error) { return 1, nil }
func (r MyResult) RowsAffected() (int64, error) { return 1, nil }
``
/*
---

## 6. **Rows**

* `SELECT` sonucunu temsil eder.
* `Next` ile satırlar tek tek okunur.
*/
``go
type Rows interface {
    Columns() []string
    Close() error
    Next(dest []Value) error
}
``

//**Örnek:**

``go
type MyRows struct{}

func (r *MyRows) Columns() []string { return []string{"id", "name"} }
func (r *MyRows) Close() error      { return nil }
func (r *MyRows) Next(dest []driver.Value) error {
    dest[0] = int64(1)
    dest[1] = "Ali"
    return io.EOF // sadece 1 satır döndürüyoruz
}
``
/*
---

## 7. **Value ve ValueConverter**

* `Value`: Go tarafında bir DB değerini temsil eder.

  * İzin verilen tipler: `int64`, `float64`, `bool`, `[]byte`, `string`, `time.Time`, `nil`.
* `ValueConverter`: Go değerlerini DB’ye uygun hale getirir.
*/
``go
type Value interface{}
``

``go
type ValueConverter interface {
    ConvertValue(v interface{}) (Value, error)
}
``

//**Örnek:**

``go
type StringConverter struct{}

func (c StringConverter) ConvertValue(v interface{}) (driver.Value, error) {
    if s, ok := v.(string); ok {
        return s, nil
    }
    return nil, fmt.Errorf("sadece string kabul edilir")
}
``
/*
---

## 8. **Execer, Queryer (Opsiyonel Arayüzler)**

* Bazı driver’lar `Prepare` yerine doğrudan `Exec`/`Query` destekler.
*/
``go
type Execer interface {
    Exec(query string, args []Value) (Result, error)
}
type Queryer interface {
    Query(query string, args []Value) (Rows, error)
}
``
/*
---

## 9. **DriverContext, ConnPrepareContext, Connector vs.** (Gelişmiş sürümler)

Go 1.8+ ile geldi.

* **DriverContext**: `OpenConnector` ile daha gelişmiş bağlantı açar.
* **Connector**: `Driver` + `DSN` birleşimini temsil eder.
* **ConnPrepareContext**, **ExecerContext**, **QueryerContext**: Context destekli sürümler.

**Örnek:**
*/
``go
type Connector interface {
    Connect(context.Context) (Conn, error)
    Driver() Driver
}
``
/*
---

# 📊 Özet Tablo

| Arayüz                  | Amaç                       |
| ----------------------- | -------------------------- |
| Driver                  | Yeni bağlantı açar         |
| Conn                    | Bağlantıyı temsil eder     |
| Stmt                    | Hazırlanmış SQL cümlesi    |
| Tx                      | Transaction yönetimi       |
| Result                  | Exec sonrası bilgi         |
| Rows                    | Select sonucu satırlar     |
| Value                   | DB değer tipi              |
| ValueConverter          | Tip dönüşümü               |
| Execer/Queryer          | Alternatif Exec/Query      |
| DriverContext/Connector | Gelişmiş bağlantı yönetimi |

---

✅ Yani `sql/driver` paketinin **tamamı**, veritabanı sürücüsü yazmak için gerekli **alt seviye API**’leri sunar.
Sen uygulama yazarken genelde `database/sql` kullanırsın, ama driver geliştirenler bu interface’leri implemente eder.

---

İstersen ben bu interface’leri kullanarak **tam bir basit driver** (INSERT + SELECT destekleyen) örneği yazabilirim. Bunu ister misin?
EVET
*/
/*
Harika 👍 O zaman sana **tam çalışan basit bir memory driver örneği** göstereyim.

Bu driver:

* `database/sql/driver` kullanıyor
* `INSERT` ve `SELECT` destekli
* `database/sql` ile kullanılabilir

---

# 📌 Basit MemoryDB Driver
*/
``go
package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"log"
)

// ---------------- DRIVER ----------------
type MemoryDriver struct{}

func (d *MemoryDriver) Open(name string) (driver.Conn, error) {
	return &MemoryConn{
		users: []User{},
	}, nil
}

// ---------------- VERİ MODELİ ----------------
type User struct {
	ID   int64
	Name string
}

// ---------------- CONNECTION ----------------
type MemoryConn struct {
	users []User
	lastID int64
}

func (c *MemoryConn) Prepare(query string) (driver.Stmt, error) {
	return &MemoryStmt{conn: c, query: query}, nil
}
func (c *MemoryConn) Close() error { return nil }
func (c *MemoryConn) Begin() (driver.Tx, error) {
	return &MemoryTx{}, nil
}

// ---------------- STATEMENT ----------------
type MemoryStmt struct {
	conn  *MemoryConn
	query string
}

func (s *MemoryStmt) Close() error { return nil }
func (s *MemoryStmt) NumInput() int { return -1 }

func (s *MemoryStmt) Exec(args []driver.Value) (driver.Result, error) {
	// Sadece INSERT destekliyoruz: INSERT INTO users VALUES(?, ?)
	if len(args) >= 1 {
		s.conn.lastID++
		s.conn.users = append(s.conn.users, User{
			ID:   s.conn.lastID,
			Name: args[0].(string),
		})
		return driver.RowsAffected(1), nil
	}
	return driver.RowsAffected(0), nil
}

func (s *MemoryStmt) Query(args []driver.Value) (driver.Rows, error) {
	// SELECT * FROM users
	return &MemoryRows{users: s.conn.users, index: 0}, nil
}

// ---------------- TRANSACTION ----------------
type MemoryTx struct{}

func (t *MemoryTx) Commit() error   { return nil }
func (t *MemoryTx) Rollback() error { return nil }

// ---------------- ROWS ----------------
type MemoryRows struct {
	users []User
	index int
}

func (r *MemoryRows) Columns() []string {
	return []string{"id", "name"}
}
func (r *MemoryRows) Close() error { return nil }
func (r *MemoryRows) Next(dest []driver.Value) error {
	if r.index >= len(r.users) {
		return io.EOF
	}
	dest[0] = r.users[r.index].ID
	dest[1] = r.users[r.index].Name
	r.index++
	return nil
}

// ---------------- MAIN ----------------
func main() {
	sql.Register("memorydb", &MemoryDriver{})

	db, err := sql.Open("memorydb", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// INSERT
	db.Exec("INSERT INTO users VALUES(?)", "Ahmet")
	db.Exec("INSERT INTO users VALUES(?)", "Mehmet")

	// SELECT
	rows, _ := db.Query("SELECT * FROM users")
	defer rows.Close()

	fmt.Println("Kullanıcılar:")
	for rows.Next() {
		var id int64
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
``
/*
---

# 📌 Açıklama

1. `MemoryDriver` → driver’ı register ediyoruz.
2. `MemoryConn` → bağlantıyı tutuyor, `users` slice’ı hafızada veritabanı görevi görüyor.
3. `MemoryStmt` → sadece `INSERT` ve `SELECT` destekli.
4. `MemoryRows` → `SELECT` sonucu satır satır okuma.
5. `MemoryTx` → transaction implementasyonu (dummy).
6. `database/sql` üzerinden tamamen normal bir veritabanı gibi kullanılabilir.

---

✅ Bu örnek **tamamen memory tabanlı bir driver**.
İstersen bir sonraki adımda bunu **UPDATE ve DELETE destekleyecek şekilde** geliştirebiliriz ve tamamen **mini CRUD driver** haline getirebiliriz.
*/
