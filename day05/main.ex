defmodule Day05 do
  def is_nice?(str) do
    hasThreeVowels = (Regex.scan(~r/[aeiou]/, str) || []) |> length >= 3
    hasDoubleLetter = String.match?(str, ~r/(.)\1/)
    hasNoBadStrings = !String.match?(str, ~r/(ab|cd|pq|xy)/)
    hasThreeVowels and hasDoubleLetter and hasNoBadStrings
  end

  def is_nicer?(str) do
    hasPairTwice = String.match?(str, ~r/(..).*\1/)
    hasSandwich = String.match?(str, ~r/(.).\1/)
    hasPairTwice and hasSandwich
  end
end

input = File.stream!("day05/input.txt", [], :line)
input |> Enum.filter(&Day05.is_nice?/1) |> length |> IO.puts()
input |> Enum.filter(&Day05.is_nicer?/1) |> length |> IO.puts()
