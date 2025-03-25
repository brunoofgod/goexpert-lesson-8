

# Etapa de construção
FROM golang:1.23.4 AS builder

# Criando diretório de trabalho
WORKDIR /app

# Copiando os arquivos do projeto
COPY go.mod ./
RUN go mod download

# Copia todo o código-fonte para dentro do container
COPY . .

# Compila o executável de forma estática para evitar dependências no Alpine
RUN CGO_ENABLED=0 go build -o /load-tester ./cmd/program

# Criando uma imagem mínima baseada em Alpine
FROM alpine:latest

# Definindo o diretório de trabalho
WORKDIR /root/

# Copiando o executável para a imagem final
COPY --from=builder /load-tester .

# Garantindo que o executável tenha permissões corretas
RUN chmod +x /root/load-tester

# Definindo a entrada corretamente
ENTRYPOINT ["/root/load-tester"]
CMD []