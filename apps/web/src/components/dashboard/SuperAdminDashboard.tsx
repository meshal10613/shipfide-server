'use client';

import { motion } from 'framer-motion';
import {
  Building,
  UserPlus,
  ChevronRight,
} from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/Button';

export function SuperAdminDashboard() {
  return (
    <div className="space-y-8 pb-12">
      {/* Super Admin Header */}
      <div className="flex items-center justify-between p-6 rounded-3xl bg-white border border-slate-200 shadow-sm">
        <div className="flex items-center gap-4">
          <div className="w-14 h-14 rounded-full bg-gradient-to-tr from-purple-600 to-[#DD0033] p-0.5 shadow-md">
            <div className="w-full h-full rounded-full bg-white flex items-center justify-center font-bold text-lg text-purple-700">
              SA
            </div>
          </div>
          <div>
            <div className="flex items-center gap-2">
              <h2 className="text-xl font-extrabold text-[#404042] tracking-tight">Shipfide Super Admin</h2>
              <span className="px-2.5 py-0.5 rounded-full text-[10px] font-bold font-mono bg-purple-50 text-purple-700 border border-purple-200">
                SUPER_ADMIN
              </span>
            </div>
            <p className="text-xs text-slate-500 font-medium">System Control & Global Operations</p>
          </div>
        </div>
      </div>

      {/* Global Revenue Card */}
      <motion.div
        initial={{ opacity: 0, y: 15 }}
        animate={{ opacity: 1, y: 0 }}
        className="p-6 rounded-3xl bg-linear-to-r from-purple-50/80 via-white to-rose-50/50 border border-purple-200 shadow-sm space-y-4"
      >
        <div className="flex items-center justify-between">
          <div>
            <span className="text-xs font-mono text-purple-700 uppercase tracking-wider block">
              Total Platform Revenue
            </span>
            <span className="text-3xl font-black text-[#404042] tracking-tight">৳ 1,482,900.00</span>
          </div>
          <Link href="/super-admin-dashboard/hubs">
            <Button variant="primary" size="sm" rightIcon={<ChevronRight className="w-4 h-4" />}>
              Manage Hubs
            </Button>
          </Link>
        </div>

        <div className="grid grid-cols-3 gap-2 pt-3 border-t border-slate-200 text-xs">
          <div>
            <span className="text-slate-500 block">System Share</span>
            <strong className="text-emerald-600 font-bold">৳ 444,870.00</strong>
          </div>
          <div>
            <span className="text-slate-500 block">Rider Share</span>
            <strong className="text-amber-600 font-bold">৳ 593,160.00</strong>
          </div>
          <div>
            <span className="text-slate-500 block">Merchant Net</span>
            <strong className="text-purple-700 font-bold">৳ 444,870.00</strong>
          </div>
        </div>
      </motion.div>

      {/* Super Admin Quick Actions */}
      <div className="grid grid-cols-2 gap-4">
        <Link href="/super-admin-dashboard/hubs">
          <div className="p-5 rounded-3xl bg-white border border-slate-200 hover:border-[#DD0033]/30 text-left space-y-2 group transition-all shadow-sm hover:shadow-md cursor-pointer">
            <div className="w-10 h-10 rounded-2xl bg-rose-50 border border-rose-100 flex items-center justify-center text-[#DD0033]">
              <Building className="w-5 h-5" />
            </div>
            <h4 className="text-sm font-bold text-[#404042]">Create Operational Hub</h4>
            <p className="text-[11px] text-slate-500">Add physical hub location</p>
          </div>
        </Link>

        <Link href="/super-admin-dashboard/staff">
          <div className="p-5 rounded-3xl bg-white border border-slate-200 hover:border-[#DD0033]/30 text-left space-y-2 group transition-all shadow-sm hover:shadow-md cursor-pointer">
            <div className="w-10 h-10 rounded-2xl bg-rose-50 border border-rose-100 flex items-center justify-center text-[#DD0033]">
              <UserPlus className="w-5 h-5" />
            </div>
            <h4 className="text-sm font-bold text-[#404042]">Invite Admin Staff</h4>
            <p className="text-[11px] text-slate-500">Generate staff credentials</p>
          </div>
        </Link>
      </div>
    </div>
  );
}
