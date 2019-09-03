# rest-api

Este projeto é uma API Exchange de Bitcoins para o teste técnico da empresa Red Ventures.
A seguir veja as dependências deste projeto e o passo a passo de como configurá-lo.

## Dependências:

* Golang 1.12.9

* Docker 18.09.7

## Configurações

Faça o clone do projeto:  
```git clone https://github.com/moromimay/rest-api.git```

Crie a imagem do banco de dados: (O arquivo Dockerfile do banco está na pasta /docker/mariadb/)  
``` docker build -t maridb-docker . ```

Execute o docker compose:  
``` docker-compose up -d```
