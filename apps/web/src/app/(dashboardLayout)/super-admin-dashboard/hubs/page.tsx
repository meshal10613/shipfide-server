
import { Building, Plus } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';

export default function SuperAdminHubsPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Hub Network Management</h1>
          <p className="text-xs text-slate-500">Create & configure physical hubs across Bangladesh</p>
        </div>
        <Button variant="primary" size="sm" leftIcon={<Plus className="w-4 h-4" />}>
          Create New Hub
        </Button>
      </div>

      <div className="bg-white rounded-3xl border border-slate-200 p-6 shadow-sm space-y-4">
        {[
          { name: 'Dhaka Central Hub', dist: 'Dhaka', div: 'Dhaka', staff: '12 Staff', status: 'ACTIVE' },
          { name: 'Chattogram Hub', dist: 'Chattogram', div: 'Chattogram', staff: '8 Staff', status: 'ACTIVE' },
          { name: 'Sylhet Divisional Hub', dist: 'Sylhet', div: 'Sylhet', staff: '5 Staff', status: 'ACTIVE' },
        ].map((hub, i) => (
          <div key={i} className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
            <div className="flex items-center gap-3">
              <div className="w-10 h-10 rounded-2xl bg-white border border-slate-200 flex items-center justify-center text-[#DD0033]">
                <Building className="w-5 h-5" />
              </div>
              <div>
                <h4 className="text-xs font-bold text-[#404042]">{hub.name}</h4>
                <p className="text-xs text-slate-500">{hub.dist} District, {hub.div} Division • {hub.staff}</p>
              </div>
            </div>

            <Badge variant="emerald">{hub.status}</Badge>
          </div>
        ))}
      </div>
    </div>
  );
}

