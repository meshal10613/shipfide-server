
import { UserPlus } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';

export default function SuperAdminStaffPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Hub Admin Staff</h1>
          <p className="text-xs text-slate-500">Invite & assign administrative staff to physical hubs</p>
        </div>
        <Button variant="primary" size="sm" leftIcon={<UserPlus className="w-4 h-4" />}>
          Invite Admin Staff
        </Button>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { name: 'Rafiqul Islam', hub: 'Dhaka Central Hub', role: 'ROLE_ADMIN', email: 'rafiq.hub@shipfide.com' },
        ].map((staff, i) => (
          <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div>
              <div className="flex items-center gap-2">
                <h4 className="text-xs font-bold text-[#404042]">{staff.name}</h4>
                <Badge variant="blue">{staff.role}</Badge>
              </div>
              <p className="text-xs text-slate-500">{staff.hub} • {staff.email}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

