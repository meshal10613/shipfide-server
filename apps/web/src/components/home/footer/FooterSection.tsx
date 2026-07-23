'use client';

import Image from 'next/image';
import Link from 'next/link';
import { Send, MapPin, Phone, Mail } from 'lucide-react';
import { Button } from '@/components/ui/Button';

export const FooterSection = () => {
  return (
    <footer className="bg-white border-t border-slate-200/80 pt-16 pb-12 text-[#404042]">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
          {/* Brand & Logo Column */}
          <div className="lg:col-span-2 space-y-4">
            <Link href="/" className="inline-block">
              <Image
                src="/logo.png"
                alt="Shipfide Logo"
                width={160}
                height={48}
                className="h-10 w-auto object-contain"
              />
            </Link>

            <p className="text-xs text-slate-500 max-w-sm leading-relaxed">
              Shipfide is Bangladesh’s premier enterprise logistics & courier operating system, providing 64-district last mile delivery, same-day COD cashouts, and AI fleet telemetry.
            </p>

            <div className="space-y-1.5 text-xs text-slate-600 font-medium pt-2">
              <p className="flex items-center gap-2">
                <MapPin className="w-3.5 h-3.5 text-[#DD0033]" />
                Dhaka Central Hub, Sector 3, Uttara, Dhaka-1230
              </p>
              <p className="flex items-center gap-2">
                <Phone className="w-3.5 h-3.5 text-[#DD0033]" />
                +880 9612 000 999 (Hotline)
              </p>
              <p className="flex items-center gap-2">
                <Mail className="w-3.5 h-3.5 text-[#DD0033]" />
                support@shipfide.com
              </p>
            </div>
          </div>

          {/* Column 2: Product & Solutions */}
          <div className="space-y-3">
            <h4 className="text-xs font-bold uppercase tracking-wider text-[#404042]">Product Solutions</h4>
            <ul className="space-y-2 text-xs text-slate-500 font-medium">
              <li><Link href="/tracking" className="hover:text-[#DD0033]">Live Parcel Telemetry</Link></li>
              <li><Link href="/pricing" className="hover:text-[#DD0033]">Rate Calculator</Link></li>
              <li><Link href="/receiver-fraud-check" className="hover:text-[#DD0033]">Receiver Fraud Shield</Link></li>
              <li><Link href="/guest-delivery" className="hover:text-[#DD0033]">Guest Parcel Booking</Link></li>
              <li><Link href="/merchant-dashboard/parcels/bulk-upload" className="hover:text-[#DD0033]">Bulk CSV Import</Link></li>
            </ul>
          </div>

          {/* Column 3: Dashboards */}
          <div className="space-y-3">
            <h4 className="text-xs font-bold uppercase tracking-wider text-[#404042]">Dashboards</h4>
            <ul className="space-y-2 text-xs text-slate-500 font-medium">
              <li><Link href="/merchant-dashboard" className="hover:text-[#DD0033]">Merchant Portal</Link></li>
              <li><Link href="/rider-dashboard" className="hover:text-[#DD0033]">Delivery Rider Portal</Link></li>
              <li><Link href="/admin-dashboard" className="hover:text-[#DD0033]">Hub Staff Desk</Link></li>
              <li><Link href="/super-admin-dashboard" className="hover:text-[#DD0033]">Super Admin Portal</Link></li>
              <li><Link href="/support" className="hover:text-[#DD0033]">Help & Support Desk</Link></li>
            </ul>
          </div>

          {/* Column 4: Newsletter */}
          <div className="space-y-3">
            <h4 className="text-xs font-bold uppercase tracking-wider text-[#404042]">Stay Updated</h4>
            <p className="text-xs text-slate-500">Subscribe for logistics industry insights & release notes.</p>
            <form onSubmit={(e) => e.preventDefault()} className="space-y-2">
              <input
                type="email"
                placeholder="Enter work email"
                className="w-full px-3 py-2 rounded-full bg-slate-50 border border-slate-200 text-xs text-[#404042] focus:outline-none focus:border-[#DD0033]"
              />
              <Button
                type="submit"
                variant="primary"
                size="sm"
                rightIcon={<Send className="w-3 h-3" />}
                className="w-full"
              >
                Subscribe
              </Button>
            </form>
          </div>
        </div>

        {/* Bottom Copyright Bar */}
        <div className="pt-8 border-t border-slate-200/60 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-slate-500">
          <p className="font-bold text-[#404042]">
            Shipfide Logistics Platform • © 2026 Shipfide. All rights reserved.
          </p>
          <div className="flex items-center gap-4 text-[11px]">
            <Link href="/privacy" className="hover:text-[#DD0033]">Privacy Policy</Link>
            <span>•</span>
            <Link href="/terms" className="hover:text-[#DD0033]">Terms of Service</Link>
            <span>•</span>
            <Link href="/sla" className="hover:text-[#DD0033]">SLA Guarantee</Link>
          </div>
        </div>
      </div>
    </footer>
  );
};
