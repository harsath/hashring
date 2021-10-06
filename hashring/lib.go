package hashring

import "github.com/google/uuid"

type nodeHolder struct {
	uuid_node   uuid.UUID
	radian_node int
}

type HashRing struct {
	node_holder  []nodeHolder
	current_size int
}

/**
Add a node into the cluster
node_uuid Input UUID identifier for a node to be inserted
*/
func (hash_ring *HashRing) Add(node_uuid uuid.UUID) {}

/**
Remove a node from the cluster
node_uuid Node to remove from the cluster
*/
func (hash_ring *HashRing) Remove(node_uuid uuid.UUID) {}

/**
Find a spot for the input BLOB
blob_data Input BLOB, which is not hashed. So, we first need to hash it and find the spot
*/
func (hash_ring *HashRing) FindSpot(blob_data []byte) {}
