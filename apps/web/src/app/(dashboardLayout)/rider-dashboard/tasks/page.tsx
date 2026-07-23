
import { ChevronRight } from 'lucide-react';
import { Badge } from '@/components/ui/Badge';
import Link from 'next/link';

export default function RiderTasksPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Assigned Delivery Tasks</h1>
        <p className="text-xs text-slate-500">Pickups and deliveries assigned for your shift</p>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { id: 'SF-20260722-04829', addr: 'House 42, Road 11, Block D, Banani', cod: '৳ 1,850', type: 'DELIVERY', status: 'OUT_FOR_DELIVERY' },
          { id: 'SF-20260722-09411', addr: 'House 14, Sector 3, Uttara', cod: '৳ 0 (Prepaid)', type: 'PICKUP', status: 'PENDING_PICKUP' },
        ].map((item) => (
          <div key={item.id} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div className="space-y-1">
              <div className="flex items-center gap-2">
                <span className="text-xs font-mono font-bold text-[#404042]">{item.id}</span>
                <Badge variant="red">{item.status}</Badge>
              </div>
              <p className="text-xs text-slate-500">{item.addr} • COD: <strong className="text-[#DD0033]">{item.cod}</strong></p>
            </div>

            <Link href="/rider-dashboard/deliveries/active">
              <button className="p-2.5 rounded-full bg-[#DD0033] text-white hover:bg-[#B30028] shadow-md">
                <ChevronRight className="w-4 h-4" />
              </button>
            </Link>
          </div>
        ))}
      </div>
    </div>
  );
}

