import Image from 'next/image';
import Link from 'next/link';

export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="min-h-screen bg-[#F9FCFE] text-[#404042] flex flex-col justify-between selection:bg-[#DD0033] selection:text-white">
      {/* Header */}
      <header className="py-6 px-4 border-b border-slate-200/60 bg-white">
        <div className="max-w-7xl mx-auto flex items-center justify-between">
          <Link href="/" className="flex items-center gap-2">
            <Image
              src="/logo.png"
              alt="Shipfide Logo"
              width={140}
              height={40}
              className="h-9 w-auto object-contain"
            />
          </Link>

          <Link
            href="/"
            className="text-xs font-bold text-slate-500 hover:text-[#DD0033] transition-colors"
          >
            ← Back to Home
          </Link>
        </div>
      </header>

      {/* Centered Auth Card Container */}
      <main className="max-w-md w-full mx-auto px-4 py-12 flex-1 flex flex-col justify-center">
        {children}
      </main>

      {/* Manual Footer */}
      <footer className="py-6 text-center text-xs text-slate-500 border-t border-slate-200 bg-white">
        <p className="font-bold text-[#404042]">Shipfide Logistics Platform</p>
        <p className="text-[11px] mt-0.5">Delivering Trust, Every Time • © 2026 Shipfide. All rights reserved.</p>
      </footer>
    </div>
  );
}
