# Login

Used to collect a Token for a registered User.

**URL** : `/api/login/`

**Method** : `POST`

**Auth required** : Yes

**Header**

```
Basic (base64<username:passwd>)
```

**Request example**
```shell
curl --location --request POST 'https://127.0.0.1:8443/v1/login' \
--header 'Authorization: Basic dXNlcmFjY291bnQ6dmFsaWRwYXNzd2Q='
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "code": 200,
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDQ4NzQyMTYsInJlYWxtIjoiIiwidXNlcl9pbnN0YW5jZSI6MTkxMjczNjk4OTYwMjM4NTkyMCwidXNlcl9zdGF0dXMiOiJ1c2VyIiwidXNlcm5hbWUiOiJ1c2VyYWNjb3VudCJ9.3s5inJM1oKqgPq3czyBo_rxxQpmOHZ7Tibj61xIepKc",
    "message": "success",
    "status": "success"
}
```

## Error Response

**Condition** : If 'username' and 'password' combination is wrong.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "code": 400,
    "data": null,
    "message": "error",
    "status": "error"
}
```