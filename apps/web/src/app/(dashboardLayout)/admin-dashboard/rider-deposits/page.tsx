import { Check } from 'lucide-react';
import { Button } from '@/components/ui/Button';

export default function AdminRiderDepositsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Rider COD Cash Deposit Approvals</h1>
        <p className="text-xs text-slate-500">Verify and confirm rider cash deposits collected on shift</p>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { rider: 'Karim Ahmed', id: 'RD-8492', amount: '৳ 12,450.00', time: '10 mins ago' },
          { rider: 'Salam Hossain', id: 'RD-9120', amount: '৳ 8,900.00', time: '25 mins ago' },
          { rider: 'Tanvir Rahman', id: 'RD-3041', amount: '৳ 15,300.00', time: '40 mins ago' },
        ].map((deposit) => (
          <div key={deposit.id} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div>
              <h4 className="text-sm font-bold text-[#404042]">{deposit.rider}</h4>
              <span className="text-xs font-mono text-slate-500">{deposit.id} • {deposit.time}</span>
            </div>

            <div className="flex items-center gap-3">
              <span className="text-sm font-bold text-emerald-600">{deposit.amount}</span>
              <Button variant="primary" size="sm" leftIcon={<Check className="w-4 h-4" />}>
                Approve Deposit
              </Button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
