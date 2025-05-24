#!/bin/bash

echo "Building frontend..."
cd frontend
npm install
npm run build
cd ..

echo "Building backend..."
go build -tags netgo -ldflags '-s -w' -o app 