'use client';

import { Send } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export const DispatchParcelsView = () => {
  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Dispatch Parcels to Rider / Hub</h1>
        <p className="text-xs text-slate-500">Assign sorted hub parcels to delivery riders or linehaul transit</p>
      </div>

      <form className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
        <Input label="Parcel Tracking Number" placeholder="SF-20260722-XXXXX" />
        <div className="space-y-1.5">
          <label className="text-xs font-semibold text-[#404042]">Assign to Rider</label>
          <select className="w-full py-3 px-4 rounded-2xl bg-white border border-slate-200 text-xs text-[#404042] focus:outline-none focus:border-[#DD0033]">
            <option value="1">Karim Ahmed (Banani Sector)</option>
            <option value="2">Salam Hossain (Uttara Sector)</option>
          </select>
        </div>
        <Button type="submit" variant="primary" className="w-full h-10 rounded-full" leftIcon={<Send className="w-4 h-4" />}>
          Dispatch Parcel Out For Delivery
        </Button>
      </form>
    </div>
  );
};
