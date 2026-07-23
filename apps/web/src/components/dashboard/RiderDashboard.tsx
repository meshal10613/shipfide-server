'use client';

import { motion } from 'framer-motion';
import {
  CheckCircle2,
  QrCode,
  Star,
  ChevronRight,
  MapPin,
  PhoneCall,
} from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/Button';

export function RiderDashboard() {
  return (
    <div className="space-y-8 pb-12">
      {/* Rider Header Bar */}
      <div className="flex items-center justify-between p-6 rounded-3xl bg-white border border-slate-200 shadow-sm">
        <div className="flex items-center gap-4">
          <div className="relative">
            <div className="w-14 h-14 rounded-full bg-gradient-to-tr from-amber-500 to-[#DD0033] p-0.5 shadow-md">
              <div className="w-full h-full rounded-full bg-white flex items-center justify-center font-bold text-lg text-amber-600">
                KA
              </div>
            </div>
            <span className="absolute bottom-0 right-0 w-4 h-4 rounded-full bg-emerald-500 border-2 border-white"></span>
          </div>
          <div>
            <div className="flex items-center gap-2">
              <h2 className="text-xl font-extrabold text-[#404042] tracking-tight">Karim Ahmed</h2>
              <span className="px-2.5 py-0.5 rounded-full text-[10px] font-bold font-mono bg-[#DD0033]/10 text-[#DD0033] border border-[#DD0033]/20">
                TOP RIDER
              </span>
            </div>
            <p className="text-xs text-slate-500 font-medium">Dhaka Central Hub • Active On Shift</p>
          </div>
        </div>

        <div className="flex items-center gap-1.5 px-3 py-1.5 rounded-2xl bg-amber-50 border border-amber-200 text-xs font-bold text-amber-700">
          <Star className="w-4 h-4 fill-amber-500 text-amber-500" />
          <span>4.95</span>
        </div>
      </div>

      {/* Cash Collection & Deposit Banner */}
      <motion.div
        initial={{ opacity: 0, y: 15 }}
        animate={{ opacity: 1, y: 0 }}
        className="p-6 rounded-3xl bg-linear-to-r from-emerald-50/80 via-white to-rose-50/50 border border-emerald-200 shadow-sm space-y-4"
      >
        <div className="flex items-center justify-between">
          <div>
            <span className="text-xs font-mono text-emerald-700 uppercase tracking-wider block">
              COD Cash in Hand
            </span>
            <span className="text-3xl font-black text-[#404042] tracking-tight">৳ 12,450.00</span>
          </div>
          <Link href="/rider-dashboard/cod-deposit">
            <Button variant="primary" size="sm" rightIcon={<ChevronRight className="w-4 h-4" />}>
              Deposit at Hub
            </Button>
          </Link>
        </div>

        <div className="pt-3 border-t border-slate-200 flex items-center justify-between text-xs text-slate-500">
          <span>Today's Deliveries: <strong className="text-[#404042]">18 Completed</strong></span>
          <span>Pending: <strong className="text-[#DD0033]">4 Remaining</strong></span>
        </div>
      </motion.div>

      {/* Active Parcel Assignment */}
      <div className="space-y-4">
        <div className="flex items-center justify-between">
          <h3 className="text-lg font-bold text-[#404042] tracking-tight">Active Delivery Task</h3>
          <span className="text-xs font-mono text-[#DD0033]">SF-20260722-04829</span>
        </div>

        <div className="p-6 rounded-3xl bg-white border border-slate-200 space-y-4 relative shadow-sm">
          <div className="flex items-center justify-between">
            <span className="px-3 py-1 rounded-full text-xs font-bold bg-amber-50 text-amber-700 border border-amber-200">
              OUT FOR DELIVERY
            </span>
            <span className="text-xs font-semibold text-slate-500">COD Amount: <strong className="text-[#404042]">৳ 1,850</strong></span>
          </div>

          <div className="space-y-3 pt-2">
            <div className="flex items-start gap-3">
              <MapPin className="w-5 h-5 text-[#DD0033] shrink-0 mt-0.5" />
              <div>
                <span className="text-xs font-semibold text-slate-500 block">Receiver Address</span>
                <p className="text-sm font-bold text-[#404042]">House 42, Road 11, Block D, Banani, Dhaka</p>
              </div>
            </div>

            <div className="flex items-center justify-between pt-2 border-t border-slate-100">
              <div>
                <span className="text-xs font-semibold text-slate-500 block">Receiver Phone</span>
                <span className="text-sm font-mono text-[#404042]">+8801711223344</span>
              </div>
              <Button variant="outline" size="sm" leftIcon={<PhoneCall className="w-3.5 h-3.5 text-[#DD0033]" />}>
                Call
              </Button>
            </div>
          </div>

          {/* OTP / Handover Actions */}
          <div className="grid grid-cols-2 gap-3 pt-3">
            <Link href="/rider-dashboard/deliveries/confirm-otp">
              <Button variant="primary" leftIcon={<CheckCircle2 className="w-4 h-4" />} className="w-full">
                Confirm Delivery
              </Button>
            </Link>
            <Button variant="outline" leftIcon={<QrCode className="w-4 h-4 text-[#DD0033]" />} className="w-full">
              Scan Parcel QR
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
