defmodule Day03 do
  def move(command, {x, y}) do
    case command do
      "<" -> {x - 1, y + 0}
      ">" -> {x + 1, y + 0}
      "^" -> {x + 0, y - 1}
      "v" -> {x + 0, y + 1}
    end
  end

  def deliver(houses, position) do
    presents = Map.get(houses, position, 0) + 1
    Map.put(houses, position, presents)
  end

  def solve_part1(input) do
    input
    |> Enum.reduce(
      {{0, 0}, %{{0, 0} => 1}},
      fn cmd, {pos, houses} ->
        pos = move(cmd, pos)
        {pos, deliver(houses, pos)}
      end
    )
    |> elem(1)
    |> map_size
  end

  def solve_part2(input) do
    input
    |> Enum.chunk_every(2)
    |> Enum.reduce(
      {{0, 0}, {0, 0}, %{{0, 0} => 2}},
      fn [cmd1, cmd2], {pos1, pos2, houses} ->
        pos1 = move(cmd1, pos1)
        pos2 = move(cmd2, pos2)
        {pos1, pos2, houses |> deliver(pos1) |> deliver(pos2)}
      end
    )
    |> elem(2)
    |> map_size
  end
end

contents =
  File.read!("day03/input.txt")
  |> String.trim()
  |> String.graphemes()

contents
|> Day03.solve_part1()
|> IO.puts()

contents
|> Day03.solve_part2()
|> IO.puts()
