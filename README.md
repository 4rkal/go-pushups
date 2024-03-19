# go-pushups
go-pushups is a simple command-line application designed to help you stay active by reminding you to do pushups at regular intervals, with a percentage increase over time.

# Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Commands](#commands)
- [Contributing](#contributing)
- [Donate](#donate)
- [License](#license)

## Installation

To install **go-pushups**, you need to have Go installed on your system. Then, you can simply run:

```bash
go install github.com/4rkal/go-pushups@latest
```

**Alternatively** download it from the [release page](https://github.com/4rkal/go-pushups/releases) 

## Usage
After installing or building the application, you can run it from the command line. Here's how to use it:
```bash
go-pushups [command] [flags]
```

## Commands
- help: Provides help about any command within the application.
```bash
go-pushups help [command]
```

- load: Loads an existing routine
Running: 
```bash
go-pushups load
```
Will prompt you to select one of the save routines

Alternatively to load a specific routine:
```bash
go-pushups load [routine-name]
```
Replace [routine-name] with the name of the routine you want to load.


- run: Runs the go-pushups application,
```bash
go-pushups run
```

- new: will create a new routine
```bash
go-pushups new
```
Alternatively

Create a new routine with a custom (non random) name

```bash
go-pushups new [name]
```
Replase [name] with the actual name you want to use

## Contributing
Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue on or submit a pull request.

## Donate
If you liked this project here are the ways that you can support me

Direct [monero](https://getmonero.org) donation:
```
8A1ympBtgTUYER42HqyJEHKrbrFD2Q94RMkhmzzWSSvBPJrWVNa3cbk4YZhe1DgPnMY35zcuuen8x58siq5D7uVRUDUZLzm
```

Other methods [here](https://4rkal.eu.org/donate/)

## License
This project is licensed under the MIT License.
