package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	text := "A 9-year-old boy looks down and his pants are wet and there is a puddle between his legs; he has peed his pants. This is an awful situation for any person to be in, but just when he thought this was the worst day in his life, Susie slips and spills a fish bowl of water right in the boy's lap. The boy is so happy that no one noticed before the water got spilt on him. The teacher rushes him to the nurse because his pants are soaked in fish water, his friends feel bad for him, and Susie whispers to him, 'I've wet my pants too.' This shows kids that we all have bad days and it is important to remember to have empathy at all times."
	// Change all words to lower characters to make it the same.
	text = strings.ToLower(text)
	// Find whole word and removes all punctuation and space
	re := regexp.MustCompile("\\w+")
	paragraph := re.FindAllString(text, -1)
	wordMap := calculateWord(paragraph)

	// Sort map by word frequncy
	pl := make(byFrequency, len(wordMap))
	i := 0
	for k, v := range wordMap {
		pl[i] = wordStructure{k, v}
		i++
	}
	sort.Sort(pl)
	// Slice map to take only the first 10 words
	topTenWordsUsed := pl[0:10]

	fmt.Println("Top ten words used:")
	fmt.Println("Frequency   Word")
	for i := 0; i < len(topTenWordsUsed); i++ {
		fmt.Println(topTenWordsUsed[i])
	}

}

func calculateWord(paragraph []string) map[string]int {
	// Initialise map
	wordCounter := make(map[string]int)
	/* Iterate through the paragraph and assign each word to a counter value
	and put the word together with the frequency value into a map
	*/
	for i := 0; i < len(paragraph); i++ {
		/* Check if the word in the string array exist in the map.
		If word exist, we will add 1 to the existing frequency value if not we will
		set frequency value to 1.
		*/
		wordExists := wordCounter[paragraph[i]]
		if wordExists != 0 {
			wordCounter[paragraph[i]] += 1
		} else {
			wordCounter[paragraph[i]] = 1
		}
	}
	return wordCounter
}

type wordStructure struct {
	word    string
	counter int
}

// Display word and frequency in this format
func (p wordStructure) String() string {
	return fmt.Sprintf("%d           %s", p.counter, p.word)
}

// Sorting pattern (by word frequency)
type byFrequency []wordStructure

func (a byFrequency) Len() int           { return len(a) }
func (a byFrequency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFrequency) Less(i, j int) bool { return a[i].counter > a[j].counter }
