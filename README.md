<p align="center">
  <a href="" rel="noopener">
 <img height=200px src="https://miro.medium.com/max/900/1*5JXt0wiQjX_FDwYvrxPN9Q.png" alt="Project logo"></a>
</p>

<h3 align="center">Golang Backend Skeleton with clean architecture DDD Hexagonal</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> Base Backend project implementing clean architecture DDD and Hexagonal with library Fx + Fiber + Inversify  (Gorm + sqlite) or (MongoDB driver) 
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

This project serves as a skeleton for a scalable application with Golang, as it has a basic implementation of the following:

- Dependency Injection
- Validation DTO
- Security (Json Web Token)
- Handlers Error
- Guards
- Module Auth
- Module Users

## 🏁 Getting Started <a name = "getting_started"></a>

### Prerequisites

```
go 1.20
air
MongoDB latest
```

### Installing

```
git clone https://github.com/sistemasnegros/backend-skelenton-golang
cd backend-skelenton-golang/src
```

Once you've cloned the project, install dependencies with

```
go mod download
```

cloud Air - Live reload for Go apps

```
go install github.com/cosmtrek/air@latest
```

Start MongoDB with docker-compose

```
docker-compose up --build -d
```

## 🎈 Usage <a name="usage"></a>

### Developing

```bash
air
```

## 🚀 Deployment <a name = "deployment"></a>

To create a production version of your app:

```bash
go build
```

## ⛏️ Built Using <a name = "built_using"></a>

- [Go](https://go.dev/) - Programming Language.
- [Fx](https://www.typescriptlang.org/) - Fx is a dependency injection system for Go..
- [Fiber](https://docs.gofiber.io/) - Fiber is an Express inspired web framework built on top of Fasthttp.
- [Gorm](https://gorm.io) - The fantastic ORM library for Golang.
- [MongoDB](https://gorm.io) - MongoDB Go Driver.

## ✍️ Authors <a name = "authors"></a>

- [@sistemasnegros](https://github.com/sistemasnegros) - Idea & Initial work

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- thanks to video [Fazt Code Fiber](https://youtu.be/8ES_ecfbZsk) - [Fazt Code Gorm](https://youtu.be/B6gQ1B0cn4s)
- thanks to video [Go Simplified](https://youtu.be/-XcyraChDUw)
