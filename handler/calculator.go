package handler 
import ( 
"github.com/gin-gonic/gin" 
"net/http" 
"strconv"
 ) 
func AddHandler(c *gin.Context) {
	 a, err := strconv.ParseFloat(c.PostForm("a"), 64)
	  b, err2 := strconv.ParseFloat(c.PostForm("b"), 64) 
	  if err != nil || err2 != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		 return 
		 } 
		 result := a + b
		  c.JSON(http.StatusOK, gin.H{"result": result})
}
