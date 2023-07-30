<h1 align="center">Git Better Commit</h1>

<p align="center">
  <img src="./data/screenshots/gbc-v110-execution.gif" alt="gbc v1.1.0 execution">
</p>

<details>
  <summary>With emojis</summary>
  <p align="center">
    <img src="./data/screenshots/gbc-v110-execution-emojis.gif" alt="gbc v1.1.0 execution with emojis">
  </p>
</details>

------------

## :book: Description ##
**Git Better Commit (gbc) is a command line interface written in Golang that provides a simple way to write commits following the [Conventional Commits](https://www.conventionalcommits.org/).**

------------

## :computer: How to use

<details open>
  <summary><h3>Install gbc</h3></summary>
  
  Follow the next steps to install **gbc**:

  1. Make sure you already have [curl](https://curl.se/) installed on your machine;
  2. Paste this command in your terminal:
     ```powershell
     bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/AllanCapistrano/gbc/main/scripts/install.sh)"
     ```
  3. You can check if **gbc** is installed by typing `gbc --version` in your terminal.
</details>

<details>
  <summary><h3>Install gbc for Termux</h3></summary>
  
  Follow the next steps to install **gbc** for [Termux](https://termux.dev/en/):
  
  1. Make sure you already have [curl](https://curl.se/) installed;
  2. Paste this command in your terminal:
     ```powershell
     bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/AllanCapistrano/gbc/main/scripts/install-termux.sh)"
     ```
  3. You can check if **gbc** is installed by typing `gbc --version` in your terminal.
</details>

<details>
  <summary><h3>Using gbc</h3></summary>
  
  Once you have **gbc** installed, after adding the file(s) to commit, type `gbc` in your terminal, select the **commit type** using the **arrows keys**, and press **enter**. After that, enter the **commit message** you want.

  You can type `gbc --help` to see the allowed command options or `gbc help <commit type>` to see a short description and basic example of each **commit type**.
</details>

<details>
  <summary><h3>Changing the defaults emojis</h3></summary>
  
  You can change the emojis in the `gbc.conf` file, available in `$HOME/.config/gbc`. Just change the `emojis = {...}`. Note that you need to use Emojis Shortcodes, you can get them in [Emoji Picker](https://github-emoji-picker.rickstaa.dev/).
  
  ###### Make sure to follow the formatting, otherwise the change will not be applied. ######
</details>

<details>
  <summary><h3>Allowing commits messages with emojis</h3></summary>
  
  You can allow **commit messages with emojis** in the `gbc.conf` file, available in `$HOME/.config/gbc`. Just change the `enableEmojis` to `true`. 
  
  ###### Make sure to follow the formatting, otherwise the change will not be applied. ######

  The result will be something like:
  ```powershell
  🐛 fix: updateUser method
  ```
</details>

<details>
  <summary><h3>Build gbc</h3></summary>
  
  To build **gbc** in your machine, you can follow the next steps:
  
  1. Make sure you already have [Go](https://go.dev/) installed;
  2. Clone this repository or download the `.zip` file;
  3. Open the project directory in your terminal;
  4. Install the dependencies:
     ```powershell
     go mod download
     ```
     or
     ```powershell
     go mod tidy
     ```
  5. Then build the project:
     ```powershell
     go build -o bin/gbc main.go
     ```
     If you get [`GLIBC not found`](https://github.com/AllanCapistrano/gbc/issues/13) error, try:
     ```powershell
     CGO_ENABLED=0 go build -o bin/gbc main.go
     ```
  To use the settings file:
  1. Create the **gbc** config directory:
     ```powershell
     mkdir -p $HOME/.config/gbc
     ```
  2. Move the `gbc.conf` file into the configuration directory::
     ```powershell
     mv ./config/gbc.conf $HOME/.config/gbc
     ```
</details>

<details>
  <summary><h3>Uninstall gbc</h3></summary>
  
  To uninstall **gbc** run the following command in your terminal:
  
  ```powershell
  bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/AllanCapistrano/gbc/main/scripts/uninstall.sh)"
  ```
</details>

<details>
  <summary><h3>Uninstall gbc from Termux</h3></summary>
  
  To uninstall **gbc** from [Termux](https://termux.dev/en/) run the following command:
  
  ```powershell
  bash -c "$(curl --fail --show-error --silent --location https://raw.githubusercontent.com/AllanCapistrano/gbc/main/scripts/uninstall-termux.sh)"
  ```
</details>

------------

## :man_technologist: Author ##

| [![Allan Capistrano](https://github.com/AllanCapistrano.png?size=100)](https://github.com/AllanCapistrano) |
| -----------------------------------------------------------------------------------------------------------|
| [Allan Capistrano](https://github.com/AllanCapistrano)                                                     |

<p>
    <h3>My socials:</h3>
    <a href="https://github.com/AllanCapistrano">
        <img src="https://github.com/AllanCapistrano/AllanCapistrano/blob/master/assets/github-square-brands.png" alt="Github icon" width="5%">
    </a>
    &nbsp
    <a href="https://www.linkedin.com/in/allancapistrano/">
        <img src="https://github.com/AllanCapistrano/AllanCapistrano/blob/master/assets/linkedin-brands.png" alt="Linkedin icon" width="5%">
    </a> 
    &nbsp
    <a href="https://mail.google.com/mail/u/0/?view=cm&fs=1&tf=1&source=mailto&to=asantos@ecomp.uefs.br">
        <img src="https://github.com/AllanCapistrano/AllanCapistrano/blob/master/assets/envelope-square-solid.png" alt="Email icon" width="5%">
    </a>
</p>

------------

## :pray: Support ##

**Please :star: this repository if this project is useful or has helped you.**

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/allancapistrano)

------------

## :balance_scale: License ##
[GPL-3.0 License](./LICENSE)
