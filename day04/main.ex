defmodule Day04 do
  def solve(key, acc, prefix) do
    hash =
      :crypto.hash(:md5, "#{key}#{acc}")
      |> Base.encode16()

    case String.starts_with?(hash, prefix) do
      true -> acc
      false -> solve(key, acc + 1, prefix)
    end
  end
end

input =
  File.read!("day04/input.txt")
  |> String.trim()

part1 = Day04.solve(input, 0, "00000")
part2 = Day04.solve(input, part1 + 1, "000000")

IO.puts(part1)
IO.puts(part2)
