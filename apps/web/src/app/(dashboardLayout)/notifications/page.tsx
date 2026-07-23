
import { Bell, Truck, DollarSign } from 'lucide-react';
import { Badge } from '@/components/ui/Badge';

export default function NotificationsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Notifications & Alerts</h1>
          <p className="text-xs text-slate-500">Real-time updates on parcels, COD, & system</p>
        </div>
        <Badge variant="red" size="md">3 NEW</Badge>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-3">
        {[
          { title: 'COD Cash Withdrawal Approved', time: '10 mins ago', desc: 'Your withdrawal request of ৳ 15,000 to bKash has been processed.', icon: DollarSign, color: 'text-emerald-600' },
          { title: 'Parcel Delivered (SF-20260722-04829)', time: '1 hour ago', desc: 'Rider Karim Ahmed completed parcel delivery with OTP verification.', icon: Truck, color: 'text-[#DD0033]' },
          { title: 'Super Admin System Announcement', time: 'Yesterday', desc: 'Dhaka Central Hub operating hours extended during Eid holiday peak.', icon: Bell, color: 'text-purple-600' },
        ].map((item, i) => (
          <div key={i} className="flex items-start gap-4 p-4 rounded-2xl bg-slate-50 border border-slate-100 hover:border-slate-200 transition-colors">
            <div className={`w-10 h-10 rounded-2xl bg-white border border-slate-200 flex items-center justify-center shrink-0 ${item.color}`}>
              <item.icon className="w-5 h-5" />
            </div>
            <div className="flex-1">
              <div className="flex items-center justify-between">
                <h3 className="text-xs font-bold text-[#404042]">{item.title}</h3>
                <span className="text-[10px] font-mono text-slate-400">{item.time}</span>
              </div>
              <p className="text-xs text-slate-500 mt-1">{item.desc}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

