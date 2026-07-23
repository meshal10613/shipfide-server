'use client';

import { useState } from 'react';
import { motion } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { SERVICES_DATA } from '@/constants/landing';
import { Check, ArrowRight, ShieldCheck } from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/Button';

export const ServicesSection = () => {
  const [activeTab, setActiveTab] = useState(SERVICES_DATA[0].id);

  const currentService = SERVICES_DATA.find((s) => s.id === activeTab) || SERVICES_DATA[0];

  return (
    <section className="py-20 bg-[#F9FCFE] border-b border-slate-200/60" id="services">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <SectionHeading
          badge="LOGISTICS SOLUTIONS"
          title="Tailored Shipping Services for Every Business Model"
          description="From small e-commerce parcel pickups to nationwide heavy freight and warehouse fulfillment."
        />

        {/* Tab Navigation Controls */}
        <div className="flex items-center justify-center gap-2 flex-wrap max-w-4xl mx-auto">
          {SERVICES_DATA.map((service) => {
            const isActive = service.id === activeTab;
            return (
              <Button
                key={service.id}
                variant={isActive ? 'primary' : 'outline'}
                size="sm"
                onClick={() => setActiveTab(service.id)}
              >
                {service.title}
              </Button>
            );
          })}
        </div>

        {/* Tab Content Display Card */}
        <motion.div
          key={currentService.id}
          initial={{ opacity: 0, y: 15 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.4 }}
          className="bg-white p-8 sm:p-12 rounded-3xl border border-slate-200 shadow-xl max-w-5xl mx-auto grid grid-cols-1 md:grid-cols-12 gap-8 items-center"
        >
          <div className="md:col-span-7 space-y-6">
            <span className="px-3 py-1 rounded-full text-[11px] font-bold text-[#DD0033] bg-[#DD0033]/10 border border-[#DD0033]/20">
              {currentService.subtitle}
            </span>

            <h3 className="text-2xl sm:text-3xl font-black text-[#404042]">
              {currentService.title}
            </h3>

            <p className="text-sm text-slate-600 leading-relaxed">
              {currentService.description}
            </p>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-3 pt-2">
              {currentService.features.map((feat, i) => (
                <div key={i} className="flex items-center gap-2 text-xs font-bold text-[#404042]">
                  <div className="w-4 h-4 rounded-full bg-emerald-100 text-emerald-600 flex items-center justify-center shrink-0">
                    <Check className="w-2.5 h-2.5" />
                  </div>
                  <span>{feat}</span>
                </div>
              ))}
            </div>

            <div className="pt-4 flex items-center gap-4">
              <Link href="/guest-delivery">
                <Button variant="primary" rightIcon={<ArrowRight className="w-3.5 h-3.5" />}>
                  Book This Service
                </Button>
              </Link>
            </div>
          </div>

          <div className="md:col-span-5 bg-[#F9FCFE] p-6 rounded-3xl border border-slate-200 space-y-4 text-center">
            <div className="w-16 h-16 rounded-3xl bg-[#DD0033]/10 text-[#DD0033] flex items-center justify-center mx-auto shadow-inner">
              <ShieldCheck className="w-8 h-8" />
            </div>
            <div>
              <span className="text-xs text-slate-400 font-semibold block">Service Guarantee</span>
              <span className="text-xl font-black text-[#404042] mt-1 block">{currentService.metrics}</span>
            </div>
            <p className="text-[11px] text-slate-500">
              Backed by Shipfide Service Level Agreement & 24/7 dedicated support desk.
            </p>
          </div>
        </motion.div>
      </div>
    </section>
  );
};
