# Go mongo jwt auth

This project is a simple REST API that demonstrates how to implement user authentication using Go, MongoDB, and JSON Web Tokens (JWTs).

## Setting up MongoDB using Docker

1. Install Docker on your machine by following the instructions on the [Docker website](https://docs.docker.com/get-docker/).

2. Pull the MongoDB Docker image by running the following command in your terminal:

   ```
   docker pull mongo
   ```

3. Run the MongoDB Docker container by running the following command in your terminal:

   ```
   docker run --name mongo -p 27018:27017 -d mongo
   ```

   This command starts a new Docker container named `mongo` and maps port `27018` on your machine to port `27017` in the container. The `-d` flag runs the container in detached mode, which means that it runs in the background.

4. Verify that the MongoDB container is running by running the following command in your terminal:

   ```
   docker ps
   ```

   This command lists all running Docker containers on your machine. You should see the `mongo` container in the list.

## Running the project

1. Clone the project repository to your machine by running the following command in your terminal:

   ```
   git clone https://github.com/zvdy/go-mongo-jwt-auth.git
   ```

2. Navigate to the project directory by running the following command in your terminal:

   ```
   cd go-mongo-jwt-auth
   ```

3. Run the `generate.go` file to generate sample user data by running the following command in your terminal:

   ```
   go run database/generate.go
   ```

   This command generates sample user data and inserts it into the MongoDB database.

4. Run the `main.go` file to start the server by running the following command in your terminal:

   ```
   go run main.go
   ```

   This command starts the server and listens for incoming requests on port `8080`.

5. Test the endpoint routes of the app using `curl` commands or a web browser.

## Testing the endpoints using `curl`

### /adduser

```bash
curl -X POST \
  http://localhost:8080/adduser \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "jane",
    "password": "s3cr3t"
  }'
```

### /login

```bash
curl -X POST \
  http://localhost:8080/login \    
  -H 'Content-Type: application/json' \  -d '{
    "username": "john",
    "password": "p4ssw0rd"
  }'  
```

### /protected

```bash
curl -X GET \ 
  http://localhost:8080/protected \
  -b 'token=$TOKEN'
```




