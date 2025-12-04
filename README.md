# Desafio Pós Go Expert Labs Google Cloud Run

## Sumário

- [Versões das ferramentas](#versões-das-ferramentas)
- [APIs utilizadas](#apis-utilizadas)
- [Link da API no Google Cloud Run](#link-da-api-no-google-cloud-run)
- [Variáveis de ambiente](#variáveis-de-ambiente)
- [Executando o projeto](#executando-o-projeto)
- [Executando o teste](#executando-o-teste)
- [Requisitos](#requisitos)

### Versões das ferramentas

[[Sumário](#sumário)]

- Go: 1.24.3
- Docker: 29.0.3, build 511dad6
- Docker Compose: 2.40.3

### APIs utilizadas

[[Sumário](#sumário)]

- CEP: https://brasilapi.com.br/
- Clima: https://www.weatherapi.com/

### Link da API no Google Cloud Run

[[Sumário](#sumário)]

:warning: Pode estar indisponível em um momento no futuro :warning:

- https://cloud-run-goexpert-748393798382.southamerica-east1.run.app

### Variáveis de ambiente

[[Sumário](#sumário)]

| Nome            | Descrição                           |
| --------------- | ----------------------------------- |
| WEB_SERVER_PORT | A porta que o webserver vai atender |
| WEATHER_API_KEY | Chave de API da API de clima        |

### Executando o projeto

[[Sumário](#sumário)]

Os comandos utilizados para operar parte do projeto estão presentes no arquivo **Makefile** na raiz do projeto. Isso foi feito para facilitar a execução de comandos extensos utilizando o **Make**. Caso não possua o **Make**, basta copiar o comando da **recipe**.

:warning: Se o comando **docker compose** não for reconhecido, substitua no arquivo **Makefile** por **docker-compose**. :warning:

Os comandos são:

```bash
# Fazer o build da imagem do webserver
make build
```

```bash
# Criar e rodar o container do webserver
make up
```

```bash
# Parar e remover o container do webserver
make down
```

```bash
# Rodar os container do webserver
make start
```

```bash
# Parar os container do webserver
make stop
```

```bash
# Imprime os logs do container do webserver
make logs
```

Para executar o projeto é necessário:

1. Realizar o build da imagem do webserver utilizando o comando **make build**.
2. Criar na raiz do projeto o arquivo **.env** com as variáveis de ambiente preenchidas.
3. Executar o comando **make up**.
4. O webserver estará de pé e pode começar a receber as requisições.
5. Para ver os logs do webserver, basta executar o comando **make logs**.

O webserver possui apenas uma rota GET em / e o CEP é informado como parâmetro da request:

Exemplo: http://webserver/11222000

### Executando o teste

[[Sumário](#sumário)]

:construction: Sem tempo, posso adicionar no futuro :construction:

### Requisitos

[[Sumário](#sumário)]

- [x] O sistema deve receber um CEP válido de 8 digitos;
- [x] O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin;
- [x] O sistema deve responder adequadamente nos seguintes cenários:
  - [x] Em caso de sucesso: Código HTTP 200 e Response Body { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 };
  - [x] Em caso de falha, caso o CEP não seja válido (com formato correto): Código HTTP 422 e Mensagem invalid zipcode;
  - [x] ​​​Em caso de falha, caso o CEP não seja encontrado: Código HTTP 404 e Mensagem can not find zipcode;
- [x] Deverá ser realizado o deploy no Google Cloud Run.
