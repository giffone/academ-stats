# excel_table
FROM golang:1.22.3

ENV TZ=Asia/Almaty

WORKDIR /web

# install golang migrate
ARG VERSION=v4.17.0
ARG OS=linux
ARG ARCH=amd64
RUN curl -L "https://github.com/golang-migrate/migrate/releases/download/${VERSION}/migrate.${OS}-${ARCH}.tar.gz" | tar xvz -C /usr/local/bin migrate

COPY . .

RUN go mod download
RUN go build -o excel_table cmd/excel_table/main.go

EXPOSE 4444

CMD ["./excel_table"]