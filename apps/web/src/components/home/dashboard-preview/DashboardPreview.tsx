'use client';

import { motion } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { BarChart3, TrendingUp, DollarSign, Package, Users, ShieldCheck } from 'lucide-react';

export const DashboardPreview = () => {
  return (
    <section className="py-20 bg-[#F9FCFE] border-b border-slate-200/60" id="dashboard-preview">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-16">
        <SectionHeading
          badge="MERCHANT DASHBOARD"
          title="Command Center for Your Entire Shipping Fleet"
          description="Track active orders, monitor live COD balances, generate bulk waybills, and manage staff permissions from one single platform."
        />

        {/* Dashboard Screen Mockup Container */}
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.6 }}
          className="bg-white rounded-3xl border border-slate-200 shadow-2xl p-6 sm:p-8 space-y-6 max-w-5xl mx-auto overflow-hidden"
        >
          {/* Top Mock Header */}
          <div className="flex items-center justify-between pb-4 border-b border-slate-100">
            <div className="flex items-center gap-3">
              <div className="w-3 h-3 rounded-full bg-rose-500" />
              <div className="w-3 h-3 rounded-full bg-amber-500" />
              <div className="w-3 h-3 rounded-full bg-emerald-500" />
              <span className="text-xs font-mono font-bold text-slate-400 ml-2">shipfide.com/merchant-dashboard</span>
            </div>
            <span className="px-3 py-1 rounded-full bg-[#DD0033]/10 text-[#DD0033] text-xs font-bold">
              LIVE METRICS
            </span>
          </div>

          {/* Metric Stat Cards */}
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
            <div className="p-4 rounded-2xl bg-[#F9FCFE] border border-slate-200/80 space-y-1">
              <span className="text-[11px] font-semibold text-slate-400">Total COD Balance</span>
              <span className="text-xl sm:text-2xl font-black text-[#DD0033] block">৳ 458,200</span>
              <span className="text-[10px] text-emerald-600 font-bold flex items-center gap-1">
                <TrendingUp className="w-3 h-3" /> +14.2% this week
              </span>
            </div>

            <div className="p-4 rounded-2xl bg-[#F9FCFE] border border-slate-200/80 space-y-1">
              <span className="text-[11px] font-semibold text-slate-400">Parcels Shipped</span>
              <span className="text-xl sm:text-2xl font-black text-[#404042] block">14,290</span>
              <span className="text-[10px] text-slate-400">99.4% Delivered</span>
            </div>

            <div className="p-4 rounded-2xl bg-[#F9FCFE] border border-slate-200/80 space-y-1">
              <span className="text-[11px] font-semibold text-slate-400">Active Riders</span>
              <span className="text-xl sm:text-2xl font-black text-blue-600 block">48</span>
              <span className="text-[10px] text-slate-400">Dhaka Hub Shift</span>
            </div>

            <div className="p-4 rounded-2xl bg-[#F9FCFE] border border-slate-200/80 space-y-1">
              <span className="text-[11px] font-semibold text-slate-400">Return Rate</span>
              <span className="text-xl sm:text-2xl font-black text-emerald-600 block">1.8%</span>
              <span className="text-[10px] text-emerald-600 font-bold">Fraud Shield Active</span>
            </div>
          </div>

          {/* Table Preview */}
          <div className="rounded-2xl border border-slate-200 overflow-hidden text-xs">
            <div className="bg-slate-50 p-3 font-bold text-[#404042] border-b border-slate-200 grid grid-cols-4 sm:grid-cols-5">
              <span>Tracking ID</span>
              <span>Receiver</span>
              <span className="hidden sm:block">Destination</span>
              <span>COD Amount</span>
              <span className="text-right">Status</span>
            </div>
            {[
              { id: 'SF-20260722-04829', receiver: 'Naimur Rahman', area: 'Banani, Dhaka', cod: '৳ 2,450.00', status: 'OUT_FOR_DELIVERY' },
              { id: 'SF-20260722-03912', receiver: 'Fatima Zohra', area: 'Agrabad, Chattogram', cod: '৳ 1,800.00', status: 'IN_TRANSIT' },
              { id: 'SF-20260721-09410', receiver: 'Arif Hossain', area: 'Dhanmondi, Dhaka', cod: '৳ 3,200.00', status: 'DELIVERED' },
            ].map((row, i) => (
              <div key={i} className="p-3 border-b border-slate-100 last:border-0 grid grid-cols-4 sm:grid-cols-5 items-center hover:bg-slate-50/50">
                <span className="font-mono font-bold text-[#404042]">{row.id}</span>
                <span className="text-slate-600 font-medium">{row.receiver}</span>
                <span className="text-slate-500 hidden sm:block">{row.area}</span>
                <span className="font-bold text-[#DD0033]">{row.cod}</span>
                <span className="text-right">
                  <span className={`px-2 py-0.5 rounded-full text-[10px] font-bold ${
                    row.status === 'DELIVERED' ? 'bg-emerald-50 text-emerald-600' : 'bg-red-50 text-[#DD0033]'
                  }`}>
                    {row.status}
                  </span>
                </span>
              </div>
            ))}
          </div>
        </motion.div>
      </div>
    </section>
  );
};
