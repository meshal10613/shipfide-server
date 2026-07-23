
import { Globe, Building, DollarSign, Users } from 'lucide-react';
import { Badge } from '@/components/ui/Badge';

export default function SuperAdminOverviewPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Super Admin System Overview</h1>
          <p className="text-xs text-slate-500">Global revenue, hub network, & platform control</p>
        </div>
        <Badge variant="purple" size="md">SUPER ADMIN</Badge>
      </div>

      <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Platform Revenue</span>
          <div className="text-2xl font-black text-purple-700">৳ 1.48M</div>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Active Hubs</span>
          <div className="text-2xl font-black text-[#404042]">64</div>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Total Merchants</span>
          <div className="text-2xl font-black text-emerald-600">3,420</div>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm space-y-1">
          <span className="text-xs font-semibold text-slate-500">Total Riders</span>
          <div className="text-2xl font-black text-amber-600">1,240</div>
        </div>
      </div>
    </div>
  );
}

