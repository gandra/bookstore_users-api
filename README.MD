# Booking API


## Stack used: 
- https://github.com/gin-gonic/gin
- https://github.com/uber-go/zap


## Run Mysql in docker

Mariadb:
```
docker run --name bookingsusersdb -e MYSQL_ROOT_PASSWORD=mysecret -e MYSQL_USER=peralozac -e MYSQL_PASSWORD=mysecret -e MYSQL_DATABASE=bookings_users -p 3306:3306 -d docker.io/library/mariadb:10.3
docker container start bookingsusersdb
```

https://www.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go/learn/lecture/16459582#overview

