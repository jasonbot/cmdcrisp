# CommandCrisp

Did you know that the main producer of apples in the United States is the state of Washington? Did you know that the main producer of Microsoft software in the United States is Microsoft, _also_ in Washington?

Really makes you think.

There are several popular varieties of apple on the market today with _Crisp_ in their name: Honeycrisp, Crimson Crisp, Evercrisp, and the coveted and adored Cosmic Crisp. Now the `cmd.exe` prompt can also pretend to be a variety of Apple: **cmdcrisp**.

At my work, our developer experience is entirely geared toward local dev on macOS. I am doing Windows, because we build native Windows clients. I'm about 90% of the way getting the stack up and running through a mix of Bash in Windows and running services in WSL, but a few of the scripts use things specific to macOS' command line utilities.

I've implemented _just enough Apple_ to make my unholy Windows/Linux build hack stack runnable.

This is _just scratching an itch_ to make the shell scripts I care about runnable on window command line and in WSL hosts: it _is not currently a full-fledged functional reimplementation_ of the whole thing. If you have further itches, you can fork it or issue a PR here. This is not serious software.

Commands currently implemented:

- `pbcopy`/`pbpaste`: They're pretty straightforward
- `open`: Just a subset of this is implemented, and it just forwards to `start` or `xdg-open`, respectively.
- `security`: Only the teeniest bit of this is implemented for password storage using the OS-native keychain (see [`go-keyring`](https://github.com/zalando/go-keyring)).

## Installing

You need a working copy of the Go compiler toolchain installed on your machine.

```shell
go install github.com/jasonbot/cmdcrisp@latest
```

If you installed Go via the native installer in Windows, the place it adds tools to should already be in your `%PATH%`.

In other systems, you'll need to add `$GOPATH/bin` to your `$PATH`, usually it's `~/go/bin/`.

## Using

Type the name of the desired command in a black text window and press enter. Hooray.
