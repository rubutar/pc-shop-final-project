package handler

import (
	"context"
	_interface "pc-shop-final-project/domain/repository"
)

type UserHandler struct {
	ctx      context.Context
	repoUser _interface.InterfaceUser
}

func NewUserHandler(ctx context.Context, repoUser _interface.InterfaceUser) *UserHandler {
	return &UserHandler{
		ctx:      ctx,
		repoUser: repoUser,
	}
}

//
//func (usr *UserHandler) DeleteDataUser(w http.ResponseWriter, r *http.Request) {
//	var vars = mux.Vars(r)
//
//	dataUser, errGet := usr.repoUser.GetUserById(usr.ctx, vars["id"])
//	if errGet != nil {
//		respErr, _ := http_response.MapResponseBuku(nil, http.StatusInternalServerError, errGet.Error())
//		w.WriteHeader(http.StatusNotFound)
//		w.Write(respErr)
//		return
//	}
//
//	errDelete := bk.repoBuku.DeleteBukuByKode(bk.ctx, dataBuku.GetKodeBuku())
//	if errDelete != nil {
//		respErr, _ := http_response.MapResponseBuku(nil, http.StatusInternalServerError, errDelete.Error())
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write(respErr)
//		return
//	}
//
//	resp, errMap := http_response.MapResponseBuku(nil, http.StatusOK, "SUCCESS DELETE DATA")
//	if errMap != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(errMap.Error()))
//		return
//	}
//
//	w.WriteHeader(200)
//	w.Write(resp)
//	return
//}