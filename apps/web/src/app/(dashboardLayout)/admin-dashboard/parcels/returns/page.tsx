import { Check } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';

export default function AdminReturnsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Hub Return Processing</h1>
        <p className="text-xs text-slate-500">Handover returned parcels to merchant or schedule return transit</p>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { id: 'SF-20260720-01923', merchant: 'TechStore BD', receiver: 'Tanvir Hossain', status: 'RETURNED_TO_HUB' },
        ].map((item) => (
          <div key={item.id} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div>
              <div className="flex items-center gap-2">
                <span className="text-xs font-mono font-bold text-[#404042]">{item.id}</span>
                <Badge variant="amber">{item.status}</Badge>
              </div>
              <p className="text-xs text-slate-500">Merchant: {item.merchant} • Receiver: {item.receiver}</p>
            </div>
            <Button variant="primary" size="sm" leftIcon={<Check className="w-4 h-4" />}>
              Handover to Merchant
            </Button>
          </div>
        ))}
      </div>
    </div>
  );
}
