package main

import (
	"testing"
)

func TestRanktemplate(t *testing.T) {
	Ranktemplate()
}

func TestInRoom(t *testing.T) {
	InRoom()
}

func TestOutRoom(t *testing.T) {
	OutRoom()
}

func TestSendchat(t *testing.T) {
	Sendchat()
}

func TestNotGift(t *testing.T) {
	NotGift()
}

func TestPay(t *testing.T) {
	Pay()
}

func TestSendGift(t *testing.T) {
	/*sendUserId := []uint64{1238179291,
		209830195,
		673955268,
		60000109,
		721452514,
		664561814,
		712144991,
		732868975,
		209829550,
		209829562,
		209829557,
		669388144,
		666358856,
		664563078,
		209829522,
		673955529,
		651892247,
		655223016,
		655223076,
		655219990,
		655222953,
		651892305,
		655223159,
		693105707,
		649896412,
		655220171,
		655220357,
		651889729,
		655220474}
	for _, id := range sendUserId {
		// SendGift(id, 71067961, "1_50095", 12345869, 20)
		SendGift(655220474, id, "1_50095", 12345870, 20)
	}*/

	//	SendGift(253746851, 596694974, "1_50095", 12345877, 10, uint64(time.Now().Unix()))
	NotGift()
}

func TestSendNotGift(t *testing.T) {
	SendNotGift()
}
