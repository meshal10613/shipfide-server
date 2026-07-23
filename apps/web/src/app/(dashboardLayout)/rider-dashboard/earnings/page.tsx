
import { DollarSign, Star } from 'lucide-react';
import { Badge } from '@/components/ui/Badge';

export default function RiderEarningsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Rider Earnings & Badges</h1>
          <p className="text-xs text-slate-500">Commission income & rider performance score</p>
        </div>
        <Badge variant="amber" size="md">TOP RIDER 🌟</Badge>
      </div>

      <div className="grid grid-cols-2 sm:grid-cols-3 gap-4">
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Today's Commission</span>
          <div className="text-2xl font-black text-[#DD0033]">৳ 1,450</div>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Monthly Earnings</span>
          <div className="text-2xl font-black text-[#404042]">৳ 34,200</div>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Rating Score</span>
          <div className="text-2xl font-black text-amber-600 flex items-center gap-1">
            <Star className="w-5 h-5 fill-amber-500 text-amber-500" /> 4.95
          </div>
        </div>
      </div>
    </div>
  );
}

