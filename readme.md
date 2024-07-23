# My Pregnancy App

## Project Setup and Commands

### Step 1: Clone the Repository

First, clone the repository:

```bash
git clone https://github.com/29ayx/csit321-mypregnancy-app.git
cd csit321-mypregnancy-app
```
### Step 2: Install Dependencies
Ensure you have Go installed. Then, install the necessary Go packages for Fiber and Swagger:

```bash
go get github.com/gofiber/fiber/v2
go get go.mongodb.org/mongo-driver/mongo
go get github.com/gofiber/fiber/v2/middleware/logger
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/fiber-swagger


## Step 2: Project Structure

Ensure your project structure looks like this:

```go

gofiber-mongodb/
├── cmd/
│   └── main.go
├── server/
│   └── database/
│       └── mongodb.go
├── handlers/
│   └── user.go
│   └── forum.go
├── models/
│   └── user.go
│   └── forum.go
├── routes/
│   └── routes.go
├── docs/
│   └── docs.go
│   └── swagger.json
│   └── swagger.yaml

```

## Step 3: Running the Server

To run the server, execute the following command:

```bash

go run cmd/main.go

```

## Step 4: Generating and Updating Swagger Documentation

To generate or update the Swagger documentation, run:

```bash

swag init -g cmd/main.go -o docs

```

## Test API:
```bash
#Inserting
curl -X POST -H "Content-Type: application/json" -d "{\"firstname\":\"test\",\"lastname\":\"user\",\"email\":\"test@example.com\"}" http://127.0.0.1:3000/api/users

#Get - Change _id_ to user's ID
curl -X GET http://127.0.0.1:3000/api/users/_id_

#Updating - Change _id_ to user's ID
curl -X PUT -H "Content-Type: application/json" -d "{\"firstname\":\"new_test\",\"lastname\":\"example_change\" }" http://127.0.0.1:3000/api/users/update/_id_
```