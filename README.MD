# Hubla-challenge

## Running and Using Locally
    In project root
        - $ docker-compose up -d 
            or 
            - $ go run cmd/server/main.go
        - go to: http://localhost:8000/users/login
        - choose "create user"
        - (after registering you will be redirected to the login screen)
        - log in
        - now you can see lists and upload file

## Project is also serverless
    - http://24.199.71.130//users/login
    - http://24.199.71.130/docs/index.html
    P.S: Working without any kind of authorization

## Run tests
    In project root
        $ go test ./...

## Documentation
    - http://localhost:8000/docs/index.html

## Comments
    - In relation to the front, SPA frameworks were not used, being very simple and containing some weak points of authorization is because i am just Backend developer
    - Authorization token expires in 500s.
    - Refresh token methodologies were not implemented
    - Navigation between pages does not require authorization
    - Only transaction data endpoints  require authorization
    - if the token expires, it will not be possible to view the data on the tables, so you need to log in again

## P.S.
    - Token test 
    - post to: http://localhost:8000/users
    - using: 
        {"name": "Name",
         "email": "email@email.com",
         "password": "password"
        }
    - post to: http://localhost:8000/users/generate_token
    - using:
        {"email": "email@email.com",
         "password": "password"
        }
    - return must be something like that:
        {"access_token": "eyJhbGciOiJIUz..."}



