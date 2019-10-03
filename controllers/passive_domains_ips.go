package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/go-pg/pg"

	"passive-dns/types"

	"passive-dns/dispatcher"

	"passive-dns/db"

	"passive-dns/models"
	// "encoding/json"
)

type updateBody struct {
	Source          string                `json:"source"`
	ResolvedEntries []types.ResolvedEntry `json:"resolved_entries"`
}

// PassiveDomainsIpsController is used to handler request from url: /passive_domains_ips
type PassiveDomainsIpsController struct {
	eleDispatcher     dispatcher.ElementDispatcher
	passiveDispatcher dispatcher.PassiveDispatcher
	db                *pg.DB
}

// Init is used to initialize controller
func (controller *PassiveDomainsIpsController) Init() error {
	var err error
	controller.eleDispatcher.Init()
	controller.passiveDispatcher.Init()
	controller.db, err = db.GetDB()
	return err
}

// Update is used to update passive ips and domains
func (controller *PassiveDomainsIpsController) Update(c *gin.Context) {
	fmt.Printf("content type is %s\n", c.ContentType())
	var body updateBody
	c.BindJSON(&body)
	var source models.Source
	fmt.Println("db", controller.db)
	controller.db.Model(&source).Where("name = ?", body.Source).Select()
	if source.Name != "" {
		fmt.Println("source: ", source.Name)
		// need error check
		err := controller.eleDispatcher.Refresh(body.Source, body.ResolvedEntries)
		if err == nil {
			// dispatch resolved entries
			err = controller.passiveDispatcher.Refresh(source.ID, body.ResolvedEntries)
		}
	}

	// jsonBytes, _ := json.Marshal(body)
	// fmt.Println(string(jsonBytes))
}

// Search is used to find the match passive dns information
func (controller *PassiveDomainsIpsController) Search(c *gin.Context) {
	// need implement
}
