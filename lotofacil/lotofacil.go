package lotofacil

import (
	"log"
	"math/rand"

	"github.com/agmguerra/lotericacli/common"
)

func GetLotofacilCards(numCards int) [][]int {
	var cards = make(map[[32]byte][]int)

	for i := 0; i < numCards; i++ {
		card := getCard()
		log.Println(card)
		hash := common.ComputeHashKeyForList(card, "0")
		cards[hash] = getCard()
	}

	return common.GetIntSliceFromMapValues(cards)
}

func getCard() []int {
	var card = make(map[int]int)
	cont := 1
	for i := 0; ; i++ {
		val := rand.Intn(25) + 1
		if _, exists := card[val]; !exists {
			card[val] = val
			cont++
		}
		if cont > 15 {
			break
		}
	}
	return common.GetSortedIntSliceFromMapValues(card)
}
