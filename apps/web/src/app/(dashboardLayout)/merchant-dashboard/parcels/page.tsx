
import { Package, Plus, Search, Filter } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';
import { Badge } from '@/components/ui/Badge';
import Link from 'next/link';

export default function MerchantParcelsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Merchant Deliveries</h1>
          <p className="text-xs text-slate-500">Track & manage all outgoing parcel orders</p>
        </div>
        <Link href="/merchant-dashboard/parcels/create">
          <Button variant="primary" size="sm" leftIcon={<Plus className="w-4 h-4" />}>
            Create Order
          </Button>
        </Link>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        <div className="flex items-center gap-3">
          <Input placeholder="Search by tracking code or phone..." leftIcon={<Search className="w-4 h-4 text-slate-400" />} />
          <Button variant="outline" size="md" leftIcon={<Filter className="w-4 h-4" />}>
            Filter
          </Button>
        </div>

        <div className="space-y-3">
          {[
            { id: 'SF-20260722-04829', receiver: 'Naimur Rahman', area: 'Banani, Dhaka', cod: '৳ 2,450.00', status: 'OUT_FOR_DELIVERY', badge: 'red' },
            { id: 'SF-20260722-03912', receiver: 'Fatima Zohra', area: 'Agrabad, Chattogram', cod: '৳ 1,800.00', status: 'IN_TRANSIT', badge: 'blue' },
            { id: 'SF-20260721-09410', receiver: 'Arif Hossain', area: 'Dhanmondi, Dhaka', cod: '৳ 3,200.00', status: 'DELIVERED', badge: 'emerald' },
          ].map((item) => (
            <div key={item.id} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100 hover:border-slate-200 transition-all">
              <div className="space-y-1">
                <div className="flex items-center gap-2">
                  <span className="text-xs font-mono font-bold text-[#404042]">{item.id}</span>
                  <Badge variant={item.badge as any}>{item.status}</Badge>
                </div>
                <p className="text-xs text-slate-500">{item.receiver} • {item.area}</p>
              </div>

              <div className="text-right">
                <span className="text-sm font-black text-[#DD0033] block">{item.cod}</span>
                <span className="text-[10px] text-slate-400 font-medium">COD Amount</span>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

