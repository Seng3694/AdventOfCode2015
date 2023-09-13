defmodule Day03 do
  defp move({command, {x, y}}) do
    case command do
      "<" -> {x - 1, y + 0}
      ">" -> {x + 1, y + 0}
      "^" -> {x + 0, y - 1}
      "v" -> {x + 0, y + 1}
    end
  end

  defp deliver(position, houses) do
    presents = Map.get(houses, position, 0) + 1
    Map.put(houses, position, presents)
  end

  def solve(input, count) do
    input
    |> Enum.chunk_every(count)
    |> Enum.reduce(
      {List.duplicate({0, 0}, count), %{{0, 0} => count}},
      fn commands, {positions, houses} ->
        positions = Enum.zip(commands, positions) |> Enum.map(&move/1)
        {positions, Enum.reduce(positions, houses, &deliver/2)}
      end
    )
    |> elem(1)
    |> map_size
  end
end

contents =
  File.read!("day03/input.txt")
  |> String.trim()
  |> String.graphemes()

contents
|> Day03.solve(1)
|> IO.puts()

contents
|> Day03.solve(2)
|> IO.puts()
