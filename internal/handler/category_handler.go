package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/wiliamvj/api-users-golang/internal/dto"
	"github.com/wiliamvj/api-users-golang/internal/handler/httperr"
	"github.com/wiliamvj/api-users-golang/internal/handler/validation"
)

// Create category
//
//	@Summary		Create new category
//	@Description	Endpoint for create category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateCategoryDto	true	"Create category dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/category [post]
func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "categoryhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "categoryhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "categoryhandler"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.categoryService.CreateCategory(r.Context(), req)
	if err != nil {
		slog.Error(fmt.Sprintf("error to create category: %v", err), slog.String("package", "categoryhandler"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}

//	 Search categories
//		@Summary		Search categories
//		@Description	Endpoint for search product
//		@Tags			categories
//		@Accept			json
//		@Produce		json
//		@Param			body	body		dto.FindProductDto	true	"Search categories"	true
//		@Success		200		{object}	response.ProductResponse
//		@Failure		400		{object}	httperr.RestErr
//		@Failure		500		{object}	httperr.RestErr
//		@Router			/categories [get]
func (h *handler) FindManyCategories(w http.ResponseWriter, r *http.Request) {
	res, err := h.categoryService.FindManyCategories(r.Context())
	if err != nil {
		slog.Error(fmt.Sprintf("error to find many users: %v", err), slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to find many users")
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
