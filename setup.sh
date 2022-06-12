#!/usr/bin/env bash

cd frontend
npm i
cd ..
go mod tidy
echo "You are set up and ready to Go."