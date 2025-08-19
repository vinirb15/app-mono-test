# Posts Platform

Este repositório contém um conjunto de microserviços que compõem a plataforma **Posts**, incluindo autenticação, gerenciamento de usuários, postagens e um **BFF** (Backend for Frontend).

## 📦 Serviços

- **bff**  
  Backend for Frontend que centraliza a comunicação entre o frontend e os demais serviços.  
  - Porta padrão: `8080`

- **auth-service**  
  Responsável pela autenticação, emissão e renovação de tokens (JWT + Refresh Token).  
  - Porta padrão: `8083`

- **profile-service**  
  Gerencia perfis de usuários (registro, dados de perfil, seguidores).  
  - Porta padrão: `8082`

- **post-service**  
  Gerencia postagens, comentários e interações.  
  - Porta padrão: `8081`

## 🗂 Estrutura do Repositório

```
.
├── bff/
├── auth-service/
├── profile-service/
├── post-service/
├── docker-compose.yml
└── README.md
```

## ⚙️ Configuração

Cada serviço utiliza variáveis de ambiente. As principais estão listadas abaixo:

### Variáveis comuns
- `JWT_SECRET` — chave usada para assinar os tokens JWT.  
- `DATABASE_URL` — URL de conexão com o PostgreSQL.  
- `ADDR` — porta em que o serviço irá escutar.  
- `ACCESS_TTL` — tempo de expiração do token de acesso (ex: `15m`).  
- `REFRESH_TTL` — tempo de expiração do refresh token (ex: `168h`).  

### Exemplo para `auth-service`

```env
JWT_SECRET=secret_key
DATABASE_URL=postgres://postgres:postgres@db:5432/dialog?sslmode=disable
ADDR=:8083
ACCESS_TTL=15m
REFRESH_TTL=168h
```

## ▶️ Como rodar

### Usando Docker Compose

1. Suba os serviços:
   ```bash
   docker compose up --build
   ```

2. Os serviços estarão disponíveis em:
   - BFF: http://localhost:8080
   - Auth Service: http://localhost:8083
   - Profile Service: http://localhost:8082
   - Post Service: http://localhost:8081

3. O banco de dados PostgreSQL ficará disponível em:
   - Host: `localhost`
   - Porta: `5432`
   - Usuário: `postgres`
   - Senha: `postgres`
   - Banco: `dialog`

### Rodando localmente

Se quiser rodar apenas um serviço:

```bash
cd auth-service
go run main.go
```

⚠️ Certifique-se de ter um banco PostgreSQL rodando e as variáveis de ambiente configuradas.

## 📖 Documentação da API

Cada serviço expõe sua própria documentação via Swagger em `/swagger/index.html`.

Exemplo:
- `http://localhost:8083/swagger/index.html` → Auth Service
- `http://localhost:8082/swagger/index.html` → Profile Service
- `http://localhost:8081/swagger/index.html` → Post Service

## 🔐 Autenticação

Os endpoints protegidos exigem o header:

```
Authorization: Bearer <token>
```

Tokens podem ser obtidos via `auth-service` (`/login` ou `/signup`).

## ✅ Próximos Passos

- [ ] Adicionar testes automatizados
- [ ] Configurar CI/CD
- [ ] Melhorar documentação de rotas
- [ ] Adicionar exemplos de requests com `curl` ou `httpie`
- [ ] Implementar cache para melhorar performance
- [ ] Implementação de WebSockets para notificações em tempo real
- [ ] Melhorar a segurança com rate limiting e CORS
- [ ] Melhoria de documentações

---

Feito com ❤️ usando Go, Gin e PostgreSQL.
