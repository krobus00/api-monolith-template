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
	}, // object | null
	"validationErrors": [
		{
			"field": "username",
			"tag": "required",
			"message": "Key: 'RegisterReq.Username' Error:Field validation for 'username' failed on the 'required' tag"
		}
	] // array of object | null
}
```
