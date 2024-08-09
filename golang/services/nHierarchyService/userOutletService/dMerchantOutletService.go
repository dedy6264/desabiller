package useroutletservice

import (
	"desabiller/configs"
	"desabiller/helpers"
	"desabiller/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (svc nHierarchyUserOutletServices) DropUserOutletService(ctx echo.Context) error {
	var (
		svcName = "DropUserOutletService"
		logErr  = "Error " + svcName
		req     models.ReqGetListNUserOutlet
	)
	_, err := helpers.BindValidate(&req, ctx)
	if err != nil {
		log.Println(logErr+"BINDING", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "FAILLED "+err.Error(), nil)
		return ctx.JSON(http.StatusOK, result)
	}
	if req.ID == 0 {
		log.Println(logErr + " id cannot be empty")
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.VALIDATE_ERROR_CODE, "id cannot be empty ", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	_, err = svc.services.ApiNHierarchy.NDropUserOutlet(req.ID)
	if err != nil {
		log.Println(logErr+"NDropUserOutlet", err.Error())
		result := helpers.ResponseJSON(configs.FALSE_VALUE, configs.DB_NOT_FOUND, "Check your data", nil)
		return ctx.JSON(http.StatusOK, result)
	}
	result := helpers.ResponseJSON(configs.TRUE_VALUE, configs.SUCCESS_CODE, "Success", true)
	return ctx.JSON(http.StatusOK, result)
}
