import { ShieldAlert } from 'lucide-react';
import { Badge } from '@/components/ui/Badge';

export default function AdminFraudAlertsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Fraud & Risk Alerts</h1>
        <p className="text-xs text-slate-500">Flagged receiver phone numbers & COD block enforcement</p>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { phone: '+8801911998877', reason: 'High Cancellation / Refused Parcel Rate', risk: 'HIGH_RISK', badge: 'red' },
          { phone: '+8801700112233', reason: 'Fake Address Reported by 2 Riders', risk: 'BLACKLISTED', badge: 'red' },
        ].map((item, i) => (
          <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div className="flex items-center gap-3">
              <ShieldAlert className="w-5 h-5 text-[#DD0033]" />
              <div>
                <h4 className="text-xs font-mono font-bold text-[#404042]">{item.phone}</h4>
                <p className="text-xs text-slate-500">{item.reason}</p>
              </div>
            </div>
            <Badge variant="red">{item.risk}</Badge>
          </div>
        ))}
      </div>
    </div>
  );
}
