# New features
[placeholder]

---

# Modified features
[placeholder]

---

# Bug fixes
[placeholder]

---

# Installation from source
https://github.com/heloint/sacct-observer?tab=readme-ov-file#installation-source

---

# Using pre-built binaries
Download the corresponding binary to your system from the assets below and put it into your PATH to make it a globally available executable.

*NOTE: Ensure that your system has a globally installed and available SSH and SQLITE3 client.*

*EXAMPLE:*
```bash
mkdir -p ~/.local/bin
cp ./path/to/sacct-observer-binary ~/.local/bin/sacct-observer
chmod u+x ~/.local/bin/sacct-observer
```

**MAKE SURE THAT "~/.local/bin" is in your PATH!**
```bash
# Check your PATH variable.
echo $PATH | grep $HOME/.local/bin

# If no result printed as output, then:
echo "export PATH=\$PATH:$HOME/.local/bin" >> $HOME/.bashrc
source $HOME/.bashrc
```

*NOTE: If you want autocomplete for your bash shell, then run the following pipe:*
```bash
rm -r /tmp/sacct-observer \
&& git clone https://github.com/heloint/sacct-observer /tmp/sacct-observer \
&& cd /tmp/sacct-observer \
&& make get-autocomplete \
&& cd -
```


# SHA256 Checksums
```
# FORMAT EXAMPLE:
# ==============
db8fb2586b6a9e234903b9b776f6d1403f6bda71aee3b0c27229719b3045f0bc  sacct-observer-386.exe
25b4676cd8fdb5dc7df75ad18fcea4366a3093a75552be865258dd7ae07b2579  sacct-observer-amd64.exe
c7f7cff07e85d4fb4e8de5a8a627d8371dc0409c0170f2c338e3b0cb97c9b7a0  sacct-observer-amd64-darwin
ca8edc15bf45756cf1ef377b5dbd5b18a83bfb49369a96ac8b242d1c898342f9  sacct-observer-arm64-darwin
688ac9a852b1184dbff04cc0e704e6040f01edd873417488ef858a78794dcfb4  sacct-observer-amd64-linux
4d69b26bed1ace4e7b0869d7981b5879be29677e51a6e4007a5a78eb8ccdf4b0  sacct-observer-386-linux
```
