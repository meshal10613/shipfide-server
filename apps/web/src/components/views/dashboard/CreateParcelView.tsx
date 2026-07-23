'use client';

import { useState } from 'react';
import { Package, User, Phone, MapPin, DollarSign, ArrowRight, CheckCircle2 } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export const CreateParcelView = () => {
  const [created, setCreated] = useState(false);

  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Create Parcel Delivery</h1>
        <p className="text-xs text-slate-500">Book single parcel for hub pickup & delivery</p>
      </div>

      {created ? (
        <div className="bg-white p-8 rounded-3xl border border-slate-200 shadow-sm text-center space-y-4">
          <CheckCircle2 className="w-12 h-12 text-emerald-600 mx-auto" />
          <h2 className="text-xl font-bold text-[#404042]">Parcel Order Created!</h2>
          <p className="text-xs text-slate-500">Tracking Code: <strong className="text-[#DD0033] font-mono">SF-20260723-99120</strong></p>
          <Button onClick={() => setCreated(false)} variant="primary" className="h-10 rounded-full px-6">Create Another Order</Button>
        </div>
      ) : (
        <form onSubmit={(e) => { e.preventDefault(); setCreated(true); }} className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
          <h3 className="text-xs font-bold text-[#404042] uppercase tracking-wider">Receiver Details</h3>
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <Input label="Receiver Name" required placeholder="Full Name" leftIcon={<User className="w-4 h-4" />} />
            <Input label="Receiver Phone" required placeholder="+8801700000000" leftIcon={<Phone className="w-4 h-4" />} />
          </div>

          <Input label="Delivery Address" required placeholder="House/Road, Area, District" leftIcon={<MapPin className="w-4 h-4" />} />

          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <Input label="COD Amount (৳)" type="number" defaultValue="1500" leftIcon={<DollarSign className="w-4 h-4" />} />
            <Input label="Parcel Weight (KG)" type="number" defaultValue="1" leftIcon={<Package className="w-4 h-4" />} />
          </div>

          <Button type="submit" variant="primary" className="w-full h-10 rounded-full" rightIcon={<ArrowRight className="w-4 h-4" />}>
            Confirm Parcel Order
          </Button>
        </form>
      )}
    </div>
  );
};
