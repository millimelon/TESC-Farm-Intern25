# Development Environment Setup
It helps to have a local environment to test your code. This document is meant to assist in getting a copy of the codebase running on your own computer for testing and development.

## Initial Setup
Before we can execute the code we need a shell to run it in. Essentially we need a terminal to run commands in. If you aren't familiar with the command line or text interfaces don't worry. Everything you need will need to type into the terminal will be listed on this page to be easily copied and pasted.

Executing commands on the terminal is just like sending a text: you compose or paste a short message and press 'return' to send it. The command line can be demanding in regards to spaces, capitalization and punctuation. Try to be patient with yourself, it gets easier the more you use it.

Once you have a terminal you can use it to install the required software:

- Git: used for tracking and merging changes to the project files. It allows us to work collaboratively, combining our edits together and keeping a history of changes, like a shared folder with an undo feature.
- Go (golang): Go is the programming language used to write the API that powers the backend. You need go installed in order to compile the api after making changes.
- NodeJS: Node is used to generate the files that compose the frontend. It's needed run a test version of frontend sites locally.

### Linux
Congratulations, Linux makes this simple. If you don't know how to open the terminal you can probably find it in your launcher menu by searching for `terminal`, there's usually also a hotkey to open the terminal: `Ctrl + Alt + t` works on many systems, including Ubuntu and Fedora.

Use the terminal to install git, golang, and nodejs:

- Debian/Ubuntu: `sudo apt install git golang nodejs`
- CentOS/RHEL: `sudo yum install git golang nodejs`
- Fedora: `sudo dnf -y install git go nodejs`
- Arch: `sudo pacman -S git go nodejs`

### Mac OS
Macs come with a great terminal, first use it to install Homebrew, which you'll use to get the packages required to run the development environment.

Open the terminal by searching for `terminal` in Spotlight (you can open Spotlight Search with `command + space`

Use the terminal to install Homebrew: `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"` (This uses the bash shell to download and run the install script)

Once Homebrew is installed you can use it to install git, golang, and nodejs: `brew install git golang nodejs`

### Android
The `termux` app is perfectly suitable for running the development environment. Simply install it from the Play Store (or the F-Droid app if you're cool)

Once Termux is installed, open the app and use the terminal to install git, golang, and nodejs: `pkg install git golang nodejs`

### Windows
It's technically possible to use the `cmd` application in Windows as a terminal, but the amount of extra effort needed to make it work is simply not worth it. Instead you should use WSL, the Windows Subsystem for Linux. WSL is an option from Microsoft that allows you to run a complete Linux terminal from within Windows.

Follow these steps to set up WSL, [you can find the official documentation here](https://learn.microsoft.com/en-us/windows/wsl/setup/environment):

1. Open a PowerShell window as administrator:
    - Open the Start menu
    - Search for `powershell`
    - Right-click on Windows PowerShell and select 'Run as administrator'
2. Run this command in PowerShell: `wsl --install`
3. Complete the installation and restart your machine
4. Open `Ubuntu` from the start menu to start setting up Linux
5. Pick a username and password for WSL
6. Update WSL: `sudo apt update && sudo apt upgrade`
7. Install git, golang, and nodejs: `sudo apt install git golang nodejs`

## Dev Environment

### Git
Navigate to your home directory (or wherever you want to put the code for the project): `cd ~` (**c**hange **d**irectory to ~, which means 'home')

Next, copy the repository from GitHub: `git clone git@github.com:absentbird/TESC-Farm.git`

Now change directory into the project: `cd TESC-Farm` and look at the files: `ls` (**l**i**s**t, or as I think of it **l**ook at **s**tuff), you should see all the same files as the main branch on GitHub.

### API
API files are located in the api directory, to enter the directory from TESC-Farm simply run `cd api`. The API is written in Go and is compiled into a binary to run on the server.

When running in development mode you can launch the API from the api directory with: `go run cmd/farmapi.go`

### Yarn
Yarn is used to manage modules for the frontend environments. Before using yarn it must be installed via NodeJS: `npm install yarn`

Once it's installed it can be used to get any necessary files by running `yarn install` in any one of the frontend directories, e.g. taskpanel.

To set up a new frontend directory you can run `yarn create vuetify` from the TESC-Farm directory.

### Database
The development database is stored in api/assets/data-dev.db in SQLite3 format. You should not need to access the database directly, but it can be done. To use the most recent data you can run `cp api/assets/data.db api/assets/data-dev.db` this turns the development database into a copy of the production database.

### Configuration

### Testing

### Launching

## Tools

### Scripts

### Vite

### Compiling

## Advanced Tools

### SOPS
