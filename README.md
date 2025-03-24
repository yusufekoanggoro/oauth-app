# OAUTH APP

## 🚀 Overview

OAUTH-APP adalah aplikasi otentikasi dan otorisasi berbasis OAuth 2.0 yang mengamankan akses ke sumber daya backend dengan sistem token. Proyek ini terdiri dari dua bagian utama:

- **Backend**: Dibangun dengan Golang, menggunakan OAuth 2.0 sebagai metode autentikasi utama.
- **Frontend**: Aplikasi berbasis React yang menyediakan UI intuitif untuk pengguna dalam mengelola sesi login mereka.

## 📌 Features

✅ Implementasi OAuth 2.0 (Authorization Code Flow)\
✅ Dukungan JWT (JSON Web Token)\
✅ Integrasi dengan penyedia OAuth populer (Google)\
✅ UI interaktif untuk login/logout

## 🛠️ Tech Stack

### Backend:

- Go Fiber
- OAuth 2.0
- JWT

### Frontend:

Pastikan Anda menggunakan Node.js versi **18**.

- React.js
- Tailwind CSS

## ⚡ Getting Started

### 1️⃣ Clone Repository

```sh
git clone https://github.com/yusufekoanggoro/oauth-app.git
cd oauth-app
```

### 2️⃣ Setup Backend

```sh
cd backend
cp .env.example .env
go mod tidy
go run main.go
```

### 3️⃣ Setup Frontend

```sh
cd ../frontend
cp .env.example .env
npm install
npm run dev
```

## 🔑 OAuth Configuration

Tambahkan kredensial OAuth dari penyedia seperti Google atau GitHub di file `.env`:

```env
GOOGLE_OAUTH_CLIENT_ID=your-client-id
JWT_CLIENT_SECRET=your-client-secret

VITE_GOOGLE_OAUTH_CLIENT_ID=your-client-id
```

### 📌 Cara Mendapatkan Google OAuth Credentials

1. Buka [Google Cloud Console](https://console.cloud.google.com/).
2. Buat proyek baru atau pilih proyek yang ada.
3. Aktifkan **Google OAuth API**.
4. Buat kredensial OAuth 2.0 dengan tipe aplikasi **Web**.
5. Tambahkan `http://localhost:3000` ke daftar redirect URIs.
6. Salin **Client ID** ke file `.env`.

## 📜 API Endpoints

| Method  | Endpoint         | Description                                        |
| ------  | ---------------- | -------------------------------------------------- |
| `POST`  | `/auth/google`   | Menukar Authorization Code dengan Access Token     |


## 🎨 UI Preview


---

⭐ **Star repo ini jika menurut Anda bermanfaat!** 🚀

