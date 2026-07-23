'use client';

import { useState } from 'react';
import { KeyRound, CheckCircle2 } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export const ConfirmOtpView = () => {
  const [otp, setOtp] = useState('');
  const [confirmed, setConfirmed] = useState(false);

  return (
    <div className="max-w-md mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Confirm OTP Handover</h1>
        <p className="text-xs text-slate-500">Ask receiver for 6-digit verification OTP code</p>
      </div>

      {confirmed ? (
        <div className="bg-white p-8 rounded-3xl border border-slate-200 shadow-sm text-center space-y-4">
          <CheckCircle2 className="w-12 h-12 text-emerald-600 mx-auto" />
          <h2 className="text-xl font-bold text-[#404042]">Parcel Handover Verified!</h2>
          <p className="text-xs text-slate-500">COD <strong>৳ 1,850.00</strong> collected into rider cash wallet.</p>
        </div>
      ) : (
        <form onSubmit={(e) => { e.preventDefault(); setConfirmed(true); }} className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
          <Input
            label="Receiver 6-Digit OTP"
            type="text"
            required
            maxLength={6}
            value={otp}
            onChange={(e) => setOtp(e.target.value)}
            placeholder="123456"
            leftIcon={<KeyRound className="w-4 h-4" />}
          />
          <Button type="submit" variant="primary" className="w-full h-10 rounded-full">
            Verify & Complete Delivery
          </Button>
        </form>
      )}
    </div>
  );
};
