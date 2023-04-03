# Use the latest Fedora image as the base
FROM fedora:latest

# Install Go and other required tools
RUN dnf -y update && \
    dnf -y install golang git

# # Set up the Go workspace
# ENV GOPATH /go
# ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# # Create a directory for your Go project
# RUN mkdir -p $GOPATH/src/my-go-project
# WORKDIR $GOPATH/src/my-go-project

# Copy the contents of your local Go project into the container
COPY . .

# # Download and install the project's dependencies
# RUN go get -d -v ./...

# # Build the Go project
# RUN go build -o my-go-app

# Expose the port your application will run on
EXPOSE 10443

# Start the Go application
CMD ["./videoweb.exe"]
