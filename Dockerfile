# Używamy oficjalnego obrazu Go
FROM golang:1.23

# Ustawiamy katalog domowy aplikacji wewnątrz kontenera
WORKDIR /app

# Kopiujemy pliki aplikacji do kontenera
COPY . .

# Pobieramy zależności
RUN go mod tidy

# Kompilujemy aplikację
RUN go build -o main .

# Domyślne polecenie do uruchomienia
CMD ["./main"]