
import { RotateCcw } from 'lucide-react';
import { Badge } from '@/components/ui/Badge';

export default function MerchantReturnsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Returned Parcels</h1>
        <p className="text-xs text-slate-500">Track undelivered shipments and return collection at hub</p>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { id: 'SF-20260720-01923', receiver: 'Tanvir Hossain', reason: 'Customer Unreachable (3 Attempts)', hub: 'Dhaka Central Hub', status: 'AWAITING_COLLECTION' },
        ].map((item) => (
          <div key={item.id} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div className="flex items-center gap-3">
              <div className="w-10 h-10 rounded-2xl bg-amber-50 text-amber-600 border border-amber-200 flex items-center justify-center">
                <RotateCcw className="w-5 h-5" />
              </div>
              <div>
                <div className="flex items-center gap-2">
                  <span className="text-xs font-mono font-bold text-[#404042]">{item.id}</span>
                  <Badge variant="amber">{item.status}</Badge>
                </div>
                <p className="text-xs text-slate-500">{item.receiver} • Reason: {item.reason}</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

