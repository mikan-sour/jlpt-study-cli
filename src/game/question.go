package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"jlpt/src/core/models"

	"github.com/manifoldco/promptui"
)

type Question interface {
	Shuffle(vals []*models.VocabWord) []*models.VocabWord
	GetResult() bool
	GetOptions() []*models.VocabWord
	FormatDefs(defs string) string
}

type QuestionImpl struct {
	Result  bool   // if guessed correctly, then `true`
	Correct string // the correct answer
	Guessed string
	Options []*models.VocabWord
}

func (g *GameImpl) AskQuestion(words []*models.VocabWord) (Question, error) {

	var question QuestionImpl

	copied := make([]*models.VocabWord, len(words))
	copy(copied, words)

	copied = question.Shuffle(copied)
	question.Options = copied
	question.Correct = question.FormatDefs(words[0].Definitions)

	prompt := promptui.Select{
		Label: fmt.Sprintf("Select the meaning of %s (%s)", words[0].Foreign1, words[0].Foreign2),
		Items: []string{
			question.FormatDefs(copied[0].Definitions),
			question.FormatDefs(copied[1].Definitions),
			question.FormatDefs(copied[2].Definitions),
			question.FormatDefs(copied[3].Definitions),
		},
		Templates: &promptui.SelectTemplates{
			Active:   " \U0001F352  {{ . | cyan }}",
			Inactive: "  {{ . | white }} ",
			Selected: "  {{ . | faint }}",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	question.Guessed = result

	if result == question.Correct {
		question.Result = true
	} else {
		question.Result = false
	}

	return &question, nil

}

func (q *QuestionImpl) Shuffle(vals []*models.VocabWord) []*models.VocabWord {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]*models.VocabWord, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func (q *QuestionImpl) GetResult() bool {
	return q.Result
}

func (q *QuestionImpl) GetOptions() []*models.VocabWord {
	return q.Options
}

func (q *QuestionImpl) FormatDefs(defs string) string {
	allDefs := strings.Split(defs, ";")
	if len(allDefs) <= 3 {
		return defs
	}

	trimmed := strings.Join(allDefs[0:3], ";")

	slashRemoved := strings.Split(trimmed, "/")

	if len(slashRemoved) < 2 {
		return trimmed
	}

	return strings.Join(slashRemoved[0:2], ";")

}
