# discord-markov-chain

A markov chain for Discord message dumps using https://github.com/fr3fou/polo

## Download

If you have a `go` installed, just clone the repo or `go install github.com/fr3fou/discord-markov-chain@latest`
Otherwise, download a binary from the [Releases section](https://github.com/fr3fou/discord-markov-chain/releases)

## Usage

```shell
# or if you have a binary downloaded ./discord-markov-chain
$ go run main.go /path/to/discord-json-dump.json
Press enter for the next generated message
	You can also enter a starting word
	Type 'quit' to quit
>
```

You can export / dump messages in JSON for a given channel using tools
like [DiscordChatExporter](https://github.com/Tyrrrz/DiscordChatExporter)
