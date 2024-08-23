# EnvWipe - Automated Python Virtual Environment Cleaner

This project provides a Golang based tool to automatically identify, delete, and manage old Python virtual environments (`venvs`) on your local computer. It also logs the paths of deleted environments and manages these logs by periodically cleaning up older entries.

## Features

- **Automatic Identification of Virtual Environments:**
  - Scans specified directories to find Python virtual environments (`venv` or `.venv` folders).
  - Deletes virtual environments that are older than a user-defined threshold (e.g., 30 days).

- **Logging:**
  - Logs the paths of all deleted virtual environments.
  - Logs are stored in a specified directory and are timestamped for easy reference.

- **Automated Log Cleanup:**
  - Periodically cleans up old log files based on user-defined retention policies.
  - Provides a summary report after each cleanup, showing the logs retained and those deleted.

- **Summary Reporting:**
  - After each run, generates a summary report detailing:
    - The paths of deleted virtual environments.
    - The number of log files cleaned up.
    - Any errors encountered during the process.

## Installation

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/Ashad/envwipe.git
   cd envwipe
   ```

2. **Build the Project:**
   ```sh
   go build -o envwipe main.go
   ```

3. **Configure the Script:**
   - Modify the `config.json` file to specify the directories to scan, the age threshold for deleting virtual environments, and log retention policies.

## Usage

1. **Run the Script:**
   ```sh
   ./envwipe
   ```

2. **Dry Run Mode:**
   - To see what would be deleted without actually deleting anything:
     ```sh
     ./envwipe --dry-run
     ```

3. **Specify a Custom Config File:**
   - Use a custom configuration file:
     ```sh
     ./envwipe --config /path/to/your/config.json
     ```

## Configuration

The `config.json` file allows you to customize the behavior of the script. Key settings include:

- **Directories to Scan:** Specify the directories where the script should look for virtual environments.
- **Age Threshold:** Define how old a virtual environment must be before it is considered for deletion.
- **Log Retention Policy:** Set how long logs should be kept before they are automatically deleted.

Example `config.json`:
```json
{
  "scanDirectories": [
    "/path/to/projects",
    "/another/path/to/scan"
  ],
  "thresholdDays": 30,
  "logDirectory": "/path/to/logs",
  "logRetentionDays": 60
}
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License.
