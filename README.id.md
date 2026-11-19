# Saragosa

**Saragosa** adalah platform pengelolaan IDS/IPS berbasis Linux yang dibangun di sekitar engine Suricata.

Saragosa menyediakan lapisan aplikasi untuk visibilitas sensor, pengumpulan event EVE, investigasi alert, pengelolaan ruleset, dan operasi yang tercatat dalam audit. Suricata tetap menjadi engine inspeksi paket.

> **Status proyek:** Fondasi awal. Kerangka dashboard, endpoint health API, dan parser EVE JSON berbasis stdin sudah tersedia. Penyimpanan database, koneksi sensor, autentikasi, dan IPS aktif belum diimplementasikan.

## Fondasi saat ini

- Kerangka dashboard Next.js dengan label status yang jujur dan bukan data live.
- Go API dengan endpoint `GET /api/v1/health`, default hanya listen di loopback.
- Fondasi collector Go yang memvalidasi EVE JSON Lines dari standard input dan melaporkan jumlah event.
- Unit test Go untuk kontrak health response dan parser event EVE.
- Dokumentasi arsitektur dan panduan pengembangan.
- GitHub Actions untuk Go tests serta typecheck/build web.

Health endpoint hanya melaporkan kondisi proses API. Collector belum membaca file Suricata secara langsung dan belum menyimpan event. Tidak ada metrik dashboard yang mengaku sebagai telemetry jaringan live.

## Teknologi

| Area | Teknologi |
|---|---|
| Detection engine | Suricata (integrasi direncanakan) |
| Backend dan collector | Go 1.27 |
| Dashboard web | Next.js 16, React 19, TypeScript |
| UI | Tailwind CSS 4 |
| Database | PostgreSQL (direncanakan) |
| Format event | Suricata EVE JSON |
| Target awal | Ubuntu Server Linux |
| Runtime | Node.js 26 LTS |

## Struktur repository

- `app/`: dashboard Next.js
- `services/api/`: Go HTTP API
- `services/collector/`: fondasi parser/collector EVE JSON
- `docs/PRD.md`: kebutuhan produk dan kriteria penerimaan per tahap
- `docs/architecture.md`: model deployment dan batas kepercayaan
- `docs/DEVELOPMENT.md`: panduan pengembangan lokal

## Pengembangan

Lihat [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md) untuk perintah dan prasyarat.

## Roadmap

1. Fondasi proyek
2. Lab Suricata IDS pasif dan verifikasi EVE
3. Collector persisten dan pipeline event
4. Alert API dan dashboard live
5. Autentikasi, RBAC, dan audit
6. Pengelolaan ruleset
7. Eksperimen IPS inline di lab terisolasi
8. Hardening dan release candidate

## Keamanan

- IDS pasif adalah mode default.
- IPS inline belum diimplementasikan dan hanya boleh diuji di lab terisolasi.
- API default listen pada `127.0.0.1:8080`.
- Jangan menjalankan aplikasi sebagai root.
- Jangan commit kredensial, private key, PCAP sensitif, atau secret produksi.
- Saragosa bukan pengganti NGFW, SIEM, EDR, atau NDR enterprise.

## Lisensi

Lihat [LICENSE](LICENSE) untuk lisensi MIT.
