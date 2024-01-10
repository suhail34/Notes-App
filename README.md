
# Notes App

### Prerequistes
- You need to have [golang version 1.20](https://go.dev/doc/install) to run the project locally without docker container

- [Docker](https://www.docker.com/get-started/) need to be installed to run it inside a container if you are running it inside the container no need to install golang

- Create .env file in root and add two variables `SECRET_KEY=<YOUR_SECRET_KEY>` and `MONGO_URI=mongodb+srv://:@/?retryWrites=true&w=majority&ssl=false`

### Run The Program
- To run the server locally use make server-local

- To run the server in docker container use make server-container-start

- To stop the server and delete the container use make server-container-stop

### Usecase of Endpoint
I have deployed my docker container you can use this urls to test

| endpoint        | Method           | Command  |
| ------------- |:-------------:| -----:|
| /api/auth/signup      | POST | curl -X POST -d '{"username": "<ENTER_USER>","email": "<ENTER_EMAIL>", "password":"<ENTER_PASS"}' https://notes-api-zy44.onrender.com/api/auth/signup |
| /api/auth/login      | POST | curl -X POST -d '{"user_name": "<ENTER_USER>", "password":"<ENTER_PASS"}' https://notes-api-zy44.onrender.com/api/auth/login |
| /api/notes | GET      |    curl -X GET -H "Authorization: <YOUR_TOKEN>" https://notes-api-zy44.onrender.com/api/notes/ |
| /api/notes | POST | curl -X POST -H "Authorization: <YOUR_TOKEN>" -d '{"title":"<YOUR_TITLE>", "description":"<YOUR_DESC>"}' https://notes-api-zy44.onrender.com/api/notes/ |
| /api/notes/:id | GET | curl -X GET -H "Authorization: <YOUR_TOKEN>" https://notes-api-zy44.onrender.com/api/notes/<YOUR_NOTES_ID> |
| /api/notes/:id | PUT | curl -X PUT -H "Authorization: <YOUR_TOKEN>" -d '{"description":"<YOUR_DESC>"}' https://notes-api-zy44.onrender.com/api/notes/<YOUR_NOTES_ID> |
| /api/notes/:id | DELETE | curl -X DELETE -H "Authorization: <YOUR_TOKEN>"  https://notes-api-zy44.onrender.com/api/notes/<YOUR_NOTES_ID> |
| /api/notes/:id/share | POST | curl -X POST -H "Authorization: <YOUR_TOKEN>" -d '{"userToShareWith":"[<YOUR_USERS>]"}' https://notes-api-zy44.onrender.com/api/notes/<YOUR_NOTES_ID>/share |
| /api/notes/search?q=<SEARCH_TERM> | GET | curl -X GET -H "Authorization: <YOUR_TOKEN>" https://notes-api-zy44.onrender.com/api/notes/search?q=<SEARCH_TERM> |
