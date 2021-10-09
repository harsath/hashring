package hashring

import "github.com/google/uuid"

type nodeHolder struct {
	uuid_node   uuid.UUID
	radian_node int
}

type HashRing struct {
	// TODO: It's better to store them in a BST with radian_node as the key
	node_holder []nodeHolder
}

/**
Add a node into the cluster
node_uuid Input UUID identifier for a node to be inserted
It returns the successor node, such that an implementation which uses this will
make move operation of the BLOBs stored in the other server.
*/
func (hash_ring *HashRing) Add(node_uuid uuid.UUID) uuid.UUID {
	var uuid_hash_output int = int(hash_ring.hashUUID(node_uuid)) % 360
	var stop_index int = 0
	for stop_index < len(hash_ring.node_holder) {
		if uuid_hash_output <= hash_ring.node_holder[stop_index].radian_node {
			break
		}
		stop_index++
	}
	// Input UUID's hash is the largest, hence we insert at the end
	if stop_index == len(hash_ring.node_holder) {
		hash_ring.node_holder = append(hash_ring.node_holder,
			nodeHolder{uuid_node: node_uuid, radian_node: uuid_hash_output})
		return hash_ring.node_holder[0].uuid_node
	}
	// Insert at the 'stop_index' spot
	temp_node_holder := make([]nodeHolder, 0)
	temp_node_holder = append(temp_node_holder, hash_ring.node_holder[:stop_index]...)
	temp_node_holder = append(temp_node_holder,
		nodeHolder{uuid_node: node_uuid, radian_node: uuid_hash_output})
	temp_node_holder = append(temp_node_holder, hash_ring.node_holder[stop_index+1:]...)
	hash_ring.node_holder = temp_node_holder
	return hash_ring.node_holder[stop_index+1].uuid_node
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
	var current_index int = hash_ring.current_size / 2
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
