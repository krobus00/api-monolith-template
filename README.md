# Monolith Service

## Response
```json
{
	"message": "validation error", // string
	"data": {
		"id": "6759e56a-fa7c-49ee-9854-f32ab38083ae",
		"username": "username17",
		"email": "email17@gmail.com",
		"level": "USER",
		"createdAt": "2024-06-16T10:46:25Z",
		"updatedAt": "2024-06-16T10:46:25Z"
	}, // object | array of object | null
	"validationErrors": [
		{
			"field": "username",
			"tag": "required",
			"message": "Key: 'RegisterReq.Username' Error:Field validation for 'username' failed on the 'required' tag"
		}
	] // array of object | null
}
```

## API Contract

### Register
#### Request

**Method:** `POST`

**URL:** `${{HOST}}/v1/auth/register`

**Headers:**
- `Content-Type: application/json`

**Body:**
```json
{
	"username":"username",
	"email": "email@gmail.com",
	"password": "strongpassword"
}
```

**Example cURL Command:**

```bash
curl --request POST \
  --url ${{HOST}}/v1/auth/register \
  --header 'Content-Type: application/json' \
  --data '{
	"username":"username",
	"email": "email@gmail.com",
	"password": "strongpassword"
}'
```

**Example Response:**
```json
{
	"message": "validation error",
	"data": null,
	"validationErrors": [
		{
			"field": "username",
			"tag": "unique_db",
			"message": "username already taken"
		},
		{
			"field": "email",
			"tag": "unique_db",
			"message": "email already taken"
		}
	]
}
```

```json
{
	"message": "ok",
	"data": null,
	"validationErrors": null
}
```

### Login
#### Request

**Method:** `POST`

**URL:** `${{HOST}}/v1/auth/login`

**Headers:**
- `Content-Type: application/json`

**Body:**
```json
{
	"identifier":"username",
	"password": "strongpassword"
}
```

**Example cURL Command:**

```bash
curl --request POST \
  --url ${{HOST}}/v1/auth/login \
  --header 'Content-Type: application/json' \
  --data '{
	"identifier":"username",
	"password": "strongpassword"
}'
```

**Example Response:**
```json
{
	"message": "user not found",
	"data": null,
	"validationErrors": null
}
```

```json
{
	"message": "ok",
	"data": {
		"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NzU5ZTU2YS1mYTdjLTQ5ZWUtOTg1NC1mMzJhYjM4MDgzYWUiLCJleHAiOjE3MTg2MTk3ODQsImlhdCI6MTcxODYxNjE4NCwianRpIjoiYzViZjAyZTctNmNkMy00MjZiLThiMjctYzk2MTUyZjc2NmU4In0.aaUAM7Hl6Z-H8kzdnrLedVmmVJEuglxes7xQYHt1HKI",
		"accessTokenExpiredAt": "2024-06-17T09:23:04Z",
		"refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NzU5ZTU2YS1mYTdjLTQ5ZWUtOTg1NC1mMzJhYjM4MDgzYWUiLCJleHAiOjE3MTg3MjQxODQsImlhdCI6MTcxODYxNjE4NCwianRpIjoiNzllNTVkZDgtMzczMS00OWU2LThjZDItNzMxNDI0MzYzZjZjIn0.KQOZGZxz8-8JiJv68Xpdj-7z1Dp6dLe0a4IC0nZ5WcA",
		"refreshTokenExpiredAt": "2024-06-17T09:23:04Z"
	},
	"validationErrors": null
}
```

### Refresh Token
#### Request

**Method:** `POST`

**URL:** `${{HOST}}/v1/auth/refresh`

**Headers:**
- `Content-Type: application/json`
- `Authorization: Bearer <REFRESH_TOKEN>`

**Example cURL Command:**

```bash
curl --request POST \
  --url ${{HOST}}/v1/auth/refresh \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <REFRESH_TOKEN>'
```

**Example Response:**
```json
{
	"message": "invalid token",
	"data": null,
	"validationErrors": null
}
```

```json
{
	"message": "ok",
	"data": {
		"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NzU5ZTU2YS1mYTdjLTQ5ZWUtOTg1NC1mMzJhYjM4MDgzYWUiLCJleHAiOjE3MTg2MTk3ODQsImlhdCI6MTcxODYxNjE4NCwianRpIjoiYzViZjAyZTctNmNkMy00MjZiLThiMjctYzk2MTUyZjc2NmU4In0.aaUAM7Hl6Z-H8kzdnrLedVmmVJEuglxes7xQYHt1HKI",
		"accessTokenExpiredAt": "2024-06-17T09:23:04Z",
		"refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiI2NzU5ZTU2YS1mYTdjLTQ5ZWUtOTg1NC1mMzJhYjM4MDgzYWUiLCJleHAiOjE3MTg3MjQxODQsImlhdCI6MTcxODYxNjE4NCwianRpIjoiNzllNTVkZDgtMzczMS00OWU2LThjZDItNzMxNDI0MzYzZjZjIn0.KQOZGZxz8-8JiJv68Xpdj-7z1Dp6dLe0a4IC0nZ5WcA",
		"refreshTokenExpiredAt": "2024-06-17T09:23:04Z"
	},
	"validationErrors": null
}
```

### Logout
#### Request

**Method:** `POST`

**URL:** `${{HOST}}/v1/auth/logout`

**Headers:**
- `Content-Type: application/json`
- `Authorization: Bearer <ACCESS_TOKEN>`

**Example cURL Command:**

```bash
curl --request POST \
  --url ${{HOST}}/v1/auth/logout \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <ACCESS_TOKEN>'
```

**Example Response:**
```json
{
	"message": "invalid token",
	"data": null,
	"validationErrors": null
}
```

```json
{
	"message": "ok",
	"data": null,
	"validationErrors": null
}
```

### User Info
#### Request

**Method:** `GET`

**URL:** `${{HOST}}/v1/auth/info`

**Headers:**
- `Content-Type: application/json`
- `Authorization: Bearer <ACCESS_TOKEN>`

**Example cURL Command:**

```bash
curl --request GET \
  --url ${{HOST}}/v1/auth/info \
  --header 'Authorization: Bearer <ACCESS_TOKEN>'
```

**Example Response:**
```json
{
	"message": "ok",
	"data": {
		"id": "6759e56a-fa7c-49ee-9854-f32ab38083ae",
		"username": "username",
		"email": "email@gmail.com",
		"level": "USER",
		"createdAt": "2024-06-16T10:46:25Z",
		"updatedAt": "2024-06-16T10:46:25Z"
	},
	"validationErrors": null
}
```
