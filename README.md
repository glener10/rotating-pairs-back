# **Rotating Pairs Back-End**

![GitHub repo size](https://img.shields.io/github/repo-size/glener10/rotating-pairs-back)
![GitHub forks](https://img.shields.io/github/forks/glener10/rotating-pairs-back)
![Bitbucket open issues](https://img.shields.io/bitbucket/issues/glener10/rotating-pairs-back)
![Bitbucket open pull requests](https://img.shields.io/bitbucket/pr-raw/glener10/rotating-pairs-back)
![Contributors](https://img.shields.io/github/contributors/glener10/rotating-pairs-back.svg)

<p align="center"> 🚀 Back-End of Rotating Pairs Web Application. Responsible for returning the positions of combinations used in the application and saving monitoring logs. </p>

🏁 Table of Contents

===================

<!--ts-->

👉 [Dependencies and Environment](#dependenciesandenvironment)

👉 [Installing](#installing)

👉 [Using](#using)

👉 [Learn More](#learnmore)

👉 [License](#license)

👉 [Author](#author)

<!--te-->

===================

You can access the production Front-End application via the link [Rotating Pair](https://rotatingpairs.online)

The Back-End production version is open only to Front-End application, you can [exec locally](#using) or you can see the source code in:

[Front-End Source Code](https://github.com/glener10/rotating-pairs-front)

[Back-End Source Code](https://github.com/glener10/rotating-pairs-back)

<div id="dependenciesandenvironment"></div>

## 💻 **Dependencies and Environment**

My dependencies and versions

**Go**: go version go1.22.0 windows/amd64

If you will exec in local environment, you will need the docker and docker-compose to create a MongoDb Instance:

**Docker**: Docker version 25.0.3, build 4debf41

**docker-compose**: Docker Compose version v2.24.5-desktop.1

<div id="installing"></div>

## 🚀 **Installing**

You will need a MongoDB instance:

```bash
# container with MongoDB instance
$ docker-compose up -d
```

This execution uploads a MongoDB container with some initialized limited values, you can see them in [init-mongo.js](./init-mongo.js)

You will need to a _.env_ file inside 'src' folder

/src/.env

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE_NAME=test
```

<div id="using"></div>

## ☕ **Using**

First, check the [dependencies](#dependenciesandenvironment) and the [installation](#installing) process:

To exec you will need enter in 'src' folder:

```bash
cd src
```

And exec the _main.go_ script:

```
go run main.go
```

Now you can open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## 📖 **Learn More**

To learn more about technologies used in the application:

- [Go](https://golang.org/) - learn about Go features and API.

- [Versel](https://vercel.com/) - learn about Versel features and API.

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
