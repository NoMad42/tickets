package v1

import (
	"net/http"

	"github.com/google/uuid"

	"homework/specs"
)

func (a apiServer) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	user := &specs.UserProfile{
		AvatarUrl: "https://example.com/avatar.jpg",
		Id:        uuid.New().String(),
		Login:     "admin",
	}

	a.writeSuccessResponse(user, w)
}
