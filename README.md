# Problem

Happy Moon's sitesi üzerindeki ürünlerin bilgilerinin saklanıp, bir API
aracılığıyla kolayca erişilebilmesi için bir uygulama yazılması istenmektedir.

> Gizlilik sınıflandırması nedeniyle, bu dokümantasyonda belirtilen kurumun adı açıkça belirtilmemiştir.

Kodun detaylı dökümantasyonu için `app/README.md` dosyasındaki yönergeleri takip ediniz.

# İsterler

- Verilerin [verilen bağlantı](https://happymoonsemaar.jacca.com/Menu?C=view_tr_happymoonsemaar_99) üzerinden çekilmesi.
- Verilerin bir veritabanında saklanması.
  - İstenilen veriler: `Kategori`, `Ürün`, `Açıklama`, `Fiyat`
- Verilerin `Go` programlama dili ile yazılmış bir API üzerinden çekilmesi.
- Servisin belirli bir konfigürasyon dosyası aracılığıyla yönetilebilmesi.
- Servisin SOLID prensiplerine uygun olması.
- Servisin çalışma portunun istemciden saklanması.
- Servisin API Gateway aracılığıyla merkezi bir noktadan yönetilebilmesi.
- Belirtilen endpointlerin çalışması:
    - `GET /happymoons` : Tüm ürünlerin listelenmesi.
    - `GET /happymoons/ex={kategori1,kategori2,...}` : Belirtilen kategorilerin hariç tutularak listelenmesi.
    - `GET /happymoons/in={kategori1,kategori2,...}` : Belirtilen kategorilerin listelenmesi.
    - `GET /happymoons/csv` : Tüm ürünlerin CSV formatında listelenmesi.
- Servisin Dockerize edilmesi.

# Çözüm

## Kullanılan Yöntemler ve Karar Süreçleri

### Genel Bakış

- Çözüm olarak geliştirilen geliştirlen API, daha stabil ve erişilebilir olması adına [REST API](https://en.wikipedia.org/wiki/Representational_state_transfer) mimarisine uygun olarak geliştirilmiştir.

- Verilerin istenilen siteden çekilmesi için [web kazıma](https://en.wikipedia.org/wiki/Web_scraping) yöntemi kullanılmıştır.

- İstenen veriler yapısal olduğundan verilerin saklanması için yüksek performanslı, tutarlı ve güvenilir bir veritabanı kullanılması gerekmektedir. Bu nedenle [PostgreSQL](https://www.postgresql.org/) veritabanı kullanılmıştır.

- Kazınan verilerin saklanması ve çekilmesi için [ORM](https://en.wikipedia.org/wiki/Object-relational_mapping) kullanılmıştır. ORM kullanılmasının sebebi, veritabanı işlemlerinin daha kolay ve hızlı ve güvenilir bir şekilde yapılmasıdır.

- Servisin [konteynerize](https://en.wikipedia.org/wiki/Containerization) edilerek daha kolay bir şekilde çalıştırılması ve dağıtılması için [Docker](https://www.docker.com/) kullanılmıştır. Docker kullanılarak aynı zamanda servisin portunun istemciden saklanması sağlanmıştır.

- Uygulama ve veri katmanları Docker aracılığıyla izole edilerek, servisin daha güvenli ve kontrol edilebilir bir şekilde çalıştırılması sağlanmıştır. Bu izole edilmiş konteynerlerin bağlantıları [Docker Compose](https://docs.docker.com/compose/) aracılığıyla yönetilmektedir.

- Servisin tek bir dosya üzerinden konfigüre edilmesi için [TOML](https://toml.io/en/) dosya formatı kullanılmıştır.

- Servis API Gateway aracılığıyla merkezi bir noktadan yönetilebilmektedir. Güvenlik için [Rate Limiting](https://en.wikipedia.org/wiki/Rate_limiting) ve [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) ayarları yapılmıştır. Servis erişimi için herhangi bir kısıtlama yapılmadığından ve bir kimlik doğrulama sistemi kullanılmadığından, Rate Limiting için basit bir [token bucket](https://en.wikipedia.org/wiki/Token_bucket) algoritması kullanılmıştır. Güvenlik için CORS ayarları ile whitelist'e eklenen domainlerden gelen isteklerin kabul edilmesi sağlanmıştır.

- Servis uluslararası standartlara uygun olarak geliştirilmiştir. (Paketlerin ve fonksiyonların isimlendirilmesi, kodlama standartları, dökümantasyon, tasarım, dil vb.)

### Kullanılan Teknolojiler ve Kütüphaneler

- [Go](https://golang.org/): REST API dizaynı ve geliştirilmesi için kullanılmıştır.
- [PostgreSQL](https://www.postgresql.org/): Verilerin saklanması için kullanılmıştır.
- [Docker](https://www.docker.com/): Servisin Dockerize edilmesi için kullanılmıştır.
- [GoQuery](https://github.com/PuerkitoBio/goquery): Web kazıma işlemleri için kullanılmıştır.
- [GORM](https://gorm.io/): Veritabanı işlemleri için kullanılmıştır.
- [go-toml](github.com/pelletier/go-toml): Konfigürasyon dosyasının parse edilmesi için kullanılmıştır.
- [gorilla/mux](https://github.com/gorilla/mux): API yönlendirme işlemleri için kullanılmıştır.


## Servisin Yapısı

- Servis, `app` dizini altında bulunmaktadır. Dosya yapısı aşağıdaki gibidir:

```
│   .dockerignore
│   .gitignore
│   config.toml
│   Dockerfile
│   go.mod
│   go.sum
│   main.go
│   README.md
├───api
│   │   gateway.go
│   ├───controllers
│   │       productController.go
│   ├───middlewares
│   │       404.go
│   │       CORS.go
│   │       ratelimit.go
│   ├───responses
│   │       error.go
│   │       json.go
│   │       responses.go
│   └───routers
│           productRouter.go
├───config
│       loader.go
├───database
│       database.go
├───models
│       models.go
│       product.go
│       scraper.go
├───scraper
│       scraper.go
└───utils
    │   utils.go
    ├───console
    │       console.go
    └───limiter
            limiter.go

```

`app` dizini altında bulunan dosyaların görevleri kısaca aşağıdaki gibidir:

- `main.go`: Servisin çalıştırılması için kullanılan dosyadır.
- `config.toml`: Servisin konfigürasyon dosyasıdır. Servis çalıştırıldığında bu dosya okunur ve servis için gerekli ayarlar yapılır.
- `Dockerfile`: Servisin Dockerize edilmesi için kullanılan dosyadır.
- `/utils`: Kullanılan kütüphaneler için yardımcı fonksiyonları içerir.
    - `/limiter`
        - `limiter.go`: Token Bucket algoritması ile Rate Limiting işlemlerini gerçekleştiren dosyadır.
    - `/console`
        - `console.go`: Konsol çıktılarını stilize eden dosyadır.
- `/scraper`:
    - `scraper.go`: Web kazıma işlemlerini gerçekleştiren dosyadır.
- `/config`:
    - `loader.go`: Servis için gerekli konfigürasyon ayarlarını yükleyen dosyadır.
- `/database`:
    - `database.go`: Veritabanı ile bağlantıyı sağlayan dosyadır.
- `/models`: Uygulama modellerini içerir.
    - `product.go`: Ürün modelini içerir.
    - `scraper.go`: Scraper arayüzünü içerir.
- `/api`: API dosyalarının bulunduğu dizindir.
  - `gateway.go`: API Gateway dosyasıdır. İsteklerin yönlendirilmesi ve kontrolü için kullanılır.
    - `/routers`: API endpointlerinin yönlendirilmesi için kullanılan dosyaları içerir.
      - `productRouter.go`: Ürünler için API yönlendirme işlemlerini gerçekleştiren dosyadır.
    - `/middlewares`: API isteklerinin kontrolü için kullanılan dosyaları içerir.
        - `404.go`: API endpointi bulunamadığında kullanılan dosyadır.
        - `CORS.go`: CORS ayarlarının yapıldığı dosyadır.
        - `ratelimit.go`: Rate Limiting işlemlerinin yapıldığı dosyadır.
    - `/controllers`: API endpointlerinin işlemlerinin gerçekleştirildiği dosyaları içerir.
        - `productController.go`: Ürünler için API endpointlerinin işlemlerini gerçekleştiren dosyadır.
    - `/responses`: API endpointlerinin cevaplarının oluşturulduğu dosyaları içerir.
        - `error.go`: Hata mesajlarının oluşturulduğu dosyadır.
        - `json.go`: JSON cevaplarının oluşturulduğu dosyadır.
  

Her bir dosyanın görevi ve sorumluluğuyla ilgli detaylı bilgi için `app/README.md` dosyasındaki yönergeleri takip ediniz.

## Uygulama Çalıştırılması

- Uygulamanın çalıştırılması için öncelikle sisteminizde Go ve Docker kurulu olmalıdır.

- Uygulamayı `go run main.go` komutu ile çalıştıramazsınız. Çünkü servis yalnızca Docker Composer ile veritabanı ile birlikte çalıştırılabilir. 

- Uygulamayı çalıştırmak için `docker-compose build && docker-compose up -d` komutunu çalıştırmanız yeterlidir. Bu komut ile uygulama Dockerize edilmiş halde çalıştırılacaktır.

- `docker-compose logs` komutu ile uygulamanın çalıştığı konsol çıktılarını görebilirsiniz. Aşağıdaki çıktılar uygulamanın başarılı bir şekilde çalıştığı anlamına gelmektedir:

```
2023-04-02 23:09:05 2023/04/02 20:09:05 /go/src/app/main.go:45
2023-04-02 23:09:05 [2.675ms] [rows:1] SELECT * FROM "gorm_products" WHERE "gorm_products"."deleted_at" IS NULL
2023-04-02 23:09:05 2023/04/02 20:09:05.505808 database.go:73: Database migration complete.
2023-04-02 23:09:05 2023/04/02 20:09:05.508726 main.go:65: Starting server on http://localhost:8080
```

> Dikkat edilmesi gereken nokta şudur ki, uygulamanın çıktısı olan `http://localhost:8080` adresine gidildiğinde erişimin engellendiği görülecektir. Bunun sebebi servisin Docker içerisinde 8080 portunda çalışmasına rağmen, görev isterleri gereği, dışarıya başka bir porttan erişim sağlanmasıdır. Bu sebeple uygulamayı çalıştırdıktan sonra `docker ps` komutu ile çalışan containerların listesini ve servisin çalıştığı containerın portlarını görebilirsiniz. Bu portlar üzerinden uygulamaya erişim sağlayabilirsiniz.