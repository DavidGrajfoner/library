# Library Management System

This project is a Library Management System built using Go and PostgreSQL. It allows users to borrow and return books while managing a collection of available books. The project is containerized using Docker for easy setup and deployment.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Postman Collection](#postmancollection)

## Features

- User creation
- Borrowing and returning books
- Managing a collection of books
- PostgreSQL as the database backend
- Easy setup with Docker

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Docker**: You need Docker installed on your machine. You can download it from [the official Docker website](https://docs.docker.com/get-docker/).
- **Docker Compose**: Docker Compose is usually included with Docker, but you can find installation instructions [here](https://docs.docker.com/compose/install/).

## Installation

Follow these steps to set up the project on your local machine:

1. **Clone the Repository**:
   Open a terminal and run:
   ```bash
   git clone https://github.com/DavidGrajfoner/library.git
   cd library

2. **Build and Start the Containers**
   Run the following command to build and start the application and database services:
   ```bash
   docker-compose up --build

## Postman Collection
I have prepared a Postman collection to help you test the API endpoints easily. You can import the Postman collection into your Postman app.
