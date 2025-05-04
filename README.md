# Encore.go + GORM + Atlas + PostgreSQL Project Template

This project demonstrates setting up a Go application using the Encore framework, GORM for database interaction, Atlas for database migrations, and PostgreSQL as the database.

## Prerequisites

*   Go installed
*   Encore framework installed (Refer to [Encore Documentation](https://encore.dev/docs/go/quick-start))
*   PostgreSQL installed and running
*   Atlas CLI installed (Download from [https://atlasgo.io/getting-started](https://atlasgo.io/getting-started))
*   Git Bash (or a similar Unix-like shell environment on Windows) for running shell scripts
*   Docker installed (Docker use for PostgreSQL instance)

## Setup Instructions

1.  **Install Go Dependencies:**
    Open your terminal and run the following commands to install the necessary Go packages:

    ```bash
    go get -u gorm.io/gorm gorm.io/driver/postgres
    go get github.com/asaskevich/govalidator
    go mod tidy
    ```

2.  **Verify Atlas Installation:**
    Ensure Atlas is correctly installed and accessible.

    *   **If Atlas is not in your system's PATH:**
        Replace the path below with your actual Atlas executable path. Run this command in your terminal:

        ```bash
        "C:\Users\cheah\atlas\bin\atlas-windows-amd64-latest.exe" version
        ```

    *   **If Atlas is added to your system's PATH:**
        You can simply run the following command in your terminal:

        ```bash
        atlas version
        ```

    *Note: Adding the Atlas bin directory to your system's PATH environment variable is recommended for easier access.*

3.  **Set Up Atlas Migration Script:**
    Configure your Atlas migration script as needed for your project structure and database connection. (Refer to Atlas documentation for details).

4.  **Generate Initial Migration:**
    Use Git Bash (or a similar shell) to navigate to the `database` directory and run the script that generates the initial database migration based on your GORM models. Make sure you are in the project's root directory initially.

    ```bash
    cd database
    ../scripts/generate-migration init
    ```
    *Note: This assumes you have a script named `generate-migration` inside a `scripts` directory relative to the project root, and your GORM models/Atlas configuration are set up relative to the `database` directory.*

## Running the Application

To run the Encore application, execute the following command in your project's root directory:

```bash
encore run
```
