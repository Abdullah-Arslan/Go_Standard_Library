# Go `os` Paketi Örnekleri

Bu proje, Go dilinde `os` paketinin farklı fonksiyonlarının nasıl kullanıldığını gösteren örneklerden oluşur.

## Çalıştırma

Her dosyayı tek tek çalıştırabilirsin:

```bash
go run examples/chmod.go
go run examples/symlink.go
```

## İçerik

- `chmod.go` → Dosya izinlerini değiştirme
- `chown.go` → Dosya sahipliğini değiştirme
- `chtimes.go` → Dosya zamanlarını değiştirme
- `symlink.go` → Sembolik link oluşturma ve okuma
- `link.go` → Hard link oluşturma
- `truncate.go` → Dosya boyutunu değiştirme
- `temp.go` → Geçici dosya/dizin oluşturma
- `expandenv.go` → Ortam değişkenlerini genişletme
- `userhomedir.go` → Kullanıcı home dizini
- `executable.go` → Çalışan binary yolu
