FROM golang:1.18-alpine

#Set the working directory inside the container
WORKDIR /app

#Copy the dependecy files
COPY go.mod ./

#DOwnload the dependencies
RUN go mod tidy

#COpy the rest of the code to workdir container
COPY . .

#Build the go application to generate binary artifcat name "main"
RUN go build -o main .

#Expose the container port
EXPOSE 8090

#Start the container by running compiled binary
CMD ["./main"]



