# Use version 37 Fedora image as the base
FROM fedora:37

# Install Go and other required tools
RUN dnf -y update && \
    dnf -y install golang

# Copy the contents of your local Go project into the container
COPY . .

# Expose the port your application will run on
EXPOSE 10443

# Start the Go application
CMD ["./videoweb.exe"]
