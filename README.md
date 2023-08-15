# Your Project Name

A simple GoLang project for a RESTful API using MongoDB, mongo-driver, and chi.

## Table of Contents
- [Your Project Name](#your-project-name)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Usage](#usage)
  - [Endpoints](#endpoints)
  - [Authentication](#authentication)

## Introduction

This project is a basic RESTful API implemented in GoLang, using MongoDB as the database and the chi package for routing. It is designed to handle HTTP requests, perform authorization, and interact with a MongoDB database to serve and store data.

## Prerequisites

Make sure you have the following installed on your system:

- GoLang: [Installation Guide](https://golang.org/doc/install)
- MongoDB: [Installation Guide](https://docs.mongodb.com/manual/installation/)
- Any required Go dependencies (`go get` them as needed).

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/your-project.git
    cd your-project
    ```

2. Install dependencies:

    ```bash
    go get -u ./...
    ```

3. Build the project:

    ```bash
    go build
    ```

## Configuration

Update the `config.yml` file with your MongoDB connection details and any other configuration parameters.

```yaml
mongo:
  uri: "mongodb://localhost:27017"
  database: "your_database_name"

```

## Usage
Run the built executable:

```bash
./coin-api
```
The API server will be running at http://localhost:8000 by default.
    
## Endpoints
- GET /account/coins: Get user's coin balance .

## Authentication
To make authorized requests, include the following headers:

- Authorization: users-auth-token.