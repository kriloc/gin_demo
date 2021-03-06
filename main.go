// gin demo from b站
package main

import (
	"context"
	handlers "gin_demo/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
	//"net/http"
	//"strings"
)

var ctx context.Context
var err error
var client *mongo.Client
var recipesHandler *handlers.RecipesHandler
var authHandler *handlers.AuthHandler

func init() {
	//recipes = make([]Recipe, 0)
	//file, _ := ioutil.ReadFile("recipes.json")
	//_ = json.Unmarshal([]byte(file), &recipes)

	// mongodb init
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.212.111:27017/test"))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Rpi4 MongoDB")

	collection := client.Database("demo").Collection("recipes")

	// insert json data to mongodb
	// 有資料就不再輸入了
	//var listOfRecipes []interface{}
	//for _, recipe := range recipes {
	//	listOfRecipes = append(listOfRecipes, recipe)
	//}
	//collection := client.Database("demo").Collection("recipes")

	//insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//log.Println("Inserted recipes:", len(insertManyResult.InsertedIDs))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.212.111:6379",
		Password: "",
		DB:       0,
	})
	status := redisClient.Ping()
	log.Println(status, "Connected to Rpi4 Redis")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection, redisClient)

	collectionUsers := client.Database("demo").Collection("users")
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
}

func main() {
	//createUsers()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, MaxAge: 12 * time.Hour,
	}))
	//r.GET("/hello", func(c *gin.Context) {
	//	c.String(http.StatusOK, "OK")
	//})
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/refresh", authHandler.RefreshHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	//router.POST("/recipes", recipesHandler.NewRecipeHandler)

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
		authorized.GET("/recipes/:id", recipesHandler.GetOneRecipeHandler)
	}

	router.Run()
	//router.GET("/recipes/search", SearchRecipesHandler)

	//http://localhost:8080/recipes/61543159f84b94bb7be3de8e

}

//func createUsers(){
//	users := map[string]string{
//		"admin":"fCRmh4Q2J7Rseqkz",
//		"packt":"RE4zfHB35VPtTkbT",
//		"mlabouardy": "L3nSFRcZzNQ67bcc",
//	}
//	ctx = context.Background()
//	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.212.111:27017/test"))
//	if err = client.Ping(context.TODO(),readpref.Primary()); err != nil{
//		log.Fatal(err)
//	}
//	collection := client.Database("demo").Collection("users")
//	h := sha256.New()
//	for username, password := range users {
//
//		collection.InsertOne(ctx, bson.M{
//
//			"username": username,
//
//			"password": string(h.Sum([]byte(password))),
//		})
//	}
//	log.Println("createed users...")
//}
//func NewRecipeHandler(c *gin.Context){
//	var recipe Recipe
//	if err := c.ShouldBindJSON(&recipe); err!=nil{
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error()})
//		return
//	}
//	//recipe.ID = xid.New().String()
//	//recipe.PublishedAt = time.Now()
//	//recipes = append(recipes, recipe)
//	recipe.ID = primitive.NewObjectID()
//
//	c.JSON(http.StatusOK, recipe)
//}

//func ListRecipesHandler(c *gin.Context){
//	//c.JSON(http.StatusOK, recipes)
//	collection := client.Database("demo").Collection("recipes")
//	cur, err := collection.Find(ctx, bson.M{})
//	if err != nil{
//		c.JSON(http.StatusInternalServerError,
//			gin.H{"error": err.Error()})
//		return
//	}
//	defer cur.Close(ctx)
//
//	recipes := make([]Recipe, 0)
//	for cur.Next(ctx){
//		var recipe Recipe
//		cur.Decode(&recipe)
//		recipes = append(recipes, recipe)
//	}
//	c.JSON(http.StatusOK, recipes)
//}

//func UpdateRecipeHandler(c *gin.Context){
//	id := c.Param("id")  // 取得URL的參數
//	var recipe Recipe
//	if err:= c.ShouldBindJSON(&recipe); err !=nil{
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error()})
//		return
//	}
//	index := -1
//	for i:=0; i<len(recipes); i++{
//		if recipes[i].ID == id{
//			index = 1
//		}
//	}
//	if index == -1{
//		c.JSON(http.StatusNotFound, gin.H{
//			"error": "Recipe not found"	})
//		return
//	}
//	recipes[index] = recipe
//	c.JSON(http.StatusOK, recipe)
//}

//func DeleteRecipeHandler(c *gin.Context) {
//	id := c.Param("id")
//	index := -1  // 初始化，不能為0，所以是-1
//	for i:=0 ; i<len(recipes); i++{
//		if recipes[i].ID == id {
//			index = i
//		}
//	}
//
//	if index == -1 {  // 表示找不到
//		c.JSON(http.StatusNotFound, gin.H{
//			"error": "Recipe not found."})
//		return
//	}
//
//	recipes = append(recipes[:index], recipes[index+1:]...) // 將此index之前跟之後加回去
//	c.JSON(http.StatusOK, gin.H{
//		"message":"Recipe has been deleted."})
//
//	// test
//	//curl -v -sX DELETE http://localhost:8080/recipes/c0283p3d0cvuglq85log
//}

//func SearchRecipesHandler(c *gin.Context) {
//	tag := c.Query("tag")
//	listOfRecipes := make([]Recipe, 0)
//
//	for i := 0; i < len(recipes); i++ {
//		found := false
//		for _, t := range recipes[i].Tags {
//			if strings.EqualFold(t, tag){
//				found = true
//			}
//		}
//		if found {
//			listOfRecipes = append(listOfRecipes, recipes[i])
//		}
//	}
//	c.JSON(http.StatusOK, listOfRecipes)
//	//http://localhost:8080/recipes/search?tag=italian
//}
