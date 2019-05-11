package aws

import (
	"net/http"
)

func (handler *GCPHandler) ComputeInstancesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_compute_instances")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetComputeInstances()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("gcp_compute_instances", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) ComputeDisksHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_compute_disks")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetDisks()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("gcp_compute_disks", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) DiskSnapshotsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_disk_snapshots")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetDiskSnapshots()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("gcp_disk_snapshots", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) ComputeImagesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_compute_images")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetComputeImages()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("gcp_compute_images", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) ComputeCPUUtilizationHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_compute_cpu_utilization")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetComputeCPUUtilization()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "monitoring:MonitoringReadScope is missing")
		} else {
			handler.cache.Set("gcp_compute_cpu_utilization", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) ComputeQuotasHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_compute_quotas")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetQuotas()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "compute:ComputeReadonlyScope is missing")
		} else {
			handler.cache.Set("gcp_compute_quotas", response)
			respondWithJSON(w, 200, response)
		}
	}
}
