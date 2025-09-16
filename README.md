# Central DevSecOps Pipeline

This repository provides a **centralized DevSecOps pipeline** for all teams, powered by [Dagger](https://dagger.io/) and Go.  
It enables consistent, secure, and scalable CI/CD workflows for projects in Python, Go, Node.js, and more.

---

## Features

- **Central pipeline logic:** All build, test, and security steps are defined and maintained here.
- **Multi-language support:** Easily extendable for Python, Go, Node.js, etc.
- **Remote execution:** Teams can run the pipeline without cloning this repo.
- **GitHub Actions integration:** Teams can call the pipeline from their own CI workflows.

---

```sh
dagger call build-default-python-image-and-publish \
  --image-tag latest \
  --github-username munichbughunter \
  --github-token env://GITHUB_TOKEN
```

---

## Usage for Dev Teams

### 1. Prerequisites

- [Dagger CLI](https://docs.dagger.io/cli/) installed (`brew install dagger` on macOS, see docs for other OS).
- Your project contains the source code (e.g., `app.py` for Python) in the root directory.

---

### 2. Run the Central Pipeline Locally

You **do not need to clone this repository**.  
Simply run the following command in your project directory:

```sh
dagger call --mod github.com/munichbughunter/central-devsecops-pipeline build-python --src=. --pythonVersion=3.12
```

- `--mod github.com/munichbughunter/central-devsecops-pipeline` loads the pipeline from GitHub.
- `build-python` is the function for Python builds (other languages available).
- `--src=.` points to your current directory.
- `--pythonVersion=3.12` sets the Python version (optional, defaults to 3.11).

---

### 3. Integrate with GitHub Actions

Add the following to your `.github/workflows/dagger.yml`:

```yaml
name: dagger
on:
  push:
    branches: [main]

jobs:
  build-python:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Call Central Dagger Function
        uses: dagger/dagger-for-github@8.0.0
        with:
          version: "latest"
          module: github.com/munichbughunter/central-devsecops-pipeline
          args: build-python --src=. --pythonVersion=3.12
          cloud-token: ${{ secrets.DAGGER_CLOUD_TOKEN }}
```

- Make sure to add your Dagger Cloud token as a secret named `DAGGER_CLOUD_TOKEN` in your repo settings.

---

## Supported Functions

- `build-python` – Build and run Python apps
- `build-go` – Build and test Go apps
- `build-node` – Build and lint Node.js apps

See the [pipeline source code](ci/build.go) for details and extend as needed.

---

## How It Works

- The pipeline mounts your source code into a container for the specified language.
- Executes build, test, and security steps in a reproducible environment.
- Results are printed to your terminal or CI logs.

---

## Extending the Pipeline

To add support for more languages or custom steps, contribute to this repository or request new features from the platform team.

---

## Troubleshooting

- Ensure your source files (e.g., `app.py`) are in the directory specified by `--src`.
- Use the debug step (`ls -l /src`) in the pipeline to verify file presence if needed.
- For help, open an issue or contact the platform engineering team.

---

## License

MIT

---

**Questions?**  
Contact the platform team or open an issue in this repository.