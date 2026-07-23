'use client';

import { Navbar } from '@/components/layout/Navbar';
import { BottomPillNav } from '@/components/layout/BottomPillNav';
import { useRouter, usePathname } from 'next/navigation';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const pathname = usePathname();

  // Determine current active role based on path
  let activeRole = 'MERCHANT';
  if (pathname.startsWith('/admin-dashboard')) activeRole = 'ADMIN';
  if (pathname.startsWith('/super-admin-dashboard')) activeRole = 'SUPER_ADMIN';
  if (pathname.startsWith('/rider-dashboard')) activeRole = 'RIDER';
  if (pathname.startsWith('/merchant-dashboard')) activeRole = 'MERCHANT';

  const handleRoleChange = (newRole: string) => {
    switch (newRole) {
      case 'SUPER_ADMIN':
        router.push('/super-admin-dashboard');
        break;
      case 'ADMIN':
        router.push('/admin-dashboard');
        break;
      case 'MERCHANT':
        router.push('/merchant-dashboard');
        break;
      case 'RIDER':
        router.push('/rider-dashboard');
        break;
      default:
        router.push('/dashboard');
    }
  };

  return (
    <div className="min-h-screen bg-[#F9FCFE] text-[#404042] flex flex-col justify-between">
      <Navbar currentRole={activeRole} onRoleChange={handleRoleChange} />

      <main className="max-w-4xl mx-auto px-4 sm:px-6 pt-6 w-full flex-1">
        {children}
      </main>

      {/* Manual Bottom Text Footer */}
      <footer className="py-6 text-center text-xs text-slate-500 border-t border-slate-200 mt-12 mb-20">
        <p className="font-bold text-[#404042]">Shipfide Logistics Platform</p>
        <p className="text-[11px] mt-0.5">Delivering Trust, Every Time • © 2026 Shipfide. All rights reserved.</p>
      </footer>

      <BottomPillNav />
    </div>
  );
}
