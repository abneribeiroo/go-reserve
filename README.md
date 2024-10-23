
<div>
    <h3 align="center">Go-reserve API</h3>
    <p align="center">Use this API to make equipment reservations securely and easily <br /> 
    Reserving and renting equipment has never been easier.
    </p>
    <p align="center">    
    <br />
    <img src="https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white" alt="GitHub" />
    <img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
    <img src="https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
    <img src="https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white" alt="Postgres" />
    <img src="https://img.shields.io/github/stars/abneribeiroo/go-reserve" alt="GitHub Repo stars" />
    <br />
    <a href="https://github.com/abneribeiroo/go-reserve/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/abneribeiroo/go-reserve/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
    </p>
</div>




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
    git clone https://github.com/abneribeiroo/go-reserve.git
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

- [GitHub](https://github.com/abneribeiroo)
- [Twitter](https://x.com/heisptol)

