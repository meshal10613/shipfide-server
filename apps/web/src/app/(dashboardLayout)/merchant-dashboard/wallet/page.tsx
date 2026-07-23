
import { Wallet, ArrowUpRight, ArrowDownLeft, ChevronRight } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import Link from 'next/link';

export default function MerchantWalletPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Merchant Wallet</h1>
          <p className="text-xs text-slate-500">Available COD balance & earnings breakdown</p>
        </div>
        <Link href="/merchant-dashboard/withdrawals">
          <Button variant="primary" size="sm" rightIcon={<ChevronRight className="w-4 h-4" />}>
            Withdraw Cash
          </Button>
        </Link>
      </div>

      <div className="p-6 rounded-3xl bg-gradient-to-r from-white via-slate-50 to-rose-50/50 border border-slate-200 shadow-sm space-y-3">
        <span className="text-xs font-mono text-slate-500 uppercase tracking-wider block">Available COD Balance</span>
        <span className="text-3xl font-black text-[#404042]">৳ 45,820.00</span>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        <h3 className="text-sm font-bold text-[#404042]">Recent Transactions</h3>
        {[
          { title: 'COD Collected (SF-20260722-04829)', time: 'Today, 09:30 AM', amount: '+৳ 2,450.00', type: 'in' },
          { title: 'Withdrawal to bKash', time: 'Yesterday, 04:00 PM', amount: '-৳ 10,000.00', type: 'out' },
        ].map((tx, i) => (
          <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div className="flex items-center gap-3">
              <div className={`w-10 h-10 rounded-2xl bg-white border border-slate-200 flex items-center justify-center ${tx.type === 'in' ? 'text-emerald-600' : 'text-[#DD0033]'}`}>
                {tx.type === 'in' ? <ArrowDownLeft className="w-5 h-5" /> : <ArrowUpRight className="w-5 h-5" />}
              </div>
              <div>
                <h4 className="text-xs font-bold text-[#404042]">{tx.title}</h4>
                <span className="text-[10px] font-mono text-slate-400">{tx.time}</span>
              </div>
            </div>

            <span className={`text-sm font-bold ${tx.type === 'in' ? 'text-emerald-600' : 'text-[#DD0033]'}`}>
              {tx.amount}
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}

