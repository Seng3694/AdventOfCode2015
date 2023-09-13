defmodule AdventOfCode do
  def parse_dimensions(line) do
    line
    |> String.trim()
    |> String.split("x")
    |> Enum.map(&String.to_integer/1)
  end

  defp calc_paper_area(dimensions) do
    dimensions
    |> (fn [l, w, h] -> [l * w, w * h, h * l] end).()
    |> (fn sides ->
          Enum.reduce(sides, 0, fn side, acc -> acc + 2 * side end) + Enum.min(sides)
        end).()
  end

  def solve_part1(dimensions) do
    Enum.reduce(dimensions, 0, fn dim, acc -> acc + calc_paper_area(dim) end)
  end

  def solve_part2(dimensions) do
    dimensions
    |> Enum.map(&Enum.sort/1)
    |> Enum.reduce(0, fn [a, b, c], acc ->
      acc + 2 * (a + b) + a * b * c
    end)
  end
end

dimensions =
  File.stream!("day02/input.txt", [], :line)
  |> Enum.map(&AdventOfCode.parse_dimensions/1)

dimensions
|> AdventOfCode.solve_part1()
|> IO.puts()

dimensions
|> AdventOfCode.solve_part2()
|> IO.puts()
