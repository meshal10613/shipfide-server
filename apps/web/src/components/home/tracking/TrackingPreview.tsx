'use client';

import { motion } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { CheckCircle2, Truck, Package, Clock, MapPin, Search } from 'lucide-react';

export const TrackingPreview = () => {
  return (
    <section className="py-20 bg-white border-b border-slate-200/60" id="tracking-preview">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-16">
        <SectionHeading
          badge="LIVE TELEMETRY DEMO"
          title="Complete Visibility from Pickup to Receiver Handover"
          description="Every parcel status update is timestamped with sub-second accuracy and broadcast via SMS, WhatsApp, & Webhooks."
        />

        {/* Mock Tracking Interactive Card */}
        <div className="max-w-4xl mx-auto bg-[#F9FCFE] rounded-3xl border border-slate-200 p-6 sm:p-10 shadow-xl space-y-8">
          {/* Top Info Bar */}
          <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4 p-4 rounded-2xl bg-white border border-slate-200 shadow-xs">
            <div>
              <span className="text-[10px] font-bold text-slate-400 uppercase tracking-widest block">Tracking Code</span>
              <span className="text-lg font-mono font-black text-[#DD0033]">SF-20260722-04829</span>
            </div>
            <div className="flex items-center gap-3">
              <span className="px-3 py-1 rounded-full bg-emerald-50 text-emerald-600 text-xs font-bold border border-emerald-200">
                OUT FOR DELIVERY
              </span>
            </div>
          </div>

          {/* Stepper Timeline */}
          <div className="grid grid-cols-1 sm:grid-cols-4 gap-4 relative">
            {[
              { title: 'Order Booked', time: 'Yesterday, 10:15 AM', location: 'Uttara Warehouse', done: true },
              { title: 'Picked Up by Rider', time: 'Yesterday, 02:30 PM', location: 'Dhaka Central Hub', done: true },
              { title: 'In Hub Transit', time: 'Today, 06:00 AM', location: 'Linehaul Route 4', done: true },
              { title: 'Rider OTP Handover', time: 'Today, 04:30 PM (Est.)', location: 'Banani Sector 11', done: false, active: true },
            ].map((step, index) => (
              <div key={index} className="flex sm:flex-col items-center sm:items-start gap-4 sm:gap-2">
                <div className={`w-8 h-8 rounded-full flex items-center justify-center font-bold text-xs shrink-0 ${
                  step.done
                    ? 'bg-emerald-600 text-white'
                    : step.active
                    ? 'bg-[#DD0033] text-white animate-pulse'
                    : 'bg-slate-200 text-slate-500'
                }`}>
                  {step.done ? <CheckCircle2 className="w-4 h-4" /> : index + 1}
                </div>
                <div>
                  <h4 className="text-xs font-bold text-[#404042]">{step.title}</h4>
                  <p className="text-[11px] text-slate-400 font-mono mt-0.5">{step.time}</p>
                  <p className="text-[10px] text-slate-500 font-medium">{step.location}</p>
                </div>
              </div>
            ))}
          </div>

          {/* Package Details Footer */}
          <div className="grid grid-cols-2 sm:grid-cols-4 gap-4 pt-6 border-t border-slate-200 text-xs">
            <div>
              <span className="text-slate-400 block font-semibold">Receiver Name</span>
              <span className="font-bold text-[#404042]">Tanvir Ahmed</span>
            </div>
            <div>
              <span className="text-slate-400 block font-semibold">Delivery Address</span>
              <span className="font-bold text-[#404042]">Banani, Dhaka</span>
            </div>
            <div>
              <span className="text-slate-400 block font-semibold">Parcel Weight</span>
              <span className="font-bold text-[#404042]">1.4 KG</span>
            </div>
            <div>
              <span className="text-slate-400 block font-semibold">COD Amount</span>
              <span className="font-bold text-[#DD0033]">৳ 2,450.00</span>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};
