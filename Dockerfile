FROM golang:1.17

# For alterations future
ENV DIR_NAME=/usr/src/app/todo-list

RUN mkdir -p ${DIR_NAME}

# Download all the dependencies
RUN go mod download && go mod verify

# For enviroment of development
RUN go get -u github.com/cosmtrek/air

# Install the package
RUN go install -v ./...

# You are in /usr/src/app
WORKDIR ${DIR_NAME}

# Get all content from root path and copy in DIR_NAME
COPY . .

# For production
# RUN go build ${DIR_NAME}/main.go


# RUN go build -v -o /usr/local/bin/app ./...
# RUN go run /usr/src/app/todo-list/main.go

EXPOSE 1289

CMD ["air"]

# docker build -t todo-list-backend .
# docker run -it --rm --name my-running-app my-golang-app