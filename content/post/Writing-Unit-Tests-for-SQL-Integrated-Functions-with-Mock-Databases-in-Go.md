---
author: Programmercave
date: "2024-01-04T00:00:00Z"
description: Unit testing is a critical part of software development that ensures
  the reliability and functionality of code modules. In this blog post, we will explore
  unit testing for Go functions that interact with SQL or PostgreSQL databases. Specifically,
  we will look at how to create strong test cases for functions that execute database
  queries. To achieve this, we will use mocking to create a mock database environment.
  This approach separates our functions under test from the actual database, enabling
  us to carefully evaluate their logic and behavior. We will take you through the
  process of creating effective unit tests that validate the functionality of Go functions
  integrated with SQL interactions using a mock database environment for seamless
  and comprehensive testing.
header-img: /assets/images/go-unit-test-cases/unit_test_case.png
tags:
- PostgreSQL
- SQL
- DataBase
- Go
title: Writing Unit Tests for SQL-Integrated Functions with Mock Databases in Go
toc: true
---

## Introduction

Unit testing is a critical part of software development that ensures the reliability and functionality of code modules. In this blog post, we will explore unit testing for Go functions that interact with SQL or PostgreSQL databases. Specifically, we will look at how to create strong test cases for functions that execute database queries. To achieve this, we will use mocking to create a mock database environment. This approach separates our functions under test from the actual database, enabling us to carefully evaluate their logic and behavior. We will take you through the process of creating effective unit tests that validate the functionality of Go functions integrated with SQL interactions using a mock database environment for seamless and comprehensive testing.

![Writing Unit Tests for SQL-Integrated Functions with Mock Databases in Go](/assets/images/go-unit-test-cases/unit_test_case.png)


## Understanding the student.go File

The `student.go` file encapsulates functionalities related to student data management in a Go application. It contains:

1. **StudentConfig Struct**

    ```go
    type StudentConfig struct {
        StudentID   int
        FirstName   string
        LastName    string
        DateOfBirth string
        Gender      string
        Email       string
        PhoneNumber string
    }
    ```

    The `StudentConfig` struct defines the structure for storing student-related information, including their unique ID, personal details, and contact information.

2. **SQL Queries**

    ```go
    const (
        insertOrUpdateStudentConfig = `
        INSERT INTO student (student_id, first_name, last_name, date_of_birth, gender, email, phone_number)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (student_id) DO UPDATE
        SET
            first_name = $2,
            last_name = $3,
            date_of_birth = $4,
            gender = $5,
            email = $6,
            phone_number = $7;
    `

        getStudentConfig = `
    SELECT student_id, first_name, last_name, date_of_birth, gender, email, phone_number
    FROM student
    WHERE student_id = $1
    `
    ```

    The file contains SQL queries represented as constants. These queries handle inserting or updating student data and retrieving student information based on a given `student_id`.

3. **Functions for Database Interactions**

    ```go
    func InsertOrUpdateStudentConfig(db *sql.DB, student StudentConfig) (StudentConfig, error) {
        tx, err := db.Begin()
        if err != nil {
            return StudentConfig{}, fmt.Errorf("could not begin transaction: %v", err)
        }

        _, err = tx.Exec(insertOrUpdateStudentConfig,
            student.StudentID,
            student.FirstName,
            student.LastName,
            student.DateOfBirth,
            student.Gender,
            student.Email,
            student.PhoneNumber,
        )

        if err != nil {
            tx.Rollback()
            return StudentConfig{}, fmt.Errorf("error executing query: %v", err)
        }

        err = tx.Commit()
        if err != nil {
            return StudentConfig{}, fmt.Errorf("error committing transaction: %v", err)
        }

        insertedConfig, err := GetStudentConfig(db, student.StudentID)
        if err != nil {
            return StudentConfig{}, fmt.Errorf("error fetching inserted student configuration: %v", err)
        }

        return insertedConfig, nil
    }

    func GetStudentConfig(db *sql.DB, studentID int) (StudentConfig, error) {

        var student StudentConfig

        row := db.QueryRow(getStudentConfig, studentID)
        err := row.Scan(
            &student.StudentID,
            &student.FirstName,
            &student.LastName,
            &student.DateOfBirth,
            &student.Gender,
            &student.Email,
            &student.PhoneNumber,
        )
        if err != nil {
            return StudentConfig{}, fmt.Errorf("error fetching student configuration: %v", err)
        }

        return student, nil
    }
    ```

    - `InsertOrUpdateStudentConfig`: Inserts a new student record into the database or updates an existing one, utilizing the SQL query `insertOrUpdateStudentConfig`.
    - `GetStudentConfig`: Retrieves student information from the database based on the provided `student_id`, using the SQL query `getStudentConfig`.

    These functions interact with the SQL database (`db`) by executing SQL queries and handling the data accordingly.

---

This file contains essential functions and SQL queries responsible for managing student data interactions with an SQL database in a Go application. We'll be focusing on creating unit tests for these functions in the subsequent sections to ensure their accuracy and reliability.

## Unit Test Cases

### TestInsertOrUpdateStudentConfig

Let's break down each line of the `TestInsertOrUpdateStudentConfig` function and explain its significance in testing the `InsertOrUpdateStudentConfig` function from the `student.go` file:

1. **Creating a Mocked Database Connection**

    ```go
    func TestInsertOrUpdateStudentConfig(t *testing.T) {
        // Create a mocked database connection
        db, mock, err := sqlmock.New()
        if err != nil {
            t.Fatalf("error creating mock database: %v", err)
        }
        defer db.Close()
    ```

    - `sqlmock.New()`: Initializes a mocked database (`db`) and a mock object (`mock`) to simulate database interactions.
    - `defer db.Close()`: Defers closing the database connection until the function exits, ensuring proper cleanup.

2. **Initializing Test Data**

    ```go
        // Initialize your StudentConfig for testing
        student := StudentConfig{
            StudentID:   1,
            FirstName:   "John",
            LastName:    "Doe",
            DateOfBirth: "1999-01-01",
            Gender:      "Male",
            Email:       "john@example.com",
            PhoneNumber: "1234567890",
        }
    ```

    - Creates a sample `StudentConfig` struct (`student`) with predefined data for testing purposes.

3. **Mocking Database Expectations**

    ```go
        // Create expected database query and mock behavior for InsertOrUpdateStudentConfig
        mock.ExpectBegin()
        mock.ExpectExec(regexp.QuoteMeta(insertOrUpdateStudentConfig)).WithArgs(student.StudentID, student.FirstName, student.LastName, student.DateOfBirth, student.Gender, student.Email, student.PhoneNumber).WillReturnResult(sqlmock.NewResult(0, 1))
        mock.ExpectCommit()
        mock.ExpectQuery(regexp.QuoteMeta(getStudentConfig)).WithArgs(student.StudentID).
            WillReturnRows(sqlmock.NewRows([]string{"student_id", "first_name", "last_name", "date_of_birth", "gender", "email", "phone_number"}).
                AddRow(student.StudentID, student.FirstName, student.LastName, student.DateOfBirth, student.Gender, student.Email, student.PhoneNumber))
    ```

    - `mock.ExpectBegin()`, `mock.ExpectExec()`, `mock.ExpectCommit()`: Sets up expectations for the sequence of SQL actions expected when calling `InsertOrUpdateStudentConfig`. It simulates the beginning of a transaction, an execution of an SQL query, and committing the transaction.
    - `mock.ExpectQuery()`: Sets the expectation for the SQL query that would be triggered internally by the `GetStudentConfig` function within the `InsertOrUpdateStudentConfig` function.
    - `regexp.QuoteMeta(getStudentConfig)`: This constructs a regex pattern to match the SQL query string (`getStudentConfig`) used within the `GetStudentConfig` function in `student.go`. It prepares the mock to expect this specific query to be executed.

    **Defining Return Rows for the Mocked Query**
    - `WillReturnRows`: Specifies the expected rows that the mocked query should return when triggered.
    - `sqlmock.NewRows([]string{...})`: Creates a set of rows with column names defined in the array. This prepares the structure of the returned rows.
    - `AddRow()`: Adds a row to the expected result set, populating it with values corresponding to the columns defined in `NewRows`.

    **Why is this Necessary?**

    In the context of testing `InsertOrUpdateStudentConfig`, it's essential to understand that this function internally calls `GetStudentConfig` after performing an insertion or update. Mocking the `GetStudentConfig` query is crucial because:
    - It ensures that the test focuses solely on the behavior of `InsertOrUpdateStudentConfig` without actually executing the `GetStudentConfig` query against a real database.
    - By setting expectations for `GetStudentConfig`, it allows the test to simulate the subsequent behavior and validate the flow within `InsertOrUpdateStudentConfig` after the database insertion or update.

    Mocking the `GetStudentConfig` query with expected return rows allows the test to confirm that after inserting or updating a student's information, the function successfully performs a subsequent fetch of that student's data without actually executing a real query against a database. This isolation is key in unit testing to verify the behavior of the function under test.

4. **Testing the Function**

    ```go
        // Test InsertOrUpdateStudentConfig function
        insertedConfig, err := InsertOrUpdateStudentConfig(db, student)
        if err != nil {
            t.Fatalf("InsertOrUpdateStudentConfig failed: %v", err)
        }
    ```

    - Calls the `InsertOrUpdateStudentConfig` function with the mock database and test data.
    - Checks for errors during function execution.

5. **Assertions**

    ```go
        // Add assertions to check the correctness of insertedConfig
        if insertedConfig.StudentID != student.StudentID {
            t.Errorf("Expected StudentID: %d, Got: %d", student.StudentID, insertedConfig.StudentID)
        }
        if insertedConfig.FirstName != student.FirstName {
            t.Errorf("Expected FirstName: %s, Got: %s", student.FirstName, insertedConfig.FirstName)
        }
        // Similarly, add assertions for other fields if needed
    ```

    - Validates the returned `insertedConfig` against the input `student` data to ensure the function behaves correctly.
    - Compares each field of the expected and actual `StudentConfig` structs to verify correctness.

6. **Verifying Mock Expectations**

    ```go
        // Ensure all expectations were met
        if err := mock.ExpectationsWereMet(); err != nil {
            t.Fatalf("mock expectations were not met: %v", err)
        }
    }
    ```

    - Ensures that all the expected interactions with the mock database were met during the test execution. This confirms that the `InsertOrUpdateStudentConfig` function behaved as expected according to the defined mock behavior.

---

The `insertOrUpdateStudentConfig` function is thoroughly tested by this comprehensive testing setup, which emulates database interactions using a mock environment to ensure its correctness and adherence to expected database operations defined in the function from `student.go` file.

### TestGetStudentConfig

The structure and assertions in the `TestGetStudentConfig` function are similar to those in `TestInsertOrUpdateStudentConfig`. Given the similarity in testing approaches and the behavior being tested, no further explanation is necessary as it mirrors the process utilized in `TestInsertOrUpdateStudentConfig`.

```go
func TestGetStudentConfig(t *testing.T) {
	// Create a mocked database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	// Initialize a StudentConfig for testing
	studentID := 1
	expectedStudent := StudentConfig{
		StudentID:   studentID,
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "1999-01-01",
		Gender:      "Male",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
	}

	// Define the expected SELECT query and mock behavior for GetStudentConfig
	mock.ExpectQuery(regexp.QuoteMeta(getStudentConfig)).WithArgs(studentID).
		WillReturnRows(sqlmock.NewRows([]string{"student_id", "first_name", "last_name", "date_of_birth", "gender", "email", "phone_number"}).
			AddRow(expectedStudent.StudentID, expectedStudent.FirstName, expectedStudent.LastName, expectedStudent.DateOfBirth, expectedStudent.Gender, expectedStudent.Email, expectedStudent.PhoneNumber))

	// Test GetStudentConfig function
	retrievedStudent, err := GetStudentConfig(db, studentID)
	if err != nil {
		t.Fatalf("GetStudentConfig failed: %v", err)
	}

	// Add assertions to check the correctness of retrievedStudent
	if retrievedStudent.StudentID != expectedStudent.StudentID {
		t.Errorf("Expected StudentID: %d, Got: %d", expectedStudent.StudentID, retrievedStudent.StudentID)
	}
	if retrievedStudent.FirstName != expectedStudent.FirstName {
		t.Errorf("Expected FirstName: %s, Got: %s", expectedStudent.FirstName, retrievedStudent.FirstName)
	}
	// Similarly, add assertions for other fields if needed

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("mock expectations were not met: %v", err)
	}
}
```

---

The complete code is available on [Github](https://github.com/{{< param "github_username" >}}/Golang-Codes/tree/main/unit_test_cases)

## Conclusion

When integrating SQL database interactions within functions, unit testing is the foundation of strong software development. In this exploration, we looked at how to test Go functions that handle SQL queries by using mock databases.

We broke down the `student.go` file in order to understand why unit tests are important for isolating functionality. These include the `StudentConfig` struct, SQL queries and database interaction functions.

The respective functions’ behavior was carefully tested by the test functions, `TestInsertOrUpdateStudentConfig` and `TestGetStudentConfig`. We made sure that these functions were correct without affecting the real database by simulating database interactions through mock environments.

Unit testing SQL-integrated functions in Go requires a lot of setup, from creating mock database connections to defining expected behaviors for SQL queries. This ensures that they work as expected and handle database interactions well.

Developers strengthen their codebase by adopting unit testing methodologies, which improves its reliability and maintainability. The purpose of this blog was to demystify unit testing in Go’s SQL integrated functions so that developers can create robust applications that are resistant to errors.