# Holiday Fetcher and CRUD API

This Go program fetches public holidays for a given country using the Calendarific API and provides a RESTful API for performing CRUD (Create, Read, Update, Delete) operations on the holiday data stored in a PostgreSQL database.

## Overview

The program consists of two main functionalities:
1. **Fetching Public Holidays**: Retrieves public holiday information for a specific country for the current year using the [Calendarific API](https://calendarific.com/).
2. **CRUD Operations**: Provides RESTful API endpoints to create, read, update, and delete holiday records in a PostgreSQL database.

By default, it fetches holidays for India (`IN`), but you can specify any country code as a command-line argument.

---

## Prerequisites

Before you run this program, make sure you have the following:

1. **Go 1.18 or later** installed on your machine.
2. **PostgreSQL** installed and running.
3. A valid **API key** for the Calendarific API. You can sign up for an API key at [https://calendarific.com/](https://calendarific.com/).
4. A `.env` file with the following environment variables for the database connection:
   ```plaintext
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_HOST=your_db_host
   DB_PORT=your_db_port
   ```

---

## Installation

1. Clone or download this repository to your local machine:
   ```sh
   git clone https://github.com/N-Jangra/cal
   ```

2. Change to the project directory:
   ```sh
   cd <project-directory>
   ```

3. Replace the `apiKey` value in the `main.go` file with your own Calendarific API key:
   ```go
   apiKey := "your_calendarific_api_key"
   ```

4. Install the required Go dependencies:
   ```sh
   go mod tidy
   ```

---

## Database Setup

1. **Create the Database** (if not already created):
   ```sh
   sudo -u postgres psql
   CREATE DATABASE calc;
   \c calc
   ```

2. **Create the `holidays` Table**:
   Run the following SQL query to create the `holidays` table:
   ```sql
   CREATE TABLE holidays (
       id SERIAL PRIMARY KEY,             -- Auto-incremented ID
       name VARCHAR(255) NOT NULL,        -- Holiday name (e.g., "Christmas")
       iso_date DATE NOT NULL,            -- ISO format date (e.g., "2025-12-25")
       international BOOLEAN NOT NULL    -- Boolean flag indicating if the holiday is international
   );
   ```

---

## Usage

### Start the Application

1. Run the program:
   ```sh
   go run main.go
   ```

2. The program will:
   - Connect to the PostgreSQL database using the credentials from the `.env` file.
   - Fetch public holidays for the default country (`IN`) or the country code provided as a command-line argument.
   - Start an HTTP server on `localhost:8080`.

### API Endpoints

The following RESTful API endpoints are available:

| HTTP Method | Endpoint          | Description                                      |
|-------------|-------------------|--------------------------------------------------|
| `GET`       | `/`               | Redirects to the root directory (home page).     |
| `GET`       | `/app`            | Fetches holidays from the Calendarific API and inserts them into the database. |
| `POST`      | `/n`              | Adds a new holiday record to the database.       |
| `GET`       | `/g/:iso_date`    | Fetches a specific holiday by its ISO date.      |
| `GET`       | `/ga`             | Fetches all holiday records from the database.   |
| `PUT`       | `/u/:id`          | Updates a specific holiday record by its ID.     |
| `DELETE`    | `/d/:iso_date`    | Deletes a specific holiday by its ISO date.      |
| `DELETE`    | `/da`             | Deletes all holiday records from the database.   |

---

### Example Requests

#### 1. Fetch Holidays from Calendarific API
```sh
curl -X GET http://localhost:8080/app
```

#### 2. Add a New Holiday
```sh
curl -X POST http://localhost:8080/n \
-H "Content-Type: application/json" \
-d '{
  "name": "New Year'\''s Day",
  "date": {
    "iso": "2025-01-01"
  },
  "international": true
}'
```

#### 3. Fetch a Specific Holiday by ISO Date
```sh
curl -X GET http://localhost:8080/g/2025-01-01
```

#### 4. Fetch All Holidays
```sh
curl -X GET http://localhost:8080/ga
```

#### 5. Update a Holiday by ID
```sh
curl -X PUT http://localhost:8080/u/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Updated Holiday Name",
  "date": {
    "iso": "2025-01-01"
  },
  "international": false
}'
```

#### 6. Delete a Specific Holiday by ISO Date
```sh
curl -X DELETE http://localhost:8080/d/2025-01-01
```

#### 7. Delete All Holidays
```sh
curl -X DELETE http://localhost:8080/da
```

---

## Error Handling

- If the program cannot retrieve holidays (e.g., due to an invalid API key or network issues), it will return an appropriate error message.
- If the Calendarific API returns a non-200 status code, the program will return an error message.
- Database errors (e.g., invalid SQL syntax or connection issues) will also be logged and returned.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Contributing

Feel free to fork this repository and submit pull requests if you wish to contribute improvements or fixes. If you have suggestions or issues, please create an issue in the repository.

---

## Notes

- Replace the `apiKey` in the `main.go` file with your own valid key from the Calendarific API before running the program.
- Ensure the `.env` file is correctly configured with your PostgreSQL database credentials.
- The program defaults to fetching holidays for India (`IN`). You can specify a different country code as a command-line argument when running the program.

## Contact 

For any questions or suggestions, feel free to reach out:

    Author: Nitin
    Email: itznitinjangra@gmail.com
    GitHub: N-Jangra
