| Fonksiyon   | Açıklama                                   | Örnek                                                                      |
|:------------|:-------------------------------------------|:---------------------------------------------------------------------------|
| Abs         | Verilen yolu mutlak yola çevirir.          | filepath.Abs("docs/readme.txt") → "/home/user/docs/readme.txt"             |
| Base        | Yolun en son öğesini döndürür.             | filepath.Base("/home/user/file.txt") → "file.txt"                          |
| Clean       | Yolu temizler (.., //, . kaldırır).        | filepath.Clean("/home/user/../docs//file.txt") → "/home/docs/file.txt"     |
| Dir         | Yolun dizin kısmını döndürür.              | filepath.Dir("/home/user/file.txt") → "/home/user"                         |
| Ext         | Dosya uzantısını döndürür.                 | filepath.Ext("report.pdf") → ".pdf"                                        |
| Join        | Parçaları güvenli şekilde birleştirir.     | filepath.Join("home","user","docs","file.txt") → "home/user/docs/file.txt" |
| Match       | Pattern ile eşleşmeyi kontrol eder.        | filepath.Match("*.txt", "notes.txt") → true                                |
| Glob        | Pattern’e uyan dosyaları bulur.            | filepath.Glob("*.go") → ["main.go", "utils.go"]                            |
| Rel         | İki yol arasındaki göreceli farkı bulur.   | filepath.Rel("/home/user","/home/user/docs/file.txt") → "docs/file.txt"    |
| Split       | Yolu dizin ve dosya adı olarak ayırır.     | filepath.Split("/home/user/file.txt") → "/home/user/", "file.txt"          |
| VolumeName  | Windows için disk sürücüsünü döndürür.     | filepath.VolumeName("C:\Users\file.txt") → "C:"                            |
| WalkDir     | Dizini derinlemesine dolaşır (yeni API).   | filepath.WalkDir(".", fn) → tüm dosyalar listelenir                        |
| Walk        | Dizini derinlemesine dolaşır (eski API).   | filepath.Walk(".", fn) → tüm dosyalar listelenir                           |
| ToSlash     | Ayracı `/` olarak değiştirir.              | filepath.ToSlash("C:\Users\Admin\file.txt") → "C:/Users/Admin/file.txt"    |
| FromSlash   | `/` işaretini sistem ayracına çevirir.     | filepath.FromSlash("C:/Users/Admin/file.txt") → "C:\Users\Admin\file.txt"  |
| IsAbs       | Yolun mutlak olup olmadığını kontrol eder. | filepath.IsAbs("/usr/local/bin") → true                                    |