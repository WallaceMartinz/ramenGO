# RamenGO

## Index

- [Technologies used](#technologies-used)
- [Running the application](#running-the-application)
  - [Locally](#locally)
  - [Using Docker](#using-docker)
- [Deploying](#deploying)

## Technologies Used

![Go](https://img.shields.io/badge/Go-1.18-blue)
![Gin](https://img.shields.io/badge/Gin-1.7.7-brightgreen)
![Docker](https://img.shields.io/badge/Docker-20.10.8-blue)
![Terraform](https://img.shields.io/badge/Terraform-1.2-purple)

RamenGO is a web application built with Go and the Gin framework, using Docker for containerization and Terraform for infrastructure management. The application is hosted on AWS and also on Render.

## Endpoints

-\broths

-\protein

-\order

- AWS: [https://ramengo.ddns.net](https://ramengo.ddns.net)

AWS student has a limitation and shuts down the machine where the api is hosted every 4 hours. If it's offline, use the link below. 

- Render: [https://ramengo-tu9g.onrender.com](https://ramengo-tu9g.onrender.com)


### Cloning the Repository

```sh
git clone https://github.com/WallaceMartinz/ramenGO.git
cd RamenGO
```

### Building the Go Application

```GO
go build -o ramengo main.go
```

### Running the Application

#### Locally

```GO
go run main.go
```

#### Using Docker

Pull the Docker image from Docker Hub:

```Docker
docker pull wallacemartinz/ramengo:latest
docker run -d --name ramengo-container -p 8080:8080 -e X_API_KEY=$X_API_KEY -e X_API_KEY_RV=$X_API_KEY_RV wallacemartinz/ramengo:latest
```

Running locally or with docker you need to add api-keys to your environment variables 

```sh
X_API_KEY_RV=ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf
X_API_KEY=“Here you must replace it with your own api-key”
```

### Deployment on AWS

### Cloning the Repository

```sh
git clone https://github.com/WallaceMartinz/TerraformRamenGO
```

####Deploying with Terraform

Configure your AWS credentials with aws configure and then execute:

```sh
terraform init
terraform apply
```
