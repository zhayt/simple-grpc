# Simple gRPC

Simple gRPC is my personal project to learn gRPC. It consists of a client that sends a new student to be saved in a MongoDB database, and a server that receives the student, generates a UUID for it, and saves it to the database. If everything goes well, the server sends a response containing the UUID of the saved student, otherwise an error is returned.

## Technologies Used

- gRPC: A high-performance, open-source universal RPC framework
- MongoDB: A popular NoSQL database that uses documents instead of tables and rows.

## How to Run
To run the project, you need to have Docker
1. Clone the repository to your local machine
2. Install the required dependencies using the following command:
   ```
   go mod tidy
   ```
3. Start the server with mongo-db by running the following command in the root directory of the project:
   ```
   docker-compose up
   ```
4. In a new terminal window, start the client by running the following command in the root directory of the project:
   ```
   go run cmd/client/client.go
   ```
5. Follow the prompts in the client to enter the student details.
6. If the student is successfully saved, the client will display the UUID of the saved student. If an error occurs, an error message will be displayed.

## API

### CreateStudent

Saves a new student to the MongoDB database.

```proto
 rpc CreateStudent(CreateStudentRequest) returns (CreateStudentResponse) {}
```

#### Request

| Field      | Type | Description                                   |
|------------| --- |-----------------------------------------------|
| `name`     | string | Name of the student                           |
| `email`    | string | Email address of the student                  |
| `degree`   | string | Student Password |

#### Response

| Field | Type | Description |
| --- | --- | --- |
| `id` | string | UUID of the saved student |

## Future Improvements

- Add additional methods to the StudentService, such as UpdateStudent and DeleteStudent.
- Improve error handling to provide more informative error messages.
- Implement authentication and authorization using gRPC's built-in security features.
