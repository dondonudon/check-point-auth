Authentication & Authorization with JWT in Golang

Built with Golang, JWT, GORM, Gin Gonic, Postgres

# Migration

    go run migrate/migrate.go

# Run the server

    go run main.go

# Build the server

    go build main.go

## Register

    curl -X POST \
      http://localhost:3000/api/auth/register \
      -H 'Content-Type: application/json' \
      -d '{
           "email":"test@gmail.com",
            "name":"test",
            "password":"testtest",
            "passwordConfirm":"testtest",
            "dob":"2000-01-01T00:00:00Z",
            "age":20,
            "gender":"male",
            "occupation":"farmer"
        }'

## Verify Email

    curl -X GET \
        http://localhost:3000/api/auth/verifyemail/{token}

## Login

    curl -X POST \
        http://localhost:3000/api/auth/login \
        -H 'Content-Type: application/json' \
        -d '{
                "email":"test@gmail.com",
                "password":"testtest"
            }'

## Get Current Profile

    curl -X GET \
        http://localhost:3000/api/users/me

## Forgot Password

    curl -X POST \
        http://localhost:3000/api/auth/forgotpassword \
        -H 'Content-Type: application/json' \
        -d '{
                "email":"test@gmail.com"
            }'

## Reset Password

    curl -X PATCH \
        http://localhost:3000/api/auth/resetpassword/{token} \
        -H 'Content-Type: application/json' \
        -d '{
                "password":"testtest",
                "passwordConfirm":"testtest"
            }'

## Logout

    curl -X GET \
        http://localhost:3000/api/auth/logout

## Refresh Token

        curl -X GET \
            http://localhost:3000/api/auth/refresh
