type MetricProps = {
  label: string;
  value: string;
  note: string;
};

function Metric({ label, value, note }: MetricProps) {
  return (
    <section className="rounded-xl border border-[var(--line)] bg-white p-5">
      <p className="text-sm text-[var(--muted)]">{label}</p>
      <p className="mt-3 text-3xl font-semibold tracking-tight">{value}</p>
      <p className="mt-2 text-xs text-[var(--muted)]">{note}</p>
    </section>
  );
}

export default function Home() {
  return (
    <main className="mx-auto min-h-screen w-full max-w-6xl px-4 py-8 sm:px-6 lg:px-8">
      <header className="flex flex-wrap items-center justify-between gap-4 border-b border-[var(--line)] pb-6">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.2em] text-[var(--accent)]">
            Network security platform
          </p>
          <h1 className="mt-2 text-2xl font-semibold tracking-tight">Saragosa</h1>
        </div>
        <span className="rounded-full border border-emerald-200 bg-emerald-50 px-3 py-1.5 text-xs font-medium text-emerald-800">
          Foundation stage
        </span>
      </header>

      <section className="py-8">
        <h2 className="text-xl font-semibold">Ringkasan sistem</h2>
        <p className="mt-2 max-w-2xl text-sm leading-6 text-[var(--muted)]">
          Fondasi dashboard sedang disiapkan. Integrasi sensor dan data alert
          belum aktif, jadi halaman ini tidak menampilkan angka simulasi.
        </p>
      </section>

      <section className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <Metric label="Mode inspeksi" value="IDS pasif" note="Mode pencegahan belum diaktifkan." />
        <Metric label="Sensor Suricata" value="Belum terhubung" note="Status sensor belum diperiksa oleh backend." />
        <Metric label="Event keamanan" value="Belum tersedia" note="Menunggu integrasi EVE JSON dan penyimpanan." />
      </section>

      <section className="mt-8 rounded-xl border border-[var(--line)] bg-white p-5">
        <h2 className="font-semibold">Status fondasi</h2>
        <ul className="mt-4 space-y-3 text-sm">
          <li className="flex items-center justify-between gap-3">
            <span>Web dashboard</span><span className="text-emerald-700">Tersedia</span>
          </li>
          <li className="flex items-center justify-between gap-3">
            <span>API health endpoint</span><span className="text-amber-700">Perlu dijalankan</span>
          </li>
          <li className="flex items-center justify-between gap-3">
            <span>EVE JSON collector foundation</span><span className="text-amber-700">CLI awal</span>
          </li>
          <li className="flex items-center justify-between gap-3">
            <span>Database dan autentikasi</span><span className="text-[var(--muted)]">Belum dibuat</span>
          </li>
        </ul>
      </section>

      <footer className="mt-10 border-t border-[var(--line)] pt-5 text-xs text-[var(--muted)]">
        Saragosa is an early development project. No live network protection is active.
      </footer>
    </main>
  );
}
