# 🚀 Xerador

**Xerador** is a fast and opinionated **Go-based CLI tool** for generating **NestJS monorepo projects**.

It helps you scaffold **repositories, projects (services), and modules** with a consistent structure, automatic wiring, and minimal manual setup.

---

## ✨ Features

- ⚡ Generate **NestJS monorepo** structure
- 📦 Create **project / service**
- 🧩 Create **module** with auto-registration
- 🔄 Automatically updates:
  - `app.module.ts`
  - `tsconfig.json` (paths)
  - `package.json` (jest moduleNameMapper)
- 🎨 Interactive CLI (select & prompt)
- 🟣 Clean task output (`✔ done`, `↷ skipped`)
- 🚀 Single binary (written in Go)

---

## 📦 Installation

### Using Go (Recommended)

```bash
go install github.com/yderana/xerador@latest
