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

## Como Executar

### Passos para Build e Execução

1. Clone o repositório:

    ```sh
    git clone <URL_DO_REPOSITORIO>
    cd load-tester
    ```

2. Construa a imagem Docker:

    ```sh
    docker build -t load-tester .
    ```

3. Execute a aplicação:

    ```sh
    docker run load-tester --url=<URL_DO_SERVICO> --requests=<NUMERO_TOTAL_DE_REQUESTS> --concurrency=<NUMERO_DE_CHAMADAS_SIMULTANEAS>
    ```

   Exemplo:

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
```


## Saída
#### Após a execução, o sistema irá gerar um relatório no console com as seguintes informações:

- Tempo total gasto na execução.
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP.
- Contribuição
- Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Contribuição
#### Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença
- Este projeto está licenciado sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.