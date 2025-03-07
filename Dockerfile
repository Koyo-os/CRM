FROM golang:1.24.0-bullseye AS builder

COPY . /workdir
WORKDIR /workdir

ENV CGO_CPPFLAGS="-D_FORTIFY_SOURCE=2 -fstack-protector-all"

ENV GOFLAGS="-buildmode=pie"

RUN go build -o bin/app cmd/main.go 

FROM gcr.io/distroless/base-debian11:nonroot
COPY --from=builder /workdir/app /bin/app

USER 65534

ENTRYPOINT ["/bin/app"]
