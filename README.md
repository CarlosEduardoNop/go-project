# Go-Project

Este projeto é destinado apenas para estudo. Projeto para estudar a linguagem e a integração do Go com RabbitMQ.

### Comandos para rodar o projeto

Para que possa rodar o projeto, será necessário subir dois containers:

RabbitMQ: `docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management`

MySQL: `docker run --name mysql-docker -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -d mysql:8`

Logo após, será necessário rodar a migration da tabela orders `go run ./cmd/migration`. No comando será criado a **database** e também a tabela de **migrations**.

### Consumidor

Depois de tudo configurado. Poderá ser utilizado o comando go run `./internal/orders_consumer`, onde irá inicializar o consumer do RabbitMQ na queue post-orders. Quando for enviado uma mensagem para essa fila contendo um json com **id**, **name** e **price**, será criado a partir do repositório uma nova linha na tabela **orders**.


##### Notas finais
Projeto criado para o aprimoramento pessoal sobre a linguagem. Tentando utilizar pacotes internos criados por mim, como a criação de migrations e o repositório.