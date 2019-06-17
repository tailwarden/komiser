package ovh

import (
	"net/http"
)

func (handler *OVHHandler) DescribeCloudProjectsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.projects")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetProjects()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.projects", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudInstancesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.instances")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetInstances()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.instances", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudStorageContainersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.storage.containers")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetStorageContainers()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.storage.containers", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudUsersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.users")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetUsers()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.users", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudVolumesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.volumes")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetVolumes()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.volumes", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudSnapshotsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.snapshots")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetSnapshots()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.snapshots", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudAlertsandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.alerts")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetAlerts()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.alerts", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCurrentUsageHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.bill.current")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetCurrentUsage()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.bill.current", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudImagesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.images")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetImages()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.images", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudIpsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.ips")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetIps()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.ips", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudPrivateNetworksHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.networks.private")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetPrivateNetworks()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.networks.private", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudPublicNetworksHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.networks.public")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetPublicNetworks()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.networks.public", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudFailoverIpsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.ip.failover")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetFailoverIps()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.ip.failover", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudVRacksHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.vrack")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetVRacks()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.vrack", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudKubeClustersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.kube.clusters")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetKubeClusters()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.kube.clusters", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudKubeNodesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.kube.nodes")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetKubeNodes()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.kube.nodes", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudSSHKeysHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.sshkeys")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetSSHKeys()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.sshkeys", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeCloudLimitsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.cloud.quotas")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetLimits()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.cloud.quotas", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeProfileHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.profile")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetProfile()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.profile", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *OVHHandler) DescribeTicketsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ovh.tickets")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.ovh.GetTicketsStats()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		} else {
			handler.cache.Set("ovh.tickets", response)
			respondWithJSON(w, 200, response)
		}
	}
}
