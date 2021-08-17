package controllers

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	stdoutBuffer sync.Map
	//out          = gin.DefaultWriter
)

type ConsoleLine struct {
	Toke string `json:"toke" binding:"required"`
	Line string `json:"line" binding:"required"`
}

// GET /picobrains/tokers
func Tokers(c *gin.Context) {
	tokers := []string{}
	stdoutBuffer.Range(func(key, value interface{}) bool {
		tokers = append(tokers, key.(string))
		return true
	})
	fmt.Println("tokers: ", tokers)
	c.JSON(http.StatusOK, gin.H{"tokers": tokers})
}

// GET /picobrains/tokeme
func TokeMe(c *gin.Context) {
	toke := uuid.New()
	outbuffer := []string{}
	stdoutBuffer.Store(toke.String(), outbuffer)
	println("toke: ", toke.String())
	c.JSON(http.StatusOK, gin.H{"toke": toke.String()})
}

// POST /picobrains
func AddStdoutLine(c *gin.Context) {

	// Validate input
	var input ConsoleLine
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buf, ok := stdoutBuffer.Load(input.Toke)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "toke better"})
		return
	}

	tokBuf, _ := buf.([]string)

	tokBuf = append(tokBuf, input.Line)
	stdoutBuffer.Store(input.Toke, tokBuf)

	c.JSON(http.StatusOK, gin.H{"data": input.Line})
	gin.Logger()
}

// GET /picobrains/:toke
func GetStdoutBuffer(c *gin.Context) {

	toke := c.Param("toke")
	if len(toke) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "must toke"})
		return
	}

	buf, ok := stdoutBuffer.Load(toke)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no smoke in toke"})
		return
	}

	tokBuf, _ := buf.([]string)

	c.JSON(http.StatusOK, gin.H{"data": tokBuf})
}
