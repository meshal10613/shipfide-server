
import { Truck, Plus } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';

export default function PickupsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Pickup Requests</h1>
          <p className="text-xs text-slate-500">Schedule hub riders to collect parcels from your warehouse</p>
        </div>
        <Button variant="primary" size="sm" leftIcon={<Plus className="w-4 h-4" />}>
          Request Pickup
        </Button>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { id: 'PK-9941', hub: 'Dhaka Central Hub', time: 'Today, 02:00 PM – 04:00 PM', count: '14 Parcels', rider: 'Salam Hossain', status: 'RIDER_ASSIGNED', badge: 'blue' },
          { id: 'PK-9910', hub: 'Dhaka Central Hub', time: 'Yesterday, 04:00 PM', count: '8 Parcels', rider: 'Karim Ahmed', status: 'COMPLETED', badge: 'emerald' },
        ].map((item) => (
          <div key={item.id} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div className="flex items-center gap-3">
              <div className="w-10 h-10 rounded-2xl bg-white border border-slate-200 flex items-center justify-center text-[#DD0033]">
                <Truck className="w-5 h-5" />
              </div>
              <div>
                <div className="flex items-center gap-2">
                  <span className="text-xs font-bold font-mono text-[#404042]">{item.id}</span>
                  <Badge variant={item.badge as any}>{item.status}</Badge>
                </div>
                <p className="text-xs text-slate-500">{item.hub} • {item.count} • Rider: {item.rider}</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

