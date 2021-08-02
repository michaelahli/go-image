# Backend Engineer Qualification Test

## Description
This application purpose is to simulate image upload process from client side to server side. Overall this application can upload image and see metadata saved in database.

## Technologies
1. Go Language (Native)
2. React.js
3. MongoDB
4. yarn

## Usage
There are 3 steps to run this application :
1. Start the Golang server <br/>
    To start the Golang server, you'll have to go to inside directory and download dependencies using : </br>
    ```
    cd server
    go mod tidy
    ```
    and then simply start the server using command :
    ```
    go run main.go
    ```
    Now the Golang server will be running inside port 8080.
2. Start the React server
    To start the React server, you'll have to go to inside directory and download dependencies using : </br>
    ```
    cd client
    yarn install
    ```
    and then simply start the server using command :
    ```
    yarn start
    ```
    Now the Website will be running inside port 3000.
## APIs
In this api, there are 2 APIs provided which are :
```
curl -X POST http://localhost:8080/api/upload-image \
 -H 'Authorization=eyJ1c2VybmFtZSI6ICJTZWVkZXIgVXNlciJ9' \
 -d @image.jpeg
```
which is API to upload image to server and also another API to check data saved in the website :
```
curl http://localhost:8080/api/check-metadata \
 -H 'Authorization=eyJ1c2VybmFtZSI6ICJTZWVkZXIgVXNlciJ9'
```

# Author
this project is authored by Michael Ahli
