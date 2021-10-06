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
func (hash_ring *HashRing) Add(node_uuid uuid.UUID) {
	// Use binary search to find the spot.
	var current_index int = hash_ring.current_size / 2
	var uuid_hash_output int = hash_ring.hashUUID(node_uuid)
	var index_to_add int
	for {
		if uuid_hash_output <= hash_ring.node_holder[current_index].radian_node {
			index_to_add = current_index
			break
		}
		current_index = current_index / 2
	}
}

/**
Remove a node from the cluster
node_uuid Node to remove from the cluster
*/
func (hash_ring *HashRing) Remove(node_uuid uuid.UUID) {}

/**
Find a spot for the input BLOB
blob_data Input BLOB, which is not hashed. So, we first need to hash it and find the spot
*/
func (hash_ring *HashRing) FindSpot(blob_data []byte) uuid.UUID {

}

func (hash_ring *HashRing) hashUUID(node_uuid uuid.UUID) uint64 {
	// djb2 hashing
	var hash uint64 = 5381
	index := 0
	for char := node_uuid[index]; ; {
		if index == len(node_uuid) {
			break
		}
		hash = ((hash << 5) + hash) + uint64(char)
		index++
	}
	return hash
}
