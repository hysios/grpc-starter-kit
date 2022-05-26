API_HOST=http://localhost:11000
API_PATH=/api/v1/users
API_URL=$API_HOST$API_PATH

function getUsers() {
    echo $(http -A bearer -a $ACCESS_TOKEN --timeout 3600 GET $API_URL$1)
}
