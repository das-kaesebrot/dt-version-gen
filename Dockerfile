FROM docker.io/library/golang:alpine@sha256:b8c9b82affbbb82a020bbcf998dad2ba9516944ac7e883524642a481a9505e9c AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app .

FROM docker.io/library/alpine:3@sha256:294b683cb724975bec92580e1e685676bd4b50bda910ddb8c51d4cabeaec77e6

ARG APP_WORKDIR="/var/opt/dt-version-gen"

RUN mkdir -pv ${APP_WORKDIR}

COPY --from=build /usr/local/bin/app /usr/local/bin/version-gen

WORKDIR ${APP_WORKDIR}

CMD ["version-gen"]
