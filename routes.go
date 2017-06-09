package main

import (
	"github.com/qiangxue/fasthttp-routing"
)

func routes() *routing.Router {
	router := routing.New()

	router.Use(setHeader)
	nominetRoutes(router)

	return router
}

func setHeader(c *routing.Context) error {
	c.SetContentType("application/json")
	return nil
}
