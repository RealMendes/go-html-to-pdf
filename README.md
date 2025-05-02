# go-html-to-pdf

Aplicação em Go para gerar PDFs a partir de HTML dinâmico usando o Gotenberg.

## Requisitos
- Go 1.20+
- Gotenberg rodando em http://gotenberg:3000

## Como rodar

```bash
docker-compose up --build
```

## Como usar

Faça um POST para `http://localhost:8080/gerar-pdf` com um JSON como este:

```json
{
  "id": "123",
  "data": {
    "titulo": "Meu PDF",
    "mensagem": "Este PDF foi gerado dinamicamente!",
    "itens": {
      "item1": "valor1",
      "item2": "valor2"
    }
  }
}
```

O job será colocado na fila e processado pelo worker, que irá gerar o PDF usando o Gotenberg.

## Como baixar o PDF gerado

Após o processamento, acesse:

```
GET http://localhost:8080/pdf/{id}
```

Substitua `{id}` pelo ID enviado no payload para baixar o PDF correspondente. 
