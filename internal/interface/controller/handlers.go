package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/fedo3nik/GamePortal_ForumService/internal/application/service"
	dto "github.com/fedo3nik/GamePortal_ForumService/internal/interface/controller/dtohttp"
	e "github.com/fedo3nik/GamePortal_ForumService/internal/util/error"
)

type HTTPAddForumHandler struct {
	forumService service.Forum
}

type HTTPGetForumHandler struct {
	forumService service.Forum
}

func handleError(w http.ResponseWriter, err error) {
	if errors.Is(err, e.ErrDB) {
		_, hError := fmt.Fprintf(w, "Error caused: %v", err)
		if hError != nil {
			log.Printf("Fprint error: %v", hError)
		}

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, hError := fmt.Fprintf(w, "Internal server error: %v", err)
	if hError != nil {
		log.Printf("Fprint error: %v", hError)
	}

	w.WriteHeader(http.StatusInternalServerError)
}

func NewHTTPAddForumHandler(forumService service.Forum) *HTTPAddForumHandler {
	return &HTTPAddForumHandler{forumService: forumService}
}

func (hh HTTPAddForumHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req dto.AddForumRequest

	var resp dto.AddForumResponse

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Body read error: %v", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	forum, err := hh.forumService.AddForum(r.Context(), req.Title, req.Topic, req.Text, req.Token)
	if err != nil {
		handleError(w, err)
		return
	}

	resp.ID = forum.ID
	resp.UserID = forum.UserID
	resp.Topic = forum.Topic
	resp.Title = forum.Title

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		log.Printf("Encode error: %v", err)
		return
	}
}

func NewHTTPGetForumHandler(forumService service.Forum) *HTTPGetForumHandler {
	return &HTTPGetForumHandler{forumService: forumService}
}

func (hh HTTPGetForumHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var resp dto.GetForumResponse

	url := r.URL.Path
	idString := path.Base(url)

	forum, err := hh.forumService.GetForum(r.Context(), idString)
	if err != nil {
		handleError(w, err)
		return
	}

	resp.ID = forum.ID
	resp.Title = forum.Title
	resp.Topic = forum.Topic
	resp.UserID = forum.UserID
	resp.Text = forum.Text

	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		handleError(w, err)
		return
	}
}
