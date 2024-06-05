# http-log-server

http-log-server is an HTTP log server built on the [CloudWeGo Hertz](https://github.com/cloudwego/hertz) framework. It receives HTTP POST requests and logs the data to a file and standard output.

## Features

- Receives HTTP POST requests and parses the log data from the requests.
- Validates the authorization information of the requests.
- Logs the data to a file and standard output.

## Usage

### Environment Variables

- `TOKEN`: Used to validate the authorization information of the requests.

### Running the Server

You can build and run the Docker image using the following commands:

```shell
docker build -t litongava/http-log-server:1.0.0 .
docker push litongava/http-log-server:1.0.0
docker run -e TOKEN=your_token -p 8888:8888 litongava/http-log-server:1.0.0
```

### Sending Log Requests

Send log data using an HTTP POST request in the following format:

```http
POST /log HTTP/1.1
Host: localhost:8080
Authorization: Bearer your_token
Content-Type: application/json

{
    "level": "INFO",
    "args": ["This is a log message."]
}
```

## Building Docker

```shell
docker build -t litongava/http-log-server:1.0.0 .
```

```shell
docker push litongava/http-log-server:1.0.0
```

## Contribution

Feel free to submit Issues and Pull Requests to improve this project.

## License

This project is licensed under the [MIT License](LICENSE).