# VRC OSC Audio Controls

A go program to control your audio using VRChat OSC Parameters

## Usage

1. Download the executable from the [releases](https://github.com/shadorki/vrc-osc-audio-controls/releases) page and double-click to run. (You may have to restart the application if you started the server before VRChat)

1. Set your avatar up with the following avatar parameters. (The names must be exact)

    - OSC_AUDIO_CONTROLS_MUTE: bool
    - OSC_AUDIO_CONTROLS_PREVIOUS: bool
    - OSC_AUDIO_CONTROLS_NEXT: bool
    - OSC_AUDIO_CONTROLS_PLAY_PAUSE: bool

1. Add the buttons to your avatar menu and you should be good to go!


## Contributing

### Requirements
- Go
- make command

## Getting Started

1. Clone the repository
    ```bash
        git clone git@github.com:shadorki/vrc-osc-audio-controls.git
    ```
2. Run the program
    ```bash
        go run main.go
    ```
1. Compile the executable for windows
    ```bash
        make build
    ```
