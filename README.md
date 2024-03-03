# **Rotating Pairs Back-End**

![GitHub repo size](https://img.shields.io/github/repo-size/glener10/rotating-pairs-back)
![Bitbucket open issues](https://img.shields.io/bitbucket/issues/glener10/rotating-pairs-back)
![Bitbucket open pull requests](https://img.shields.io/bitbucket/pr-raw/glener10/rotating-pairs-back)
![Contributors](https://img.shields.io/github/contributors/glener10/rotating-pairs-back.svg)
![Build Status](https://github.com/glener10/rotating-pairs-back/workflows/go/badge.svg)
[![License](https://img.shields.io/github/license/glener10/rotating-pairs-back)](/LICENSE)

<p align="center"> 🚀 Back-End of Rotating Pairs Web Application. Responsible for saving monitoring logs and
Returning the positions of combinations used in the application (Limited beetwen 2 and 20), like this. 👇🏽 </p>

![Return Example](/documentation/readmeImages/return.png)

🏁 Table of Contents

===================

<!--ts-->

👉 [Dependencies and Environment](#dependenciesandenvironment)

👉 [Installing](#installing)

👉 [Testing](#testing)

👉 [Using](#using)

👉 [Learn More](#learnmore)

👉 [License](#license)

👉 [Author](#author)

<!--te-->

===================

The production application is in [rotating-pairs-back-production](https://rotating-pairs-back-production.up.railway.app/), the only route that is available to any origin is '/', where you will see a Hello World! The other routes are private to the Front-End application, you can [exec locally](#using) the back-end with some limited values.

You can access the production Front-End application via the link [Rotating Pair](https://rotatingpairs.online)
The front-end source code is in [hear](https://github.com/glener10/rotating-pairs-front).

<div id="dependenciesandenvironment"></div>

## 💻 **Dependencies and Environment**

My dependencies and versions

**Go**: go version go1.22.0 windows/amd64

If you will exec in local environment, you will need the docker and docker-compose to create a MongoDb Instance:

**Docker**: Docker version 25.0.3, build 4debf41

**docker-compose**: Docker Compose version v2.24.5-desktop.1

<div id="installing"></div>

## 🚀 **Installing**

You will need a MongoDB instance, we have a ready image and docker-compose configuration in _data_ folder, you can exec duying this:

```bash
#Go to 'data' folder
cd data

# container with MongoDB instance
$ docker-compose up -d
```

This execution uploads a MongoDB container with some initialized limited values, you can see them in [init-mongo.js](./data/init-mongo.js)

You will need to a _.env_ file in root folder

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE_NAME=test
ALLOW_URLS=*
ENV=development
SECRET=banana
```

<div id="testing"></div>

## 🧪 **Testing**

**OBSERVATION**: To run the tests, [MongoDB](#installing) must be running

A truncate script is executed each time the tests are run and the database is reset to an initial state with limited data values. That is, if you are working and running the tests, the data will be lost.

```
go test -p 1 ./src/...
```

<div id="using"></div>

## ☕ **Using**

First, check the [dependencies](#dependenciesandenvironment) and the [installation](#installing) process:

Going to root folder and exec:

```
go run .\main.go
```

Now you can open [http://localhost:8080](http://localhost:8080) with your browser to see the result.

## 📖 **Learn More**

To learn more about technologies used in the application:

- [Go](https://golang.org/) - learn about Go features and API.

- [Railway](https://railway.app/) - learn about Railway features and API.

- [Docker](https://www.docker.com/) - learn about Docker features and API.

- [Docker Compose](https://docs.docker.com/compose/) - learn about Docker Compose features and API.

<div id="license"></div>

## 🔒 **License**

Projeto contêm [GNU GENERAL PUBLIC LICENSE](LICENSE).

<div id="author"></div>

#### **👷 Author**

Made by Glener Pizzolato! 🙋

[![Linkedin Badge](https://img.shields.io/badge/-Glener-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/glener-pizzolato/)](https://www.linkedin.com/in/glener-pizzolato-6319821b0/)
[![Gmail Badge](https://img.shields.io/badge/-glenerpizzolato@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:glenerpizzolato@gmail.com)](mailto:glenerpizzolato@gmail.com)
