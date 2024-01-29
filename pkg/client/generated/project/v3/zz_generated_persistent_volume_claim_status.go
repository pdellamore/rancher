package client

const (
	PersistentVolumeClaimStatusType                           = "persistentVolumeClaimStatus"
	PersistentVolumeClaimStatusFieldAccessModes               = "accessModes"
	PersistentVolumeClaimStatusFieldAllocatedResourceStatuses = "allocatedResourceStatuses"
	PersistentVolumeClaimStatusFieldAllocatedResources        = "allocatedResources"
	PersistentVolumeClaimStatusFieldCapacity                  = "capacity"
	PersistentVolumeClaimStatusFieldConditions                = "conditions"
	PersistentVolumeClaimStatusFieldPhase                     = "phase"
)

type PersistentVolumeClaimStatus struct {
	AccessModes               []string                         `json:"accessModes,omitempty" yaml:"accessModes,omitempty"`
	AllocatedResourceStatuses map[string]string                `json:"allocatedResourceStatuses,omitempty" yaml:"allocatedResourceStatuses,omitempty"`
	AllocatedResources        map[string]string                `json:"allocatedResources,omitempty" yaml:"allocatedResources,omitempty"`
	Capacity                  map[string]string                `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	Conditions                []PersistentVolumeClaimCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Phase                     string                           `json:"phase,omitempty" yaml:"phase,omitempty"`
}
