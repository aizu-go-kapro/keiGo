# backend

## How to dev API server

**Local run**

```
git clone https://github.com/aizu-go-kapro/keiGo.git
cd backend
go mod download
go run main.go
```

**Example API call**

Request: POST /keigo

```
curl --request POST \
  --url 'http://localhost:3000/api/v1/keigo?kind=teinei' \
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

## Testing
**Unit Test convert logic**

Testing all.
```
go test ./... -v -cover
```
