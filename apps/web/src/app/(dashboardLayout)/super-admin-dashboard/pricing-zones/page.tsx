
import { Plus } from 'lucide-react';
import { Button } from '@/components/ui/Button';

export default function SuperAdminPricingZonesPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Pricing & Delivery Zones</h1>
          <p className="text-xs text-slate-500">Configure base rates & weight multipliers across 64 districts</p>
        </div>
        <Button variant="primary" size="sm" leftIcon={<Plus className="w-4 h-4" />}>
          New Pricing Rule
        </Button>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { zone: 'INSIDE_DHAKA', base: '৳ 60.00', extraPerKg: '৳ 15.00', time: '24 Hours' },
          { zone: 'DHAKA_SUBURB', base: '৳ 80.00', extraPerKg: '৳ 15.00', time: '24–48 Hours' },
          { zone: 'OUTSIDE_DHAKA_DISTRICT', base: '৳ 130.00', extraPerKg: '৳ 25.00', time: '48–72 Hours' },
        ].map((rule, i) => (
          <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div>
              <h4 className="text-xs font-bold font-mono text-[#404042]">{rule.zone}</h4>
              <p className="text-xs text-slate-500">Est. Time: {rule.time} • Extra KG: {rule.extraPerKg}</p>
            </div>
            <span className="text-sm font-black text-[#DD0033]">{rule.base} Base</span>
          </div>
        ))}
      </div>
    </div>
  );
}

