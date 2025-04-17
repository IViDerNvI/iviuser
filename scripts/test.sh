curl -XPOST http://127.0.0.1:8080/v1/user/ -i \
-d '{
    "username": "admin2",
    "email": "xxx@domain.com",
    "phone" : "11122223333",
    "password": "admin2",
    "status": "admin"
}' \
-H 'Content-Type: application/json' \
-H 'Authorization: Basic YWRtaW46YWRtaW4K'
