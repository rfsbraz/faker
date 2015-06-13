package provider

import (
	"fmt"
	"time"
)

var card_types map[string]*Card

type Card struct {
	Name               string
	Length             int
	PrefixList         [][]int
	SecurityCode       string
	SecurityCodeLength int
}

func NewCreditCard(faker *Faker) *CreditCard {
	visa_prefix_list := [][]int{
		[]int{4, 5, 3, 9},
		[]int{4, 5, 5, 6},
		[]int{4, 9, 1, 6},
		[]int{4, 5, 3, 2},
		[]int{4, 9, 2, 9},
		[]int{4, 0, 2, 4, 0, 0, 7, 1},
		[]int{4, 4, 8, 6},
		[]int{4, 7, 1, 6},
		[]int{4},
	}
	mastercard_prefix_list := [][]int{
		[]int{5, 1},
		[]int{5, 2},
		[]int{5, 3},
		[]int{5, 4},
		[]int{5, 5},
	}

	amex_prefix_list := [][]int{
		[]int{3, 4},
		[]int{3, 7},
	}

	discover_prefix_list := [][]int{[]int{6, 0, 1, 1}}

	diners_prefix_list := [][]int{
		[]int{3, 0, 0},
		[]int{3, 0, 1},
		[]int{3, 0, 2},
		[]int{3, 6},
		[]int{3, 8},
	}

	enroute_prefix_list := [][]int{
		[]int{2, 0, 1, 4},
		[]int{2, 1, 4, 9},
	}

	jcb16_prefix_list := [][]int{
		[]int{3, 0, 8, 8},
		[]int{3, 0, 9, 6},
		[]int{3, 1, 1, 2},
		[]int{3, 1, 5, 8},
		[]int{3, 3, 3, 7},
		[]int{3, 5, 2, 8},
	}

	jcb15_prefix_list := [][]int{
		[]int{2, 1, 0, 0},
		[]int{1, 8, 0, 0},
	}
	voyager_prefix_list := [][]int{[]int{8, 6, 9, 9}}

	mastercard := &Card{Name: "Mastercard", PrefixList: mastercard_prefix_list, Length: 16, SecurityCode: "CVV", SecurityCodeLength: 4}
	visa16 := &Card{Name: "VISA 16 digit", PrefixList: visa_prefix_list, Length: 16, SecurityCode: "CVC", SecurityCodeLength: 3}
	visa13 := &Card{Name: "VISA 13 digit", PrefixList: visa_prefix_list, Length: 13, SecurityCode: "CVC", SecurityCodeLength: 3}
	amex := &Card{Name: "American Express", PrefixList: amex_prefix_list, Length: 15, SecurityCode: "CVC", SecurityCodeLength: 3}
	discover := &Card{Name: "Discover", PrefixList: discover_prefix_list, Length: 16, SecurityCode: "CVC", SecurityCodeLength: 3}
	diners := &Card{Name: "Diners Club / Carte Blanche", PrefixList: diners_prefix_list, Length: 14, SecurityCode: "CVC", SecurityCodeLength: 3}
	enroute := &Card{Name: "enRoute", PrefixList: enroute_prefix_list, Length: 15, SecurityCode: "CVC", SecurityCodeLength: 3}
	jcb15 := &Card{Name: "JCB 15 digit", PrefixList: jcb15_prefix_list, Length: 15, SecurityCode: "CVC", SecurityCodeLength: 3}
	jcb16 := &Card{Name: "JCB 16 digit", PrefixList: jcb16_prefix_list, Length: 16, SecurityCode: "CVC", SecurityCodeLength: 3}
	voyager := &Card{Name: "Voyager", PrefixList: voyager_prefix_list, Length: 15, SecurityCode: "CVC", SecurityCodeLength: 3}

	card_types = map[string]*Card{
		"mastercard": mastercard,
		"visa16":     visa16,
		"visa":       visa16,
		"visa13":     visa13,
		"amex":       amex,
		"discover":   discover,
		"diners":     diners,
		"enroute":    enroute,
		"jcb15":      jcb15,
		"jcb16":      jcb16,
		"jcb":        jcb16,
		"voyager":    voyager,
	}

	return &CreditCard{faker, "credit_card"}
}

// Returns the corresponding card provider if one is passed,
// or a random provider when an empty string is passed as an argument
func (credit_card *CreditCard) CardProvider(card_type string) string {
	return credit_card.cardType(card_type).Name
}

func (credit_card *CreditCard) Number(card_type string) string {
	card := credit_card.cardType(card_type)
	return credit_card.generateNumber(card.PrefixList[credit_card.RandomNumber(len(card.PrefixList))], card.Length)
}

func (credit_card *CreditCard) generateNumber(prefixes []int, length int) string {

	// We cam append all the digits directly in prefixes
	for len(prefixes) < length-1 {
		prefixes = append(prefixes, credit_card.RandomDigit())
	}

	tot := 0

	// Revert number
	reversed_number := make([]int, len(prefixes))
	copy(reversed_number, prefixes)
	for i := 0; i < len(prefixes)/2; i++ {
		reversed_number[i], reversed_number[len(prefixes)-1-i] = reversed_number[len(prefixes)-1-i], reversed_number[i]
	}

	for i := 0; i < len(prefixes); i++ {
		if i%2 == 0 {
			reversed_number[i] *= 2
			if reversed_number[i] > 9 {
				reversed_number[i] -= 9
			}
		}
		tot += reversed_number[i]
	}

	// Calculate check digit
	check_digit := 10 - tot%10

	prefixes = append(prefixes, check_digit)

	//Convert number back to string
	number := ""
	for _, value := range prefixes {
		number = fmt.Sprintf("%v%v", number, value)
	}
	return number
}

func (credit_card *CreditCard) cardType(card_type string) *Card {
	if val, ok := card_types[card_type]; ok && card_type != "" {
		return val
	} else {
		types := make([]string, 0, len(card_types))
		for k := range card_types {
			types = append(types, k)
		}
		return card_types[types[credit_card.RandomNumber(len(types))]]
	}
}

func (credit_card *CreditCard) SecurityCode(card_type string) string {
	card := credit_card.cardType(card_type)
	return fmt.Sprintf("%v", credit_card.RandomNumberOfSize(card.SecurityCodeLength))
}

func (credit_card *CreditCard) ExpireDate() string {
	expire := time.Now().AddDate(0, credit_card.RandomNumber(5*12), credit_card.RandomNumber(31))
	return expire.Format("01/06")
}

func (credit_card *CreditCard) Full(provider string) string {
	card_provider := credit_card.cardType(provider)
	return fmt.Sprintf("%v\n%v\n%v %v\n%v: %v",
		card_provider.Name,
		credit_card.Faker.Person.Name(),
		credit_card.Number(provider),
		credit_card.ExpireDate(),
		card_provider.SecurityCode,
		credit_card.SecurityCode(provider),
	)
}
