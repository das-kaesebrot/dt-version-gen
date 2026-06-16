FROM docker.io/library/golang:alpine@sha256:f1ddd9fe14fffc091dd98cb4bfa999f32c5fc77d2f2305ea9f0e2595c5437c14 AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app .

FROM docker.io/library/alpine:3@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b

ARG APP_WORKDIR="/var/opt/dt-version-gen"

RUN mkdir -pv ${APP_WORKDIR}

COPY --from=build /usr/local/bin/app /usr/local/bin/version-gen

WORKDIR ${APP_WORKDIR}

CMD ["version-gen"]
