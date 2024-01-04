#!/bin/bash

for((num=10; num>0; num--)); do
  output=$(curl -X POST -d '{"username": "user1", "email":"user1@gmail.com", "password":"123"}' http://localhost:8080/api/auth/signup)
  echo"$output"
done
