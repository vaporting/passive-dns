package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"passive-dns/types"

	"encoding/json"
)

type updateBody struct {
	Source          string                `json:"source"`
	ResolvedEntries []types.ResolvedEntry `json:"resolved_entries"`
}

// PassiveDomainsIpsController is used to handler request from url: /passive_domains_ips
type PassiveDomainsIpsController struct {
}

// Update is used to update passive ips and domains
func (controller *PassiveDomainsIpsController) Update(c *gin.Context) {
	fmt.Printf("content type is %s\n", c.ContentType())
	var body updateBody
	c.BindJSON(&body)
	jsonBytes, _ := json.Marshal(body)
	fmt.Println(string(jsonBytes))
	// check source is existed or not
	// if existed send to dispatcher
}

// Search is used to find the match passive dns information
func (controller *PassiveDomainsIpsController) Search(c *gin.Context) {

}
