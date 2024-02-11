#!/bin/bash

# Wait for the database to be ready
./wait-for-it.sh db:3306 --timeout=30 -- echo "db is up"

# Command to start the application
exec go run main.go
