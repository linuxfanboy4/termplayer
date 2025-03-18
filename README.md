# TermPlayer - A Command-Line Music Player

TermPlayer is a sophisticated, lightweight, and highly customizable command-line music player designed for those who demand precision and control over their audio experience. Built with Go, TermPlayer offers a seamless integration with your terminal environment, enabling you to manage playlists, control playback, and fine-tune audio settings with ease. Whether you're a developer, a system administrator, or a power user, TermPlayer empowers you to enjoy music without leaving the command line.

## Features

- **Play Music**: Play audio files directly from the terminal.
- **Playlist Management**: Create, switch, and manage multiple playlists with ease.
- **Shuffle and Loop**: Toggle shuffle and loop modes for dynamic playback.
- **Volume Control**: Adjust volume levels with precision.
- **Pause and Resume**: Pause and resume playback at any time.
- **Next and Previous Tracks**: Navigate through your playlist effortlessly.
- **Mute Functionality**: Instantly mute or unmute audio playback.
- **Cross-Platform**: Works on any system with Go and `mpv` installed.

## Installation

To install TermPlayer, simply run the following command in your terminal:

```bash
curl -s https://raw.githubusercontent.com/linuxfanboy4/termplayer/refs/heads/main/src/termplayer.go | bash go run termplayer.go
```

This command will download and execute the TermPlayer script, allowing you to start using the player immediately.

## Usage

TermPlayer is designed to be intuitive and powerful. Below is a comprehensive guide to its commands and functionalities.

### Commands

- **play `<filename>`**: Play the specified audio file.
- **list**: Display the current playlist.
- **add `<filename>`**: Add a file to the current playlist.
- **remove `<filename>`**: Remove a file from the current playlist.
- **create `<playlist_name>`**: Create a new playlist.
- **switch `<playlist_name>`**: Switch to the specified playlist.
- **shuffle**: Toggle shuffle mode.
- **pause**: Pause the current playback.
- **resume**: Resume the paused playback.
- **volume `<level>`**: Set the volume level (0-100).
- **loop**: Toggle loop mode.
- **next**: Play the next track in the playlist.
- **prev**: Play the previous track in the playlist.
- **mute**: Toggle mute functionality.
- **quit**: Exit TermPlayer.

### Example Usage

1. **Playing a Song**:
   ```bash
   play song.mp3
   ```

2. **Creating and Switching Playlists**:
   ```bash
   create my_playlist
   switch my_playlist
   add song1.mp3
   add song2.mp3
   list
   ```

3. **Shuffling and Looping**:
   ```bash
   shuffle
   loop
   ```

4. **Adjusting Volume**:
   ```bash
   volume 75
   ```

5. **Pausing and Resuming**:
   ```bash
   pause
   resume
   ```

6. **Navigating Tracks**:
   ```bash
   next
   prev
   ```

7. **Muting**:
   ```bash
   mute
   ```

8. **Exiting TermPlayer**:
   ```bash
   quit
   ```

## Dependencies

TermPlayer relies on the following dependencies:

- **Go**: Ensure you have Go installed on your system.
- **mpv**: TermPlayer uses `mpv` for audio playback. Install `mpv` using your package manager.

### Installing Dependencies

- **Go**: Follow the official [Go installation guide](https://golang.org/doc/install).
- **mpv**: Install `mpv` using your package manager:
  - **Debian/Ubuntu**: `sudo apt-get install mpv`
  - **Fedora**: `sudo dnf install mpv`
  - **macOS**: `brew install mpv`

## License

TermPlayer is licensed under the **MIT License**. This permissive license allows you to use, modify, and distribute the software with minimal restrictions. For more details, see the [LICENSE](LICENSE) file.

## Contributing

We welcome contributions from the community. Whether you're fixing bugs, adding features, or improving documentation, your efforts are appreciated. Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Submit a pull request with a detailed description of your changes.

## Support

If you encounter any issues or have suggestions for improvements, please open an issue on the [GitHub repository](https://github.com/linuxfanboy4/termplayer). Your feedback is invaluable in making TermPlayer better.

## Acknowledgments

TermPlayer was developed with the goal of providing a robust and efficient command-line music player. Special thanks to the Go community and the developers of `mpv` for their incredible tools and libraries.

---

TermPlayer is more than just a music player; it's a testament to the power and flexibility of the command line. With its rich feature set and intuitive interface, TermPlayer is the ultimate choice for anyone who values control and efficiency in their audio experience. Start using TermPlayer today and elevate your command-line music experience to new heights.
