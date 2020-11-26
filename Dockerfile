FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .

# For china user
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# see the build env
# https://forums.docker.com/t/standard-init-linux-go-195-exec-user-process-caused-no-such-file-or-directory/43777/7
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/exe .

FROM scratch AS bin
COPY --from=build /out/exe /exe
ENTRYPOINT ["/exe"]


