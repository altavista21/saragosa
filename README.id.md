# Saragosa

**Saragosa** adalah platform pengelolaan Network Intrusion Detection and Prevention System (IDS/IPS) yang berfokus pada Linux dan menggunakan engine Suricata.

Tujuan proyek ini adalah menyediakan antarmuka web praktis untuk memantau kesehatan sensor, melihat event keamanan, menyelidiki alert, mengelola ruleset, dan menjalankan tindakan respons secara terkontrol.

> **Status proyek:** Tahap perencanaan / PRD. Fitur yang direncanakan tidak boleh dianggap sudah tersedia.

## Arah proyek

Saragosa tidak akan membuat ulang packet inspection engine. Saragosa menggunakan **Suricata** untuk inspeksi jaringan dan membangun lapisan aplikasi untuk pengumpulan event, visibilitas, pengelolaan konfigurasi, serta operasi yang tercatat dalam audit log.

### Fitur yang direncanakan

- Status kesehatan sensor dan service Suricata
- Ingesti event EVE JSON
- Pencarian, filter, dan investigasi alert
- Import dan validasi ruleset
- Role-based access control
- Audit log administrator
- Mode IDS pasif sebagai pengaturan default
- Alur IPS inline eksperimental dan opt-in

## Teknologi yang diusulkan

| Area | Teknologi |
|---|---|
| Detection engine | Suricata |
| Backend dan collector | Go |
| Dashboard web | Next.js, TypeScript |
| UI | Tailwind CSS, Lucide React |
| Database | PostgreSQL |
| Format event | Suricata EVE JSON |
| Target awal | Ubuntu Server Linux |
| Pengujian | Go tests, integration tests, PCAP fixtures |

Versi dependency akan dikunci saat implementasi. Tidak ada klaim performa atau tingkat deteksi sebelum benchmark yang dapat diulang tersedia.

## Roadmap pengembangan

Implementasi dibagi menjadi beberapa tahap dengan gerbang kelulusan:

1. Pemeriksaan repository dan perencanaan
2. Fondasi proyek
3. Lab Suricata dan verifikasi EVE JSON
4. Event collector
5. Alert API dan dashboard
6. Autentikasi, RBAC, dan audit
7. Pengelolaan ruleset
8. IPS inline eksperimental di lab terisolasi
9. Hardening dan release candidate

Lihat [docs/PRD.md](docs/PRD.md) untuk ruang lingkup, kebutuhan, kriteria penerimaan, strategi pengujian, dan aturan anti AI slop yang wajib dipatuhi AI Agent.

## Catatan keamanan

- IDS pasif adalah mode default.
- IPS inline dapat mengganggu konektivitas jaringan dan harus opt-in.
- Uji pemblokiran hanya di lab terisolasi dengan jalur pemulihan.
- Saragosa bukan pengganti NGFW, SIEM, EDR, atau NDR enterprise.
- Jangan membuka antarmuka manajemen ke jaringan yang tidak tepercaya.
- Jangan pernah melakukan commit terhadap kredensial, private key, PCAP sensitif, atau secret produksi.

## Batasan saat ini

Pada tahap PRD, integrasi sensor, dashboard aktif, tindakan pemblokiran, maupun deployment produksi belum boleh dianggap sudah berfungsi. Status fitur hanya boleh diperbarui setelah implementasi dan pengujian memberikan bukti.

## Dokumentasi

- PRD bahasa Indonesia: [docs/PRD.md](docs/PRD.md)
- README bahasa Inggris: [README.md](README.md)

## Lisensi

Lisensi belum ditentukan. Sebelum lisensi ditambahkan, jangan menganggap kode boleh digunakan ulang, dimodifikasi, atau didistribusikan.
