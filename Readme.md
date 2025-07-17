# CSV Checker – Robust CLI Tool for Validating CSV Files

CSV Checker is a fast, extensible, and professional-grade command-line tool written in Go. It is designed to validate the structure and content of CSV files before importing, processing, or integrating them into your data pipeline.

This project comes in two tiers: **Free** and **Pro**.\
The Free version includes basic validation functionality, while the Pro version adds advanced CLI control, report exporting, config support, and CI-ready integrations – all powered by the [Cobra](https://github.com/spf13/cobra) CLI framework.

---

## 🚀 Features

### ✅ Free Version (Open Source)

- Validates consistent column count across rows
- Detects empty fields
- Detects forbidden characters (excluding `@`)
- Simple CLI (`go run main.go file.csv`)
- No external dependencies

### 🔒 Pro Version (Coming Soon)

- Fully featured `Cobra` CLI with subcommands & flags
- Enable/disable rules via flags
- Export validation report (JSON or HTML)
- YAML/JSON config file support
- GitHub Actions / CI-ready integration
- Docker support
- REST API version for remote CSV validation (optional)

---

## 🧐 What It Checks

| Rule                     | Description                                                           |
| ------------------------ | --------------------------------------------------------------------- |
| Column Count Consistency | Ensures all rows have the same number of columns as the header        |
| Empty Field Detection    | Flags any fields that are empty                                       |
| Forbidden Characters     | Flags characters like `#`, `$`, `%`, `^`, etc. (Note: `@` is allowed) |

---

## 💻 Example (Free Version)

```bash
go run main.go testdata/sample.csv
```

Sample output:

```
Validation failed: row 3: column 1 is empty
```

---

## 📦 Installation

**Free Version (Local)**

```bash
git clone https://github.com/yourusername/csv-checker
cd csv-checker
go mod tidy
go run main.go testdata/sample.csv
```

**Pro Version (when released)**

```bash
go install github.com/yourusername/csv-checker@latest
csv-checker validate data.csv --output=report.json
```

---

## ✨ Planned CLI (Cobra – Pro)

```bash
csv-checker validate data.csv
csv-checker validate data.csv --skip-empty-check
csv-checker config init
csv-checker version
```

Subcommands will include:

- `validate` — Validate a CSV file
- `config` — Generate or load config file
- `version` — Show version number
- `report` — Export validation result (Pro only)

---

## 🧪 Sample CSV

```csv
name,email,age
John,john@example.com,30
Jane,jane@example.com,25
,missing@domain.com,22
Invalid,!@bademail.com,40
```

---

## 💰 Pricing Plans

| Plan    | Price | Includes                                                                |
| ------- | ----- | ----------------------------------------------------------------------- |
| Free    | \$0   | Basic CLI validation (column count, empty fields, forbidden characters) |
| Pro     | \$40  | Full-featured Cobra CLI, flag controls, report export, config support   |
| Premium | \$90  | REST API server, Docker support, CI/CD integrations                     |

Need a custom integration or private support?\
→ **Contact below.**

---


---

## 📅 Status

> The Free version is production-ready.
>
> The Pro version is actively under development.

---

## 📧 Contact

Questions, bugs, or feature requests?

- Email: [zviadi.mamardashvili.1@gmail.com](mailto\:zviadi.mamardashvili.1@gmail.com)
- GitHub: [github.com/ziko0/csv-checker](https://github.com/ziko0/csv-checker)

---

## 📄 License

**Free version** is released under the MIT License.\
**Pro & Premium versions** are licensed per user or company purchase.

