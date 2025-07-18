package types

// HTTPResponse represents a generic HTTP response
type HTTPResponse struct {
	StatusCode int
	Body       []byte
	Error      error
}

// AuthHeader represents authentication headers for API requests
type AuthHeader map[string]string

// APIError represents an error from an API call
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (e APIError) Error() string {
	if e.Details != "" {
		return e.Message + ": " + e.Details
	}
	return e.Message
}
