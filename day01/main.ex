defmodule Day01 do
  def solve_part1(input) do
    input
    |> Enum.map(
      &case &1 do
        "(" -> 1
        ")" -> -1
      end
    )
    |> Enum.sum()
  end

  def solve_part2(contents) do
    contents
    |> Enum.with_index()
    |> Enum.reduce({0, -1}, fn {char, index}, {floor, basement} ->
      case char do
        "(" ->
          {floor + 1, basement}

        ")" ->
          {floor - 1,
           if basement === -1 && floor === 0 do
             index + 1
           else
             basement
           end}
      end
    end)
    |> elem(1)
  end
end

contents =
  File.read!("day01/input.txt")
  |> String.trim()
  |> String.graphemes()

contents
|> Day01.solve_part1()
|> IO.puts()

contents
|> Day01.solve_part2()
|> IO.puts()
