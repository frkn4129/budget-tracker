# Budget Tracker Microservices

## English

### Overview
Budget-Tracker is a Domain-Driven-Design (DDD) oriented sample project that demonstrates a microservice architecture written in Go.  
The system currently consists of the following services:

| Service | Responsibility | Port | Consul Service-Name |
|---------|----------------|------|---------------------|
| discovery-service | Service discovery façade in front of Consul | 4001 | discovery-service |
| auth-service | User registration & JWT based authentication | 3000 | auth-service |
| budget-service | Expense CRUD & business logic | 3002 | budget-service |
| gateway-service | Public API gateway / reverse-proxy | 4000 | gateway-service |

Service discovery & health-checks are handled via [HashiCorp Consul](https://www.consul.io/).

### Prerequisites
* Go ≥ 1.21
* Docker (optional but recommended for PostgreSQL & Consul)
* PostgreSQL 14+
* Consul agent ( `docker run -p 8500:8500 hashicorp/consul` )

### Quick Start (Local)
```bash
# 0) clone repo & cd
export APP_ENV=dev         # enable verbose logging

# 1) start Consul (if not already running)
docker run --name consul -p 8500:8500 -d hashicorp/consul agent -dev

# 2) start PostgreSQL for auth & budget services
# (example using docker-compose)
cd docker-composes && docker compose up -d postgres

# 3) run discovery-service
cd ../discovery-service
go run ./cmd

# 4) run auth-service
cd ../auth-service
cp .env.sample .env        # edit if necessary
./init_db.sh               # create tables
go run ./cmd

# 5) run budget-service
cd ../budget-service
export DB_*                # set your DB env vars or use .env
go run .

# 6) run gateway-service
cd ../gateway-service
export JWT_SECRET=mysecret
export DISCOVERY_URL=http://localhost:4001
go run ./cmd
```
Gateway is now listening on http://localhost:4000

### Example Requests
```bash
# register & login
curl -X POST http://localhost:4000/api/auth/register -d '{"email":"a@a.com","password":"pass"}' -H 'Content-Type: application/json'

curl -X POST http://localhost:4000/api/auth/login    -d '{"email":"a@a.com","password":"pass"}' -H 'Content-Type: application/json'
# => { "token": "<JWT>" }

# create expense
curl -X POST http://localhost:4000/api/expenses \
     -H "Authorization: Bearer <JWT>" \
     -d '{"amount":100,"description":"Coffee"}' -H 'Content-Type: application/json'
```

### Environment Variables
| Variable | Default | Description |
|----------|---------|-------------|
| APP_ENV | dev | dev ⇒ verbose logs, prod ⇒ compact logs |
| GATEWAY_PORT | 4000 | Gateway exposed port |
| DISCOVERY_URL | http://localhost:4001 | URL of discovery-service |
| JWT_SECRET |   | Shared secret between services |
| (each service has its own DB_*, JWT_SECRET etc.) |

### Deployment / Production
Set `APP_ENV=prod` and use the provided docker-compose files under `docker-composes/` to bootstrap PostgreSQL, Consul and all Go services.  
Health-checks are exposed at `/health` for every service and registered in Consul.

---

## Türkçe

### Genel Bakış
Budget-Tracker, Domain-Driven-Design (DDD) yaklaşımı ile geliştirilmiş, Go dilinde yazılmış örnek bir mikroservis mimarisidir.  
Sistemde şu mikroservisler bulunur:

| Servis | Sorumluluk | Port | Consul Adı |
|--------|------------|------|------------|
| discovery-service | Consul önünde servis keşif API’si | 4001 | discovery-service |
| auth-service | Kullanıcı kayıt & JWT tabanlı kimlik doğrulama | 3000 | auth-service |
| budget-service | Harcama işlemleri & iş kuralları | 3002 | budget-service |
| gateway-service | Dış dünyaya açılan API geçidi / proxy | 4000 | gateway-service |

Servis keşfi ve sağlık kontrolleri [HashiCorp Consul](https://www.consul.io/) ile yapılmaktadır.

### Ön Koşullar
* Go ≥ 1.21
* Docker (PostgreSQL & Consul için önerilir)
* PostgreSQL 14+
* Consul agent (`docker run -p 8500:8500 hashicorp/consul`)

### Hızlı Başlangıç (Lokal)
```bash
# 0) repoyu klonla & dizine gir
export APP_ENV=dev       # ayrıntılı loglar

# 1) Consul’u başlat
docker run --name consul -p 8500:8500 -d hashicorp/consul agent -dev

# 2) PostgreSQL başlat
cd docker-composes && docker compose up -d postgres

# 3) discovery-service
cd ../discovery-service
go run ./cmd

# 4) auth-service
cd ../auth-service
cp .env.sample .env   # gerekirse düzenle
./init_db.sh          # tabloları oluştur
go run ./cmd

# 5) budget-service
cd ../budget-service
export DB_*           # env değişkenlerini ayarla veya .env kullan
go run .

# 6) gateway-service
cd ../gateway-service
export JWT_SECRET=mysecret
export DISCOVERY_URL=http://localhost:4001
go run ./cmd
```
Gateway artık http://localhost:4000 adresinde çalışıyor.

### Örnek İstekler
```bash
# kayıt & giriş
curl -X POST http://localhost:4000/api/auth/register -d '{"email":"a@a.com","password":"pass"}' -H 'Content-Type: application/json'

curl -X POST http://localhost:4000/api/auth/login -d '{"email":"a@a.com","password":"pass"}' -H 'Content-Type: application/json'
# => { "token": "<JWT>" }

# harcama oluştur
curl -X POST http://localhost:4000/api/expenses \
     -H "Authorization: Bearer <JWT>" \
     -d '{"amount":100,"description":"Kahve"}' -H 'Content-Type: application/json'
```

### Ortam Değişkenleri
| Değişken | Varsayılan | Açıklama |
|----------|-----------|----------|
| APP_ENV | dev | dev ⇒ detaylı log, prod ⇒ sade log |
| GATEWAY_PORT | 4000 | Gateway portu |
| DISCOVERY_URL | http://localhost:4001 | discovery-service adresi |
| JWT_SECRET |   | Servisler arası paylaşılan gizli anahtar |
| (Her servisin kendi DB_* ve JWT_SECRET değişkenleri vardır) |

### Prod Dağıtımı
`APP_ENV=prod` ayarlayın ve `docker-composes/` altındaki compose dosyaları ile PostgreSQL, Consul ve tüm Go servislerini ayağa kaldırın.  
Tüm servisler `/health` endpoint’i ile sağlık durumunu raporlar ve Consul’a kayıtlıdır. 