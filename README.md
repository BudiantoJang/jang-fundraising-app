# Jang's Fundraising Application

This is a simple back-end system for fundraising application that allows users to create campaigns and contribute to other campaigns. To further improve the system, I have added JWT authentication and authorization to the system to make it more secure. I have also added Midtrans as a payment gateway to the system to make it more realistic and make it easier for contributors to make a payment.

## Developing
### Prerequisites
- Go 1.18
- MySQL 8.0.26
- Midtrans Account
- Direnv 

## Steps
### Creating Database
1. Create database in MySQL following this ERD:

![Jang's Fundraising Database ERD](https://drive.google.com/uc?export=view&id=15kT07qnPErJvz8tgla5Vdsa1l3S8N_lG)

2. Create a `.envrc` file in the root directory of the project and fill it with the following environment variables:

```
   DB_NAME : database name
   DB_PASSWORD : database password
   DB_PORT : database port
   DB_USER : database user 
```

3. If you want to deploy the project, put all the environtment variables into you repository or deployment variables

### Adding Midtrans Credentials
1. Create a `.envrc` file in the root directory of the project and fill it with the following environment variables:

```
   MIDTRANS_SERVER_KEY : Midtrans server key
   MIDTRANS_CLIENT_KEY : Midtrans client key
```

### Running the Application
1. Run `direnv allow` to load the environment variables
2. Run `go run main.go` or `make run` to run the application


## API Documentation

The API documentation for this project can be accessed from : [POSTMAN COLLECTION](https://documenter.getpostman.com/view/15849991/TzXzDq8o)

## Future Update

- [ ] Add unit test
- [ ] Add integration test
- [ ] Add more features
- [ ] Created Swagger documentation using swaggo package : [swaggo](https://github.com/swaggo/swag)
- [ ] Creating frontend to consume the API and build full-stack application