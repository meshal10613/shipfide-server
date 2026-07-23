'use client';

import { useState } from 'react';
import { Search, CheckCircle2 } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';
import { Badge } from '@/components/ui/Badge';

export const ReceiverFraudCheckView = () => {
  const [phone, setPhone] = useState('');
  const [checked, setChecked] = useState(false);

  return (
    <div className="max-w-3xl mx-auto px-4 sm:px-6 py-12 w-full space-y-6">
      <div className="text-center space-y-3">
        <Badge variant="purple">FRAUD PROTECTION</Badge>
        <h1 className="text-3xl font-black text-[#404042] tracking-tight">
          Receiver Fraud & Risk Check
        </h1>
        <p className="text-xs text-slate-500 max-w-md mx-auto">
          Check receiver phone number cancellation history before booking parcel dispatch
        </p>
      </div>

      <form
        onSubmit={(e) => {
          e.preventDefault();
          if (phone.trim()) setChecked(true);
        }}
        className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4"
      >
        <Input
          label="Receiver Phone Number"
          required
          placeholder="+8801700000000"
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
          leftIcon={<Search className="w-4 h-4 text-slate-400" />}
        />
        <Button type="submit" variant="primary" className="w-full h-10 rounded-full">
          Analyze Receiver Risk Profile
        </Button>
      </form>

      {checked && (
        <div className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
          <div className="flex items-center justify-between">
            <div>
              <h3 className="text-sm font-bold font-mono text-[#404042]">{phone}</h3>
              <p className="text-xs text-slate-500">Total Orders: 24 • Delivered: 23 • Cancelled: 1</p>
            </div>
            <Badge variant="emerald" size="md">SAFE RECEIVER (95.8%)</Badge>
          </div>

          <div className="p-4 rounded-2xl bg-emerald-50 border border-emerald-200 text-xs text-emerald-800 flex items-center gap-2">
            <CheckCircle2 className="w-5 h-5 text-emerald-600 shrink-0" />
            <span>Low return risk. Safe for Cash on Delivery (COD) dispatch.</span>
          </div>
        </div>
      )}
    </div>
  );
};
