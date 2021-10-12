package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


type mockPokemonController struct {}

func (mpc *mockPokemonController) GetPokemons() gin.HandlerFunc {
  gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return func(ctx *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	}
}
func (mpc *mockPokemonController) CreatePokemons() gin.HandlerFunc {
  gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return func(ctx *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	}
}

func newPokemonController(t *testing.T) *mockPokemonController{
  return &mockPokemonController{}
}

func TestPokemonControllerGet(t *testing.T){
  pCtrl := newPokemonController(t)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	r.GET("/pokemon", pCtrl.GetPokemons())

	c.Request, _ = http.NewRequest(http.MethodGet, "/pokemon", bytes.NewBuffer([]byte("{}")))

  fmt.Println(c.Request)
  fmt.Println(w)

  r.ServeHTTP(w, c.Request)

	assert.EqualValues(t, w.Code, http.StatusOK)
}

const (
  ID = "151"
  NAME = "MEW"
)


func getPokemonPayload() string {
  params:= url.Values{}
  params.Add("id", ID)
  params.Add("name", NAME)

  return params.Encode()
}

func TestPokemonControllerPost(t *testing.T){
  var pCtrl = newPokemonController(t)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

  r.POST("/pokemon", pCtrl.CreatePokemons(), nil)

  pkmPayload := getPokemonPayload()

  req, _ := http.NewRequest(http.MethodPost, "/pokemon", bytes.NewBuffer([]byte("{}")))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(pkmPayload)))

  r.ServeHTTP(w,req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

  p, err := ioutil.ReadAll(w.Body)

  if err != nil{
    t.Fail()
  }

  fmt.Println(p)
}