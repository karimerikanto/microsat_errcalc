package models

// Sample includes all replicas
type Sample struct {
	ReplicaArray []Replica
}

// IsSingle says that is the sample single type (only one replica)
func (sample Sample) IsSingle() bool {
	return len(sample.ReplicaArray) == 1
}

// GetReplicaNames returns names of the replicas
func (sample Sample) GetReplicaNames() string {
	replicaNames := ""

	for _, replica := range sample.ReplicaArray {
		if len(replicaNames) > 0 {
			replicaNames += ", "
		}

		replicaNames += replica.Name
	}

	return replicaNames
}
