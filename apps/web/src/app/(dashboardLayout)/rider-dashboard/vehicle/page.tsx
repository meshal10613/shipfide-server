
import { Bike, Save } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export default function RiderVehiclePage() {
  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div>
        <h1 className="text-2xl font-black text-[#404042] tracking-tight">Vehicle Registration</h1>
        <p className="text-xs text-slate-500">Manage registered delivery vehicle specs</p>
      </div>

      <form className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-4">
        <Input label="Vehicle Type" defaultValue="Motorcycle (Two Wheeler)" leftIcon={<Bike className="w-4 h-4" />} />
        <Input label="Registration Number" defaultValue="Dhaka Metro LA-84-1920" />
        <Input label="Model / Make" defaultValue="Yamaha FZ-S V3 (150cc)" />
        <Button type="button" variant="primary" className="py-3 px-6" leftIcon={<Save className="w-4 h-4" />}>
          Update Vehicle Details
        </Button>
      </form>
    </div>
  );
}

