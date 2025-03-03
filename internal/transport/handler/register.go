package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/crm/internal/data/models"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
    if err != nil{
        http.Error(w, "cant read request body", http.StatusBadRequest)
        return
    }

    var user models.User

    if err := sonic.Unmarshal(body, &user);err != nil{
        http.Error(w, "cant read request body", http.StatusBadRequest)
        return
    }

    token, id, err := h.service.CreateUser(&user)
    if err != nil{
        http.Error(w, "cant create user", http.StatusInternalServerError)
        return
    }

    fmt.Fprint(w, map[string]string{
        "id" : fmt.Sprintf("%d", id),
        "token" : token,
    })
}