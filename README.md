# Repository scanner tool

This project was thrown together to scan a list of repositories for commit messages for an outcome tag similar to the repository tool, may make a web interface based on this one day™.

## Usage

To use this program you should provide it with a file with a list of repository urls (1 per line) and a private key file.

I have provided an example [file](./repositories.txt) for reference.

### From source


```
go run . <repositoryListFile> <privateKeyFile>

>>

Scanning for commits in <repo>

Outcome 1.1.1.1 has 1 commits
Outcome 1.1.1.2 has 2 commits
Outcome 1.1.1.3 has 3 commits
Outcome 1.1.1.4 has 4 commits
Outcome 1.1.1.5 has 5 commits

... (for each module)

```

Where repositoryListFile is the path of your file containing a list of repository files and privateKeyFile being the path to your private key file.

For me the command looks like:

```
go run . repositories.txt C:/Users/Jordan/.ssh/id_rsa

```

### From executable 

If don't have go installed and you are naive enough to blindly execute an exe find it in the [releases](https://github.com/IllogicalTree/repository-scanner-tool/releases/) section.


``` 
scanner.exe <repositoryListFile> <privateKeyFile>
```

While I haven't done anything malicious I won't go into how bad of an idea it is to blindly run an exe file, if you would like an executable you can build from source using the following command and then run the exe file as above.

```
go build -o "scanner.exe" .
```

## Notes

~~Currently only one repository is scanned, the majority of mine are only in 1 repository but I may enhance this in future.~~
A file with a list of repositories is expected, if you have only one repository only use 1 line.

Ensure the repository URL provided is of the SSH (git@bitbucket) format as opposed to the http address.

WARNING - Do not blindly give programs access to your private SSH key file, review the code and ensure you are happy with it
While this program should not do anything malicious, it makes use of external (trusted and respected) packages that could potentially be doing anything with your private keys.
