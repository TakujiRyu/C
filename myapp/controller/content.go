package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"

	"github.com/gorilla/mux"
)

func AddContent(w http.ResponseWriter, r *http.Request) {
	var cont model.Content
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cont)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	saveErr := cont.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	// no error
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"Status": "Content Added"})
}

func GetContent(w http.ResponseWriter, r *http.Request) {
	// get url parameter
	ctid := mux.Vars(r)["cid"]
	cont := model.Content{ContentId: ctid}
	getErr := cont.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Content not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, cont)
	}
	// fmt.Println(stud)
}

func UpdateContent(w http.ResponseWriter, r *http.Request) {
	ctid := mux.Vars(r)["cid"]
	var cont model.Content
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cont)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid body")
		return
	}
	updateErr := cont.Update(ctid)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Content Not Found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, cont)
	}
}

func DeleteContent(w http.ResponseWriter, r *http.Request) {
	ctid := mux.Vars(r)["cid"]
	c := model.Content{ContentId: ctid}
	if err := c.Delete(); err != nil {
		switch err {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusBadRequest, "Content has been either deleted or doesnt exist")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		// return
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
	}
}

func GetAllContent(w http.ResponseWriter, r *http.Request) {
	conts, getErr := model.GetAllContent()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, conts)
}
