package controllers

import (
	"encoding/json"
	"github.com/ArturoAguilar1/b_items-api/domain/items"
	"github.com/ArturoAguilar1/b_items-api/services"
	"github.com/ArturoAguilar1/b_items-api/utils/http_utils"
	"github.com/federicoleon/bookstore_oauth-go/oauth"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"io/ioutil"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {}


func (cont *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w,err)
		return
	}

	var itemRequest items.Item
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close() //Closing the body under the request
	if err := json.Unmarshal(requestBody,&itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = oauth.GetClientId(r)
	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w,createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)

}


func (cont *itemsController)Get(w http.ResponseWriter, r *http.Request) {}
