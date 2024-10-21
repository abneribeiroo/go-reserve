# Project go-reserve

Go-reserve is an equipment reservation management system where users can register, log in, and reserve available equipment. The system allows administrators to manage equipment and reservations, providing an efficient platform for resource utilization.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install/source) (version 1.16 or higher)
- [Docker](https://docs.docker.com/get-docker/)
- [PostgreSQL](https://www.postgresql.org/download/) (or another compatible database)

### Setting Up the Environment

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/go-reserve.git
    cd go-reserve
    ```

2. Create a `.env` file in the root of the project with the following variables:
    ```env
    DB_DATABASE=yourdatabase
    DB_USERNAME=yourusername
    DB_PASSWORD=yourpassword
    DB_HOST=localhost
    DB_PORT=5432
    DB_SCHEMA=public
    PORT=8080
    SECRET_KEY=yoursecretkey
    ```

### Makefile

Run the following commands to manage the project:

- **Build and run all tests:**
    ```bash
    make all
    ```

- **Build the application:**
    ```bash
    make build
    ```

- **Run the application:**
    ```bash
    make run
    ```

- **Create the database container:**
    ```bash
    make docker-run
    ```

- **Shut down the database container:**
    ```bash
    make docker-down
    ```

- **Database integration tests:**
    ```bash
    make itest
    ```

- **Live reload the application:**
    ```bash
    make watch
    ```

- **Run the test suite:**
    ```bash
    make test
    ```

- **Clean the binaries from the last build:**
    ```bash
    make clean
    ```

## Project Structure

- `cmd/` - Source code for the main application.
- `internal/` - Contains internal components of the project, such as servers, controllers, and models.
- `database/` - Handles database connection and operations.
- `middlewares/` - Middleware implementations for access control and other functionalities.
- `models/` - Data structures and functions related to models (users, equipment, reservations).
- `controllers/` - HTTP request handlers.

## How to Contribute

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/new-feature`).
3. Make your changes and commit (`git commit -m 'Adding new feature'`).
4. Push to your branch (`git push origin feature/new-feature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for more details.

## Contact

If you have questions or feedback, feel free to reach out:

- [GitHub](https://github.com/yourusername)

