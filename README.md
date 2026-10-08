# kmp-gen

> A CLI tool for scaffolding Kotlin Multiplatform projects — so you spend less time creating files and more time writing features.

---

## The problem

Starting a new feature in a Kotlin Multiplatform project means creating the same folder structure every time — domain models, repositories, use cases, data sources, ViewModels, Compose screens. It's repetitive, easy to get inconsistent, and adds friction every time a new feature needs to exist.

`kmp-gen` eliminates that entirely.

---

## What it does

One command scaffolds a complete, production-ready feature structure following **Domain Driven Design**:

```bash
kmp-gen feature auth
```

```
composeApp/src/commonMain/kotlin/com/yourapp/
└── feature/
    └── auth/
        ├── domain/
        │   ├── model/          Auth.kt
        │   ├── repository/     AuthRepository.kt
        │   └── usecase/        GetAuthUseCase.kt
        ├── data/
        │   ├── repository/     AuthRepositoryImpl.kt
        │   └── source/         AuthRemoteSource.kt
        │                       AuthLocalSource.kt
        └── presentation/
            ├── AuthViewModel.kt
            └── AuthScreen.kt
```

Every file is pre-populated with the correct package declarations and imports — ready to build on immediately.

---

## Getting started

### Install

```bash
go install github.com/jnawaz/kmp-gen@latest
```

### Initialise in your project

Run this once from your KMP project root:

```bash
kmp-gen init
```

`kmp-gen` will auto-detect your source root and walk you through a short setup, writing a `kmp-gen.yaml` config file.

### Scaffold a feature

```bash
kmp-gen feature <name>
```

That's it.

---

## Why Domain Driven Design?

DDD structures code around **business domains** rather than technical layers. The result is a codebase where:

- Each feature is self-contained and independently navigable
- The domain layer has zero external dependencies — fully testable and portable across platforms
- New team members can find their way around without a guide

`kmp-gen` enforces this structure consistently across every feature, every time.

---

## Built for Kotlin Multiplatform

`kmp-gen` understands KMP project conventions out of the box — `commonMain`, shared source sets, and Compose Multiplatform screen scaffolding are first-class citizens. No configuration gymnastics required.

---

## Extensible by design

The architecture is pluggable. DDD is supported today, with additional architecture patterns planned. One config line is all it takes to switch.

```yaml
architecture: ddd
```

---

## Tech

- Built in Go using [Cobra](https://github.com/spf13/cobra)
- Zero runtime dependencies — single binary, works anywhere
- Templates embedded directly in the binary — no external files needed

---

## Contributing

Issues and PRs welcome at [github.com/jnawaz/kmp-gen](https://github.com/jnawaz/kmp-gen).
