'use client';

import { useState } from 'react';
import { Navbar } from '@/components/layout/Navbar';
import { BottomPillNav } from '@/components/layout/BottomPillNav';
import { MerchantDashboard } from '@/components/dashboard/MerchantDashboard';
import { RiderDashboard } from '@/components/dashboard/RiderDashboard';
import { AdminDashboard } from '@/components/dashboard/AdminDashboard';
import { SuperAdminDashboard } from '@/components/dashboard/SuperAdminDashboard';

export function DashboardView() {
  const [role, setRole] = useState<string>('MERCHANT');

  return (
    <div className="min-h-screen bg-[#F9FCFE] text-[#404042] flex flex-col justify-between">
      <Navbar currentRole={role} onRoleChange={setRole} />

      <main className="max-w-4xl mx-auto px-4 sm:px-6 pt-6 w-full flex-1">
        {role === 'SUPER_ADMIN' && <SuperAdminDashboard />}
        {role === 'ADMIN' && <AdminDashboard />}
        {role === 'MERCHANT' && <MerchantDashboard />}
        {role === 'RIDER' && <RiderDashboard />}
      </main>

      {/* Manual Bottom Text Footer */}
      <footer className="py-6 text-center text-xs text-slate-500 border-t border-slate-200 mt-12 mb-20">
        <p className="font-bold text-[#404042]">Shipfide Logistics Platform</p>
        <p className="text-[11px] mt-0.5">Delivering Trust, Every Time • © {new Date().getFullYear()} Shipfide. All rights reserved.</p>
      </footer>

      <BottomPillNav />
    </div>
  );
}
