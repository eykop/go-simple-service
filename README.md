# go-simple-service

A simple microservice with Golang, Mainly done while learning Golang

- [x] CRUD APIs
  - [x] List
  - [x] Get
  - [x] Update
  - [x] Create
  - [x] Delete

# TODO - Order is not pririty:

- [ ] Connect to DB (PostgreSQL).
- [ ] Dockerfile and Docker compose.
- [x] Better setup repository and improve documenation.
- [ ] Add more unit tests.
  - [ ] utils
  - [x] create product
  - [ ] list products
  - [x] get product
  - [x] delete product
  - [ ] update product
- [ ] Add some authentication (proably some JWT).
- [ ] Add some permissions for calling APIs.
- [x] Refactor code and make more order.

# Swagger

In order to view the swagger docs run the service and then browse the url paths of:

- `http://localhost:3000/docs/`
- `http://localhost:3000/redocs`

In order for the swagger to work as expected, several step need to be followed:

We need to install some swagger helper utilities:

```shell
go get -u github.com/go-swagger/go-swagger/cmd/swagger
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
```

Make sure the installed go modules are in your PATH, for example:

`export PATH=~/go/bin/:${PATH}`

Then we need to generate the swagger documentations:

`swag init`

This will generate the following files:

```shell
├── docs --> the swagger docs folder ( already found in repo)
│ ├── docs.go --> the swagger docs module ( already found in repo), this will help http serve the docs
│ ├── swagger.json --> the swagger specification json file ( already found in repo)
│ └── swagger.yaml --> the swagger specification yaml file ( already found in repo)
```

Finally for running the endpoint calls with a cli client, we can use swagger to auto generated a go client:

`swagger generate client -f docs/swagger.yaml -t client`

This will generate the following client package and modules,
Please note this is not and should not be part of the repository as this is an auto generated code that should not
be maintained nor modified by us.

```shell
├── client
│ ├── client
│ │ ├── products
│ │ │ ├── delete_products_id_parameters.go
│ │ │ ├── delete_products_id_responses.go
│ │ │ ├── get_products_id_parameters.go
│ │ │ ├── get_products_id_responses.go
│ │ │ ├── get_products_parameters.go
│ │ │ ├── get_products_responses.go
│ │ │ ├── patch_products_id_parameters.go
│ │ │ ├── patch_products_id_responses.go
│ │ │ ├── post_products_parameters.go
│ │ │ ├── post_products_responses.go
│ │ │ └── products_client.go
│ │ └── swagger_products_api_client.go
│ └── models
│ ├── data_product.go
│ └── handlers_http_error.go
```

Special thanks to [Nic Jackson](https://www.youtube.com/c/NicJackson) and for his greate tutorial on Go and Microservices.
