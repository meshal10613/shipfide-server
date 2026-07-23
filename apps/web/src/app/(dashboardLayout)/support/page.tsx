
import { PhoneCall, Mail, MessageSquare, Plus } from 'lucide-react';
import { Button } from '@/components/ui/Button';

export default function SupportPage() {
  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-black text-[#404042] tracking-tight">Support & Help Desk</h1>
          <p className="text-xs text-slate-500">Contact Shipfide support or submit a ticket</p>
        </div>
        <Button variant="primary" size="sm" leftIcon={<Plus className="w-4 h-4" />}>
          New Support Ticket
        </Button>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm text-center space-y-2">
          <PhoneCall className="w-6 h-6 text-[#DD0033] mx-auto" />
          <h3 className="text-xs font-bold text-[#404042]">Call Hotline</h3>
          <p className="text-xs font-mono text-slate-500">+880 9612 000 999</p>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm text-center space-y-2">
          <Mail className="w-6 h-6 text-[#DD0033] mx-auto" />
          <h3 className="text-xs font-bold text-[#404042]">Email Support</h3>
          <p className="text-xs font-mono text-slate-500">support@shipfide.com</p>
        </div>
        <div className="bg-white p-5 rounded-3xl border border-slate-200 shadow-sm text-center space-y-2">
          <MessageSquare className="w-6 h-6 text-[#DD0033] mx-auto" />
          <h3 className="text-xs font-bold text-[#404042]">Live Chat</h3>
          <p className="text-xs text-slate-500">09:00 AM – 10:00 PM</p>
        </div>
      </div>
    </div>
  );
}

