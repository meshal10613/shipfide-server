'use client';

import { useState } from 'react';
import { DollarSign, CheckCircle2 } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export const WalkInPaymentView = () => {
  const [collected, setCollected] = useState(false);

  return (
    <div className="max-w-md mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Walk-In Customer Payment</h1>
        <p className="text-xs text-slate-500">Collect OTC cash / digital payment at hub desk</p>
      </div>

      {collected ? (
        <div className="bg-white p-8 rounded-3xl border border-slate-200 shadow-sm text-center space-y-4">
          <CheckCircle2 className="w-12 h-12 text-emerald-600 mx-auto" />
          <h2 className="text-xl font-bold text-[#404042]">Walk-In Payment Recorded!</h2>
          <p className="text-xs text-slate-500">Receipt printed & status set to COLLECTED.</p>
          <Button onClick={() => setCollected(false)} variant="primary" className="h-10 rounded-full px-6">Process Another Payment</Button>
        </div>
      ) : (
        <form onSubmit={(e) => { e.preventDefault(); setCollected(true); }} className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
          <Input label="Parcel Tracking ID" required placeholder="SF-20260722-XXXXX" />
          <Input label="Payment Amount (৳)" type="number" required defaultValue="120" leftIcon={<DollarSign className="w-4 h-4" />} />
          <Button type="submit" variant="primary" className="w-full h-10 rounded-full">
            Confirm Payment Collection
          </Button>
        </form>
      )}
    </div>
  );
};
