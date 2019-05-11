package aws

import (
	"net/http"
)

func (handler *AWSHandler) IAMRolesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_role")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMRoles(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:ListRoles is missing")
		} else {
			handler.cache.Set("aws_role", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) IAMGroupsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_group")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMGroups(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:ListGroups is missing")
		} else {
			handler.cache.Set("aws_group", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) IAMPoliciesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_policy")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMPolicies(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:ListPolicies is missing")
		} else {
			handler.cache.Set("aws_policy", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) IAMUsersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_users")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMUsers(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:ListUsers is missing")
		} else {
			handler.cache.Set("aws_users", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) IAMUserHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_user")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeIAMUser(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "iam:GetUser is missing")
		} else {
			handler.cache.Set("aws_user", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) DescribeOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_organization")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeOrganization(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "organizations:DescribeOrganization is missing")
		} else {
			handler.cache.Set("aws_organization", response)
			respondWithJSON(w, 200, response)
		}
	}
}
