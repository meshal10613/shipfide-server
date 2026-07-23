'use client';

import { Badge } from '@/components/ui/Badge';

export const HubOverviewView = () => {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Dhaka Central Hub Operations</h1>
          <p className="text-xs text-slate-500">Live parcel throughput, rider shift, & cash status</p>
        </div>
        <Badge variant="blue" size="md">HUB STAFF</Badge>
      </div>

      <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Parcels Today</span>
          <div className="text-2xl font-black text-[#DD0033]">1,420</div>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Active Riders</span>
          <div className="text-2xl font-black text-blue-600">48</div>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Pending COD</span>
          <div className="text-2xl font-black text-amber-600">৳ 84.2K</div>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Fraud Flagged</span>
          <div className="text-2xl font-black text-red-600">3</div>
        </div>
      </div>
    </div>
  );
};
