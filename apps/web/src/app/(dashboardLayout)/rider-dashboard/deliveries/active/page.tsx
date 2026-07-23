
import { MapPin, PhoneCall, CheckCircle2, QrCode } from 'lucide-react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';
import Link from 'next/link';

export default function ActiveDeliveryPage() {
  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Active Delivery Task</h1>
          <p className="text-xs font-mono text-[#DD0033]">SF-20260722-04829</p>
        </div>
        <Badge variant="amber">OUT FOR DELIVERY</Badge>
      </div>

      <div className="bg-white p-6 rounded-3xl border border-slate-200 shadow-sm space-y-6">
        <div className="space-y-3">
          <div className="flex items-start gap-3">
            <MapPin className="w-5 h-5 text-[#DD0033] shrink-0 mt-0.5" />
            <div>
              <span className="text-xs font-semibold text-slate-500 block">Receiver Address</span>
              <p className="text-sm font-bold text-[#404042]">House 42, Road 11, Block D, Banani, Dhaka</p>
            </div>
          </div>

          <div className="flex items-center justify-between pt-3 border-t border-slate-100">
            <div>
              <span className="text-xs font-semibold text-slate-500 block">Receiver Phone</span>
              <span className="text-sm font-mono font-bold text-[#404042]">+8801711223344</span>
            </div>
            <button className="flex items-center gap-1.5 px-4 py-2 rounded-full bg-slate-50 border border-slate-200 text-xs font-bold text-[#404042] hover:text-[#DD0033]">
              <PhoneCall className="w-4 h-4 text-[#DD0033]" />
              <span>Call Customer</span>
            </button>
          </div>
        </div>

        <div className="grid grid-cols-2 gap-3 pt-3 border-t border-slate-100">
          <Link href="/rider-dashboard/deliveries/confirm-otp" className="w-full">
            <Button variant="primary" className="w-full py-3.5" leftIcon={<CheckCircle2 className="w-4 h-4" />}>
              Confirm Delivery
            </Button>
          </Link>
          <Button variant="outline" className="w-full py-3.5" leftIcon={<QrCode className="w-4 h-4" />}>
            Scan QR Code
          </Button>
        </div>
      </div>
    </div>
  );
}

