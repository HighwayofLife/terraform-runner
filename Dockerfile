FROM golang:1.16.5-alpine3.14 as base

RUN apk add --no-cache git

WORKDIR /tmp/terraform-runner-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN go build -o ./bin/terraform-runner-api ./cmd/...

FROM alpine:3.14.0

ARG APP_VERSION=0.1
ARG TERRAGRUNT_VERSION=0.30.70

RUN apk add --no-cache ca-certificates git curl \
  && apk add --update --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/community \
  terraform \
  && curl -fL -o /usr/local/bin/terragrunt \
  https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 \
  && chmod +x /usr/local/bin/terragrunt

COPY --from=base /tmp/terraform-runner-api/bin/terraform-runner-api /app/terraform-runner-api

EXPOSE 8080

CMD ["/app/terraform-runner-api"]
