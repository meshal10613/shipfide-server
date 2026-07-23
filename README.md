# Shipfide Turborepo Monorepo

Shipfide is an enterprise logistics and parcel delivery platform structured as a high-performance **Turborepo Monorepo**.

---

## 🏗️ Monorepo Structure

```text
shipfide/
├── apps/
│   ├── server/           # Go Fiber v3 REST API Backend (@shipfide/server)
│   ├── web/              # Next.js / React Web Application (@shipfide/web) [Placeholder]
│   └── app/              # React Native / Expo Mobile Application (@shipfide/app) [Placeholder]
├── packages/
│   ├── types/            # Shared TypeScript DTO interfaces (@shipfide/types)
│   └── config/           # Shared configuration settings (@shipfide/config)
├── turbo.json            # Turborepo task pipeline configuration
├── pnpm-workspace.yaml   # pnpm workspace packages definition
└── package.json          # Monorepo root scripts & dependencies
```

---

## 🚀 Getting Started

### Backend Server (`@shipfide/server`)

The Go REST API is located in [`apps/server`](file:///apps/server).

To run the backend server:

```bash
# Using root monorepo script
pnpm server:dev

# OR navigate directly into server directory
cd apps/server
air
```

### Running Tests

```bash
# Run tests across all monorepo apps
pnpm test

# Run Go server unit tests
pnpm server:test
```

### Building for Production

```bash
# Build all apps via Turborepo
pnpm build

# Build Go server binary (outputs to apps/server/bin/api)
pnpm server:build
```
