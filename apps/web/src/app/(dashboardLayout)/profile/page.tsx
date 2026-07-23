
import { User, Mail, Phone, Shield, Camera, Save } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';
import { Badge } from '@/components/ui/Badge';

export default function ProfilePage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Account Profile</h1>
          <p className="text-xs text-slate-500">Manage personal details and credentials</p>
        </div>
        <Badge variant="red" size="md">MERCHANT</Badge>
      </div>

      <div className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-6">
        <div className="flex items-center gap-4 pb-6 border-b border-slate-100">
          <div className="relative">
            <div className="w-20 h-20 rounded-full bg-gradient-to-tr from-[#DD0033] to-rose-400 p-0.5 shadow-md flex items-center justify-center">
              <div className="w-full h-full rounded-full bg-white flex items-center justify-center font-black text-2xl text-[#DD0033]">
                SJ
              </div>
            </div>
            <button className="absolute bottom-0 right-0 p-1.5 rounded-full bg-[#DD0033] text-white shadow-md hover:bg-[#B30028]">
              <Camera className="w-3.5 h-3.5" />
            </button>
          </div>
          <div>
            <h2 className="text-lg font-bold text-[#404042]">Sri Julaekha</h2>
            <span className="text-xs text-slate-500 font-mono">superadmin.shipfide@gmail.com</span>
            <div className="mt-1">
              <Badge variant="emerald" size="sm">ACTIVE ACCOUNT</Badge>
            </div>
          </div>
        </div>

        <form className="space-y-4">
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <Input label="Full Name" defaultValue="Sri Julaekha" leftIcon={<User className="w-4 h-4" />} />
            <Input label="Email Address" defaultValue="superadmin.shipfide@gmail.com" disabled leftIcon={<Mail className="w-4 h-4" />} />
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <Input label="Phone Number" defaultValue="+8801700000000" leftIcon={<Phone className="w-4 h-4" />} />
            <Input label="Account Role" defaultValue="Merchant / Partner" disabled leftIcon={<Shield className="w-4 h-4" />} />
          </div>

          <Button type="button" variant="primary" className="py-3 px-6" leftIcon={<Save className="w-4 h-4" />}>
            Save Profile Changes
          </Button>
        </form>
      </div>
    </div>
  );
}

