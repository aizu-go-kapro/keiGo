# keiGo

## How to dev API server

**Local run**

```
git clone https://github.com/aizu-go-kapro/keiGo.git
cd backend
go mod download
go run main.go
```

**Example API call**

Request: GET /keigo

```
curl --request GET \
  --url http://localhost:3000/api/v1/keigo \
  --header 'content-type: application/json' \
  --data '{
  "body": "寿司が食べたい。"
}'
```

Response:

```
{
  "converted_body": "寿司 が 食べ たい 。 "
}
```
