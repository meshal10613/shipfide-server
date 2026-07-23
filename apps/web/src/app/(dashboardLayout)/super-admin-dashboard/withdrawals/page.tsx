
import { Check } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';

export default function SuperAdminWithdrawalsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Platform Cashout Approvals</h1>
        <p className="text-xs text-slate-500">Approve & disburse merchant & rider cashout requests</p>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { user: 'TechStore BD (Merchant)', amount: '৳ 25,000.00', method: 'bKash Merchant', id: 'WD-001', status: 'REQUESTED' },
          { user: 'Karim Ahmed (Rider)', amount: '৳ 4,500.00', method: 'Nagad', id: 'WD-002', status: 'REQUESTED' },
        ].map((item) => (
          <div key={item.id} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div>
              <div className="flex items-center gap-2">
                <h4 className="text-xs font-bold text-[#404042]">{item.user}</h4>
                <Badge variant="amber">{item.status}</Badge>
              </div>
              <p className="text-xs text-slate-500">{item.method} • Ref: {item.id}</p>
            </div>

            <div className="flex items-center gap-3">
              <span className="text-sm font-bold text-[#DD0033]">{item.amount}</span>
              <Button variant="primary" size="sm" leftIcon={<Check className="w-4 h-4" />}>
                Approve & Pay
              </Button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

