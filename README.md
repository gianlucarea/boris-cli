# Boris-cli

**Boris-cli** is a fun terminal application written in **Go** that prints ASCII art along with a random quote from the TV show *Boris* every time your terminal starts. Perfect for adding some humor and flair to your command-line experience!

---

## Features

* Shows a **ASCII art quote** from the cult Italian TV series *Boris*.
* Lightweight and easy to install.
* Written entirely in **Go**.

---

### From source

```bash
git clone https://github.com/yourusername/boris-cli.git
cd boris-cli
go build -o boris-cli
./boris-cli
```

---

## Usage

Simply run:

```bash
boris-cli
```

It will display an ASCII art of Boris with a random quote like:

```
   ____            _         _ 
  / ___| ___ _ __ (_) ___   | |
 | |  _ / _ \ '_ \| |/ _ \  | |
 | |_| |  __/ | | | | (_) | |_|
  \____|\___|_| |_|_|\___/  (_)
                           
```

---

## Make it run at terminal startup

Add the command to you lib

```
cd boris-cli
mv boris-cli /usr/local/bin
boris-cli
```

Then add the command to your shell config file:

### Bash (`~/.bashrc`)

```bash
boris-cli
```

### Zsh (`~/.zshrc`)

```zsh
boris-cli
```

### Fish (`~/.config/fish/config.fish`)

```fish
boris-cli
```

---

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes (`git commit -am 'Add feature'`)
4. Push to the branch (`git push origin feature/foo`)
5. Open a Pull Request

---
