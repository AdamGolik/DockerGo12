# **Go Application with PostgreSQL and Docker**

## **Opis projektu**
Prosta aplikacja stworzona w Go, która współpracuje z bazą danych PostgreSQL. Projekt wykorzystuje Dockera do zarządzania bazą danych oraz aplikacją. W trakcie programowania możesz również używać narzędzia `Air`, które automatycznie przeładowuje aplikację po zmianach w kodzie.

## **Wymagania wstępne**
Przed rozpoczęciem pracy z projektem upewnij się, że masz:
- [Go 1.23]() (lub nowszą wersję)
- [Docker]() i [Docker Compose]()
- Narzędzie [Air]() (opcjonalne, do lokalnego uruchamiania w czasie developmentu)

## **Zawartość projektu**
- **Backend w Go**: Aplikacja zarządzająca użytkownikami.
- **PostgreSQL**: Baza danych używana do zarządzania i przechowywania danych aplikacji.
- **Docker**: Używane do konteneryzacji aplikacji i bazy danych.
- **Air**: Narzędzie do automatycznego restartowania aplikacji podczas developmentu.

## **Konfiguracja**

### **1. Plik `.env`**
Upewnij się, że w katalogu głównym projektu znajduje się plik `.env` z poniższą zawartością:

```plaintext
# JWT token secret key
SECRET_KEY=qzawxsecdrvftbgyhnujimko,1234567890-=

# Database is PostgreSQL
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=adamgolik113
DB_NAME=go2
```

### **2. Plik `docker-compose.yml`**
`docker-compose.yml` zawiera konfigurację dla zarówno aplikacji, jak i bazy danych PostgreSQL. Plik wymaga poniższej konfiguracji:

```yaml
version: '3.8'

services:
  # Baza danych PostgreSQL
  postgres:
    image: postgres:15
    container_name: postgres_container
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: adamgolik113
      POSTGRES_DB: go2
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - app_network

  # Aplikacja Go
  app:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: go_app_container
    restart: always
    ports:
      - "8080:8080"
    environment:
      DB_HOST: postgres
      DB_PORT: 5432
      DB_USER: postgres
      DB_PASSWORD: adamgolik113
      DB_NAME: go2
      SECRET_KEY: qzawxsecdrvftbgyhnujimko,1234567890-=
    depends_on:
      - postgres
    networks:
      - app_network

volumes:
  postgres_data:

networks:
  app_network:
    driver: bridge
```

### **3. Plik `air.toml`**
Plik `air.toml` może być używany podczas developmentu z narzędziem Air:

```toml
# Config file for Air (https://github.com/cosmtrek/air)

[build]
cmd = "go build -o ./tmp/main"
bin = "./tmp/main"
run = "./tmp/main"

[watch]
dirs = ["."]
exclude_dirs = ["tmp", "vendor", "node_modules"]
extensions = [".go", ".tpl", ".html"]

[log]
level = "debug"
```

### **4. Plik `Dockerfile`**
Plik `Dockerfile` opisuje, jak budowany jest obraz Dockera dla aplikacji Go:

```Dockerfile
# Używamy oficjalnego obrazu Go
FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
```

---

## **Uruchamianie aplikacji**

### **1. Uruchomienie z Dockerem**
Aby uruchomić aplikację i bazę danych PostgreSQL w kontenerach Dockera:

1. Zbuduj obrazy i uruchom kontenery:

   ```bash
   docker-compose up --build
   ```

2. Po uruchomieniu:
   - PostgreSQL jest dostępny na porcie **5432**.
   - Aplikacja Go jest dostępna na porcie **8080**:  
     ```
     http://localhost:8080
     ```

3. Aby zatrzymać kontenery:

   ```bash
   docker-compose down
   ```
   ---

### **2. Uruchomienie w trybie developmentu z Air**

Jeśli chcesz szybko modyfikować i testować aplikację lokalnie:

1. Uruchom PostgreSQL w kontenerze Dockera:
   ```bash
   docker-compose up postgres
   ```

2. Zainstaluj `Air` (jeśli jeszcze tego nie masz):
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

3. Uruchom aplikację za pomocą `Air`:
   ```bash
   air
   ```

---

## **API**

### **Podstawowe endpointy**

1. **Dodanie nowego użytkownika**
   - **POST** `/user`

   Body (JSON):
   ```json
   {
     "name": "John",
     "last_name": "Doe",
     "age": 25,
     "email": "john.doe@example.com",
     "username": "johndoe",
     "password": "password"
   }
   ```

2. **Pobranie listy użytkowników**:
   - **GET** `/users`

3. **Aktualizacja użytkownika po ID**:
   - **PUT** `/user/{id}`

   Body (JSON):
   ```json
   {
     "name": "Jane",
     "last_name": "Smith",
     "age": 30,
     "email": "jane.smith@example.com",
     "username": "janesmith",
     "password": "newpassword"
   }
   ```

4. **Usunięcie użytkownika po ID**:
   - **DELETE** `/user/{id}`

---

## **Baza danych**

- **Typ**: PostgreSQL
- **Domyślne połączenie**:
  - **Host**: `postgres`
  - **Port**: `5432`
  - **Użytkownik**: `postgres`
  - **Hasło**: `adamgolik113`
  - **Nazwa bazy danych**: `go2`

---

## **Uwagi**  
- Aplikacja używa Gorm do komunikacji z bazą danych.
- Narzędzie Air automatycznie przeładowuje aplikację podczas developmentu.
- PostgreSQL jest dostępny w kontenerze Dockera. Upewnij się, że plik `.env` zawiera poprawne dane konfiguracji.

---

## **Autor**
- Autor `adam Golik`
Utworzono z pomocą. 😊
