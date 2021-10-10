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
	fmt.Println("radian_node of 1st server: ",
		hash_ring.node_holder[0].radian_node)
	server_two_uuid := uuid.New()
	hash_ring.Add(server_two_uuid)
	fmt.Println(len(hash_ring.node_holder))
	if len(hash_ring.node_holder) != 2 {
		t.Errorf("hash_ring len() not 2")
	}
	fmt.Println("radian_node of 2nd server: ",
		hash_ring.node_holder[1].radian_node)
}
