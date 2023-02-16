# usertTest
User micro service in go with a REST API

## Requirements
- Docker
- Docker-Compose
- An internet connection

## Composition
- Main userTest microservice
- Postgres server
- RabbitMQ queue system
- Consumer app to get and print the Domain Events that userTest throws

## Execution
### Run tests
Execute this commands in the main folder:
```bash
go mod vendor
go test ./...
```

### Build
Execute :
```bash
docker-compose -f Docker-compose.yaml build
```
### Up and Run
```
docker-compose -f Docker-compose.yaml up -d
```

## Using it
The microservice has 5 endpoints and the http server is listen the port 8088

### HeathCheck
GET - localhost:8088/health_check
This can be used in Kubernetes to check the availability of the service.
It returns 200 as status if the service is alive.

---
### Create an User
POST - localhost:8088/api/v1/user
The payload is a JSON object like this:
```JSON
{
    "first_name": "Name",
    "last_name" : "LastName",
    "nickname": "test",
    "password": "password",
    "email": "test@test.com",
    "country": "ES"
}
```
---
### Update an User
PUT - localhost:8088/api/v1/user/{user uuid}
The payload is the same:
```JSON
{
    "first_name": "NameUpdated",
    "last_name" : "LastNameUpdated",
    "nickname": "testUpdated",
    "password": "passwordUpdated",
    "email": "test@testUpdated.com",
    "country": "FR"
}
```
---
### Delete an User
DELETE - localhost:8088/api/v1/user/{user uuid}

---
### List users
This enpoint return a list of users, it uses filters and pagination

GET - localhost:8088/api/v1/user?p=0&ps=10&first_name="xx"....
- p is the page do you want to get
- ps is the page size, the default is 10 items per page.
- first_name, last_name, nickname, email, country are the filters allowed in this endpoint.
---
## Note:
In the 'postman' folder are all the possible calls.

### Curls
List users
```bash
curl --location --request GET 'localhost:8088/api/v1/user?ps=2&p=0&first_name=NameUpdated&nickname=testUpdated&country=FR'
```

HealthCheck
```bash
curl --location --request GET 'localhost:8088/health_check'
```

Create User
```bash
curl --location --request POST 'localhost:8088/api/v1/user' 
--header 'Content-Type: application/json' 
--data-raw '{
    "first_name": "Name",
    "last_name" : "LastName",
    "nickname": "test",
    "password": "password",
    "email": "test@test.com",
    "country": "ES"
}'
```

Update User
```bash
curl --location --request PUT 'localhost:8088/api/v1/user/9610093d-bc7a-4ab5-bbd6-55f97df3df8f' 
--header 'Content-Type: application/json' 
--data-raw '{
    "first_name": "NameUpdated",
    "last_name" : "LastNameUpdated",
    "nickname": "testUpdated",
    "password": "passwordUpdated",
    "email": "test@testUpdated.com",
    "country": "FR"
}'
```

Delete User
```bash
curl --location --request DELETE 'localhost:8088/api/v1/user/6690366c-4d16-4569-82ca-b9a436bef152'
```

## RabbitMQ
In localhost:15672 you can access to the RabbitMQ UI, the user is `guest` and password `guest` if you want to monitor the queue status

## Check the event consumer
If you want to see the events consumed you can execute in the terminal :
```
docker-compose -f Docker-compose.yaml logs -f consumer
```

If you prefer watch all the logs you can run:
```
docker-compose -f Docker-compose.yaml logs -f
```
