# quote-cli

**quote-cli** is a fun terminal application written in **Go** that prints ASCII art along with a random quote every time your terminal starts. Perfect for adding some humour and flair to your command-line experience!

---

## Features

* Shows a **ASCII art quote**.
* Lightweight and easy to install.
* Written entirely in **Go**.

---

### From source

```bash
git clone https://github.com/yourusername/quote-cli.git
cd quote-cli
go build -o quote-cli
./quote-cli
```

---

## Usage

Simply run:

```bash
quote-cli
```

It will display an ASCII art of Quote with a random quote like:

```
 __   __                  ___              _         _   
 \ \ / /__  _   _ _ __   / _ \  ___  _   _| |_ ___  | |  
  \ V / _ \| | | | '__| | | | |/ _ \| | | | __/ _ \ | |  
   | | (_) | |_| | |    | |_| | (_) | |_| | ||  __/ |_|  
   |_|\___/ \__,_|_|     \__\_\\___/ \__,_|\__\___| (_)  
                           
```

---

## Make it run at terminal startup

Add the command to you lib

```
cd quote-cli
mv quote-cli /usr/local/bin
quote-cli
```

Then add the command to your shell config file:

### Bash (`~/.bashrc`)

```bash
quote-cli
```

### Zsh (`~/.zshrc`)

```zsh
quote-cli
```

### Fish (`~/.config/fish/config.fish`)

```fish
quote-cli
```

---

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes (`git commit -am 'Add feature'`)
4. Push to the branch (`git push origin feature/foo`)
5. Open a Pull Request

---
