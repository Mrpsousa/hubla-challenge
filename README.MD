# Hubla-challenge

## Running and Using
    In project root
        - $ docker-compose up -d 
            or 
            - $ go run cmd/server/main.go
        - go to: http://24.199.71.130/users/login
        - choose "create user"
        - (after registering you will be redirected to the login screen)
        - log in
        - now you can see lists and upload file

## Run tests
    In project root
        $ go test ./...

## Documentation
    - http://24.199.71.130/docs/index.html

## Comments
    - Authorization token expires in 500s.
    - Fefresh token methodologies were not implemented
    - Navigation between pages does not require authorization
    - Only transaction data endpoints  require authorization
    - if the token expires, it will not be possible to view the data on the tables, so you need to log in again

## P.S.
    - Token test 
    - post to: http://24.199.71.130/users
    - using: 
        {"name": "Name",
         "email": "email@email.com",
         "password": "password"
        }
    - post to: http://24.199.71.130/users/generate_token
    - using:
        {"email": "email@email.com",
         "password": "password"
        }
    - return must be something like that:
        {"access_token": "eyJhbGciOiJIUz..."}