# gounseen

A simple encoding/decoding project to translate messages from [The Unseen (Dauntless)](https://twitter.com/dntls_TheUnseen)

# Pre-requisites

To run this, you must have Golang binaries v1.20 installed on your machine

# How to run

First, obviously, clone this repository to your machine.

## Decoding messages

With the terminal initialized on the root path of this project, call `go run . decode <message>`, replacing `<message>` with the encoded message you want to translate.

For example, to translate [this tweet](https://twitter.com/dntls_TheUnseen/status/1714719346159944070), just execute
```bash
go run . decode ΔㅏㄹㅋNЭΣΣ­ ­ИΣ ­NИΓㅎ. ­ㅏㄹЭ ­YΩU ­ㄹЭㅏΔY, ­ΣΛㅏYЭㄹ?
```

Which results in:
```bash
decoding 'ΔㅏㄹㅋNЭΣΣ­ ­ИΣ ­NИΓㅎ. ­ㅏㄹЭ ­YΩU ­ㄹЭㅏΔY ­ΣΛㅏYЭㄹ?':
result: 'DARKNESS­ ­IS ­NIGH. ­ARE ­YOU ­READY ­SLAYER?'
```

## Encoding messages

With the terminal initialized on the root path of this project, call `go run . encode <message>`, replacing `<message>` with the message you want to encode.

For example, to generate [this tweet](https://twitter.com/dntls_TheUnseen/status/1714719346159944070) encoded phrase, just execute
```bash
go run . encode DARKNESS­ ­IS ­NIGH. ­ARE ­YOU ­READY ­SLAYER?
```

Which results in:
```bash
encoding 'DARKNESS­ ­IS ­NIGH. ­ARE ­YOU ­READY ­SLAYER?':
result: 'Δㅏㄹㅋ𑀦ЭΣΣ­ ­ИΣ ­𑀦ИΓㅎ. ­ㅏㄹЭ ­YΩU ­ㄹЭㅏΔY ­ΣΛㅏYЭㄹ?'
```

Please note that some letters can have distinct results randomly, because they have more than one possible values. Those letters are:

- C: can be encoded to `C` or `𑀘`;
- J: can be encoded to `𑀛` or `ㅈ`;
- N: can be encoded to `N` or `𑀦`.

On decoding, there's no problem to translate those possibilities. But I don't have idea what kind of logic it should fulfill on the encoding process, so it's random based on the letter map initialization process.