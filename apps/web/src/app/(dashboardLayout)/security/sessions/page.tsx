
import { Laptop, Smartphone, LogOut } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';

export default function SecuritySessionsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Active Login Sessions</h1>
          <p className="text-xs text-slate-500">Manage devices logged into your Shipfide account</p>
        </div>
        <Button variant="danger" size="sm" leftIcon={<LogOut className="w-3.5 h-3.5" />}>
          Revoke All Sessions
        </Button>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { device: 'Windows PC • Chrome Browser', loc: 'Dhaka, Bangladesh', ip: '103.205.142.12', active: true, icon: Laptop },
          { device: 'iPhone 15 Pro • Shipfide Mobile App', loc: 'Dhaka, Bangladesh', ip: '103.205.142.99', active: false, icon: Smartphone },
        ].map((session, i) => (
          <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-200">
            <div className="flex items-center gap-3">
              <div className="w-10 h-10 rounded-2xl bg-white border border-slate-200 flex items-center justify-center text-[#DD0033]">
                <session.icon className="w-5 h-5" />
              </div>
              <div>
                <h3 className="text-xs font-bold text-[#404042]">{session.device}</h3>
                <span className="text-[11px] font-mono text-slate-500">{session.loc} • IP: {session.ip}</span>
              </div>
            </div>

            <div className="flex items-center gap-2">
              {session.active ? (
                <Badge variant="emerald" size="sm">CURRENT DEVICE</Badge>
              ) : (
                <Button variant="ghost" size="sm">Revoke</Button>
              )}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

