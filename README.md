# Monolith Service

## Response
```json
{
	"message": "validation error", // string
	"data": null, // object
	"error": null, // object
	"validationErrors": [
		{
			"field": "username",
			"tag": "required",
			"message": "Key: 'RegisterReq.Username' Error:Field validation for 'username' failed on the 'required' tag"
		}
	] // array of object
}
```
