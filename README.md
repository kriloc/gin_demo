# Building.Distributed.Applications.in.Gin notes

##＃　distributed microservices

測試語句：

`
curl --location --request POST 'http://localhost:8080/recipes' --header 'Content-Type: application/json' --data-raw '{ "name": "Homemade Pizza", "tags" : ["italian", "pizza", "dinner"], "ingredients": ["1 1/2 cups (355 ml) warm water (105°F-115°F)","1 package (2 1/4 teaspoons) of active dry yeast","3 3/4 cups (490 g) bread flour","feta cheese, firm mozzarella cheese, grated"],"instructions": ["Step 1.","Step 2.","Step 3."]}' | jq -r
`

update:

```
post http://localhost:8080/recipes/
Body / RAW / json
{
  "name": "Homemade Pepperoni Pizza",
  "tags": ["italian", "pizza", "dinner"],
  "ingredients": [
        "pizza dough",
        "1/2 tsp salt",
        "2 tablespoon extra-virgin olive oil",
        "1 lemon, juiced"
    ],
  "instructions": ["step1", "step2"]
}

然後取得 _id 後，再 update data
put http://localhost:8080/recipes/600dcc85a65917cbd1f201b0
更改其中的 "name":"xxxxxxx"
```

出現:
`missing go.sum entry for module providing package`

使用`go mod tidy`

go get go.mongodb.org/mongo-driver/mongo
go get github.com/go-redis/redis/v8
go get github.com/dgrijalva/jwt-go

## Auth: JWT
先 POST http://localhost:8080/signin
其中Body
```json
{
    "username":"admin",
    "password":"fCRmh4Q2J7Rseqkz"
}
```
會取得一token

新增食譜 POST http://localhost:8080/recipes
資料如同前面範例，但Headers欄位新增：
KEY:Authorization  
VALUE: 剛剛取得的 token