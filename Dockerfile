FROM golang:1.17

# For alterations future
ENV DIR_NAME=/usr/src/app/blog-backend

RUN mkdir -p ${DIR_NAME}

# Download all the dependencies
# RUN go mod download && go mod verify

# Install the package
# RUN go install -v ./...

# You are in /usr/src/app/blogbackend
WORKDIR ${DIR_NAME}

# Get all content from root path and copy in DIR_NAME
# current directory to ${DIR_NAME} = /usr/src/app/blog-backend
COPY . .

# Install the dependecies
RUN go get -d -v ./...
RUN go mod tidy

# For enviroment of development
# RUN go get -u github.com/cosmtrek/air

# For production
# RUN go run ${DIR_NAME}/main.go
# RUN go build ${DIR_NAME}/main.go


# RUN go build -v -o /usr/local/bin/app ./...
# RUN go run /usr/src/app/todo-list/main.go

EXPOSE 1289

CMD ["go", "run", "main.go"]
# CMD [ "air" ]
# CMD [ "main" ]


# build image
# docker build -t platform-backend:1.0 .

# platform-backend is the name of image
# :1.0 is tag
# . (dot) is local directory

# You are have error in this image?
# docker rmi -f platform-backend:1.0

# RUN SERVER AND GOOOOOOOO!
# docker run -p 1289:1289 platform-backend:1.0
