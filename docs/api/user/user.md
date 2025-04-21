# POST:/

## Request Header

| Header | Content | Description | r   |
| ------ | ------- | ----------- | --- |

## Request Body

| Field       | Type                  | Description              | required   |
| ----------- | --------------------- | ------------------------ | ---------- |
| username    | string                | The username of the user | required   |
| password    | string                | The user's password      | required   |
| status      | enum("user", "admin") | The status of the user   | required   |
| email       | string                | The user's email address | required   |
| phone       | string                | The user's phone number  | optional   |
| bio         | string                | The user's bio           | optional   |
| company     | string                | The user's company       | optional   |
| location    | string                | The user's location      | optional   |
| profile_url | string                | The user's profile url   | optional   |
| nickname    | string                | The user's nickname      | deprecated |
| avatar      | string                | The user's avatar url    | deprecated |

## Response Body

| Field            | Type      | Description                   |
| ---------------- | --------- | ----------------------------- |
| code             | int       | error code                    |
| message          | string    | error message                 |
| status           | string    | error status                  |
| data             | object    | user info                     |
| data.ID          | int       | Unique identifier of the user |
| data.instanceID  | int       | Instance ID of the user       |
| data.createdAt   | timestamp | Creation timestamp            |
| data.updatedAt   | timestamp | Last update timestamp         |
| data.DeletedAt   | timestamp | Deletion timestamp (nullable) |
| data.username    | string    | The username of the user      |
| data.password    | string    | Hashed password of the user   |
| data.status      | string    | The status of the user        |
| data.nickname    | string    | The user's nickname           |
| data.email       | string    | The user's email address      |
| data.phone       | string    | The user's phone number       |
| data.avatar      | string    | The user's avatar URL         |
| data.bio         | string    | The user's bio                |
| data.company     | string    | The user's company            |
| data.location    | string    | The user's location           |
| data.profile_url | string    | The user's profile URL        |

# PUT:/{username}

## Request Header

## Request Body

| Field       | Type                  | Description              | required   |
| ----------- | --------------------- | ------------------------ | ---------- |
| password    | string                | The user's password      | optional   |
| status      | enum("user", "admin") | The status of the user   | optional   |
| email       | string                | The user's email address | optional   |
| phone       | string                | The user's phone number  | optional   |
| bio         | string                | The user's bio           | optional   |
| company     | string                | The user's company       | optional   |
| location    | string                | The user's location      | optional   |
| profile_url | string                | The user's profile url   | optional   |
| nickname    | string                | The user's nickname      | deprecated |
| avatar      | string                | The user's avatar url    | deprecated |

## Response Body

# GET:/{username}

## Request Header

## Request Body

## Response Body

# DELETE:/{username}

## Request Header

## Request Body

## Response Body

# GET:/

## Request Header

## Request Body

## Response Body

# PUT:/{username}/avatar

## Request Header

## Request Body

## Response Body

# GET:/{username}/avatar

## Request Header

## Request Body

## Response Body
