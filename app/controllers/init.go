package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(Initialize)
	revel.InterceptMethod((*DatabaseController).Begin, revel.BEFORE)
	//revel.InterceptMethod((*Profile).Index, revel.BEFORE)
	revel.InterceptMethod((*DatabaseController).Commit, revel.AFTER)
	revel.InterceptMethod((*DatabaseController).Rollback, revel.FINALLY)
}
