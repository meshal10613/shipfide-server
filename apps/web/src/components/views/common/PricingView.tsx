'use client';

import { useState } from 'react';
import { ArrowRight } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';
import { Badge } from '@/components/ui/Badge';

export const PricingView = () => {
  const [weight, setWeight] = useState('1');
  const [zone, setZone] = useState('INSIDE_DHAKA');
  const [cod, setCod] = useState('1500');

  const basePrice = zone === 'INSIDE_DHAKA' ? 60 : zone === 'DHAKA_SUBURB' ? 80 : 130;
  const extraWeightPrice = Math.max(0, (parseFloat(weight) || 1) - 1) * (zone === 'OUTSIDE_DHAKA' ? 25 : 15);
  const totalRate = basePrice + extraWeightPrice;

  return (
    <div className="max-w-4xl mx-auto px-4 sm:px-6 py-12 w-full space-y-8">
      <div className="text-center space-y-3">
        <Badge variant="blue">TRANSPARENT RATES</Badge>
        <h1 className="text-3xl sm:text-4xl font-black text-[#404042] tracking-tight">
          Delivery Rate & Pricing Calculator
        </h1>
        <p className="text-sm text-slate-500 max-w-lg mx-auto">
          Calculate instant shipping charges based on destination zone and parcel weight
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-12 gap-8 items-start">
        <form className="md:col-span-7 bg-white p-6 sm:p-8 rounded-3xl border border-slate-200 shadow-sm space-y-4">
          <div className="space-y-1.5">
            <label className="text-xs font-semibold text-[#404042]">Destination Delivery Zone</label>
            <select
              value={zone}
              onChange={(e) => setZone(e.target.value)}
              className="w-full py-3 px-4 rounded-2xl bg-white border border-slate-200 text-xs text-[#404042] focus:outline-none focus:border-[#DD0033]"
            >
              <option value="INSIDE_DHAKA">Inside Dhaka City (24 Hours)</option>
              <option value="DHAKA_SUBURB">Dhaka Suburbs (Savar/Gazipur/Narayanganj)</option>
              <option value="OUTSIDE_DHAKA">Outside Dhaka District (All 64 Districts)</option>
            </select>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <Input label="Parcel Weight (KG)" type="number" min="0.5" step="0.5" value={weight} onChange={(e) => setWeight(e.target.value)} />
            <Input label="COD Amount (৳)" type="number" value={cod} onChange={(e) => setCod(e.target.value)} />
          </div>
        </form>

        <div className="md:col-span-5 bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-6">
          <div className="space-y-1">
            <span className="text-xs font-semibold text-slate-400">Estimated Delivery Fee</span>
            <div className="text-4xl font-black text-[#DD0033]">৳ {totalRate.toFixed(2)}</div>
          </div>

          <div className="space-y-2 text-xs text-slate-600 border-t border-slate-100 pt-4">
            <div className="flex justify-between">
              <span>Base Charge (Up to 1 KG):</span>
              <span className="font-bold">৳ {basePrice}</span>
            </div>
            <div className="flex justify-between">
              <span>Extra Weight Fee:</span>
              <span className="font-bold">৳ {extraWeightPrice}</span>
            </div>
            <div className="flex justify-between pt-2 border-t border-slate-100 font-bold text-[#404042]">
              <span>Total Delivery Charge:</span>
              <span className="text-[#DD0033]">৳ {totalRate.toFixed(2)}</span>
            </div>
          </div>

          <Button variant="primary" className="w-full h-10 rounded-full" rightIcon={<ArrowRight className="w-4 h-4" />}>
            Create Shipment Order
          </Button>
        </div>
      </div>
    </div>
  );
};
