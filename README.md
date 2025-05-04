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

4.  **Set Database Name & Generate Initial Migration:**
    Before generating the migration, ensure you update the database name in **two** places:
    *   Modify the **string literal** passed as the first argument to `sqldb.NewDatabase` in the `database/database.go` file (around line 15) from `"my_db_name"` to your desired database name (must be in `snake_case`).
    *   Modify the `DB_NAME` variable in the `database/scripts/generate-migration` script (around line 6) from `my_db_name` to the **same** desired database name used in `database.go`.

    **Important:** The database name must be identical (and in `snake_case`) in both `database/database.go` (as a string literal) and `database/scripts/generate-migration` (as a variable) for the migration process to work correctly.

    Also, configure how the script finds the Atlas executable in `database/scripts/generate-migration` (around lines 18-23):
    *   **If Atlas is NOT in your system's PATH:** Update the `ATLAS_PATH` variable (around line 22) with the full path to your `atlas.exe`.
        ```bash
        # Example:
        ATLAS_PATH="C:/path/to/your/atlas/bin/atlas-windows-amd64-latest.exe"
        "$ATLAS_PATH" migrate diff ... # Keep this line active
        ```
    *   **If Atlas IS in your system's PATH:** Comment out the `ATLAS_PATH` variable definition and the line that uses it, and uncomment the line that calls `atlas` directly.
        ```bash
        # Comment out these two lines:
        # ATLAS_PATH="C:/Users/cheah/atlas/bin/atlas-windows-amd64-latest.exe"
        # "$ATLAS_PATH" migrate diff ...

        # Uncomment this line (around line 21):
        atlas migrate diff "$MIGRATION_NAME" --env local --dev-url "$(encore db conn-uri --shadow $DB_NAME)&search_path=public"
        ```

    Then, use Git Bash (or a similar shell) to navigate to the `database` directory and run the script that generates the initial database migration based on your GORM models. Make sure you are in the project's root directory initially.

    ```bash
    cd database
    ../scripts/generate-migration init
    ```
    *Note: This assumes you have a script named `generate-migration` inside a `scripts` directory relative to the project root, and your GORM models/Atlas configuration are set up relative to the `database` directory.*

## Docker (POSTGRESQL related information)

1.  The default user is `postgres` and the password is `postgres`.
2.  Normally, the database is running on port `5432`.

## Running the Application

To run the Encore application, execute the following command in your project's root directory:

```bash
encore run
```


