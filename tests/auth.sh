function auth() {
    username=$1
    password=$2

    echo $(http --ignore-stdin --form --follow --timeout 3600 POST 'http://localhost:11000/token' \
        'grant_type'='password' \
        'client_id'='000000' \
        'client_secret'='999999' \
        'scope'='read' \
        'username'=$username \
        'password'=$password \
        'response_type'='token')
}

AUTH_JSON=$(auth admin admin)

export ACCESS_TOKEN=$(echo $AUTH_JSON | jq -r '.access_token')
export REFRESH_TOKEN=$(echo $AUTH_JSON | jq -r '.refresh_token')

