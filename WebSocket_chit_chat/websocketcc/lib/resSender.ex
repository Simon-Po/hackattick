defmodule ResSender do
  use GenServer

  def start_link do 
   GenServer.start_link(__MODULE__,nil,[name: __MODULE__]) 
  end

  def init(_init_arg) do 
    {:ok,nil}
  end


  def handle_cast({:send,{:token, sToken}},_) do 
    payload = %{:secret => sToken}

    resp = Req.post!("https://hackattic.com/challenges/websocket_chit_chat/solve?access_token=be8466c39c975877&playground=1", json: payload)

    IO.puts(inspect(resp.body))
    {:noreply,nil}
  end



end
