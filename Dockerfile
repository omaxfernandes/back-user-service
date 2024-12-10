# Etapa 1: Construção da aplicação
FROM golang:1.20-alpine

# Definir diretório de trabalho
WORKDIR /app

# Copiar o código da aplicação para dentro do contêiner
COPY . .

# Baixar dependências do Go
RUN go mod tidy

# Compilar a aplicação
RUN go build -o back-user-service cmd/main.go

# Etapa 2: Rodar a aplicação
FROM alpine:latest

# Instalar dependências do sistema, como o pacote para rodar a aplicação
RUN apk --no-cache add ca-certificates

# Copiar a aplicação compilada da etapa de build
COPY --from=build /app/back-user-service /usr/local/bin/

# Expor a porta que o servidor vai ouvir
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["back-user-service"]
