package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/joisandresky/go-chi-clean-starter/internal/application/dto"
	"github.com/joisandresky/go-chi-clean-starter/internal/application/usecases"
	"github.com/joisandresky/go-chi-clean-starter/pkg/guy"
)

type PostHttpApi interface {
	RegisterRoutes(r *chi.Mux)
}

type postHttpApi struct {
	uc usecases.PostUsecase
}

func NewPostHttpApi(uc usecases.PostUsecase) PostHttpApi {
	return &postHttpApi{uc: uc}
}

func (api *postHttpApi) RegisterRoutes(r *chi.Mux) {
	r.Route("/api/v1/posts", func(r chi.Router) {
		r.Get("/", api.GetAll)
		r.Get("/{id}", api.GetById)
		r.Post("/", api.CreatePost)
		r.Put("/{id}", api.UpdatePostByid)
		r.Delete("/{id}", api.DeletePostById)
	})
}

func (api *postHttpApi) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := api.uc.GetAll(r.Context())
	if err != nil {
		guy.HandleError(w, r, err)
		return
	}

	guy.Ok(w, r, "", posts)
}

func (api *postHttpApi) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	post, err := api.uc.GetById(r.Context(), id)
	if err != nil {
		guy.HandleError(w, r, err)
		return
	}

	guy.Ok(w, r, "", post)
}

func (api *postHttpApi) CreatePost(w http.ResponseWriter, r *http.Request) {
	req := dto.CreatePost{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		guy.BadRequest(w, r, "failed to get request body", err)
		return
	}

	if err := guy.Validate(req); err != nil {
		guy.BadRequest(w, r, "Validation errors", err)
		return
	}

	if err := api.uc.Create(r.Context(), &req); err != nil {
		guy.HandleError(w, r, err)
		return
	}

	guy.Created(w, r, "post created successfully", nil)
}

func (api *postHttpApi) UpdatePostByid(w http.ResponseWriter, r *http.Request) {
	req := dto.CreatePost{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		guy.BadRequest(w, r, "failed to get request body", err)
		return
	}

	if err := guy.Validate(req); err != nil {
		guy.BadRequest(w, r, "Validation errors", err)
		return
	}

	if err := api.uc.UpdateById(r.Context(), chi.URLParam(r, "id"), &req); err != nil {
		guy.HandleError(w, r, err)
		return
	}

	guy.Ok(w, r, "post updated successfully", nil)
}

func (api *postHttpApi) DeletePostById(w http.ResponseWriter, r *http.Request) {
	if err := api.uc.DeleteById(r.Context(), chi.URLParam(r, "id")); err != nil {
		guy.HandleError(w, r, err)
		return
	}

	guy.Ok(w, r, "post deleted successfully", nil)
}
