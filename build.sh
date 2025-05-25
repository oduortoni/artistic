#!/bin/bash

echo "Building frontend..."
cd frontend
npm install
npm run build
cd ..

echo "Building backend..."
go build -tags netgo -ldflags '-s -w' -o rtstk

echo "Build complete!"
echo "To run the application:"
echo "1. With default port (8000):"
echo "   ./rtstk"
echo "2. With custom port:"
echo "   PORT=<port> ./rtstk"
echo "Eg: PORT=3000 ./rtstk" 