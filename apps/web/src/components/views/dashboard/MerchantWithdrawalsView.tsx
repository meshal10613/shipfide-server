'use client';

import { useState } from 'react';
import { DollarSign, ArrowRight, CheckCircle2 } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export const MerchantWithdrawalsView = () => {
  const [amount, setAmount] = useState('5000');
  const [method, setMethod] = useState('BKASH');
  const [requested, setRequested] = useState(false);

  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Request Cash Withdrawal</h1>
        <p className="text-xs text-slate-500">Transfer available COD funds to Bank or Mobile Wallet</p>
      </div>

      {requested ? (
        <div className="bg-white p-8 rounded-3xl border border-slate-200 shadow-sm text-center space-y-4">
          <CheckCircle2 className="w-12 h-12 text-emerald-600 mx-auto" />
          <h2 className="text-xl font-bold text-[#404042]">Withdrawal Requested!</h2>
          <p className="text-xs text-slate-500">Amount <strong>৳ {amount}</strong> via <strong>{method}</strong>. Processing within 24 hours.</p>
          <Button onClick={() => setRequested(false)} variant="primary" className="h-10 rounded-full px-6">New Request</Button>
        </div>
      ) : (
        <form onSubmit={(e) => { e.preventDefault(); setRequested(true); }} className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
          <Input label="Withdrawal Amount (৳)" type="number" required value={amount} onChange={(e) => setAmount(e.target.value)} leftIcon={<DollarSign className="w-4 h-4" />} />

          <div className="space-y-1.5">
            <label className="text-xs font-semibold text-[#404042]">Payment Method</label>
            <select value={method} onChange={(e) => setMethod(e.target.value)} className="w-full py-3 px-4 rounded-2xl bg-white border border-slate-200 text-xs text-[#404042] focus:outline-none focus:border-[#DD0033]">
              <option value="BKASH">bKash Merchant Wallet</option>
              <option value="NAGAD">Nagad Wallet</option>
              <option value="BANK_TRANSFER">Bank Account (BEFTN/NPSB)</option>
            </select>
          </div>

          <Button type="submit" variant="primary" className="w-full h-10 rounded-full" rightIcon={<ArrowRight className="w-4 h-4" />}>
            Submit Cashout Request
          </Button>
        </form>
      )}
    </div>
  );
};
