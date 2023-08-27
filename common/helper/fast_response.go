package helper

import (
	"github.com/jun-chiang/simple-douyin/common/consts"
	"github.com/jun-chiang/simple-douyin/common/models"
)

func GetSuccesReponse(msg string) (resp *models.BasicResponse) {
	resp = &models.BasicResponse{
		StatusCode: consts.StatusOk,
		StatusMsg:  msg,
	}
	return
}

func GetFailReponse(msg string) (resp *models.BasicResponse) {
	resp = &models.BasicResponse{
		StatusCode: consts.StatusError,
		StatusMsg:  msg,
	}
	return
}
