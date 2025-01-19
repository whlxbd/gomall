FROM golang:1.23.5-alpine

RUN go env -w GOPROXY=https://goproxy.cn,direct
ENV GO_ENV=dev

WORKDIR /

# 复制项目文件
COPY . .

ARG NAME
ARG PORT

# 设置工作目录
WORKDIR /app/${NAME}

# 下载依赖
RUN go mod download

# 暴露端口（根据您的微服务需要调整）
EXPOSE ${PORT}

CMD go run .