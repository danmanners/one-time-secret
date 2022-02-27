# Start by building the application.
FROM docker.io/golang:1.17-bullseye as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./... \
    && go build -o /go/bin/ots main.go

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/ots /
CMD ["/ots"]
