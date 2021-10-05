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

