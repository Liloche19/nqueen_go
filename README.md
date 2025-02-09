# nqueen - Go

## A program that solves the nqueen problem in Go

The program try all possibility, and count the ones that follows the laws of the problem.

In order to have the maximum performances, the program is multithreaded at the size of the plate that the program is trying to solve.

More informations can be found [here](https://en.wikipedia.org/wiki/Eight_queens_puzzle).

## 🛠️ Installation

### Make sure to have the necessary dependencies :

**On Ubuntu / Debian :**

```bash
sudo apt install go
```

**On Archlinux :**

```bash
sudo pacman -S install go
```

### This is how to install the project :

Clone the repository :

```bash
git clone https://github.com/Liloche19/nqueen_go.git
cd nqueen_go
```

Then, compile the project :

```bash
go build
```

**(OPTIONAL) :**

After use, you can delete the executable file :

```bash
rm nqueen_go
```

## This is how to launch the project :

```bash
./nqueen_go size
```

The outpout will look like this :

```bash
Solutions : [number]
```

## Contribution

If you have any proposition, you can contact me. I will be happy to improve this projet.
