# Chaos Tasks
A simple party game that assigns every guest a random and secret task he or she has to fulfill. Inspired by [Chaoskarten](https://www.chaoskarten.de/).
| Start | Task received |
|-------|---------------|
|![Start Screen](doc/img/start.png) | ![Task Screen](doc/img/task.png) |

## Information
- No JavaScript
- Everyone should use his/her own phone. Application uses cookies to show own task.

### Software Requirements
- postgresql (13.13 is tested)
- go (1.15 is tested)
- bash

## Setup
1. Edit `etc/chaostasks-env.sh` with the details of your postgres database you want to create/use.
2. Create the database with its content execute `scripts/1-create-db.sh`.
3. Import tasks from a text file:
    ```bash
    ./2-impot-tasks.sh <Path>
    ```
    Example tasks (in german) are stored in `example_tasks_german.txt`.
    One line per task.
4. Compile the program and start it:
    ```bash
    cd src/chaostasks/
    go build
    ./chaostasks
    ```
5. Browse to `http://127.0.0.1:3000`

## Things planned
- German version
- docker build
- user management
- site customisation
- task management in web frontend.
