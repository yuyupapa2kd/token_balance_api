package api

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/the-medium/token-balance-api/docs"
	"github.com/the-medium/token-balance-api/internal/config"
	"github.com/the-medium/token-balance-api/internal/core"
	"github.com/the-medium/token-balance-api/internal/resource"
)

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func setupSwagger(r *gin.Engine) {

	localAddr := getOutboundIP()
	IPNPortString := localAddr.String() + ":" + config.RuntimeConf.Server.Port

	docs.SwaggerInfo.Title = "GNDChain Token Balance API"
	docs.SwaggerInfo.Description = "This is a temporary Tool for BoT on GNDChain."
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.Host = IPNPortString
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	swaggerUrlString := "http://" + IPNPortString + "/swagger/doc.json"
	fmt.Println("swaggerUrlString : ", swaggerUrlString)
	// url := ginSwagger.URL("http://localhost:4000/swagger/doc.json")
	url := ginSwagger.URL(swaggerUrlString)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func SetRouter() *gin.Engine {
	r := gin.Default()

	setupSwagger(r)

	r.GET("/coin/:address", getCoinBalance)
	r.GET("/token/:symbol/:address", getTokenBalance)

	return r
}

// TxDecoder 	godoc
// @Tags         Coin
// @Summary      Get Coin Balance
// @Description  Get Coin Balance of given Address
// @Produce      json
// @Param        address  path      string  true  "address"
// @Success      200      {object}  resource.ResJSON{data=resource.ResBalance}
// @Failure      400      {object}  resource.ResJSON{data=resource.ResErr}
// @Router       /coin/{address} [get]
func getCoinBalance(c *gin.Context) {
	address := c.Param("address")
	fmt.Println("param : ", address)
	var res resource.ResBalance
	res, err := core.GetCoinBalance(address)
	if err != nil {
		c.JSON(400, gin.H{"result": res, "error": string(err.Error())})
		return
	}
	c.JSON(200, gin.H{"result": res, "error": ""})
}

// TxDecoder	 godoc
// @Tags         Token
// @Summary      Get Token Balance of given Address
// @Description  Allowed Token Symbol : SOP, LOUI, ksETH, ksUSDT
// @Produce      json
// @Param        symbol   path      string  true  "symbol"
// @Param        address  path      string  true  "address"
// @Success      200      {object}  resource.ResJSON{data=resource.ResBalance}
// @Failure      400      {object}  resource.ResJSON{data=resource.ResErr}
// @Router       /token/{symbol}/{address} [get]
func getTokenBalance(c *gin.Context) {
	symbol := c.Param("symbol")
	address := c.Param("address")
	var res resource.ResBalance
	res, err := core.GetTokenBalance(symbol, address)
	if err != nil {
		c.JSON(400, gin.H{"result": res, "error": string(err.Error())})
		return
	}
	c.JSON(200, gin.H{"result": res, "error": ""})
}
