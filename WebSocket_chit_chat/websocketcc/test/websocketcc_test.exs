defmodule WebsocketccTest do
  use ExUnit.Case
  doctest Websocketcc

  test "greets the world" do
    assert Websocketcc.hello() == :world
  end
end
