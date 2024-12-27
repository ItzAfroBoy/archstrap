<div align="center">
<pre>
 ______               __      ____    __                             
/\  _  \             /\ \    /\  _`\ /\ \__                          
\ \ \L\ \  _ __   ___\ \ \___\ \,\L\_\ \ ,_\  _ __    __     _____   
 \ \  __ \/\`'__\/'___\ \  _ `\/_\__ \\ \ \/ /\`'__\/'__`\  /\ '__`\ 
  \ \ \/\ \ \ \//\ \__/\ \ \ \ \/\ \L\ \ \ \_\ \ \//\ \L\.\_\ \ \L\ \
   \ \_\ \_\ \_\\ \____\\ \_\ \_\ `\____\ \__\\ \_\\ \__/.\_\\ \ ,__/
    \/_/\/_/\/_/ \/____/ \/_/\/_/\/_____/\/__/ \/_/ \/__/\/_/ \ \ \/ 
                                                               \ \_\ 
                                                                \/_/ 
<br>
Arch Linux bootstrapper
<br>
<img alt="GitHub License" src="https://img.shields.io/github/license/ItzAfroBoy/archstrap"> <a href="https://www.codefactor.io/repository/github/itzafroboy/archstrap"><img src="https://www.codefactor.io/repository/github/itzafroboy/archstrap/badge" alt="CodeFactor" /></a> <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/ItzAfroBoy/inv">
</pre>
</div>

## Installation  

### Install with go

```shell
go install github.com/ItzAfroBoy/archstrap@latest
archstrap ...
```

### Build from source

```shell
git clone https://github.com/ItzAfroBoy/archstrap
cd archstrap
go install
archstrap ...
```

## Usage

`Usage: archstrap [--skip-git] [--skip-pacman] [--skip-yay]`

- `--skip-git`: Skips cloning of git repos
- `--skip-pacman`: Skips installation of pacman packages
- `--skip-yay`: Skips installation of yay packages
