package models

// Sample includes all replicas
type Sample struct {
	ReplicaArray []Replica
}

// IsSingle says that is the sample single type (only one replica)
func (sample Sample) IsSingle() bool {
	return len(sample.ReplicaArray) == 1
}
