package main

import (
	"net/http"

	"github.com/valentedev/elements/internal/data"
	"github.com/valentedev/elements/internal/validator"
)

func (app *application) listElementsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AtomicWeight string `json:"atomic_weight"`
		Name         string `json:"name"`
		Filters      data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.AtomicWeight = app.readString(qs, "atomic_weight", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	elements, metadata, err := app.models.Elements.GetAll(input.Name, input.AtomicWeight, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, mailbox{"elements": elements, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
