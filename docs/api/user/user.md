# User

| fieldname  | JSON name   | database field | type   | constrict                        | description  |
| ---------- | ----------- | -------------- | ------ | -------------------------------- | ------------ |
| ID         | id          |                |        |                                  |              |
| InstanceID | instance_id |                |        |                                  |              |
| CreateAt   | update_at   |                |        |                                  |              |
| UpdateAt   | update_at   |                |        |                                  |              |
| DeleteAt   |             |                |        |                                  |              |
| UserName   | username    | username       | string | required, min=8, max=16          | user name    |
| Password   | password    | password       | string | required, min=8, max=16          | password     |
| Status     | status      | status         | string | required, oneof=admin user guest | user status  |
| NickName   | nickname    | nickname       | string | -                                | nickname     |
| Email      | email       | email          | string | required, email                  | email        |
| Phone      | phone       | phone          | string | -                                | user phone   |
| Avatar     | avatar      | avatar         | string | -                                | avator       |
| Bio        | bio         | bio            | string | -                                | user bio     |
| Company    | company     | company        | string | -                                | company      |
| Location   | location    | location       | string | -                                | location     |
| ProfileURL | profile_url | profile_url    | string | -                                | user profile |
