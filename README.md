# File Organizer
It is a CLI developed in Go that allows organizing all files into directories by format in a directory passed as a parameter.

## Install 
1. Clone the repo
```bash
   git clone https://github.com/EzequielK-source/filesOrganizer
```
2. Navigate to the project directory:
```bash
   cd filesOrganizer
```
3. Compile the code
```bash
   go build -o organizer main.go
```

## How to use
The only thing you need to do to organize a directory is to run the executable, passing the directory path as the value of the 'dir' flag. In case you don't provide a directory address as a parameter, it takes the directory from where the command is being executed.
```bash
   ./organizer -dir=<directoryPath>
```
## Optional: Make 'fileOrganizer' Executable from Anywhere
If you want to be able to use the 'fileOrganizer' program in the terminal from anywhere.
### Linux 
Once created the build of fileOrganizer set the terminal in the directory of executable, and follow the next steps
1. Move the executable to /usr/local/bin
```bash
   sudo mv organizer /usr/local/bin/
```
2. Add the executable to the PATH in the .bashrc or .zshrc file:
```bash
   echo export PATH=$PATH:/usr/local/bin/organizer > .bashrc
   #For zsh
   echo export PATH=$PATH:/usr/local/bin/organizer > .zshrc
```
3. Apply the changes:
```bash
   source .bashrc
   #For zsh
   source .zshrc
```
### Windows 
1. Add the executable for 'fileOrganizer' to the PATH:
```bash
   setx PATH "%PATH%;<pathToExecutable>"
```
3. Restart the terminal or apply the changes.

These steps ensure that the 'fileOrganizer' executable is accessible from any location in the terminal.
