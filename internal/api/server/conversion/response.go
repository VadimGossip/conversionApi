package conversion

import "time"

type HealthCheckResponse struct {
	Status       string    `json:"status"`
	AppStartedAt time.Time `json:"app_started_at"`
}
