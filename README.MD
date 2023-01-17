# DashBoard Personal Finance API with Go

This is a project to try my skills in Golang

## 1- Run the project
```sh
go tidy
```

## 2- After  create in the root of project file **.env** set the environment variables
```sh
PORT=:3000
JWT_SECRET=example
DATABASE_URL=postgres://example:example@localhost:5432/example?sslmode=disable
USER_EMAIL=example@gmail.com
PASS_EMAIL=example
```
## 3- Build docker image 
```
cd database
docker build. -t name 
docker run -p 5432:5432 name
```

## 4- Run  Server 
```
go run main.go
```




## DataBase Diagram 

![admin-finanzas (1)](https://user-images.githubusercontent.com/108156933/209610220-a7ef4af6-1c49-45ce-97ca-7e6d3b78d5cd.png)


