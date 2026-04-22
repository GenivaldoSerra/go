# Projeto de Cursos em Go

Um repositório educacional para aprender os fundamentos da linguagem Go, com exemplos práticos e exercícios.

## 📁 Estrutura do Projeto

```
.
├── hello/              # Módulo com exemplos básicos de Go
│   ├── main.go        # Programa Hello World com argumentos
│   ├── array.go       # Exemplos de manipulação de arrays
│   ├── maps.go        # Exemplos de uso de maps (dicionários)
│   ├── ponteiros.go   # Conceitos de ponteiros
│   ├── ponterios-funcao.go  # Ponteiros em funções
│   └── go.mod         # Arquivo de módulo Go
│
├── httpserver/         # Servidor HTTP com autenticação
│   ├── main.go        # Servidor HTTP com Basic Auth
│   └── go.mod         # Arquivo de módulo Go
│
└── README.md          # Este arquivo
```

## 🚀 Módulos

### Hello - Fundamentos de Go

Um módulo introdutório que cobre os conceitos básicos da linguagem Go:

#### **main.go** - Hello World
Exemplo simples de um programa que recebe um argumento de linha de comando e exibe uma saudação.

```bash
cd hello
go run main.go "Seu Nome"
# Saída: Hello, Seu Nome!
```

#### **array.go** - Manipulação de Arrays
Demonstra operações comuns com arrays em Go:
- Criação e inicialização de arrays
- Adicionar elementos com `append()`
- Remover elementos
- Acessar elementos por índice

#### **maps.go** - Dicionários
Exemplos de uso de maps (estruturas chave-valor):
- Criação e inicialização de maps
- Iteração sobre elementos com `range`

#### **ponteiros.go** - Conceitos de Ponteiros
Introdução a ponteiros em Go:
- Declaração e uso de ponteiros
- Desreferenciação
- Modificação de valores através de ponteiros

#### **ponterios-funcao.go** - Ponteiros em Funções
Exemplo de como usar ponteiros como parâmetros de funções para modificar valores.

### HTTP Server - Servidor com Autenticação

Um servidor HTTP com autenticação básica (Basic Auth).

#### **main.go** - Servidor HTTP Autenticado

Servidor de arquivos estáticos com proteção por autenticação básica.

**Credenciais padrão:**
- Usuário: `admin`
- Senha: `admin123`

**Como executar:**

```bash
cd httpserver
go run main.go <diretorio> <porta>
```

**Exemplo:**

```bash
go run main.go /tmp 8080
```

Depois acesse `http://localhost:8080` no navegador e entre com as credenciais.

## 🛠️ Pré-requisitos

- **Go 1.16+** instalado em sua máquina
- Terminal/Console para executar os comandos

### Instalação do Go

Se você ainda não tem Go instalado, baixe em: [golang.org](https://golang.org/dl/)

## 📦 Dependências

O projeto usa as seguintes dependências externas:

### httpserver
- `github.com/abbot/go-http-auth` - Autenticação HTTP básica

Para instalar as dependências:

```bash
go mod download
```

## 🎯 Como Usar

### Executar Exemplos do Módulo Hello

```bash
# Navegar para o diretório
cd hello

# Executar o Hello World
go run main.go "Seu Nome"

# Executar exemplos de arrays
go run array.go

# Executar exemplos de maps
go run maps.go

# Executar exemplos de ponteiros
go run ponteiros.go
go run ponterios-funcao.go
```

### Compilar e Executar

```bash
# Compilar o programa
go build main.go

# Executar o binário compilado
./main "Seu Nome"
```

### Executar Testes

```bash
# Executar testes (se disponíveis)
go test ./...
```

## 💡 Conceitos Aprendidos

- ✅ Variáveis e tipos de dados
- ✅ Arrays e slices
- ✅ Maps (dicionários)
- ✅ Ponteiros
- ✅ Funções
- ✅ Packages e módulos
- ✅ HTTP server
- ✅ Autenticação básica

## 📚 Recursos Adicionais

- [Documentação Oficial do Go](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## 📝 Licença

Este projeto é fornecido como material educacional.

## 👤 Autor

Projeto de aprendizado em Go.

---

**Última atualização:** Abril de 2026
