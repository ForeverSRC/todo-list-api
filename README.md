# todo-list-api
Api server of todo-list app. 

Created for practicing hexagonal architecture of golang project.

## Structure
```shell
.
├── README.md
├── cmd
│   └── todo-list-server
│       └── main.go
├── conf
│   └── default.yml
├── go.mod
├── go.sum
└── pkg
    ├── config
    │   └── config.go
    ├── http
    │   └── rest
    │       ├── handler.go
    │       └── router_api.go
    ├── model
    │   └── item.go
    ├── repository
    │   └── item.go
    ├── service
    │   └── item
    │       ├── creating
    │       │   └── service.go
    │       ├── listing
    │       │   └── service.go
    │       └── managing
    │           └── service.go
    ├── storage
    │   └── mongodb
    │       ├── init.go
    │       └── item.go
    ├── utils
    │   ├── array
    │   │   └── array.go
    │   └── math
    │       └── math.go
    └── vo
        └── item.go
```

* cmd
  * used to create an app/server
  * combine the service, storage defined in domain
* pkg
  * definite adaptors and domain objects

in `pkg`:

* model
  * define entity
* repository
  * define the repository adapter(s)
* service
  * define service 
  * divided by business logic
* storage
  * implementation of repository
  * divided by storages and entities
* vo
  * value object
  * used for transfer info
* utils
  * divided by usage
* http
  * http adapter


