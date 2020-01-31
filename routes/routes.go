package routes

import (
	"passive-dns/controllers"

	"github.com/gin-gonic/gin"
)

// InitRoutes is used to initialize server routes
func InitRoutes() *gin.Engine {
	router := gin.Default()
	controller := controllers.PassiveDomainsIpsController{}
	controller.Init()
	// router.PUT("/passive_domains_ips", controller.Update)
	router.POST("passive_domains_ips/search", controller.Search)
	return router
}
