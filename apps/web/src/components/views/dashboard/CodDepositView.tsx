'use client';

import { useState } from 'react';
import { DollarSign, Building2, CheckCircle2 } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export const CodDepositView = () => {
  const [deposited, setDeposited] = useState(false);

  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Hub COD Cash Deposit</h1>
        <p className="text-xs text-slate-500">Deposit collected cash-in-hand to Dhaka Central Hub manager</p>
      </div>

      {deposited ? (
        <div className="bg-white p-8 rounded-3xl border border-slate-200 shadow-sm text-center space-y-4">
          <CheckCircle2 className="w-12 h-12 text-emerald-600 mx-auto" />
          <h2 className="text-xl font-bold text-[#404042]">Deposit Request Submitted!</h2>
          <p className="text-xs text-slate-500">Amount <strong>৳ 12,450.00</strong> submitted for Hub Manager approval.</p>
        </div>
      ) : (
        <form onSubmit={(e) => { e.preventDefault(); setDeposited(true); }} className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
          <Input label="COD Cash in Hand (৳)" defaultValue="12450.00" disabled leftIcon={<DollarSign className="w-4 h-4" />} />
          <Input label="Target Operational Hub" defaultValue="Dhaka Central Hub" disabled leftIcon={<Building2 className="w-4 h-4" />} />
          <Button type="submit" variant="primary" className="w-full h-10 rounded-full">
            Submit Cash Deposit to Hub Manager
          </Button>
        </form>
      )}
    </div>
  );
};
