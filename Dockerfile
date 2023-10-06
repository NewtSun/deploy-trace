FROM golang:alpine AS builder
LABEL authors="NewtSun"
WORKDIR /go/src/app
COPY . .
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN go build -o main .

FROM python:3.8-slim
LABEL authors="NewtSun"
WORKDIR /app
RUN pip3 config set global.index-url https://mirrors.aliyun.com/pypi/simple
RUN pip3 install requests
RUN pip3 install argparse
COPY --from=builder /go/src/app/main /app/main
COPY ./static/plugin/linux_test.py .

ENTRYPOINT ["./main"]