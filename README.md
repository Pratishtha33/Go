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

2. change the directory
   ```bash
   cd .\Go

4. Install Dependencies:
   ```bash
   go get -u ./...

5. Build and Run the Application:
   ```bash
   go run main.go


## Gofr
- GoFr seamlessly integrates with GORM for efficient database management. GoFr has ingeniously incorporated various datastores, including PostgreSQL, Redis, MongoDB, SQL and Cassandra, simplifying the process of connecting to and managing different databases.

# 🚀  Project Functionalities:


## Overview

The Employee Leave Management System is an HTTP API that facilitates the management of employee leaves. It provides endpoints to perform CRUD operations on leave records, allowing employee handle employee leave requests.

## Code Structure

The codebase is organized into several packages, each serving a specific purpose:

- **datastore:** Implements the data storage and retrieval logic for leave records.
- **handler:** Defines HTTP handlers for processing leave-related requests.
- **model:** Contains data models representing the entities used in the system.

## Packages

### datastore

#### `leave` Struct

- Represents the data store for leave records.

#### `New()` Function

- Creates a new instance of the leave data store.

#### `GetByID` Function

- Retrieves a leave record by its ID from the data store.

#### `Create` Function

- Adds a new leave record to the data store.

#### `Update` Function

- Modifies an existing leave record in the data store.

#### `Delete` Function

- Removes a leave record from the data store by its ID.

#### `handler` Struct

- Implements HTTP handlers for processing leave-related requests.

#### `New()` Function

- Creates a new instance of the leave management handler.

#### `GetByID` Function

- Handles GET requests to retrieve leave details by ID.

#### `Create` Function

- Handles POST requests to create a new leave record.

#### `Update` Function

- Handles PUT requests to update an existing leave record.

#### `Delete` Function

- Handles DELETE requests to delete a leave record.

## Model

### `Leave` Struct

- Represents the data model for a leave record.

- Fields:
  - `ID`: Unique identifier for the leave record.
  - `EmployeeID`: Identifier of the associated employee.
  - `StartDate`: Start date of the leave.
  - `EndDate`: End date of the leave.
  - `Reason`: Reason for taking the leave.

# 🔧 Postman collection for trying out the APIs
1. Update

![WhatsApp Image 2023-12-17 at 15 14 19_35c4ed4a](https://github.com/Pratishtha33/Go/assets/77717155/b272dded-ddf4-4b73-807c-78913b15fa9a)

2. Read

![WhatsApp Image 2023-12-17 at 15 15 50_df87a6e5](https://github.com/Pratishtha33/Go/assets/77717155/0066ab51-a851-40cf-888b-6a42ce2a21a9)

3. Create

![WhatsApp Image 2023-12-17 at 15 16 31_12d45f5f](https://github.com/Pratishtha33/Go/assets/77717155/4bab2e92-d4ac-4c49-a4aa-43f19795ef88)

4. Delete

![WhatsApp Image 2023-12-17 at 15 14 37_065f6d8c](https://github.com/Pratishtha33/Go/assets/77717155/df849497-df9e-433b-b997-af9e4945d297)


# 📊 Sequence diagram

▶️ Client sends a request to the server

◀️ Server sends a response back to the client

![WhatsApp Image 2023-12-17 at 15 38 39_6de52914](https://github.com/Pratishtha33/Go/assets/77717155/ebf25a9d-7b53-4b75-8a0c-45e7a7aad8dd)

#  🚦 Unit Test Coverage
:white_check_mark: Run the tests to ensure everything is working correctly.


- Run the following commands to test the coverage, here's a 100% coverage

1. Test Handler:
   ```bash
     go test ./handler

2.  Test Datastore: 
    ```bash
     go test ./datastore


