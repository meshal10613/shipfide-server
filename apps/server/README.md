# Shipfide Server API

A production-ready, high-performance REST API built with Go, Fiber v3, GORM, PostgreSQL, JWT Authentication, Cloudinary Image Storage, and SMTP transactional emails.

---

## Key Features

- **Framework**: [Fiber v3](https://github.com/gofiber/fiber) (highly optimized for performance and lightweight Go APIs).
- **ORM**: [GORM](https://gorm.io/) with a native PostgreSQL driver and auto-migrations for all 15 relational entities.
- **Cloudinary Image Upload & Auto-Cleanup**:
  - Direct Cloudinary integration for user avatars, rider document verification (NID, Driving License, Vehicle RC), and merchant NID upload.
  - **Automatic Image Replacement**: When updating an existing profile or document image, after successfully uploading the new image to Cloudinary, the previous image is automatically destroyed from Cloudinary storage.
- **Admin Domain & Hub Scope**:
  - Hubs represent physical operations locations created by Super Admins.
  - Administrative staff accounts (`RoleAdmin`) are created strictly by Super Admins and scoped to physical hubs.
  - **First-Time Password Change**: Admins are created with an automatic cryptographically secure temporary password and have `NeedsPasswordChange: true`. First-time login forces them to change their password via standard change password endpoints, which clears the flag.
- **3-Token Authentication System**:
  - `AccessToken` (JWT, duration configured via `JWT_DURATION` in `.env`, e.g. 30 days)
  - `RefreshToken` (UUID, 30-day expiry, tracked in GORM `Account` model)
  - `SessionToken` (UUID, 30-day expiry, tracked in GORM `Session` model for active device tracking)
- **Authenticating Options**: Reads tokens from `access_token` cookies, `token` cookies, or HTTP `Authorization: Bearer <token>` headers.
- **Shipment Management & State Engine**:
  - Full lifecycle shipment tracking from `PENDING` to `DELIVERED`, `FAILED_DELIVERY`, or `CANCELLED`.
  - Auto-generated unique tracking codes (`SF-YYYYMMDD-XXXXX`).
  - Automated revenue split calculation on delivery (`RiderShare`, `SystemShare`, `MerchantNet`).
- **Pricing & Zone Fee Engine**: Dynamic fee calculation based on 6 zone types (`INSIDE_DHAKA`, `DHAKA_SUBURB`, `SAME_CITY_NON_DHAKA`, `OUTSIDE_DHAKA_DIVISIONAL`, `OUTSIDE_DHAKA_DISTRICT`, `OUTSIDE_DHAKA_UPAZILA`), weight tiers, parcel types, and surcharges.
- **Receiver Fraud Score & Risk Assessment**: Dynamic fraud profile scoring (0–100) and automated COD blocking for high-risk receiver phone numbers based on delivery history.
- **Delivery Confirmation & COD Deposit**:
  - OTP verification for prepaid parcels with 48-hour expiration and 3-attempt locking mechanism.
  - Rider COD collection verification and hub cash deposit pipeline with Admin approval.
- **Merchant & Rider Profiles**: Full profile management, emergency contacts, operational zones, vehicle classification, and KYC approval workflows.
- **Ratings & Badging**: Receiver-to-rider delivery ratings (auto-updates rider average rating and badge status: `NEW_RIDER`, `TOP_RIDER`, `GOOD`, `AVERAGE`, `UNDER_REVIEW`) and merchant platform ratings.
- **Withdrawal / Cashout Engine**: Support for Rider and Merchant withdrawal requests with ৳100 minimum threshold enforcement.
- **SMTP Email Verification & Credentials Dispatch**:
  - Integrated SMTP transactional email delivery using EJS template parsing.
  - Asynchronous background email delivery to maintain low-latency HTTP responses.
- **Load Balancing (5 Replicas)**: Pre-configured **Nginx** reverse proxy load-balancing traffic across **5 separate backend API containers** on port 80.
- **Default Super Admin Seeding**: Automatically seeds a Super Admin user on startup if not already registered.

---

## Project Structure

```text
├── cmd/
│   └── api/
│       └── main.go                  # Application entry point
├── internal/
│   ├── config/                      # Config loader via godotenv
│   ├── database/                    # Postgres DB client setup
│   ├── domain/                      # Domain layers
│   │   ├── address/                 # Saved address management
│   │   ├── admin/                   # Admin CRUD logic, password generation, invitations
│   │   ├── auth/                    # Auth endpoints, login, register, password reset
│   │   ├── deliveryConfirmation/    # Delivery OTP verification and COD deposit approval
│   │   ├── guestSender/             # Walk-in guest sender registration
│   │   ├── hub/                     # Operations hub registration and mapping
│   │   ├── merchant/                # Merchant profiles, KYC, max COD limits
│   │   ├── rating/                  # Rider & platform rating submissions
│   │   ├── receiverFraud/           # Receiver fraud scoring and COD block management
│   │   ├── rider/                   # Rider profiles, vehicle info, hub assignment
│   │   ├── session/                 # Active user session revoking and tracking
│   │   ├── shipment/                # Shipment operations, pricing, tracking, state transitions
│   │   ├── upload/                  # Image upload to Cloudinary & auto-deletion
│   │   ├── user/                    # User profiles management
│   │   ├── walkInPayment/           # Hub walk-in payment collection records
│   │   └── withdrawal/              # Rider & Merchant cashout requests
│   ├── models/                      # GORM models (User, Account, Session, Hub, Admin, Address,
│   │                                #  Merchant, Rider, GuestSender, Shipment, deliveryConfirmation,
│   │                                #  codDeliveryConfirmation, walkInPayment, rating,
│   │                                #  receiverFraudProfile, withdrawal, Enums)
│   ├── routes/                      # Core route registration mapping
│   └── server/                      # Server bootstrapping & auto-migrations
├── nginx/
│   └── nginx.conf                   # Nginx load balancer configuration (Round-robin across 5 replicas)
├── pkg/                             # Shared core packages
│   ├── cloudinary/                  # Cloudinary upload client and public ID extraction
│   ├── email/                       # SMTP mail delivery & EJS template parsing
│   ├── httpResponse/                # Consistent success and error response shapes
│   ├── middlewares/                 # Authentication & Authorization middlewares
│   ├── pricing/                     # Shipping charge & revenue split calculator
│   ├── queryBuilder/                # Dynamic GORM database query builder
│   ├── seed/                        # Default data seeds (Super Admin user seeding)
│   ├── templates/                   # Transactional email layouts
│   ├── utils/                       # JWT, cookie, and tracking code utilities
│   └── validation/                  # Custom validator helpers
├── .air.toml                        # Air dev hot-reload settings
├── Dockerfile                       # Multi-stage Go build image configuration
├── docker-compose.yml               # Service composition for Nginx and 5 API backends
└── .env                             # Configured environment parameters
```

---

## Getting Started

### Option 1: Local Development

1. Ensure Go (1.26+) and PostgreSQL are installed and running.
2. Clone the repository and navigate to the project directory.
3. Set up environment variables in `.env`:
   ```env
   PORT=5000
   DSN="host=localhost user=postgres password=postgres dbname=shipfide port=5432 sslmode=disable TimeZone=Asia/Dhaka"
   JWT_SECRETKEY="your-jwt-secret-key-must-be-long-and-secure"
   JWT_DURATION=30d
   
   FRONTEND_URL="http://localhost:3000"

   SUPER_ADMIN_NAME="Shipfide Super Admin"
   SUPER_ADMIN_EMAIL="superadmin.shipfide@gmail.com"
   SUPER_ADMIN_PASSWORD="SuperAdmin@123"

   SMTP_HOST=smtp.mailtrap.io
   SMTP_PORT=2525
   SMTP_USERNAME=your_username
   SMTP_PASSWORD=your_password
   SMTP_FROM_EMAIL=no-reply@shipfide.com

   CLOUDINARY_CLOUD_NAME=your_cloud_name
   CLOUDINARY_API_KEY=your_api_key
   CLOUDINARY_API_SECRET=your_api_secret
   ```
4. Run code auto-reloader:
   ```bash
   air
   ```

### Option 2: Docker Compose (Nginx + 5 Replicas)

```bash
docker compose up --build -d
```
Access the API on port `80` (`http://localhost/`). Nginx automatically round-robins requests across the 5 container replicas.

---

## API Documentation

All request payloads and response bodies follow strict **camelCase** key naming conventions.

### Core & Health Check

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `GET` | `/` | Root health check | No |

### Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/auth/register` | Register new user (`PENDING` status) | No |
| `POST` | `/api/v1/auth/login` | User login with password | No |
| `POST` | `/api/v1/auth/verify-email` | Verify email with 6-digit OTP | No |
| `POST` | `/api/v1/auth/send-verification` | Resend email verification OTP | No (Optional) |
| `POST` | `/api/v1/auth/forgot-password` | Send password reset OTP | No |
| `POST` | `/api/v1/auth/verify-otp` | Validate OTP code | No |
| `POST` | `/api/v1/auth/reset-password` | Reset password using OTP | No |
| `POST` | `/api/v1/auth/refresh-token` | Obtain new access, refresh, and session tokens | No |
| `GET` | `/api/v1/auth/me` | Fetch authenticated user profile | Yes |
| `POST` | `/api/v1/auth/change-password` | Change password (clears `needsPasswordChange` flag) | Yes |
| `POST` | `/api/v1/auth/logout` | Revoke active session and clear cookies | Yes |

### Image & File Upload Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/upload/image` | Upload image to Cloudinary (optional `oldImageUrl` form param deletes previous image) | Yes |
| `DELETE` | `/api/v1/upload/image` | Delete image from Cloudinary by `?url=` query parameter | Yes |

### Admin Management Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/admins` | Create new admin staff with temporary credentials & auto welcome email | Yes (SUPER_ADMIN) |
| `GET` | `/api/v1/admins` | List admins with pagination, sorting, and search | Yes (SUPER_ADMIN) |
| `GET` | `/api/v1/admins/:id` | Get admin details by ID | Yes (SUPER_ADMIN) |
| `PUT` | `/api/v1/admins/:id` | Update admin settings & hub assignment | Yes (SUPER_ADMIN) |
| `DELETE` | `/api/v1/admins/:id` | Delete administrative user | Yes (SUPER_ADMIN) |

### Hub Management Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/hubs` | Create a physical operations hub | Yes (SUPER_ADMIN) |
| `GET` | `/api/v1/hubs` | List all hubs with paginated filters | Yes (ADMIN / SUPER_ADMIN) |
| `GET` | `/api/v1/hubs/:id` | Retrieve hub details by ID | Yes (ADMIN / SUPER_ADMIN) |
| `PUT` | `/api/v1/hubs/:id` | Update physical hub details | Yes (SUPER_ADMIN) |
| `DELETE` | `/api/v1/hubs/:id` | Delete hub | Yes (SUPER_ADMIN) |

### User Management Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `GET` | `/api/v1/users` | List users with search, page, limit, and sorting | Yes (ADMIN / SUPER_ADMIN) |
| `GET` | `/api/v1/users/:id` | Retrieve user by ID | Yes |
| `PUT` | `/api/v1/users/:id` | Update user profile (`name`, `phone`, `image`). Automatically deletes old avatar from Cloudinary when replaced. | Yes (Self / Admin / SuperAdmin) |
| `DELETE` | `/api/v1/users/:id` | Delete user | Yes (Self / Admin / SuperAdmin) |

### Session Management Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `GET` | `/api/v1/sessions` | List active device sessions for current user | Yes |
| `DELETE` | `/api/v1/sessions/:id` | Revoke a specific active session | Yes |

### Address Management Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/addresses` | Create a saved reusable address | Yes |
| `GET` | `/api/v1/addresses` | List saved addresses | Yes |
| `GET` | `/api/v1/addresses/:id` | Get address details by ID | Yes |

### Merchant Management Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `GET` | `/api/v1/merchants/me` | Get current merchant profile | Yes |
| `POST` | `/api/v1/merchants` | Create merchant profile | Yes |
| `GET` | `/api/v1/merchants` | List merchants | Yes (ADMIN / SUPER_ADMIN) |
| `GET` | `/api/v1/merchants/:id` | Get merchant profile by ID | Yes |
| `PUT` | `/api/v1/merchants/:id` | Update merchant profile (autocleans previous NID image on replace) | Yes |
| `PUT` | `/api/v1/merchants/:id/kyc` | Update merchant KYC verification & COD limits | Yes (ADMIN / SUPER_ADMIN) |

### Rider Management Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `GET` | `/api/v1/riders/me` | Get current rider profile | Yes |
| `POST` | `/api/v1/riders` | Create rider profile | Yes |
| `GET` | `/api/v1/riders` | List riders | Yes (ADMIN / SUPER_ADMIN) |
| `GET` | `/api/v1/riders/:id` | Get rider details by ID | Yes |
| `PUT` | `/api/v1/riders/:id` | Update rider details (autocleans previous NID, Driving License, or Vehicle RC images on replace) | Yes |
| `PUT` | `/api/v1/riders/:id/status` | Toggle rider active availability status | Yes |
| `PUT` | `/api/v1/riders/:id/kyc` | Update rider KYC review status | Yes (ADMIN / SUPER_ADMIN) |

### Guest Sender Management Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/guest-senders` | Register walk-in guest sender | Yes (ADMIN / SUPER_ADMIN) |
| `GET` | `/api/v1/guest-senders` | List guest senders | Yes (ADMIN / SUPER_ADMIN) |
| `GET` | `/api/v1/guest-senders/:id` | Get guest sender details | Yes (ADMIN / SUPER_ADMIN) |
| `PUT` | `/api/v1/guest-senders/:id/flag` | Flag or unflag guest sender phone | Yes (ADMIN / SUPER_ADMIN) |

### Shipment Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/shipments/calculate-price` | Calculate shipping fee for zone, weight, and parcel type | No |
| `GET` | `/api/v1/shipments/track/:trackingCode` | Track parcel status by tracking code | No |
| `POST` | `/api/v1/shipments` | Create new shipment (Merchant/Admin) | Yes |
| `GET` | `/api/v1/shipments` | List shipments with status, sender, and search filters | Yes |
| `GET` | `/api/v1/shipments/:id` | Get shipment details | Yes |
| `PUT` | `/api/v1/shipments/:id/status` | Update shipment status (triggers revenue split on `DELIVERED`) | Yes |
| `PUT` | `/api/v1/shipments/:id/assign-rider` | Assign rider to shipment | Yes (ADMIN / SUPER_ADMIN) |
| `PUT` | `/api/v1/shipments/:id/assign-hub` | Assign operations hub to shipment | Yes (ADMIN / SUPER_ADMIN) |

### Delivery Confirmation Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/delivery/otp/verify` | Verify delivery OTP code for prepaid shipment | Yes |
| `POST` | `/api/v1/delivery/otp/regenerate` | Regenerate delivery OTP | Yes (ADMIN / SUPER_ADMIN) |
| `POST` | `/api/v1/delivery/cod/confirm` | Confirm COD cash collection at handover | Yes |
| `POST` | `/api/v1/delivery/cod/deposit` | Submit collected COD cash at hub | Yes |
| `PUT` | `/api/v1/delivery/cod/deposit/:id/approve` | Approve rider COD deposit | Yes (ADMIN / SUPER_ADMIN) |

### Walk-In Payment Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/payments/walk-in` | Record hub walk-in parcel payment | Yes (ADMIN / SUPER_ADMIN) |
| `GET` | `/api/v1/payments/walk-in` | List walk-in payments | Yes (ADMIN / SUPER_ADMIN) |

### Receiver Fraud Profile Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `GET` | `/api/v1/fraud-profiles/check` | Check receiver fraud score and COD block status by phone | Yes |
| `GET` | `/api/v1/fraud-profiles` | List receiver fraud profiles | Yes (ADMIN / SUPER_ADMIN) |
| `GET` | `/api/v1/fraud-profiles/:id` | Get receiver fraud profile details | Yes (ADMIN / SUPER_ADMIN) |
| `PUT` | `/api/v1/fraud-profiles/:id/cod-status` | Update COD block status for receiver | Yes (ADMIN / SUPER_ADMIN) |

### Rating Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/ratings/rider` | Submit receiver-to-rider rating (1–5 stars) | No |
| `GET` | `/api/v1/ratings/rider/:riderId` | List ratings received by rider | Yes |
| `POST` | `/api/v1/ratings/merchant` | Submit merchant-to-platform rating | Yes (MERCHANT) |
| `GET` | `/api/v1/ratings/merchant` | List merchant platform ratings | Yes (ADMIN / SUPER_ADMIN) |

### Withdrawal Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/api/v1/withdrawals` | Submit cashout request (minimum ৳100) | Yes (RIDER / MERCHANT) |
| `GET` | `/api/v1/withdrawals` | List withdrawal requests | Yes |
| `GET` | `/api/v1/withdrawals/:id` | Get withdrawal details | Yes |
| `PUT` | `/api/v1/withdrawals/:id/status` | Update withdrawal status (`APPROVED`, `PAID`, `REJECTED`) | Yes (ADMIN / SUPER_ADMIN) |
