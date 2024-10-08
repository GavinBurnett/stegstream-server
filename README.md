Overview

Stegstream is made up of a server and client program. The server program uses steganography to hide a file inside a MP3 music file. The music file is then made available via the server program streaming it, and it can be listened to using a web browser in the same way as normal streaming. Running the client program will extract the hidden file from the music file streamed by the server. This method allows files to be transferred in an invisible way, as any observer will simply see a music streaming service.

Design Goal

To provide a practical application of steganography that is easy for non technical people to use in order to privately distribute files to a large number of people.

Installation

Download the archive file containing the server executable for the relevant operating system from the releases. Validate the PGP signature if an integrity check is advisable, and extract the executable from the archive file. Place the executable in a suitable directory; E.G /home/username/stegstream

GnuPG Signing Key: http://pgp.mit.edu/pks/lookup?op=get&search=0x203092F792253A6F

Setting up a server

Open a command prompt, and cd into the directory containing the server executable.

Enter the following (Using the example files):

./stegstream-server Waves.mp3 HideFile.txt

This will hide the file HideFile.txt in the music file Waves.mp3

The server URL that Waves.mp3 will now be streamed on will be displayed. Open this URL in a web browser to listen to this music file - note that the stream is http based, so any browser warnings about lack of https security can be ignored. The client program will need to be given this URL to extract the hidden file.

Press CTRL+C or kill the stegstream-server process to stop streaming. The server can also be set to shut down at a given time and date by setting the AutoShutdown entry in the configuration file.

Things to consider

The hidden file is stored in the music file as plaintext. If additional security is required, it is recommended to encrypt the hidden file before hiding it in the music file.

Recommended encryption applications are:

The GNU Privacy Guard (GnuPG): https://gnupg.org/index.html
VeraCrypt: https://www.veracrypt.fr/en/Home.html

The hidden and music files are left on disk after the server program has shut down. It is the responsibility of the user to delete these files if secrecy is important. The files can be securely deleted when the server shuts down by setting the WipeAudio and WipeHidden entries in the configuration file. It is also possible to securely delete the hidden file after it has been hidden in the music file by setting the WipeAfterHide entry in the configuration file.

The larger the hidden file in relation to the music file, the more audible distortions will be noticed on streaming playback. The limit for a hidden file is 10% of the size of the music file, as this limit keeps the distortions to an acceptable level. If the hidden file is large, a large music file will be needed.

Usage of configuration file

Placing a file named StegstreamServerConfig.txt in the same directory as the executable allows changes to made to the server configuration.

It is possible to change the following:

The port number the server listens on:

Port=8080

Stream the music without hiding the hidden file in it:

StreamOnly=true

Hide the hidden file in the music file without starting the web server:

HideOnly=true

Wipe music file when server shuts down:

WipeAudio=true

Wipe hidden file when server shuts down:

WipeHidden=true

Set a date and time to automatically shut down the server:

Format: dd/mm/yyyy hh:ss

24 Hour clock

AutoShutdown=11/04/2024 12:02

Set to wipe hidden file after hiding the file in the music file:

WipeAfterHide=true


Any lines in the configuration file that start with a # character will be ignored.

See the release files for example configuration files.
