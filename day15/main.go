package main

import (
	"aocutil"
	"strconv"
	"strings"
)

/*
--- Day 15: Science for Hungry People ---

Today, you set out on the task of perfecting your milk-dunking cookie recipe. All you have to do is find the right balance of ingredients.

Your recipe leaves room for exactly 100 teaspoons of ingredients. You make a list of the remaining ingredients you could use to finish the recipe (your puzzle input) and their properties per teaspoon:

    capacity (how well it helps the cookie absorb milk)
    durability (how well it keeps the cookie intact when full of milk)
    flavor (how tasty it makes the cookie)
    texture (how it improves the feel of the cookie)
    calories (how many calories it adds to the cookie)

You can only measure ingredients in whole-teaspoon amounts accurately, and you have to be accurate so you can reproduce your results in the future. The total score of a cookie can be found by adding up each of the properties (negative totals become 0) and then multiplying together everything except calories.

For instance, suppose you have these two ingredients:

Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3

Then, choosing to use 44 teaspoons of butterscotch and 56 teaspoons of cinnamon (because the amounts of each ingredient must add up to 100) would result in a cookie with the following properties:

    A capacity of 44*-1 + 56*2 = 68
    A durability of 44*-2 + 56*3 = 80
    A flavor of 44*6 + 56*-2 = 152
    A texture of 44*3 + 56*-1 = 76

Multiplying these together (68 * 80 * 152 * 76, ignoring calories for now) results in a total score of 62842880, which happens to be the best score possible given these ingredients. If any properties had produced a negative total, it would have instead become zero, causing the whole score to multiply to zero.

Given the ingredients in your kitchen and their properties, what is the total score of the highest-scoring cookie you can make?

--- Part Two ---

Your cookie recipe becomes wildly popular! Someone asks if you can make another recipe that has exactly 500 calories per cookie (so they can use it as a meal replacement). Keep the rest of your award-winning process the same (100 teaspoons, same ingredients, same scoring system).

For example, given the ingredients above, if you had instead selected 40 teaspoons of butterscotch and 60 teaspoons of cinnamon (which still adds to 100), the total calorie count would be 40*8 + 60*3 = 500. The total score would go down, though: only 57600000, the best you can do in such trying circumstances.

Given the ingredients in your kitchen and their properties, what is the total score of the highest-scoring cookie you can make with a calorie total of 500?

*/

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
