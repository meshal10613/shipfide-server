
import { MapPin, Plus } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';

export default function MerchantAddressesPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Saved Addresses</h1>
          <p className="text-xs text-slate-500">Warehouse & pickup locations for rider collection</p>
        </div>
        <Button variant="primary" size="sm" leftIcon={<Plus className="w-4 h-4" />}>
          Add Address
        </Button>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { label: 'Main Warehouse', addr: 'House 14, Road 5, Sector 3, Uttara, Dhaka', phone: '+8801700000000', default: true },
          { label: 'Chattogram Store', addr: 'Plot 4, GEC Circle, Chattogram', phone: '+8801800000000', default: false },
        ].map((item, i) => (
          <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div className="flex items-start gap-3">
              <MapPin className="w-5 h-5 text-[#DD0033] shrink-0 mt-0.5" />
              <div>
                <div className="flex items-center gap-2">
                  <h4 className="text-xs font-bold text-[#404042]">{item.label}</h4>
                  {item.default && <Badge variant="emerald" size="sm">DEFAULT</Badge>}
                </div>
                <p className="text-xs text-slate-500 mt-0.5">{item.addr} • {item.phone}</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

