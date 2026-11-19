# SARAGOSA
## Product Requirements Document (PRD)

- **Versi:** 1.0.0
- **Status:** Draft untuk pengembangan bertahap
- **Jenis:** Network Intrusion Detection and Prevention System (NIDS/NIPS)
- **Target awal:** Ubuntu Server Linux
- **Bahasa:** Indonesia

## 1. Ringkasan Produk

Saragosa adalah platform pengelolaan IDS/IPS untuk Linux. Suricata menjadi engine inspeksi jaringan; Saragosa menyediakan lapisan aplikasi untuk konfigurasi, pemantauan sensor, ingest event, investigasi alert, pengelolaan ruleset, audit, dan respons yang terkontrol.

Saragosa tidak membuat packet-inspection engine sendiri dan bukan pengganti NGFW, SIEM, EDR, atau NDR enterprise.

## 2. Masalah dan Pengguna

Administrator membutuhkan satu antarmuka untuk memantau sensor, menyelidiki event jaringan, mengelola aturan secara aman, dan mengaktifkan pencegahan tanpa memutus jaringan secara tidak sengaja.

Peran sasaran:
- Network Administrator: status sensor, trafik, konfigurasi, dan rules.
- Security Analyst: alert, filter, investigasi, dan riwayat event.
- System Administrator: instalasi, service health, backup, dan pembaruan.
- Auditor: riwayat perubahan dan identitas pelaksana.

## 3. Tujuan dan Batasan

### Tujuan v1
- Dashboard status API, collector, database, dan Suricata.
- Membaca EVE JSON dan menyimpan event aktual.
- Daftar/detail alert dengan filter dan pagination.
- Import serta validasi ruleset sebelum aktivasi.
- Autentikasi, RBAC, dan audit log.
- IDS pasif sebagai default.
- Rancangan IPS inline eksperimental yang opt-in dan hanya diuji di lab.

### Bukan tujuan v1
- Membuat engine inspeksi paket sendiri.
- Menjanjikan deteksi semua zero-day atau malware.
- Menggunakan AI/ML untuk memutuskan pemblokiran.
- Menggantikan firewall/NGFW/SIEM/EDR/NDR.
- Menjamin throughput tanpa benchmark perangkat keras.
- SSL/TLS interception otomatis.
- Pemblokiran otomatis tanpa kebijakan administrator.

## 4. Prinsip Produk

1. IDS-first: mode pasif adalah default.
2. Fail-safe: kegagalan dashboard tidak mengubah jalur forwarding.
3. Least privilege: aplikasi tidak berjalan sebagai root kecuali helper terbatas yang dibutuhkan.
4. Explicit action: IPS, pemblokiran, dan perubahan rules memerlukan otorisasi administrator.
5. Observable: operasi penting memiliki status, log, dan alasan kegagalan.
6. No false claims: UI hanya menampilkan fitur yang benar-benar tersedia.
7. Reversible: perubahan konfigurasi divalidasi dan dapat di-rollback.
8. Secure by default: secret tidak masuk source control.

## 5. Arsitektur dan Stack Usulan

```text
Network Interface / PCAP
        |
   Suricata Engine
   (IDS default, IPS opt-in)
        |
      EVE JSON
        |
 Saragosa Collector (Go)
        |
   PostgreSQL
        |
 Backend API (Go)
        |
 Next.js + TypeScript Dashboard
```

| Lapisan | Teknologi | Peran |
|---|---|---|
| Detection engine | Suricata | Packet inspection dan signature detection |
| Collector/API | Go | EVE ingestion, validasi, API, health, audit |
| Frontend | Next.js + TypeScript | Dashboard responsif |
| UI | Tailwind CSS + Lucide React | Komponen antarmuka |
| Database | PostgreSQL | Event, pengguna, konfigurasi, audit |
| Event format | Suricata EVE JSON | Kontrak event |
| Target awal | Ubuntu Server + systemd | Operasional awal |
| Lab opsional | Docker Compose | Lingkungan pengembangan terisolasi |

Versi dependency harus diverifikasi dan dikunci ketika implementasi dimulai. Jangan mengarang versi.

## 6. Kebutuhan Fungsional

### F-01 Status sistem
- Tampilkan status API, collector, database, dan Suricata.
- Bedakan `healthy`, `degraded`, `unavailable`, dan `unknown`.
- Sertakan waktu pemeriksaan terakhir.
- API hidup tidak boleh otomatis berarti engine sehat.

### F-02 Event ingestion
- Baca EVE JSON secara streaming.
- Dukung minimal event `alert`, `flow`, `dns`, dan `http` jika tersedia.
- Validasi JSON dan tipe event.
- Baris rusak tidak boleh menghentikan seluruh collector.
- Simpan timestamp, sensor, source/destination IP dan port, protokol, signature, severity, serta metadata yang tersedia.
- Tangani duplikasi secara idempotent bila identitas event mendukung.
- Catat jumlah diterima, ditolak, dan gagal disimpan.

### F-03 Alert dashboard
- Daftar dan detail alert dengan pagination.
- Filter waktu, severity, signature, source IP, destination IP, dan protokol.
- Jangan menampilkan field rekaan.
- Pemetaan severity harus terdokumentasi.

### F-04 Rule management
- Tampilkan sumber dan status ruleset.
- Import melalui jalur yang diizinkan dan dengan batas ukuran.
- Validasi konfigurasi/rules sebelum aktivasi.
- Simpan versi, waktu, dan hasil validasi.
- Aktivasi memerlukan konfirmasi admin.
- Validasi gagal berarti konfigurasi aktif tidak berubah.
- Jangan mengambil rules dari URL sembarang tanpa allowlist dan validasi integritas.

### F-05 Mode IDS/IPS
- IDS pasif adalah default.
- IPS inline nonaktif sampai admin mengaktifkannya secara eksplisit.
- Tampilkan peringatan bahwa IPS inline dapat mengganggu konektivitas.
- Lakukan preflight: interface, konfigurasi, hak akses, dan kesiapan rollback.
- Jangan mengubah firewall/nftables diam-diam.
- Bypass/fail-open hanya boleh diklaim bila didukung dan telah diuji pada target.

### F-06 IP response
- Pemblokiran otomatis bukan default.
- Blokir manual hanya boleh tersedia setelah adapter firewall Linux diuji.
- Catat alasan, aktor, waktu mulai, masa berlaku, dan pencabutan.
- Validasi IP/CIDR dan lindungi alamat management host.
- Jangan mengklaim berhasil sebelum OS mengonfirmasi tindakan.

### F-07 Audit
Catat login berhasil/gagal, perubahan konfigurasi, import/aktivasi rules, perubahan mode, block/unblock, serta perubahan pengguna/peran. Record minimal: actor, action, target, timestamp, outcome, dan request ID. Jangan simpan password/token.

### F-08 Autentikasi dan otorisasi
- Login wajib untuk dashboard operasional.
- Role minimum: `admin` dan `analyst`.
- Analyst hanya membaca; admin melakukan mutasi yang diizinkan.
- Gunakan password hashing yang sesuai dan session/token dengan expiry serta revoke.
- Semua endpoint sensitif memeriksa otorisasi di server.
- Terapkan proteksi CSRF/CORS sesuai mekanisme autentikasi.

### F-09 Konfigurasi dan retensi
- Validasi konfigurasi database, lokasi EVE JSON, interface, dan retensi.
- Secret hanya dari environment/secret store.
- Retensi dapat dikonfigurasi.
- Penghapusan data lama tidak boleh menghapus audit log tanpa sengaja.

## 7. Kebutuhan Non-Fungsional

- Keamanan: validasi input, auth, RBAC, secret hygiene, dan dependency review.
- Reliabilitas: collector pulih setelah restart dan tahan terhadap event rusak.
- Kinerja: pagination, indeks, batch insert terukur, dan backpressure.
- Skalabilitas: sensor, collector, API, dan database dapat dipisah host di masa depan.
- Portabilitas: Ubuntu Server adalah target resmi awal.
- Usability: desktop dan layar ponsel.
- Maintainability: package kecil, interface jelas, konfigurasi terdokumentasi, dan test otomatis.

## 8. Model Data Minimum

- `alerts`: id, sensor_id, event_timestamp, signature_id, signature, severity, src_ip/src_port, dest_ip/dest_port, protocol, category, raw_event JSONB, created_at.
- `sensors`: id, name, hostname, engine_version, last_seen_at, status, created_at.
- `users`: id, email, password_hash, role, created_at, updated_at.
- `audit_logs`: id, actor_user_id, action, target, outcome, details JSONB, request_id, created_at.

Skema final ditentukan melalui migration yang ditinjau. Jangan menambah tabel tanpa kebutuhan jelas.

## 9. API Minimum yang Diusulkan

- `GET /api/v1/health`
- `GET /api/v1/sensors`
- `GET /api/v1/alerts`
- `GET /api/v1/alerts/{id}`
- `GET /api/v1/rules`
- `POST /api/v1/rules/import`
- `POST /api/v1/rules/{id}/validate`
- `POST /api/v1/rules/{id}/activate`
- `GET /api/v1/settings`
- `PATCH /api/v1/settings`
- `GET /api/v1/audit-logs`

Endpoint mutasi wajib auth, role check, validasi, dan audit. Block/unblock baru ditambahkan setelah adapter OS dan test disetujui.

## 10. Roadmap Bertahap dan Gerbang Kelulusan

### Tahap 0: Repository reconnaissance & planning
Baca struktur, README, konfigurasi, dependency, test, dan CI. Laporkan stack, risiko, asumsi, rencana file, serta perintah test. **Dilarang mengedit file.**
Lulus bila laporan berbasis bukti tersedia dan working tree tidak berubah.

### Tahap 1: Project foundation
Buat fondasi Go dan Next.js, contoh environment tanpa secret, health endpoint, dashboard placeholder yang jujur, logging terstruktur, dan konfigurasi tervalidasi.
Lulus bila backend/frontend berjalan, health benar, build/test lulus, dan tidak ada credential di repo.

### Tahap 2: Suricata lab & EVE JSON
Dokumentasikan instalasi Suricata Ubuntu, jalankan IDS pasif di lab, gunakan PCAP legal, dan verifikasi EVE JSON.
Lulus bila versi/status engine dapat diperiksa, PCAP menghasilkan event yang diharapkan, dan tidak ada inline blocking.

### Tahap 3: EVE collector
Buat parser streaming, validasi field, retry/backoff, penanganan malformed line, batch persistence, dan metrik.
Lulus bila unit test event valid/opsional/rusak serta integration test database lulus; satu baris rusak tidak menghentikan collector.

### Tahap 4: Alert API & dashboard
Buat endpoint list/detail, pagination/filter, tabel/detail alert, dan status sensor berbasis query aktual.
Lulus bila data UI cocok dengan database, empty/error state ada, dan tidak ada angka dummy.

### Tahap 5: Authentication, RBAC & audit
Implementasikan session lifecycle, role admin/analyst, perlindungan endpoint, audit log, dan rate limit login.
Lulus bila akses anonim ditolak, analyst tidak dapat mutasi, secret tidak muncul di log, dan audit tercatat.

### Tahap 6: Rule management
Implementasikan daftar/import ruleset, batas ukuran, validasi Suricata, versioning, aktivasi atomik, dan rollback.
Lulus bila rules invalid ditolak dan kegagalan aktivasi tidak merusak konfigurasi aktif.

### Tahap 7: IPS inline eksperimental
Hanya di lab terisolasi. Dokumentasikan topologi dua interface/NFQUEUE yang didukung, preflight, backup, konfirmasi eksplisit, rule drop terkontrol, rollback, dan service crash.
Lulus bila trafik uji yang cocok diblokir, trafik lain lewat, dan rollback terbukti. Jangan aktifkan di host produksi secara default.

### Tahap 8: Hardening & release candidate
Review threat model/permission, dependency, backup/restore, beban, dokumentasi install/upgrade/rollback, dan limitasi.
Lulus bila test suite lulus, build dapat diulang, secret scan bersih, smoke test tercatat, dan limitasi dinyatakan.

Agent tidak boleh melompat tahap sebelum acceptance criteria tahap aktif lulus.

## 11. Aturan Anti AI Slop (Wajib)

1. Baca file terkait dan konteks pemanggil sebelum mengubah.
2. Kerjakan satu tahap per pekerjaan; jangan melompat tanpa persetujuan.
3. Lakukan perubahan minimal; jangan refactor luas tanpa alasan.
4. Jangan membuat tombol, grafik, angka, status, atau API palsu.
5. Jangan mengklaim build, test, integrasi, IPS, atau deployment berhasil tanpa menjalankan dan melaporkan bukti.
6. Verifikasi dokumentasi resmi, versi dependency, dan API.
7. Baca schema/migration aktual sebelum query.
8. Tangani error; jangan mengganti kegagalan dengan data dummy.
9. Jangan menghapus, melemahkan, atau melewati test agar pipeline hijau.
10. Jelaskan dan minta persetujuan untuk perubahan auth, RBAC, firewall, nftables, capability, atau permission.
11. Semua uji IPS hanya di lab terisolasi.
12. Jangan commit .env, password, token, private key, atau PCAP sensitif.
13. Laporkan command, exit code, ringkasan output, dan test yang tidak dijalankan.
14. Jika konteks kurang, ajukan satu pertanyaan spesifik daripada menebak.
15. Jangan menambah library/service tanpa kebutuhan dan trade-off tertulis.

### Format laporan wajib

```text
TAHAP:
STATUS: LULUS / GAGAL / TERBLOKIR
FILE DIBACA:
- ...
FILE DIBUAT/DIUBAH:
- ...
IMPLEMENTASI:
- ...
PENGUJIAN:
- Perintah:
- Hasil:
- Exit code:
LIMITASI / RISIKO:
- ...
LANGKAH BERIKUTNYA:
- ...
```

## 12. Strategi Pengujian

- Unit: parser, validator, filter, authorization, business logic.
- Integration: collector-PostgreSQL, API-database, Suricata EVE ingestion.
- PCAP: dataset uji legal dengan hasil yang diketahui.
- Security: auth bypass, RBAC, input validation, injection, secret leakage.
- Failure: database down, malformed event, disk-full simulation, collector restart, engine stopped.
- IPS: hanya jaringan terisolasi dengan akses pemulihan console.

Sebutkan cakupan setiap test. Test yang tidak dijalankan harus ditandai `NOT RUN`.

## 13. Risiko dan Mitigasi

| Risiko | Mitigasi |
|---|---|
| False positive | IDS default, tuning rules, review alert |
| False negative | Dokumentasikan cakupan ruleset dan batas inspeksi |
| IPS memutus koneksi | Lab test, preflight, backup, rollback |
| Event volume tinggi | Batch, indeks, retensi, backpressure |
| Collector berhenti | systemd restart policy, health check, monitoring |
| Privilege escalation | Least privilege dan helper terbatas |
| Credential bocor | Environment/secret store dan secret scanning |
| Dashboard menampilkan data palsu | Semua metrik berasal dari API/database nyata |

## 14. Kriteria MVP Selesai

MVP selesai hanya bila Suricata IDS pasif berjalan, collector menyimpan EVE JSON, dashboard menampilkan event aktual dengan filter/pagination, auth/RBAC berfungsi, perubahan tercatat di audit, ruleset divalidasi sebelum aktivasi, dokumentasi operasi tersedia, dan test wajib memiliki bukti lulus.

IPS inline adalah milestone eksperimental terpisah, bukan syarat MVP dashboard IDS.

## 15. Referensi Teknis

- Suricata: https://docs.suricata.io/
- Suricata IPS: https://docs.suricata.io/en/latest/ips/ips-concept.html
- Suricata EVE JSON: https://docs.suricata.io/en/latest/output/eve/eve-json-output.html
- Snort: https://www.snort.org/
- Go: https://go.dev/doc/
- PostgreSQL: https://www.postgresql.org/docs/
- Next.js: https://nextjs.org/docs

Periksa kembali dokumentasi saat implementasi. Dokumentasi `latest` dapat berubah.
