/**
  Copyright 2016 Brandon Martinez
  Based off of Bradley Kennedy "NiceThings.py"
  https://github.com/co60ca
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package niceThings

import (
	"math/rand"
	"strings"
	"time"
)

func GeneratePhrase() string {
	if len(runTimeCache) == 0 {
		static := []string{
			"<3 u {0}",
			"<3 {0} <3",
			"You the best {0}.",
			"Need some love {0}? I love you!",
		}
		start := []string{
			"Hey,", "Hi,", "Hello,", "Hi {0}!", "Hey {0},",
		}
		mid := []string{
			"{0} is the best <3.", "I love you {0}.", "you are important {0}.", "{0} is so kind.", "I love you.", "I value you, {0}.", "never give up, {0} you can do it.",
		}
		end := []string{
			"Have a great day.", "Have a great day {0}.", "Can't wait to hear from you again.", "I miss you {0}!", "<3", "<3 <3 <3!", "<3 from me!",
		}
		mid = append(mid, NiceThings...)
		gen := make(map[int][]string)
		gen[0] = start
		gen[1] = mid
		gen[2] = end

		genList := cartesianProductPhrase(gen)

		static = append(static, genList...)
		static = append(static, NiceThings...)
		runTimeCache = static
	}
	return RandomSliceItem(runTimeCache)
}

func cartesianProductPhrase(gen map[int][]string) []string {
	ret := []string{}
	for i := 0; i < len(gen[0]); i++ {
		for j := 0; j < len(gen[1]); j++ {
			for jj := 0; jj < len(gen[2]); jj++ {
				ret = append(ret, gen[0][i]+" "+firstToLower(gen[1][j])+" "+gen[2][jj])
			}
		}
	}
	return ret
}

func firstToLower(s string) string {
	if string(s[0]) == "I" {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

func RandomSliceItem(s []string) string {
	rand.Seed(time.Now().UnixNano())
	k := rand.Int() % len(s)
	return s[k]
}

var runTimeCache []string

var NiceThings = []string{
	"Your smile is contagious.",
	"You look great today.",
	"You're a smart cookie.",
	"I bet you make babies smile.",
	"You have impeccable manners.",
	"I like your style.",
	"You have the best laugh.",
	"I appreciate you.",
	"You are the most perfect you there is.",
	"You are enough.",
	"You're strong.",
	"Your perspective is refreshing.",
	"You're an awesome friend.",
	"You light up the room.",
	"You deserve a hug right now.",
	"You should be proud of yourself.",
	"You're more helpful than you realize.",
	"You have a great sense of humor.",
	"You've got all the right moves!",
	"Is that your picture next to \"charming\" in the dictionary?",
	"Your kindness is a balm to all who encounter it.",
	"You're all that and a super-size bag of chips.",
	"On a scale from 1 to 10, you're an 11.",
	"You are brave.",
	"You're even more beautiful on the inside than you are on the outside.",
	"You have the courage of your convictions.",
	"Your eyes are breathtaking.",
	"You are making a difference.",
	"You're like sunshine on a rainy day.",
	"You bring out the best in other people.",
	"You're a great listener.",
	"How is it that you always look great, even in sweatpants?",
	"Everything would be better if more people were like you!",
	"I bet you sweat glitter.",
	"You were cool way before hipsters were cool.",
	"That color is perfect on you.",
	"Hanging out with you is always a blast.",
	"You smell really good.",
	"Being around you makes everything better!",
	"When you say, \"I meant to do that,\" I totally believe you.",
	"When you're not afraid to be yourself is when you're most incredible.",
	"Colors seem brighter when you're around.",
	"That thing you don't like about yourself is what makes you so interesting.",
	"You're wonderful.",
	"You have cute elbows. For reals!",
	"Jokes are funnier when you tell them.",
	"You're better than a triple-scoop ice cream cone. With sprinkles.",
	"Your bellybutton is kind of adorable.",
	"Your hair looks stunning.",
	"You're one of a kind!",
	"You're inspiring.",
	"You should be thanked more often. So thank you!!",
	"Our community is better because you're in it.",
	"You have the best ideas.",
	"You always know how to find that silver lining.",
	"You're a candle in the darkness.",
	"You're a great example to others.",
	"Being around you is like being on a happy little vacation.",
	"You always know just what to say.",
	"You're more fun than bubble wrap.",
	"When you make a mistake, you fix it.",
	"Who raised you? They deserve a medal for a job well done.",
	"You're great at figuring stuff out.",
	"Your voice is magnificent.",
	"The people you love are lucky to have you in their lives.",
	"You're like a breath of fresh air.",
	"You're gorgeous -- and that's the least interesting thing about you, too.",
	"You're so thoughtful.",
	"Your creative potential seems limitless.",
	"Your name suits you to a T.",
	"You're irresistible when you blush.",
	"Actions speak louder than words, and yours tell an incredible story.",
	"Somehow you make time stop and fly at the same time.",
	"When you make up your mind about something, nothing stands in your way.",
	"You seem to really know who you are.",
	"Any team would be lucky to have you on it.",
	"In high school I bet you were voted \"most likely to keep being awesome.\"",
	"I bet you do the crossword puzzle in ink.",
	"Babies and small animals probably love you.",
	"There's ordinary, and then there's you.",
	"You're someone's reason to smile.",
	"You're even better than a unicorn, because you're real.",
	"How do you keep being so funny and making everyone laugh?",
	"You have a good head on your shoulders.",
	"Has anyone ever told you that you have great posture?",
	"The way you treasure your loved ones is incredible.",
	"You're really something special.",
	"You're a gift to those around you.",
	"You're a compelling image.",
}
