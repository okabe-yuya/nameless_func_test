defmodule AggregaterTest do
  use ExUnit.Case
  doctest Aggregater

  test "test for Aggregater.counter" do
    # lst -> []string
    # expects -> map[string]int
    debuger = fn lst, expects ->
      res = Aggregater.counter(lst)
      Enum.each(expects, fn {key, val} ->
        if Map.has_key?(res, key) do
          if val != Map.get(res, key) do
            IO.puts("#{val} == #{Map.get(res, key)}")
            raise :error
          else
            :ok
          end
        else
          raise :error
        end
      end)
    end

    # シンプルなケース
    assert debuger.(["A", "B", "C", "D"], %{"A" => 1, "B" => 1, "C" => 1, "D" => 1}) == :ok

    # 重複した値がうまくカウントされているか
    assert debuger.(["A", "B", "A", "C"], %{"A" => 2, "B" => 1, "C" => 1}) == :ok

    # リストが空の場合
    assert debuger.([], %{}) == :ok

    # 複数の値が重複する場合
    assert debuger.(["A", "B", "A", "C", "D", "C", "E", "F", "D", "F", "G"], %{"A" => 2, "B" => 1, "C" => 2, "D" => 2, "E" => 1, "F" => 2, "G" => 1}) == :ok

    # 重複の発生が3回以上
    assert debuger.(["A", "A", "A"], %{"A" => 3}) == :ok
  end
end
