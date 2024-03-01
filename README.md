sacct-observer
==============
An observer that fetches the job history with the "sacct" command of the Slurm workload manager and saves it in a SQLite3 database.
Can be useful to integrate it in microservices which launch jobs in HPC clusters where Slurm is used as a job scheduler.
The reason why it saves the records in a light database, because usually the content of the "sacct" command gets deleted periodically.

# Table of content
- [Installation with a pre-built binary](#installation-prebuilt-binary)
- [Installation from source](#installation-source)
    - [1. Clone the repository](#clone-repository)
    - [2. Build the binary (latest Go compiler is required)](#build-binary)
    - [3. Make the binary available for the session's user.](#binary-to-path)
    - [4. (optional) Enable autocomplete for the CLI flags.](#enable-autocomplete)

- [Usage](#usage)

---

# Installation with a pre-built binary <a id="installation-prebuilt-binary" />
Download the corresponding binary to your system from the assets below and put it into your PATH to make it a globally available executable.

[Download binary here](https://github.com/heloint/sacct-observer/releases/latest)

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

---

# Installation from source <a id="installation-source"/>

*NOTE: Ensure that your system has a globally installed and available SSH and SQLITE3 client.*

## 1. Clone the repository <a id="clone-repository" />

```bash
git clone https://github.com/heloint/sacct-observer
```

## 2. Build the binary (latest Go compiler is required) <a id="build-binary" />

```bash
make build
```

*OR use Docker to build the binary*

```bash
docker build -t sacct-observer:latest . \
&& mkdir -p ./bin \
&& id=$(docker create sacct-observer:latest) \
&& docker cp $id:/app/bin/sacct-observer - > ./bin/sacct-observer \
&& docker rm -v $id
```

## 3. Make the binary available for the session's user. <a id="binary-to-path" />

```bash
make install
```

## 4. (optional) Enable autocomplete for the CLI flags. <a id="enable-autocomplete />

```bash
make get-autocomplete
```

---

# Usage <a id="usage" />
```bash
sacct-observer --username "my_remote_username" \
               --remote-address "my_cluster.org" \
               --output-sqlite-db "/path/to/create/my_sqlite.db" \
               --update-frequency 30
```
