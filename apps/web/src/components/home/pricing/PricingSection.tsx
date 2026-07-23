'use client';

import { useState } from 'react';
import { motion } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { PRICING_PLANS } from '@/constants/landing';
import { Check, ArrowRight, Sparkles } from 'lucide-react';
import Link from 'next/link';

export const PricingSection = () => {
  const [isYearly, setIsYearly] = useState(true);

  return (
    <section className="py-20 bg-white border-b border-slate-200/60" id="pricing">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <SectionHeading
          badge="TRANSPARENT PRICING"
          title="Simple & Predictable Shipping Rates for Sellers of All Sizes"
          description="No hidden fuel surcharges, volume penalties, or secret fees. Choose the tier that matches your monthly shipping volume."
        />

        {/* Monthly / Yearly Billing Toggle */}
        <div className="flex items-center justify-center gap-3">
          <span className={`text-xs font-bold ${!isYearly ? 'text-[#404042]' : 'text-slate-400'}`}>
            Monthly Billing
          </span>
          <button
            onClick={() => setIsYearly(!isYearly)}
            className="w-14 h-8 rounded-full bg-[#404042] p-1 relative transition-colors cursor-pointer"
            aria-label="Toggle billing frequency"
          >
            <div
              className={`w-6 h-6 rounded-full bg-[#DD0033] transition-transform ${
                isYearly ? 'translate-x-6' : 'translate-x-0'
              }`}
            />
          </button>
          <div className="flex items-center gap-1.5">
            <span className={`text-xs font-bold ${isYearly ? 'text-[#404042]' : 'text-slate-400'}`}>
              Annual Billing
            </span>
            <span className="px-2 py-0.5 rounded-full text-[10px] font-black bg-emerald-100 text-emerald-700 uppercase">
              Save 20%
            </span>
          </div>
        </div>

        {/* Pricing Cards Grid */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 items-stretch pt-4">
          {PRICING_PLANS.map((plan, index) => {
            const price = isYearly ? plan.priceYearly : plan.priceMonthly;

            return (
              <motion.div
                key={plan.id}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ duration: 0.5, delay: index * 0.1 }}
                className={`rounded-3xl p-8 flex flex-col justify-between space-y-8 relative transition-all ${
                  plan.isPopular
                    ? 'bg-white border-2 border-[#DD0033] shadow-2xl shadow-red-500/10 scale-105 z-10'
                    : 'bg-[#F9FCFE] border border-slate-200 shadow-sm hover:border-slate-300'
                }`}
              >
                {plan.isPopular && (
                  <div className="absolute -top-4 left-1/2 -translate-x-1/2 px-4 py-1 rounded-full bg-[#DD0033] text-white text-[10px] font-black uppercase tracking-wider shadow-md flex items-center gap-1">
                    <Sparkles className="w-3 h-3" />
                    Most Popular Choice
                  </div>
                )}

                <div className="space-y-6">
                  <div>
                    <h3 className="text-xl font-black text-[#404042]">{plan.name}</h3>
                    <p className="text-xs text-slate-500 mt-1 leading-relaxed">{plan.description}</p>
                  </div>

                  <div className="flex items-baseline gap-1">
                    <span className="text-sm font-bold text-slate-400">৳</span>
                    <span className="text-4xl sm:text-5xl font-black text-[#404042]">
                      {price}
                    </span>
                    <span className="text-xs font-semibold text-slate-400">/ base rate</span>
                  </div>

                  <div className="space-y-3 pt-4 border-t border-slate-200/80">
                    <span className="text-[11px] font-bold uppercase tracking-wider text-slate-400 block">Features Included</span>
                    {plan.features.map((feat, i) => (
                      <div key={i} className="flex items-start gap-2.5 text-xs text-[#404042] font-medium">
                        <Check className="w-4 h-4 text-[#DD0033] shrink-0 mt-0.5" />
                        <span>{feat}</span>
                      </div>
                    ))}
                  </div>
                </div>

                <Link href="/auth/register">
                  <button
                    className={`w-full py-4 rounded-2xl font-bold text-xs flex items-center justify-center gap-2 transition-all cursor-pointer ${
                      plan.isPopular
                        ? 'bg-[#DD0033] hover:bg-[#B30028] text-white shadow-lg shadow-red-500/25'
                        : 'bg-white hover:bg-slate-50 text-[#404042] border border-slate-200'
                    }`}
                  >
                    <span>{plan.ctaText}</span>
                    <ArrowRight className="w-4 h-4" />
                  </button>
                </Link>
              </motion.div>
            );
          })}
        </div>
      </div>
    </section>
  );
};
