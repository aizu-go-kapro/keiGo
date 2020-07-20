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
  --url 'http://localhost:8080/api/v1/keigo?kind=teinei' \
  --header 'content-type: application/json' \
  --data '{
  "body": "私は寿司が食べたい。"
}'
```

Response:

```
{
  "converted_body": "私は寿司が食べたいです。"
}
```

**Test convert logic**

Testing all.
```
go test ./... -v -cover
```
