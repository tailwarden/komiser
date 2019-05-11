package aws

import (
	"net/http"
)

func (handler *GCPHandler) IAMRolesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_iam_roles")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetIamRoles()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_iam_roles", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) IAMUsersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_iam_users")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetIamUsers()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_iam_users", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) IAMServiceAccountsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_iam_service_accounts")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetServiceAccounts()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_iam_service_accounts", response)
			respondWithJSON(w, 200, response)
		}
	}
}
