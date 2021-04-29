# KPZ Contactless Restaurant Backend


<!-- ABOUT THE PROJECT -->
## About The Project

KPZ Contactless Restaurant Backend is a student project for KPZ - Conference of Team Projects of University of technology of Wrocław.
This is a backend part of the app which is supposed to help restaurants with contactless customer service.

The app is to provide functionalities such as:
* Listing associated restaurants
* Ordering food
* Collecting points

### Built With

#### Main technology
* [Golang 1.15](https://golang.org/)
* [Echo - Go web framework](https://echo.labstack.com/)
* [ORM library for Golang](https://gorm.io/index.html)
* [PostgresSQL](https://www.postgresql.org/)


#### Sub technology
* [Echo Swagger - Automatically generated documentation](https://github.com/swaggo/echo-swagger)
* [Docker](https://www.docker.com/)

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* Docker
  ```
  Install Docker for your system
  ```
* Create .env file write
  ```env
  #DB_DSN="user=restaurant password=restaurant dbname=restaurant port=5432 sslmode=disable"

  DB_DSN="host=PRIVATE-IP user=restaurant password=restaurant dbname=restaurant port=5432 sslmode=disable"
  SECRET_KEY="XXXX" -> JUST ANY https://randomkeygen.com 256-bit WEP Keys
  ```


### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/angrynerds-pl/kpz-contactless-restaurant-backend.git
   ```
2. Build docker images
   ```sh
   docker-compose build
   ```
3. Start docker service
   ```sh
   docker-compose up
   ```



<!-- USAGE EXAMPLES -->
## Usage

To use swagger run docker and go on https://PRIVATE-IP:8585/swagger/index.html


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Your Name - [@your_twitter](https://twitter.com/your_username) - email@example.com

Project Link: [https://github.com/your_username/repo_name](https://github.com/your_username/repo_name)

