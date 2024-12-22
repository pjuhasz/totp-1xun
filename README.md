# TOTP 1×űN

## ... or, TOTP made simple

This program serves a very specific use case:

You have elderly relatives who don't have a smartphone and don't
_want_ to have a smartphone, yet they have to use an online service that
insists on TOTP for 2-factor authentication (ahem, Ügyfélkapu+).

All you have to do to help them is save the QR code the service generates
when enrolling the user, then run this program with the path to the image
as an argument. The program will generate a passcode and also copy it
to the clipboard for maximum comfort. You can even create a launcher
shortcut on the desktop for the true one-(double-)click 2FA experience.

Yes, this decreases security, and in a way, defeats the purpose of 2FA.
The QR code containing the secret has to be stored on the computer in plaintext
(or plain image, anyway). But for the very specific use case outlined above
(your relative accessing the 2FA service always from their dedicated
computer, inside a locked house) this is considered better than the alternatives.

## Features

Run the program as `totp1xun /path/to/qr_code.png` on Linux
or `totp1xun C:\path\to\qr_code.png` on Windows.

This will generate a standard 6-digit passcode that is valid for (at most)
30 seconds, print it to the terminal, and then exit.

Alternatively, run the program with the `-persist` command line switch,
in which case it will keep running and generating passcodes every 30 seconds.
You can kill it with Ctrl-C or closing the window or whatever.

The program will also attempt to copy the passcodes to the clipboard.
On Windows this supposedly Just Works, on Linux you unfortunately have to
install some helper programs that work with your windowing system's clipboard
(`xclip` for X, `wl-copy` for Wayland).

## Compilation

Just run `go build`.

Pre-built binaries for Linux and Windows are included in the bin/ directory.

## License

MIT License. (C) Juhász Péter, 2024
