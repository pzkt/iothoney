FROM golang:1.26.5-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /example .

FROM scratch
COPY --from=build /example /example
USER 65534:65534
EXPOSE 8080
ENTRYPOINT ["/example"]