## E-commerce made in Golang, React and Next

This is repository for simple boilerplate for the the e-commerce site made in Golang, React, Next and Stripe.

## Backend

- Built in Go, version 1.17.5
- Uses [docker images](https://docker.io)
- Uses docker compose
- Uses [Gorm](https://gorm.io/index.html)
- Uses [MySql docker image](https://www.mysql.com)
- Uses [Fiber framework](https://github.com/gofiber/fiber)
- Uses [JWT]("https://jwt.io)
- [Postman documentation](https://documenter.getpostman.com/view/11120225/UVeMH3aQ)

## Setup instruction

- Git clone this project in your repository
- It is necessary to have installed docker and docker for desktop to run docker compose
- in docker-compose.yaml file set your credentials
- cd to backend folder and run docker compose up or docker compose up -d if you want to run deamon in the background, it will spin up docker container and images
- to see loggs run docker ps first to see container_id and than docker loggs containerID

- to run functions in commands folder, to generates fake data, you have to exec backend image when docker is running and database is connected

- cd backend from root folder
- docker compose exec backend sh
- go run src/commands/folder_name/file_name

- Open postman and pick endpoint that you want to use

- If you want to stop docker images run docker composer down