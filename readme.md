
# Radio Garden - Terminal UI Based


This project is a terminal application written in Go for listen all available [Radio.Garden 🌏](https://radio.garden/) stations.

| ![Example](https://i.imgur.com/a9WyNEO.png) |
|:--:| 
| *Theme: Black and White* |

## Installation 💻

Garden-Tui uses [Beep](https://github.com/faiface/beep) , which uses Oto under the hood to interact with speakers across multi platforms. Check [Oto Documentation](https://github.com/ebitengine/oto) to know if you need all necessary dependencies and how to install them.

### Go Install 🌸

You can use `go install` to easily compile and install the package and use the `garden-tui` command directly in your terminal emulator.

```
$ go install github.com/vergonha/garden-tui@latest
```

### Build 🔨

1 - Clone the project in your machine. 

```
$ git clone https://github.com/vergonha/garden-tui
```

2 - Install required dependencies

```
# On Ubuntu
apt-get install libasound2-dev
```
```
# On Arch
pacman -S alsa-lib
```

3 - Build Project 
```
$ cd garden-tui 
$ go build .
```
⚠ If you get the `undefined newDriver` error, try to `export CGO_ENABLED=1`. May be required for building in Linux distributions.

## Usage

Use your keyboard to interact with interface. Check the keybindings below.

|           Key            | Action                                                     |
| :----------------------: | ---------------------------------------------------------- |
| <kbd>p</kbd> | Pause                                       |
|       <kbd>s</kbd>       | Search |
|       <kbd>+</kbd>       | Increase Volume  |
|       <kbd>-</kbd>       | Decrease Volume  |
|       <kbd>&lt;ENTER&gt;</kbd>       | Select Station |




### Tasks

- [x] Add volume controls.
- [ ] Improve interface usability.
- [ ] Review code. 
