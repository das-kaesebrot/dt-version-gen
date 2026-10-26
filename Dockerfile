FROM docker.io/library/golang:alpine@sha256:f92b6ef800e499660581efdabdf25d9d817a9d124eaf900924f0504e7e27e12d AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app .

FROM docker.io/library/alpine:3@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

ARG APP_WORKDIR="/var/opt/dt-version-gen"

RUN mkdir -pv ${APP_WORKDIR}

COPY --from=build /usr/local/bin/app /usr/local/bin/version-gen

WORKDIR ${APP_WORKDIR}

CMD ["version-gen"]
