package model

type CreateOptions struct {
	// DryRun when set, the object will be created in "dry-run" mode. This means the object will not be persisted.
	DryRun bool `json:"dryRun"`
}

type DeleteOptions struct {
	// DryRun when set, the object will be deleted in "dry-run" mode. This means the object will not be deleted.
	DryRun bool `json:"dryRun"`

	// GracePeriodSeconds the duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used.
	GracePeriodSeconds int64 `json:"gracePeriodSeconds"`
}

type GetOptions struct {
	// DryRun when set, the object will be fetched in "dry-run" mode. This means the object will not be fetched.
	DryRun bool `json:"dryRun"`
}

type ListOptions struct {
	// DryRun when set, the object will be listed in "dry-run" mode. This means the object will not be listed.
	DryRun bool `json:"dryRun"`
	// Limit is the maximum number of items to return. If not specified, it will be defaulted to 100.
	Limit int64 `json:"limit"`
	// Offset is a query offset value for the result set.
	Offset int64 `json:"offset"`
	// FieldSelector is a query against a resource's fields. A comma-delimited list of fields to set in the result.
	FieldSelector string `json:"fieldSelector"`
	// LabelSelector is a query against a set of labels. A comma-delimited list of labels to query for.
	LabelSelector string `json:"labelSelector"`
	// ResourceVersion sets a constraint on what resource versions a request may be served from.
	ResourceVersion string `json:"resourceVersion"`
	// TimeoutSeconds sets the timeout for this list request.
	TimeoutSeconds int64 `json:"timeoutSeconds"`
}

type UpdateOptions struct {
	// DryRun when set, the object will be updated in "dry-run" mode. This means the object will not be persisted.
	DryRun bool `json:"dryRun"`
}
