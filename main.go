package main

import (
	"fmt"
	"math/rand"
	"time"
)

var sampleText = []string{
	"Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptate soluta rerum tempore iusto. Nostrum minima magni dolor, minus eum quidem voluptatibus ipsum, labore assumenda libero sed nemo repellendus? Dolore, laborum.",
	"Lorem ipsum dolor sit amet consectetur adipisicing elit. Nam quisquam recusandae sunt eum deleniti temporibus sed atque magni, quas sequi nemo similique fugiat nostrum? Quidem ad maxime maiores necessitatibus distinctio!",
	"Lorem ipsum, dolor sit amet consectetur adipisicing elit. Eligendi aliquam aperiam culpa molestias explicabo pariatur rerum ipsa beatae dolorum! Error optio sapiente aliquam provident, perferendis numquam unde repudiandae itaque nisi.",
	"Lorem ipsum dolor sit amet consectetur adipisicing elit. Itaque velit, officia impedit quia nobis, harum ipsam omnis natus quam beatae cum quae quidem nihil facere minima laudantium porro saepe neque.",
	"Lorem ipsum dolor sit amet, consectetur adipisicing elit. Minima, reiciendis molestias magni ea repudiandae vel quo nostrum rerum, quae repellat quis magnam laboriosam, tenetur explicabo error maiores facilis debitis aspernatur.",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	m := NewMarkovChain()
	for _, line := range sampleText {
		m.ReadText(line)
	}
	m.MakeProbabilities()

	fmt.Println(m.Sentence())
}
