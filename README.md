# Employee Leave Management API

The Employee Leave Management System API is a Go language HTTP API built using the GoFr framework. It provides endpoints for managing employee leaves, including operations such as creating a leave request, updating leave details, retrieving leave information, and deleting leave records.

## Prerequisites

Before running the project, ensure you have the following installed:

- [Go](https://golang.org/dl/) (Version 1.16 or later)
- [Git](https://git-scm.com/downloads) 
- [GoFr](https://gofr.dev/docs)
- Make sure to install it using the following command:
  ```
  go get gofr.dev
- [Docker](https://hub.docker.com/_/mysql)
- To simplify the setup and management of the database for your REST API, consider using a Docker image for SQL.

## Running the Project

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Pratishtha33/Go.git

2. ```bash
   cd .\Go

3. Install Dependencies:
   ```bash
   go get -u ./...

4. Build and Run the Application:
   ```bash
   go run main.go


## Gofr
- GoFr seamlessly integrates with GORM for efficient database management. GoFr has ingeniously incorporated various datastores, including PostgreSQL, Redis, MongoDB, SQL and Cassandra, simplifying the process of connecting to and managing different databases.

# Postman collection for trying out the APIs
1. Update

![WhatsApp Image 2023-12-17 at 15 14 19_35c4ed4a](https://github.com/Pratishtha33/Go/assets/77717155/b272dded-ddf4-4b73-807c-78913b15fa9a)

2. Read

![WhatsApp Image 2023-12-17 at 15 15 50_df87a6e5](https://github.com/Pratishtha33/Go/assets/77717155/0066ab51-a851-40cf-888b-6a42ce2a21a9)

3. Create

![WhatsApp Image 2023-12-17 at 15 16 31_12d45f5f](https://github.com/Pratishtha33/Go/assets/77717155/4bab2e92-d4ac-4c49-a4aa-43f19795ef88)

4. Delete

![WhatsApp Image 2023-12-17 at 15 14 37_065f6d8c](https://github.com/Pratishtha33/Go/assets/77717155/df849497-df9e-433b-b997-af9e4945d297)
