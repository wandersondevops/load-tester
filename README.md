# Load Tester

Este projeto é um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário pode fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

## Funcionalidades

- Realiza requests HTTP para a URL especificada.
- Distribui os requests de acordo com o nível de concorrência definido.
- Garante que o número total de requests seja cumprido.
- Gera um relatório com:
  - Tempo total gasto na execução.
  - Quantidade total de requests realizados.
  - Quantidade de requests com status HTTP 200.
  - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Pré-requisitos

- Docker instalado na máquina.
- Go instalado para compilar a aplicação.

## Como Executar

### Passos para Build e Execução

1. Clone o repositório:

    ```sh
    git clone <URL_DO_REPOSITORIO>
    cd load-tester
    ```

2. Configure o módulo Go e compile a aplicação:

    ```sh
    go mod init load-tester
    go mod tidy
    go build -o load-tester
    ```

3. Construa a imagem Docker:

    ```sh
    docker build -t load-tester .
    ```

# ATENÇAO

4. Execute o serviço com o comando `docker run`, fornecendo os parâmetros necessários:

    ```sh
    docker run load-tester --url=http://google.com --requests=1000 --concurrency=10
    ```

### Parâmetros

- `--url`: URL do serviço a ser testado. (Obrigatório)
- `--requests`: Número total de requests. (Obrigatório)
- `--concurrency`: Número de chamadas simultâneas. (Obrigatório)

### Exemplo de Uso

Para testar a URL `http://example.com` com um total de 1000 requests e 10 chamadas simultâneas:

```sh
docker run load-tester --url=http://example.com --requests=1000 --concurrency=10
