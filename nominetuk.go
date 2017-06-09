package main

import (
	"errors"
	"fmt"
	"github.com/pangpondpon/golang-epp/nominetuk"
	"github.com/qiangxue/fasthttp-routing"
)

func nominetUkRoutes(nominet *routing.RouteGroup) {
	nominet.Get("/uk/info", nominetUkDomainInfo)
}

func nominetUkDomainInfo(c *routing.Context) error {
	domainName := string(c.FormValue("name")) // Use scottishslategifts.co.uk for testing
	if domainName == "" {
		return errors.New("No domain name")
	}

	nominetUk, err := nominetuk.NewNominetUk()
	if err != nil {
		return err
	}

	domain := nominetUk.NewDomain()

	info, err := domain.Info(domainName)
	if err != nil {
		return err
	}

	fmt.Fprintf(c, response(info))
	return nil
}

var handler routing.Handler

func requestHandler(c *routing.Context) {

}
