package main

import (
	"aocutil"
	"strconv"
	"strings"
)

type CookieIngredient struct {
	capacity, durability, flavor, texture, calories int
}

func parse_line(str string) (ingredient CookieIngredient) {
	split := strings.Split(strings.Replace(str, ",", "", -1), " ")
	ingredient.capacity, _ = strconv.Atoi(split[2])
	ingredient.durability, _ = strconv.Atoi(split[4])
	ingredient.flavor, _ = strconv.Atoi(split[6])
	ingredient.texture, _ = strconv.Atoi(split[8])
	ingredient.calories, _ = strconv.Atoi(split[10])
	return
}

var (
	ingredients []CookieIngredient
)

type IngredientAmount struct {
	id     int
	amount int
}

func calculate_cookie_score(amount []IngredientAmount, calorieConstraint int) (score int) {
	result := CookieIngredient{}
	for _, e := range amount {
		current := &ingredients[e.id]
		result.capacity += (current.capacity * e.amount)
		result.durability += (current.durability * e.amount)
		result.flavor += (current.flavor * e.amount)
		result.texture += (current.texture * e.amount)
		result.calories += (current.calories * e.amount)
	}

	if calorieConstraint > 0 {
		if result.calories != calorieConstraint {
			return 0
		}
	}

	result.capacity = aocutil.ClampLower(result.capacity, 0)
	result.durability = aocutil.ClampLower(result.durability, 0)
	result.flavor = aocutil.ClampLower(result.flavor, 0)
	result.texture = aocutil.ClampLower(result.texture, 0)

	return result.capacity * result.durability * result.flavor * result.texture
}

func solution_recursive(calorieConstraint, index, remaining int, amounts *[]IngredientAmount, maxScore *int) {
	if index == len(*amounts) {
		return
	}
	score := 0
	for i := 0; i <= remaining; i++ {
		(*amounts)[index].amount = remaining - i
		r := remaining - (*amounts)[index].amount
		solution_recursive(calorieConstraint, index+1, r, amounts, maxScore)
		score = calculate_cookie_score(*amounts, calorieConstraint)
		if score > *maxScore {
			*maxScore = score
		}
	}
}

func solution(calorieConstraint int) int {
	amounts := make([]IngredientAmount, len(ingredients))
	for i := 0; i < len(amounts); i++ {
		amounts[i].id = i
		amounts[i].amount = 0
	}
	score := 0
	solution_recursive(calorieConstraint, 0, 100, &amounts, &score)
	return score
}

func main() {
	ingredients = make([]CookieIngredient, 0, 10)
	aocutil.FileReadAllLines("day15/input.txt", func(s string) {
		ingredients = append(ingredients, parse_line(s))
	})
	aocutil.AOCFinish(solution(0), solution(500))
}
