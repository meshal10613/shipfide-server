'use client';

import { useState } from 'react';
import { Search, CheckCircle2 } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';
import { Badge } from '@/components/ui/Badge';

export const TrackingView = () => {
  const [code, setCode] = useState('');
  const [searched, setSearched] = useState(false);

  return (
    <div className="max-w-4xl mx-auto px-4 sm:px-6 py-12 w-full space-y-8">
      <div className="text-center space-y-3">
        <Badge variant="red">LIVE TELEMETRY</Badge>
        <h1 className="text-3xl sm:text-4xl font-black text-[#404042] tracking-tight">
          Track Your Shipfide Parcel
        </h1>
        <p className="text-sm text-slate-500 max-w-lg mx-auto">
          Sub-second real-time tracking across 64 districts in Bangladesh
        </p>
      </div>

      <div className="bg-white p-4 sm:p-6 rounded-3xl border border-slate-200 shadow-sm max-w-xl mx-auto">
        <form
          onSubmit={(e) => {
            e.preventDefault();
            if (code.trim()) setSearched(true);
          }}
          className="flex gap-2"
        >
          <div className="flex-1">
            <Input
              placeholder="Enter Tracking ID (e.g. SF-20260722-04829)"
              value={code}
              onChange={(e) => setCode(e.target.value)}
              leftIcon={<Search className="w-4 h-4 text-slate-400" />}
            />
          </div>
          <Button type="submit" variant="primary" className="h-10 rounded-full px-6">
            Track
          </Button>
        </form>
      </div>

      {searched && (
        <div className="bg-white rounded-3xl border border-slate-200 p-6 sm:p-8 shadow-md space-y-6">
          <div className="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4 pb-6 border-b border-slate-100">
            <div>
              <span className="text-xs text-slate-400 font-mono">TRACKING NUMBER</span>
              <h3 className="text-xl font-black font-mono text-[#DD0033]">{code || 'SF-20260722-04829'}</h3>
            </div>
            <Badge variant="emerald" size="md">OUT FOR DELIVERY</Badge>
          </div>

          <div className="space-y-6">
            <h4 className="text-xs font-bold text-[#404042] uppercase tracking-wider">Shipment History</h4>
            <div className="space-y-4">
              {[
                { title: 'Out For Delivery with Rider Karim Ahmed', time: 'Today, 09:30 AM', active: true },
                { title: 'Arrived at Dhaka Central Hub', time: 'Today, 06:15 AM', done: true },
                { title: 'Picked Up from Merchant Warehouse', time: 'Yesterday, 04:00 PM', done: true },
                { title: 'Parcel Order Booked by Merchant', time: 'Yesterday, 02:00 PM', done: true },
              ].map((step, i) => (
                <div key={i} className="flex items-start gap-4">
                  <div className={`w-6 h-6 rounded-full flex items-center justify-center shrink-0 mt-0.5 ${
                    step.active ? 'bg-[#DD0033] text-white animate-pulse' : 'bg-emerald-600 text-white'
                  }`}>
                    <CheckCircle2 className="w-4 h-4" />
                  </div>
                  <div>
                    <h5 className="text-xs font-bold text-[#404042]">{step.title}</h5>
                    <span className="text-[11px] font-mono text-slate-400">{step.time}</span>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
