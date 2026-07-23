'use client';

import { ArrowRight, ShieldCheck, Zap } from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/Button';

export const CtaSection = () => {
  return (
    <section className="py-20 bg-linear-to-b from-white to-[#F9FCFE] relative overflow-hidden">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="bg-[#404042] text-white rounded-3xl p-8 sm:p-14 lg:p-16 relative overflow-hidden shadow-2xl space-y-8 text-center sm:text-left">
          {/* Subtle Background Glow */}
          <div className="absolute top-0 right-0 w-96 h-96 bg-[#DD0033]/20 rounded-full blur-3xl pointer-events-none" />

          <div className="max-w-2xl space-y-4 relative z-10">
            <span className="inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-[10px] font-black uppercase tracking-widest bg-[#DD0033] text-white">
              <Zap className="w-3 h-3" />
              Instant Merchant Onboarding
            </span>

            <h2 className="text-3xl sm:text-4xl lg:text-5xl font-black text-white tracking-tight leading-tight">
              Ready to Upgrade Your Logistics Infrastructure?
            </h2>

            <p className="text-sm sm:text-base text-slate-300 leading-relaxed">
              Join 12,500+ Bangladeshi merchants using Shipfide for guaranteed last-mile delivery, same-day COD cashouts, and AI fraud prevention.
            </p>
          </div>

          <div className="flex flex-col sm:flex-row items-center gap-4 relative z-10">
            <Link href="/auth/register" className="w-full sm:w-auto">
              <Button
                variant="primary"
                size="lg"
                rightIcon={<ArrowRight className="w-4 h-4" />}
                className="w-full sm:w-auto shadow-lg shadow-red-500/30"
              >
                Create Free Account
              </Button>
            </Link>
            <Link href="/support" className="w-full sm:w-auto">
              <Button
                variant="outline"
                size="lg"
                className="w-full sm:w-auto text-white border-slate-600 bg-white/10 hover:bg-white/20 hover:text-white"
              >
                Talk to Enterprise Sales
              </Button>
            </Link>
          </div>

          <div className="flex flex-wrap items-center justify-center sm:justify-start gap-6 pt-4 text-xs text-slate-400 border-t border-slate-700/60 relative z-10">
            <span className="flex items-center gap-1.5">
              <ShieldCheck className="w-4 h-4 text-emerald-400" />
              No Setup Fee or Hidden Costs
            </span>
            <span className="flex items-center gap-1.5">
              <ShieldCheck className="w-4 h-4 text-emerald-400" />
              Instant API Keys
            </span>
            <span className="flex items-center gap-1.5">
              <ShieldCheck className="w-4 h-4 text-emerald-400" />
              24/7 Priority Support
            </span>
          </div>
        </div>
      </div>
    </section>
  );
};
