'use client';

import { useState } from 'react';
import { QrCode, CheckCircle2 } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export const ReceiveParcelsView = () => {
  const [code, setCode] = useState('');
  const [scanned, setScanned] = useState(false);

  return (
    <div className="max-w-md mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Receive Parcels at Hub</h1>
        <p className="text-xs text-slate-500">Scan incoming barcode / QR code from merchant pickup</p>
      </div>

      {scanned ? (
        <div className="bg-white p-8 rounded-3xl border border-slate-200 shadow-sm text-center space-y-4">
          <CheckCircle2 className="w-12 h-12 text-emerald-600 mx-auto" />
          <h2 className="text-xl font-bold text-[#404042]">Parcel Received at Hub!</h2>
          <p className="text-xs text-slate-500">Tracking Code: <strong className="text-[#DD0033] font-mono">{code.toUpperCase()}</strong></p>
          <Button onClick={() => setScanned(false)} variant="primary" className="h-10 rounded-full px-6">Scan Next Parcel</Button>
        </div>
      ) : (
        <form onSubmit={(e) => { e.preventDefault(); setScanned(true); }} className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
          <Input label="Parcel Tracking ID / QR Code" required value={code} onChange={(e) => setCode(e.target.value)} placeholder="SF-20260722-XXXXX" leftIcon={<QrCode className="w-4 h-4" />} />
          <Button type="submit" variant="primary" className="w-full h-10 rounded-full">
            Confirm Received at Hub
          </Button>
        </form>
      )}
    </div>
  );
};
