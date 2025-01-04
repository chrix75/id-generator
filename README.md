# ID Generator

This project provides a RESTful API for generating global unique IDs.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

* Go programming language installed.
* Gin web framework (`go get github.com/gin-gonic/gin`).

### Installing

1. Clone the repository:
   ```bash
   git clone https://github.com/chrix75/id-generator.git
   ```
2. Navigate to the project directory:
   ```bash
   cd id-generator
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running

1. Set the machine ID enrivonment variable
    ```bash
    export ID_GENERATOR_MACHINE_ID=1
    ```
2. Start the server:
   ```bash
   go run main.go
   ```
3. The API will be accessible at `http://localhost:8080`.

## Run by using Docker image
   ```bash
    docker run -e ID_GENERATOR_MACHINE_ID=1 -p 8080:8080  csperandio:id-generator
   ```

The command above starts an instance of id-generator with a machine ID set to 1.

## API Endpoints

### `/api/id`

* **Method:** GET
* **Description:** Retrieves a new unique ID.
* **Response:**
    ```json
    {
        "id": 123456789
    }
    ```

## Built With

* [Gin](https://github.com/gin-gonic/gin) - The web framework used

## Authors

* **Christian Sperandio** - *Developer* - [My Github Profile](https://github.com/chrix75) - [My LinkedIn Profile](https://www.linkedin.com/in/christian-sperandio-25182a12)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.