# Teste de carga de requisições em GO

Esse projeto implementa um mecanismo para enviar várias requisições de forma simultaneas em Go 

## Descrição

O  programa é um console application que envia várias requisições de forma simultâneas em GO utilizando de GO Rotines

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação para implementar o serviço e as Go Routines.

## **Instruções de Execução**
### **Rodando com Docker Compose**
1. Clone o repositório:
   ```sh
   git clone https://github.com/brunoofgod/goexpert-lesson-8.git
   cd goexpert-lesson-8
   ```

2. Inicie o projeto com Docker Compose:
   ```sh
   docker build -t load-tester .
   docker run --rm load-tester --url=http://google.com --requests=1000 --concurrency=10
   ```

