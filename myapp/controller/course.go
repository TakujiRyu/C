package controller

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"myapp/model"
// 	"myapp/utils/httpResp"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func AddCourse(w http.ResponseWriter, r *http.Request) {
// 	var course model.Course
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&course)
// 	if err != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
// 		return
// 	}
// 	defer r.Body.Close()
// 	saveErr := course.Create()
// 	if saveErr != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
// 		return
// 	}
// 	// no error
// 	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"Status": "Course Data Added"})
// }

// func GetCourse(w http.ResponseWriter, r *http.Request) {
// 	// get url parameter
// 	csid := mux.Vars(r)["cid"]
// 	course := model.Course{CourseId: csid}
// 	getErr := course.Read()
// 	if getErr != nil {
// 		switch getErr {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusNotFound, "Course not found")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
// 		}
// 	} else {
// 		httpResp.RespondWithJSON(w, http.StatusOK, course)
// 	}
// 	// fmt.Println(stud)
// }

// func UpdateCourse(w http.ResponseWriter, r *http.Request) {
// 	csid := mux.Vars(r)["cid"]
// 	var course model.Course
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&course)
// 	if err != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid body")
// 		return
// 	}
// 	updateErr := course.Update(csid)
// 	if updateErr != nil {
// 		switch updateErr {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusNotFound, "Course Not Found")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
// 		}
// 	} else {
// 		httpResp.RespondWithJSON(w, http.StatusOK, course)
// 	}
// }

// func DeleteCourse(w http.ResponseWriter, r *http.Request) {
// 	csid := mux.Vars(r)["cid"]
// 	c := model.Course{CourseId: csid}
// 	if err := c.Delete(); err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusBadRequest, "Course has been either deleted or doesnt exist")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		// return
// 	} else {
// 		httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
// 	}
// }

// func GetAllCourse(w http.ResponseWriter, r *http.Request) {
// 	courses, getErr := model.GetAllCourses()
// 	if getErr != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusOK, courses)
// }
