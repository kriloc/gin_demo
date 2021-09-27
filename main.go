// gin demo from b站
package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"time"
)

type Recipe struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init(){
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func NewRecipeHandler(c *gin.Context){
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func ListRecipesHandler(c *gin.Context){
	c.JSON(http.StatusOK, recipes)
}

func UpdateRecipeHandler(c *gin.Context){
	id := c.Param("id")  // 取得URL的參數
	var recipe Recipe
	if err:= c.ShouldBindJSON(&recipe); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	index := -1
	for i:=0; i<len(recipes); i++{
		if recipes[i].ID == id{
			index = 1
		}
	}
	if index == -1{
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"	})
		return
	}
	recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	//r.GET("/hello", func(c *gin.Context) {
	//	c.String(http.StatusOK, "OK")
	//})
	router.POST("/recipes",NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("recipes/:id", DeleteRecipeHandler)
	router.Run()

}

func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	index := -1  // 初始化，不能為0，所以是-1
	for i:=0 ; i<len(recipes); i++{
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {  // 表示找不到
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found."})
		return
	}

	recipes = append(recipes[:index], recipes[index+1:]...) // 將此index之前跟之後加回去
	c.JSON(http.StatusOK, gin.H{
		"message":"Recipe has been deleted."})

	// test
	//curl -v -sX DELETE http://localhost:8080/recipes/c0283p3d0cvuglq85log
}
