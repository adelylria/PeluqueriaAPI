# Utiliza una imagen base de Go
FROM golang:1.22.1-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /PeluqueriaAPI

# Copia los archivos de tu proyecto al directorio de trabajo
COPY . .

RUN go build -o main

EXPOSE 8080

CMD ["go", "run", "main.go"]