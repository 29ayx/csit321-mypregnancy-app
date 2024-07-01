# Project Setup and Commands

# Step 1: Install Dependencies

Ensure you have Go installed. Then, install the necessary Go packages for Fiber and Swagger:

```bash
bashCopy code
go get github.com/gofiber/fiber/v2
go get go.mongodb.org/mongo-driver/mongo
go get github.com/gofiber/fiber/v2/middleware/logger
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/fiber-swagger

```

## Step 2: Project Structure

Ensure your project structure looks like this:

```go
goCopy code
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
bashCopy code
go run cmd/main.go

```

## Step 4: Generating and Updating Swagger Documentation

To generate or update the Swagger documentation, run:

```bash
bashCopy code
swag init -g cmd/main.go -o docs

```