import Link from 'next/link';

export default function NotFound() {
  return (
    <div className="min-h-screen bg-[#F9FCFE] text-[#404042] flex flex-col items-center justify-center p-4 text-center space-y-4">
      <h1 className="text-6xl font-black text-[#DD0033]">404</h1>
      <h2 className="text-xl font-bold text-[#404042]">Page Not Found</h2>
      <p className="text-xs text-slate-500 max-w-sm">
        The page you are looking for does not exist or has been moved.
      </p>
      <Link
        href="/"
        className="px-6 py-3 rounded-2xl bg-[#DD0033] hover:bg-[#B30028] text-white font-bold text-xs shadow-md transition-colors"
      >
        Return to Home Page
      </Link>
    </div>
  );
}
