defmodule Aggregater do
  def counter(lst) do
    Enum.reduce(lst, %{}, fn val, acc ->
      if Map.has_key?(acc, val) do
        Map.put(acc, val, Map.get(acc, val)+1)
      else
        Map.put(acc, val, 1)
      end
    end)
  end
end
