
import { Users, Search } from 'lucide-react';
import { Input } from '@/components/ui/Input';
import { Badge } from '@/components/ui/Badge';

export default function SuperAdminUsersPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Global User Accounts</h1>
        <p className="text-xs text-slate-500">Manage all registered platform accounts and statuses</p>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        <Input placeholder="Search user by name, email, or role..." leftIcon={<Search className="w-4 h-4 text-slate-400" />} />

        <div className="space-y-3">
          {[
            { name: 'Sri Julaekha', email: 'superadmin.shipfide@gmail.com', role: 'SUPER_ADMIN', status: 'ACTIVE', badge: 'purple' },
            { name: 'TechStore BD', email: 'techstore@example.com', role: 'MERCHANT', status: 'ACTIVE', badge: 'emerald' },
            { name: 'Karim Ahmed', email: 'karim.rider@gmail.com', role: 'RIDER', status: 'ACTIVE', badge: 'amber' },
          ].map((usr, i) => (
            <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
              <div>
                <div className="flex items-center gap-2">
                  <h4 className="text-xs font-bold text-[#404042]">{usr.name}</h4>
                  <Badge variant={usr.badge as any}>{usr.role}</Badge>
                </div>
                <span className="text-xs font-mono text-slate-500">{usr.email}</span>
              </div>
              <Badge variant="emerald">{usr.status}</Badge>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

