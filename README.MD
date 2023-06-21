# Hubla-challenge

## Running and Using
    In project root
        - $ docker-compose up -d 
            or 
            - $ go run cmd/server/main.go
        - go to: http://localhost:8000/users/login
        - choose "create user"
        - (after registering you will be redirected to the login screen)
        - log in
        - now you can see lists and upload file

## Run tests
    In project root
        $ go test ./...

## Documentation
    - http://localhost:8000/docs/index.html

## Comments
    - Authorization token expires in 300s.
    - Fefresh token methodologies were not implemented
    - Navigation between pages does not require authorization
    - Only transaction data endpoints  require authorization

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