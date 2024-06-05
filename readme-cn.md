# http-log-server

http-log-server 是一个基于 [CloudWeGo Hertz](https://github.com/cloudwego/hertz) 框架构建的 HTTP 日志服务器。它接收 HTTP
POST 请求，并将日志记录到文件和标准输出中。

## 功能

- 接收 HTTP POST 请求，解析请求中的日志数据。
- 验证请求的授权信息。
- 将日志记录到文件和标准输出中。

## 使用方法

### 环境变量

- `TOKEN`: 用于验证请求的授权信息。

### 运行服务器

你可以通过以下命令构建和运行 Docker 镜像：

```shell
docker build -t litongava/http-log-server:1.0.0 .
docker push litongava/http-log-server:1.0.0
docker run -e TOKEN=your_token -p 8888:8888 litongava/http-log-server:1.0.0
```

### 发送日志请求

使用以下格式的 HTTP POST 请求发送日志：

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

## 构建Docker

```shell
docker build -t litongava/http-log-server:1.0.0 .
```

```shell
docker push litongava/http-log-server:1.0.0
```

## 贡献

欢迎提交 Issue 和 Pull Request 来改进该项目。

## 许可证

本项目使用 [MIT 许可证](LICENSE)。
