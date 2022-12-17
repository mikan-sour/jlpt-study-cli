package game

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"jlpt/src/core/models"
	"jlpt/src/core/service"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed, color.Bold).SprintFunc()

type GameState int

const (
	noGame GameState = iota
	inProgress
	win
	lose
)

type Game interface {
	Run(ctx context.Context, questions *[][]*models.VocabWord) error
	AskQuestion(words []*models.VocabWord) (Question, error)
	GetWordsForQuestion(ctx context.Context, low, high int) ([]*models.VocabWord, error)
}

type GameImpl struct {
	State       GameState
	Score       int
	Misses      int
	IsRestart   bool
	AlreadySeen [][]*models.VocabWord
	Missed      [][]*models.VocabWord
	Service     service.MatchingGameService
}

func NewGame(service service.MatchingGameService) Game {
	return &GameImpl{Service: service, State: noGame}
}

func (g *GameImpl) GetWordsForQuestion(ctx context.Context, low, high int) ([]*models.VocabWord, error) {

	rand := g.Service.GetRandomIds(low, high)
	words, err := g.Service.GetWordsForAQuestion(ctx, rand)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (g *GameImpl) Shuffle(arr [][]*models.VocabWord) [][]*models.VocabWord {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([][]*models.VocabWord, len(arr))
	perm := r.Perm(len(arr))
	for i, randIndex := range perm {
		ret[i] = arr[randIndex]
	}
	return ret
}

func (g *GameImpl) Run(ctx context.Context, lastRoundQuestions *[][]*models.VocabWord) error {
	g.State = inProgress
	g.Score = 0
	g.Misses = 0

	low, high, err := g.Service.GetIDRange(ctx)
	if err != nil {
		return err
	}

	numberOfQuestionsToGet := 12
	questions := [][]*models.VocabWord{}

	if lastRoundQuestions != nil && len(*lastRoundQuestions) > 0 {
		questions = append(questions, *lastRoundQuestions...)
		numberOfQuestionsToGet = numberOfQuestionsToGet - len(questions)
	}

	forThisRound := make([][]*models.VocabWord, 12)

	wg := &sync.WaitGroup{}
	wg.Add(numberOfQuestionsToGet)

	errs := make(chan error, 1)

	for i := 0; i < numberOfQuestionsToGet; i++ {
		go func(i int) {
			defer wg.Done()
			words, err := g.GetWordsForQuestion(ctx, low, high)
			if err != nil {
				errs <- err
			}
			forThisRound[i] = words
		}(i)
	}
	close(errs)

	if err := <-errs; err != nil {
		return err
	}
	wg.Wait()

	questions = append(questions, forThisRound...)
	questions = g.Shuffle(questions)
	for _, q := range questions {
		res, err := g.AskQuestion(q)
		if err != nil {
			return err
		}

		if res.GetResult() {
			g.Score += 1
			fmt.Printf("%s You got it right! %d points, %d misses\n", green("\U00002714"), g.Score, g.Misses)
		} else {
			g.Misses += 1
			g.Missed = append(g.Missed, q)
			fmt.Printf("%s You got it wrong... %d points, %d misses\n", red("X"), g.Score, g.Misses)
		}

		g.AlreadySeen = append(g.AlreadySeen, res.GetOptions())

		if g.Score == 10 {
			g.State = win
			break
		}
		if g.Misses == 3 {
			g.State = lose
			break
		}

	}

	promptOptions := []string{
		"1. redo round with same questions",
		"2. redo round with missed questions",
		"3. redo round with new questions",
		"4. quit",
	}

	prompt := promptui.Select{
		Label: "The game is over - what's next?",
		Items: promptOptions,
		Templates: &promptui.SelectTemplates{
			Active:   " \U0001F352  {{ . | cyan }}",
			Inactive: "  {{ . | white }} ",
			Selected: "  {{ . | faint }}",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch result {
	case promptOptions[0]:
		g.Run(ctx, &g.AlreadySeen)
	case promptOptions[1]:
		g.Run(ctx, &g.Missed)
	case promptOptions[2]:
		g.Run(ctx, nil)
	default:
		fmt.Println("thanks for playing! さよなら!")
	}

	return nil

}
