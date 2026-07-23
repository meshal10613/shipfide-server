
import { ShieldCheck, FileText } from 'lucide-react';
import { Badge } from '@/components/ui/Badge';

export default function RiderKycPage() {
  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Rider KYC Verification</h1>
          <p className="text-xs text-slate-500">NID card & driving license document verification</p>
        </div>
        <Badge variant="emerald" size="md">VERIFIED RIDER</Badge>
      </div>

      <div className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
        <div className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
          <div className="flex items-center gap-3">
            <FileText className="w-5 h-5 text-emerald-600" />
            <div>
              <h4 className="text-xs font-bold text-[#404042]">National ID (NID)</h4>
              <p className="text-[10px] text-slate-400">Front & Back Uploaded • Verified</p>
            </div>
          </div>
          <Badge variant="emerald">APPROVED</Badge>
        </div>

        <div className="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-100">
          <div className="flex items-center gap-3">
            <FileText className="w-5 h-5 text-emerald-600" />
            <div>
              <h4 className="text-xs font-bold text-[#404042]">Driving License</h4>
              <p className="text-[10px] text-slate-400">Valid until 2029 • Verified</p>
            </div>
          </div>
          <Badge variant="emerald">APPROVED</Badge>
        </div>
      </div>
    </div>
  );
}

