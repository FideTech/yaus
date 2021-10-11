FROM golang:1.16.2-alpine as build

RUN apk add --no-cache ca-certificates git

WORKDIR /go/src/github.com/FideTech/yaus

COPY . .

RUN go mod vendor

RUN GIT_COMMIT=$(git log --pretty=format:"%h" -n 1) && \
    GOOS=linux && \
    CGO_ENABLED=0 \
    go build -ldflags "-X github.com/FideTech/yaus/core.Commit=$GIT_COMMIT" -o yaus

FROM scratch as runtime

WORKDIR /root/

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/FideTech/yaus/yaus /root/yaus

ENV GIN_MODE=release

EXPOSE 4568

CMD ["/root/yaus"]
