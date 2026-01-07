# 🚀 Xerador

**Xerador** adalah CLI tool berbasis **Go** untuk meng-generate **NestJS monorepo** dengan cepat dan konsisten.

Dirancang untuk workflow **create repo → create project → create module** tanpa ribet, lengkap dengan struktur folder, schema, repository, provider, dan integrasi otomatis ke `app.module.ts`.

---

## ✨ Fitur

- ⚡ Generate **NestJS monorepo** dengan struktur standar
- 📦 Create **project/service**
- 🧩 Create **module** (auto register ke app.module)
- 🧠 Update otomatis:
  - `app.module.ts`
  - `tsconfig.json`
  - `package.json` (jest paths)
- 🎨 CLI interaktif (select & prompt)
- 🟣 Status task rapi (`✔ done`, `↷ skipped`)
- 🧰 Ditulis full di **Go** (single binary, cepat & portable)

---

## 📦 Instalasi

### Menggunakan Go (Recommended)

```bash
go install github.com/yderana/xerador@latest
