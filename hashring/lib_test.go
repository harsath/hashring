package hashring

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestAdd(t *testing.T) {
	hash_ring := &HashRing{}
	server_one_uuid := uuid.New()
	hash_ring.Add(server_one_uuid)
	if len(hash_ring.node_holder) != 1 {
		t.Errorf("hash_ring len() not 1")
	}
	server_two_uuid := uuid.New()
	hash_ring.Add(server_two_uuid)
	server_three_uuid := uuid.New()
	hash_ring.Add(server_three_uuid)
	server_four_uuid := uuid.New()
	hash_ring.Add(server_four_uuid)
	server_five_uuid := uuid.New()
	hash_ring.Add(server_five_uuid)
	server_six_uuid := uuid.New()
	hash_ring.Add(server_six_uuid)
	if len(hash_ring.node_holder) != 6 {
		t.Errorf("hash_ring len() not 6")
	}
	fmt.Println("radian_node of 1st server: ",
		hash_ring.node_holder[0].radian_node)
	fmt.Println("UUID of 1st server: ",
		hash_ring.node_holder[0].uuid_node)
	fmt.Println("radian_node of 2nd server: ",
		hash_ring.node_holder[1].radian_node)
	fmt.Println("UUID of 2nd server: ",
		hash_ring.node_holder[1].uuid_node)
	fmt.Println("radian_node of 3st server: ",
		hash_ring.node_holder[2].radian_node)
	fmt.Println("UUID of 3st server: ",
		hash_ring.node_holder[2].uuid_node)
	fmt.Println("radian_node of 4nd server: ",
		hash_ring.node_holder[3].radian_node)
	fmt.Println("UUID of 4nd server: ",
		hash_ring.node_holder[3].uuid_node)
	fmt.Println("radian_node of 5st server: ",
		hash_ring.node_holder[4].radian_node)
	fmt.Println("UUID of 5st server: ",
		hash_ring.node_holder[4].uuid_node)
	fmt.Println("radian_node of 6nd server: ",
		hash_ring.node_holder[5].radian_node)
	fmt.Println("UUID of 6nd server: ",
		hash_ring.node_holder[5].uuid_node)
	blob_one := []byte("Hello world, this is bytes!")
	spot_one_uuid := hash_ring.FindSpot(blob_one)
	fmt.Println("blob_one spot: ", spot_one_uuid)
}
