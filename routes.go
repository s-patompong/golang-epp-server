package main

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
)

func routes() *routing.Router {
	router := routing.New()

	router.Use(setHeader, handleResponse)
	nominetRoutes(router)

	return router
}

func setHeader(c *routing.Context) error {
	c.SetContentType("application/json")
	return nil
}

func handleResponse(c *routing.Context) error {
	err := c.Next()

	if err != nil {
		_, _ = fmt.Fprintf(c, err.Error())
		return err
	}

	return nil
}
