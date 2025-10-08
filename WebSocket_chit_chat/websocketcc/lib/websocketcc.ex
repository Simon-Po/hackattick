defmodule Websocketcc do
  @moduledoc """
  Documentation for `Websocketcc`.
  """

  @doc """
  Hello world.

  ## Examples

      iex> Websocketcc.hello()
      :world

  """
  def start do
    # make call to hackattic
    ResSender.start_link()
    
    resp = Req.get!("https://hackattic.com/challenges/websocket_chit_chat/problem?access_token=be8466c39c975877")
    # call webserver with url
    url = "wss://hackattic.com/_/ws/#{resp.body["token"]}"
    Client.start_link(url)
  end
end
