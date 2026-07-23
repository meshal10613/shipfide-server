'use client';

import { motion } from 'framer-motion';
import { Button } from '@/components/ui/Button';

export function AdminDashboard() {
  return (
    <div className="space-y-8 pb-12">
      {/* Admin Header */}
      <div className="flex items-center justify-between p-6 rounded-3xl bg-white border border-slate-200 shadow-sm">
        <div className="flex items-center gap-4">
          <div className="w-14 h-14 rounded-full bg-gradient-to-tr from-blue-600 to-[#DD0033] p-0.5 shadow-md">
            <div className="w-full h-full rounded-full bg-white flex items-center justify-center font-bold text-lg text-blue-600">
              DH
            </div>
          </div>
          <div>
            <div className="flex items-center gap-2">
              <h2 className="text-xl font-extrabold text-[#404042] tracking-tight">Dhaka Hub Staff</h2>
              <span className="px-2.5 py-0.5 rounded-full text-[10px] font-bold font-mono bg-blue-50 text-blue-700 border border-blue-200">
                ROLE_ADMIN
              </span>
            </div>
            <p className="text-xs text-slate-500 font-medium">Dhaka Central Hub Operations</p>
          </div>
        </div>
      </div>

      {/* Hub Metrics Grid */}
      <div className="grid grid-cols-2 gap-4">
        {[
          { title: 'Hub Parcels Today', val: '1,420', sub: '85% Dispatched', color: 'text-[#DD0033]' },
          { title: 'Pending COD Deposits', val: '৳ 84,200', sub: '12 Rider Deposits', color: 'text-amber-600' },
          { title: 'Assigned Riders', val: '48 Active', sub: '95% On Shift', color: 'text-blue-600' },
          { title: 'Fraud Alerts', val: '3 Flagged', sub: 'COD Block Enforced', color: 'text-red-600' },
        ].map((m, i) => (
          <motion.div
            key={m.title}
            initial={{ opacity: 0, y: 15 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.05 * i }}
            className="p-5 rounded-3xl bg-white border border-slate-200 space-y-2 shadow-sm"
          >
            <span className="text-xs font-semibold text-slate-500 block">{m.title}</span>
            <div className={`text-2xl font-black ${m.color}`}>{m.val}</div>
            <span className="text-[11px] text-slate-400 block">{m.sub}</span>
          </motion.div>
        ))}
      </div>

      {/* Rider Cash Deposit Approvals List */}
      <div className="p-6 rounded-3xl bg-white border border-slate-200 space-y-4 shadow-sm">
        <div className="flex items-center justify-between">
          <h3 className="text-lg font-bold text-[#404042] tracking-tight">Rider COD Cash Deposit Approvals</h3>
          <span className="text-xs text-[#DD0033] font-bold">12 Pending</span>
        </div>

        <div className="space-y-3">
          {[
            { rider: 'Karim Ahmed', id: 'RD-8492', amount: '৳ 12,450.00', time: '10 mins ago' },
            { rider: 'Salam Hossain', id: 'RD-9120', amount: '৳ 8,900.00', time: '25 mins ago' },
            { rider: 'Tanvir Rahman', id: 'RD-3041', amount: '৳ 15,300.00', time: '40 mins ago' },
          ].map((deposit) => (
            <div
              key={deposit.id}
              className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-200 transition-colors"
            >
              <div>
                <h4 className="text-sm font-bold text-[#404042]">{deposit.rider}</h4>
                <span className="text-xs font-mono text-slate-500">{deposit.id} • {deposit.time}</span>
              </div>

              <div className="flex items-center gap-3">
                <span className="text-sm font-bold text-emerald-600">{deposit.amount}</span>
                <Button variant="primary" size="sm">
                  Approve
                </Button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
