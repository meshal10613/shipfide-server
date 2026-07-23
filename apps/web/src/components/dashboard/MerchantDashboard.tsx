'use client';

import { motion } from 'framer-motion';
import {
  Wallet,
  ArrowUpRight,
  Truck,
  Zap,
  PackageCheck,
  RotateCcw,
  Plus,
  Search,
  ShieldCheck,
  ChevronRight,
} from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/Button';

export function MerchantDashboard() {
  return (
    <div className="space-y-8 pb-12">
      {/* Merchant Header Profile Bar */}
      <div className="flex items-center justify-between p-6 rounded-3xl bg-white border border-slate-200 shadow-sm">
        <div className="flex items-center gap-4">
          <div className="relative">
            <div className="w-14 h-14 rounded-full bg-gradient-to-tr from-[#DD0033] to-rose-400 p-0.5 shadow-md">
              <div className="w-full h-full rounded-full bg-white flex items-center justify-center font-bold text-lg text-[#DD0033]">
                SJ
              </div>
            </div>
            <span className="absolute bottom-0 right-0 w-4 h-4 rounded-full bg-emerald-500 border-2 border-white"></span>
          </div>
          <div>
            <h2 className="text-xl font-extrabold text-[#404042] tracking-tight">Sri Julaekha</h2>
            <p className="text-xs text-slate-500 font-medium">Good Morning • Merchant Account</p>
          </div>
        </div>

        <div className="flex items-center gap-2">
          <Button variant="ghost" size="sm" className="w-10 h-10 p-0 rounded-2xl bg-slate-50 border border-slate-200 text-[#404042] hover:text-[#DD0033]">
            <Search className="w-4.5 h-4.5" />
          </Button>
          <Button variant="ghost" size="sm" className="w-10 h-10 p-0 rounded-2xl bg-slate-50 border border-slate-200 text-[#DD0033]">
            <ShieldCheck className="w-4.5 h-4.5" />
          </Button>
        </div>
      </div>

      {/* Balance Card */}
      <motion.div
        initial={{ opacity: 0, y: 15 }}
        animate={{ opacity: 1, y: 0 }}
        className="p-6 rounded-3xl bg-linear-to-r from-white via-slate-50 to-rose-50/50 border border-slate-200 shadow-md relative overflow-hidden group"
      >
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="w-12 h-12 rounded-2xl bg-[#DD0033]/10 border border-[#DD0033]/20 flex items-center justify-center text-[#DD0033]">
              <Wallet className="w-6 h-6" />
            </div>
            <div>
              <span className="text-xs font-mono text-slate-500 uppercase tracking-wider block">Available Balance</span>
              <span className="text-2xl font-black text-[#404042] tracking-tight">৳ 45,820.00</span>
            </div>
          </div>

          <Link href="/merchant-dashboard/withdrawals">
            <Button variant="primary" rightIcon={<ChevronRight className="w-4 h-4" />}>
              Withdraw Cash
            </Button>
          </Link>
        </div>
      </motion.div>

      {/* Services Section */}
      <div className="space-y-4">
        <div className="flex items-center justify-between">
          <h3 className="text-lg font-bold text-[#404042] tracking-tight">Services</h3>
          <Link href="/merchant-dashboard/parcels">
            <span className="text-xs font-semibold text-[#DD0033] hover:underline cursor-pointer">See all</span>
          </Link>
        </div>

        <div className="grid grid-cols-2 gap-4">
          {[
            {
              title: 'Pickup Request',
              sub: 'Schedule hub rider pickup',
              icon: Truck,
              href: '/merchant-dashboard/pickups',
            },
            {
              title: 'Express Delivery',
              sub: 'Same-day urgent delivery',
              icon: Zap,
              href: '/merchant-dashboard/parcels/create',
            },
            {
              title: 'Pick & Drop',
              sub: 'Point-to-point courier',
              icon: PackageCheck,
              href: '/merchant-dashboard/parcels/create',
            },
            {
              title: 'Return Request',
              sub: 'Manage parcel returns',
              icon: RotateCcw,
              href: '/merchant-dashboard/returns',
            },
          ].map((service, i) => (
            <Link key={service.title} href={service.href}>
              <motion.div
                initial={{ opacity: 0, y: 15 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.05 * i }}
                className="p-5 rounded-3xl bg-white border border-slate-200 hover:border-[#DD0033]/40 hover:shadow-md transition-all duration-200 relative group cursor-pointer"
              >
                <div className="flex items-start justify-between mb-4">
                  <div className="w-10 h-10 rounded-2xl bg-rose-50 border border-rose-100 flex items-center justify-center text-[#DD0033]">
                    <service.icon className="w-5 h-5" />
                  </div>
                  <div className="w-8 h-8 rounded-full bg-slate-100 flex items-center justify-center text-slate-400 group-hover:bg-[#DD0033] group-hover:text-white transition-colors">
                    <ArrowUpRight className="w-4 h-4" />
                  </div>
                </div>

                <h4 className="text-sm font-bold text-[#404042] mb-0.5">{service.title}</h4>
                <p className="text-[11px] text-slate-500 font-medium">{service.sub}</p>
              </motion.div>
            </Link>
          ))}
        </div>
      </div>

      {/* Quick Action Category Pills */}
      <div className="p-4 rounded-3xl bg-white border border-slate-200 shadow-sm space-y-3">
        <div className="grid grid-cols-4 gap-2">
          {[
            { label: 'Parcels', icon: '📦', href: '/merchant-dashboard/parcels' },
            { label: 'Payments', icon: '💳', href: '/merchant-dashboard/wallet' },
            { label: 'Support', icon: '🎧', href: '/support' },
            { label: 'Tickets', icon: '🎫', href: '/support' },
          ].map((cat) => (
            <Link key={cat.label} href={cat.href}>
              <div className="flex flex-col items-center gap-1.5 p-3 rounded-2xl bg-slate-50 hover:bg-rose-50/50 transition-colors text-center border border-slate-100 cursor-pointer">
                <span className="text-lg">{cat.icon}</span>
                <span className="text-[11px] font-semibold text-[#404042]">{cat.label}</span>
              </div>
            </Link>
          ))}
        </div>

        <Link href="/merchant-dashboard/parcels/create" className="block w-full">
          <Button variant="primary" leftIcon={<Plus className="w-4 h-4" />} className="w-full">
            Create New Delivery Order
          </Button>
        </Link>
      </div>
    </div>
  );
}
