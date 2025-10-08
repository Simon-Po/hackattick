# Hackattic WebSocket Chit Chat – How I Solved It

I built a small Elixir app that automates the WebSocket chit chat challenge on Hackattic. The flow starts in `Websocketcc.start/0`, where I call Hackattic's REST endpoint with `Req` to obtain a fresh session token. With that token I assemble the `wss://` URL and spin up a `WebSockex` client module named `Client`.

`Client` keeps the round-trip timing logic. On each `ping!` frame from Hackattic I compare the elapsed milliseconds against the allowed answers (`[700, 1500, 2000, 2500, 3000]`) and pick the closest match before replying. That same module listens for the final `congratulations!` message, extracts the embedded secret token, and hands it off to a helper process.

The helper is a lightweight GenServer (`ResSender`). Once it receives the secret it posts the result back to Hackattic's `/solve` endpoint, again using `Req`, and prints the JSON response so I can confirm that the submission landed.

Thisfetch token, run the WebSocket timing loop, then report the answer—covers the entire handshake required by the challenge, so kicking off `Websocketcc.start/0` completes the exercise end to end.

Beam is the perfect plattform for a Challenge like this the ease of using Elixir keeps suprising me.

I will try to solve the same challenge with go to figure out how i would go about it there.
