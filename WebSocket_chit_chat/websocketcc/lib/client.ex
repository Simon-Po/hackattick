defmodule Client do
  use WebSockex

  def start_link(url, _state \\ %{}) do
    WebSockex.start_link(url, __MODULE__, :os.system_time(:millisecond),[name: __MODULE__])
  end
  @doc """
  Approximate the Closetst value from targets
  """
  def approximateClosest(diff) do
    targets = [700, 1500,2000,2500,3000]
    Enum.min_by(targets, fn t -> abs(diff - t) end) 
  end

  @impl true
  def handle_frame({:text, "hello! " <> rest}, state) do
    IO.puts("Got text: hello! #{rest}")
    {:ok, state}
  end
  @impl true
  def handle_frame({:text, "good!"}, state) do
    IO.puts("Got text: good!")
    {:ok, state}
  end
  @impl true
  def handle_frame({:text, "ping!"}, state) do
    
    IO.puts("Got ping")
    now = :os.system_time(:millisecond)
    IO.puts("Time: #{now - state}")
    closest = approximateClosest(now - state)
    WebSockex.cast(self(),{:send, {:text, closest}})
    {:ok, now}
  end

def handle_frame({:text, "congratulations! the solution to this challenge is \"" <> token_and_quote}, state) do
  token = String.trim_trailing(token_and_quote, "\"")
  IO.puts("Got text: #{token}")
  GenServer.cast(ResSender, {:send, {:token, token}})
  {:ok, state}
end


  @impl true
  def handle_cast({:send, {:text, msg}}, state) do
    IO.puts("Sending #{msg}")
    {:reply, {:text, Integer.to_string(msg)}, state}
  end
end
