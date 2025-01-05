# ID Generator

This project provides a RESTful API for generating global unique IDs.

## Rational

This API does not generate UUIDs (Universally Unique Identifiers) because of the following considerations:

1. **Compact Identifier Size**
    - UUIDs are 16 bytes long, which can be unnecessarily large for many use cases.
    - Our solution generates identifiers that fit within **8 bytes**, reducing storage requirements and improving
      efficiency for systems that need to handle a high volume of identifiers.

2. **Ordered Identifiers**
    - Unlike standard UUIDs, our identifiers are **ordered**.
    - This is essential for use cases where maintaining an inherent order improves indexing, querying, or sorting
      performance, such as in distributed systems or databases.

3. **Scalability Without a Centralized Sequence**
    - Relying on a database sequence or any centralized mechanism for generating unique identifiers can become a *
      *bottleneck** in highly scalable systems.
    - Our solution is designed to be **scalable**, avoiding the need for a single point of failure or contention.

4. **Optimized for Distributed Systems**
    - The design supports distributed architectures, ensuring identifiers can be generated independently across multiple
      nodes while remaining unique and ordered.

This approach makes the API suitable for modern, high-performance systems requiring compact, ordered, and scalable
identifiers without the overhead of traditional UUIDs or database-based sequences.

## Overview

### ID value

Each ID value is an 8 bytes ordered value whose content is shown below:

![id-generator-ID View.drawio.png](id-generator-ID%20View.drawio.png)

The ordering is managed by the timestamp of the ID request. For each millisecond, we can ask until 16384 values.
If we make more requests, the service pauses for one millisecond.

It's possible to manage more 16384 requests per millisecond by running many service instances. This is managed by
setting
a different machine ID for each service instance.

**⚠️ Each service instance must have its own machine ID.**

### Architecture with one id-generator service instance

When we don't need to provide too many ID values at the same time, we deploy only one instance of the id-generator
service.

![id-generator-Architecture Simple.drawio.png](id-generator-Architecture%20Simple.drawio.png)

1. A client asks an ID value by an HTTP request to a Web server (like Nginx or Apache). The web server manages SSL
   connection and rate limiting.
2. The web server call the id-generator service to get an ID and returns the result

> If we don't need manage SSL or rate limiting, we can omit the web server and let clients to directly call the
> id-generator service.

### Architecture with many id-generator service instances

id-generator is scalable, in case of tens or hundreds of thousands ID requests is made at the same time we can dispatch
requests between many id-generator service instances.

![id-generator-Architecture with LB.drawio.png](id-generator-Architecture%20with%20LB.drawio.png)

1. A client sends the ID request to the load balancer. Each request is dedicated to one web server by using a dispatch
   algorithm like round-robin. The request should be uniformly dispatched between all instances.
2. Each web server is linked to one id-generator service which provide the ID value.

**⚠️ Don't forget each service instance must have its own machine ID. If we have many intances with the same machine ID,
we may have ID collisions**

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

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
   docker run -e ID_GENERATOR_MACHINE_ID=1 -p 8080:8080  csperandio/id-generator:1.0.0
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

* **Christian Sperandio** -
  *Developer* - [My Github Profile](https://github.com/chrix75) - [My LinkedIn Profile](https://www.linkedin.com/in/christian-sperandio-25182a12)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
