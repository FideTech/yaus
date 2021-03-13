FROM golang:1.16.2 as backend

WORKDIR /go/src/github.com/github.com/FideTech/yaus

COPY . .

RUN go mod vendor

RUN GIT_COMMIT=$(git log --pretty=format:"%h" -n 1) && \
    GOOS=linux && \
    go build -mod=vendor -ldflags "-X github.com/github.com/FideTech/yaus/core.Commit=$GIT_COMMIT" -o yaus

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=backend /go/src/github.com/github.com/FideTech/yaus/yaus .

ENV GIN_MODE=release

EXPOSE 4568

CMD ["/root/yaus"]
