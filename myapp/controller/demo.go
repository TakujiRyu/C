package controller

// import (
// 	"encoding/json"
// 	"myapp/model"
// 	"myapp/utils/httpResp"
// 	"net/http"
// )

// func AddContent(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintln(w, "Student data successfully saved")
// 	var stud model.Student
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&stud)
// 	if err != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
// 		return
// 	}
// 	defer r.Body.Close()
// 	saveErr := stud.Create()
// 	if saveErr != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
// 		return
// 	}
// 	// no error
// 	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Student Data Added"})
// }
