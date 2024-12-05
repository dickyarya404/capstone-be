FROM golang:1.22-alpine3.19 as build-stage

WORKDIR /projects/recything-be

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o recything-be ./cmd/api/

FROM gcr.io/distroless/base-debian11 AS build-release-stage

COPY --from=build-stage /projects/recything-be/web/ /web/
COPY --from=build-stage /projects/recything-be/docs/ /docs/
COPY --from=build-stage /projects/recything-be/recything-be /recything-be

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./recything-be"]