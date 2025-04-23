# 🏷️ Leilão com Concorrência em Go

Este é um projeto de um sistema de leilão com fechamento automático, desenvolvido em Go (Golang). Ele utiliza o MongoDB como banco de dados e foca em conceitos de concorrência e testes automatizados.

## 📦 Funcionalidades

- Cadastro de leilões (Auction)
- Fechamento automático de leilões ao fim do tempo de expiração
- Persistência com MongoDB
- Testes de integração com validação do comportamento assíncrono
- Arquitetura limpa e separação entre domínio e infraestrutura

## 🛠️ Tecnologias Utilizadas

- **Go 1.21+**
- **MongoDB**
- **Mongo Go Driver**
- **Testify** (assert) para testes
- **Context API** para controle de concorrência
- **Docker** (opcional)

## 🚀 Como rodar o projeto

### 1. Clone o repositório
```bash
git clone https://github.com/vitorlrrcamargo/seu-repo.git
cd seu-repo
```

### 2. Suba o MongoDB com Docker (se necessário)
Se você não possui o MongoDB rodando localmente, pode subir o container com o seguinte comando:
```bash
docker run --name mongo -p 27017:27017 -d mongo:latest
```

### 3. Execute os testes
Para rodar os testes automatizados, utilize o comando:
```bash
go test ./...
```

## 🧪 Teste de Fechamento Automático

Este projeto inclui um teste de integração que valida se o leilão é fechado automaticamente após o tempo limite. O arquivo de teste está localizado em:
```text
tests/integration/auction_auto_close_test.go
```
⚠️ **Certifique-se de que o MongoDB está rodando corretamente para que o teste funcione.**

## 📁 Estrutura de Pastas
```text
.
├── internal/
│   ├── domain/         # Lógica de negócio (entidades, interfaces)
│   └── infra/          # Conexões com MongoDB e repositórios
├── tests/
│   └── integration/    # Testes de integração
├── go.mod / go.sum
└── README.md
```

## 📌 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues, sugerir melhorias ou criar pull requests.

## 👨‍💻 Autor

Feito com 💛 por [vitorlrrcamargo](https://github.com/vitorlrrcamargo)
