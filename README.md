# GitHub-Sync (ghs)

Ever coded daily for hours and finally pushed your changes after a while, and GitHub was like: **"Oh, so you only worked today, huh?"**
And it counts as a single contribution for the whole week! 
GitHub activity tracking is a bit... let's say, "unreliable" (*cough* unfair *cough*)
And the contribution graph might be the only way for others to know your coding activity.


That's where ghs comes to the rescue

## Overview
GitHub-Sync is a command-line tool designed to accurately track your local coding activity and automatically push contributions to your GitHub repository. It ensures your GitHub activity graph reflects your actual work by logging coding sessions and committing progress at regular intervals.

### How It Works?
When you first run ghs start, the tool prompts you for a Personal Access Token (PAT), your current activity, and the desired commit frequency. A background process then monitors whether your editor is open. If it detects active coding, it increments the tracker until the set time interval is reached. Once that happens, it automatically commits your progress with an informational message and appends the details to log.txt, before pushing the update to GitHub.

---

## Setting Up Your GitHub Personal Access Token (PAT)
To use GitHub-Sync, you need to generate a GitHub PAT (Personal Access Token). The tool uses this token **locally** to authenticate and push commits automaticly on your behalf.

The tool require PAT with only these permision:
- repo, user, read:org.

When you run GitHub-Sync for the first time, it will ask for this token. The token is stored **only on your local machine** and is never transmitted elsewhere.

**PAT will be stored in plain text in config.json so make sure to use this tool only on your personal device and it has only the basic permisions specified above**
---

## Installation

### Using Go
```sh
# Install GitHub-Sync
go install github.com/mostafa-mahmood/GitHub-Sync
```

### Manual Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/mostafa-mahmood/GitHub-Sync.git
   ```
2. Navigate into the directory:
   ```sh
   cd GitHub-Sync
   ```
3. Build the binary:
   ```sh
   go build -o ghs.exe
   ```
4. 
   ```sh
   go install
   ```
---

### Commands:

#### start
```sh
ghs start
```
- starts tracking background process.
- Asks for your GitHub PAT (if not already set).
- Creates necessary configuration files.
- Prompts for:
  - **Activity description** (e.g., "Building authentication system")
  - **Commit frequency** (minimum 100 minutes)
- Detects active code editor (e.g., VS Code, Vim, JetBrains IDEs).
- Begins tracking coding activity every 5 minutes.

#### stop
```sh
ghs stop
```
- Stops the background tracking process.

#### reset
```sh
ghs reset
```
- Resets tracked session time to zero but keeps the total session history.

#### status
```sh
ghs status
```
- Displays current tracked minutes, total session time, commit count, and last update.

#### config
```sh
ghs config // Displays configuration data
ghs config --activity=<""> // Update Activity (string)
ghs config --frequency=< > // Update Commit Frequency (int)
ghs config --pat=<""> // Update PAT (string)
```
- Use flags to update configuration data.
---

## Potential Improvements
I welcome contributions and suggestions! Some ideas for future improvements:
- Enhanced editor detection for more coding environments.
- Custom commit templates.

If you have ideas or want to contribute, feel free to open an issue or PR!

---

## 📜 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

