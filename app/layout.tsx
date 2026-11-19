import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Saragosa | Network Security Monitoring",
  description: "Saragosa network intrusion detection management dashboard.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="id">
      <body>{children}</body>
    </html>
  );
}
