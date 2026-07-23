import { Bike, Star } from 'lucide-react';
import { Badge } from '@/components/ui/Badge';

export default function RiderRosterPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Hub Rider Roster</h1>
        <p className="text-xs text-slate-500">Manage hub rider shifts, assignments, & ratings</p>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { name: 'Karim Ahmed', area: 'Banani Sector', vehicle: 'Motorcycle', rating: '4.95', status: 'ACTIVE_SHIFT' },
          { name: 'Salam Hossain', area: 'Uttara Sector', vehicle: 'Pickup Van', rating: '4.88', status: 'ACTIVE_SHIFT' },
        ].map((rider, i) => (
          <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div className="flex items-center gap-3">
              <div className="w-10 h-10 rounded-2xl bg-white border border-slate-200 flex items-center justify-center text-[#DD0033]">
                <Bike className="w-5 h-5" />
              </div>
              <div>
                <h4 className="text-xs font-bold text-[#404042]">{rider.name}</h4>
                <p className="text-xs text-slate-500">{rider.area} • {rider.vehicle}</p>
              </div>
            </div>

            <div className="flex items-center gap-3">
              <span className="text-xs font-bold text-amber-600 flex items-center gap-1">
                <Star className="w-3.5 h-3.5 fill-amber-500 text-amber-500" /> {rider.rating}
              </span>
              <Badge variant="emerald">{rider.status}</Badge>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
