package crowdapiv1

import (
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"strconv"
)

func NewUploadInputDataHandler(service projectService) runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		projectIDstr := pathParams["id"]
		projectID, err := strconv.ParseInt(projectIDstr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("failed to get project id: %s", err.Error())))
			return
		}
		file, _, err := r.FormFile("file")
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("failed to get file data: %s", err.Error())))
			return
		}

		err = service.UploadProjectData(r.Context(), int(projectID), file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("failed to upload project data: %s", err.Error())))
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
