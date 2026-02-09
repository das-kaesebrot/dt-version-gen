FROM docker.io/library/golang:alpine@sha256:f6751d823c26342f9506c03797d2527668d095b0a15f1862cddb4d927a7a4ced AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app .

FROM docker.io/library/alpine:3@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659

ARG APP_WORKDIR="/var/opt/dt-version-gen"

RUN mkdir -pv ${APP_WORKDIR}

COPY --from=build /usr/local/bin/app /usr/local/bin/version-gen

WORKDIR ${APP_WORKDIR}

CMD ["version-gen"]
