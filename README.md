Implementation of a simple RESTful API with [Mux](https://github.com/gorilla/mux).

The data is a list of shopping items.

The API allows us to:

- GET the list of all shopping items
- POST a shopping item to the list
- DELETE a shopping item from the list (using its id)

POST:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"name": "toilet paper"}' http://localhost:8080/shopping-items
```

GET:

```sh
curl -X GET http://localhost:8080/shopping-items
```

DELETE:

```sh
curl -X DELETE http://localhost:8080/shopping-items/845d4a2c-67ce-42e2-bc90-1717d2bd1f10
```

