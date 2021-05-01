# Repository scanner tool

This project was thrown together to scan a given repository for commit messages for an outcome tag similar to the repository tool, may make a web interface based on this one day™.

## Usage

### From source

```
go run . <repoUrl> <privateKeyFile>

>>

Outcome 1.1.1.1 has 1 commits
Outcome 1.1.1.2 has 2 commits
Outcome 1.1.1.3 has 3 commits
Outcome 1.1.1.4 has 4 commits
Outcome 1.1.1.5 has 5 commits

... (for each module)

```

For me the command looks like:

```
go run . git@bitbucket.org:15027887/personal-reflections.git C:/Users/Jordan/.ssh/id_rsa

```

### From executable 

If don't have go installed and you are naive enough to blindly execute an exe find it in the [releases](https://github.com/IllogicalTree/repository-scanner-tool/releases/) section.


``` 
scanner.exe <repoUrl> <privateKeyFile>
```

While I haven't done anything malicious I won't go into how bad of an idea it is to blindly run an exe file, if you would like an executable you can build from source using the following command and then run the exe file as above.

```
go build -o "scanner.exe" .
```

## Notes

Currently only one repository is scanned, the majority of mine are only in 1 repository but I may enhance this in future.

Ensure the repository URL provided is of the SSH (git@bitbucket) format as opposed to the http address.

WARNING - Do not blindly give programs access to your private SSH key file, review the code and ensure you are happy with it
While this program should not do anything malicious, it makes use of external (trusted and respected) packages that could potentially be doing anything with your private keys.
