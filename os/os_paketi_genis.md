| Fonksiyon                         | Açıklama                                        | Örnek Kod                                                        |
|:----------------------------------|:------------------------------------------------|:-----------------------------------------------------------------|
| os.Create(name)                   | Yeni dosya oluşturur, yazma modunda açar        | f, _ := os.Create("a.txt")                                       |
| os.Open(name)                     | Dosyayı okuma modunda açar                      | f, _ := os.Open("a.txt")                                         |
| os.OpenFile(name, flag, perm)     | Dosyayı özel modlarla açar                      | f, _ := os.OpenFile("a.txt", os.O_RDWR, 0644)                    |
| os.Remove(name)                   | Dosyayı siler                                   | os.Remove("a.txt")                                               |
| os.Rename(old, new)               | Dosyayı yeniden adlandırır/taşır                | os.Rename("a.txt", "b.txt")                                      |
| os.Stat(name)                     | Dosya bilgilerini döndürür                      | info, _ := os.Stat("a.txt")                                      |
| os.Lstat(name)                    | Sembolik bağlantı bilgilerini döndürür          | info, _ := os.Lstat("symlink")                                   |
| os.ReadFile(name)                 | Dosya içeriğini []byte olarak okur              | data, _ := os.ReadFile("a.txt")                                  |
| os.WriteFile(name, data, perm)    | Dosyaya []byte yazar                            | os.WriteFile("a.txt", []byte("Merhaba"), 0644)                   |
| os.Truncate(name, size)           | Dosyayı belirtilen boyuta keser                 | os.Truncate("a.txt", 0)                                          |
| os.Chmod(name, mode)              | Dosyanın izinlerini değiştirir                  | os.Chmod("a.txt", 0644)                                          |
| os.Chown(name, uid, gid)          | Dosyanın kullanıcı/grup sahipliğini değiştirir  | os.Chown("a.txt", 1000, 1000)                                    |
| os.Chtimes(name, atime, mtime)    | Dosya erişim/değişiklik zamanlarını değiştirir  | os.Chtimes("a.txt", time.Now(), time.Now())                      |
| os.Symlink(oldname, newname)      | Sembolik bağlantı oluşturur                     | os.Symlink("a.txt", "link.txt")                                  |
| os.Readlink(name)                 | Sembolik bağlantının hedefini döndürür          | target, _ := os.Readlink("link.txt")                             |
| os.Link(oldname, newname)         | Sabit bağlantı (hard link) oluşturur            | os.Link("a.txt", "hardlink.txt")                                 |
| os.TempDir()                      | Geçici dizin yolunu döndürür                    | fmt.Println(os.TempDir())                                        |
| os.CreateTemp(dir, pattern)       | Geçici dosya oluşturur                          | f, _ := os.CreateTemp("", "example-*.txt")                       |
| os.MkdirTemp(dir, pattern)        | Geçici dizin oluşturur                          | d, _ := os.MkdirTemp("", "example-*")                            |
| os.Mkdir(name, perm)              | Yeni bir dizin oluşturur                        | os.Mkdir("dir", 0755)                                            |
| os.MkdirAll(path, perm)           | İç içe dizinler oluşturur                       | os.MkdirAll("a/b/c", 0755)                                       |
| os.RemoveAll(path)                | Dizini ve içindekileri siler                    | os.RemoveAll("dir")                                              |
| os.Chdir(dir)                     | Çalışma dizinini değiştirir                     | os.Chdir("/tmp")                                                 |
| os.Getwd()                        | Mevcut çalışma dizinini döndürür                | dir, _ := os.Getwd()                                             |
| os.ReadDir(name)                  | Dizin içeriğini listeler                        | entries, _ := os.ReadDir(".")                                    |
| os.Getenv(key)                    | Ortam değişkeni değerini alır                   | os.Getenv("PATH")                                                |
| os.Setenv(key, value)             | Ortam değişkeni tanımlar                        | os.Setenv("API_KEY", "123")                                      |
| os.Unsetenv(key)                  | Ortam değişkenini siler                         | os.Unsetenv("API_KEY")                                           |
| os.Environ()                      | Tüm ortam değişkenlerini döndürür               | envs := os.Environ()                                             |
| os.ExpandEnv(s)                   | Ortamdaki değişkenleri string içinde genişletir | fmt.Println(os.ExpandEnv("$HOME/config"))                        |
| os.UserHomeDir()                  | Kullanıcı home dizinini döndürür                | home, _ := os.UserHomeDir()                                      |
| os.StartProcess(name, argv, attr) | Yeni işlem başlatır                             | os.StartProcess("/bin/ls", []string{"ls", "-l"}, &os.ProcAttr{}) |
| os.FindProcess(pid)               | PID ile işlem bulur                             | p, _ := os.FindProcess(1234)                                     |
| os.Getpid()                       | Mevcut işlem PID'si                             | pid := os.Getpid()                                               |
| os.Getppid()                      | Parent PID                                      | ppid := os.Getppid()                                             |
| os.ErrNotExist                    | Dosya/dizin bulunamadı hatası                   | errors.Is(err, os.ErrNotExist)                                   |
| os.ErrExist                       | Dosya zaten var hatası                          | errors.Is(err, os.ErrExist)                                      |
| os.ErrPermission                  | İzin hatası                                     | errors.Is(err, os.ErrPermission)                                 |
| os.ErrInvalid                     | Geçersiz argüman hatası                         | errors.Is(err, os.ErrInvalid)                                    |
| os.Exit(code)                     | Programı belirtilen kod ile sonlandırır         | os.Exit(0)                                                       |
| os.Args                           | Komut satırı argümanlarını döndürür             | fmt.Println(os.Args)                                             |
| os.Executable()                   | Çalışan binary yolunu döndürür                  | exe, _ := os.Executable()                                        |
| os.Hostname()                     | Sistemin hostname bilgisini döndürür            | host, _ := os.Hostname()                                         |