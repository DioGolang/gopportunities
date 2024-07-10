# Gopportunities

Gopportunities é um projeto de API para gerenciamento de oportunidades de emprego, desenvolvido em Go (Golang), utilizando PostgreSQL como banco de dados.

## Requisitos

Antes de começar, verifique se você tem os seguintes requisitos instalados em seu ambiente de desenvolvimento:

- Go (Golang)
- Docker
- Git

## Clonando o Repositório

Clone o repositório do GitHub para seu ambiente local:

```bash
git clone https://github.com/seu-usuario/gopportunities.git
cd gopportunities
```
---

## Configurando o Banco de Dados
O projeto utiliza PostgreSQL como banco de dados. Certifique-se de ter o Docker instalado e em execução. A configuração do banco de dados pode ser feita utilizando Docker Compose:

```
# No diretório raiz do projeto
docker-compose up -d

```
Isso iniciará um container Docker com PostgreSQL configurado de acordo com o arquivo docker-compose.yml fornecido.

---

## Configuração do Ambiente


Crie um arquivo .env na raiz do projeto com as seguintes variáveis de ambiente:



```
POSTGRES_HOST=db
DB_PORT=5432
POSTGRES_USER=yourusername
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=gopportunities

```
Substitua yourusername e yourpassword pelos seus próprios valores de usuário e senha do PostgreSQL.

---

## Instalando Dependências
Certifique-se de que todas as dependências do projeto estejam instaladas. Use o go mod para isso:

```
go mod tidy

```

## Executando o Projeto
Com todas as configurações e dependências prontas, você pode executar o projeto:

```
go run cmd/api/main.go

```
Isso iniciará a API na porta padrão 8080.

---

## Testando a API
Você pode testar a API usando ferramentas como cURL ou Postman. Aqui está um exemplo de como criar um novo job:

```
POST http://localhost:8080/jobs
Content-Type: application/json

{
  "title": "Desenvolvedor Full Stack",
  "role": "Desenvolvedor",
  "description": "Descrição da vaga...",
  "requirements_qualifications": ["Experiência com Go e PostgreSQL"],
  "working_model": "remote",
  "process_steps": ["Entrevista técnica", "Avaliação prática"],
  "company": "Minha Empresa",
  "location": "Remoto",
  "salary": 8000,
  "link": "https://minhaempresa.com/vaga",
  "expiration_date": "2024-12-31"
}

```

## Contribuindo
Sinta-se à vontade para contribuir com melhorias no projeto. Abra uma issue para discutir novos recursos ou problemas encontrados.





