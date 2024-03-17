# stori
System that processes transactions

install swagger cmd

- go get -u github.com/swaggo/swag/cmd/swag
- go install github.com/swaggo/swag/cmd/swag@latest
- swag init

>http://localhost:8080/api/stori/swagger/index.html

RUN PROJECT

- install docker
- change credentials of your own postgres database in postgres/postgres.go
- execute ./deploy.sh
- or go run main.go

>http://52.202.149.44/api/stori/swagger/index.html

File Template

> docs/fileTemplate.csv

POST to CLOUD AWS

``
curl --location '52.202.149.44/api/stori/transaction/upload' \
--form 'email="tivetind23@gmail.com"' \
--form 'file=@"/Users/tizon/Documents/github/stori/docs/fileTemplate.csv"'
``
