/*
Go’da **`os/user`** paketi, işletim sistemi üzerinde **kullanıcı bilgilerini** almayı sağlar.
Özellikle **kimlik doğrulama, dosya izinleri, kullanıcıya özel ayarlar** gibi yerlerde kullanılır.

---

# 🔹 `os/user` Paketinin Temel İşlevleri

### 1. `user.Current()`

Programı çalıştıran **mevcut kullanıcıyı** döndürür.

### 2. `user.Lookup(username string)`

Belirtilen kullanıcı adını arar, `*User` döner.

### 3. `user.LookupId(uid string)`

Kullanıcıyı UID ile bulur.

### 4. `user.LookupGroup(name string)`

Grup bilgilerini isimle arar.

### 5. `user.LookupGroupId(gid string)`

Grubu GID ile arar.

### 6. `user.CurrentGroup()`

Mevcut kullanıcının **birincil grubunu** verir. (Go 1.19+ ile geldi.)

### 7. `user.CurrentGroups()`

Mevcut kullanıcının **tüm grup ID’lerini** verir.

---

# 🔹 `User` Yapısı

`os/user.User` struct’ı şunları içerir:

* `Username` → Kullanıcı adı
* `Uid` → Kullanıcının UID’si
* `Gid` → Birincil grup ID’si
* `Name` → Kullanıcının tam adı (genellikle `/etc/passwd` içinden)
* `HomeDir` → Kullanıcının ana dizini

---

# 🔹 Örnekler

## 1. Mevcut Kullanıcıyı Bulma
/*
``go
package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Kullanıcı adı:", u.Username)
	fmt.Println("UID:", u.Uid)
	fmt.Println("GID:", u.Gid)
	fmt.Println("Tam Ad:", u.Name)
	fmt.Println("Home dizini:", u.HomeDir)
}
``

//📌 Çıktı (Linux/Mac):

``
Kullanıcı adı: myuser
UID: 1000
GID: 1000
Tam Ad: My User
Home dizini: /home/myuser
``
/*
---

## 2. Kullanıcı Adı ile Kullanıcı Bulma
*/
``go
package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, err := user.Lookup("root")
	if err != nil {
		panic(err)
	}
	fmt.Println("Root UID:", u.Uid)
	fmt.Println("Root GID:", u.Gid)
}
``
/*
---

## 3. UID ile Kullanıcı Bulma
*/
``go
package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, err := user.LookupId("1000")
	if err != nil {
		panic(err)
	}
	fmt.Println("UID 1000 kullanıcı adı:", u.Username)
}
``
/*
---

## 4. Grup Bilgisi Alma
*/
``go
package main

import (
	"fmt"
	"os/user"
)

func main() {
	g, err := user.LookupGroup("sudo")
	if err != nil {
		panic(err)
	}
	fmt.Println("Grup adı:", g.Name)
	fmt.Println("GID:", g.Gid)
}
``
/*
---

## 5. Mevcut Kullanıcının Grubunu Bulma
*/
``go
package main

import (
	"fmt"
	"os/user"
)

func main() {
	g, err := user.CurrentGroup()
	if err != nil {
		panic(err)
	}
	fmt.Println("Kullanıcının birincil grubu:", g.Name, " (GID:", g.Gid, ")")
}
``
/*
---

## 6. Mevcut Kullanıcının Tüm Gruplarını Listeleme
*/
``go
package main

import (
	"fmt"
	"os/user"
)

func main() {
	groups, err := user.CurrentGroups()
	if err != nil {
		panic(err)
	}
	fmt.Println("Kullanıcının grup ID’leri:", groups)
}
``
/*
---

# 🔹 Kullanım Alanları

* **Kimlik doğrulama** (kimin giriş yaptığını bilmek)
* **Dosya izinleri ve sahiplik** kontrolü
* **Kullanıcıya özel config dosyaları** bulmak
* **Multi-user sistemlerde yetki yönetimi**

---

# 🔹 Özet

* `os/user` paketi kullanıcı ve grup bilgilerini sağlar.
* `Current`, `Lookup`, `LookupId`, `LookupGroup`, `LookupGroupId`, `CurrentGroup`, `CurrentGroups` gibi fonksiyonlar bulunur.
* Kullanıcıya ait: **UID, GID, Username, Full Name, HomeDir** bilgileri alınabilir.

---

👉 İstersen ben sana bu paketin fonksiyonlarını tablo (Fonksiyon + Açıklama + Örnek Kod) halinde **Excel + PDF** formatında da hazırlayabilirim.
Bunu ister misin?
EVET

👉 İstersen sana bunun üzerine bir de “kullanıcı tabanlı erişim kontrolü” mini proje örneği hazırlayayım (ör. sadece belirli kullanıcı çalıştırabilsin). İstiyor musun?
EVET
*/
/*
Süper 👍 O zaman sana `os/user` paketini gerçek bir senaryoda göstereyim:
Bir **“kullanıcı tabanlı erişim kontrolü”** mini proje yapacağız.

👉 Senaryo:

* Program sadece belirli kullanıcılar tarafından çalıştırılabilsin.
* Eğer yetkisiz bir kullanıcı çalıştırırsa, program kapanacak.
* Ayrıca yetkili kullanıcıya özel mesaj gösterecek.

---

# 🔹 Örnek: Kullanıcı Tabanlı Erişim Kontrolü

## 📄 `main.go`
*/
``go
package main

import (
	"fmt"
	"os"
	"os/user"
)

// Yetkili kullanıcı listesi
var allowedUsers = []string{"root", "mabdu", "admin"}

func isAllowed(username string) bool {
	for _, u := range allowedUsers {
		if u == username {
			return true
		}
	}
	return false
}

func main() {
	// Mevcut kullanıcıyı al
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Kullanıcı bilgisi alınamadı:", err)
		os.Exit(1)
	}

	fmt.Println("🔹 Giriş yapan kullanıcı:", currentUser.Username)

	// Kullanıcı kontrolü
	if !isAllowed(currentUser.Username) {
		fmt.Println("❌ Yetkisiz kullanıcı! Programdan çıkılıyor...")
		os.Exit(1)
	}

	// Yetkili kullanıcı için işlem
	fmt.Println("✅ Hoş geldin", currentUser.Username, "!")
	fmt.Println("Programı kullanmaya devam edebilirsin...")
}
``
/*
---

## 🖥 Kullanım

1. Programı çalıştır:
*/
   ``bash
   go run main.go
   ``

//2. Eğer **yetkili kullanıcı** (örneğin `mabdu`):

   ``
   🔹 Giriş yapan kullanıcı: mabdu
   ✅ Hoş geldin mabdu !
   Programı kullanmaya devam edebilirsin...
   ``

//3. Eğer **yetkisiz kullanıcı**:

   ``
   🔹 Giriş yapan kullanıcı: testuser
   ❌ Yetkisiz kullanıcı! Programdan çıkılıyor...
   ``
/*
---

# 🔹 Geliştirme Fikirleri

Bunu daha ileri taşıyabiliriz:

* Kullanıcıyı **UID** veya **GID** ile kontrol etmek.
* Belirli bir **grup** üyesi olanlara izin vermek (ör. `sudo` grubundakiler).
* Kullanıcı adına göre farklı **izin seviyeleri** (ör. admin → tüm yetkiler, user → sınırlı).
* Bir **config dosyası** ile yetkili kullanıcıları dışarıdan tanımlamak.

---

👉 İstersen sana bu örneğin **“grup bazlı yetkilendirme”** sürümünü de yazabilirim (ör. sadece `sudo` grubundaki kullanıcılar çalıştırabilsin). Onu da yapayım mı?
EVET
*/
/*
Harika 👍 Şimdi kullanıcı kontrolünü bir adım ileri taşıyalım:
Bu sürümde **“grup bazlı yetkilendirme”** yapacağız.

👉 Senaryo:

* Program sadece belirli bir gruptaki kullanıcılar tarafından çalıştırılabilsin (örneğin `sudo` grubu).
* Kullanıcı o gruba dahil değilse program kapanacak.
* Böylece yetkilendirmeyi kullanıcı yerine **grup üyeliği** ile yapmış olacağız.

---

# 🔹 Örnek: Grup Bazlı Yetkilendirme

## 📄 `main.go`
*/
``go
package main

import (
	"fmt"
	"os"
	"os/user"
)

// Yetkili grup adı
const allowedGroup = "sudo"

func main() {
	// Mevcut kullanıcı bilgisi
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Kullanıcı bilgisi alınamadı:", err)
		os.Exit(1)
	}

	fmt.Println("🔹 Giriş yapan kullanıcı:", currentUser.Username)

	// Kullanıcının grup ID’lerini al
	groupIDs, err := user.CurrentGroups()
	if err != nil {
		fmt.Println("Grup bilgisi alınamadı:", err)
		os.Exit(1)
	}

	// Kullanıcı bu gruplardan birine üye mi?
	authorized := false
	for _, gid := range groupIDs {
		g, err := user.LookupGroupId(gid)
		if err == nil && g.Name == allowedGroup {
			authorized = true
			break
		}
	}

	if !authorized {
		fmt.Println("❌ Kullanıcı", currentUser.Username, "grup yetkisine sahip değil! Programdan çıkılıyor...")
		os.Exit(1)
	}

	// Yetkili kullanıcıya mesaj
	fmt.Println("✅ Hoş geldin", currentUser.Username, "!")
	fmt.Println("Bu programı çalıştırmaya iznin var çünkü", allowedGroup, "grubundasın.")
}
``
/*
---

# 🖥 Kullanım

1. Programı çalıştır:
*/
   ``bash
   go run main.go
   ``

//2. Eğer kullanıcı `sudo` grubundaysa:

   ``
   🔹 Giriş yapan kullanıcı: mabdu
   ✅ Hoş geldin mabdu !
   Bu programı çalıştırmaya iznin var çünkü sudo grubundasın.
   ``

//3. Eğer kullanıcı `sudo` grubunda değilse:

   ``
   🔹 Giriş yapan kullanıcı: testuser
   ❌ Kullanıcı testuser grup yetkisine sahip değil! Programdan çıkılıyor...
   ``
/*
---

# 🔹 Geliştirme Fikirleri

* Birden fazla **izinli grup** tanımlanabilir (örn: `admins`, `devops`, `sudo`).
* Kullanıcıya **farklı yetki seviyeleri** verilebilir (örn: admin → tam yetki, user → sadece okuma).
* **Config dosyasından** yetkili grup(lar) okunabilir.
* `os/signal` ile birleşip, `SIGHUP` alındığında izinli grup listesi yeniden yüklenebilir.

---

👉 İstersen sana bir sonraki adımda **“çoklu grup + farklı yetki seviyeleri”** (ör. admin/devops/user) destekli sürümü de hazırlayabilirim. Onu da yapayım mı?
EVET
*/
/*
Harika 👍 Şimdi kullanıcı kontrolünü bir adım ileri taşıyalım:
Bu sürümde **“grup bazlı yetkilendirme”** yapacağız.

👉 Senaryo:

* Program sadece belirli bir gruptaki kullanıcılar tarafından çalıştırılabilsin (örneğin `sudo` grubu).
* Kullanıcı o gruba dahil değilse program kapanacak.
* Böylece yetkilendirmeyi kullanıcı yerine **grup üyeliği** ile yapmış olacağız.

---

# 🔹 Örnek: Grup Bazlı Yetkilendirme

## 📄 `main.go`
*/
``go
package main

import (
	"fmt"
	"os"
	"os/user"
)

// Yetkili grup adı
const allowedGroup = "sudo"

func main() {
	// Mevcut kullanıcı bilgisi
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Kullanıcı bilgisi alınamadı:", err)
		os.Exit(1)
	}

	fmt.Println("🔹 Giriş yapan kullanıcı:", currentUser.Username)

	// Kullanıcının grup ID’lerini al
	groupIDs, err := user.CurrentGroups()
	if err != nil {
		fmt.Println("Grup bilgisi alınamadı:", err)
		os.Exit(1)
	}

	// Kullanıcı bu gruplardan birine üye mi?
	authorized := false
	for _, gid := range groupIDs {
		g, err := user.LookupGroupId(gid)
		if err == nil && g.Name == allowedGroup {
			authorized = true
			break
		}
	}

	if !authorized {
		fmt.Println("❌ Kullanıcı", currentUser.Username, "grup yetkisine sahip değil! Programdan çıkılıyor...")
		os.Exit(1)
	}

	// Yetkili kullanıcıya mesaj
	fmt.Println("✅ Hoş geldin", currentUser.Username, "!")
	fmt.Println("Bu programı çalıştırmaya iznin var çünkü", allowedGroup, "grubundasın.")
}
``
/*
---

# 🖥 Kullanım

1. Programı çalıştır:
*/
   ``bash
   go run main.go
   ``

//2. Eğer kullanıcı `sudo` grubundaysa:

   ``
   🔹 Giriş yapan kullanıcı: mabdu
   ✅ Hoş geldin mabdu !
   Bu programı çalıştırmaya iznin var çünkü sudo grubundasın.
   ``

/73. Eğer kullanıcı `sudo` grubunda değilse:

   ``
   🔹 Giriş yapan kullanıcı: testuser
   ❌ Kullanıcı testuser grup yetkisine sahip değil! Programdan çıkılıyor...
   ``
/*
---

# 🔹 Geliştirme Fikirleri

* Birden fazla **izinli grup** tanımlanabilir (örn: `admins`, `devops`, `sudo`).
* Kullanıcıya **farklı yetki seviyeleri** verilebilir (örn: admin → tam yetki, user → sadece okuma).
* **Config dosyasından** yetkili grup(lar) okunabilir.
* `os/signal` ile birleşip, `SIGHUP` alındığında izinli grup listesi yeniden yüklenebilir.

---

👉 İstersen sana bir sonraki adımda **“çoklu grup + farklı yetki seviyeleri”** (ör. admin/devops/user) destekli sürümü de hazırlayabilirim. Onu da yapayım mı?
EVET
*/
/*
Harika 👍 Şimdi son sürümü biraz daha profesyonel hale getirelim:

👉 **Senaryo: Çoklu grup + farklı yetki seviyeleri**

* Kullanıcı hangi gruptaysa ona göre **farklı rol** atanacak.
* Örneğin:

  * `admin` → Tam yetki
  * `devops` → Sunucu yönetim yetkisi
  * `user` → Sadece okuma yetkisi
* Eğer kullanıcı bu gruplardan hiçbirine üye değilse program kapanacak.

---

# 🔹 Örnek: Çoklu Grup + Yetki Seviyeleri

## 📄 `main.go`
*/
``go
package main

import (
	"fmt"
	"os"
	"os/user"
)

// Yetkili gruplar ve roller
var groupRoles = map[string]string{
	"admin":  "Tam Yetki (Admin)",
	"devops": "Sunucu Yönetim Yetkisi",
	"user":   "Okuma Yetkisi",
}

func main() {
	// Mevcut kullanıcıyı al
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Kullanıcı bilgisi alınamadı:", err)
		os.Exit(1)
	}

	fmt.Println("🔹 Giriş yapan kullanıcı:", currentUser.Username)

	// Kullanıcının grup ID'lerini al
	groupIDs, err := user.CurrentGroups()
	if err != nil {
		fmt.Println("Grup bilgisi alınamadı:", err)
		os.Exit(1)
	}

	// Kullanıcıya atanacak rol
	var role string

	for _, gid := range groupIDs {
		g, err := user.LookupGroupId(gid)
		if err != nil {
			continue
		}
		if r, ok := groupRoles[g.Name]; ok {
			role = r
			break
		}
	}

	if role == "" {
		fmt.Println("❌ Yetkisiz kullanıcı! Hiçbir izinli grupta bulunmuyor.")
		os.Exit(1)
	}

	fmt.Println("✅ Hoş geldin", currentUser.Username, "!")
	fmt.Println("Rolün:", role)
}
``
/*
---

# 🖥 Kullanım

1. Kullanıcı `admin` grubundaysa:
*/
   ``
   🔹 Giriş yapan kullanıcı: mabdu
   ✅ Hoş geldin mabdu !
   Rolün: Tam Yetki (Admin)
   ``

/72. Kullanıcı `devops` grubundaysa:

   ``
   🔹 Giriş yapan kullanıcı: testdev
   ✅ Hoş geldin testdev !
   Rolün: Sunucu Yönetim Yetkisi
   ``

//3. Kullanıcı `user` grubundaysa:

   ``
   🔹 Giriş yapan kullanıcı: normaluser
   ✅ Hoş geldin normaluser !
   Rolün: Okuma Yetkisi
   ``

//4. Kullanıcı bu gruplarda değilse:

   ``
   🔹 Giriş yapan kullanıcı: guest
   ❌ Yetkisiz kullanıcı! Hiçbir izinli grupta bulunmuyor.
   ``
/*
---

# 🔹 Geliştirme Fikirleri

* Yetkili grupları **config dosyasından** okumak (dinamik yapı).
* `SIGHUP` sinyali ile **rolleri runtime’da yeniden yüklemek**.
* Kullanıcıya göre farklı **menüler veya komut erişimleri** açmak.
* Loglama eklemek (kim girdi, hangi rol aldı).

---

👉 İstersen sana bu sistemi bir adım daha ileri taşıyıp **config dosyası tabanlı dinamik rol yönetimi + SIGHUP ile canlı reload** sürümünü de yazabilirim.
Bunu da yapayım mı?
*/